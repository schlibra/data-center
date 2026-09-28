<script setup>
import { NButton, NFlex, NSwitch, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import { computed, h, onMounted, ref } from 'vue'
import {
  createAdminGroup, deleteAdminGroup,
  getAdminGroupList,
  getAdminPermissionList,
  getUserInfo,
  updateAdminGroup,
} from '@/api'
import { useAdminGroupStore, useAdminPermissionStore, useUserStore } from '@/stores'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const user = useUserStore()
const adminGroup = useAdminGroupStore()
const adminPermission = useAdminPermissionStore()

const showModal = ref(false)
const createGroup = ref(true)
const groupId = ref('')
const groupName = ref('')
const groupAdmin = ref(false)
const permission = ref([])

const permissionTree = computed(() =>
  adminPermission.permissions.map((item) => {
    return {
      label: item.name,
      key: item.id,
      disabled: true,
      children: item.children.map((item) => {
        return {
          label: item.name,
          key: item.id,
        }
      }),
    }
  }),
)
const dialogError = (content) => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}
const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 60,
  },
  {
    title: '用户组名',
    key: 'name',
    width: 100,
  },
  {
    title: '管理员',
    key: 'admin',
    width: 100,
    render(row) {
      return h(NSwitch, {
        value: row.admin === 1,
      })
    },
  },
  {
    title: '权限数量',
    key: 'permission',
    width: 100,
    render(row) {
      try {
        let permissionList = JSON.parse(row.permission)
        return h('span', {}, permissionList.length)
      } catch (e) {
        return h('span', {}, '0')
      }
    },
  },
  {
    title: '操作',
    key: 'action',
    width: 150,
    render(row) {
      return h(NFlex, {}, [
        h(
          NButton,
          {
            type: 'primary',
            onClick() {
              openUpdateModal(row)
            },
          },
          '编辑',
        ),
        h(
          NButton,
          {
            type: 'error',
            onClick() {
              deleteGroup(row)
            },
          },
          '删除',
        ),
      ])
    },
  },
]
function openCreateModal() {
  groupId.value = ''
  groupName.value = ''
  groupAdmin.value = false
  permission.value = []
  createGroup.value = true
  showModal.value = true
}
function openUpdateModal(row) {
  groupId.value = row.id
  groupName.value = row.name
  groupAdmin.value = row.admin === 1
  try {
    permission.value = JSON.parse(row.permission)
  } catch (e) {
    permission.value = []
  }
  createGroup.value = false
  showModal.value = true
}
async function submitModal() {
  showModal.value = false
  let _permission
  try {
    _permission = JSON.stringify(permission.value)
  } catch (e) {
    _permission = '[]'
  }
  if (createGroup.value) {
    let [status, data] = await createAdminGroup(
      groupName.value,
      groupAdmin.value ? 1 : 0,
      _permission,
    )
    if (!status) {
      dialog.error({
        title: '创建失败',
        content: data,
        positiveText: '确定',
      })
    } else {
      message.success('创建成功')
    }
  } else {
    let [status, data] = await updateAdminGroup(
      groupId.value,
      groupName.value,
      groupAdmin.value ? 1 : 0,
      _permission,
    )
    if (!status) {
      dialog.error({
        title: '修改失败',
        content: data,
        positiveText: '确定',
      })
    } else {
      message.success('修改成功')
    }
  }
  await loadDataList()
}
function deleteGroup(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除用户组${row.name}？`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await deleteAdminGroup(row.id)
      if (!status) {
        dialog.error({
          title: '删除失败',
          content: data,
          positiveText: '确定'
        })
      } else {
        message.success('删除成功')
      }
      await loadDataList()
    }
  })
}
async function loadUserInfo() {
  let [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
  return true
}
async function loadAdminGroupList() {
  let [status, data] = await getAdminGroupList()
  if (!status) {
    return dialogError(data)
  }
  adminGroup.groups = data
  return true
}
async function loadAdminPermissionList() {
  let [status, data] = await getAdminPermissionList()
  if (!status) {
    return dialogError(data)
  }
  adminPermission.permissions = data
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!(await loadAdminGroupList())) return
  if (!(await loadAdminPermissionList())) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">管理员 用户组管理</n-h1>
    </template>
    <n-flex>
      <n-button size="large" type="primary" @click="openCreateModal()">创建用户组</n-button>
      <n-data-table :data="adminGroup.groups" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card size="large" style="max-width: 400px">
      <template #header>
        <n-h1>{{ createGroup ? '创建' : '编辑' }}用户组</n-h1>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createGroup">
          <n-input :value="groupId" disabled readonly></n-input>
        </n-form-item>
        <n-form-item label="用户组名">
          <n-input v-model:value="groupName"></n-input>
        </n-form-item>
        <n-form-item label="管理员">
          <n-switch v-model:value="groupAdmin"></n-switch>
        </n-form-item>
        <n-form-item label="权限">
          <n-scrollbar style="max-height: calc(100vh - 600px)">
            <n-tree
              v-model:checked-keys="permission"
              default-expand-all
              checkable
              :data="permissionTree"
            ></n-tree>
          </n-scrollbar>
        </n-form-item>
      </n-form>
      <template #action>
        <n-flex justify="end">
          <n-button size="large" @click="showModal = false">取消</n-button>
          <n-button size="large" type="primary" @click="submitModal()">确定</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped></style>
