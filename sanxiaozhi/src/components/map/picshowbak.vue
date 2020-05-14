<template>
	<div>
		<el-upload
			action="myUrl"
			:on-change="
				(file, fileList) => {
					handleChange(file, fileList, 1);
				}
			"
			:on-remove="
				(file, fileList) => {
					handleRemove(file, fileList, 1);
				}
			"
			:auto-upload="false"
			:file-list="fileList[0]"
			ref="file1"
		>
			<el-button size="small">选择图片</el-button>
		</el-upload>
		<el-upload
			action="myUrl"
			:on-change="
				(file, fileList) => {
					handleChange(file, fileList, 2);
				}
			"
			:on-remove="
				(file, fileList) => {
					handleRemove(file, fileList, 2);
				}
			"
			:auto-upload="false"
			:file-list="fileList[1]"
			ref="file2"
		>
			<el-button size="small">选择图片</el-button>
		</el-upload>
		<el-upload
			action="myUrl"
			:on-change="
				(file, fileList) => {
					handleChange(file, fileList, 3);
				}
			"
			:on-remove="
				(file, fileList) => {
					handleRemove(file, fileList, 3);
				}
			"
			:auto-upload="false"
			:file-list="fileList[2]"
			ref="file3"
		>
			<el-button size="small">选择图片</el-button>
		</el-upload>
		<el-button @click="submitData">确认</el-button>
	</div>
</template>

<script>
import axios from 'axios';

export default {
	data() {
		return {
			fileList: [[], [], []], //缓存区文件
			uploadFile: [[], [], []] ,//  上传用文件
			temp: ''
		};
	},
	methods: {
		handleChange(file, fileList, type) {
			//  限制单张上传，超过限制即覆盖
			if (fileList.length > 1) {
				fileList.splice(0, 1);
			}
			
			//  校验
			const isLt2M = file.size / 1024 / 1024 < 5;
			if (!isLt2M) {
				this.$message.error('上传图片大小不能超过 5MB!');
				this.removeUploadedFile(type); //  不符合要求删除文件
				return false;
			}

			const isImage = file.raw.type.includes('image');
			if (!isImage) {
				this.$message.error('上传的格式必须是图片!');
				this.removeUploadedFile(type);
				return false;
			}
			//  验证通过之后，将缓存区文件存入上传区文件中
			
			this.temp = file.raw;
		},
		//  从缓存区移除文件
		removeUploadedFile(type) {
			if (type === 0) {
				this.$refs.file1.clearFiles();
			}
			if (type === 1) {
				this.$refs.file2.clearFiles();
			}
			if (type === 2) {
				this.$refs.file3.clearFiles();
			}
		},
		//  删除文件
		handleRemove(file, fileList, type) {
			//  删除文件时要移除缓存区文件和上传区文件
			this.fileList[type] = 0;
			this.uploadFile[type] = [];
		},
		//  上传文件
		submitData() {
			//  校验是否选择文件
				/*
			let fileNum = this.flatten(this.uploadFile).length;
			if (fileNum === 0) {
				this.$error('未选择任何文件!');
				return false;
			}

		
			if (this.formData.files[0]) {
				this.formData.append('file1', this.formData.files[0]);
			}
			if (this.formData.files[1]) {
				this.formData.append('file2', this.formData.files[1]);
			}
			if (this.formData.files[2]) {
				this.formData.append('file2', this.formData.files[2]);
			}
				*/
			//  请求:在headers上务必加上content-Type,指定表单形式发送
			//  使用formdata格式
			let formData = new FormData();
			formData.append('file', this.temp);
			console.log("bak formData = ", formData);

			axios
				.post('http://127.0.0.1:9090/pic/upload', formData, { headers: { 'Content-Type': 'multipart/form-data' } })
				.then(res => {
					console.log(res);
					this.$success('上传图片成功!');
					this.fileList = [0, 0, 0];
					this.uploadFile = [[], [], []];
				})
				.catch(e => {
					console.log(e);
				});
		},
		//  扁平化数组
		flatten(arr) {
			let res = [];
			for (let i = 0; i < arr.length; i++) {
				if (Array.isArray(arr[i])) {
					res = res.concat(this.flatten(arr[i]));
				} else {
					res.push(arr[i]);
				}
			}
			return res;
		}
	}
};
</script>

<style></style>
