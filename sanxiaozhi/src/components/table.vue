<template>
	<div>
		<!--
		<div class="demo-input-size" style="float:left;padding:5px;color:white;margin-left:2%;width:15%">
			<el-input placeholder="请输入内容" suffix-icon="el-icon-date" v-model="input1"></el-input>
			<el-input size="medium" placeholder="请输入内容" suffix-icon="el-icon-date" v-model="input2" ></el-input>
			<el-input size="small" placeholder="请输入内容" suffix-icon="el-icon-date" v-model="input3"></el-input>
			<el-input size="mini" placeholder="请输入内容" suffix-icon="el-icon-date" v-model="input4"></el-input>
		</div>
		-->
		<el-row>
			<el-col :span="12" style=" width=50;float:left;padding:15px; margin-left:5%"><el-input v-model="username" placeholder="姓名"></el-input></el-col>
			<el-col :span="12" style="float:left;padding:15px; margin-left:5%"><el-button type="primary" icon="el-icon-search" @click="searchInfo">搜索</el-button></el-col>
		</el-row>
		<el-table ref="multipleTable" :data="tableData" border tooltip-effect="dark" style="width: 100%" @selection-change="handleSelectionChange">
			<el-table-column type="selection" width="50"></el-table-column>
			<el-table-column prop="uid" label="序号" width="80"></el-table-column>
			<el-table-column prop="create_time" label="日期" width="180"></el-table-column>
			<el-table-column prop="username" label="姓名" width="100"></el-table-column>
			<el-table-column prop="province" label="省份" width="100"></el-table-column>
			<el-table-column prop="city" label="市区" width="100"></el-table-column>
			<el-table-column prop="address" label="地址" width="400"></el-table-column>
			<el-table-column prop="code" label="邮编" width="120"></el-table-column>
			<el-table-column label="操作" width="300">
				<template slot-scope="scope">
					<el-button @click="handleClick(scope.row, 'search')" type="text" size="small">查看</el-button>
					<el-button @click="handleClick(scope.row, 'modify')" type="text" size="small">编辑</el-button>
					<el-button @click="handleClick(scope.row, 'delete')" type="text" size="small">删除</el-button>
				</template>
			</el-table-column>
			<el-table-column prop="detail" label="备注" width="200"></el-table-column>
		</el-table>
	</div>
</template>

<script>
import { showMsg } from '../api/message.js';
export default {
	data() {
		// 查询后端的数据
		var dataArray = this.searchInfo();
		return {
			username: '',
			tableData: dataArray,
			multipleSelection: []
		};
	},

	methods: {
		searchInfo() {
			var tmpDataArray = [];
			this.$post('/user/detail').then(response => {
				// console.log(response);
				var tmpDataInfo = response.data;
				for (var key in tmpDataInfo) {
					var item = tmpDataInfo[key];
					tmpDataArray.push(item);
				}
			});
			return tmpDataArray;
		},
		// 信息的增删改查
		handleClick(object, action) {
			if (action == 'search') {
				showMsg(this, '查询 uid = ' + object.uid + '的信息');
			} else if (action == 'delete') {
				showMsg(this, '删除 uid = ' + object.uid + '的信息');
			} else if (action == 'modify') {
				showMsg(this, '修改 uid = ' + object.uid + '的信息');
			} else {
				showMsg(this, '未知的action = ' + action);
			}
		},
		toggleSelection(rows) {
			if (rows) {
				rows.forEach(row => {
					this.$refs.multipleTable.toggleRowSelection(row);
				});
			} else {
				this.$refs.multipleTable.clearSelection();
			}
		},
		handleSelectionChange(val) {
			this.multipleSelection = val;
		}
	}
};
</script>

<style>
.el-row {
	margin-bottom: 20px;
	text-align: left;
	&:last-child {
		margin-bottom: 0;
	}
}
</style>
