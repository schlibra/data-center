<script setup>
import { computed, ref } from 'vue'

const urlPath = ref(location.pathname)
const titleList = {
  "/": "首页",
  "/login": "登录",
  "/register": "注册",
  "/frp/token": "Frp Token管理",
  "/frp/rule": "Frp 端口规则管理",
  "/frp/client": "Frp 客户端管理",
  "/frp/proxy": "Frp 映射管理",
  "/frp/config": "Frp 配置生成",
  "/admin/user": "管理员 用户管理",
  "/user": "用户中心"
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
  </n-page-header>
</template>

<style scoped>
.header {
  padding-left: 16px;
}
</style>
