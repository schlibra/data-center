<script setup>
import {onMounted} from "vue";
import {getAutoLoginLink, getAutoLoginList, getUserInfo} from "@/api/index.js";
import {useDialog, useLoadingBar} from "naive-ui";
import {useAutoLoginStore, useUserStore} from "@/stores/index.js";

const dialog = useDialog()
const loadingBar = useLoadingBar()

const user = useUserStore()
const autoLogins = useAutoLoginStore()

const dialogError = content => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定'
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
async function loadAutoLoginList() {
  const [status, data] = await getAutoLoginList()
  if (!status) {
    return dialogError(data)
  }
  autoLogins.autoLogins = data
  return true
}

async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadAutoLoginList()) return
  loadingBar.finish()
}

async function openLogin(item) {
  const [status, data] = await getAutoLoginLink(item.id)
  if (!status) {
    dialog.error({
      title: '链接获取失败',
      content: data,
      positiveText: '确定'
    })
  } else {
    const targetWindow = window.open(data.url, '_blank')
    let timer = setInterval(() => {
      targetWindow.postMessage({autologin: data.data}, data.url)
    }, 100)
    window.addEventListener("message", e => {
	  if (e.data.autoLoginReady) {
		clearInterval(timer)
	  }
    })
  }
}

onMounted(() => {
  loadDataList()
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">快速登录</n-h1>
    </template>
    <n-grid>
      <n-grid-item v-for="item in autoLogins.autoLogins" :key="item.id">
        <n-button size="large" style="height: 80px!important;" @click="openLogin(item)">
          <n-space justify="center">
            <n-avatar :src="item.icon"></n-avatar>
            <n-text>{{item.name}}</n-text>
          </n-space>
        </n-button>
      </n-grid-item>
    </n-grid>
  </n-card>
</template>

<style scoped>

</style>