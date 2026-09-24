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

export { getAdminPermissionList }