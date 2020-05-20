import fetch from '../utils/fetch.js';

// 发送请求到后端进行用户名、密码验证
export function loginByEmail(email, password) {
	const data = {
		email,
		password
	};
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
	return fetch({
		url: '/user/info',
		method: 'get',
		params: {
			token
		}
	});
}
