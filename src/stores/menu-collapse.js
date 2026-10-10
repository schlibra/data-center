import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMenuCollapseStore = defineStore("menu-collapse", {
  state() {
    const collapse = ref(true)
    return {
      collapse
    }
  },
  persist: true
})