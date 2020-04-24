<template>
	<div class="">
		<el-header>Header</el-header>
		<el-container style="border: 1px solid #eee">
			<el-aside style=":rgb(238, 241, 246)">
				<el-radio-group v-model="isCollapse" style="margin-bottom: 20px;">
					<el-radio-button :label="false">展开</el-radio-button>
					<el-radio-button :label="true">收起</el-radio-button>
				</el-radio-group>
				<el-menu default-active="1-4-1" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" :collapse="isCollapse">
					<el-submenu index="1">
						<template slot="title"><i class="el-icon-loading"></i><span slot="title">导航一</span></template>
						<el-menu-item-group>
							<template slot="title">分组一</template>
							<el-menu-item index="1-1">选项1</el-menu-item>
							<el-menu-item index="1-2">选项2</el-menu-item>
						</el-menu-item-group>
						<el-menu-item-group title="分组2">
							<el-menu-item index="1-3">选项3</el-menu-item>
						</el-menu-item-group>
						<el-submenu index="1-4">
							<template slot="title">选项4</template>
							<el-menu-item index="1-4-1">选项4-1</el-menu-item>
						</el-submenu>
					</el-submenu>
					<el-submenu index="2">
						<template slot="title"><i class="el-icon-menu"></i><span slot="title">导航二</span></template>
						<el-menu-item-group>
							<template slot="title">分组一</template>
							<el-menu-item index="2-1">选项1</el-menu-item>
							<el-menu-item index="2-2">选项2</el-menu-item>
						</el-menu-item-group>
						<el-menu-item-group title="分组2">
							<el-menu-item index="2-3">选项3</el-menu-item>
						</el-menu-item-group>
						<el-submenu index="2-4">
							<template slot="title">选项4</template>
							<el-menu-item index="2-4-1">选项4-1</el-menu-item>
						</el-submenu>
					</el-submenu>
					<el-submenu index="3">
						<template slot="title"><i class="el-icon-setting"></i><span slot="title">导航三</span></template>
						<el-menu-item-group>
							<template slot="title">分组一</template>
							<el-menu-item index="3-1">选项1</el-menu-item>
							<el-menu-item index="3-2">选项2</el-menu-item>
						</el-menu-item-group>
						<el-menu-item-group title="分组2">
							<el-menu-item index="3-3">选项3</el-menu-item>
						</el-menu-item-group>
						<el-submenu index="3-4">
							<template slot="title">选项4</template>
							<el-menu-item index="3-4-1">选项4-1</el-menu-item>
						</el-submenu>
					</el-submenu>
					<el-submenu index="4">
						<template slot="title"><i class="el-icon-s-promotion"></i><span slot="title">导航四</span></template>
						<el-menu-item-group>
							<template slot="title">分组一</template>
							<el-menu-item index="4-1">选项1</el-menu-item>
							<el-menu-item index="4-2">选项2</el-menu-item>
						</el-menu-item-group>
					</el-submenu>
					<el-submenu index="5">
						<template slot="title"><i class="el-icon-s-promotion"></i><span slot="title">日历</span></template>
						<el-calendar v-model="value">
						</el-calendar>
					</el-submenu>
				</el-menu>
			</el-aside>

			<el-container>
				<el-header style="text-align: right; font-size: 12px">
					<el-dropdown>
						<i class="el-icon-setting" style="margin-right: 15px"></i>
						<el-dropdown-menu slot="dropdown">
							<el-dropdown-item>查看</el-dropdown-item>
							<el-dropdown-item>新增</el-dropdown-item>
							<el-dropdown-item>删除</el-dropdown-item>
						</el-dropdown-menu>
					</el-dropdown>
					<!--从login.vue获取到的登录人信息，后续会有权限认证等-->
					<span>{{ this.$userInfo.username }}</span>
				</el-header>

				<el-main>
					<el-table :data="tableData" border style="width: 100%">
						<el-table-column fixed prop="uid" label="序号" width="80">
						</el-table-column>
						<el-table-column fixed prop="create_time" label="日期" width="180">
						</el-table-column>
						<el-table-column prop="username" label="姓名" width="100">
						</el-table-column>
						<el-table-column prop="province" label="省份" width="100">
						</el-table-column>
						<el-table-column prop="city" label="市区" width="100">
						</el-table-column>
						<el-table-column prop="address" label="地址" width="400">
						</el-table-column>
						<el-table-column prop="code" label="邮编" width="120">
						</el-table-column>
						<el-table-column label="操作" width="300">
							<template slot-scope="scope">
								<el-button @click="handleClick(scope.row, 'search')" type="text" size="small">查看</el-button>
								<el-button @click="handleClick(scope.row, 'modify')" type="text" size="small">编辑</el-button>
								<el-button @click="handleClick(scope.row, 'delete')" type="text" size="small">删除</el-button>
							</template>
						</el-table-column>
						<el-table-column prop="detail" label="备注" width="200">
						</el-table-column>
					</el-table>
				</el-main>
			</el-container>
		</el-container>
		<el-footer>Footer</el-footer>
		<router-view></router-view>
	</div>
</template>

<style>
	.el-menu-vertical-demo:not(.el-menu--collapse) {
		width: 200px;
		min-height: 500px;
		text-align: left;
	}

	.el-header,
	.el-footer {
		background-color: #909399;
		color: #333;
		text-align: center;
		line-height: 60px;
	}

	.el-aside {
		width: 200px;
		background-color: #909399;
		color: #333;
	}
</style>

<script>
	// import axios from 'axios'
	export default {
		data() {
			// 在login的时候已经读取了后端的数据
			var dataArray = [];
			for (var key in this.$dataInfo) {
				var item = this.$dataInfo[key];
				console.log(item); //AA,BB,CC,DD
				dataArray.push(item)
			}
			return {
				// tableData: Array(10).fill(item),
				tableData: dataArray,
				isCollapse: true,
				value: new Date()
			}
		},
		methods: {
			handleOpen(key, keyPath) {
				console.log(key, keyPath);
			},
			handleClose(key, keyPath) {
				console.log(key, keyPath);
			},
			// 信息的增删改查
			handleClick(object, action) {
				if (action == 'search') {
					//alert("查询 id = " + object.id + "的信息");
					this.showMsg("查询 uid = " + object.uid + "的信息");
				} else if (action == 'delete') {
					this.showMsg("删除 uid = " + object.uid + "的信息");
				} else if (action == 'modify') {
					this.showMsg("修改 uid = " + object.uid + "的信息");
				} else {
					this.showMsg("未知的action = " + action)
				}

			},
			showMsg(msg) {
				this.$alert(msg, '提示信息', {
					confirmButtonText: '确定',
					/*
					callback: action => {
						this.$message({
							type: 'info',
							message: `action: ${ action }`
						});
					}*/
				});
			},
			randomNum(min, max) {
				return Math.floor(Math.random() * (max - min) + min)
			},

		}
	};
</script>
