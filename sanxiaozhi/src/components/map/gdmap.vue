<template>
	<div id="container" style="height: 800px;"></div>
</template>

<script>
export default {
	data() {
		return {
			gAmap: ''
		};
	},
	methods: {
		amapInit() {
			// 避免全局引用
			delete window.amapInit;
			this.gAmap = new window.AMap.Map('container', {
				resizeEnable: true,
				center: [116.397428, 39.90923], //地图中心点
				zoom: 13 //地图显示的缩放级别
			});
			this.addMarker();
			this.addDragRoute();
		},
		// 路线规划
		addRoutePlan() {
			var driving = new window.AMap.Driving({
				map: this.gAmap,
				container: 'container'
			});
			// 根据起终点经纬度规划驾车导航路线
			driving.search(new window.AMap.LngLat(116.379028, 39.865042), new window.AMap.LngLat(116.427281, 39.903719), function(status, result) {
				if (status === 'complete') {
					alert('绘制驾车路线完成');
				} else {
					alert('获取驾车数据失败：' + result);
				}
			});
		},
		// 允许拖拽的路线规划
		addDragRoute() {
			var path = [];
			path.push([116.303843, 39.983412]);
			path.push([116.321354, 39.896436]);
			path.push([116.407012, 39.992093]);

			var route = new window.AMap.DragRoute(this.gAmap, path, window.AMap.DrivingPolicy.LEAST_FEE);
			// 查询导航路径并开启拖拽导航
			route.search();
		},
		// 添加一个坐标标识点
		addMarker() {
			var marker = new window.AMap.Marker({
				position: new window.AMap.LngLat(116.39, 39.9), // 经纬度对象，也可以是经纬度构成的一维数组[116.39, 39.9]
				title: '北京'
			});
			this.gAmap.add(marker);
			// this.gAmap.remove(marker);
		},
		loadAmap() {
			var url = 'https://webapi.amap.com/maps?v=1.4.15&key=4b5f2cf2cba25200cc6b68c398468899&callback=amapInit&plugin=AMap.DragRoute';
			var jsapi = document.createElement('script');
			jsapi.charset = 'utf-8';
			jsapi.src = url;
			document.head.appendChild(jsapi);
		}
	},
	mounted() {
		window.amapInit = this.amapInit;
		this.loadAmap();
	}
};
</script>

<style>
.style {
	position: absolute;
	z-index: 99;
	height: 123px;
	width: 821px;
	top: 0px;
	left: 0px;
}
</style>
