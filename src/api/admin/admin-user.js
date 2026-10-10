import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getAdminUserList() {
  try {
    return apiBack(await axios.get('/api/admin/user', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getAdminUserInfo(id) {
  try {
    return apiBack(await axios.get(`/api/admin/user/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function createAdminUser(username, password, nickname, group, enable) {
  try {
    return apiBack(await axios.post(
      '/api/admin/user',
      {
        username,
        password,
        nickname,
        group,
        enable,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function updateAdminUser(id, nickname, group, enable) {
  try {
    return apiBack(await axios.put(
      `/api/admin/user/${id}`,
      {
        nickname,
        group,
        enable,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function passwordAdminUser(id, password) {
  try {
    return apiBack(await axios.patch(
      `/api/admin/user/${id}`,
      {
        password,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function deleteAdminUser(id) {
  try {
    return apiBack(await axios.delete(`/api/admin/user/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}

export {
  getAdminUserList,
  getAdminUserInfo,
  createAdminUser,
  deleteAdminUser,
  passwordAdminUser,
  updateAdminUser,
}