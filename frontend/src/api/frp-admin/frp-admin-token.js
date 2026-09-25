import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getFrpAdminTokenList() {
  try {
    return apiBack(await axios.get('/api/frp/admin/token', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getFrpAdminTokenInfo(id) {
  try {
    return apiBack(await axios.get(`/api/frp/admin/token/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function createFrpAdminToken(name, user) {
  try {
    return apiBack(await axios.post(
      '/api/frp/admin/token',
      {
        user,
        name,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function updateFrpAdminToken(id, name, user, enable) {
  try {
    return apiBack(await axios.put(
      `/api/frp/admin/token/${id}`,
      {
        name,
        user,
        enable,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function deleteFrpAdminToken(id) {
  try {
    return apiBack(await axios.delete(`/api/frp/admin/token/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function generateFrpAdminToken(id) {
  try {
    return apiBack(await axios.post(`/api/frp/admin/token/${id}`, {}, token.config))
  } catch (e) {
    return apiError(e)
  }
}

export {
  getFrpAdminTokenInfo,
  getFrpAdminTokenList,
  createFrpAdminToken,
  updateFrpAdminToken,
  deleteFrpAdminToken,
  generateFrpAdminToken,
}