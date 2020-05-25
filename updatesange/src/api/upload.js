import fetch from '../utils/fetch.js';
import toFormData from '../utils/formdata.js'

// 发送上传图片请求到后端
export function PicUpload(sendfileList) {	
	let data = toFormData({
		"picfile": sendfileList
	});
	return fetch({
		url: '/pic/upload',
		method: 'post',
		data
	});
}

// 发送上传图片请求到后端
export function DataUpload(sendfileList) {	
	let data = toFormData({
		"filelist": sendfileList
	});
	return fetch({
		url: '/file/upload',
		method: 'post',
		data
	});
}


// 下载图片
export function PicDownload() {	
	return fetch({
		url: '/pic/download',
		method: 'get',
		params: '',
	});
}
