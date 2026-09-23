import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useFrpClientStore = defineStore('frp-client', {
  state() {
    const client = ref([])
    const count = computed(() => client.value.length)
    return {
      client,
      count,
    }
  },
  persist: true
})
