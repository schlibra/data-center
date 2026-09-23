import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useAdminUserStore = defineStore("admin-user", {
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