import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack } from '@/utils'

const token = useTokenStore()

async function getFrpAdminTokenList() {
  try {
    const res = await axios.get('/api/frp/admin/token', token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function getFrpAdminTokenInfo(id) {
  try {
    const res = await axios.get(`/api/frp/admin/token/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function createFrpAdminToken(name, user) {
  try {
    const res = await axios.post(
      '/api/frp/admin/token',
      {
        user,
        name,
      },
      token.config,
    )
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function updateFrpAdminToken(id, name, user, enable) {
  try {
    const res = await axios.put(
      `/api/frp/admin/token/${id}`,
      {
        name,
        user,
        enable,
      },
      token.config,
    )
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function deleteFrpAdminToken(id) {
  try {
    const res = await axios.delete(`/api/frp/admin/token/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function generateFrpAdminToken(id) {
  try {
    const res = await axios.post(`/api/frp/admin/token/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
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