import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import * as ElementUI from 'element-ui';
import echarts from 'echarts'
Vue.config.productionTip = false

Vue.prototype.$echarts = echarts


Vue.use(ElementUI);
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
