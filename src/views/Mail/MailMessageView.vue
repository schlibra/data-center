<script setup>
import {getUserInfo, listMailDomain, listMailList, listMailMessage} from "@/api/index.js";
import router from "@/router/index.js";
import {computed, onMounted, onUnmounted, ref} from "vue";
import {event} from "@/utils/event.js";
import {useMailDomainStore, useMailListStore, useMailMessageStore, useUserStore} from "@/stores/index.js";
import {useDialog, useLoadingBar} from "naive-ui";

const dialog = useDialog()
const loadingBar = useLoadingBar()

const mailDomainSelectOptions = computed(() => mailDomain.domain.map(item => {
  return {
    label: `${item.domain}（${item.name}）`,
    value: item.id
  }
}))
const mailListSelectOptions = computed(() => mailList.mail.map(item => {
  return {
    label: `${item.email}`,
    value: item.id
  }
}))
const dialogError = content => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',

  })
  loadingBar.error()
}

const user = useUserStore()
const mailDomain = useMailDomainStore()
const mailList = useMailListStore()
const mailMessage = useMailMessageStore()

const showModal = ref(false)
const mailHtml = ref('')
const mailSubject = ref('')
const mailFromAddr = ref('')
const mailFromName = ref('')
const mailDate = ref('')

function openMessage(item) {
  mailHtml.value = item.body.html
  mailSubject.value = item.subject
  mailFromAddr.value = item.from.address
  mailFromName.value = item.from.name
  mailDate.value = item.date
  showModal.value = true
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
  if (mailDomain.count === 0) {
    return dialog.error({
      title: '没有创建邮箱域名',
      content: '请先创建邮箱域名',
      positiveText: '确定',
      onClose() {
        router.push('/mail/domain')
      },
      onPositiveClick() {
        router.push('/mail/domain')
      }
    })
  }
  return true
}
async function loadMailList() {
  if (mailList.domainId === '') {
    mailList.domainId = mailDomain.domain[0].id
  }
  const [status, data] = await listMailList(mailList.domainId)
  if (!status) {
    return dialogError(data)
  }
  mailList.mail = data
  if (mailList.count === 0) {
    return dialog.error({
      title: '没有创建邮箱',
      content: '请先创建邮箱',
      positiveText: '确定',
      onClose() {
        router.push('/mail/list')
      },
      onPositiveClick() {
        router.push('/mail/list')
      }
    })
  }
  return true
}
async function loadMailMessage() {
  if (mailMessage.mailRecord === '') {
    mailMessage.mailRecord = mailList.mail[0].id
    console.log(mailList.mail[0])
  }
  const [status, data] = await listMailMessage(mailList.domainId, mailMessage.mailRecord)
  if (!status) {
    return dialogError(data)
  }
  if (data.count === 0) {
    mailMessage.message = []
  } else {
    mailMessage.message = data.mails
  }
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadMailDomainList()) return
  if (!await loadMailList()) return
  if (!await loadMailMessage()) return
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
      <n-h1 prefix="bar">邮箱 邮件管理</n-h1>
    </template>
    <n-form>
      <n-form-item label="域名">
        <n-select @update:value="loadDataList()" v-model:value="mailList.domainId" :options="mailDomainSelectOptions" style="max-width: 300px"></n-select>
      </n-form-item>
      <n-form-item label="邮箱">
        <n-select @update:value="loadDataList()" v-model:value="mailMessage.mailRecord" :options="mailListSelectOptions" style="max-width: 300px"></n-select>
      </n-form-item>
    </n-form>
    <n-list hoverable clickable bordered>
      <n-list-item v-for="item in mailMessage.message" :key="item.id" @click="openMessage(item)">
        <n-thing :title="item.subject" :title-extra="item.date" :description="item.body.text.substr(1, 100)"></n-thing>
      </n-list-item>
    </n-list>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card style="max-width: 82vw">
      <template #header>
        <n-h1>邮件详情</n-h1>
      </template>
      <n-descriptions>
        <n-descriptions-item label="主题"><n-text>{{mailSubject}}</n-text></n-descriptions-item>
        <n-descriptions-item label="发件人"><n-text>{{mailFromName}}（{{mailFromAddr}}）</n-text></n-descriptions-item>
        <n-descriptions-item label="时间"><n-text>{{mailDate}}</n-text></n-descriptions-item>
      </n-descriptions>
      <iframe style="width: 80vw; height: calc(100vh - 300px);border: none" :srcdoc="mailHtml" sandbox="allow-same-origin allow-scripts"></iframe>
      <template #action>
        <n-flex justify="end">
          <n-button size="large" type="primary" @click="showModal = false">关闭</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped>

</style>