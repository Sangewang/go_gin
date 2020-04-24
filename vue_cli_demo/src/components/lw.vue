<template>
	<h2 class="red">lw test {{dynamicSegment}}</h2>
</template>

<script>
// 脚本区域
export default {
	name: 'lw',
	props: {
		msg: {
			type: String,
			default: "test msg"
		}
	},
	methods: {
		
	},
	data() {
		return {
			dynamicSegment: ''
		}
	},
	// 组件切换的监听
	watch: {
		$route (to, from) {
			// to表示要去的组件、from表示从哪个组件来
			console.log(to);
			console.log(from);
			this.dynamicSegment = to.params.id
		}
	}
}
</script>

<style>
.red{color:#ff0000;}
</style>
