<template>
	<div id="app">
		<header>
			<router-link class="green" to="/home">Home</router-link>
			|
			<router-link class="blue" to="/about">About</router-link>
			|
			<router-link class="red" to="/login">login</router-link>
			<!--
			<router-link class="red" to="/lw/123">lw123</router-link>
			<br>
			<router-link class="red" to="/lw/456">lw456</router-link>
			-->
		</header>
		<router-view></router-view>   
	</div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
// import lw from './components/lw.vue'
// import getGinData from './components/getGinData.vue'
// import login from './components/login.vue'

export default {
  name: 'App',
  components: {
    // HelloWorld,
	// lw,
	// getGinData,
	// login,
  }
}
</script>

<style>
	#app {
		font-family: Avenir, Helvetica, Arial, sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: center;
		color: #2c3e50;
		margin-top: 0px;
	}
	body{
		margin: 0px;
		padding: 0;
	}
</style>
