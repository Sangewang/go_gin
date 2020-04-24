<template>
	<div id="example-1">
		<input v-model="message" placeholder="edit me">
		<p>Message is : {{ message }}</p>
		<textarea v-model="message2" placeholder="add multiple lines"></textarea>
		<p style="white-space: pre-line;">{{ message2 }}</p>
		<h1>路由 : {{this.mainpath}}</h1>
		<h1>姓名 : {{name}}</h1>
		<h1>性别 : {{sex}}</h1>
		<h1>地址 : {{address}}</h1>
	</div>
	<!--
	写一个提交 一个查询 两个接口
	<div class="getGinData">
		<template v-for="value,index in 5" :key="index">
			<div>{{index}}---{{value}}</div>
		</template>
	</div>
	-->	
</template>

<script>
import axios from 'axios'
export default{
	name: 'getGinData',
	data() {
		return {
			testdata: [],
			name: "",
			sex: "",
			address: "",
			message: "",
			message2: "",
			mainpath:""
		}
	},
	// 钩子函数
	created() {
		axios.get("http://172.22.200.31:9090/user/get/7/admin").then((res) => {
			console.log(res);
			this.testdata = res.data;
			this.name = this.testdata['用户名'];
			console.log(this.name);
			this.sex = this.testdata['性别'];
			this.address = this.testdata['地址'];
		})
	},
	methods: {
		
	},
	// 组件切换的监听
	watch: {
		$route (to, from) {
			// to表示要去的组件、from表示从哪个组件来
			console.log(to);
			console.log(from);
			this.mainpath = to.path;
		}
	}
	
}
</script>

<style>
</style>
