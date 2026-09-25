import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack } from '@/utils'

const token = useTokenStore()

async function getAdminGroupList() {
  try {
    const res = await axios.get("/api/admin/group", token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function getAdminGroupInfo(id) {
  try {
    const res = await axios.get(`/api/admin/group/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function createAdminGroup(name, admin, permission) {
  try {
    const res = await axios.post("/api/admin/group", {
      name, admin, permission
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function updateAdminGroup(id, name, admin, permission) {
  try {
    const res = await axios.put(`/api/admin/group/${id}`, {
      name, admin, permission
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function deleteAdminGroup(id) {
  try {
    const res = await axios.delete(`/api/admin/group/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
export {
  getAdminGroupList,
  getAdminGroupInfo,
  updateAdminGroup,
  createAdminGroup,
  deleteAdminGroup,
}