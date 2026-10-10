<script setup>
import {computed, onMounted, ref} from "vue";
import {useDialog, useLoadingBar, useMessage} from "naive-ui";
import {getConfig} from "@/api/config.js";
import {useAdminGroupStore} from "@/stores/index.js";
import {getAdminGroupList, setAdminSettings} from "@/api/index.js";

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const adminGroup = useAdminGroupStore()
const groupOptions = computed(() => adminGroup.groups.map(item => {
  return {
    label: item.name,
    value: item.id,
  }
}))

const siteName = ref('')
const defaultGroup = ref('')
const registerConfirm = ref(false)

const dialogError = content => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}
const dialogChangeError = content => {
  dialog.error({
    title: '数据更新失败',
    content,
    positiveText: '确定',
  })
}
async function loadSiteName() {
  let [status, data] = await getConfig('site.name')
  siteName.value = status && data ? data : 'Data-Center'
}
async function loadDefaultGroup() {
  let [status, data] = await getConfig('group.default')
  try {
    data = parseInt(data)
  } catch (e) {}
  defaultGroup.value = status ? data : 1
}
async function loadRegisterConfirm() {
  let [status, data] = await getConfig('register.confirm')
  if (!status) {
    registerConfirm.value = true
  } else {
    registerConfirm.value = data === "1"
  }
}
async function loadGroupList() {
  let [status, data] = await getAdminGroupList()
  if (!status) {
    return dialogError(data)
  }
  adminGroup.groups = data
  return true
}
async function changeSiteName() {
  let [status, data] = await setAdminSettings('site.name', '站点名称', siteName.value)
  if (!status) {
    dialogChangeError(data)
  } else {
    message.success('修改成功')
  }
  await loadDataList()
}
async function changeDefaultGroup() {
  let [status, data] = await setAdminSettings('group.default', '默认用户组', defaultGroup.value+"")
  if (!status) {
    dialogChangeError(data)
  } else {
    message.success('修改成功')
  }
  await loadDataList()
}
async function changeRegisterConfirm(value) {
  let [status, data] = await setAdminSettings('register.confirm', '注册需要确认', value ? "1" : "0")
  if (!status) {
    dialogChangeError(data)
  } else {
    message.success('修改成功')
  }
  await loadDataList()
}
async function loadDataList() {
  loadingBar.start()
  await loadSiteName()
  if (!await loadGroupList()) return
  await loadDefaultGroup()
  await loadRegisterConfirm()
  loadingBar.finish()
}

onMounted(() => {
  loadDataList()
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">管理员 系统设置</n-h1>
    </template>
    <n-form>
      <n-form-item label="站点名称">
        <n-space>
          <n-input style="min-width: 300px" v-model:value="siteName"></n-input>
          <n-button @click="changeSiteName" type="primary">保存</n-button>
        </n-space>
      </n-form-item>
      <n-form-item label="默认用户组">
        <n-space>
          <n-select style="min-width: 300px" :options="groupOptions" v-model:value="defaultGroup"></n-select>
          <n-button @click="changeDefaultGroup" type="primary">保存</n-button>
        </n-space>
      </n-form-item>
      <n-form-item label="注册需要确认">
        <n-switch @update:value="changeRegisterConfirm" :value="registerConfirm"></n-switch>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<style scoped>

</style>