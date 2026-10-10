import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useUserStore = defineStore("user", {
  state() {
    const userId = ref("")
    const username = ref("")
    const nickname = ref("")
    const groupId = ref("")
    const groupName = ref("")
    const isAdmin = ref(false)
    const permissionList = ref([])
    const permissionKeys = computed(() => permissionList.value.map(item => {
      return item.key
    }))
    const hasPermission = key => {
      return permissionKeys.value.indexOf(key) > -1 || isAdmin.value
    }

    const setUserInfo = data => {
      userId.value = data.id
      username.value = data.username
      nickname.value = data.nickname
      groupId.value = data.group
      groupName.value = data.group_info.name
      isAdmin.value = data.group_info.admin === 1
      permissionList.value = data.permissions
    }
    return {
      userId,
      username,
      nickname,
      groupId,
      groupName,
      isAdmin,
      permissionList,
      permissionKeys,
      setUserInfo,
      hasPermission
    }
  },
  persist: true
})