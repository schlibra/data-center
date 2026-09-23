<script setup>
import { onMounted, ref } from 'vue'
import { loginKey, loginUser } from '@/api/user.js'
import { useDialog, useMessage } from 'naive-ui'
import router from '@/router/index.js'
import encrypt from '@/utils/encrypt.js'

const message = useMessage()
const dialog = useDialog()

const usernameRef = ref(null)
const passwordRef = ref(null)

const username = ref('')
const password = ref('')

const dialogError = msg => {
  dialog.error({
    title: '登录失败',
    content: msg,
    positiveText: "确定",
    onPositiveClick() {
      usernameRef.value.focus()
    }
  })
}

const doLogin = async () => {
  let [status, data] = await loginKey(username.value)
  if (!status) {
    dialogError(data)
    return
  }
  ;[status, data] = encrypt(password.value, data)
  if (!status) {
    dialogError(data)
    return
  }
  ;[status, data] = await loginUser(username.value, data)
  if (status) {
    message.success('登录成功')
    setTimeout(() => {
      router.push("/")
    }, 1500)
  } else {
    dialogError(data)
  }
}
onMounted(() => {
  usernameRef.value.focus()
})
</script>

<template>
  <n-card style="max-width: 400px; margin-top: 10%;">
    <template #header>
      <h3>登录账号</h3>
    </template>
    <n-form>
      <n-form-item label="用户名">
        <n-input v-model:value="username" ref="usernameRef" @keydown.enter="passwordRef.focus()"></n-input>
      </n-form-item>
      <n-form-item label="密码">
        <n-input v-model:value="password" type="password" ref="passwordRef" @keydown.enter="doLogin()"></n-input>
      </n-form-item>
    </n-form>
    <template #footer>
      <n-flex justify="center">
        <n-button size="large" type="primary" @click="doLogin()">登录</n-button>
        <n-button size="large" type="info" @click="router.push('/register')">去注册</n-button>
      </n-flex>
    </template>
  </n-card>
</template>

<style scoped></style>
