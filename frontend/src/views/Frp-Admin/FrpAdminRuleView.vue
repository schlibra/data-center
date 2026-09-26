<script setup>
import {
  useAdminUserStore,
  useFrpAdminRuleStore,
  useFrpAdminTokenStore,
  useUserStore,
} from '@/stores'
import {
  createFrpAdminRule,
  deleteFrpAdminRule,
  getAdminUserList,
  getFrpAdminRuleList,
  getFrpAdminTokenList,
  getUserInfo,
  updateFrpAdminRule,
} from '@/api'
import { NButton, NFlex, useDialog, useLoadingBar, useMessage } from 'naive-ui'
import { computed, h, onMounted, ref } from 'vue'

const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()
const user = useUserStore()
const frpAdminRule = useFrpAdminRuleStore()
const frpAdminToken = useFrpAdminTokenStore()
const adminUser = useAdminUserStore()

const showModal = ref(false)
const createRule = ref(true)
const rulePortRange = ref(true)
const ruleId = ref('')
const ruleMin = ref('')
const ruleMax = ref('')
const ruleToken = ref('')
const ruleUser = ref('')

const tokenSelectOption = computed(() =>
  frpAdminToken.tokens.map((item) => {
    return {
      label: item.name,
      value: item.id,
    }
  }),
)
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
    width: 60,
  },
  {
    title: '最小端口',
    key: 'min',
    width: 100,
  },
  {
    title: '最大端口',
    key: 'max',
    width: 100,
  },
  {
    title: 'Token',
    key: 'token',
    width: 200,
    render(row) {
      return h('span', {}, `${row.token_info.name}(${row.token})`)
    },
  },
  {
    title: '用户',
    key: 'user',
    width: 150,
    render(row) {
      return h('span', {}, `${row.user_info.nickname}(${row.user_info.username})`)
    },
  },
  {
    title: '操作',
    key: 'action',
    width: 150,
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
            type: 'error',
            onClick() {
              deleteRule(row)
            },
          },
          '删除',
        ),
      ])
    },
  },
]
function openCreateModal() {
  ruleId.value = ''
  ruleMin.value = ''
  ruleMax.value = ''
  ruleToken.value = ''
  ruleUser.value = ''
  createRule.value = true
  showModal.value = true
}
function openUpdateModal(row) {
  ruleId.value = row.id
  ruleMin.value = row.min
  ruleMax.value = row.max
  rulePortRange.value = ruleMin.value !== ruleMax.value
  ruleToken.value = row.token
  ruleUser.value = row.user
  createRule.value = false
  showModal.value = true
}
async function submitModal() {
  showModal.value = false
  if (!rulePortRange.value) {
    ruleMax.value = ruleMin.value
  }
  if (createRule.value) {
    let [status, data] = await createFrpAdminRule(
      ruleMin.value,
      ruleMax.value,
      ruleToken.value,
      ruleUser.value,
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
    let [status, data] = await updateFrpAdminRule(
      ruleId.value,
      ruleMin.value,
      ruleMax.value,
      ruleToken.value,
      ruleUser.value,
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
async function deleteRule(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除端口规则？"`,
    negativeText: '取消',
    positiveText: '确定',
    async onPositiveClick() {
      let [status, data] = await deleteFrpAdminRule(row.id)
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
const dialogError = (content) => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
  })
  loadingBar.error()
}

async function loadUserInfo() {
  let [status, data] = await getUserInfo()
  if (!status) {
    return dialogError(data)
  }
  user.setUserInfo(data)
  return true
}
async function loadFrpAdminTokenList() {
  let [status, data] = await getFrpAdminTokenList()
  if (!status) {
    return dialogError(data)
  }
  frpAdminToken.tokens = data
  return true
}
async function loadFrpAdminRuleList() {
  let [status, data] = await getFrpAdminRuleList()
  if (!status) {
    return dialogError(data)
  }
  frpAdminRule.rules = data
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
async function loadDataList() {
  loadingBar.start()
  if (!(await loadUserInfo())) return
  if (!(await loadFrpAdminTokenList())) return
  if (!(await loadFrpAdminRuleList())) return
  if (!(await loadAdminUserList())) return
  loadingBar.finish()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>Frp 端口规则管理（管理员）</h3>
    </template>
    <n-flex>
      <n-button size="large" type="primary" @click="openCreateModal()">创建端口规则</n-button>
      <n-data-table :columns="columns" :data="frpAdminRule.rules"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card style="max-width: 400px">
      <template #header>
        <h3>{{ createRule ? '创建' : '编辑' }}端口规则</h3>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createRule">
          <n-input :value="ruleId" readonly disabled></n-input>
        </n-form-item>
        <n-form-item label="端口类型">
          <n-radio-group v-model:value="rulePortRange">
            <n-radio :value="false">单个端口</n-radio>
            <n-radio :value="true">端口范围</n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item :label="(rulePortRange ? '最小' : '') + '端口'">
          <n-input-number v-model:value="ruleMin"></n-input-number>
        </n-form-item>
        <n-form-item label="最大端口" v-if="rulePortRange">
          <n-input-number v-model:value="ruleMax"></n-input-number>
        </n-form-item>
        <n-form-item label="Token">
          <n-select v-model:value="ruleToken" :options="tokenSelectOption"></n-select>
        </n-form-item>
        <n-form-item label="用户">
          <n-select v-model:value="ruleUser" :options="userSelectOption"></n-select>
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
