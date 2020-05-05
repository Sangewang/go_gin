import axios from 'axios'
import Qs from 'qs'

export function getServer() {
	axios
		.post('http://127.0.0.1:9090/monitor/local', {
			serverlist: "local"
		})
		.then(successResponse => {
			console.log(successResponse);
			if (successResponse.data.code === 0) {
				return successResponse.data.data;
			} else {
				alert("not ok")
			}
		})
		.catch(failResponse => {
			alert(failResponse);
		});
}


/**
 * post方法，对应post请求
 * @param {String} url [请求的url地址]
 * @param {Object} params [请求时携带的参数]
 */
export function getNServer(url, params) {
	return new Promise((resolve, reject) => {
		axios.post(url, Qs.stringify(params))
			.then(res => {
				resolve(res.data);
			})
			.catch(err => {
				reject(err.data)
			})
	});
}
