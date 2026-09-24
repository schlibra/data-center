import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useFrpTokenStore = defineStore("frp-token", {
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