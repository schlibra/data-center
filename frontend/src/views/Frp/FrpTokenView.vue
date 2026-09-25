<script setup>
import { h, onMounted, ref } from 'vue'
import { NButton, NFlex, NSwitch, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import router from '@/router'
import {
  createFrpToken,
  deleteFrpToken,
  generateFrpToken,
  getFrpTokenList,
  updateFrpToken,
  getUserInfo,
} from '@/api'
import { useFrpTokenStore, useFrpConfigStore, useUserStore } from '@/stores'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const user = useUserStore()
const frpToken = useFrpTokenStore()
const frpConfig = useFrpConfigStore()

const showModal = ref(false)
const createToken = ref(true)
const createTokenId = ref('')
const createTokenName = ref('')
const createTokenValue = ref('')
let tokenRow = {}

const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: '名称',
    key: 'name',
  },
  {
    title: 'Token',
    key: 'token',
  },
  {
    title: '启用',
    key: 'enable',
    render(row) {
      return h(NSwitch, {
        value: row.enable === 1,
        onUpdateValue(value) {
          changeEnable(row, value)
        },
      })
    },
  },
  {
    title: '操作',
    key: 'action',
    render(row) {
      return h(NFlex, {}, [
        h(
          NButton,
          {
            type: 'primary',
            onClick() {
              openUpdateModal(row)
            },
          },
          '编辑',
        ),
        h(
          NButton,
          {
            type: 'warning',
            onClick() {
              generateToken(row)
            },
          },
          '生成Token',
        ),
        h(
          NButton,
          {
            type: 'info',
            onClick() {
              goConfig(row)
            },
          },
          '生成配置',
        ),
        h(
          NButton,
          {
            type: 'error',
            onClick() {
              deleteToken(row)
            },
          },
          '删除',
        ),
      ])
    },
  },
]
const goConfig = (row) => {
  frpConfig.setAuth(row.name, row.token)
  router.push('/frp/config')
}
const changeEnable = async (row, value) => {
  let [status, data] = await updateFrpToken(row.id, row.name, value ? 1 : 0)
  if (!status) {
    return dialog.error({
      title: '数据修改失败',
      content: data,
      positiveText: '确定',
    })
  }
  message.success('修改成功')
  await loadDataList()
}

const dialogError = (msg) => {
  dialog.error({
    title: '数据获取失败',
    content: msg,
    positiveText: '确定',
  })
  loadingBar.error()
}

async function loadFrpTokenList() {
  let [status, data] = await getFrpTokenList()
  if (!status) {
    return dialogError(data)
  }
  frpToken.tokens = data
  return true
}

async function loadUserInfo() {
  let [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
  return true
}

async function generateToken(row) {
  dialog.info({
    title: '是否生成Token',
    content: `是否重新生成Token: "${row.name}？"`,
    negativeText: '取消',
    positiveText: '确定',
    async onPositiveClick() {
      let [status, data] = await generateFrpToken(row.id)
      if (!status) {
        dialog.error({
          title: '生成失败',
          content: data,
          negativeText: '确定',
        })
      } else {
        message.success('生成成功')
      }
      await loadDataList()
    },
  })
}

async function deleteToken(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除Token: "${row.name}？"`,
    negativeText: '取消',
    positiveText: '确定',
    async onPositiveClick() {
      let [status, data] = await deleteFrpToken(row.id)
      if (!status) {
        dialog.error({
          title: '删除失败',
          content: data,
          negativeText: '确定',
        })
      } else {
        message.success('删除成功')
      }
      await loadDataList()
    },
  })
}

function openCreateModal() {
  createTokenId.value = ''
  createTokenName.value = ''
  createTokenValue.value = ''
  createToken.value = true
  showModal.value = true
}
function openUpdateModal(row) {
  createToken.value = false
  createTokenId.value = row.id
  createTokenName.value = row.name
  createTokenValue.value = row.token
  showModal.value = true
  tokenRow = row
}

async function submitModal() {
  showModal.value = false
  if (createToken.value) {
    let [status, data] = await createFrpToken(createTokenName.value)
    if (!status) {
      return dialog.error({
        title: '创建失败',
        content: data,
        negativeText: '确定',
      })
    }
    message.success('创建成功')
  } else {
    let [status, data] = await updateFrpToken(
      createTokenId.value,
      createTokenName.value,
      tokenRow.enable,
    )
    if (!status) {
      return dialog.error({
        title: '修改失败',
        content: data,
        negativeText: '确定',
      })
    }
    message.success('修改成功')
  }
  await loadDataList()
}

async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!(await loadFrpTokenList())) return
  loadingBar.finish()
}

onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>Frp Token管理</h3>
    </template>
    <n-flex>
      <n-button type="primary" size="large" @click="openCreateModal()">创建Token</n-button>
      <n-data-table :data="frpToken.tokens" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card style="max-width: 400px">
      <template #header>
        <h3>{{ createToken ? '创建' : '编辑' }}Token</h3>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createToken">
          <n-input :value="createTokenId" disabled readonly></n-input>
        </n-form-item>
        <n-form-item label="名称">
          <n-input v-model:value="createTokenName" @keydown.enter="submitModal()"></n-input>
        </n-form-item>
        <n-form-item label="Token" v-if="!createToken">
          <n-input :value="createTokenValue" disabled readonly></n-input>
        </n-form-item>
      </n-form>
      <template #footer>
        <n-flex justify="end">
          <n-button @click="showModal = false">取消</n-button>
          <n-button type="primary" @click="submitModal()">确定</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped></style>
