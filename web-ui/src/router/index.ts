import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import EnterSecretPage from '../views/EnterSecretPage.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: EnterSecretPage
  },
  {
    path: '/sec/:id',
    name: "Secret",
    component: () => import(/* webpackChunkName: "show-secret-page" */ '../views/ShowSecretPage.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
