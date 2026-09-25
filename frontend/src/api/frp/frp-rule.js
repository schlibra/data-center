import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getFrpRuleList() {
  try {
    return apiBack(await axios.get('/api/frp/rule', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getFrpRuleInfo(id) {
  try {
    return apiBack(await axios.get(`/api/frp/rule/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function createFrpRule(min, max, tokenId) {
  try {
    return apiBack(await axios.post(
      '/api/frp/rule',
      {
        min,
        max,
        token: tokenId,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function updateFrpRule(id, min, max, tokenId) {
  try {
    return apiBack(await axios.put(
      `/api/frp/rule/${id}`,
      {
        min,
        max,
        token: tokenId,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function deleteFrpRule(id) {
  try {
    return apiBack(await axios.delete(`/api/frp/rule/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
export {
  getFrpRuleList,
  getFrpRuleInfo,
  updateFrpRule,
  createFrpRule,
  deleteFrpRule
}
