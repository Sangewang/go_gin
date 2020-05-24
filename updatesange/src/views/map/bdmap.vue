<template>
	<div><ve-bmap :settings="chartSettings" :after-set-option-once="afterSet" :series="chartSeries" :tooltip="chartTooltip"></ve-bmap></div>
</template>
<script>
export default {
	data() {
		this.chartSettings = {
			key: 'Syuf2OnG5RaSOO6Bsx5LaY5vPhpfdFZj',
			bmap: {
				center: [120, 30],
				zoom: 14,
				roam: true,
				mapStyle: {}
			}
		};
		this.chartTooltip = { show: true };
		return {
			chartSeries: [
				{
					type: 'scatter',
					coordinateSystem: 'bmap',
					data: [
						[120, 30, 1] // 经度，维度，value，...
					]
				}
			],
			gBmap: ''
		};
	},
	methods: {
		afterSet: function(echarts) {
			var bmap = echarts
				.getModel()
				.getComponent('bmap')
				.getBMap();
			bmap.addControl(new window.BMap.MapTypeControl());
			console.log(window.BMap);
			this.gBmap = bmap;
			this.addDiffControl();
			this.addRoutePlan();
		},
		// 添加一些乱七八糟的控件，BMap需要使用window.BMap才生效
		addDiffControl() {
			this.gBmap.addControl(new window.BMap.NavigationControl()); // 添加平移缩放控件
			this.gBmap.addControl(new window.BMap.ScaleControl()); // 添加比例尺控件
			this.gBmap.enableContinuousZoom(); //启用地图惯性拖拽，默认禁用
			this.gBmap.addControl(new window.BMap.OverviewMapControl()); //添加缩略地图控件
			this.gBmap.enableScrollWheelZoom(); //启用滚轮放大缩小
			this.gBmap.addControl(new window.BMap.MapTypeControl()); //添加地图类型控件
			this.gBmap.disable3DBuilding();
		},
		// 路线规划
		addRoutePlan() {
			this.gBmap.centerAndZoom(new window.BMap.Point(116.404, 39.915), 12);

			var driving = new window.BMap.DrivingRoute(this.gBmap, {
				renderOptions: {
					map: this.gBmap,
					autoViewport: true,
					enableDragging: true //起终点可进行拖拽
				}
			});
			var start = new window.BMap.Point(116.310791, 40.003419);
			var end = new window.BMap.Point(116.486419, 39.877282);
			driving.search(start, end);
		}
	}
};
</script>
