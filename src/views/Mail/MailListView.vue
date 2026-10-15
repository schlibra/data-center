<script setup>
import {
  createMailList,
  deleteMailList,
  getUserInfo,
  listMailDomain,
  listMailList,
  updateMailList
} from "@/api/index.js";
import {computed, h, onMounted, onUnmounted, ref} from "vue";
import {event} from "@/utils/event.js";
import {useMailDomainStore, useMailListStore, useMailMessageStore, useUserStore} from "@/stores/index.js";
import {NButton, NFlex, useDialog, useLoadingBar, useMessage} from "naive-ui";
import router from "@/router/index.js";

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()

const user = useUserStore()
const mailDomain = useMailDomainStore()
const mailList = useMailListStore()
const mailMessage = useMailMessageStore()

const columns = [
  {
    title: 'ID',
    key: 'id'
  },
  {
    title: '自增ID',
    key: 'ID'
  },
  {
    title: '邮箱',
    key: 'email'
  },
  {
    title: '用户名',
    key: 'username'
  },
  {
    title: '密码',
    key: 'password'
  },
  {
    title: '操作',
    key: 'action',
    render(row) {
      return h(NFlex, {}, [
          h(NButton, {
            type: 'primary',
            size: 'large',
            onClick() {
              openUpdateModal(row)
            }
          }, '编辑'),
          h(NButton, {
            type: 'info',
            size: 'large',
            onClick() {
              goMailMessage(row)
            }
          }, '邮件列表'),
          h(NButton, {
            type: 'error',
            size: 'large',
            onClick() {
              deleteModal(row)
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
    positiveText: '确定',
  })
  loadingBar.error()
}

const domainSelectOptions = computed(() => mailDomain.domain.map(item => {
  return {
    label: `${item.domain}（${item.name}）`,
    value: item.id
  }
}))

const showModal = ref(false)
const createMail = ref(true)
const mailId = ref('')
const mailIId = ref('')
const mailValue = ref('')
const mailUser = ref('')
const mailPassword = ref('')

function openCreateModal() {
  mailId.value = ''
  mailIId.value = ''
  mailValue.value = ''
  mailUser.value = ''
  mailPassword.value = ''
  showModal.value = true
  createMail.value = true
}

function openUpdateModal(row) {
  mailId.value = row.id
  mailIId.value = row.ID
  mailValue.value = row.email
  mailUser.value = row.username
  mailPassword.value = row.password
  showModal.value = true
  createMail.value = false
}
async function submitModal() {
  showModal.value = false
  if (createMail.value) {
    const [status, data] = await createMailList(mailList.domainId, mailValue.value, mailUser.value, mailPassword.value)
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
    const [status, data] = await updateMailList(mailList.domainId, mailId.value, mailValue.value, mailUser.value, mailPassword.value)
    if (!status) {
      dialog.error({
        title: '编辑失败',
        content: data,
        positiveText: '确定'
      })
    } else {
      message.success('编辑成功')
    }
  }
  await loadDataList()
}
function deleteModal(row) {
  dialog.info({
    title: '删除邮箱',
    content: `确定要删除邮箱 ${row.email} 吗？`,
    negativeText: '取消',
    positiveText: '确定',
    async onPositiveClick() {
      const [status, data] = await deleteMailList(mailList.domainId, row.id)
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
function goMailMessage(row) {
  mailMessage.mailRecord = row.id
  router.push('/mail/message')
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
  return true
}

async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadMailDomainList()) return
  if (!await loadMailList()) return
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
      <n-h1 prefix="bar">邮箱 邮箱列表</n-h1>
    </template>
    <n-flex>
      <n-select @click="loadDataList()" :options="domainSelectOptions" v-model:value="mailList.domainId"></n-select>
      <n-button size="large" type="primary" @click="openCreateModal()">创建邮箱</n-button>
      <n-data-table :columns="columns" :data="mailList.mail"></n-data-table>
    </n-flex>
    <n-modal v-model:show=showModal>
      <n-card size="large" style="max-width: 400px">
        <template #header>
          <n-h1>{{ createMail ? '创建邮箱' : '编辑邮箱' }}</n-h1>
        </template>
        <n-form>
          <n-form-item label="ID" v-if="!createMail">
            <n-input :value="mailId" disabled readonly></n-input>
          </n-form-item>
          <n-form-item label="自增ID" v-if="!createMail">
            <n-input-number :value="mailIId" disabled readonly></n-input-number>
          </n-form-item>
          <n-form-item label="邮箱">
            <n-input v-model:value="mailValue"></n-input>
          </n-form-item>
          <n-form-item label="用户名">
            <n-input v-model:value="mailUser"></n-input>
          </n-form-item>
          <n-form-item label="密码">
            <n-input v-model:value="mailPassword"></n-input>
          </n-form-item>
        </n-form>
        <template #action>
          <n-flex justify="end">
            <n-button size="large" @click="showModal = false">取消</n-button>
            <n-button size="large" type="primary" @click="submitModal()">确定</n-button>
          </n-flex>
        </template>
      </n-card>
    </n-modal>
  </n-card>
</template>

<style scoped>

</style>