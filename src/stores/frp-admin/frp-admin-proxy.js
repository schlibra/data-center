import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useFrpAdminProxyStore = defineStore("frp-admin-proxy", {
  state() {
    const proxy = ref([])
    const count = computed(() => proxy.value.length)
    return {
      proxy,
      count
    }
  },
  persist: true
})