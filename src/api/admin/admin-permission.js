import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getAdminPermissionList() {
  try {
    return apiBack(await axios.get('/api/admin/permission', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getAdminPermissionInfo(id) {
  try {
    return apiBack(await axios.get(`/api/admin/permission/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function createAdminPermission(key, name, parent) {
  try {
    return apiBack(await axios.post("/api/admin/permission", {
      key, name, parent
    }, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function updateAdminPermission(id, key, name, parent) {
  try {
    return apiBack(await axios.put(`/api/admin/permission/${id}`, {
      key, name, parent
    }, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function deleteAdminPermission(id) {
  try {
    return apiBack(await axios.delete(`/api/admin/permission/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}

export {
  getAdminPermissionList,
  getAdminPermissionInfo,
  updateAdminPermission,
  createAdminPermission,
  deleteAdminPermission,
}