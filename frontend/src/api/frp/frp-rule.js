import { useTokenStore } from '@/stores/token.js'
import axios from 'axios'
import { apiBack } from '@/utils/api.js'

const token = useTokenStore()

async function getFrpRuleList() {
  try {
    const res = await axios.get('/api/frp/rule', token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function createFrpRule(min, max, tokenId) {
  try {
    const res = await axios.post('/api/frp/rule', {
      min, max, token: tokenId
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function updateFrpRule(id, min, max, tokenId) {
  try {
    const res = await axios.put(`/api/frp/rule/${id}`, {
      min,
      max,
      token: tokenId
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function deleteFrpRule(id) {
  try {
    const res = await axios.delete(`/api/frp/rule/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
export {
  getFrpRuleList,
  updateFrpRule,
  createFrpRule,
  deleteFrpRule
}