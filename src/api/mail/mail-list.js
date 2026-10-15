import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

async function listMailList(domain) {
    try {
        return apiBack(await axios.get(`/api/mail/list/${domain}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function getMailList(domain, id) {
    try {
        return apiBack(await axios.get(`/api/mail/list/${domain}/${id}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function createMailList(domain, email, password, username) {
    try {
        return apiBack(await axios.post(`/api/mail/list/${domain}`, {
            email, password, username
        }, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function updateMailList(domain, recordId, email, password, username) {
    try {
        return apiBack(await axios.put(`/api/mail/list/${domain}/${recordId}`, {
            email, password, username
        }, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function deleteMailList(domain, id) {
    try {
        return apiBack(await axios.delete(`/api/mail/list/${domain}/${id}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
export {
    listMailList,
    getMailList,
    createMailList,
    updateMailList,
    deleteMailList
}