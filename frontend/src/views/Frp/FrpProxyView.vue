<script setup>
import { h, onMounted } from 'vue'
import { NTag, useDialog, useLoadingBar } from 'naive-ui'
import { calcSize } from '@/utils'
import { getFrpProxyList, getUserInfo } from '@/api'
import { useFrpProxyStore, useUserStore } from '@/stores'

const dialog = useDialog()
const loadingBar = useLoadingBar()
const user = useUserStore()
const frpProxy = useFrpProxyStore()

const columns = [
  {
    title: '状态',
    key: 'status.phase',
    render(row) {
      const _status = row.status.phase
      return h(
        NTag,
        {
          type: _status === 'online' ? 'success' : _status === 'offline' ? 'default' : 'warning',
          round: true,
        },
        _status === 'online' ? '在线' : _status === 'offline' ? '离线' : '未知',
      )
    },
  },
  {
    title: '名称',
    key: 'name',
    width: 200,
    render(row) {
      let full_name = row.name
      const user = row.user
      full_name = full_name.replace(`${user}.`, '')
      return h('span', {}, full_name)
    },
  },
  {
    title: '远程端口',
    key: 'remotePort',
    width: 100,
    render(row) {
      const type = row.spec.type
      const port = row.spec[type].remotePort
      return h('spane', {}, port)
    },
  },
  {
    title: 'Token',
    key: 'user',
    width: 100,
  },
  {
    title: '类型',
    key: 'spec.type',
    width: 70,
    render(row) {
      return h(
        NTag,
        {
          type: 'info',
          round: true,
        },
        row.spec.type,
      )
    },
  },
  {
    title: '今日入站流量',
    key: 'status.todayTrafficIn',
    width: 110,
    render(row) {
      return h('span', {}, calcSize(row.status.todayTrafficIn))
    },
  },
  {
    title: '今日出站流量',
    key: 'status.todayTrafficOut',
    width: 110,
    render(row) {
      return h('span', {}, calcSize(row.status.todayTrafficOut))
    },
  },
  {
    title: '当前连接数量',
    key: 'status.curConns',
    width: 110,
  },
  {
    title: '上次连接时间',
    key: 'status.lastStartAt',
    width: 200,
    render(row) {
      const time = new Date()
      time.setTime(row.status.lastStartAt * 1000)
      return h('span', {}, time.toLocaleString())
    },
  },
  {
    title: '上次断开时间',
    key: 'status.lastCloseAt',
    width: 200,
    render(row) {
      const _t = row.status.lastCloseAt
      if (_t) {
        const time = new Date()
        time.setTime(_t * 1000)
        return h('span', {}, time.toLocaleString())
      } else {
        return h('span', {}, '-')
      }
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
async function loadFrpProxyList() {
  let [status, data] = await getFrpProxyList()
  if (!status) {
    return dialogError(data)
  }
  frpProxy.proxy = data
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadFrpProxyList()) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>Frp 映射管理</h3>
    </template>
    <n-data-table :data="frpProxy.proxy" :columns="columns"></n-data-table>
  </n-card>
</template>

<style scoped></style>
