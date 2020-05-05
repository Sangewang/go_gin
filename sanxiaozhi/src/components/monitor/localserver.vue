<template>
	<div class="app-container">
		<el-row>
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
									<td><div class="cell">服务器架构</div></td>
									<td>
										<div v-if="server.Arch" class="cell">{{ server.Arch }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">操作系统</div></td>
									<td>
										<div v-if="server.Operate" class="cell">{{ server.Operate }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">核心数</div></td>
									<td>
										<div v-if="server.Cpunum" class="cell">{{ server.Cpunum }}</div>
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
					<div slot="header"><span>go运行环境</span></div>
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
									<td><div class="cell">go version</div></td>
									<td>
										<div v-if="server.GoVersion" class="cell">{{ server.GoVersion }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">Goroutine</div></td>
									<td>
										<div v-if="server.NumGoroutine" class="cell">{{ server.NumGoroutine }}</div>
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
									<td><div class="cell">总共</div></td>
									<td>
										<div class="cell">{{ server.DiskTotal }}</div>
									</td>
								</tr>
								<tr>
									<td><div class="cell">可用</div></td>
									<td>
										<div class="cell">{{ server.DiskFree }}</div>
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
			server: [],
			// 请求地址
			url: 'http://127.0.0.1:9090/monitor/local'
		};
	},
	created() {
		this.getList();
		this.openLoading();
	},
	methods: {
		/* 查询服务器信息
		getList() {
			getServer().then(response => {
				this.server = response;
				this.loading.close();
			});
		}, */
		getList() {
			getNServer(this.url, "").then(response => {
				console.log(response);
				this.server = response.old_data;
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
