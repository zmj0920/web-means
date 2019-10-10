import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home'
import VuexStore from './components/VuexStore'
import Parent from './components/Parent'
Vue.use(Router)
const router = new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '*',   // 错误路由
      redirect: '/'   //重定向
    },
    {
      path: '/vuex',
      name: 'vuex',
      component: VuexStore
    },
    {
      path: '/parent',
      name: 'parent',
      component: Parent
    },
  ]
});
export default router
