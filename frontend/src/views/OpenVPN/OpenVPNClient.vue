<script setup>
import { getUserInfo, getOpenVPNClientList, kickOpenVPNClient } from '@/api/index.js'
import { NButton, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import { useOpenVPNClientStore, useUserStore } from '@/stores/index.js'
import { h, onMounted } from 'vue'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()

const user = useUserStore()
const openVPNClient = useOpenVPNClientStore()

const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: '用户名',
    key: 'username',
  },
  {
    title: 'IP地址',
    key: 'ip_addr',
  },
  {
    title: '认证时间',
    key: 'auth_time',
    render(row) {
      const date = new Date(row.auth_time * 1000)
      return date.toLocaleString()
    }
  },
  {
    title: '操作',
    key: 'action',
    render(row) {
      return h(NButton, {
        type: 'error',
        size: 'large',
        onClick() {
          kickClient(row)
        }
      }, '踢出客户端')
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
function kickClient(row) {
  dialog.info({
    title: '是否踢出该客户端',
    content: `是否确认踢出客户端 ${row.username}(${row.ip_addr})？`,
    negativeText: '取消',
    positiveText: '确定',
    async onPositiveClick() {
      let [status, data] = await kickOpenVPNClient(row.id)
      if (!status) {
        dialog.error({
          title: '踢出失败',
          content: data,
          positiveText: '确定',
        })
      } else {
        message.success('踢出成功')
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
async function loadOpenVPNClientList() {
  let [status, data] = await getOpenVPNClientList()
  if (!status) {
    return dialogError(data)
  }
  openVPNClient.clients = data.results.data
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!(await loadOpenVPNClientList())) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">OpenVPN 客户端管理</n-h1>
    </template>
    <n-data-table :data="openVPNClient.clients" :columns="columns"></n-data-table>
  </n-card>
</template>

<style scoped></style>
