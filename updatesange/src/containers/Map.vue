<template>
	<div class="app">
		<AppHeader />
		<div class="app-body">
			<Sidebar />
			<main class="main">
				<breadcrumb :list="list" />
				<div class="container-fluid"><router-view></router-view></div>
			</main>
			<AppAside />
		</div>
		<AppFooter />
	</div>
</template>

<script>
	import AppHeader from '../components/VueHeader.vue'
	import AppFooter from '../components/VueFooter.vue'
	import AppAside from '../components/VueASide.vue'
	import Sidebar from '../components/VueSidebar'
	import Breadcrumb from '../components/Breadcrumb'
	export default {
		name: 'full',
		components: {
			AppHeader,
			AppFooter,
			AppAside,
			Sidebar,
			Breadcrumb
		},
		computed: {
			name () {
				return this.$route.name
			},
			
			list () {
				return this.$route.matched
			}
		}
		
	}
</script>

<style></style>
