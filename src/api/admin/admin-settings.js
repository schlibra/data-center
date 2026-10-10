import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";
import {useTokenStore} from "@/stores/index.js";

const token = useTokenStore()

async function listAdminSettings() {
    try {
        return apiBack(await axios.get('/api/admin/settings', token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function getAdminSettings(id) {
    try {
        return apiBack(await axios.get(`/api/admin/settings/${id}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function setAdminSettings(key, name, value) {
    try {
        return apiBack(await axios.put('/api/admin/settings', {key, name, value}, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function createAdminSettings(key, name, value) {
    try {
        return apiBack(await axios.post('/api/admin/settings', {key, name, value}, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function updateAdminSettings(id, name, value) {
    try {
        return apiBack(await axios.put(`/api/admin/settings/${id}`, {name, value}, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function deleteAdminSettings(id) {
    try {
        return apiBack(await axios.delete(`/api/admin/settings/${id}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
export {
    listAdminSettings,
    getAdminSettings,
    setAdminSettings,
    createAdminSettings,
    updateAdminSettings,
    deleteAdminSettings,
}