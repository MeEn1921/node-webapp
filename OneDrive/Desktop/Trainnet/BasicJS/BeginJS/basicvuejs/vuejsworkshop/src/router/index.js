import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '',
    name: 'Toolbar',
    component: () => import(/* webpackChunkName: "about" */ '../views/Toolbar.vue'),
    children: [
      {
          path: '/ME',
          name: 'ME',
          component: () => import(/* webpackChunkName: "about" */ '../views/ME.vue')
      },
      {
        path: '/about',
        name: 'about',
        component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
