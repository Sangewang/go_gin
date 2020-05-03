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
			ref="upload"
			action="https://jsonplaceholder.typicode.com/posts/"
			:on-preview="handlePreview"
			:on-remove="handleRemove"
			:file-list="fileList"
			:auto-upload="false"
			:multiple="true"
			list-type="picture"
		>
			<el-button slot="trigger" size="medium" type="primary">选取文件</el-button>
			<el-button style="margin-left: 10px;" size="medium" type="success" @click="submitUpload">上传到服务器</el-button>
			<el-button style="margin-left: 10px;" size="medium" type="danger" @click="clearFiles">重置</el-button>
			<div style="margin: 20px;"></div>
			<div slot="tip" class="el-upload__tip">文件列表:</div>
		</el-upload>
		
	</div>
</template>

<script>
export default {
	data() {
		return {
			labelPosition: 'right',
			formLabelAlign: {
				name: '',
				region: '',
				type: ''
			},
			fileList: [],
		};
	},
	methods: {
		submitUpload() {
			this.$refs.upload.submit();
			// 文件上传的接口 去go的file模块进行controller的处理
		},
		handleRemove(file, fileList) {
			console.log(file, fileList);
		},
		handlePreview(file) {
			console.log(file);
		},
		clearFiles() {
			this.fileList = [];
		}
	}
};
</script>

<style>
/*
.el-input__inner {
	width: 40% !important;
}
*/
.el-upload__tip {
	font-size: 20px;
	color: #606266;
	margin-top: 7px;
}



</style>
