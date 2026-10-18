import {useTokenStore} from "@/stores/index.js";
import {apiBack, apiError} from "@/utils/index.js";
import axios from "axios";

const token = useTokenStore()

async function listPveNode() {
    try {
        return apiBack(await axios.get('/api/pve/node', token.config))
    } catch (e) {
        return apiError(e)
    }
}
export {
    listPveNode
}