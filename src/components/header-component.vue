<script setup>
import {computed, onMounted, ref} from 'vue'
import logo from '@/assets/logo.svg'

import {
  MoonOutline as DarkIcon,
  SunnyOutline as LightIcon,
  RefreshSharp as RefreshIcon,
} from '@vicons/ionicons5'
import { useThemeStore } from '@/stores/theme.js'
import {getConfig} from "@/api/config.js";
import {event} from "@/utils/event.js";

const theme = useThemeStore()
const siteName = ref('')

const urlPath = ref(location.pathname)
const titleList = {
  '/': '首页',
  '/login': '登录',
  '/register': '注册',
  '/developer': '开发者功能',
  '/frp/token': 'Frp Token管理',
  '/frp/rule': 'Frp 端口规则管理',
  '/frp/client': 'Frp 客户端管理',
  '/frp/proxy': 'Frp 映射管理',
  '/frp/config': 'Frp 配置生成',
  '/frp/admin/token': 'Frp Token管理（管理员）',
  '/frp/admin/rule': 'Frp 端口规则管理（管理员）',
  '/frp/admin/client': 'Frp 客户端管理（管理员）',
  '/frp/admin/proxy': 'Frp 映射管理（管理员）',
  '/mail/domain': '邮箱 域名管理',
  '/mail/list': '邮箱 邮箱列表',
  '/mail/message': '邮箱 邮件管理',
  '/openvpn/user': 'OpenVPN 用户管理',
  '/openvpn/group': 'OpenVPN IP组管理',
  '/openvpn/client': 'OpenVPN 客户端管理',
  '/admin/user': '管理员 用户管理',
  '/admin/group': '管理员 用户组管理',
  '/admin/permission': '管理员 权限管理',
  '/admin/system': '管理员 系统设置',
  '/admin/settings': '管理员 所有设置',
  '/admin/version': '管理员 版本管理',
  '/user': '用户中心',
}
const titleText = computed(() => titleList[urlPath.value])
const showRefresh = computed(() => urlPath.value !== '/login' && urlPath.value !== '/register' && urlPath.value !== '/404')
setInterval(() => {
  urlPath.value = location.pathname
  document.title = titleText.value
}, 100)
const handleRefresh = () => {
  event.emit('trigger-refresh')
  event.emit('trigger-menu-refresh')
  loadDataList()
}
async function loadSiteName() {
  let [status, data] = await getConfig("site.name")
  siteName.value = status && data ? data : 'Data-Center'
}

async function loadDataList() {
  await loadSiteName()
}

onMounted(() => {
  loadDataList()
})
</script>

<template>
  <n-page-header class="header">
    <template #avatar>
      <n-avatar size="large" color="#0000" :src="logo"></n-avatar>
    </template>
    <template #title>
      <h3>{{ siteName }}</h3>
    </template>
    <template #subtitle>
      <span>{{ titleText }}</span>
    </template>
    <template #extra>
      <n-flex>
        <n-button @click="handleRefresh" v-if="showRefresh">
          <n-icon>
            <RefreshIcon></RefreshIcon>
          </n-icon>
        </n-button>
        <n-button @click="theme.switchTheme()">
          <n-icon>
            <LightIcon v-if="theme.dark"></LightIcon>
            <DarkIcon v-else></DarkIcon>
          </n-icon>
        </n-button>
      </n-flex>
    </template>
  </n-page-header>
</template>

<style scoped>
.header {
  padding-left: 16px;
}
</style>
<style>
.n-page-header__extra {
  padding-right: 16px;
}
</style>
