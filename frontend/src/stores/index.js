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
}