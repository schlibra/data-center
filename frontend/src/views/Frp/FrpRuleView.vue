<script setup>
import { useFrpTokenStore } from '@/stores/frp/frp-token.js'
import { useFrpRuleStore } from '@/stores/frp/frp-rule.js'
import { getFrpTokenList } from '@/api/frp/frp-token.js'
import { createFrpRule, deleteFrpRule, getFrpRuleList, updateFrpRule } from '@/api/frp/frp-rule.js'
import { NButton, NFlex, useDialog, useMessage } from 'naive-ui'
import { computed, h, onMounted, ref } from 'vue'
import { getUserInfo } from '@/api/user.js'
import { useUserStore } from '@/stores/user.js'

const dialog = useDialog()
const message = useMessage()
const user = useUserStore()
const frpToken = useFrpTokenStore()
const frpRule = useFrpRuleStore()

const showModal = ref(false)
const createRule = ref(true)
const ruleId = ref('')
const ruleMin = ref('')
const ruleMax = ref('')
const ruleToken = ref('')
const options = computed(() =>
  frpToken.tokens.map((item) => {
    return {
      label: item.name,
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
    title: '最小端口',
    key: 'min',
  },
  {
    title: '最大端口',
    key: 'max',
  },
  {
    title: 'Token',
    key: 'token',
    render(row) {
      return h('span', {}, `${row.token_info.name} ( ${row.token} )`)
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

const dialogError = (content) => {
  dialog.error({
    title: '数据获取失败',
    content,
    positiveText: '确定',
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
async function loadFrpTokenList() {
  let [status, data] = await getFrpTokenList()
  if (!status) {
    return dialogError(data)
  }
  frpToken.tokens = data
  return true
}

async function loadFrpRuleList() {
  let [status, data] = await getFrpRuleList()
  if (!status) {
    return dialogError(data)
  }
  frpRule.rules = data
  console.log(frpRule.rules)
  return true
}
async function loadDataList() {
  if (!await loadUserInfo()) return
  if (!await loadFrpTokenList()) return
  await loadFrpRuleList()
}
async function deleteRule(row) {
  dialog.info({
    title: '是否删除',
    content: `是否删除端口规则？"`,
    negativeText: '取消',
    positiveText: '确定',
    async onPositiveClick() {
      let [status, data] = await deleteFrpRule(row.id)
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
function openCreateModal() {
  createRule.value = true
  ruleId.value = ''
  ruleMin.value = ''
  ruleMax.value = ''
  ruleToken.value = ''
  showModal.value = true
}
function openUpdateModal(row) {
  createRule.value = false
  ruleId.value = row.id
  ruleMin.value = row.min
  ruleMax.value = row.max
  ruleToken.value = row.token
  showModal.value = true
}
async function submitModal() {
  showModal.value = false
  if (createRule.value) {
    let [status, data] = await createFrpRule(ruleMin.value, ruleMax.value, ruleToken.value)
    if (!status) {
      return dialog.error({
        title: '创建失败',
        content: data,
        positiveText: '确定',
      })
    }
    message.success('创建成功')
  } else {
    let [status, data] = await updateFrpRule(
      ruleId.value,
      ruleMin.value,
      ruleMax.value,
      ruleToken.value,
    )
    if (!status) {
      return dialog.error({
        title: '修改失败',
        content: data,
        positiveText: '确定',
      })
    }
    message.success('修改成功')
  }
  await loadDataList()
}
onMounted(async () => {
  await loadDataList()
})
</script>

<template>
  <n-card>
    <template #header>
      <h3>Frp 端口规则管理</h3>
    </template>
    <n-flex>
      <n-button type="primary" size="large" @click="openCreateModal()">创建端口规则</n-button>
      <n-data-table :data="frpRule.rules" :columns="columns"></n-data-table>
    </n-flex>
  </n-card>
  <n-modal v-model:show="showModal">
    <n-card style="max-width: 400px">
      <template #header>
        <h3>{{ createRule ? '创建' : '编辑' }}端口规则</h3>
      </template>
      <n-form>
        <n-form-item label="ID" v-if="!createRule">
          <n-input :value="ruleId" disabled readonly></n-input>
        </n-form-item>
        <n-form-item label="最小端口">
          <n-input-number v-model:value="ruleMin"></n-input-number>
        </n-form-item>
        <n-form-item label="最大端口">
          <n-input-number v-model:value="ruleMax"></n-input-number>
        </n-form-item>
        <n-form-item label="Token">
          <n-select :options="options" v-model:value="ruleToken"></n-select>
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
