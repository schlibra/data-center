import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useFrpRuleStore = defineStore("frp-rule", {
  state() {
    const rules = ref([])
    const count = computed(() => rules.value.length)
    return {
      rules,
      count,
    }
  },
  persist: true
})