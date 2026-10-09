import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useFrpConfigStore = defineStore("frp-config", {
  state() {
    const info = ref({})
    const preConfig = ref("")
    const mainConfig = ref([])
    const setAuth = (user, token) => {
      preConfig.value = `serverHost = "${info.value.host}"
serverPort = ${info.value.port}      
user = "${user}"
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
localPort = ${item.localPort}
remotePort = ${item.remotePort}
`
      })
      return res
    })
    const clear = () => {
      preConfig.value = ""
      mainConfig.value = []
    }
    return {
      info,
      preConfig,
      mainConfig,
      resultConfig,
      setAuth,
      clear,
    }
  },
  persist: true
})