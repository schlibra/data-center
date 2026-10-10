import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useOpenVPNGroupStore = defineStore('open-vpn-group', {
  state() {
    const groups = ref([])
    const count = computed(() => groups.value.length)
    return {
      groups,
      count
    }
  },
  persist: true
})