<script setup>
import { useAdminUserStore } from '@/stores/admin/admin-user.js'
import { adminGetUserList } from '@/api/admin/admin-user.js'
import { NButton, NFlex, useDialog } from 'naive-ui'
import { h, onMounted } from 'vue'

const dialog = useDialog()
const adminUser = useAdminUserStore()
const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: '用户名',
    key: 'username',
  },
  {
    title: '昵称',
    key: 'nickname',
  },
  {
    title: '用户组',
    key: 'group_info.name',
  },
  {
    title: '操作',
    key: 'action',
    render() {
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
          '修改密码',
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

async function loadAdminUserList() {
  let [status, data] = await adminGetUserList()
  if (!status) {
    return dialog.error({
      title: '数据获取失败',
      content: data,
      positiveText: '确定',
    })
  }
  adminUser.users = data
}
onMounted(async () => {
  await loadAdminUserList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>管理员 用户管理</h3>
    </template>
    <n-flex>
      <n-button type="primary" size="large">创建用户</n-button>
      <n-data-table :data="adminUser.users" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
</template>

<style scoped></style>
