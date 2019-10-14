import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Parent from './components/Parent.vue'
import Child from './components/Child.vue'


Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/child',
      name: 'child',
      component: Child
    },
    {
      path: '/parent',
      name: 'parent',
      component: Parent
    }
  ]
})
