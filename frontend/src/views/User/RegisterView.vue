<script setup>
import { ref } from 'vue'
import { registerKey, registerUser } from '@/api/user.js'
import { useDialog, useMessage } from 'naive-ui'
import encrypt from '@/utils/encrypt.js'
import router from '@/router/index.js'

const dialog = useDialog()
const message = useMessage()

const usernameRef = ref(null)
const passwordRef = ref(null)
const confirmPasswordRef = ref(null)
const nicknameRef = ref(null)

const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const nickname = ref('')

const doRegister = async () => {
  if (password.value !== confirmPassword.value) {
    return dialog.error({
      title: '注册失败',
      content: '两次密码不一致',
      positiveText: '确定',
      onPositiveClick() {
        usernameRef.value.focus()
      },
    })
  }
  let [status, data] = await registerKey(username.value)
  if (!status) {
    return dialog.error({
      title: '注册失败',
      content: data,
      positiveText: '确定',
      onPositiveClick() {
        usernameRef.value.focus()
      },
    })
  }
  ;[status, data] = encrypt(password.value, data)
  if (!status) {
    return dialog.error({
      title: '注册失败',
      content: data,
      positiveText: '确定',
      onPositiveClick() {
        usernameRef.value.focus()
      },
    })
  }
  ;[status, data] = await registerUser(username.value, data, nickname.value)
  if (!status) {
    return dialog.error({
      title: '注册失败',
      content: data,
      positiveText: '确定',
      onPositiveClick() {
        usernameRef.value.focus()
      },
    })
  }
  message.success("注册成功")
  setTimeout(() => {
    router.push('/login')
  }, 1500)
}
</script>

<template>
  <n-card size="large" style="max-width: 600px">
    <template #header>
      <span>注册账号</span>
    </template>
    <n-form>
      <n-form-item label="用户名">
        <n-input
          v-model:value="username"
          ref="usernameRef"
          @keydown.enter="passwordRef.focus()"
        ></n-input>
      </n-form-item>
      <n-form-item label="密码">
        <n-input
          ref="passwordRef"
          type="password"
          v-model:value="password"
          @keydown.enter="confirmPasswordRef.focus()"
        ></n-input>
      </n-form-item>
      <n-form-item label="确认密码">
        <n-input
          ref="confirmPasswordRef"
          type="password"
          v-model:value="confirmPassword"
          @keydown.enter="nicknameRef.focus()"
        ></n-input>
      </n-form-item>
      <n-form-item label="昵称">
        <n-input ref="nicknameRef" v-model:value="nickname" @keydown.enter="doRegister()"></n-input>
      </n-form-item>
    </n-form>
    <template #footer>
      <n-button @click="doRegister()" size="large" type="primary">注册</n-button>
    </template>
  </n-card>
</template>

<style scoped></style>
