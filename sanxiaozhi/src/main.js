import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import router from './router/index.js'
import Pagination from './components/Pagination'

import { Scrollbar, Menu, Submenu, MenuItem, MenuItemGroup } from 'element-ui'
 
import { post, fetch, patch, put } from './api/http'

//定义全局变量
Vue.prototype.$post = post;
Vue.prototype.$fetch = fetch;
Vue.prototype.$patch = patch;
Vue.prototype.$put = put;
 
Vue.use(Scrollbar)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)
Vue.use(MenuItemGroup)

Vue.config.productionTip = false
Vue.use(ElementUI);

// 全局组件挂载
Vue.component('Pagination', Pagination)

new Vue({
  render: h => h(App),
  router, //挂在route.js 实例化的路由
}).$mount('#app')
