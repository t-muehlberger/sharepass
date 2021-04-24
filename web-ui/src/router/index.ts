import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import EnterSecretPage from '../views/EnterSecretPage.vue'
import ShowSecretPage from '../views/ShowSecretPage.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: EnterSecretPage
  },
  {
    path: '/sec/:id',
    name: "Secret",
    component: ShowSecretPage
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
