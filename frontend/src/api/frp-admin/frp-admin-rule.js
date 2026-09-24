import { useTokenStore } from '@/stores/token.js'
import axios from 'axios'
import { apiBack } from '@/utils/api.js'

const token = useTokenStore()

async function getFrpAdminRuleList() {
  try {
    const res = await axios.get("/api/frp/admin/rule", token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function getFrpAdminRuleInfo(id) {

}
async function createFrpAdminRule(min, max, token, user) {

}
async function updateFrpAdminRule(id, min, max, token, user) {

}
async function deleteFrpAdminRule(id) {

}
export {
  getFrpAdminRuleInfo,
  getFrpAdminRuleList,
  createFrpAdminRule,
  updateFrpAdminRule,
  deleteFrpAdminRule
}
