<script setup>
import {onMounted} from "vue";
import {useDialog, useLoadingBar, useMessage} from "naive-ui";
import {useAdminVersionStore, useUserStore} from "@/stores/index.js";
import {getAdminVersion, getUserInfo, upgradeAdminVersion} from "@/api/index.js";

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const user = useUserStore()
const adminVersion = useAdminVersionStore()

const dialogError = content => {
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

async function loadAdminVersion() {
  let [status, data] = await getAdminVersion()
  if (!status) {
    return dialogError(data)
  }
  adminVersion.set(data)
  return true
}

async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadAdminVersion()) return
  loadingBar.finish()
}
async function doUpgrade() {
  let [status, data] = await upgradeAdminVersion()
  if (!status) {
    dialog.error({
      title: '版本更新失败',
      content: data,
      positiveText: '确定',
    })
  } else {
    message.success('版本更新成功')
  }
}
onMounted(() => {
  loadDataList()
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">管理员 版本管理</n-h1>
    </template>
    <n-descriptions>
      <n-descriptions-item label="版本号">
        <n-text>{{ adminVersion.version }}</n-text>
      </n-descriptions-item>
      <n-descriptions-item label="提交Hash">
        <n-text>{{ adminVersion.commit }}</n-text>
      </n-descriptions-item>
      <n-descriptions-item label="构建时间">
        <n-text>{{ adminVersion.buildTime }}</n-text>
      </n-descriptions-item>
    </n-descriptions>
    <template #action>
      <n-flex>
        <n-button size="large" type="warning" @click="doUpgrade()">更新版本</n-button>
      </n-flex>
    </template>
  </n-card>
</template>

<style scoped>

</style>