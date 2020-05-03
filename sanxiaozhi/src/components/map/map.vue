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
								<i class="el-icon-s-home"></i>
								系统简介
							</el-menu-item>
							<el-menu-item index="2">
								<i class="el-icon-upload"></i>
								数据上传
							</el-menu-item>
							<el-menu-item index="3">
								<i class="el-icon-s-promotion"></i>
								任务创建
							</el-menu-item>
							<el-menu-item index="4">
								<i class="el-icon-search"></i>
								任务查询
							</el-menu-item>
							<el-menu-item index="5">
								<i class="el-icon-setting"></i>
								设备状态
							</el-menu-item>
							<el-menu-item index="6">
								<i class="el-icon-s-data"></i>
								评测结果
							</el-menu-item>
							<el-submenu index="7">
								<template slot="title">
									<i class="el-icon-location"></i>
									<span slot="title">备用功能</span>
								</template>
								<el-menu-item-group>
									<el-submenu index="7-1">
										<span class="el-icon-picture" slot="title">图片展示</span>
										<el-menu-item index="7-1-1">图片列表</el-menu-item>
										<el-menu-item index="7-1-2">备用</el-menu-item>
									</el-submenu>
								</el-menu-item-group>
								<el-menu-item-group>
									<el-submenu index="7-2">
										<span slot="title">备用2</span>
										<el-menu-item index="7-2-1">备用2-1</el-menu-item>
										<el-menu-item index="7-2-2">备用2-2</el-menu-item>
									</el-submenu>
								</el-menu-item-group>
							</el-submenu>
						</el-menu>
					</div>
				</el-col>
				<el-col :xs="20" :sm="20" :md="20" :lg="20">
					<div style="border: 1px; padding:5px; margin:2px; background-color: antiquewhite">
						<el-breadcrumb separator="/">
							<el-breadcrumb-item :to="{ path: '/map' }">首页</el-breadcrumb-item>
							<el-breadcrumb-item>{{ breadcrumbItems }}</el-breadcrumb-item>
						</el-breadcrumb>
					</div>
					<div style="margin-top:10px"><router-view></router-view></div>
				</el-col>
			</el-row>
		</div>
		<AppFooter />
	</div>
</template>
	
<script type="text/ecmascript-6">
	import AppHeader from '../Header'
	import AppFooter from '../Footer'
export default {
  data(){
    return {
      searchCriteria: '',
      breadcrumbItems: '系统简介',
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
				this.$router.push('/map/introduction');
				this.breadcrumbItems  = '系统简介'
				break;
			case '2':
				this.$router.push('/map/upload');
				this.breadcrumbItems  = '数据上传'
				break;
			case '7-1-1':
				this.$router.push('/map/picshow');
				this.breadcrumbItems  = '图片列表'
				break;			
      }
    },

  },
}
</script>
