<template>
	<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
		<el-menu-item index="1"> <router-link tag="div" to='/page1'>定位分析平台</router-link></el-menu-item>
		<el-menu-item index="5"> <router-link tag="div" to='/'>地图分析平台</router-link></el-menu-item>
		<el-menu-item index="3" disabled>消息中心</el-menu-item>
		<el-menu-item index="4"><a href="https://www.baidu.com" target="_blank">百度知道</a></el-menu-item>
		<el-submenu index="2">
			<template slot="title">
			<router-link tag="div" to='/'>我的工作台</router-link>
			</template>
			<el-menu-item index="2-1">选项1</el-menu-item>
			<el-menu-item index="2-2">选项2</el-menu-item>
			<el-menu-item index="2-3">选项3</el-menu-item>
			<el-submenu index="2-4">
				<template slot="title">
					选项4
				</template>
				<el-menu-item index="2-4-1">选项1</el-menu-item>
				<el-menu-item index="2-4-2">选项2</el-menu-item>
				<el-menu-item index="2-4-3">选项3</el-menu-item>
			</el-submenu>
		</el-submenu>
	</el-menu>
</template>
<script>
export default {
	data() {
		return {
			activeIndex: '1',
			activeIndex2: '1'
		};
	},
	methods: {
		handleSelect(key, keyPath) {
			console.log(key, keyPath);
		}
	}
};
</script>
