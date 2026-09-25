<script setup>
import { useDialog, useLoadingBar } from 'naive-ui'
import { useAdminGroupStore, useAdminPermissionStore, useUserStore } from '@/stores/index.js'
import { onMounted } from 'vue'
import { getAdminPermissionList, getUserInfo } from '@/api/index.js'

const dialog = useDialog()
const loadingBar = useLoadingBar()
const user = useUserStore()
const adminPermission = useAdminPermissionStore()

const columns = [
  {
    title: 'ID',
    key: 'id',
  },
]
const dialogError = content => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定'
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
  if (!await loadUserInfo()) return
  if (!await loadAdminPermissionList()) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>管理员 权限管理</h3>
    </template>
    <n-flex>
      <n-button size="large" type="primary">创建权限</n-button>
      <n-data-table :data="adminPermission.permissions" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
</template>

<style scoped></style>
