<script setup>

import {h, onMounted, onUnmounted, ref} from "vue";
import {event} from "@/utils/event.js";
import {createMailDomain, deleteMailDomain, getUserInfo, listMailDomain, updateMailDomain} from "@/api/index.js";
import {useMailDomainStore, useMailListStore, useUserStore} from "@/stores/index.js";
import {NButton, NFlex, NSwitch, useDialog, useLoadingBar, useMessage} from "naive-ui";
import router from "@/router/index.js";

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()

const columns = [
  {
    title: 'ID',
    key: 'id'
  },
  {
    title: '域名',
    key: 'domain'
  },
  {
    title: '名称',
    key: 'name'
  },
  {
    title: '数据表ID',
    key: 'table_id'
  },
  {
    title: '服务器',
    key: 'server'
  },
  {
    title: '端口',
    key: 'port'
  },
  {
    title: '启用SSL',
    key: 'ssl',
    render(row) {
      return h(NSwitch, {
        value: row.ssl === "1" || row.ssl === "Y",
        onUpdateValue(value) {
          updateSSL(row, value)
        }
      })
    }
  },
  {
    title: '操作',
    key: 'action',
    render(row) {
      return h(NFlex, {}, [
          h(NButton, {
            size: 'large',
            type: 'primary',
            onClick() {
              openUpdateModal(row)
            }
          }, '编辑'),
          h(NButton, {
            size: 'large',
            type: 'info',
            onClick() {
              goMailList(row)
            }
          }, '邮箱列表'),
          h(NButton, {
            size: 'large',
            type: 'error',
            onClick() {
              deleteDomain(row)
            }
          }, '删除'),
      ])
    }
  }
]

const dialogError = content => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定'
  })
  loadingBar.error()
}

const user = useUserStore()
const mailDomain = useMailDomainStore()
const mailList = useMailListStore()

const showModal = ref(false)
const createDomain = ref(true)
const domainId = ref('')
const domainValue = ref('')
const domainName = ref('')
const domainTableId = ref('')
const domainServer = ref('')
const domainPort = ref('')
const domainSSL = ref(false)

function openCreateModal() {
  domainId.value = ''
  domainValue.value = ''
  domainName.value = ''
  domainTableId.value = ''
  domainServer.value = ''
  domainPort.value = ''
  domainSSL.value = false
  showModal.value = true
  createDomain.value = true
}
function openUpdateModal(row) {
  domainId.value = row.id
  domainValue.value = row.domain
  domainName.value = row.name
  domainTableId.value = row.table_id
  domainServer.value = row.server
  domainPort.value = row.port
  domainSSL.value = row.ssl === "1" || row.ssl === "Y"
  showModal.value = true
  createDomain.value = false
}
function deleteDomain(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除域名 ${row.domain}（${row.name}）？`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      const [status, data] = await deleteMailDomain(row.id)
      if (!status) {
        dialog.error({
          title: '删除失败',
          content: data,
          positiveText: '确定'
        })
      } else {
        message.success('删除成功')
      }
      await loadDataList()
    }
  })
}
async function submitModal() {
  showModal.value = false
  if (createDomain.value) {
    const [status, data] = await createMailDomain(domainValue.value, domainName.value, domainTableId.value, domainServer.value, domainPort.value, domainSSL.value ? "1" : "0")
    if (!status) {
      dialog.error({
        title: '创建失败',
        content: data,
        positiveText: '确定'
      })
    } else {
      message.success('创建成功')
    }
  } else {
    const [status, data] = await updateMailDomain(domainId.value, domainValue.value, domainName.value, domainTableId.value, domainServer.value, domainPort.value, domainSSL.value ? "1" : "0")
    if (!status) {
      dialog.error({
        title: '更新失败',
        content: data,
        positiveText: '确定'
      })
    } else {
      message.success('更新成功')
    }
  }
  await loadDataList()
}
async function updateSSL(row, value) {
  const [status, data] = await updateMailDomain(row.id, row.domain, row.name, row.table_id, row.server, row.port, value ? "1" : "0")
  if (!status) {
    return dialog.error({
      title: '修改失败',
      content: data,
      positiveText: '确定'
    })
  } else {
    message.success('修改成功')
  }
  await loadDataList()
}
async function goMailList(row) {
  mailList.domainId = row.id
  router.push('/mail/list')
}
async function loadUserInfo() {
  const [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
  return true
}

async function loadMailDomainList() {
  const [status, data] = await listMailDomain()
  if (!status) {
    return dialogError(data)
  }
  mailDomain.domain = data
  return true
}

async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadMailDomainList()) return
  loadingBar.finish()
}

onMounted(() => {
  loadDataList()
  event.on('trigger-refresh', loadDataList)
})
onUnmounted(() => {
  event.off('trigger-refresh', loadDataList)
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">邮箱 域名管理</n-h1>
    </template>
    <n-flex>
      <n-button size="large" type="primary" @click="openCreateModal()">创建域名</n-button>
      <n-data-table :data="mailDomain.domain" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card size="large" style="max-width: 400px">
      <template #header>
        <n-h1>{{ createDomain ? '创建域名' : '编辑域名' }}</n-h1>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createDomain">
          <n-input :value="domainId" disabled readonly></n-input>
        </n-form-item>
        <n-form-item label="域名">
          <n-input v-model:value="domainValue"></n-input>
        </n-form-item>
        <n-form-item label="名称">
          <n-input v-model:value="domainName"></n-input>
          </n-form-item>
        <n-form-item label="数据表ID">
          <n-input v-model:value="domainTableId"></n-input>
        </n-form-item>
        <n-form-item label="服务器">
          <n-input v-model:value="domainServer"></n-input>
        </n-form-item>
        <n-form-item label="端口">
          <n-input-number v-model:value="domainPort"></n-input-number>
        </n-form-item>
        <n-form-item label="启用SSL">
          <n-switch v-model:value="domainSSL"></n-switch>
        </n-form-item>
      </n-form>
      <template #action>
        <n-flex justify="end">
          <n-button size="large" @click="showModal = true">取消</n-button>
          <n-button size="large" type="primary" @click="submitModal">确定</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped>

</style>