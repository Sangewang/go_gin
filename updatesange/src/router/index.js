import Vue from 'vue';
import Router from 'vue-router';

import login from '../views/login/index.vue'
import Full from '../containers/Full.vue'
import Dashboard from '../views/Dashboard.vue'
import Introduction from '../views/Introduction.vue'
import Local from '../views/monitor/localserver.vue'
import Remote from '../views/monitor/remoteserver.vue'
import Map from '../containers/Map.vue'
import MapIntroduction from '../views/map/introduction.vue'
import DataUpload from '../views/map/dataupload.vue'
import PicUpload from '../views/map/picupload.vue'
import MapCreate from '../views/map/create.vue'
import MapSearch from '../views/map/search.vue'
import MapDevice from '../views/map/device.vue'
import MapResult from '../views/map/result.vue'
import BdMap from '../views/map/bdmap.vue'
import GdMap from '../views/map/gdmap.vue'
import LogSearch from '../views/system/log.vue'
import LogDownload from '../views/system/log.vue'
import Create from '../views/system/create.vue'
import Table from '../views/system/table.vue'
import Company from '../views/system/company.vue'

Vue.use(Router)

export const constantRouterMap = [{
	path: '/login',
	component: login,
	hidden: true
}]


export default new Router({
	mode: 'hash',
	linkActiveClass: 'open active',
	scollBehavior: () => ({
		y: 0
	}),
	routes: constantRouterMap
})

export const asyncRouterMap = [
	{
		path: '/',
		redirect: '/dashboard',
		name: '首页',
		component: Full,
		hidden: false,
		children: [
			{ path: '/dashboard', name: 'Dashboard', icon:'md-home', component: Dashboard },
			{ path: '/introduction', name: '介绍', icon:'md-information-circle', component: Introduction },
			{ path: '/monitor', name: '系统监控', icon:'md-eye', redirect: '/monitor/local',
				component: { render (c) {return c('router-view') }},
				children: [ { path: 'local', name: '本机状态', icon:'md-bonfire', component: Local, } ,
							{ path: 'remote', name: '远程状态', icon:'md-flame', component: Remote, } ,
				]
			},
			{ path: '/system', name: '系统工具', icon:'md-construct', redirect: '/system/task',
				component: { render (c) {return c('router-view') }},
				children: [ { path: 'task', name: '系统配置', icon: 'md-construct', redirect: '/system/task/create', 
								component: { render (c) {return c('router-view') }},
								children: [
									{ path: 'create', name: '任务创建', icon:'md-color-filter', component: Create, } ,
									{ path: 'table', name: '图表操作', icon:'md-apps', component: Table, } ,
									{ path: 'company', name: '我爱我司', icon:'md-compass', component: Company, } ,
								]
				},
						{ path: 'log', name: '日志管理', icon: 'md-recording', redirect: '/system/log/search' ,
								component: { render (c) {return c('router-view') }},
								children: [
									{ path: 'search', name: '日志查询', icon:'md-search', component: LogSearch, } ,
									{ path: 'download', name: '日志下载', icon:'md-download', component: LogDownload, } ,
								]
				},
				]
			},
		]
	},
	
	{
		path: '/map',
		redirect: '/map/introduction',
		name: '地图',
		component: Map,
		hidden: false,
		children: [
			{ path: '/map/introduction', name: '系统简介', icon: 'md-locate', component: MapIntroduction },
			{ path: '/upload', name: '数据上传', icon: 'md-cloud-upload', redirect: '/map/upload', 
				component: { render (c) {return c('router-view') }},
				children: [ { path: 'dataupload', name: '文件上传', icon: 'md-repeat', component: DataUpload, }, 
							{ path: 'picupload', name: '图片上传', icon: 'md-image', component: PicUpload, }, 
				]
			},
			//{ path: '/map/upload', name: '数据上传', icon: 'md-cloud-upload', component: MapUpload },
			{ path: '/map/create', name: '任务创建', icon: 'md-send', component: MapCreate },
			{ path: '/map/search', name: '任务查询', icon: 'md-search', component: MapSearch },
			{ path: '/map/device', name: '设备状态', icon: 'logo-android', component: MapDevice },
			{ path: '/map/result', name: '评测结果', icon: 'md-podium', component: MapResult},
			{ path: '/bdgd', name: '地图系统', icon:'md-locate', redirect: '/map/bdgd',
				component: { render (c) {return c('router-view') }},
				children: [ { path: 'bdmap', name: '百度地图', icon:'md-paw', component: BdMap, } ,
							{ path: 'gdmap', name: '高德地图', icon:'md-planet', component: GdMap, } ,
				]
			},
		]
	}
]
