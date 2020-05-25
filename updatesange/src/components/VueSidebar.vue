<template>
	<div class="sidebar">
		<nav class="sidebar-nav"><sidebar-item :routes="siderbar_routers.children"></sidebar-item></nav>
	</div>
</template>

<script>
import { mapGetters } from 'vuex';
import SidebarItem from './SidebarItem';
export default {
	name: 'sidebar',
	data() {
		return {
			
		}
	},
	components: { SidebarItem },
	computed: {
		// console.log("sidebar sidebar_routers = ", this.$store.getters.sidebar_routers)siderbar_routers
		...mapGetters(['siderbar_routers'])
		// sidebar_routers: sidebar_routers
	},
	methods: {
		handleClick(e) {
			e.preventDefault();
			e.target.parentElement.classList.toggle('open');
		}
	},
	mounted(){
		//var sidebar_routers=this.$store.getters.sidebar_routers
		//console.log("sidebar_routers = ", sidebar_routers)
	}
};
</script>

<style lang="css">
.nav-link {
  cursor:pointer;
}
/*
.nav-item {
	padding-left: 10px !important;
}*/
</style>
