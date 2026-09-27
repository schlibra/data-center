import { useTokenStore } from '@/stores/index.js'
import { apiBack, apiError } from '@/utils/index.js'
import axios from 'axios'

const token = useTokenStore()

async function getOpenVPNGroupList() {
  try {
    return apiBack(await axios.get('/api/openvpn/group', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getOpenVPNGroupInfo(id) {
  try {
    return apiBack(await axios.get(`/api/openvpn/group/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function createOpenVPNGroup(name, value) {
  try {
    return apiBack(await axios.post('/api/openvpn/group', { name, value }, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function updateOpenVPNGroup(id, name, value) {
  try {
    return apiBack(await axios.put(`/api/openvpn/group/${id}`, { name, value }, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function deleteOpenVPNGroup(id) {
  try {
    return apiBack(await axios.delete(`/api/openvpn/group/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
export {
  getOpenVPNGroupInfo,
  getOpenVPNGroupList,
  createOpenVPNGroup,
  updateOpenVPNGroup,
  deleteOpenVPNGroup,
}