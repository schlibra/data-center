<script setup>
import { onMounted } from 'vue'
import { useDialog } from 'naive-ui'
import {
  useAdminGroupStore,
  useAdminPermissionStore,
  useAdminUserStore,
  useFrpAdminClientStore,
  useFrpAdminProxyStore,
  useFrpAdminRuleStore,
  useFrpAdminTokenStore,
  useFrpClientStore,
  useFrpProxyStore,
  useFrpRuleStore,
  useFrpTokenStore,
  useUserStore,
} from '@/stores'
import {
  getAdminGroupList,
  getAdminPermissionList,
  getAdminUserList,
  getFrpAdminClientList,
  getFrpAdminProxyList,
  getFrpAdminRuleList,
  getFrpAdminTokenList,
  getFrpClientList,
  getFrpProxyList,
  getFrpRuleList,
  getFrpTokenList,
  getUserInfo,
} from '@/api'

const dialog = useDialog()
const user = useUserStore()
const frpToken = useFrpTokenStore()
const frpRule = useFrpRuleStore()
const frpClient = useFrpClientStore()
const frpProxy = useFrpProxyStore()
const adminUser = useAdminUserStore()
const adminGroup = useAdminGroupStore()
const adminPermission = useAdminPermissionStore()
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
async function loadAdminGroupList() {
  let [status, data] = await getAdminGroupList()
  if (!status) {
    return dialogError(data)
  }
  adminGroup.groups = data
  return true
}
async function loadAdminPermissionList() {
  let [status, data] = await getAdminPermissionList()
  if (!status) {
    return dialogError(data)
  }
  adminPermission.permissions = data
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
  if (!(await loadAdminGroupList())) return
  if (!(await loadAdminPermissionList())) return
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
        <n-statistic label="用户组数量">
          <n-number-animation :from="0" :to="adminGroup.count"></n-number-animation>
        </n-statistic>
        <n-statistic label="用户权限数量">
          <n-number-animation :from="0" :to="adminPermission.count"></n-number-animation>
        </n-statistic>
      </n-flex>
    </template>
  </n-card>
</template>

<style scoped></style>