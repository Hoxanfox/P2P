import { createRouter, createWebHashHistory } from 'vue-router'

import Dashboard from '../components/Dashboard.vue'
import Users from '../components/Users.vue'
import Groups from '../components/Groups.vue'

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard,
  },
  {
    path: '/users',
    name: 'Users',
    component: Users,
  },
  {
    path: '/groups',
    name: 'Groups',
    component: Groups,
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes, 
})

export default router
