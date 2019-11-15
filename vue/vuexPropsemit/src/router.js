import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home'
import VuexStore from './components/VuexStore'
import Parent from './components/Parent'
import Input from './components/Input'
import Slot1 from '@/components/slot1/Slot1'
import Slot2 from '@/components/slot2/Slot2'
import Slot3 from '@/components/slot3/Slot3'
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
    {
      path: '/input',
      name: 'input',
      component: Input
    },
    {
      path: '/slot1',
      name: 'Slot1',
      component: Slot1
    },
    {
      path: '/slot2',
      name: 'Slot2',
      component: Slot2
    },
    {
      path: '/slot3',
      name: 'Slot3',
      component: Slot3
    },
  ]
});
export default router
