import fetch from '../utils/fetch.js';
import toFormData from '../utils/formdata.js'

// 发送请求到后端进行用户名、密码验证
export function loginByEmail(email, password) {	
	let data = toFormData({
		"email": email,
		"password": password
	});
	return fetch({
		url: '/login/loginbyemail',
		method: 'post',
		data
	});
}

// 退出登陆
export function logout() {
	return fetch({
		url: '/login/logout',
		method: 'post'
	});
}

// 获取用户信息
export function getInfo(token) {
	console.log("getInfo token = ", token)
	/*
	let data = toFormData({
		"token": token
	});
	*/
	return fetch({
		url: '/user/info',
		method: 'get',
		params: { token }
	});
}


