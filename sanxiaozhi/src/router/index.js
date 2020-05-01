import Vue from 'vue'
import Router from 'vue-router'
import home from '../components/home'
import login from '../components/login'
import page1 from '../components/page1'
import sysdemoperate from '../components/sysdemoperate'
import sysdemocreate from '../components/sysdemocreate'
import calendar from '../components/calendar'
import table from '../components/table'
import chart from '../components/chart'
import Full from '../components/containers/Full'

const routerPush = Router.prototype.push
Router.prototype.push = function push(location) {
	return routerPush.call(this, location).catch(error => error)
}

Vue.use(Router)

export default new Router({
	routes: [ 
		/*
		{ path: "/home", name: "home", component: home },
		{ path: "/login", name: "login", component: login },
		{ path: "/page1", name: "page1", component: page1 },
		{ path: "/page2", name: "page2", component: page2 },
		{ path: "/calendar", name: "calendar", component: calendar },
		{ path: "/table", name: "table", component: table },
		{ path: "/chart", name: "chart", component: chart },*/
		{ path: "/dashboard", name: "dashboard", component: Full,  redirect: "/page1"},
		{ path: "/login", name: "login", component: login },
		{ 
			path: '/',
			name: 'home',
			component: home,
			redirect: "/login", //重定向，第一次进入就进入login，不添加的话第一次进入右侧是空白
			children: [{
					path: '/page1',
					name: 'page1',
					component: page1
				},
				{
					path: '/sysdemocreate',
					name: 'sysdemocreate',
					component: sysdemocreate
				},
				{
					path: '/sysdemoperate',
					name: 'sysdemoperate',
					component: sysdemoperate
				},
				{
					path: '/calendar',
					name: 'calendar',
					component: calendar
				},
				{
					path: '/table',
					name: 'table',
					component: table
				},
				{
					path: '/chart',
					name: 'chart',
					component: chart
				}
			]
		}
	]
})
