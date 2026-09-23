<script setup>
import { onMounted } from 'vue'
import { useUserStore } from '@/stores/user.js'
import { getUserInfo } from '@/api/user.js'
import { useDialog } from 'naive-ui'
import { getFrpTokenList } from '@/api/frp-auth/frp-token.js'
import { useFrpTokenStore } from '@/stores/frp-auth/frp-token.js'
import { useFrpRuleStore } from '@/stores/frp-auth/frp-rule.js'
import { getFrpRuleList } from '@/api/frp-auth/frp-rule.js'
import { getFrpClientList, getFrpProxyList } from '@/api/frp-auth/frp-api.js'
import { useFrpClientStore } from '@/stores/frp-auth/frp-client.js'
import { useFrpProxyStore } from '@/stores/frp-auth/frp-proxy.js'

const dialog = useDialog()
const user = useUserStore()
const frpToken = useFrpTokenStore()
const frpRule = useFrpRuleStore()
const frpClient = useFrpClientStore()
const frpProxy = useFrpProxyStore()

const dialogError = (msg) => {
  dialog.error({
    title: '数据获取失败',
    content: msg,
    positiveText: '确定',
  })
}

async function loadUserInfo() {
  let [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
}
async function loadFrpTokenList() {
  let [status, data] = await getFrpTokenList()
  if (!status) {
    return dialogError(data)
  }
  frpToken.tokens = data
}
async function loadFrpRuleList() {
  let [status, data] = await getFrpRuleList()
  if (!status) {
    return dialogError(data)
  }
  frpRule.rules = data
}
async function loadFrpClientList() {
  let [status, data] = await getFrpClientList()
  if (!status) {
    return dialogError(data)
  }
  frpClient.client = data
}
async function loadFrpProxyList() {
  let [status, data] = await getFrpProxyList()
  if (!status) {
    return dialogError(data)
  }
  frpProxy.proxy = data
}

onMounted(async () => {
  await loadUserInfo()
  await loadFrpTokenList()
  await loadFrpRuleList()
  await loadFrpClientList()
  await loadFrpProxyList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>首页</h3>
    </template>
    <span>用户信息</span>
    <n-flex>
      <n-statistic label="用户ID" :value="user.userId"></n-statistic>
      <n-statistic label="用户名" :value="user.username"></n-statistic>
      <n-statistic label="昵称" :value="user.nickname"></n-statistic>
      <n-statistic label="用户组名称" :value="user.groupName"></n-statistic>
      <n-statistic label="用户组ID" :value="user.groupId"></n-statistic>
    </n-flex>
    <span>Frp信息</span>
    <n-flex>
      <n-statistic label="Token数量">
        <n-number-animation :from="0" :to="frpToken.count"></n-number-animation>
      </n-statistic>
      <n-statistic label="端口规则数量">
        <n-number-animation :from="0" :to="frpRule.count"></n-number-animation>
      </n-statistic>
      <n-statistic label="客户端数量" :value="frpClient.count">
        <n-number-animation :from="0" :to="frpClient.count"></n-number-animation>
      </n-statistic>
      <n-statistic label="映射数量">
        <n-number-animation :from="0" :to="frpProxy.count"></n-number-animation>
      </n-statistic>
    </n-flex>
  </n-card>
</template>

<style scoped></style>