/**
 * @file
 */


export default function toFormData(data) {
	console.log("data = ", data)
    let formData = new FormData();
	// let gCount = 0;
    for (let key in data) {
        let value = data[key];
			
        // 确保数组可以在formdata中传递成功：
        if (Array.isArray(value) || value instanceof FileList) {
            key = "files";
            for (const ele of value) {
				console.log(key)
				// gCount += 1;
                formData.append(key, ele);
            }
        } else {
            formData.append(key, value);
        }
    }
	// formData.append("gCount", gCount);
	console.log("formData = ", formData)
    return formData;
}



// export default function blobToFormData(data) {
// 	console.log("data = ", data)
// 	let formData = new FormData();
// 	// let gCount = 0
// 	for (let key in data) {
// 		let value = data[key]; // picfile
// 		if (Array.isArray(value) || value instanceof FileList) {
// 			key = key + '[]';
// 			for (const ele of value) {
// 				// formData.append(ele['name'], ele['url']);
// 				formData.append(key, ele['raw']);
// 				// console.log("ele = ", ele)
// 				// console.log("ele['name'] = ", ele['name'])
// 			}
// 		} else {
// 			formData.append(key, value);
// 		}
// 	}
// 	console.log("formData = ", formData.values())
// 	return formData;
// }
