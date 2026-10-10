import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useFrpProxyStore = defineStore("frp-proxy", {
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