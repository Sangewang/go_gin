<template>
	<div class="app-container">
		<el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
			<el-form-item label="字典名称" prop="dictName">
				<el-input v-model="queryParams.dictName" placeholder="请输入字典名称" clearable size="small" style="width: 240px" @keyup.enter.native="handleQuery" />
			</el-form-item>
			<el-form-item label="字典类型" prop="dictType">
				<el-input v-model="queryParams.dictType" placeholder="请输入字典类型" clearable size="small" style="width: 240px" @keyup.enter.native="handleQuery" />
			</el-form-item>
			<el-form-item label="状态" prop="status">
				<el-select v-model="queryParams.status" placeholder="字典状态" clearable size="small" style="width: 240px">
					<el-option v-for="dict in statusOptions" :key="dict.dictValue" :label="dict.dictLabel" :value="dict.dictValue" />
				</el-select>
			</el-form-item>

			<el-form-item>
				<el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
				<el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
			</el-form-item>
		</el-form>

		<el-row :gutter="10" class="mb8">
			<el-col :span="1.5"><el-button type="primary" icon="el-icon-plus" size="mini" @click="handleAdd">新增</el-button></el-col>
			<el-col :span="1.5"><el-button type="success" icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate">修改</el-button></el-col>
			<el-col :span="1.5"><el-button type="danger" icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete">删除</el-button></el-col>
			<el-col :span="1.5"><el-button type="warning" icon="el-icon-download" size="mini" @click="handleExport">导出</el-button></el-col>
		</el-row>

		<el-table v-loading="loading" :data="typeList" @selection-change="handleSelectionChange">
			<el-table-column type="selection" width="55" align="center" />
			<el-table-column label="字典编号" width="80" align="center" prop="dictId" />
			<el-table-column label="字典名称" align="center" prop="dictName" :show-overflow-tooltip="true"/>
			<el-table-column label="字典类型" align="center" prop="dictType" :show-overflow-tooltip="true"/>
			<el-table-column label="状态" align="center" prop="status" :formatter="statusFormat" />
			<el-table-column label="备注" align="center" prop="remark" :show-overflow-tooltip="true" />
			<el-table-column label="创建时间" align="center" prop="create_time" width="180"></el-table-column>
			<el-table-column label="操作" align="center" class-name="small-padding fixed-width">
				<template slot-scope="scope">
					<el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)">修改</el-button>
					<el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)">删除</el-button>
				</template>
			</el-table-column>
		</el-table>
		<pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageIndex" :limit.sync="queryParams.pageSize" @pagination="getList"/>
		<!--
		<el-pagination :page-size="10" :pager-count="2" layout="prev, pager, next" :total="2"></el-pagination>
		-->
	</div>
</template>

<script>
import { showMsg } from '../api/message.js';
export default {
	name: 'Dict',
	data() {
		return {
			loading: true,
			// 选中数组
			ids: [],
			// 非单个禁用
			single: true,
			// 非多个禁用
			multiple: true,
			// 总条数
			total: 0,
			// 状态数据字典
			statusOptions: [],
			// 字典表格数据
			typeList: [],
			// 查询参数
			queryParams: {
				pageIndex: 1,
				pageSize: 10,
				dictName: undefined,
				dictType: undefined,
				status: undefined
			}
		};
	},
	created() {
		this.getList();
	},
	methods: {
		// 把状态栏进行格式化
		statusFormat(row) {
			// showMsg(this, row);
			var status_dict = {'0': '正常', '1':'异常'};
			return status_dict[row.status];
		},
		/** 查询字典类型列表 */
		getList() {
			// 当前的查询效率一般，应该去分页查询，直接查询某一页的数据
			var tmpDataArray = [];
			this.$fetch('/system/dict').then(response => {
				console.log(response);
				var tmpDataInfo = response.data;
				var datalen = Object.keys(tmpDataInfo).length;
				// 是不是可以取固定长度的呢
				var start = this.queryParams.pageSize * (this.queryParams.pageIndex - 1)
				var end   = start + 10
				if (end > datalen)
				{
					end = datalen
				}
				for (var key = start; key < end; key ++ )
				{
					var item = tmpDataInfo[key];
					tmpDataArray.push(item);
				}
				/*
				for (var key in tmpDataInfo) {
					var item = tmpDataInfo[key];
					tmpDataArray.push(item);
				}*/
				console.log(tmpDataArray);
				this.total = datalen;
				this.loading = false;
				this.typeList = tmpDataArray;
			});
			
			
			
		},
		/** 搜索按钮操作 */
		handleQuery() {
			this.queryParams.pageIndex = 1;
			showMsg(this, '搜索功能未完成');
		},
		/** 重置按钮操作 */
		resetQuery() {
			this.dateRange = [];
			showMsg(this, '重置功能未完成');
		},
		/** 新增按钮操作 */
		handleAdd() {
			showMsg(this, '新增功能未完成');
		},
		/** 修改按钮操作 */
		handleUpdate() {
			showMsg(this, '修改功能未完成');
		},
		/** 删除按钮操作 */
		handleDelete() {
			showMsg(this, '删除功能未完成');
		},
		handleExport() {
			showMsg(this, '导出功能未完成');
		},
		// 多选框选中数据
		handleSelectionChange(selection) {
			this.ids = selection.map(item => item.dictId);
			this.single = selection.length !== 1;
			this.multiple = !selection.length;
			// showMsg(this, this.ids);
		}
	}
};
</script>
