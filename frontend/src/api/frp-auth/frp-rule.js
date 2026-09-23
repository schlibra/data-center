import { useTokenStore } from '@/stores/token.js'
import axios from 'axios'

const token = useTokenStore()

async function getFrpRuleList() {
  try {
    const res = await axios.get('/api/frp/rule', {
      headers: {
        Authorization: `Bearer ${token.token}`,
      },
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
async function createFrpRule(min, max, tokenId) {
  try {
    const res = await axios.post('/api/frp/rule', {
      min,
      max,
      token: tokenId
    }, {
      headers: {
        Authorization: `Bearer ${token.token}`,
      },
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
async function updateFrpRule(id, min, max, tokenId) {
  try {
    const res = await axios.put(`/api/frp/rule/${id}`, {
      min,
      max,
      token: tokenId
    }, {
      headers: {
        Authorization: `Bearer ${token.token}`,
      },
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
async function deleteFrpRule(id) {
  try {
    const res = await axios.delete(`/api/frp/rule/${id}`, {
      headers: {
        Authorization: `Bearer ${token.token}`,
      },
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
  getFrpRuleList,
  updateFrpRule,
  createFrpRule,
  deleteFrpRule
}