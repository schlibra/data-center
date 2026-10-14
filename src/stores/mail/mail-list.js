import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const useMailListStore = defineStore('mail-list', {
    state() {
        const domainId = ref('')
        const mail = ref([])
        const count = computed(() => mail.value.length)
        return {
            domainId,
            mail,
            count
        }
    },
    persist: true
})