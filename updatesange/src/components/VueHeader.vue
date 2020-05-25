<template>
	<navbar>
		<button class="navbar-toggler mobile-sidebar-toggler d-lg-none" type="button" @click="mobileSidebarToggle">&#9776;</button>
		<a class="navbar-brand"></a>
		<ul class="nav navbar-nav d-md-down-none">
			<li class="nav-item"><a class="nav-link navbar-toggler sidebar-toggler" @click="sidebarMinimize">&#9776;</a></li>
		</ul>

		<ul class="nav navbar-nav d-md-down-none">
			<li class="nav-item header-item">
				<router-link tag="div" to="/" style="height:55px;" class="nav-link">
					<p><Icon type="md-analytics" size="30" color="#2d8cf0"></Icon></p>
					<p style="color:white">系统定位</p>
				</router-link>
			</li>
			<li class="nav-item header-item">
				<router-link tag="div" to="/map" style="height:55px;" class="nav-link">
					<p><Icon type="md-map" size="30" color="#2d8cf0"></Icon></p>
					<p style="color:white">地图平台</p>
				</router-link>
			</li>

			<li class="nav-item header-item">
				<Dropdown>
					<router-link tag="div" to="/" style="height:55px;" class="nav-link">
						<p><Icon type="md-hand" size="30" color="#2d8cf0"></Icon></p>
						<p style="color:white">我的工作台</p>
					</router-link>
					<DropdownMenu slot="list">
						<DropdownItem><a href="https://github.com/Sangewang/go_gin" target="_blank">
						<p style="color: #000000"><Icon type="logo-github"></Icon>github</p></a>
						</DropdownItem>
						<DropdownItem><a href="https://blog.csdn.net/u011559236" target="_blank"></a>
						<p style="color: #000000">csdn</p>
						</DropdownItem>
						<DropdownItem>Vue Js</DropdownItem>
					</DropdownMenu>
				</Dropdown>
			</li>

			<li class="nav-item header-item">
				<router-link tag="div" to="/home1" style="height:55px;" class="nav-link">
					<p><Icon type="md-trophy" size="30" color="#2d8cf0"></Icon></p>
					<p style="color:white">课外知识</p>
				</router-link>
			</li>
		</ul>
		<ul class="nav navbar-nav ml-auto">
			<li class="nav-item d-md-down-none">
				<a class="nav-link">
					<Icon type="android-notifications" size="20"></Icon>
					<span class="badge badge-pill badge-danger">5</span>
				</a>
			</li>

			<Dropdown class="nav-item">
				<a href="javascript:void(0)">
					<span slot="button">
						<img src="static/img/avatars/6.jpg" class="img-avatar" alt="o" />
						<span class="d-md-down-none">admin</span>
					</span>
				</a>
				<Dropdown-menu slot="list">
					<Dropdown-item>
						<p class="dropdown-itemp">
							<Icon type="md-alert"></Icon>
							Updates
							<span class="badge badge-info">42</span>
						</p>
					</Dropdown-item>
					<Dropdown-item>
						<p class="dropdown-itemp">
							<Icon type="md-chatboxes"></Icon>
							Messages
							<span class="badge badge-success">42</span>
						</p>
					</Dropdown-item>
					<Dropdown-item>
						<p class="dropdown-itemp">
							<Icon type="md-chatboxes"></Icon>
							Messages
							<span class="badge badge-danger">42</span>
						</p>
					</Dropdown-item>
					<Dropdown-item divided>
						<p class="dropdown-itemp">
							<Icon type="md-contact"></Icon>
							Profile
						</p>
					</Dropdown-item>
					<Dropdown-item>
						<p class="dropdown-itemp">
							<Icon type="md-settings"></Icon>
							Settings
						</p>
					</Dropdown-item>

					<Dropdown-item>
						<a href="" @click="Logout">
							<p class="dropdown-itemp">
								<Icon type="md-log-out"></Icon>
								Logout
							</p>
						</a>
					</Dropdown-item>
				</Dropdown-menu>
			</Dropdown>
			<li class="nav-item d-md-down-none"><a class="nav-link navbar-toggler aside-menu-toggler" @click="asideToggle">&#9776;</a></li>
		</ul>
	</navbar>
</template>

<script>
import navbar from './VueNavbar';

export default {
	name: 'vueheader',
	components: {
		navbar
	},
	methods: {
		Logout(e) {
			e.preventDefault();
			this.$store
				.dispatch('LogOut')
				.then(() => {
					this.$router.push({ path: '/login' });
				})
				.catch(err => {
					this.$message.error(err);
				});
		},
		click() {
			// do nothing
		},
		sidebarToggle(e) {
			e.preventDefault();
			document.body.classList.toggle('sidebar-hidden');
		},
		sidebarMinimize(e) {
			e.preventDefault();
			document.body.classList.toggle('sidebar-minimized');
		},
		mobileSidebarToggle(e) {
			e.preventDefault();
			document.body.classList.toggle('sidebar-mobile-show');
		},
		asideToggle(e) {
			e.preventDefault();
			document.body.classList.toggle('aside-menu-hidden');
		}
	}
};
</script>

<style type="text/css" scoped>
.dropdown-itemp {
	text-align: left;
	font-size: 15px;
	padding: 10px;
}
.header-item .ivu-dropdown-item {
	padding: 15px;
}
.header-item {
	width: 130px;
	/*background-color: #20a8d8;*/
	/*background-color: white;*/

	height: 55px;
}
.header-item a {
	color: white !important;
}
</style>
