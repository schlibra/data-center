<script setup>
import { NButton, NFlex, NSwitch, useDialog, useLoadingBar } from 'naive-ui'
import { h, onMounted } from 'vue'
import { getAdminGroupList, getAdminPermissionList, getUserInfo } from '@/api/index.js'
import { useAdminGroupStore, useAdminPermissionStore, useUserStore } from '@/stores/index.js'

const dialog = useDialog()
const loadingBar = useLoadingBar()
const user = useUserStore()
const adminGroup = useAdminGroupStore()
const adminPermission = useAdminPermissionStore()

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
  },
  {
    title: '用户组名',
    key: 'name',
  },
  {
    title: '管理员',
    key: 'admin',
    render(row) {
      return h(NSwitch, {
        value: row.admin === 1,
      })
    },
  },
  {
    title: '权限数量',
    key: 'permission',
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
    render(row) {
      return h(NFlex, {}, [
        h(NButton, {type: 'primary'}, '编辑'),
        h(NButton, {type: 'error'}, '删除'),
      ])
    },
  },
]

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
  <n-card>
    <template #header>
      <h3>管理员 用户组管理</h3>
    </template>
    <n-flex>
      <n-button size="large" type="primary">创建用户组</n-button>
      <n-data-table :data="adminGroup.groups" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
</template>

<style scoped></style>
