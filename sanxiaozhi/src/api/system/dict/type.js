// 发送请求到gin后端查询数据
// 查询字典类型列表
/*
export function listType(query) {
  return request({
    url: '/api/v1/dict/typelist',
    method: 'get',
    params: query
  })
}*/
// 
export function listType(query) {
	alert(query);
	var tmpDataArray = [];
	this.$fetch(query).then(response => {
		// console.log(response);
		var tmpDataInfo = response.data;
		for (var key in tmpDataInfo) {
			var item = tmpDataInfo[key];
			tmpDataArray.push(item);
		}
	});
	return tmpDataArray;
}
