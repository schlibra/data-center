import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import("@/views/HomeView.vue"),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import("@/views/User/LoginView.vue")
    },
    {
      path: '/register',
      name: 'register',
      component: () => import("@/views/User/RegisterView.vue")
    },
    {
      path: '/frp/token',
      name: 'frp token',
      component: () => import("@/views/Frp/FrpTokenView.vue")
    },
    {
      path: '/frp/rule',
      name: 'frp rule',
      component: () => import("@/views/Frp/FrpRuleView.vue")
    },
    {
      path: '/frp/client',
      name: 'frp client',
      component: () => import("@/views/Frp/FrpClientView.vue")
    },
    {
      path: '/frp/proxy',
      name: 'frp proxy',
      component: () => import("@/views/Frp/FrpProxyView.vue")
    },
    {
      path: '/frp/config',
      name: 'frp config',
      component: () => import("@/views/Frp/FrpConfigView.vue")
    },
    {
      path: '/admin/user',
      name: 'admin user',
      component: () => import("@/views/Admin/AdminUserView.vue")
    },
    {
      path: '/admin/group',
      name: 'admin group',
      component: () => import("@/views/Admin/AdminGroupView.vue")
    },
    {
      path: '/admin/permission',
      name: 'admin permission',
      component: () => import("@/views/Admin/AdminPermissionView.vue")
    },
    {
      path: '/admin/settings',
      name: 'admin settings',
      component: () => import("@/views/Admin/AdminSettingsView.vue")
    },
    {
      path: '/admin/system',
      name: 'admin system',
      component: () => import("@/views/Admin/AdminSystemView.vue")
    },
    {
      path: '/admin/version',
      name: 'admin version',
      component: () => import("@/views/Admin/AdminVersionView.vue")
    },
    {
      path: '/frp/admin/token',
      name: 'frp token admin',
      component: () => import("@/views/Frp-Admin/FrpAdminTokenView.vue")
    },
    {
      path: '/frp/admin/rule',
      name: 'frp rule admin',
      component: () => import("@/views/Frp-Admin/FrpAdminRuleView.vue")
    },
    {
      path: '/frp/admin/client',
      name: 'frp client admin',
      component: () => import('@/views/Frp-Admin/FrpAdminClientView.vue')
    },
    {
      path: '/frp/admin/proxy',
      name: 'frp proxy admin',
      component: () => import('@/views/Frp-Admin/FrpAdminProxyView.vue')
    },
    {
      path: '/openvpn/user',
      name: 'openvpn user',
      component: () => import('@/views/OpenVPN/OpenVPNUser.vue')
    },
    {
      path: '/openvpn/client',
      name: 'openvpn client',
      component: () => import('@/views/OpenVPN/OpenVPNClient.vue')
    },
    {
      path: '/openvpn/group',
      name: 'openvpn group',
      component: () => import('@/views/OpenVPN/OpenVPNGroup.vue')
    },
    {
      path: '/user',
      name: 'user',
      component: () => import("@/views/User/UserView.vue")
    },
    {
      path: '/developer',
      name: 'developer',
      component: () => import("@/views/DeveloperView.vue")
    }
  ],
})

export default router
