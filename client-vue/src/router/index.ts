import { createRouter, createWebHistory } from 'vue-router';
import * as layouts from '@/layouts';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      meta: {
        layout: 'Landing'
      },
      component: () => import('@/views/HomeView.vue')
    },
    {
      path: '/sign-up',
      name: 'sign-up',
      meta: {
        layout: 'Landing'
      },
      component: () => import('@/views/SignUpView.vue')
    },
    {
      path: '/sign-in',
      name: 'sign-in',
      meta: {
        layout: 'Landing'
      },
      component: () => import('@/views/SignInView.vue')
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      meta: {
        layout: 'Landing'
      },
      component: () => import('@/views/Dashboard.vue')
    }
  ]
})

export default router
