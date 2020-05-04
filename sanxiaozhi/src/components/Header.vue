<template>
	<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
		<el-menu-item index="1"><router-link tag="div" to="/page1">定位分析平台</router-link></el-menu-item>
		<el-menu-item index="5"><router-link tag="div" to="/map">地图分析平台</router-link></el-menu-item>
		<el-menu-item index="3" disabled>消息中心</el-menu-item>
		<el-menu-item index="4"><a href="https://www.baidu.com" target="_blank">百度知道</a></el-menu-item>
		<el-submenu index="2">
			<template slot="title">
				<router-link tag="div" to="/map/introduction">我的工作台</router-link>
			</template>
			<el-menu-item index="2-1"><a href="https://github.com/Sangewang/go_gin" target="_blank">github</a></el-menu-item>
			<el-menu-item index="2-2"><a href="https://blog.csdn.net/u011559236" target="_blank">csdn</a></el-menu-item>
			<el-submenu index="2-3">
				<template slot="title">
					map
				</template>
				<el-menu-item index="2-3-1"><a href="http://lbsyun.baidu.com/" target="_blank">bdmap</a></el-menu-item>
				<el-menu-item index="2-3-2"><a href="https://lbs.amap.com/" target="_blank">gdmap</a></el-menu-item>
			</el-submenu>
			<el-submenu index="2-4">
				<template slot="title">
					go/vue/...
				</template>
				<el-menu-item index="2-4-1"><a href="http://www.topgoer.com/" target="_blank">go</a></el-menu-item>
				<el-menu-item index="2-4-2"><a href="https://cn.vuejs.org/" target="_blank">vue</a></el-menu-item>
				<el-menu-item index="2-4-3"><a href="https://element.eleme.cn/" target="_blank">element-ui</a></el-menu-item>
				<el-menu-item index="2-4-4"><a href="https://v-charts.js.org/" target="_blank">vcharts</a></el-menu-item>
				<el-menu-item index="2-4-5"><a href="https://www.iviewui.com/" target="_blank">iview</a></el-menu-item>
			</el-submenu>
		</el-submenu>
	</el-menu>
</template>
<script>
export default {
	data() {
		return {
			activeIndex: '1'
		};
	},
	methods: {
		handleSelect(key, keyPath) {
			console.log(key, keyPath);
		}
	}
};
</script>
<style></style>
