<script setup>
import { onMounted } from 'vue'
import { useUserStore } from '@/stores/user.js'
import { getUserInfo } from '@/api/user.js'
import { useDialog } from 'naive-ui'
import { getFrpTokenList } from '@/api/frp/frp-token.js'
import { useFrpTokenStore } from '@/stores/frp/frp-token.js'
import { useFrpRuleStore } from '@/stores/frp/frp-rule.js'
import { getFrpRuleList } from '@/api/frp/frp-rule.js'
import { getFrpClientList, getFrpProxyList } from '@/api/frp/frp-api.js'
import { useFrpClientStore } from '@/stores/frp/frp-client.js'
import { useFrpProxyStore } from '@/stores/frp/frp-proxy.js'
import { getFrpAdminTokenList } from '@/api/frp-admin/frp-admin-token.js'
import { useFrpAdminTokenStore } from '@/stores/frp-admin/frp-admin-token.js'
import { useFrpAdminRuleStore } from '@/stores/frp-admin/frp-admin-rule.js'
import { getFrpAdminRuleList } from '@/api/frp-admin/frp-admin-rule.js'
import { useFrpAdminClientStore } from '@/stores/frp-admin/frp-admin-client.js'
import { useFrpAdminProxyStore } from '@/stores/frp-admin/frp-admin-proxy.js'
import { getFrpAdminClientList, getFrpAdminProxyList } from '@/api/frp-admin/frp-admin-api.js'
import { useAdminUserStore } from '@/stores/admin/admin-user.js'
import { getAdminUserList } from '@/api/admin/admin-user.js'

const dialog = useDialog()
const user = useUserStore()
const frpToken = useFrpTokenStore()
const frpRule = useFrpRuleStore()
const frpClient = useFrpClientStore()
const frpProxy = useFrpProxyStore()
const adminUser = useAdminUserStore()
const frpAdminToken = useFrpAdminTokenStore()
const frpAdminRule = useFrpAdminRuleStore()
const frpAdminClient = useFrpAdminClientStore()
const frpAdminProxy = useFrpAdminProxyStore()

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
  return true
}
async function loadFrpTokenList() {
  let [status, data] = await getFrpTokenList()
  if (!status) {
    return dialogError(data)
  }
  frpToken.tokens = data
  return true
}
async function loadFrpRuleList() {
  let [status, data] = await getFrpRuleList()
  if (!status) {
    return dialogError(data)
  }
  frpRule.rules = data
  return true
}
async function loadFrpClientList() {
  let [status, data] = await getFrpClientList()
  if (!status) {
    return dialogError(data)
  }
  frpClient.client = data
  return true
}
async function loadFrpProxyList() {
  let [status, data] = await getFrpProxyList()
  if (!status) {
    return dialogError(data)
  }
  frpProxy.proxy = data
  return true
}
async function loadAdminUserList() {
  let [status, data] = await getAdminUserList()
  if (!status) {
    return dialogError(data)
  }
  adminUser.users = data
  return true
}
async function loadFrpAdminTokenList() {
  let [status, data] = await getFrpAdminTokenList()
  if (!status) {
    return dialogError(data)
  }
  frpAdminToken.tokens = data
  return true
}
async function loadFrpAdminRuleList() {
  let [status, data] = await getFrpAdminRuleList()
  if (!status) {
    return dialogError(data)
  }
  frpAdminRule.rules = data
  return true
}
async function loadFrpAdminClientList() {
  let [status, data] = await getFrpAdminClientList()
  if (!status) {
    return dialogError(data)
  }
  frpAdminClient.client = data
  return true
}
async function loadFrpAdminProxyList() {
  let [status, data] = await getFrpAdminProxyList()
  if (!status) {
    return dialogError(data)
  }
  frpAdminProxy.proxy = data
  return true
}
async function loadBasicData() {
  if (!(await loadUserInfo())) return
  if (!(await loadFrpTokenList())) return
  if (!(await loadFrpRuleList())) return
  if (!(await loadFrpClientList())) return
  if (!(await loadFrpProxyList())) return
  return true
}
async function loadAdminData() {
  if (!(await loadAdminUserList())) return
  if (!(await loadFrpAdminTokenList())) return
  if (!(await loadFrpAdminRuleList())) return
  if (!(await loadFrpAdminClientList())) return
  await loadFrpAdminProxyList()
}

onMounted(async () => {
  if (!(await loadBasicData())) return
  if (user.isAdmin) {
    await loadAdminData()
  }
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
      <n-statistic label="客户端数量">
        <n-number-animation :from="0" :to="frpClient.count"></n-number-animation>
      </n-statistic>
      <n-statistic label="映射数量">
        <n-number-animation :from="0" :to="frpProxy.count"></n-number-animation>
      </n-statistic>
    </n-flex>
    <template v-if="user.isAdmin">
      <span>Frp信息（管理员）</span>
      <n-flex>
        <n-statistic label="Token数量">
          <n-number-animation :from="0" :to="frpAdminToken.count"></n-number-animation>
        </n-statistic>
        <n-statistic label="端口规则数量">
          <n-number-animation :from="0" :to="frpAdminRule.count"></n-number-animation>
        </n-statistic>
        <n-statistic label="客户端数量">
          <n-number-animation :from="0" :to="frpAdminClient.count"></n-number-animation>
        </n-statistic>
        <n-statistic label="映射数量">
          <n-number-animation :from="0" :to="frpAdminProxy.count"></n-number-animation>
        </n-statistic>
      </n-flex>
    </template>
    <template v-if="user.isAdmin">
      <span>管理员信息</span>
      <n-flex>
        <n-statistic label="用户数量">
          <n-number-animation :from="0" :to="adminUser.count"></n-number-animation>
        </n-statistic>
      </n-flex>
    </template>
  </n-card>
</template>

<style scoped></style>