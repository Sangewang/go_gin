import router from './router'
import store from './store'
import NProgress from 'nprogress' // Progress 进度条
import 'nprogress/nprogress.css' // Progress 进度条样式


// permissiom judge
function hasPermission(roles, permissionRoles) {
  if (roles.indexOf('admin') >= 0) return true // admin权限 直接通过
  if (!permissionRoles) return true
  return roles.some(role => permissionRoles.indexOf(role) >= 0)
}

// register global progress.
const whiteList = ['/login', '/authredirect']// 不重定向白名单

router.beforeEach((to, from, next) => {
	// alert("router.beforeEach")
	NProgress.start() // 开启Progress
	// 先判断是否有token
	if (store.getters.token) {
		// alert(store.getters.token)
		console.log("store.getters.token = ", store.getters.token)
		console.log("to.path = ", to.path)
		if (to.path === '/login') {
			next({
				path: '/'
			})
		} else {
			// alert(store.getters.roles)
			if (store.getters.roles.length === 0) { // 判断当前用户是否已拉取完user_info信息
				store.dispatch('GetInfo').then(res => { // 拉取user_info
					const data = res.data.data
					// console.log("data = ", data)
					const role = data.role
					store.dispatch('GenerateRoutes', {
						role
					}).then(() => { // 生成可访问的路由
						// console.log("store.getters.addRouters = ", store.getters.addRouters)
						router.addRoutes(store.getters.addRouters) // 动态添加可访问路由表
						next({ ...to
						}) // hack方法 确保addRoutes已完成
					})
				}).catch(() => {
					store.dispatch('LogOut').then(() => {
						next({
							path: '/login'
						})
					})
				})
			} else {
				// alert("获取当前路由")
				store.dispatch('getNowRoutes', to);
				if (hasPermission(store.getters.roles, to.meta.role)) {
					next() //
					console.log("has userinfo")
				} else {
					next({
						path: '/',
						query: {
							noGoBack: true
						}
					})
				}
			}
		}
	} else {
		if (whiteList.indexOf(to.path) !== -1) { // 在免登录白名单，直接进入
			// alert("先到白名单")
			// alert(to.path)
			next()
		} else {
			next('/login') // 否则全部重定向到登录页
			NProgress.done() // 在hash模式下 改变手动改变hash 重定向回来 不会触发afterEach 暂时hack方案 ps：history模式下无问题，可删除该行！
		}
	}
})

router.afterEach(() => {
	// alert("router.afterEach")
	NProgress.done() // 结束Progress
})
