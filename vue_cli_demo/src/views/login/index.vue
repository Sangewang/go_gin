<template>
	<div class="login-container" style="background-color: #141a48;margin: 0px;overflow: hidden;">
	<div id="canvascontainer" ref='can'></div>

	<Form ref="loginForm" autoComplete="on" :model="loginForm" :rules="loginRules"  class="card-box login-form">
		<Form-item prop="email">
			<Input type="text" v-model="loginForm.email" placeholder="Username" autoComplete="on">
				<Icon type="ios-person-outline" slot="prepend" ></Icon>
			</Input>
		</Form-item>
		<Form-item prop="password">
			<Input type="password" v-model="loginForm.password" placeholder="Password" @keyup.enter.native="handleLogin">
				<Icon type="ios-locked-outline" slot="prepend"></Icon>
			</Input>
		</Form-item>
		<Form-item>
			<Button type="primary" @click="handleLogin('loginForm')" long>登录</Button>
		</Form-item>
		<div class='tips'>admin账号为:admin@wz.com 密码123456</div>
		<div class='tips'>editor账号:editor@wz.com 密码123456</div>
	</Form>
	</div>
</template>

<script>
</script>

<style>
</style>
