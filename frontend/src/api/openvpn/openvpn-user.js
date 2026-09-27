import { useTokenStore } from '@/stores/index.js'
import { apiBack, apiError } from '@/utils/index.js'
import axios from 'axios'

const token = useTokenStore()

async function getOpenVPNUserList() {
  try {
    return apiBack(await axios.get('/api/openvpn/user', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getOpenVPNUserInfo(id) {
  try {
    return apiBack(await axios.get(`/api/openvpn/user/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function createOpenVPNUser(username, password, enabled='yes', share=1, src_addr=[]) {
  try {
    return apiBack(await axios.post('/api/openvpn/user', {username, password, enabled, share, src_addr}, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function updateOpenVPNUser(id, username, password, enabled='yes', share=1, src_addr=[]) {
  try {
    return apiBack(await axios.put(`/api/openvpn/user/${id}`, {username, password, enabled, share, src_addr}, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function deleteOpenVPNUser(id) {
  try {
    return apiBack(await axios.delete(`/api/openvpn/user/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
export {
  getOpenVPNUserInfo,
  getOpenVPNUserList,
  createOpenVPNUser,
  updateOpenVPNUser,
  deleteOpenVPNUser,
}