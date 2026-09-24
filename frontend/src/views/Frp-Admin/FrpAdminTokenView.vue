<script setup>
import { NButton, NFlex, NSwitch, useDialog } from 'naive-ui'
import { getFrpAdminTokenList, getUserInfo } from '@/api'
import { useFrpAdminTokenStore, useUserStore } from '@/stores/'
import { h, onMounted } from 'vue'

const dialog = useDialog()
const frpAdminToken = useFrpAdminTokenStore()
const user = useUserStore()

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
        value: row.enable === 1
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
          },
          '编辑',
        ),
        h(
          NButton,
          {
            type: 'warning',
          },
          '生成Token',
        ),
        h(
          NButton,
          {
            type: 'error',
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
}

async function loadFrpAdminTokenList() {
  let [status, data] = await getFrpAdminTokenList()
  if (!status) {
    return dialogError(data)
  }
  frpAdminToken.tokens = data
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
  if (!await loadFrpAdminTokenList()) return
  await loadUserInfo()
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
      <n-button size="large" type="primary">创建Token</n-button>
      <n-data-table :data="frpAdminToken.tokens" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
</template>

<style scoped></style>
