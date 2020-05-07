<template>
	<ve-amap :settings="chartSettings" :after-set-option-once="afterSet" :series="chartSeries" :tooltip="chartTooltip"></ve-amap>
</template>

<script>
export default {
	data() {
		this.chartSettings = {
			key: '4b5f2cf2cba25200cc6b68c398468899',
			v: '1.4.3',
			amap: {
				resizeEnable: true,
				center: [116.39, 39.9],
				zoom: 10
			}
		};
		this.chartTooltip = { show: true };
		return {
			chartSeries: [
				{
					type: 'scatter',
					coordinateSystem: 'amap',
					data: [
						[116.39, 39.9, 1] // 经度，维度，value，...
					]
				}
			],
			gAmap: ''
		};
	},
	methods: {
		afterSet: function(echarts) {
			// console.log("A = ", echarts)
			var amap = echarts
				.getModel()
				.getComponent('amap')
				.getAMap();

			this.gAmap = amap;
			this.addMarker();
			// console.log('this.gAmap = ', this.gAmap);
			// this.addRoutePlan();
			// this.routePlan();
			// amap.addControl(new window.AMap.MapTypeControl());
		},
		// 路线规划
		addRoutePlan() {
			window.Amap.plugin('AMap.Driving', function() {
				alert("路线规划")
			});
		},
		addMarker() {
			var marker = new window.AMap.Marker({
				position: new window.AMap.LngLat(116.39, 39.9), // 经纬度对象，也可以是经纬度构成的一维数组[116.39, 39.9]
				title: '北京'
			});
			this.gAmap.add(marker);
			// this.gAmap.remove(marker);
		},
		//  发起路线规划
		routePlan() {
			console.log("window.Amap = ", window.Amap)
		}
	}
};
</script>
