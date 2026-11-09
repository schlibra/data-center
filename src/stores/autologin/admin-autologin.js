import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const useAdminAutoLoginStore = defineStore('admin-autologin', {
    state() {
        const autoLogins = ref([])
        const count = computed(() => autoLogins.value.length)
        return {
            autoLogins,
            count
        }
    },
    persist: true
})