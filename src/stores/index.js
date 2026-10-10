import { useUserStore } from '@/stores/user.js'
import { useMenuCollapseStore } from '@/stores/menu-collapse.js'
import { useTokenStore } from '@/stores/token.js'
import {useAdminGroupStore} from '@/stores/admin/admin-group.js'
import {useAdminPermissionStore} from '@/stores/admin/admin-permission.js'
import {useAdminUserStore} from '@/stores/admin/admin-user.js'
import {useFrpClientStore} from '@/stores/frp/frp-client.js'
import {useFrpConfigStore} from '@/stores/frp/frp-config.js'
import {useFrpProxyStore} from '@/stores/frp/frp-proxy.js'
import {useFrpAdminClientStore} from '@/stores/frp-admin/frp-admin-client.js'
import {useFrpAdminProxyStore} from '@/stores/frp-admin/frp-admin-proxy.js'
import {useFrpAdminRuleStore} from '@/stores/frp-admin/frp-admin-rule.js'
import {useFrpRuleStore} from '@/stores/frp/frp-rule.js'
import {useFrpAdminTokenStore} from '@/stores/frp-admin/frp-admin-token.js'
import {useFrpTokenStore} from '@/stores/frp/frp-token.js'
import {useOpenVPNUserStore} from '@/stores/openvpn/openvpn-user.js'
import {useOpenVPNClientStore} from '@/stores/openvpn/openvpn-client.js'
import {useOpenVPNGroupStore} from '@/stores/openvpn/openvpn-group.js'
import {useAdminSettingsStore} from "@/stores/admin/admin-settings.js";
import {useConfigStore} from "@/stores/config.js";
import {useAdminVersionStore} from "@/stores/admin/admin-version.js";

export {
  useTokenStore,
  useUserStore,
  useMenuCollapseStore,
  useFrpTokenStore,
  useFrpAdminTokenStore,
  useFrpRuleStore,
  useFrpAdminRuleStore,
  useFrpAdminProxyStore,
  useFrpAdminClientStore,
  useFrpProxyStore,
  useFrpConfigStore,
  useFrpClientStore,
  useAdminUserStore,
  useAdminPermissionStore,
  useAdminGroupStore,
  useOpenVPNClientStore,
  useOpenVPNGroupStore,
  useOpenVPNUserStore,
  useAdminSettingsStore,
  useConfigStore,
  useAdminVersionStore
}