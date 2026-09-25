<script setup>
import { useDialog, useLoadingBar, useMessage } from 'naive-ui'
import { onMounted, ref } from 'vue'
import router from '@/router'
import { useTokenStore, useUserStore } from '@/stores'
import { getUserInfo, logoutUser, setUserPassword, updateUser } from '@/api'

const loadingBar = useLoadingBar()
const dialog = useDialog()
const message = useMessage()
const user = useUserStore()
const token = useTokenStore()
const password = ref("")

const dialogError = (data) => {
  dialog.error({
    title: '数据获取失败',
    content: data,
    positiveText: '确定',
  })
  loadingBar.error()
}
async function updateUserInfo() {
  let [status, data] = await updateUser(user.nickname)
  if (!status) {
    return dialog.error({
      title: "修改失败",
      content: data,
      positiveText: "确定"
    })
  }
  message.success("修改成功")
  await loadDataList()
}
async function updateUserPassword() {
  let [status, data] = await setUserPassword(password.value)
  if (!status) {
    return dialog.error({
      title: "修改失败",
      content: data,
      positiveText: "确定"
    })
  }
  dialog.success({
    title: "修改成功",
    content: data,
    positiveText: "去登录",
    onPositiveClick() {
      router.push("/login")
    },
    closable: false,
    closeFocusable: false,
    closeOnEsc: false
  })
}
async function loadUserInfo() {
  let [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
  return true
}
async function logout() {
  dialog.info({
    title: "是否退出",
    content: "是否退出登录",
    positiveText: "确定",
    negativeText: "取消",
    async onPositiveClick() {
      await logoutUser()
      token.token = ""
      router.push("/login")
    }
  })
}
async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>用户中心</h3>
    </template>
    <n-form>
      <n-form-item label="用户ID">
        <n-input :value="user.userId" readonly disabled></n-input>
      </n-form-item>
      <n-form-item label="用户名">
        <n-input :value="user.username" readonly disabled></n-input>
      </n-form-item>
      <n-form-item label="昵称">
        <n-flex :wrap="false">
          <n-input v-model:value="user.nickname" @keydown.enter="updateUserInfo()"></n-input>
          <n-button type="primary" @click="updateUserInfo()">修改昵称</n-button>
        </n-flex>
      </n-form-item>
      <n-form-item label="用户组">
        <n-input :value="user.groupName" readonly disabled></n-input>
      </n-form-item>
      <n-form-item label="密码">
        <n-flex :wrap="false">
          <n-input v-model:value="password" type="password" placeholder="设置密码" @keydown.enter="updateUserPassword()"></n-input>
          <n-button type="warning" @click="updateUserPassword()">修改密码</n-button>
        </n-flex>
      </n-form-item>
    </n-form>
    <template #footer>
      <n-button type="error" @click="logout()">退出登录</n-button>
    </template>
  </n-card>
</template>

<style scoped></style>
