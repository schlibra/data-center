<script setup>
import { createOpenVPNGroup, deleteOpenVPNGroup, getOpenVPNGroupList, getUserInfo, updateOpenVPNGroup } from '@/api/index.js'
import { useOpenVPNGroupStore, useUserStore } from '@/stores/index.js'
import { NButton, NFlex, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import { h, onMounted, ref } from 'vue'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()

const user = useUserStore()
const openVPNGroup = useOpenVPNGroupStore()

const showModal = ref(false)
const createGroup = ref(true)
const groupId = ref('')
const groupName = ref('')
const groupRule = ref([])

const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: 'IP组名称',
    key: 'group_name',
  },
  {
    title: '引用次数',
    key: 'ref_count',
  },
  {
    title: 'IP规则',
    key: 'group_value',
    render(row) {
      return row.group_value
        .map((item) => {
          return item.ip
        })
        .join(', ')
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
            size: 'large',
            onClick() {
              openUpdateModal(row)
            },
          },
          '编辑',
        ),
        h(NButton, {
          type: 'error',
          size: 'large',
          onClick() {
            deleteGroup(row)
          }
        }, '删除'),
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
function openCreateModal() {
  createGroup.value = true
  groupId.value = ''
  groupName.value = ''
  groupRule.value = []
  showModal.value = true
}
function openUpdateModal(row) {
  createGroup.value = false
  groupId.value = row.id
  groupName.value = row.group_name
  groupRule.value = row.group_value.map((item) => {
    return item.ip
  })
  showModal.value = true
}
async function submitModal() {
  showModal.value = false
  if (createGroup.value) {
    let [status, data] = await createOpenVPNGroup(
      groupName.value,
      groupRule.value.map((item) => {
        return { ip: item }
      }),
    )
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
    let [status, data] = await updateOpenVPNGroup(groupId.value, groupName.value, groupRule.value.map((item) => {
      return { ip: item }
    }))
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
function deleteGroup(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除IP组"${row.group_name}"？`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await deleteOpenVPNGroup(row.id)
      if (!status) {
        dialog.error({
          title: '删除失败',
          content: data,
          positiveText: '确定',
        })
      } else {
        message.success('删除成功')
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
async function loadOpenVPNGroupList() {
  let [status, data] = await getOpenVPNGroupList()
  if (!status) {
    return dialogError(data)
  }
  openVPNGroup.groups = data.ip_data
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!(await loadOpenVPNGroupList())) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card size="large">
    <template #header>
      <n-h1 prefix="bar">OpenVPN IP组管理</n-h1>
    </template>
    <n-flex>
      <n-button size="large" type="primary" @click="openCreateModal()">创建IP组</n-button>
      <n-data-table :data="openVPNGroup.groups" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card prefix="bar" style="max-width: 400px">
      <template #header>
        <n-h1>{{ createGroup ? '创建' : '编辑' }}IP组</n-h1>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createGroup">
          <n-input :value="groupId" disabled readonly></n-input>
        </n-form-item>
        <n-form-item label="IP组名称">
          <n-input v-model:value="groupName"></n-input>
        </n-form-item>
        <n-form-item label="IP规则">
          <n-dynamic-input v-model:value="groupRule"></n-dynamic-input>
        </n-form-item>
      </n-form>
      <template #action>
        <n-flex justify="end">
          <n-button size="large" @click="showModal = false">取消</n-button>
          <n-button size="large" type="primary" @click="submitModal()">确定</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped></style>
