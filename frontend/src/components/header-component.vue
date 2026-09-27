<script setup>
import { computed, ref } from 'vue'

import { MoonOutline as DarkIcon, SunnyOutline as LightIcon } from '@vicons/ionicons5'
import { useThemeStore } from '@/stores/theme.js'

const theme = useThemeStore()

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
  '/openvpn/user': 'OpenVPN 用户管理',
  '/openvpn/group': 'OpenVPN IP组管理',
  '/openvpn/client': 'OpenVPN 客户端管理',
  '/admin/user': '管理员 用户管理',
  '/admin/group': '管理员 用户组管理',
  '/admin/permission': '管理员 权限管理',
  '/user': '用户中心',
}
const titleText = computed(() => titleList[urlPath.value])
setInterval(() => {
  urlPath.value = location.pathname
  document.title = titleText.value
}, 100)
</script>

<template>
  <n-page-header class="header">
    <template #title>
      <h3>Data-Center</h3>
    </template>
    <template #subtitle>
      <span>{{ titleText }}</span>
    </template>
    <template #extra>
      <n-button @click="theme.switchTheme()">
        <n-icon>
          <LightIcon v-if="theme.dark"></LightIcon>
          <DarkIcon v-else></DarkIcon>
        </n-icon>
      </n-button>
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
