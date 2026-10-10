import {defineStore} from "pinia";
import {ref} from "vue";

export const useAdminVersionStore = defineStore('admin-version', {
    state() {
        const version = ref('')
        const commit = ref('')
        const buildTime = ref('')
        const latestVersion = ref('')
        const set = data => {
            version.value = data.version
            commit.value = data.commit
            buildTime.value = data.buildTime
            latestVersion.value = data.latestVersion
        }
        return {
            version,
            commit,
            buildTime,
            latestVersion,
            set
        }
    },
    persist: true
})