import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const useMailDomainStore = defineStore('mail-domain', {
    state() {
        const domain = ref([])
        const count = computed(() => domain.value.length)
        return {
            domain,
            count
        }
    },
    persist: true
})