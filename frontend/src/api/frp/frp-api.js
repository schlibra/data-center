import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getFrpClientList() {
  try {
    return apiBack(await axios.get('/api/frp/api/client', token.config))
  } catch (e) {
    return apiError(e)
  }
}

async function getFrpProxyList() {
  try {
    return apiBack(await axios.get('/api/frp/api/proxy', token.config))
  } catch (e) {
    return apiError(e)
  }
}
export { getFrpProxyList, getFrpClientList }