<template>
	<div class="app-container">
		<el-row :gutter="20">
			<el-col :span="12" class="card-box">
				<el-card>
					<div slot="header"><span>服务器信息</span></div>
					<div class="el-table el-table--enable-row-hover el-table--medium">
						<table cellspacing="0" style="width: 100%;">
							<thead>
								<tr>
									<th class="is-leaf"><div class="cell">属性</div></th>
									<th class="is-leaf"><div class="cell">值</div></th>
								</tr>
							</thead>
							<tbody>
								<tr>
									<td><div class="cell">服务器名字</div></td>
									<td>
										<div v-if="server.hostname" class="cell">{{ server.hostname }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">服务器架构</div></td>
									<td>
										<div v-if="server.kernelArch" class="cell">{{ server.kernelArch }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">启动时间</div></td>
									<td>
										<div v-if="server.bootTime" class="cell">{{ server.bootTime }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">运行时间</div></td>
									<td>
										<div v-if="server.uptime" class="cell">{{ server.uptime }}天</div>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				</el-card>
			</el-col>
			<el-col :span="12" class="card-box">
				<el-card>
					<div slot="header"><span>磁盘状态</span></div>
					<div class="el-table el-table--enable-row-hover el-table--medium">
						<table cellspacing="0" style="width: 100%;">
							<thead>
								<tr>
									<th class="is-leaf"><div class="cell">属性</div></th>
									<th class="is-leaf"><div class="cell">值</div></th>
								</tr>
							</thead>
							<tbody v-if="server.DiskTotal">
								<tr>
									<td><div class="cell">路径</div></td>
									<td>
										<div class="cell">{{ server.DiskPath }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">总共</div></td>
									<td>
										<div class="cell">{{ server.DiskTotal }}G</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">已用</div></td>
									<td>
										<div class="cell">{{ server.DiskUsed }}G</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">使用率</div></td>
									<td>
										<div class="cell">{{ server.DiskUsage }}%</div>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				</el-card>
			</el-col>
			<el-col :span="12" class="card-box">
				<el-card>
					<div slot="header"><span>内存</span></div>
					<div class="el-table el-table--enable-row-hover el-table--medium">
						<table cellspacing="0" style="width: 100%;">
							<thead>
								<tr>
									<th class="is-leaf"><div class="cell">属性</div></th>
									<th class="is-leaf"><div class="cell">值</div></th>
								</tr>
							</thead>
							<tbody>
								<tr>
									<td><div class="cell">总内存</div></td>
									<td>
										<div v-if="server.MemTotal" class="cell">{{ server.MemTotal }}G</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">已用内存</div></td>
									<td>
										<div v-if="server.MemUsed" class="cell">{{ server.MemUsed }}G</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">使用率</div></td>
									<td>
										<div v-if="server.MemUsage" class="cell" :class="{ 'text-danger': server.MemUsage > 10 }">{{ server.MemUsage }}%</div>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				</el-card>
			</el-col>

			<el-col :span="12" class="card-box">
				<el-card>
					<div slot="header"><span>CPU</span></div>
					<div class="el-table el-table--enable-row-hover el-table--medium">
						<table cellspacing="0" style="width: 100%;">
							<thead>
								<tr>
									<th class="is-leaf"><div class="cell">属性</div></th>
									<th class="is-leaf"><div class="cell">值</div></th>
								</tr>
							</thead>
							<tbody>
								<tr>
									<td><div class="cell">名字</div></td>
									<td>
										<div v-if="server.CpuName" class="cell">{{ server.CpuName }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">数量</div></td>
									<td>
										<div v-if="server.CpuCores" class="cell">{{ server.CpuCores }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">使用率</div></td>
									<td>
										<div v-if="server.CpuUsage" class="cell">{{ server.CpuUsage }}%</div>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				</el-card>
			</el-col>
		</el-row>
	</div>
</template>

<script>
import { getNServer } from '../../api/monitor/server';

export default {
	name: 'Server',
	data() {
		return {
			// 加载层信息
			loading: [],
			// 服务器信息
			server: []
		};
	},
	created() {
		this.getList();
		this.openLoading();
	},
	methods: {
		getList() {
			getNServer().then(response => {
				// 后期这里应该做个异常保护，读取不到如何处理
				this.server = response.data.old_data;
				this.loading.close();
			});
		},
		// 打开加载层
		openLoading() {
			this.loading = this.$loading({
				lock: true,
				text: '拼命读取中',
				spinner: 'el-icon-loading',
				background: 'rgba(0, 0, 0, 0.7)'
			});
		}
	}
};
</script>

<style>
</style>
