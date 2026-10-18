import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

async function getPveQemuStatus(node, vmId) {
    try {
        return apiBack(await axios.get(`/api/pve/qemu-status/${node}/${vmId}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function setPveQemuStatus(node, vmId, action) {
    try {
        return apiBack(await axios.post(`/api/pve/qemu-status/${node}/${vmId}`, {action}, token.config))
    } catch (e) {
        return apiError(e)
    }
}
async function rebootPveQemuStatus(node, vmId) {
    return await setPveQemuStatus(node, vmId, "reboot")
}
async function resetPveQemuStatus(node, vmId) {
    return await setPveQemuStatus(node, vmId, "reset")
}
async function resumePveQemuStatus(node, vmId) {
    return await setPveQemuStatus(node, vmId, "resume")
}
async function shutdownPveQemuStatus(node, vmId) {
    return await setPveQemuStatus(node, vmId, "shutdown")
}
async function startPveQemuStatus(node, vmId) {
    return await setPveQemuStatus(node, vmId, "start")
}
async function stopPveQemuStatus(node, vmId) {
    return await setPveQemuStatus(node, vmId, "stop")
}
async function suspendPveQemuStatus(node, vmId) {
    return await setPveQemuStatus(node, vmId, "suspend")
}
export {
    getPveQemuStatus,
    rebootPveQemuStatus,
    resetPveQemuStatus,
    resumePveQemuStatus,
    shutdownPveQemuStatus,
    startPveQemuStatus,
    stopPveQemuStatus,
    suspendPveQemuStatus
}
