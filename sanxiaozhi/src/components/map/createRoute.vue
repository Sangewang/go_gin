<template>
	<div class="app">
		<AppHeader />
		<div class="app-body">
			路线生成页面
		</div>
		<AppFooter />
	</div>
</template>
<script type="text/ecmascript-6">
import AppHeader from '../Header.vue'
import AppFooter from '../Footer.vue'
// import Sidebar from '../sidebar'
export default {
	name: 'full',
	components: {
		AppHeader,
		AppFooter
		// Sidebar
	},
	data(){
	return {
		searchCriteria: '',
		breadcrumbItems: '路线生产',
	}
  },

  methods:{
    handleIconClick(ev) {
      console.log(ev);
    }
  },
}
</script>
