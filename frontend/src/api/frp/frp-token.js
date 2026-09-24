import axios from 'axios'
import { useTokenStore } from '@/stores/token.js'
import router from '@/router/index.js'
import { apiBack } from '@/utils/api.js'

const token = useTokenStore()

async function getFrpTokenList() {
  try {
    const res = await axios.get("/api/frp/token", token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}

async function updateFrpToken(id, name, enable) {
  try {
    const res = await axios.put(`/api/frp/token/${id}`, {
      enable, name
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}

async function createFrpToken(name) {
  try {
    const res = await axios.post("/api/frp/token", {
      name
    }, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}

async function generateFrpToken(id) {
  try {
    const res = await axios.post(`/api/frp/token/${id}`, {}, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}

async function deleteFrpToken(id) {
  try {
    const res = await axios.delete(`/api/frp/token/${id}`, token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}

export {
  getFrpTokenList,
  updateFrpToken,
  createFrpToken,
  generateFrpToken,
  deleteFrpToken,
}