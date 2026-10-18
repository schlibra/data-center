import {
  getAdminGroupList,
  getAdminGroupInfo,
  updateAdminGroup,
  createAdminGroup,
  deleteAdminGroup,
} from './admin/admin-group.js'
import {
  getAdminPermissionList,
  getAdminPermissionInfo,
  updateAdminPermission,
  createAdminPermission,
  deleteAdminPermission,
} from './admin/admin-permission.js'
import {
  getAdminUserList,
  getAdminUserInfo,
  createAdminUser,
  deleteAdminUser,
  passwordAdminUser,
  updateAdminUser,
} from './admin/admin-user.js'
import {
  getFrpProxyList,
  getFrpClientList,
  getFrpInfo
} from './frp/frp-api.js'
import {
  getFrpRuleList,
  getFrpRuleInfo,
  updateFrpRule,
  createFrpRule,
  deleteFrpRule
} from './frp/frp-rule.js'
import {
  getFrpTokenList,
  updateFrpToken,
  createFrpToken,
  generateFrpToken,
  deleteFrpToken,
} from './frp/frp-token.js'
import {
  getFrpAdminClientList,
  getFrpAdminProxyList,
  clearOfflineProxyFrpAdmin,
} from './frp-admin/frp-admin-api.js'
import {
  getFrpAdminRuleInfo,
  getFrpAdminRuleList,
  createFrpAdminRule,
  updateFrpAdminRule,
  deleteFrpAdminRule,
} from './frp-admin/frp-admin-rule.js'
import {
  getFrpAdminTokenInfo,
  getFrpAdminTokenList,
  createFrpAdminToken,
  updateFrpAdminToken,
  deleteFrpAdminToken,
  generateFrpAdminToken,
} from './frp-admin/frp-admin-token.js'
import {
  loginKey,
  loginUser,
  registerKey,
  registerUser,
  getUserInfo,
  updateUser,
  logoutUser,
  setUserPassword,
  generateUserApiKey,
  getUserApiKey,
} from './user.js'
import {
  getOpenVPNClientInfo,
  getOpenVPNClientList,
  kickOpenVPNClient,
} from './openvpn/openvpn-client.js'
import {
  getOpenVPNUserInfo,
  getOpenVPNUserList,
  createOpenVPNUser,
  updateOpenVPNUser,
  deleteOpenVPNUser,
} from './openvpn/openvpn-user.js'
import {
  getOpenVPNGroupInfo,
  getOpenVPNGroupList,
  createOpenVPNGroup,
  updateOpenVPNGroup,
  deleteOpenVPNGroup,
} from './openvpn/openvpn-group.js'
import {
  listAdminSettings,
  getAdminSettings,
  setAdminSettings,
  createAdminSettings,
  updateAdminSettings,
  deleteAdminSettings,
} from './admin/admin-settings.js'
import {
    getAdminVersion,
    upgradeAdminVersion
} from "@/api/admin/admin-version.js";
import {
  listMailDomain,
  getMailDomain,
  createMailDomain,
  updateMailDomain,
  deleteMailDomain
} from './mail/mail-domain.js'
import {
  listMailList,
  getMailList,
  createMailList,
  updateMailList,
  deleteMailList
} from './mail/mail-list.js'
import {
  listMailMessage,
  createMailMessage,
} from './mail/mail-message.js'
import {
  listPveNode
} from './pve/pve-node.js'
import {
  listPveQemu
} from './pve/pve-qemu.js'
import {
  getPveQemuStatus,
  rebootPveQemuStatus,
  resetPveQemuStatus,
  resumePveQemuStatus,
  shutdownPveQemuStatus,
  startPveQemuStatus,
  stopPveQemuStatus,
  suspendPveQemuStatus
} from './pve/pve-qemu-status.js'
import {
  listPveLxc
} from './pve/pve-lxc.js'
import {
  getPveLxcStatus,
  rebootPveLxcStatus,
  resumePveLxcStatus,
  suspendPveLxcStatus,
  stopPveLxcStatus,
  startPveLxcStatus,
  shutdownPveLxcStatus
} from './pve/pve-lxc-status.js'

export {
  getAdminGroupList,
  getAdminGroupInfo,
  updateAdminGroup,
  createAdminGroup,
  deleteAdminGroup,
  getAdminPermissionList,
  getAdminPermissionInfo,
  updateAdminPermission,
  createAdminPermission,
  deleteAdminPermission,
  getAdminUserList,
  getAdminUserInfo,
  createAdminUser,
  deleteAdminUser,
  passwordAdminUser,
  updateAdminUser,
  getFrpProxyList,
  getFrpClientList,
  getFrpInfo,
  getFrpRuleList,
  getFrpRuleInfo,
  updateFrpRule,
  createFrpRule,
  deleteFrpRule,
  getFrpTokenList,
  updateFrpToken,
  createFrpToken,
  generateFrpToken,
  deleteFrpToken,
  getFrpAdminClientList,
  getFrpAdminProxyList,
  getFrpAdminRuleInfo,
  getFrpAdminRuleList,
  createFrpAdminRule,
  updateFrpAdminRule,
  deleteFrpAdminRule,
  getFrpAdminTokenInfo,
  getFrpAdminTokenList,
  createFrpAdminToken,
  updateFrpAdminToken,
  deleteFrpAdminToken,
  generateFrpAdminToken,
  loginKey,
  loginUser,
  registerKey,
  registerUser,
  getUserInfo,
  updateUser,
  logoutUser,
  generateUserApiKey,
  getUserApiKey,
  setUserPassword,
  getOpenVPNClientInfo,
  getOpenVPNClientList,
  kickOpenVPNClient,
  getOpenVPNUserInfo,
  getOpenVPNUserList,
  createOpenVPNUser,
  updateOpenVPNUser,
  deleteOpenVPNUser,
  getOpenVPNGroupInfo,
  getOpenVPNGroupList,
  createOpenVPNGroup,
  updateOpenVPNGroup,
  deleteOpenVPNGroup,
  listAdminSettings,
  getAdminSettings,
  setAdminSettings,
  createAdminSettings,
  updateAdminSettings,
  deleteAdminSettings,
  getAdminVersion,
  upgradeAdminVersion,
  listMailDomain,
  getMailDomain,
  createMailDomain,
  updateMailDomain,
  deleteMailDomain,
  listMailList,
  getMailList,
  createMailList,
  updateMailList,
  deleteMailList,
  listMailMessage,
  createMailMessage,
  clearOfflineProxyFrpAdmin,
  getPveQemuStatus,
  rebootPveQemuStatus,
  resetPveQemuStatus,
  resumePveQemuStatus,
  shutdownPveQemuStatus,
  startPveQemuStatus,
  stopPveQemuStatus,
  suspendPveQemuStatus,
  listPveNode,
  listPveQemu,
  getPveLxcStatus,
  rebootPveLxcStatus,
  resumePveLxcStatus,
  suspendPveLxcStatus,
  stopPveLxcStatus,
  startPveLxcStatus,
  shutdownPveLxcStatus,
  listPveLxc
}