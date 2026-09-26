import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useFrpConfigStore = defineStore("frp-config", {
  state() {
    const preConfig = ref("")
    const mainConfig = ref([])
    const setAuth = (user, token) => {
      preConfig.value = `user = "${user}"
metadatas.token = "${token}"
      `
    }
    const resultConfig = computed(() => {
      let res = ""
      res += preConfig.value
      mainConfig.value.forEach(item => {
        res += `
[[proxies]]
name = "${item.name}"
type = "${item.type}"
localIP = "${item.localIP}"
localPort = "${item.localPort}"
remotePort = "${item.remotePort}"
        `
      })
      return res
    })
    const clear = () => {
      preConfig.value = ""
      mainConfig.value = []
    }
    return {
      preConfig,
      mainConfig,
      resultConfig,
      setAuth,
      clear,
    }
  },
  persist: true
})