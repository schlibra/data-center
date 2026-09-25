import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack } from '@/utils'

const token = useTokenStore()

async function getFrpAdminRuleList() {
  try {
    const res = await axios.get('/api/frp/admin/rule', token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function getFrpAdminRuleInfo(id) {
  try {
    const res = await axios.get(`/api/frp/admin/rule/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function createFrpAdminRule(min, max, tokenId, user) {
  try {
    const res = await axios.post('/api/frp/admin/rule', {
      min, max, token: tokenId, user
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function updateFrpAdminRule(id, min, max, tokenId, user) {
  try {
    const res = await axios.put(`/api/frp/admin/rule/${id}`, {
      min, max, token: tokenId, user
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function deleteFrpAdminRule(id) {
  try {
    const res = await axios.delete(`/api/frp/admin/rule/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
export {
  getFrpAdminRuleInfo,
  getFrpAdminRuleList,
  createFrpAdminRule,
  updateFrpAdminRule,
  deleteFrpAdminRule,
}
