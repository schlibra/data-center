import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const usePveLxcStore = defineStore('pve-lxc', {
    state() {
        const node = ref('')
        const lxc = ref([])
        const count = computed(() => lxc.value.length)
        return {
            node,
            lxc,
            count
        }
    },
    persist: true
})