<script setup>
import {
  getUserInfo,
  listPveNode,
  listPveQemu, rebootPveQemuStatus, resetPveQemuStatus, resumePveQemuStatus, shutdownPveQemuStatus,
  startPveQemuStatus,
  stopPveQemuStatus,
  suspendPveQemuStatus
} from "@/api/index.js";
import {NButton, NFlex, NTag, useDialog, useLoadingBar, useMessage} from "naive-ui";
import {usePveNodeStore, usePveQemuStore, useUserStore} from "@/stores/index.js";
import {onMounted, onUnmounted, h, computed, ref} from "vue";
import {event} from "@/utils/event.js";
import {calcSize} from "@/utils/index.js";

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()

const user = useUserStore()
const pveNode = usePveNodeStore()
const pveQemu = usePveQemuStore()

const checkedRows = ref([])

const pveNodeSelectOptions = computed(() => pveNode.nodes.map(item => {
  return {
    label: item.node,
    value: item.node,
  }
}))

const columns = [
  {
    type: 'selection',
    key: 'vmid'
  },
  {
    title: '状态',
    key: 'status',
    width: 80,
    render(row) {
      return h(NTag, {
        type: row.status === 'running' ? 'success' :
            row.status === 'stopped' ? 'error' : 'warning',
        round: true
      }, row.status === 'running' ? '正常' :
      row.status === 'stopped' ? '停止' : row.status)
    }
  },
  {
    title: '虚拟机ID',
    key: 'vmid',
    width: 100,
    sorter(rowA, rowB) {
      return rowA.vmid - rowB.vmid
    }
  },
  {
    title: '虚拟机名称',
    key: 'name',
    width: 100,
  },
  {
    title: '标签',
    key: 'tags',
    width: 100,
    render(row) {
      if (!row.tags) return "-"
      return h(NFlex, {}, row.tags.split(";").map(item => {
        return h(NTag, {type: item.indexOf('172.16.') > -1 ? 'success' : 'info'}, item)
      }))
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
    width: 100,
    render(row) {
      return h('span', {}, row.disk ? `${calcSize(row.disk)} / ${calcSize(row.maxdisk)}` : calcSize(row.maxdisk))
    }
  },
  {
    title: '操作',
    key: 'action',
    width: 400,
    render(row) {
      let child = []
      if (row.status === 'stopped') {
        child.push(h(NButton, {
          type: 'success',
          onClick() {
            actionQemu(row, '启动', startPveQemuStatus)
          }
        }, '开机'));
        child.push(h(NButton, {
          type: 'info',
          onClick() {
            actionQemu(row, '恢复', resumePveQemuStatus)
          }
        }, '恢复'))
      }
      if (row.status === 'running') {
        child.push(h(NButton, {
          type: 'error',
          onClick() {
            actionQemu(row, '关闭', shutdownPveQemuStatus)
          }
        }, '关机'));
        child.push(h(NButton, {
          type: 'warning',
          onClick() {
            actionQemu(row, '重启', rebootPveQemuStatus)
          }
        }, '重启'))
        child.push(h(NButton, {
          type: 'warning',
          onClick() {
            actionQemu(row, '重置', resetPveQemuStatus)
          }
        }, '重置'))
        child.push(h(NButton, {
          type: 'error',
          onClick() {
            actionQemu(row, '停止', stopPveQemuStatus)
          }
        }, '停止'))
        child.push(h(NButton, {
          type: 'warning',
          onClick() {
            actionQemu(row, '暂停', suspendPveQemuStatus)
          }
        }, '暂停'))
      }
      return h(NFlex, {}, child)
    }
  }
]

const dialogError = content => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}
function actionQemu(row, actionText, actionFunction) {
  dialog.info({
    title: `${actionText}虚拟机`,
    content: `确定要${actionText}虚拟机${row.name}吗？`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await actionFunction(pveQemu.node, row.vmid)
      if (!status) {
        return dialog.error({
          title: `${actionText}失败`,
          content: data,
          positiveText: '确定',
        })
      } else {
        message.success(`${actionText}成功`)
      }
      await loadDataList()
    }
  })
}
function multiActionQemu(actionText, actionFunction) {
  if (checkedRows.value.length === 0) {
    message.warning('请选择要批量操作的虚拟机')
  } else {
    dialog.info({
      title: `批量${actionText}虚拟机`,
      content: `确定要批量${actionText}虚拟机吗？`,
      positiveText: '确定',
      negativeText: '取消',
      async onPositiveClick() {
        let successCount = 0
        let failCount = 0
        let error_msg = []
        for (const item of checkedRows.value) {
          let [status, data] = await actionFunction(pveQemu.node, item)
          if (!status) {
            failCount++
            error_msg.push(data)
          } else {
            successCount++
          }
        }
        dialog.info({
          title: `批量${actionText}完成`,
          content: `批量${actionText}完成，成功${successCount}个，失败${failCount}个`+ (failCount ? `，失败原因：${error_msg.join(';')}` : ''+``),
          positiveText: '确定'
        })
        await loadDataList()
      }
    })
  }
}
async function loadUserInfo() {
  let [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
  return true
}
async function loadPveQemuList() {
  let [status, data] = await listPveQemu(pveQemu.node)
  if (!status) {
    return dialogError(data)
  }
  pveQemu.qemu = data
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
  if (pveQemu.node === '') {
    pveQemu.node = pveNode.nodes[0].node
  }
  if (!await loadPveQemuList()) return
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
  <n-card>
    <template #header>
      <n-h1 prefix="bar">PVE QEMU管理</n-h1>
    </template>
    <n-flex>
      <n-select :options="pveNodeSelectOptions" v-model:value="pveQemu.node" @update:value="loadDataList()"></n-select>
      <n-space>
        <n-button type="success" @click="multiActionQemu('启动', startPveQemuStatus)">批量开机</n-button>
        <n-button type="error" @click="multiActionQemu('关闭', shutdownPveQemuStatus)">批量关机</n-button>
        <n-button type="error" @click="multiActionQemu('停止', stopPveQemuStatus)">批量停止</n-button>
        <n-button type="warning" @click="multiActionQemu('重启', rebootPveQemuStatus)">批量重启</n-button>
        <n-button type="warning" @click="multiActionQemu('重置', resetPveQemuStatus)">批量重置</n-button>
        <n-button type="info" @click="multiActionQemu('暂停', suspendPveQemuStatus)">批量暂停</n-button>
        <n-button type="info" @click="multiActionQemu('恢复', resumePveQemuStatus)">批量恢复</n-button>
      </n-space>
      <n-data-table @update:checked-row-keys="keys => checkedRows = keys" :row-key="row => row.vmid" :columns="columns" :data="pveQemu.qemu"></n-data-table>
    </n-flex>
  </n-card>
</template>

<style scoped>

</style>