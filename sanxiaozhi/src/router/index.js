import Vue from 'vue'
import Router from 'vue-router'
import home from '../components/home'
import login from '../view/login/login'
//import dashboard from '../components/dashboard'
import sysdemoperate from '../components/sysdemoperate'
import sysdemocreate from '../components/sysdemocreate'
import calendar from '../components/calendar'
import table from '../components/table'
import chart from '../components/chart'
import Full from '../components/containers/Full'
import map from '../components/map/map'
import introduction from '../components/map/introduction'
import upload from '../components/map/upload'
import picshow from '../components/map/picshow'
import result from '../components/map/result'
import bdmap from '../components/map/bdmap'
import gdmap from '../components/map/gdmap'
import monitor from '../components/monitor/localserver'
import picshowbak from '../components/map/picshowbak'

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
		{ path: "/dashboard", name: "dashboard", component: Full},
		{ path: "/login", name: "login", component: login },
		{ 
			path: '/',
			name: 'home',
			component: home,
			redirect: "/login", //重定向，第一次进入就进入login，不添加的话第一次进入右侧是空白
			children: [
				/*
				{
					path: '/dashboard',
					name: 'dashboard',
					component: dashboard
				},*/
				{
					path: '/monitor',
					name: 'monitor',
					component: monitor
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
		},
		{
			path: '/map',
			name: 'map',
			component: map,
			redirect: "/map/introduction", //重定向，第一次进入就进入login，不添加的话第一次进入右侧是空白
			children: [{
					path: '/map/introduction',
					name: 'introduction',
					component: introduction
				},
				{
					path: '/map/upload',
					name: 'upload',
					component: upload
				},
				{
					path: '/map/picshow',
					name: 'picshow',
					component: picshow
				},
				{
					path: '/map/result',
					name: 'result',
					component: result
				},
				{
					path: '/map/bdmap',
					name: 'bdmap',
					component: bdmap
				},
				{
					path: '/map/gdmap',
					name: 'gdmap',
					component: gdmap
				},
				{
					path: '/map/picshowbak',
					name: 'picshowbak',
					component: picshowbak
				},
			]
		}
	]
})
