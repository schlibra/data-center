<script setup>
import { useFrpConfigStore } from '@/stores'
import { ref } from 'vue'

const configRef = ref(null)

const copyConfig = () => {
  configRef.value.focus()
  document.execCommand("copy")
}

const frpConfig = useFrpConfigStore()
const selectOptions = [
  {
    label: "TCP",
    value: "tcp",
  },
  {
    label: "UDP",
    value: "udp",
  }
]
function onCreate() {
  return {
    name: "",
    type: "",
    localIP: "",
    localPort: "",
    remotePort: "",
  }
}
</script>

<template>
  <n-card>
    <template #header>
      <h3>配置生成</h3>
    </template>
    <n-form>
      <n-form-item label="前置配置">
        <n-input type="textarea" :rows="2" v-model:value="frpConfig.preConfig"></n-input>
      </n-form-item>
      <n-form-item label="映射配置">
        <n-dynamic-input v-model:value="frpConfig.mainConfig" @create="onCreate()">
          <template #create-button-default>
            <span>创建映射配置</span>
          </template>
          <template #default="{ value }">
            <n-input v-model:value="value.name" placeholder="名称" style="max-width: 150px"></n-input>
            <n-input v-model:value="value.localIP" placeholder="本地IP" style="max-width: 150px"></n-input>
            <n-select :options="selectOptions" v-model:value="value.type" placeholder="映射类型" style="max-width: 150px"></n-select>
            <n-input-number v-model:value="value.localPort" placeholder="本地端口" style="max-width: 150px"></n-input-number>
            <n-input-number v-model:value="value.remotePort" placeholder="远程端口" style="max-width: 150px"></n-input-number>
          </template>
        </n-dynamic-input>
      </n-form-item>
      <n-form-item label="完整配置">
        <n-input type="textarea" ref="configRef" :value="frpConfig.resultConfig" :rows="8"></n-input>
      </n-form-item>
    </n-form>
    <template #action>
      <n-flex justify="center">
        <n-button type="warning" size="large" @click="frpConfig.clear()">清空</n-button>
        <n-button type="primary" size="large" @click="copyConfig()">复制</n-button>
      </n-flex>
    </template>
  </n-card>
</template>

<style scoped></style>
