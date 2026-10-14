import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const useMailMessageStore = defineStore('mail-message', {
    state() {
        const mailRecord = ref('')
        const message = ref([])
        const count = computed(() => message.value.length)
        return {
            mailRecord,
            message,
            count
        }
    },
    persist: true
})