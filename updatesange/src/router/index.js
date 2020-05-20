import Vue from 'vue';
import Router from 'vue-router';

import login from '../views/login/index.vue'

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
