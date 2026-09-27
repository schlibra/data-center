<script setup>
import {
  createOpenVPNUser,
  deleteOpenVPNUser,
  getOpenVPNGroupList,
  getOpenVPNUserList,
  getUserInfo,
  updateOpenVPNUser,
} from '@/api/index.js'
import { useOpenVPNGroupStore, useOpenVPNUserStore, useUserStore } from '@/stores/index.js'
import { NButton, NFlex, NSwitch, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import { h, onMounted, ref } from 'vue'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()

const user = useUserStore()
const openVPNUser = useOpenVPNUserStore()
const openVPNGroup = useOpenVPNGroupStore()

const showModal = ref(false)
const createUser = ref(true)
const userId = ref('')
const username = ref('')
const userpass = ref('')
const userShare = ref(0)
const userEnable = ref(true)
const userGroup = ref([])

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
    title: '密码',
    key: 'passwd',
    render(row) {
      return h(
        'span',
        {
          class: 'hide-password',
        },
        row.passwd,
      )
    },
  },
  {
    title: '共享数量',
    key: 'share',
  },
  {
    title: '可用IP',
    key: 'ip',
    render(row) {
      const obj = row.src_addr.object
      if (obj) {
        return obj
          .map((user_group) => {
            return openVPNGroup.groups
              .find((group) => user_group.gid === `IPGP${group.id}`)
              .group_value.map((gv_item) => {
                return gv_item.ip
              })
              .join(', ')
          })
          .join(', ')
      } else {
        return '-'
      }
    },
  },
  {
    title: '连接时间',
    key: 'start_time',
    render(row) {
      if (row.start_time === 0) {
        return '-'
      }
      const date = new Date(row.start_time * 1000)
      return date.toLocaleString()
    },
  },
  {
    title: '断开时间',
    key: 'last_offtime',
    render(row) {
      if (row.last_offtime === 0) {
        return '-'
      }
      const date = new Date(row.last_offtime * 1000)
      return date.toLocaleString()
    },
  },
  {
    title: '启用',
    key: 'enabled',
    render(row) {
      return h(NSwitch, {
        value: row.enabled === 'yes',
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
            size: 'large',
            onClick() {
              openUpdateModal(row)
            },
          },
          '编辑',
        ),
        h(
          NButton,
          {
            type: 'error',
            size: 'large',
            onClick() {
              deleteUser(row)
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
    content: content,
    positiveText: '确定',
  })
  loadingBar.error()
}
function openCreateModal() {
  createUser.value = true
  userId.value = ''
  username.value = ''
  userpass.value = ''
  userShare.value = 1
  userEnable.value = true
  userGroup.value = []
  showModal.value = true
}
function openUpdateModal(row) {
  createUser.value = false
  userId.value = row.id
  username.value = row.username
  userpass.value = row.passwd
  userShare.value = row.share
  userEnable.value = row.enabled === 'yes'
  userGroup.value = row.src_addr.object.map((item) => {
    return parseInt(item.gid.replace('IPGP', ''))
  })
  console.log(userGroup.value)
  showModal.value = true
}
async function submitModal() {
  showModal.value = false
  const src_addr = {
    object: userGroup.value.map((group_id) => {
      return {
        type: 0,
        gp_name: openVPNGroup.groups.find((item) => item.id === group_id).group_name,
        gid: `IPGP${group_id}`,
      }
    }),
  }
  if (createUser.value) {
    let [status, data] = await createOpenVPNUser(
      username.value,
      userpass.value,
      userEnable.value ? 'yes' : 'no',
      userShare.value,
      src_addr,
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
    let [status, data] = await updateOpenVPNUser(
      userId.value,
      username.value,
      userpass.value,
      userEnable.value ? 'yes' : 'no',
      userShare.value,
      src_addr,
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
function deleteUser(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除用户 ${row.username} ?`,
    positiveText: '确定',
    negativeText: '取消',
    async onPositiveClick() {
      let [status, data] = await deleteOpenVPNUser(row.id)
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
    },
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
async function loadOpenVPNUserList() {
  let [status, data] = await getOpenVPNUserList()
  if (!status) {
    return dialogError(data)
  }
  openVPNUser.users = data.results.data
  return true
}
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!(await loadOpenVPNGroupList())) return
  if (!(await loadOpenVPNUserList())) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>OpenVPN 用户管理</h3>
    </template>
    <n-flex>
      <n-button size="large" type="primary" @click="openCreateModal()">创建用户</n-button>
      <n-data-table :data="openVPNUser.users" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card style="max-width: 400px">
      <template #header>
        <h3>{{ createUser ? '创建' : '编辑' }}用户</h3>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createUser">
          <n-input :value="userId" readonly disabled></n-input>
        </n-form-item>
        <n-form-item label="用户名">
          <n-input v-model:value="username"></n-input>
        </n-form-item>
        <n-form-item label="密码">
          <n-input v-model:value="userpass" type="password" show-password-on="click"></n-input>
        </n-form-item>
        <n-form-item label="共享数量">
          <n-input-number v-model:value="userShare"></n-input-number>
        </n-form-item>
        <n-form-item label="启用">
          <n-switch v-model:value="userEnable"></n-switch>
        </n-form-item>
        <n-form-item label="IP组">
          <n-scrollbar style="height: 100px">
            <n-checkbox-group v-model:value="userGroup">
              <n-flex>
                <n-checkbox v-for="item in openVPNGroup.groups" :value="item.id">
                  {{ item.group_value.map((item) => item.ip).join(', ') }} ({{ item.group_name }})
                </n-checkbox>
              </n-flex>
            </n-checkbox-group>
          </n-scrollbar>
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

<style>
.hide-password {
  filter: blur(5px) opacity(0.8);
  transition: all 5s;
}
.hide-password:hover {
  filter: blur(0) opacity(1);
  transition: all 0.3s;
}
</style>
