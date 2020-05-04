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
			]
		};
	},
	methods: {
		afterSet: function(echarts) {
			var bmap = echarts
				.getModel()
				.getComponent('bmap')
				.getBMap();
			bmap.addControl(new window.BMap.MapTypeControl());
		}
	}
};
</script>
