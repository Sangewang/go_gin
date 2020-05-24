import fetch from '../../utils/fetch.js';
import toFormData from '../../utils/formdata.js'

export function getNServer(params) {
	// 当前params参数无用
	let data = toFormData({
		"params": params,
	});
	return fetch({
		url: '/monitor/local',
		method: 'post',
		data
	});
}



