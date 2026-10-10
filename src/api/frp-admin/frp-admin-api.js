import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getFrpAdminClientList() {
  try {
    return apiBack(await axios.get('/api/frp/admin/api/client', token.config))
  } catch (e) {
    return apiError(e)
  }
}

async function getFrpAdminProxyList() {
  try {
    return apiBack(await axios.get('/api/frp/admin/api/proxy', token.config))
  } catch (e) {
    return apiError(e)
  }
}
export { getFrpAdminClientList, getFrpAdminProxyList }