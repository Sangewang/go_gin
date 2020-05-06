<template>
	<body id="poster">
		<el-form :model="loginForm" ref="loginForm" :label-position="labelPosition" label-width="80px" class="login-container" status-icon :rules="loginRules">
			<h3 class="login_title">系统定位分析平台</h3>
			<el-form-item label="用户名" prop="username"><el-input v-model="loginForm.username" autocomplete="off"></el-input></el-form-item>
			<el-form-item label="密码" prop="password"><el-input type="password" v-model="loginForm.password" autocomplete="off"></el-input></el-form-item>
			<el-button type="primary" style="width: 100%;background: #505458;border: none" @click="submitForm('loginForm')">
				<span v-if="!loading">登 录</span>
				<span v-else>登 录 中...</span>
			</el-button>
		</el-form>
	</body>
</template>

<script>
import axios from 'axios';
import Vue from 'vue';
import Global from '../../util/global/glabalval'
export default {
	name: 'login',
	data() {
		var usernamePass = (rule, value, callback) => {
			if (value === '') {
				callback(new Error('请输入用户名, 提示: admin'));
			} else {
				callback(); // 会掉必须有
			}
		};
		return {
			loginForm: {
				username: 'admin',
				password: '123456'
			},
			responseResult: [],
			loading: false,
			labelPosition: 'left',
			loginRules: {
				// username: [{ required: true, trigger: 'blur', message: '用户名不能为空' }],
				username: [{ validator: usernamePass, trigger: 'blur' }],
				password: [{ required: true, trigger: 'blur', message: '密码不能为空' }]
			}
		};
	},
	methods: {
		submitForm(formName) {
			// form 那边有了ref后这边才可以使用this.$refs
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.handleLogin();
				} else {
					console.log('error submit!!');
					return false;
				}
			});
		},
		resetForm(formName) {
			this.$refs[formName].resetFields();
		},
		handleLogin() {
			axios
				.post('http://127.0.0.1:9090/user/login', {
					username: this.loginForm.username,
					password: this.loginForm.password
				})
				.then(successResponse => {
					console.log(successResponse);
					if (successResponse.data.code === 0) {
						// alert(successResponse.data.msg)
						// 将全局变量模块挂载到Vue.prototype中、Global中
						Global.username = this.loginForm.username;
						Vue.prototype.$userInfo = this.loginForm;
						// 登录成功后初始化数据
						Vue.prototype.$dataInfo = successResponse.data.data;
						// 登陆成功后进行页面跳转
						this.$router.replace({ path: '/monitor' });
					} else {
						alert(successResponse.data.msg);
					}
				})
				.catch(failResponse => {
					alert(failResponse);
				});
		}
	}
};
</script>

<style>
#poster {
	background: url('../../assets/login.png') no-repeat;
	background-position: center;
	height: 100%;
	width: 100%;
	background-size: cover;
	position: fixed;
}
body {
	margin: 0px;
	padding: 0;
}

.login-container {
	border-radius: 15px;
	background-clip: padding-box;
	margin: 90px auto;
	width: 350px;
	padding: 35px 35px 15px 35px;
	background: #fff;
	border: 1px solid #eaeaea;
	box-shadow: 0 0 25px #cac6c6;
}

.login_title {
	margin: 0px auto 40px auto;
	text-align: center;
	color: #505458;
}

/*
.el-input__inner {
	width: 40%;
}*/
</style>
