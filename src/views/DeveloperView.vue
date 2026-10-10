<script setup>
import { onMounted, ref } from 'vue'
import { useDialog, useLoadingBar, useMessage } from 'naive-ui'
import {generateUserApiKey, getUserApiKey, getUserInfo} from '@/api'
import { useUserStore, useTokenStore } from '@/stores'
import router from '@/router'

const loadingBar = useLoadingBar()
const message = useMessage()
const dialog = useDialog()
const tokenRef = ref(null)
const apiKeyRef = ref(null)
const user = useUserStore()
const token = useTokenStore()

const dialogError = (content) => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}

const copyToken = () => {
  tokenRef.value.select()
  document.execCommand('copy')
}
const copyApiKey = () => {
  apiKeyRef.value.select()
  document.execCommand('copy')
}
function generateApiKey() {
  dialog.info({
    title: "是否生成API Key",
    content: '是否生成API Key，之前生成的API Key将失效',
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await generateUserApiKey()
      if (!status) {
        dialog.error({
          title: 'API Key生成失败',
          content: data,
          positiveText: '确定'
        })
      } else {
        message.success('API Key生成成功')
        token.apiKey = data.token
      }
    }
  })
}
async function getApiKey() {
  let [status, data] = await getUserApiKey()
  if (!status) {
    dialogError(data)
  } else {
    message.success('API Key获取成功')
    token.apiKey = data.token
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
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  loadingBar.finish()
}
function clearData() {
  dialog.info({
    title: '是否清空数据',
    content: '清空数据后将需要重新登录',
    positiveText: '确定',
    negativeText: '取消',
    positiveButtonProps: {
      type: 'error',
    },
    onPositiveClick() {
      localStorage.clear()
      message.success('数据清空成功')
      router.push('/login')
    },
  })
}
function clearCache() {
  caches.keys().then(
      keys => keys.forEach(item => {
        caches.delete(item)
      })
  )
  message.success('缓存清空成功')
  location.reload()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">开发者功能</n-h1>
    </template>
    <n-form>
      <n-form-item label="用户Token">
        <n-flex>
          <n-input
            ref="tokenRef"
            v-model:value="token.token"
            type="textarea"
            :rows="5"
            style="min-width: 600px"
            placeholder="用户Token"
          ></n-input>
          <n-button size="large" type="primary" @click="copyToken()">复制</n-button>
        </n-flex>
      </n-form-item>
      <n-form-item label="API Key">
        <n-space vertical>
          <n-input style="min-width: 600px" :rows="5" ref="apiKeyRef" type="textarea" :value="token.apiKey" placeholder="API Key"></n-input>
          <n-flex>
            <n-button type="warning" size="large" @click="generateApiKey()">生成Key</n-button>
            <n-button type="info" size="large" @click="getApiKey()">获取Key</n-button>
            <n-button type="primary" size="large" @click="copyApiKey()">复制</n-button>
          </n-flex>
        </n-space>
      </n-form-item>
      <n-form-item label="清空缓存">
        <n-space vertical>
          <n-text>用于在网站更新时清理缓存以同步最新版本页面</n-text>
          <n-button type="warning" size="large" @click="clearCache()">清空缓存</n-button>
        </n-space>
      </n-form-item>
      <n-form-item label="清空数据">
        <n-space vertical>
          <n-text style="color: gray"
            >这个操作将会清空该站点在本地存储的所有数据，可能可以解决一些问题，清空后需要重新登录</n-text
          >
          <n-button type="error" size="large" @click="clearData()">清空数据</n-button>
        </n-space>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<style scoped></style>
