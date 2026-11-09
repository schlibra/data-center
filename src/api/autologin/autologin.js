import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

export async function getAutoLoginList() {
    try {
        return apiBack(await axios.get('/api/autologin', token.config))
    } catch (e) {
        return apiError(e)
    }
}
export async function getAutoLoginLink(id) {
    try {
        return apiBack(await axios.get(`/api/autologin/${id}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}