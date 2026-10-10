import {defineStore} from "pinia";
import {computed, ref} from "vue";

export const useAdminSettingsStore = defineStore('adminSettings', {
    state() {
        const settings = ref([])
        const count = computed(() => settings.value.length)
        return {
            settings,
            count
        }
    }
})