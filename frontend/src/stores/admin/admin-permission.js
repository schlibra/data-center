import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useAdminPermissionStore = defineStore("admin-permission", {
  state() {
    const permissions = ref([])
    const count = computed(() => permissions.value.length)
    return {
      permissions,
      count
    }
  },
  persist: true
})