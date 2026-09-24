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

export {
  getAdminUserList
}