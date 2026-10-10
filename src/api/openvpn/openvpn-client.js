import { useTokenStore } from '@/stores/index.js'
import { apiBack, apiError } from '@/utils/index.js'
import axios from 'axios'

const token = useTokenStore()

async function getOpenVPNClientList() {
  try {
    return apiBack(await axios.get('/api/openvpn/client', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getOpenVPNClientInfo(id) {
  try {
    return apiBack(await axios.get(`/api/openvpn/client/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function kickOpenVPNClient(id) {
  try {
    return apiBack(await axios.delete(`/api/openvpn/client/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
export {
  getOpenVPNClientInfo,
  getOpenVPNClientList,
  kickOpenVPNClient,
}