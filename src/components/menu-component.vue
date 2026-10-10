<script setup>
import {computed, h, onMounted, ref} from 'vue'
import {NIcon, useMessage} from 'naive-ui'
import router from '@/router'
import {
  HomeOutline as HomeIcon,
  PlanetOutline as FrpAuthIcon,
  KeyOutline as TokenIcon,
  GlobeOutline as PortIcon,
  DesktopOutline as ClientIcon,
  RepeatOutline as ProxyIcon,
  PersonOutline as UserIcon,
  SaveOutline as ConfigIcon,
  PlanetSharp as FrpAdminIcon,
  SettingsOutline as AdminIcon,
  PeopleOutline as GroupIcon,
  LockClosedOutline as PermissionIcon,
  MailOutline as MailIcon,
  ListOutline as ListIcon,
  AlbumsOutline as InfoIcon,
  MailOpenOutline as OutlookIcon,
  CodeSlashOutline as DeveloperIcon,
  RocketOutline as OpenVPNIcon,
  FolderOpenOutline as IPGroupIcon,
  AppsOutline as AppsIcon,
} from '@vicons/ionicons5'
import { useUserStore, useMenuCollapseStore } from '@/stores'
import {getConfig} from "@/api/config.js";
import {useConfigStore} from "@/stores/";

const config = useConfigStore()
const menuCollapse = useMenuCollapseStore()
const user = useUserStore()
const message = useMessage()


const renderIcon = (icon) => () => h(NIcon, null, { default: () => h(icon) })

const menuChange = (key) => {
  urlPath.value = key
  router.push(key)
}

const urlPath = ref(location.pathname)
const menuOptions = computed(() => [
  {
    label: '首页',
    key: '/',
    icon: renderIcon(HomeIcon),
  },
  {
    label: 'Frp',
    icon: renderIcon(FrpAuthIcon),
    key: '/frp',
    show: user.hasPermission('menu.frp') && !!config.config["frp.enable"],
    children: [
      {
        label: 'Token管理',
        icon: renderIcon(TokenIcon),
        show: user.hasPermission('frp.token.get'),
        key: '/frp/token',
      },
      {
        label: '端口规则管理',
        icon: renderIcon(PortIcon),
        show: user.hasPermission('frp.rule.get'),
        key: '/frp/rule',
      },
      {
        label: '客户端管理',
        icon: renderIcon(ClientIcon),
        show: user.hasPermission('frp.api.client'),
        key: '/frp/client',
      },
      {
        label: '映射管理',
        icon: renderIcon(ProxyIcon),
        show: user.hasPermission('frp.api.proxy'),
        key: '/frp/proxy',
      },
      {
        label: '配置生成',
        icon: renderIcon(ConfigIcon),
        key: '/frp/config',
      },
    ],
  },
  {
    label: 'Frp（管理员）',
    key: '/frp/admin',
    show: user.isAdmin && !!config.config["frp.enable"],
    icon: renderIcon(FrpAdminIcon),
    children: [
      {
        label: 'Token管理',
        icon: renderIcon(TokenIcon),
        key: '/frp/admin/token',
      },
      {
        label: '端口规则管理',
        icon: renderIcon(PortIcon),
        key: '/frp/admin/rule',
      },
      {
        label: '客户端管理',
        icon: renderIcon(ClientIcon),
        key: '/frp/admin/client',
      },
      {
        label: '映射管理',
        icon: renderIcon(ProxyIcon),
        key: '/frp/admin/proxy',
      },
    ],
  },
  {
    label: 'OpenVPN管理',
    key: '/openvpn',
    show: user.hasPermission('menu.openvpn') && !!config.config["openvpn.enable"],
    icon: renderIcon(OpenVPNIcon),
    children: [
      {
        label: '用户管理',
        key: '/openvpn/user',
        icon: renderIcon(UserIcon),
      },
      {
        label: 'IP组管理',
        key: '/openvpn/group',
        icon: renderIcon(IPGroupIcon)
      },
      {
        label: '客户端管理',
        key: '/openvpn/client',
        icon: renderIcon(ClientIcon)
      }
    ]
  },
  {
    label: '邮箱管理',
    key: '/mail',
    show: user.hasPermission('menu.mail') && !!config.config["mail.enable"],
    icon: renderIcon(MailIcon),
    children: [
      {
        label: '邮箱列表',
        key: '/mail/list',
        icon: renderIcon(ListIcon),
      },
      {
        label: '邮件管理',
        key: '/mail/info',
        icon: renderIcon(InfoIcon),
      },
    ],
  },
  {
    label: 'Outlook管理',
    key: '/outlook',
    show: user.hasPermission('menu.outlook') && !!config.config["outlook.enable"],
    icon: renderIcon(OutlookIcon),
    children: [
      {
        label: '邮箱列表',
        key: '/outlook/list',
        icon: renderIcon(ListIcon),
      },
      {
        label: '邮件管理',
        key: '/outlook/info',
        icon: renderIcon(InfoIcon),
      },
    ],
  },
  {
    label: '管理员设置',
    key: '/admin',
    show: user.isAdmin,
    icon: renderIcon(AdminIcon),
    children: [
      {
        label: '用户管理',
        key: '/admin/user',
        icon: renderIcon(UserIcon),
      },
      {
        label: '用户组管理',
        key: '/admin/group',
        icon: renderIcon(GroupIcon),
      },
      {
        label: '权限管理',
        key: '/admin/permission',
        icon: renderIcon(PermissionIcon),
      },
      {
        label: '系统设置',
        key: '/admin/system',
        icon: renderIcon(AppsIcon)
      },
      {
        label: '所有设置',
        key: '/admin/settings',
        icon: renderIcon(ListIcon)
      }
    ],
  },
  {
    label: '用户中心',
    icon: renderIcon(UserIcon),
    key: '/user',
  },
  {
    label: '开发者功能',
    icon: renderIcon(DeveloperIcon),
    key: '/developer',
    show: user.hasPermission('menu.developer'),
  },
])

setInterval(() => {
  urlPath.value = location.pathname
}, 100)
async function loadConfig(key) {
  let [status, data] = await getConfig(key)
  if (!status) {
    message.error(`加载配置失败：${data}`)
  } else {
    try {
      config.config[key] = parseInt(data)
    } catch (e) {
      config.config[key] = data
    }
  }
}
async function loadDataList() {
  await loadConfig("frp.enable")
  await loadConfig("openvpn.enable")
  await loadConfig("mail.enable")
  await loadConfig("outlook.enable")
}

onMounted(() => {
  loadDataList()
})
</script>

<template>
  <n-layout-sider
    bordered
    :width="240"
    collapse-mode="width"
    :collapsed-width="48"
    :collapsed="menuCollapse.collapse"
    @collapse="menuCollapse.collapse = true"
    @expand="menuCollapse.collapse = false"
    show-trigger
  >
    <n-menu
      :options="menuOptions"
      class="menu"
      :value="urlPath"
      :collapsed="menuCollapse.collapse"
      :collapsed-width="48"
      :collapsed-icon-size="20"
      @update-value="menuChange"
      accordion
    ></n-menu>
  </n-layout-sider>
</template>

<style scoped>
.menu {
  height: calc(100vh - 75px);
}
</style>
