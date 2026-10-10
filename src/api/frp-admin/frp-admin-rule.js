import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getFrpAdminRuleList() {
  try {
    return apiBack(await axios.get('/api/frp/admin/rule', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getFrpAdminRuleInfo(id) {
  try {
    return apiBack(await axios.get(`/api/frp/admin/rule/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function createFrpAdminRule(min, max, tokenId, user) {
  try {
    return apiBack(await axios.post('/api/frp/admin/rule', {
      min, max, token: tokenId, user
    }, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function updateFrpAdminRule(id, min, max, tokenId, user) {
  try {
    return apiBack(await axios.put(`/api/frp/admin/rule/${id}`, {
      min, max, token: tokenId, user
    }, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function deleteFrpAdminRule(id) {
  try {
    return apiBack(await axios.delete(`/api/frp/admin/rule/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
export {
  getFrpAdminRuleInfo,
  getFrpAdminRuleList,
  createFrpAdminRule,
  updateFrpAdminRule,
  deleteFrpAdminRule,
}
