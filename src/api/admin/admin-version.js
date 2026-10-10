import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

async function getAdminVersion() {
    try {
        return apiBack(await axios.get('/api/admin/version', token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function upgradeAdminVersion() {
    try {
        return apiBack(await axios.post('/api/admin/version', {}, token.config))
    } catch (e) {
        return apiError(e)
    }
}
export {
    getAdminVersion,
    upgradeAdminVersion,
}