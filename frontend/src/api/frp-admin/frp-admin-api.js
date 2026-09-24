import { useTokenStore } from '@/stores/token.js'
import axios from 'axios'
import { apiBack } from '@/utils/api.js'

const token = useTokenStore()

async function getFrpAdminClientList() {
  try {
    const res = await axios.get("/api/frp/admin/api/client", token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}

async function getFrpAdminProxyList() {
  try {
    const res = await axios.get('/api/frp/admin/api/proxy', token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
export {
  getFrpAdminClientList,
  getFrpAdminProxyList
}