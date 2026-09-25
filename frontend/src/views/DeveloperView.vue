<script setup>
import { onMounted, ref } from 'vue'
import { useDialog, useLoadingBar } from 'naive-ui'
import { getUserInfo } from '@/api'
import { useUserStore, useTokenStore } from '@/stores'

const loadingBar = useLoadingBar()
const dialog = useDialog()
const tokenRef = ref(null)
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
  if (!await loadUserInfo()) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>开发者功能</h3>
    </template>
    <n-form>
      <n-form-item label="用户Token">
        <n-flex>
          <n-input
            ref="tokenRef"
            :value="token.token"
            type="textarea"
            :rows="5"
            style="min-width: 600px"
            placeholder="用户Token"
          ></n-input>
          <n-button size="large" type="primary" @click="copyToken()">复制</n-button>
        </n-flex>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<style scoped></style>
