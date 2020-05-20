import Vue from 'vue'
import App from './App.vue'
import router from './router/index.js'
import store from './store';
import './ui.js' 

Vue.config.productionTip = false


new Vue({
  render: h => h(App),
  router, //挂在route.js 实例化的路由
  store
}).$mount('#app')