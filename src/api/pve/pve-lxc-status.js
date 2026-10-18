import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

async function getPveLxcStatus(node, vmId) {
    try {
        return apiBack(await axios.get(`/api/pve/lxc-status/${node}/${vmId}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function setPveLxcStatus(node, vmId, action) {
    try {
        return apiBack(await axios.post(`/api/pve/lxc-status/${node}/${vmId}`, {action: action}, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function rebootPveLxcStatus(node, vmId) {
    return await setPveLxcStatus(node, vmId, "reboot")
}
async function resumePveLxcStatus(node, vmId) {
    return await setPveLxcStatus(node, vmId, "resume")
}
async function suspendPveLxcStatus(node, vmId) {
    return await setPveLxcStatus(node, vmId, "suspend")
}
async function stopPveLxcStatus(node, vmId) {
    return await setPveLxcStatus(node, vmId, "stop")
}
async function startPveLxcStatus(node, vmId) {
    return await setPveLxcStatus(node, vmId, "start")
}
async function shutdownPveLxcStatus(node, vmId) {
    return await setPveLxcStatus(node, vmId, "shutdown")
}

export {
    getPveLxcStatus,
    rebootPveLxcStatus,
    resumePveLxcStatus,
    suspendPveLxcStatus,
    stopPveLxcStatus,
    startPveLxcStatus,
    shutdownPveLxcStatus
}