import {useTokenStore} from "@/stores/index.js";
import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const useAutoLoginStore = defineStore('autologin', {
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