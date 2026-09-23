<script setup>
import { darkTheme, zhCN } from 'naive-ui'
import HeaderComponent from '@/components/header-component.vue'
import MenuComponent from '@/components/menu-component.vue'
import { computed, ref } from 'vue'

const noMenuList = ['/login', '/register', '/404']
const urlPath = ref(location.pathname)

const showMenu = computed(() => noMenuList.indexOf(urlPath.value) === -1)
setInterval(() => {
  urlPath.value = location.pathname
}, 100)
</script>

<template>
  <n-config-provider :theme="darkTheme" :locale="zhCN">
    <n-message-provider>
      <n-dialog-provider>
        <n-modal-provider>
          <n-layout class="main">
            <n-layout-header>
              <header-component></header-component>
            </n-layout-header>
            <n-layout :has-sider="showMenu">
              <menu-component v-if="showMenu"></menu-component>
              <n-layout-content>
                <n-flex align="center" justify="center">
                  <RouterView />
                </n-flex>
              </n-layout-content>
            </n-layout>
          </n-layout>
        </n-modal-provider>
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>
.main {
  height: 100vh;
}
</style>
