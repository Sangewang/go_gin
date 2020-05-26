import fetch from '../../utils/fetch.js';

// 发送请求到后端进行用户名、密码验证
export function getSystemDict(data) {	
	return fetch({
		url: '/system/dict',
		method: 'post',
		data
	});
}

// 发送请求到后端进行用户名、密码验证
export function getUserDetail(data) {	
	return fetch({
		url: '/user/detail',
		method: 'post',
		data
	});
}


