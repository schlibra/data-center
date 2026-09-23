import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore("user", {
  state() {
    const userId = ref("")
    const username = ref("")
    const nickname = ref("")
    const groupId = ref("")
    const groupName = ref("")

    const setUserInfo = data => {
      userId.value = data.id
      username.value = data.username
      nickname.value = data.nickname
      groupId.value = data.group
      groupName.value = data.group_info.name
    }
    return {
      userId,
      username,
      nickname,
      groupId,
      groupName,
      setUserInfo,
    }
  },
  persist: true
})