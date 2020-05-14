<template>
	<div>
		<el-radio-group v-model="labelPosition" size="medium">
			<el-radio-button label="left">左对齐</el-radio-button>
			<el-radio-button label="right">右对齐</el-radio-button>
			<el-radio-button label="top">顶部对齐</el-radio-button>
		</el-radio-group>
		<div style="margin: 20px;"></div>
		<el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign">
			<el-form-item label="名称"><el-input v-model="formLabelAlign.name" style="width: 400px"></el-input></el-form-item>
			<el-form-item label="活动区域"><el-input v-model="formLabelAlign.region" style="width: 400px"></el-input></el-form-item>
			<el-form-item label="活动形式"><el-input v-model="formLabelAlign.type" style="width: 400px"></el-input></el-form-item>
		</el-form>
		<div style="margin: 20px 0;"></div>

		<el-upload
			class="upload-demo"
			action="http://127.0.0.1:9090/pic/upload"
			:on-preview="handlePreview"
			:on-remove="handleRemove"
			:file-list="fileList"
			:multiple="true"
			:auto-upload="false"
			:on-change="handleAddPic"
			list-type="pictures"
		>
			<el-button slot="trigger" size="medium" type="primary">选取文件</el-button>
			<el-button style="margin-left: 10px;" size="medium" type="success" @click="submitUpload">上传到服务器</el-button>
			<el-button style="margin-left: 10px;" size="medium" type="danger" @click="clearFiles">重置</el-button>
			<div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
		</el-upload>

		<div style="margin: 100px 0;"></div>
		<span font-size="20px">图片列表：</span>
		<el-button style="margin-left: 10px;" size="medium" type="success" @click="refreshFiles">刷新</el-button>
		<div class="demo-image__preview" v-for="fit in showfileList" :key="fit.name">
			<span class="demonstration">{{ fit.name }}</span>
			<el-image style="width: 100px; height: 100px" :src="fit.url[0]" :preview-src-list="fit.url"></el-image>
		</div>
	</div>
</template>

<script>
import axios from 'axios';

import blobToFormData from '../../util/toFormData';
export default {
	data() {
		return {
			labelPosition: 'left',
			formLabelAlign: {
				name: '',
				region: '',
				type: ''
			},
			url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
			fileList: [],
			// 图片发送的数据
			sendfileList: [],
			// 图片显示的数据
			showfileList: [
				{ name: '1', url: ['https://fuss10.elemecdn.com/8/27/f01c15bb73e1ef3793e64e6b7bbccjpeg.jpeg'] },
				{ name: '2', url: ['https://fuss10.elemecdn.com/1/8e/aeffeb4de74e2fde4bd74fc7b4486jpeg.jpeg'] }
			],
			gCount: 0
		};
	},
	methods: {
		submitUpload() {
			// console.log(this.fileList);
			// alert(this.file-list);
			// this.$refs.upload.submit();
			let data = blobToFormData({
				picfile: this.sendfileList
			});
			console.log("real data = ", data)
			axios
				.post('http://127.0.0.1:9090/pic/upload', data, {
			'Content-Type': 'multipart/form-data'
			
		})
				.then(successResponse => {
					console.log(successResponse);
					if (successResponse.data.code === 0) {
						// alert(successResponse.data);
					} else {
						// alert(successResponse.data.msg);
					}
				})
				.catch(failResponse => {
					alert(failResponse);
				});
			/*
			axios
				.post('http://127.0.0.1:9090/pic/upload', {
					picfile: this.sendfileList
				})
				.then(successResponse => {
					console.log(successResponse);
					if (successResponse.data.code === 0) {
						alert(successResponse.data);
					} else {
						alert(successResponse.data.msg);
					}
				})
				.catch(failResponse => {
					alert(failResponse);
				});
			*/
		},
		handleAddPic(file, fileList) {
			console.log(fileList); //代表上传的全部图片列表
			console.log(file); // 代表当前最新上传的文件
			
			//  校验
			const isLt2M = file.size / 1024 / 1024 < 5;
			if (!isLt2M) {
				alert('上传图片大小不能超过 5MB!');
				return false;
			}
			this.gCount += 1
			this.sendfileList = fileList.map(file => file.raw);
			// 通过uid唯一标示是可以做到对某一张图片进行赋值的，接口已经准备好
		},
		handleRemove(file, fileList) {
			console.log(file, fileList);
		},
		handlePreview(file) {
			console.log(file);
		},
		//https://element.eleme.cn/#/zh-CN/component/upload 还有一些钩子函数可以进行检查
		clearFiles() {
			this.fileList = [];
			this.showfileList = [];
			this.sendfileList = [];
		},
		refreshFiles() {
			axios
				.get('http://127.0.0.1:9090/pic/download', {})
				.then(successResponse => {
					console.log(successResponse);
					if (successResponse.data.code === 0) {
						// alert(successResponse.data)
						this.showfileList = successResponse.data.data;
					} else {
						alert(successResponse.data.msg);
					}
				})
				.catch(failResponse => {
					alert(failResponse);
				});
		}
	}
};
</script>

<style>
/*
.el-input__inner {
	width: 40% !important;
}

.el-upload__tip {
	font-size: 20px;
	color: #606266;
	margin-top: 7px;
}
*/
</style>
