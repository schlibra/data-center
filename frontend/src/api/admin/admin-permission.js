import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack } from '@/utils'

const token = useTokenStore()

async function getAdminPermissionList() {
  try {
    const res = await axios.get('/api/admin/permission', token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function getAdminPermissionInfo(id) {
  try {
    const res = await axios.get(`/api/admin/permission/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function createAdminPermission(key, name, parent) {
  try {
    const res = await axios.post("/api/admin/permission", {
      key, name, parent
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function updateAdminPermission(id, key, name, parent) {
  try {
    const res = await axios.put(`/api/admin/permission/${id}`, {
      key, name, parent
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function deleteAdminPermission(id) {
  try {
    const res = await axios.delete(`/api/admin/permission/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}

export {
  getAdminPermissionList,
  getAdminPermissionInfo,
  updateAdminPermission,
  createAdminPermission,
  deleteAdminPermission,
}