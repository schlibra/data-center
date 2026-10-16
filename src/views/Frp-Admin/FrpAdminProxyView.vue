<script setup>
import {computed, h, onMounted, onUnmounted} from 'vue'
import {NTag, useDialog, useLoadingBar, useMessage} from 'naive-ui'
import { calcSize } from '@/utils'
import {clearOfflineProxyFrpAdmin, getFrpAdminProxyList, getFrpProxyList, getUserInfo} from '@/api'
import { useFrpProxyStore, useUserStore } from '@/stores'
import {event} from "@/utils/event.js";

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const user = useUserStore()
const frpProxy = useFrpProxyStore()

const columns = computed(() => [
  {
    title: '状态',
    key: 'status.phase',
    width: 100,
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
    filterOptions: [
      {
        label: '在线',
        value: 'online'
      },
      {
        label: '离线',
        value: 'offline'
      }
    ],
    filter(value, row) {
      return row.status.phase === value
    }
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
    sorter(rowA, rowB) {
      return rowA.spec[rowB.spec.type].remotePort - rowB.spec[rowB.spec.type].remotePort
    }
  },
  {
    title: 'Token',
    key: 'user',
    width: 100,
    filterOptions: function() {
      let list = {}
      frpProxy.proxy.forEach(item => {
        list[item.user] = item.user
      })
      let options = []
      for (let key in list) {
        options.push({
          label: key,
          value: key,
        })
      }
      return options
    }(),
    filter(value, row) {
      return row.user === value
    }
  },
  {
    title: '类型',
    key: 'spec.type',
    width: 80,
    render(row) {
      return h(
        NTag,
        {
          type: row.spec.type === 'tcp' ? 'success' : row.spec.type === 'udp' ? 'info' : 'default',
          round: true,
        },
        row.spec.type,
      )
    },
    filterOptions: function() {
      let list = {}
      frpProxy.proxy.forEach(item => {
        list[item.spec.type] = item.spec.type
      })
      let options = []
      for (let key in list) {
        options.push({
          label: key,
          value: key,
        })
      }
      return options
    }(),
    filter(value, row) {
      return row.spec.type === value
    }
  },
  {
    title: '今日入站流量',
    key: 'status.todayTrafficIn',
    width: 130,
    render(row) {
      return h('span', {}, calcSize(row.status.todayTrafficIn))
    },
    sorter(rowA, rowB) {
      return rowA.status.todayTrafficIn - rowB.status.todayTrafficIn
    }
  },
  {
    title: '今日出站流量',
    key: 'status.todayTrafficOut',
    width: 130,
    render(row) {
      return h('span', {}, calcSize(row.status.todayTrafficOut))
    },
    sorter(rowA, rowB) {
      return rowA.status.todayTrafficOut - rowB.status.todayTrafficOut
    }
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
])
const dialogError = (content) => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}
function clearOffline() {
  dialog.info({
    title: '清理离线映射',
    content: '确定要清理离线映射吗？',
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await clearOfflineProxyFrpAdmin()
      if (!status) {
        dialog.error({
          title: '清理失败',
          content: data,
          positiveText: '确定',
        })
      } else {
        message.success("清理成功")
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
async function loadFrpProxyList() {
  let [status, data] = await getFrpAdminProxyList()
  if (!status) {
    return dialogError(data)
  }
  frpProxy.proxy = data
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!await loadFrpProxyList()) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
  event.on('trigger-refresh', loadDataList)
})
onUnmounted(() => {
  event.off('trigger-refresh', loadDataList)
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">Frp 映射管理（管理员）</n-h1>
    </template>
    <n-flex>
      <n-button size="large" type="warning" @click="clearOffline">清理离线</n-button>
      <n-data-table :data="frpProxy.proxy" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
</template>

<style scoped></style>
