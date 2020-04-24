import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)
// const _import = require('./_import_' + process.env.NODE_ENV);
import lw from '../components/lw'
import getGinData from '../components/getGinData'
import login from '../components/login'
import home from '../components/home'
import about from '../components/about'
import sidebar from "../components/Sidebar/sidebar"

// 配置路由
const routes=[
	{path:"/login", component: login},
	{path:"/show", component: getGinData},
	{path:"/lw/:id", component: lw},
	{
		path:"/home", 
		component: home,
		// 子路由
		children: [
			{
				path: "phone",
				component: getGinData
			},
			{
				path: "tablet",
				component: getGinData
			},
			{
				path: "computer",
				component: getGinData
			}
		]
	},
	{path:"/about", component: about},
	{path:"/Sidebar/sidebar", component: sidebar},
]

// 实例化路由
const router = new VueRouter({
	routes // 简写相当 routes: routes
})

// console.log(routes);
// 导出路由模块
export default router