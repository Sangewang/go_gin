<template>
	<div style="background-color: #EBEBEB;min-height:900px">
		<div style="width:100%;background-color: #636363; overflow: hidden">
			<AppHeader />
			<!--
			<span class="demonstration" style="float:left;padding-top:10px;color:white;margin-left:1%">后台管理系统</span>
			<span class="demonstration" style="float:left;padding:5px;color:white;margin-left:2%;width:15%">
				<el-input placeholder="请输入" icon="search" v-model="searchCriteria" :on-icon-click="handleIconClick"></el-input>
			</span>
			<span class="demonstration" style="float:right;padding-top:10px;margin-right:1%">
				<el-dropdown trigger="click">
					<span class="el-dropdown-link" style="color:white">
						admin
						<i class="el-icon-caret-bottom el-icon--right"></i>
					</span>
					<el-dropdown-menu slot="dropdown">
						<el-dropdown-item>个人信息</el-dropdown-item>
						<el-dropdown-item>退出登录</el-dropdown-item>
					</el-dropdown-menu>
				</el-dropdown>
			</span>
			-->
		</div>
		
		<div style="margin-top:5px">
			<el-row :gutter="10">
				<el-col :xs="4" :sm="4" :md="4" :lg="4">
					<div>
						<el-menu default-active="1" class="el-menu-vertical-demo" style="min-height:900px" @select="handleSelect">
							<el-menu-item index="1">
								<i class="el-icon-setting"></i>
								基础信息
							</el-menu-item>
							<el-submenu  index="2">
								<template slot="title">
									<i class="el-icon-s-promotion"></i>
									<span slot="title">系统工具</span>
								</template>
								<el-menu-item-group>
									<el-submenu index="2-1">
										<span slot="title">工具模板</span>
										<el-menu-item index="2-1-1">模板创建</el-menu-item>
										<el-menu-item index="2-1-2">模板操作</el-menu-item>
									</el-submenu>
								</el-menu-item-group>
							</el-submenu >
							<el-menu-item index="3">
								<i class="el-icon-s-help"></i>
								系统日历
							</el-menu-item>
							<el-submenu index="4">
								<template slot="title">
									<i class="el-icon-location"></i>
									<span slot="title">系统管理</span>
								</template>
								<el-menu-item-group>
									<el-submenu index="4-1">
										<span slot="title">系统配置</span>
										<el-menu-item index="4-1-1">任务创建</el-menu-item>
										<el-menu-item index="4-1-2">图表操作</el-menu-item>
									</el-submenu>
								</el-menu-item-group>
								<el-menu-item-group>
									<el-submenu index="4-2">
										<span slot="title">日志管理</span>
										<el-menu-item index="4-2-1">登陆日志</el-menu-item>
										<el-menu-item index="4-2-2">操作日志</el-menu-item>
									</el-submenu>
								</el-menu-item-group>
							</el-submenu>
						</el-menu>
					</div>
				</el-col>
				<el-col :xs="20" :sm="20" :md="20" :lg="20">
					<div style="border: 1px; padding:5px; margin:2px; background-color: antiquewhite">
						<el-breadcrumb separator="/">
							<el-breadcrumb-item :to="{ path: '/dashboard' }">首页</el-breadcrumb-item>
							<el-breadcrumb-item>{{ breadcrumbItems }}</el-breadcrumb-item>
						</el-breadcrumb>
					</div>
					<div style="margin-top:10px"><router-view></router-view></div>
				</el-col>
				<!--
				<el-col :xs="20" :sm="20" :md="20" :lg="20">
					<div>
						<div style="border: 1px solid #A6A6A6; border-radius:6px; padding:5px; margin:2px; background-color: white">
							<el-breadcrumb separator="/">
								<el-breadcrumb-item>{{ breadcrumbItems }}</el-breadcrumb-item>
							</el-breadcrumb>
						</div>
					</div>
					<div style="margin-top:10px"><router-view></router-view></div>
				</el-col>
				-->
			</el-row>
		</div>
		<AppFooter />
	</div>
</template>
	
<script type="text/ecmascript-6">
	import AppHeader from './Header'
	import AppFooter from './Footer'
export default {
  data(){
    return {
      searchCriteria: '',
      breadcrumbItems: '基础信息',
    }
  },
  components: {
		AppHeader,
		AppFooter
	},
  methods:{
    handleIconClick(ev) {
      console.log(ev);
    },

    handleSelect(key, keyPath){
		console.log(keyPath);
		switch(key){
			case '1':
				// this.$router.push('/page1');
				this.$router.push('/monitor');
				this.breadcrumbItems  = '基础信息'
				break;
			case '2-1-1':
				this.$router.push('/sysdemocreate')
				this.breadcrumbItems  = '模板创建'
				break;
			case '2-1-2':
				this.$router.push('/sysdemoperate')
				this.breadcrumbItems  = '模板操作'
				break;
			case '3':
				this.$router.push('/calendar')
				this.breadcrumbItems  = '系统日历'
				break;
			// 隶属于 系统管理(4) 下的 系统配置(4-1)
			case '4-1-1':
				this.$router.push('/table')
				this.breadcrumbItems  = '表格操作'
				break;
			case '4-1-2':
				this.$router.push('/chart')
				this.breadcrumbItems  = '图表展示'
				break;
      }
    },

  },
}
</script>
