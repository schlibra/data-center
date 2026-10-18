import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

async function listPveQemu(node) {
    try {
        return apiBack(await axios.get(`/api/pve/qemu/${node}`, token.config))
    } catch (e) {
        return apiError(e)
    }
}
export {
    listPveQemu
}