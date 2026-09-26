<script setup>
import { computed, h, onMounted, ref } from 'vue'
import { NButton, NFlex, NSwitch, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import {
  createAdminUser,
  deleteAdminUser,
  getAdminUserList, passwordAdminUser,
  updateAdminUser,
  getUserInfo,
  getAdminGroupList,
} from '@/api'
import { useAdminUserStore, useUserStore, useAdminGroupStore } from '@/stores'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const user = useUserStore()
const adminUser = useAdminUserStore()
const adminGroup = useAdminGroupStore()
const groupSelectOptions = computed(() =>
  adminGroup.groups.map((item) => {
    return {
      label: item.name,
      value: item.id,
    }
  }),
)
const showModal = ref(false)
const showPasswordModal = ref(false)
const createUser = ref(true)
const userId = ref('')
const username = ref('')
const password = ref('')
const nickname = ref('')
const userEnable = ref(false)
const userGroup = ref('')
const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 60,
  },
  {
    title: '用户名',
    key: 'username',
    width: 100
  },
  {
    title: '昵称',
    key: 'nickname',
    width: 100
  },
  {
    title: '用户组',
    key: 'group_info.name',
    width: 100,
  },
  {
    title: '启用',
    key: 'enable',
    width: 100,
    render(row) {
      return h(NSwitch, {
        value: row.enable === 1,
        onUpdateValue(value) {
          updateEnable(row, value)
        },
      })
    },
  },
  {
    title: '操作',
    key: 'action',
    width: 250,
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
            type: 'warning',
            onClick() {
              openPasswordModal(row)
            }
          },
          '修改密码',
        ),
        h(
          NButton,
          {
            type: 'error',
            onClick() {
              deleteUser(row)
            }
          },
          '删除',
        ),
      ])
    },
  },
]
const dialogError = (content) => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}
async function loadUserInfo() {
  let [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
  return true
}
async function loadAdminUserList() {
  let [status, data] = await getAdminUserList()
  if (!status) {
    return dialogError(data)
  }
  adminUser.users = data
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
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!(await loadAdminGroupList())) return
  if (!await loadAdminUserList()) return
  loadingBar.finish()
}
function openPasswordModal(row) {
  userId.value = row.id
  username.value = row.username
  password.value = ""
  showPasswordModal.value = true
}
function openCreateModal() {
  userId.value = ''
  username.value = ''
  nickname.value = ''
  password.value = ''
  userEnable.value = true
  userGroup.value = ''
  createUser.value = true
  showModal.value = true
}
function openUpdateModal(row) {
  userId.value = row.id
  username.value = row.username
  nickname.value = row.nickname
  password.value = row.password
  userEnable.value = row.enable === 1
  userGroup.value = row.group
  createUser.value = false
  showModal.value = true
}
async function submitPassword() {
  showPasswordModal.value = false
  let [status, data] = await passwordAdminUser(userId.value, password.value)
  if (!status) {
    dialog.error({
      title: "用户修改失败",
      content: data,
      positiveText: "确定"
    })
  } else {
    message.success("用户修改成功")
  }
  await loadDataList()
}
async function submitModal() {
  showModal.value = false
  if (createUser.value) {
    let [status, data] = await createAdminUser(
      username.value,
      password.value,
      nickname.value,
      userGroup.value,
      userEnable.value ? 1 : 0,
    )
    if (!status) {
      dialog.error({
        title: '用户创建失败',
        content: data,
        positiveText: '确定',
      })
    } else {
      message.success('用户创建成功')
    }
  } else {
    let [status, data] = await updateAdminUser(
      userId.value,
      nickname.value,
      userGroup.value,
      userEnable.value ? 1 : 0,
    )
    if (!status) {
      dialog.error({
        title: '用户修改失败',
        content: data,
        positiveText: '确定',
      })
    } else {
      message.success('用户修改成功')
    }
  }
  await loadDataList()
}
async function deleteUser(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除用户${row.username}(${row.id})`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await deleteAdminUser(row.id)
      if (!status) {
        dialog.error({
          title: '用户删除失败',
          content: data,
          positiveText: '确定',
        })
      } else {
        message.success('用户删除成功')
      }
      await loadDataList()
    },
  })
}
async function updateEnable(row, value) {
  let [status, data] = await updateAdminUser(row.id, row.nickname, row.group, value ? 1 : 0)
  if (!status) {
    dialog.error({
      title: '用户修改失败',
      content: data,
      positiveText: '确定',
    })
  } else {
    message.success('用户修改成功')
  }
  await loadDataList()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>管理员 用户管理</h3>
    </template>
    <n-flex>
      <n-button type="primary" size="large" @click="openCreateModal()">创建用户</n-button>
      <n-data-table :data="adminUser.users" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card style="max-width: 400px">
      <template #header>
        <h3>{{ createUser ? '创建' : '编辑' }}用户</h3>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createUser">
          <n-input :value="userId" readonly disabled></n-input>
        </n-form-item>
        <n-form-item label="用户名">
          <n-input
            v-model:value="username"
            :disabled="!createUser"
            :readonly="!createUser"
          ></n-input>
        </n-form-item>
        <n-form-item label="昵称">
          <n-input v-model:value="nickname"></n-input>
        </n-form-item>
        <n-form-item label="密码" v-if="createUser">
          <n-input v-model:value="password" type="password"></n-input>
        </n-form-item>
        <n-form-item label="启用">
          <n-switch v-model:value="userEnable"></n-switch>
        </n-form-item>
        <n-form-item label="用户组">
          <n-select v-model:value="userGroup" :options="groupSelectOptions"></n-select>
        </n-form-item>
      </n-form>
      <template #footer>
        <n-flex justify="end">
          <n-button size="large" @click="showModal = false">取消</n-button>
          <n-button size="large" type="primary" @click="submitModal()">确定</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
  <n-modal v-model:show="showPasswordModal">
    <n-card style="max-width: 400px">
      <template #header>
        <h3>设置密码</h3>
      </template>
      <n-form>
        <n-form-item label="用户ID">
          <n-input readonly disabled :value="userId"></n-input>
        </n-form-item>
        <n-form-item label="用户名">
          <n-input readonly disabled v-model:value="username"></n-input>
        </n-form-item>
        <n-form-item label="密码">
          <n-input v-model:value="password"></n-input>
        </n-form-item>
      </n-form>
      <template #footer>
        <n-flex justify="end">
          <n-button size="large" @click="showPasswordModal = false">取消</n-button>
          <n-button size="large" type="primary" @click="submitPassword()">确定</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped></style>
