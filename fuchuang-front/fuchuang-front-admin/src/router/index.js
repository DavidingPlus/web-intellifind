import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/api/login',
      name: '登录',
      component: () => import('../views/login.vue')
    },
  ]
})

export default router
