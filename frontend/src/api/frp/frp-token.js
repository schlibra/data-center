import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getFrpTokenList() {
  try {
    return apiBack(await axios.get('/api/frp/token', token.config))
  } catch (e) {
    return apiError(e)
  }
}

async function updateFrpToken(id, name, enable) {
  try {
    return apiBack(await axios.put(
      `/api/frp/token/${id}`,
      {
        enable,
        name,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}

async function createFrpToken(name) {
  try {
    return apiBack(await axios.post(
      '/api/frp/token',
      {
        name,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}

async function generateFrpToken(id) {
  try {
    return apiBack(await axios.post(`/api/frp/token/${id}`, {}, token.config))
  } catch (e) {
    return apiError(e)
  }
}

async function deleteFrpToken(id) {
  try {
    return apiBack(await axios.delete(`/api/frp/token/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}

export { getFrpTokenList, updateFrpToken, createFrpToken, generateFrpToken, deleteFrpToken }