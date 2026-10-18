import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const usePveQemuStore = defineStore('pve-qemu', {
    state() {
        const node = ref("")
        const qemu = ref([])
        const count = computed(() => qemu.value.length)
        return {
            node,
            qemu,
            count
        }
    },
    persist: true
})