<script setup>
import {h, onUnmounted} from 'vue';
import {getUserInfo, listPveNode} from "@/api/index.js";
import {usePveLxcStore, usePveNodeStore, usePveQemuStore, useUserStore} from "@/stores/index.js";
import {NButton, NFlex, NTag, useDialog, useLoadingBar} from "naive-ui";
import {computed, onMounted} from "vue";
import {calcSize, calcTime} from "@/utils/index.js";
import {event} from "@/utils/event.js";
import router from "@/router/index.js";

const dialog = useDialog()
const loadingBar = useLoadingBar()

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
const pveQemu = usePveQemuStore()
const pveLxc = usePveLxcStore()

const columns = computed(() => [
  {
    title: '节点ID',
    key: 'node',
    width: 100
  },
  {
    title: '节点状态',
    key: 'status',
    width: 100,
    render(row) {
      return h(NTag, {type: 'success', round: true}, row.status === 'online' ? '在线' : '离线')
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
    title: '运行时长',
    key: 'uptime',
    width: 150,
    render(row) {
      return h('span', {}, calcTime(row.uptime))
    }
  },
  {
    title: '操作',
    key: 'action',
    width: 250,
    render(row) {
      return h(NFlex, {}, [
          h(NButton, {
            type: 'info',
            onClick() {
              goQemu(row)
            }
          }, 'QEMU管理'),
          h(NButton, {
            type: 'primary',
            onClick() {
              goLxc(row)
            }
          }, 'LXC管理'),
      ])
    }
  }
])
function goQemu(row) {
  pveQemu.node = row.node
  router.push("/pve/qemu")
}
function goLxc(row) {
  pveLxc.node = row.node
  router.push("/pve/lxc")
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
      <n-h1 prefix="bar">PVE 节点管理</n-h1>
    </template>
    <n-data-table :columns="columns" :data="pveNode.nodes"></n-data-table>
  </n-card>
</template>

<style scoped>

</style>