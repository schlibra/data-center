<script setup>
import { NButton, NFlex, NSwitch, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import {
  createFrpAdminToken,
  deleteFrpAdminToken,
  generateFrpAdminToken,
  getAdminUserList,
  getFrpAdminTokenList,
  getUserInfo,
  updateFrpAdminToken,
} from '@/api'
import { useAdminUserStore, useFrpAdminTokenStore, useUserStore } from '@/stores/'
import { computed, h, onMounted, ref, useId } from 'vue'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const frpAdminToken = useFrpAdminTokenStore()
const adminUser = useAdminUserStore()
const user = useUserStore()

const showModal = ref(false)
const createToken = ref(true)
const tokenId = ref('')
const tokenName = ref('')
const tokenUser = ref('')
const tokenEnable = ref(true)
const tokenValue = ref('')

const userSelectOption = computed(() =>
  adminUser.users.map((item) => {
    return {
      label: `${item.username}(${item.nickname})`,
      value: item.id,
    }
  }),
)
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
    title: '用户',
    key: 'user',
    render(row) {
      return h('span', {}, `${row.user_info.username} ( ${row.user_info.nickname} )`)
    },
  },
  {
    title: '启用',
    key: 'enable',
    render(row) {
      return h(NSwitch, {
        value: row.enable === 1,
        onUpdateValue(value) {
          changeEnable(row, value)
        }
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
            }
          },
          '生成Token',
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

const dialogError = (content) => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}
function openCreateModal() {
  tokenId.value = ''
  tokenName.value = ''
  tokenUser.value = ''
  tokenValue.value = ''
  tokenEnable.value = true
  createToken.value = true
  showModal.value = true
}
function openUpdateModal(row) {
  tokenId.value = row.id
  tokenName.value = row.name
  tokenUser.value = row.user
  tokenValue.value = row.token
  tokenEnable.value = row.enable === 1
  createToken.value = false
  showModal.value = true
}
async function submitModal() {
  showModal.value = false
  if (createToken.value) {
    let [status, data] = await createFrpAdminToken(tokenName.value, tokenUser.value)
    if (!status) {
      dialog.error({
        title: '创建失败',
        content: data,
        positiveText: '确定',
      })
    } else {
      message.success('创建成功')
    }
  } else {
    let [status, data] = await updateFrpAdminToken(
      tokenId.value,
      tokenName.value,
      tokenUser.value,
      tokenEnable.value ? 1 : 0,
    )
    if (!status) {
      dialog.error({
        title: '修改失败',
        content: data,
        positiveText: '确定',
      })
    } else {
      message.success('修改成功')
    }
  }
  await loadDataList()
}
async function generateToken(row) {
  dialog.info({
    title: '是否生成Token',
    content: `是否重新生成Token: "${row.name}？"`,
    negativeText: '取消',
    positiveText: '确定',
    async onPositiveClick() {
      let [status, data] = await generateFrpAdminToken(row.id)
      if (!status) {
        dialog.error({
          title: '生成失败',
          content: data,
          negativeText: '确定',
        })
      }
      message.success('生成成功')
      await loadDataList()
    },
  })
}
async function changeEnable(row, value) {
  let [status, data] = await updateFrpAdminToken(row.id, row.name, row.user, value ? 1 : 0)
  if (!status) {
    dialog.error({
      title: '修改失败',
      content: data,
      positiveText: '确定'
    })
  } else {
    message.success('修改成功')
  }
  await loadDataList()
}
async function deleteToken(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除Token: "${row.name}？"`,
    negativeText: '取消',
    positiveText: '确定',
    async onPositiveClick() {
      let [status, data] = await deleteFrpAdminToken(row.id)
      if (!status) {
        dialog.error({
          title: '删除失败',
          content: data,
          negativeText: '确定',
        })
      }
      message.success('删除成功')
      await loadDataList()
    },
  })
}
async function loadFrpAdminTokenList() {
  let [status, data] = await getFrpAdminTokenList()
  if (!status) {
    return dialogError(data)
  }
  frpAdminToken.tokens = data
  return true
}
async function loadAdminUserList() {
  let [status, data] = await getAdminUserList()
  if (!status) {
    return dialogError(data)
  }
  adminUser.users = data
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
async function loadDataList() {
  loadingBar.start()
  if (!await loadUserInfo()) return
  if (!await loadAdminUserList()) return
  if (!await loadFrpAdminTokenList()) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>Frp Token管理（管理员）</h3>
    </template>
    <n-flex>
      <n-button @click="openCreateModal()" size="large" type="primary">创建Token</n-button>
      <n-data-table :data="frpAdminToken.tokens" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card style="max-width: 400px">
      <template #header>
        <h3>{{ createToken ? '创建' : '编辑' }}Token</h3>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createToken">
          <n-input :value="tokenId" disabled readonly></n-input>
        </n-form-item>
        <n-form-item label="名称">
          <n-input v-model:value="tokenName"></n-input>
        </n-form-item>
        <n-form-item label="用户">
          <n-select :options="userSelectOption" v-model:value="tokenUser"></n-select>
        </n-form-item>
        <n-form-item label="Token" v-if="!createToken">
          <n-input :value="tokenValue" readonly disabled></n-input>
        </n-form-item>
        <n-form-item label="启用">
          <n-switch v-model:value="tokenEnable"></n-switch>
        </n-form-item>
      </n-form>
      <template #footer>
        <n-flex justify="end">
          <n-button size="large" @click="showModal = false">取消</n-button>
          <n-button size="large" type="primary" @click="submitModal()">确定</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped></style>
