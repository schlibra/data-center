import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

async function listMailDomain() {
    try {
        return apiBack(await axios.get('/api/mail/domain', token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function getMailDomain(id) {
    try {
        return apiBack(await axios.get(`/api/mail/domain/${id}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function createMailDomain(domain, name, tableId, server, port, ssl) {
    try {
        return apiBack(await axios.post('/api/mail/domain', {
            domain, name, tableId, server, port, ssl
        }, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function updateMailDomain(id, domain, name, tableId, server, port, ssl) {
    try {
        return apiBack(await axios.put(`/api/mail/domain/${id}`, {
            domain, name, tableId, server, port, ssl
        }, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function deleteMailDomain(id) {
    try {
        return apiBack(await axios.delete(`/api/mail/domain/${id}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
export {
    listMailDomain,
    getMailDomain,
    createMailDomain,
    updateMailDomain,
    deleteMailDomain
}