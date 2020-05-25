<template>
	<div class="sidebar">
		<nav class="sidebar-nav">
			<ul class="nav">
				<template v-for="item in routes">
					<!--进行子路由的递归展示 {{ item.name }} {{ item.children }} {{ item.hidden }}-->
					
					<router-link
						tag="li"
						class="nav-item nav-dropdown"
						v-if="!item.hidden && item.children && item.children.length > 0"
						:key="item.name"
						:to="item.path + '' + item.children[0].path"
						disabled
						>
						
						<div class="nav-link nav-dropdown-toggle" @click="handleClick">
							<Icon :type="item.icon" color="white" />
							{{ item.name }}
						</div>
						
						<ul class="nav-dropdown-items">
							<li class="nav-item" v-for="child in item.children" :key="child.name" @click="addActive">
								<!--二级子循环嵌套，场景很少-->
								<router-link 
									v-if='!child.hidden && child.children' 
									:to="child.path + '/' + child.children[0].path"
									class="nav-link"
									>
								</router-link>
								<!--打开二级目录-->
								<router-link
									:to="item.path + '/' + child.path"
									class="nav-link"
								>
									<Icon :type="child.icon" color="white" />
										{{ child.name }}
								</router-link>
								
								
							</li>
						</ul>
						
					</router-link>
					<!--没有子路由，只有根结点-->
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
