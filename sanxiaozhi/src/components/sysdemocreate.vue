<template>
	<div class="app-container">
		<el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
			<el-form-item label="字典名称" prop="dictName">
				<el-input v-model="queryParams.dictName" placeholder="请输入字典名称" clearable size="small" style="width: 200px" @keyup.enter.native="handleQuery" />
			</el-form-item>
			<el-form-item label="字典类型" prop="dictType">
				<el-input v-model="queryParams.dictType" placeholder="请输入字典类型" clearable size="small" style="width: 200px" @keyup.enter.native="handleQuery" />
			</el-form-item>
			<el-form-item label="字典状态 : " prop="status">
				<el-select v-model="queryParams.status" placeholder="字典状态" clearable size="small" style="width: 200px">
					<el-option v-for="item in statusOptions" :value="item.value" :key="item.label" name="sendValue">{{ item.value }}</el-option>
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
			<el-col :span="1.5"><el-button type="warning" icon="el-icon-check" size="mini" @click="clickAll">全选/反选</el-button></el-col>
		</el-row>

		<el-table ref="multipleTable" v-loading="loading" :data="typeList" @selection-change="handleSelectionChange">
			<el-table-column type="selection" width="55" align="center" />
			<el-table-column label="字典编号" width="80" align="center" prop="dictId" />
			<el-table-column label="字典名称" align="center" prop="dictName" :show-overflow-tooltip="true" />
			<el-table-column label="字典类型" align="center" prop="dictType" :show-overflow-tooltip="true" />
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
		<pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageIndex" :limit.sync="queryParams.pageSize" @pagination="getList" />
		<!--
		<el-pagination :page-size="10" :pager-count="2" layout="prev, pager, next" :total="2"></el-pagination>
		-->
		<!-- 添加或修改参数配置对话框 -->
		<el-dialog :title="title" :visible.sync="open" width="500px">
			<el-form ref="form" :model="form" :rules="rules" label-width="80px">
				<el-form-item label="字典名称" prop="dictName"><el-input v-model="form.dictName" placeholder="请输入字典名称" /></el-form-item>
				<el-form-item label="字典类型" prop="dictType"><el-input v-model="form.dictType" placeholder="请输入字典类型" /></el-form-item>
				<el-form-item label="状态" prop="status">
					<el-radio-group v-model="form.status">
						<el-radio v-for="item in statusOptions" :key="item.value" :label="item.label">{{ item.value }}</el-radio>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="备注" prop="remark"><el-input v-model="form.remark" type="textarea" placeholder="请输入内容" /></el-form-item>
			</el-form>

			<div slot="footer" class="dialog-footer">
				<el-button type="primary" @click="submitForm">确 定</el-button>
				<el-button @click="cancel">取 消</el-button>
			</div>
		</el-dialog>
	</div>
</template>

<script>
import { showMsg } from '../api/message';
import { formatJson } from '../util/converjson';
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
			// 全选、反选
			checkAll: false,
			// 字典表格数据
			typeList: [],
			// 全部数据为了全选等功能使用
			alltypeList: [],
			sendValue: '',
			// 是否显示弹出层
			open: false,
			// 弹出层标题
			title: '',
			// 表单参数
			form: {},
			// 表单校验
			rules: {
				dictName: [{ required: true, message: '字典名称不能为空', trigger: 'blur' }],
				dictType: [{ required: true, message: '字典类型不能为空', trigger: 'blur' }]
			},
			// 查询的条件
			query_params: {
				qdictName: 'All',
				qdictType: 'All',
				qstatus: 'All'
			},
			// 状态数据字典
			statusOptions: [{ value: 'All', label: '-1' }, { value: '正常', label: '0' }, { value: '异常', label: '1' }],
			// 查询参数
			queryParams: {
				pageIndex: 1,
				pageSize: 10,
				dictName: 'All',
				dictType: 'All',
				status: 'All'
			}
		};
	},
	created() {
		this.getList();
	},
	methods: {
		// 把状态栏进行格式化
		statusFormat(row) {
			var status_dict = { '0': '正常', '1': '异常' };
			return status_dict[row.status];
		},
		// 取消按钮
		cancel() {
			this.open = false;
			this.reset();
		},
		/** 查询字典类型列表 */
		getList() {
			// 当前的查询效率一般，应该去分页查询，直接查询某一页的数据
			var tmpDataArray = [];
			var alltmpDataArray = [];
			// showMsg(this, this.query_params);
			this.$post('/system/dict', this.query_params).then(response => {
				// console.log(response);
				var tmpDataInfo = response.data;
				var datalen = Object.keys(tmpDataInfo).length;
				// 是不是可以取固定长度的呢
				var start = this.queryParams.pageSize * (this.queryParams.pageIndex - 1);
				var end = start + 10;
				if (end > datalen) {
					end = datalen;
				}
				// 只取固定某一页的数据
				for (var key = start; key < end; key++) {
					var item = tmpDataInfo[key];
					tmpDataArray.push(item);
				}

				for (var allkey in tmpDataInfo) {
					var allitem = tmpDataInfo[allkey];
					alltmpDataArray.push(allitem);
				}
				// console.log(tmpDataArray);
				this.total = datalen;
				this.loading = false;
				this.typeList = tmpDataArray;
				this.alltypeList = alltmpDataArray;
			});
		},
		/** 搜索按钮操作 */
		handleQuery() {
			this.queryParams.pageIndex = 1;

			// showMsg(this, this.typeList);
			// 根据条件进行过滤
			// alert(this.queryParams.dictName)
			this.query_params['qdictName'] = this.queryParams.dictName === '' ? 'All' : this.queryParams.dictName;
			this.query_params['qdictType'] = this.queryParams.dictType === '' ? 'All' : this.queryParams.dictType;
			this.query_params['qstatus'] = this.queryParams.status === '' ? 'All' : this.queryParams.status;
			this.getList();
		},
		/** 重置按钮操作 */
		resetQuery() {
			this.dateRange = [];
			// this.resetForm('queryForm')
			this.resetForm('queryForm');
			this.handleQuery();
			// showMsg(this, '重置功能未完成');
		},
		// 表单重置
		reset() {
			this.form = {
				dictId: 'All',
				dictName: 'All',
				dictType: 'All',
				status: '0',
				remark: undefined
			};
			this.resetForm('form');
		},
		/** 新增按钮操作 */
		handleAdd() {
			this.reset();
			this.open = true;
			this.title = '添加字典类型';
		},
		/** 修改按钮操作 */
		handleUpdate(row) {
			this.reset();
			this.open = true;
			this.title = '修改字典类型';
			console.log(row);
			// const dictId = row.dictId || this.ids;
			/*
			getType(dictId).then(response => {
				this.form = response.data;
				this.open = true;
				this.title = '修改字典类型';
			});*/
			showMsg(this, '请求发送接口自己完善即可');
		},
		/** 删除按钮操作 */
		handleDelete(row) {
			const dictIds = row.dictId || this.ids;
			this.$confirm('是否确认删除字典编号为"' + dictIds + '"的数据项?', '警告', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			}).then(() => {
				showMsg(this, '删除功能未完成');
			});
		},
		/** 导出按钮操作  row*/
		handleExport() {
			var tmpDataList = [];
			if (this.checkAll == true) {
				// 这个就是全选了
				tmpDataList = this.alltypeList;	
			} else {
				// 选择当前页的某一部分 => this.typeList
				/*
				alert(this.ids);
				var dictIds = new String(row.dictId || this.ids);
				alert(dictIds);
				var idArr = dictIds.split(',');
				var idLength = idArr.length;
				//var sum = 0;
				console.log("alltypeList = ", this.alltypeList);
				for (var i = 0; i < idLength; i++) {
					// alert(idArr[i]);
					tmpDataList.push(this.alltypeList[idArr[i]])
				}
				//alert(sum);*/
				tmpDataList = this.typeList;
			}
			// console.log("tmpDataList = ", tmpDataList);

			this.$confirm('是否确认导出所有类型数据项?', '警告', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			}).then(() => {
				this.downloadLoading = true;
				import('../export/Export2Excel').then(excel => {
					const tHeader = ['字典编号', '字典名称', '字典类型', '状态', '备注'];
					const filterVal = ['dictId', 'dictName', 'dictType', 'status', 'remark'];
					const list = tmpDataList;
					// console.log("list = ", list);

					const data = formatJson(filterVal, list);
					// console.log("data = ", data);
					excel.export_json_to_excel({
						header: tHeader,
						data,
						filename: '字典管理',
						autoWidth: true, // Optional
						bookType: 'xlsx' // Optional
					});
					this.downloadLoading = false;
				});
			});
		},
		// 多选框选中数据
		handleSelectionChange(selection) {
			this.ids = selection.map(item => item.dictId);
			this.single = selection.length !== 1;
			this.multiple = !selection.length;
			// showMsg(this, this.ids);
		},
		// 全选按钮
		clickAll() {
			if (this.checkAll == false) {
				// alert('全选');
				this.checkAll = true;
			} else {
				// alert('反选');
				this.checkAll = false;
			}
			// alert(this.alltypeList.length);
			this.$refs.multipleTable.toggleAllSelection();
		},
		/** 提交按钮 */
		submitForm: function() {
			this.$refs['form'].validate(valid => {
				if (valid) {
					if (this.form.dictId !== undefined) {
						showMsg(this, '请求发送接口自己完善即可');
						console.log(this.form);
					} else {
						showMsg(this, this.form.dictId + ' => 异常');
						console.log(this.form);
					}
				}
			});
		}
	}
};
</script>
