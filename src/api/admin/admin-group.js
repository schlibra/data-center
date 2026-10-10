import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function getAdminGroupList() {
  try {
    return apiBack(await axios.get("/api/admin/group", token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getAdminGroupInfo(id) {
  try {
    return apiBack(await axios.get(`/api/admin/group/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function createAdminGroup(name, admin, permission) {
  try {
    return apiBack(await axios.post("/api/admin/group", {
      name, admin, permission
    }, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function updateAdminGroup(id, name, admin, permission) {
  try {
    return apiBack(await axios.put(`/api/admin/group/${id}`, {
      name, admin, permission
    }, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function deleteAdminGroup(id) {
  try {
    return apiBack(await axios.delete(`/api/admin/group/${id}`, token.config))
  } catch (e) {
    return apiError(e)
  }
}
export {
  getAdminGroupList,
  getAdminGroupInfo,
  updateAdminGroup,
  createAdminGroup,
  deleteAdminGroup,
}