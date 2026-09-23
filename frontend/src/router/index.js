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
      path: '/user',
      name: 'user',
      component: () => import("@/views/User/UserView.vue")
    }
  ],
})

export default router
