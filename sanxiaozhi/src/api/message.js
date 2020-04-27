/**
 * 封装showMsg方法
 * @param url
 * @param data
 * @returns {Promise}
 */
export function showMsg (object, msg) {
	object.$alert(msg, '提示信息', {
		confirmButtonText: '确定',
		/*
		callback: action => {
			this.$message({
				type: 'info',
				message: `action: ${ action }`
			});
		}*/
	});
}
