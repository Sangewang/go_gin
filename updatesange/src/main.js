import Vue from 'vue'
import App from './App.vue'
import router from './router/index.js'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import store from './store';
import './ui.js' 
import './login.js' 
import VCharts from 'v-charts'
import Pagination from './components/Pagination.vue'

Vue.config.productionTip = false
Vue.use(ElementUI);
Vue.use(VCharts)

// 全局组件挂载
Vue.component('Pagination', Pagination)

new Vue({
  render: h => h(App),
  router, //挂在route.js 实例化的路由
  store
}).$mount('#app')