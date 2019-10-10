import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Index from '@/components/Index'
import Detail from '@/components/Detail'
import Page from '@/components/Page'
Vue.use(Router)
export default new Router({
    routes: [{
            path: '*',
            component: Home,
        },
        {
            path: '/home',
            name: 'Home',
            component: Home,
            children: [{
                    name: '首页',
                    path: '/home/index',
                    component: Index,
                    meta: {
                        keepAlive: true,
                        isBack: false
                    }
                },
                {
                    name: '详情页',
                    hidden: true,
                    path: '/detail',
                    component: Detail
                },
                {
                    name: '展示页',
                    path: '/page',
                    component: Page
                }
            ]
        }
    ]
})