import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useOpenVPNUserStore = defineStore('open-vpn-user', {
  state() {
    const users = ref([])
    const count = computed(() => users.value.length)
    return {
      users,
      count
    }
  },
  persist: true
})