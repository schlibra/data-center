import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useTokenStore = defineStore("token", {
  state() {
    const token = ref("")
    const config = computed(() => {
      return {
        headers: {
          Authorization: `Bearer ${token.value}`,
        },
      }
    })
    const apiKey = ref("")
    return {
      token,
      config,
      apiKey
    }
  },
  persist: true
})