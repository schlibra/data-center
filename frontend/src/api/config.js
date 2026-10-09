import {apiBack, apiError} from "@/utils";
import axios from "axios";

export async function getConfig(key) {
    try {
        return apiBack(await axios.get(`/api/config/${key}`))
    } catch (e) {
        return apiError(e)
    }
}