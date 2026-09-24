import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack } from '@/utils'

const token = useTokenStore()

async function getFrpClientList() {
  try {
    const res = await axios.get('/api/frp/api/client', token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}

async function getFrpProxyList() {
  try {
    const res = await axios.get('/api/frp/api/proxy', token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
export { getFrpProxyList, getFrpClientList }