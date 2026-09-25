import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useThemeStore = defineStore("theme", {
  state() {
    const dark = ref(true)
    const switchTheme = () => {
      dark.value = !dark.value
    }
    return {
      dark,
      switchTheme
    }
  },
  persist: true
})