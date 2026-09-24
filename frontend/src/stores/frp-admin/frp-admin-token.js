import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useFrpAdminTokenStore = defineStore("frp-admin-token", {
  state() {
    const tokens = ref([])
    const count = computed(() => tokens.value.length)
    return {
      tokens,
      count
    }
  },
  persist: true
})