import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

async function listMailMessage(domain, record) {
    try {
        return apiBack(await axios.get(`/api/mail/domain/${domain}/${record}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function createMailMessage(domain, record, id) {
    try {
        return apiBack(await axios.get(`/api/mail/domain/${domain}/${record}/${id}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}

export {
    listMailMessage,
    createMailMessage,
}