<template>
  <body id="poster">
    <el-form class="login-container" label-position="left" label-width="0px">
      <h3 class="login_title">系统登录</h3>
      <el-form-item>
        <el-input type="text" v-model="loginForm.username" auto-complete="off" placeholder="账号"></el-input>
      </el-form-item>
 
      <el-form-item>
        <el-input type="password" v-model="loginForm.password" auto-complete="off" placeholder="密码"></el-input>
      </el-form-item>
 
      <el-form-item style="width: 100%">
        <el-button type="primary" style="width: 100%;background: #505458;border: none" v-on:click="login">登录</el-button>
      </el-form-item>
    </el-form>
  </body>
</template>
 
 
<script>
	import axios from 'axios'
	import Vue from 'vue' 
    export default {
        name: "login",
        data() {
            return {
                loginForm: {
                    username: 'admin',
                    password: '123456'
                },
                responseResult: []
            }
        },
        methods: {
            login() {
                axios
                    .post('http://127.0.0.1:9090/user/login', {
                        username: this.loginForm.username,
                        password: this.loginForm.password
                    })
                    .then(successResponse => {
						console.log(successResponse)
                        if (successResponse.data.code === 0) {
							// alert(successResponse.data.msg)
							// 将全局变量模块挂载到Vue.prototype中
							Vue.prototype.$userInfo = this.loginForm; 
							// 登录成功后初始化数据
							Vue.prototype.$dataInfo = successResponse.data.data;
							// 登陆成功后进行页面跳转
                            this.$router.replace({path: '/page1'})
                        } else {
							alert(successResponse.data.msg)
						}
                    })
                    .catch(failResponse => {
						alert(failResponse);
                    })
            }
        }
    }
</script>
 
<style>
  #poster {
    background:url("../assets/login.png") no-repeat;
    background-position: center;
    height: 100%;
    width: 100%;
    background-size: cover;
    position: fixed;
  }
  body{
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
</style>