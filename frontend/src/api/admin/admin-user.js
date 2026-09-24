import { useTokenStore } from '@/stores/token.js'
import axios from 'axios'
import { apiBack } from '@/utils/api.js'

const token = useTokenStore()

async function getAdminUserList() {
  try {
    const res = await axios.get("/api/admin/user", token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function getAdminUserInfo(id) {
  try {
    const res = await axios.get(`/api/admin/user/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function createAdminUser(username, password, nickname, group, enable) {
  try {
    const res = await axios.post("/api/admin/user", {
      username, password, nickname, group, enable
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function updateAdminUser(id, nickname, group, enable) {
  try {
    const res = await axios.put(`/api/admin/user/${id}`, {
      nickname,
      group,
      enable
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function passwordAdminUser(id, password) {
  try {
    const res = await axios.patch(`/api/admin/user/${id}`, {
      password
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function deleteAdminUser(id) {
  try {
    const res = await axios.delete(`/api/admin/user/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
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