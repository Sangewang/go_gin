<template>
	<div class="sidebar">
		<nav class="sidebar-nav">
			<ul class="nav">
				<template v-for="item in routes">
					<!--判断一级目录下是否有二级目录-->
					<router-link
						tag="li"
						class="nav-item nav-dropdown"
						v-if="!item.hidden && item.children && item.children.length > 0"
						:key="item.name"
						:to="item.path + '' + item.children[0].path"
						disabled
						>
						<!--二级目录展示-->
						<div class="nav-link nav-dropdown-toggle" @click="handleClick">
							<Icon :type="item.icon" color="white" />
							{{ item.name }}
						</div>
						<!--二级目录进行展开-->
						<ul class="nav-dropdown-items">
							<li class="nav-item" v-for="child in item.children" :key="child.name" @click="addActive">
								<!--判断二级目录下是否有三级目录-->
								<router-link
									tag="li"
									class="nav-item nav-dropdown"
									v-if="!child.hidden && child.children && item.children.length > 0"
									:key="child.name"
									:to="child.path + '' + child.children[0].path"
									disabled
									>
									<!-- 三级目录展示 -->
									<div class="nav-link nav-dropdown-toggle" @click="handleClick">
										<Icon :type="child.icon" color="white" />
										{{ child.name }}
									</div>
									
									<ul class="nav-dropdown-items">
										<li class="nav-item" v-for="son in child.children" :key="son.name" @click="addActive">
											<router-link
												:to="item.path + '/' + child.path + '/' + son.path"
												class="nav-link"
											>
												<Icon :type="son.icon" color="white" />
													{{ son.name }}
											</router-link>
										</li>
									</ul>
								</router-link>
								<!-- 二级目录下没有子路由-->
								<router-link
									v-else
									:to="item.path + '/' + child.path"
									class="nav-link"
								>
									<Icon :type="child.icon" color="white" />
										{{ child.name }}
								</router-link>
							</li>
						</ul>
					</router-link>
					<!--只有一级目录-->
					<li class="nav-item" v-if="!item.hidden && !item.children" :key="item.name">
						<router-link :to="item.path" class="nav-link" exact>
							<Icon :type="item.icon" color="white" />
							{{ item.name }}
						</router-link>
					</li>
				</template>
			</ul>
		</nav>
	</div>
</template>

<script>
export default {
	name: 'SidebarItem',
	props: {
		routes: {
			type: Array
		}
	},
	methods: {
		handleClick(e) {
			e.preventDefault();
			e.target.parentElement.classList.toggle('open');
		},
		addActive(e) {
			e.preventDefault();
			e.target.parentElement.parentElement.parentElement.classList.add('open');
		}
	},
	mounted() {
		console.log("this.routes = ", this.routes);
	}
};
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.wscn-icon {
	margin-right: 10px;
}
.hideSidebar .menu-indent {
	display: block;
	text-indent: 10px;
}

</style>
