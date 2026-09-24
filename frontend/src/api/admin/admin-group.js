import { useTokenStore } from '@/stores/token.js'
import axios from 'axios'
import { apiBack } from '@/utils/api.js'

const token = useTokenStore()

async function getAdminGroupList() {
  try {
    const res = await axios.get("/api/admin/group", token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
export {
  getAdminGroupList
}