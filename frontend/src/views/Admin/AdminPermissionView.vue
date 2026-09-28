<script setup>
import { NButton, NFlex, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import { useAdminGroupStore, useAdminPermissionStore, useUserStore } from '@/stores/index.js'
import { computed, h, onMounted, ref } from 'vue'
import {
  createAdminPermission,
  deleteAdminPermission,
  getAdminPermissionList,
  getUserInfo,
  updateAdminPermission,
} from '@/api/index.js'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const user = useUserStore()
const adminPermission = useAdminPermissionStore()

const showModal = ref(false)
const createPermission = ref(true)
const isGroup = ref(true)
const permissionId = ref('')
const permissionKey = ref('')
const permissionName = ref('')
const permissionGroup = ref('')

const groupSelectOption = computed(() =>
  adminPermission.permissions.map((item) => {
    return {
      label: `${item.name} (${item.key})`,
      value: item.id,
    }
  }),
)
const permissionTree = computed(() =>
  adminPermission.permissions.map((item) => {
    return {
      label: `${item.name} (${item.key})`,
      key: item.id,
      suffix() {
        return h(NFlex, {}, [
          h(
            NButton,
            {
              type: 'info',
              size: 'tiny',
              onClick() {
                openCreateSubModal(item)
              }
            },
            '添加'
          ),
          h(
            NButton,
            {
              type: 'primary',
              size: 'tiny',
              onClick() {
                openUpdateModal(item)
              },
            },
            '编辑',
          ),
          h(
            NButton,
            {
              type: 'error',
              size: 'tiny',
              onClick() {
                deletePermission(item)
              },
            },
            '删除',
          ),
        ])
      },
      children: item.children.map((item) => {
        return {
          label: `${item.name} (${item.key})`,
          key: item.id,
          suffix() {
            return h(NFlex, {}, [
              h(
                NButton,
                {
                  type: 'primary',
                  size: 'tiny',
                  onClick() {
                    openUpdateModal(item)
                  },
                },
                '编辑',
              ),
              h(
                NButton,
                {
                  type: 'error',
                  size: 'tiny',
                  onClick() {
                    deletePermission(item)
                  },
                },
                '删除',
              ),
            ])
          },
        }
      }),
    }
  }).sort((a, b) => a.key - b.key),
)

const dialogError = (content) => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}
function openCreateModal() {
  permissionId.value = ''
  permissionKey.value = ''
  permissionName.value = ''
  isGroup.value = true
  permissionGroup.value = ''
  createPermission.value = true
  showModal.value = true
}
function openCreateSubModal(row) {
  permissionId.value = ''
  permissionKey.value = row.key
  permissionName.value = ''
  isGroup.value = false
  permissionGroup.value = row.id
  createPermission.value = true
  showModal.value = true
}
function openUpdateModal(row) {
  permissionId.value = row.id
  permissionKey.value = row.key
  permissionName.value = row.name
  isGroup.value = row.parent === 0
  permissionGroup.value = row.parent
  createPermission.value = false
  showModal.value = true
}
async function submitModal() {
  showModal.value = false
  if (isGroup.value) {
    permissionGroup.value = 0
  }
  if (createPermission.value) {
    let [status, data] = await createAdminPermission(
      permissionKey.value,
      permissionName.value,
      permissionGroup.value,
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
    let [status, data] = await updateAdminPermission(
      permissionId.value,
      permissionKey.value,
      permissionName.value,
      permissionGroup.value,
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
async function deletePermission(row) {
  dialog.info({
    title: '是否删除权限',
    content: `是否删除权限${row.name}(${row.key})？`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await deleteAdminPermission(row.id)
      if (!status) {
        dialog.error({
          title: '删除失败',
          content: data,
          positiveText: '确定',
        })
      } else {
        message.success('删除成功')
      }
      await loadDataList()
    },
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
async function loadAdminPermissionList() {
  let [status, data] = await getAdminPermissionList()
  if (!status) {
    return dialogError(data)
  }
  adminPermission.permissions = data
  console.log(JSON.stringify(adminPermission.tree))
  return true
}

async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
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
      <n-h1 prefix="bar">管理员 权限管理</n-h1>
    </template>
    <n-space vertical>
      <n-button @click="openCreateModal()" size="large" type="primary">创建权限</n-button>
      <n-tree
        style="max-width: 600px"
        default-expand-all
        :data="permissionTree"
        block-line
        show-line
      ></n-tree>
    </n-space>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card size="large" style="max-width: 400px">
      <template #header>
        <n-h1 prefix="bar">{{ createPermission ? '创建' : '编辑' }}权限</n-h1>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createPermission">
          <n-input :value="permissionId" readonly disabled></n-input>
        </n-form-item>
        <n-form-item label="Key">
          <n-input v-model:value="permissionKey"></n-input>
        </n-form-item>
        <n-form-item label="权限名称">
          <n-input v-model:value="permissionName" @keydown.enter="submitModal()"></n-input>
        </n-form-item>
        <n-form-item label="是否为分组">
          <n-switch v-model:value="isGroup"></n-switch>
        </n-form-item>
        <n-form-item label="分组" v-if="!isGroup">
          <n-select v-model:value="permissionGroup" :options="groupSelectOption"></n-select>
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
