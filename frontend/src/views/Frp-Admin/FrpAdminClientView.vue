<script setup>
import { h, onMounted } from 'vue'
import { NTag, useDialog, useLoadingBar } from 'naive-ui'
import { getFrpAdminClientList, getFrpClientList, getUserInfo } from '@/api'
import { useFrpClientStore, useUserStore } from '@/stores'

const dialog = useDialog()
const loadingBar = useLoadingBar()
const user = useUserStore()
const frpClient = useFrpClientStore()
const columns = [
  {
    title: '状态',
    key: 'online',
    width: 70,
    render(row) {
      return h(
        NTag,
        {
          size: 'large',
          type: row.online ? 'success' : 'default',
          round: true,
        },
        row.online ? '在线' : '离线',
      )
    },
  },
  {
    title: 'Token',
    key: 'user',
    width: 100,
  },
  {
    title: '客户端IP',
    key: 'clientIP',
    width: 150,
  },
  {
    title: '主机名',
    key: 'hostname',
    width: 150,
  },
  {
    title: '客户端ID',
    key: 'clientID',
    width: 160,
  },
  {
    title: '客户端版本',
    key: 'version',
    width: 100,
  },
  {
    title: '首次连接时间',
    key: 'firstConnectedAt',
    width: 200,
    render(row) {
      const time = new Date()
      time.setTime(row.firstConnectedAt * 1000)
      return h('span', {}, time.toLocaleString())
    },
  },
  {
    title: '上次连接时间',
    key: 'lastConnectedAt',
    width: 200,
    render(row) {
      const time = new Date()
      time.setTime(row.lastConnectedAt * 1000)
      return h('span', {}, time.toLocaleString())
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
async function loadFrpClientList() {
  let [status, data] = await getFrpAdminClientList()
  if (!status) {
    return dialogError(data)
  }
  frpClient.client = data
  return true
}
async function loadUserInfo() {
  let [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!await loadFrpClientList()) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">Frp 客户端管理（管理员）</n-h1>
    </template>
    <n-data-table :data="frpClient.client" :columns="columns"></n-data-table>
  </n-card>
</template>

<style scoped></style>
