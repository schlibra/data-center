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
    return {
      token,
      config
    }
  },
  persist: true
})