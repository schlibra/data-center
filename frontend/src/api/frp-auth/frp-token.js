import axios from 'axios'
import { useTokenStore } from '@/stores/token.js'

const token = useTokenStore()

async function getFrpTokenList() {
  try {
    const res = await axios.get("/api/frp/token", {
      headers: {
        Authorization: `Bearer ${token.token}`
      }
    })
    if (res.data.code === 200) {
      return [true, res.data.data]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}

async function updateFrpToken(id, name, enable) {
  try {
    const res = await axios.put(`/api/frp/token/${id}`, {
      enable: enable ? 1 : 0,
      name
    }, {
      headers: {
        Authorization: `Bearer ${token.token}`
      }
    })
    if (res.data.code === 200) {
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}

async function createFrpToken(name) {
  try {
    const res = await axios.post("/api/frp/token", {
      name
    }, {
      headers: {
        Authorization: `Bearer ${token.token}`
      }
    })
    if (res.data.code === 200) {
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}

async function generateFrpToken(id) {
  try {
    const res = await axios.post(`/api/frp/token/${id}`, {}, {
      headers: {
        Authorization: `Bearer ${token.token}`
      }
    })
    if (res.data.code === 200) {
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}

async function deleteFrpToken(id) {
  try {
    const res = await axios.delete(`/api/frp/token/${id}`, {
      headers: {
        Authorization: `Bearer ${token.token}`
      }
    })
    if (res.data.code === 200) {
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
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