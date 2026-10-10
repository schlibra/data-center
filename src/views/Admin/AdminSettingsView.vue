<script setup>
import {createAdminSettings, deleteAdminSettings, getUserInfo, listAdminSettings, updateAdminSettings} from '@/api'
import {useAdminSettingsStore, useUserStore} from '@/stores'
import {
  useDialog,
  useLoadingBar,
  NCard,
  NInput,
  NButton,
  NModal,
  NForm,
  NFormItem,
  NH1,
  NDataTable,
  NFlex,
  useMessage
} from "naive-ui";
import {h, onMounted, ref} from "vue";

const loadingBar = useLoadingBar()
const dialog = useDialog()
const message = useMessage()
const user = useUserStore()
const adminSettings = useAdminSettingsStore()

const showModal = ref(false)
const createSettings = ref(false)
const settingsId = ref("")
const settingsKey = ref("")
const settingsName = ref("")
const settingsValue = ref("")

const columns = [
  {
    title: 'ID',
    key: 'id'
  },
  {
    title: '设置键名',
    key: 'key'
  },
  {
    title: '设置名称',
    key: 'name'
  },
  {
    title: '设置值',
    key: 'value'
  },
  {
    title: '操作',
    key: 'action',
    render(row) {
      return h(NFlex, {}, [
          h(NButton, {
            type: 'primary',
            size: 'large',
            onClick() {
              openUpdateModal(row)
            }
          }, '编辑'),
          h(NButton, {
            type: 'error',
            size: 'large',
            onClick() {
              deleteSettings(row)
            }
          }, '删除'),
      ])
    }
  }
]

const dialogError = content => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}

function openCreateModal() {
  settingsId.value = ""
  settingsKey.value = ""
  settingsName.value = ""
  settingsValue.value = ""
  createSettings.value = true
  showModal.value = true
}
function openUpdateModal(row) {
  settingsId.value = row.id
  settingsKey.value = row.key
  settingsName.value = row.name
  settingsValue.value = row.value
  createSettings.value = false
  showModal.value = true
}
async function submitModal() {
  showModal.value = false
  let status, data
  if (createSettings.value) {
    [status, data] = await createAdminSettings(settingsKey.value, settingsName.value, settingsValue.value)
  } else {
    [status, data] = await updateAdminSettings(settingsId.value, settingsName.value, settingsValue.value)
  }
  const actionText = createSettings.value ? '创建' : '编辑'
  if (!status) {
    dialog.error({
      title: `${actionText}失败`,
      content: data,
      positiveText: '确定',
    })
  } else {
    message.success(`${actionText}成功`)
  }
  await loadDataList()
}
function deleteSettings(row) {
  dialog.info({
    title: '确认删除',
    content: `确认删除设置 ${row.name} 吗？`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await deleteAdminSettings(row.id)
      if (!status) {
        dialog.error({
          title: '删除失败',
          content: data,
          positiveText: '确定'
        })
      } else {
        message.success(`删除设置 ${row.name} 成功`)
      }
      await loadDataList()
    }
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
async function loadSettingsList() {
  let [status, data] = await listAdminSettings()
  if (!status) {
    return dialogError(data)
  }
  adminSettings.settings = data
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadSettingsList()) return
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
    <n-flex>
      <n-button size="large" type="primary" @click="openCreateModal()">创建设置</n-button>
      <n-data-table :columns="columns" :data="adminSettings.settings"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card size="large" style="max-width: 400px">
      <template #header>
        <n-h1>{{createSettings ? '创建' : '编辑'}}设置</n-h1>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createSettings">
          <n-input v-model:value="settingsId" readonly disabled></n-input>
        </n-form-item>
        <n-form-item label="设置键名">
          <n-input v-model:value="settingsKey" :disabled="!createSettings" :readonly="!createSettings"></n-input>
        </n-form-item>
        <n-form-item label="设置名称">
          <n-input v-model:value="settingsName"></n-input>
        </n-form-item>
        <n-form-item label="设置值">
          <n-input v-model:value="settingsValue"></n-input>
        </n-form-item>
      </n-form>
      <template #action>
        <n-flex justify="end">
          <n-button size="large" @click="showModal = false">取消</n-button>
          <n-button type="primary" size="large" @click="submitModal()">确定</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped>

</style>