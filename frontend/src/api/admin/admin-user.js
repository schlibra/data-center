import { useTokenStore } from '@/stores/token.js'
import axios from 'axios'

const token = useTokenStore()

async function adminGetUserList() {
  try {
    const res = await axios.get("/api/admin/user", token.config)
    if (res.data.code === 200) {
      return [true, res.data.data]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}

export {
  adminGetUserList
}