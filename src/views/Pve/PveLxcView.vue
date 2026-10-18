<script setup>

import {
  getUserInfo,
  listPveLxc,
  listPveNode,
  rebootPveLxcStatus, resumePveLxcStatus,
  shutdownPveLxcStatus, startPveLxcStatus,
  stopPveLxcStatus, suspendPveLxcStatus
} from "@/api/index.js";
import {onMounted, onUnmounted, h, computed} from "vue";
import {event} from "@/utils/event.js";
import {NButton, NFlex, NTag, useDialog, useLoadingBar, useMessage} from "naive-ui";
import {usePveLxcStore, usePveNodeStore, useUserStore} from "@/stores/index.js";
import {calcSize} from "@/utils/index.js";

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()

const pveNodeSelectOptions = computed(() => pveNode.nodes.map(item => {
  return {
    label: item.node,
    value: item.node
  }
}))

const dialogError = content => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定'
  })
  loadingBar.error()
}

const user = useUserStore()
const pveNode = usePveNodeStore()
const pveLxc = usePveLxcStore()

const columns = [
  {
    title: '容器ID',
    key: 'vmid',
    width: 80,
  },
  {
    title: '容器名称',
    key: 'name',
    width: 100
  },
  {
    title: '状态',
    key: 'status',
    width: 100,
    render(row) {
      return h(NTag, {
        type: row.status === 'running' ? 'success' : row.status === 'stopped' ? 'error' : 'warning',
        round: true
      }, row.status === 'running' ? '运行' : row.status === 'stopped' ? '停止' : row.status)
    }
  },
  {
    title: '内存',
    key: 'mem',
    width: 150,
    render(row) {
      return h('span', {}, `${calcSize(row.mem)} / ${calcSize(row.maxmem)}`)
    }
  },
  {
    title: '硬盘',
    key: 'disk',
    width: 150,
    render(row) {
      return h('span', {}, `${calcSize(row.disk)} / ${calcSize(row.maxdisk)}`)
    }
  },
  {
    title: 'Swap',
    key: 'swap',
    width: 150,
    render(row) {
      return h('span', {}, `${calcSize(row.swap)} / ${calcSize(row.maxswap)}`)
    }
  },
  {
    title: '操作',
    key: 'action',
    width: 300,
    render(row) {
      let child = []
      if (row.status === 'running') {
        child.push(h(NButton, {
          type: 'error',
          onClick() {
            actionLxc(row, '关闭', shutdownPveLxcStatus)
          }
        }, '关机'))
        child.push(h(NButton, {
          type: 'error',
          onClick() {
            actionLxc(row, '停止', stopPveLxcStatus)
          }
        }, '停止'))
        child.push(h(NButton, {
          type: 'warning',
          onClick() {
            actionLxc(row, '重启', rebootPveLxcStatus)
          }
        }, '重启'))
        child.push(h(NButton, {
          type: 'info',
          onClick() {
            actionLxc(row, '暂停', suspendPveLxcStatus)
          }
        }, '暂停'))
      }
      if (row.status === 'stopped') {
        child.push(h(NButton, {
          type: 'success',
          onClick() {
            actionLxc(row, '启动', startPveLxcStatus)
          }
        }, '开机'))
        child.push(h(NButton, {
          type: 'info',
          onClick() {
            actionLxc(row, '恢复', resumePveLxcStatus)
          }
        }, '恢复'))
      }
      return h(NFlex, {}, child)
    }
  }
]
function actionLxc(row, actionText, actionFunction) {
  dialog.info({
    title: `${actionText}容器`,
    content: `确定要${actionText}容器吗？`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await actionFunction(row)
      if (!status) {
        dialog.error({
          title: `${actionText}容器失败`,
          content: data,
          positiveText: '确定'
        })
      } else {
        message.success(`${actionText}容器成功`)
      }
      await loadDataList()
    }
  })
}
async function loadPveLxcList() {
  let [status, data] = await listPveLxc(pveLxc.node)
  if (!status) {
    return dialogError(data)
  }
  pveLxc.lxc = data
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
async function loadPveNodeList() {
  let [status, data] = await listPveNode()
  if (!status) {
    return dialogError(data)
  }
  pveNode.nodes = data
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadPveNodeList()) return
  if (pveLxc.node === '') {
    pveLxc.node = pveNode.nodes[0].node
  }
  if (!await loadPveLxcList()) return
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
      <n-h1 prefix="bar">PVE LXC管理</n-h1>
    </template>
    <n-flex>
      <n-select :options="pveNodeSelectOptions" v-model:value="pveLxc.node" @update:value="loadDataList()"></n-select>
      <n-data-table :columns="columns" :data="pveLxc.lxc"></n-data-table>
    </n-flex>
  </n-card>
</template>

<style scoped>

</style>