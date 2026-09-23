import { useTokenStore } from '@/stores/token.js'
import axios from 'axios'

const token = useTokenStore()

async function getFrpClientList() {
  try {
    const res = await axios.get("/api/frp/api/client", {
      headers: {
        Authorization: `Bearer ${token.token}`
      }
    })
    if (res.data.code === 200) {
      return [true, res.data.data]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}

async function getFrpProxyList() {
  try {
    const res = await axios.get('/api/frp/api/proxy', {
      headers: {
        Authorization: `Bearer ${token.token}`,
      },
    })
    if (res.data.code === 200) {
      return [true, res.data.data]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}
export {
  getFrpProxyList,
  getFrpClientList
}