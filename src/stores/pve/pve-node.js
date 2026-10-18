import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const usePveNodeStore = defineStore('pve-node', {
    state() {
        const nodes = ref([])
        const count = computed(() => nodes.value.length)
        return {
            nodes,
            count
        }
    },
    persist: true
})