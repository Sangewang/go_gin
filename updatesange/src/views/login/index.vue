<template>
	<div class="login-container" style="background-color: #141a48;margin: 0px;overflow: hidden; height: 100%;">
		<div id="canvascontainer" ref='can'></div>
		<Form ref="loginForm" autoComplete="on" :model="loginForm" :rules="loginRules" class="card-box login-form">
			<Form-item prop="email">
				<Input type="text" v-model="loginForm.email" placeholder="Username" autoComplete="on">
					<Icon type="ios-person-outline" slot="prepend"></Icon>
				</Input>
			</Form-item>
			<Form-item prop="password">
				<Input type="password" v-model="loginForm.password" placeholder="Password" @keyup.enter.native="handleLogin">
					<Icon type="md-lock" slot="prepend"></Icon>
				</Input>
			</Form-item>
			<Form-item>
				<Button type="primary" @click="handleLogin('loginForm')" long>登录</Button>
				<div class='tips'>admin账号为:admin@bytedance.com 密码123456</div>
			</Form-item>
		</Form>
	</div>
</template>

<script>
	import { isByteDanceEmail } from '../../utils/validate.js'
	// import Stats from './jsm/libs/stats.module.js';
	import * as THREE from 'three'
	var SEPARATION = 100, AMOUNTX = 50, AMOUNTY = 50;
	// var stats;
	var container;
	var camera, scene, renderer;
	var particles, count = 0;
	var mouseX = 0, mouseY = 0;
	var windowHalfX = window.innerWidth / 2;
	var windowHalfY = window.innerHeight / 2;
	// import * as THREE from '../../../static/js/three.min.js';
	
	export default {
		name: 'login',
		data() {
			const validateEmail = (rule, value, callback) => {
				if(!isByteDanceEmail(value)) {
					callback(new Error("请输入正确且合法的公司邮箱"))
				} else {
					callback();
				}
			};
			const validatePass = (rule, value, callback) => {
				if (value.length < 6) {
					callback(new Error("密码不能小于6位"));
				} else {
					callback();
				}
			};
			return {
				loginForm: {
					email: 'admin@bytedance.com',
					password: ''
				},
				loginRules: {
					email: [
						{ required: true, trigger: 'blur', validator: validateEmail }
					],
					password: [
						{ required: true, trigger: 'blur', validator: validatePass }
					]
				},
				loading: false,
				showDialog: false
			}
		},
		// 挂在成功后执行的函数
		mounted() {
			// alert("mounted");
			this.init();
			animate();
		},
		methods: {
			init,
			handleLogin() {
				// alert("登陆");
				this.$refs.loginForm.validate(valid => {
					if (valid) {
						this.loading = true;
						this.$store.dispatch('LoginByEmail', this.loginForm).then(() => {
						this.$Message.success('登录成功');
						this.loading = false;
						this.$router.push({ path: '/' });
						}).catch(err => {
							this.$message.error(err);
							this.loading = false;
						});
					} else {
						console.log('error submit!!');
						return false;
					}
				});
			},
		},
	}
	
	function init() {
		container = document.createElement( 'div' );
		// document.body.appendChild( container ); // => 这个是整个生命周期都有效
		this.$refs.can.appendChild( container );  
		camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 1, 10000 );
		camera.position.z = 1000;
		scene = new THREE.Scene();
		var numParticles = AMOUNTX * AMOUNTY;
		var positions = new Float32Array( numParticles * 3 );
		var scales = new Float32Array( numParticles );
		var i = 0, j = 0;
		for ( var ix = 0; ix < AMOUNTX; ix ++ ) {
			for ( var iy = 0; iy < AMOUNTY; iy ++ ) {
				positions[ i ] = ix * SEPARATION - ( ( AMOUNTX * SEPARATION ) / 2 ); // x
				positions[ i + 1 ] = 0; // y
				positions[ i + 2 ] = iy * SEPARATION - ( ( AMOUNTY * SEPARATION ) / 2 ); // z
				scales[ j ] = 1;
				i += 3;
				j ++;
			}
		}

		var geometry = new THREE.BufferGeometry();
		geometry.setAttribute( 'position', new THREE.BufferAttribute( positions, 3 ) );
		geometry.setAttribute( 'scale', new THREE.BufferAttribute( scales, 1 ) );

		var material = new THREE.ShaderMaterial( {
			uniforms: {
				color: { value: new THREE.Color( 0x0078de ) },
			},
			vertexShader: document.getElementById( 'vertexshader' ).textContent,
			fragmentShader: document.getElementById( 'fragmentshader' ).textContent
		} );

		particles = new THREE.Points( geometry, material );
		scene.add( particles );
		renderer = new THREE.WebGLRenderer( { antialias: true } );
		renderer.setPixelRatio( window.devicePixelRatio );
		renderer.setSize( window.innerWidth, window.innerHeight );
		container.appendChild( renderer.domElement );

		// stats = new Stats();
		// container.appendChild( stats.dom );

		document.addEventListener( 'mousemove', onDocumentMouseMove, false );
		document.addEventListener( 'touchstart', onDocumentTouchStart, false );
		document.addEventListener( 'touchmove', onDocumentTouchMove, false );
		window.addEventListener( 'resize', onWindowResize, false );
	}

	function onWindowResize() {
		windowHalfX = window.innerWidth / 2;
		windowHalfY = window.innerHeight / 2;
		camera.aspect = window.innerWidth / window.innerHeight;
		camera.updateProjectionMatrix();
		renderer.setSize( window.innerWidth, window.innerHeight );
	}

	function onDocumentMouseMove( event ) {
		mouseX = event.clientX - windowHalfX;
		mouseY = event.clientY - windowHalfY;
	}

	function onDocumentTouchStart( event ) {
		if ( event.touches.length === 1 ) {
			event.preventDefault();
			mouseX = event.touches[ 0 ].pageX - windowHalfX;
			mouseY = event.touches[ 0 ].pageY - windowHalfY;
		}
	}

	function onDocumentTouchMove( event ) {
		if ( event.touches.length === 1 ) {
			event.preventDefault();
			mouseX = event.touches[ 0 ].pageX - windowHalfX;
			mouseY = event.touches[ 0 ].pageY - windowHalfY;
		}
	}

	function animate() {
		requestAnimationFrame( animate );
		render();
		// stats.update();
	}

	function render() {
		camera.position.x += ( mouseX - camera.position.x ) * .05;
		camera.position.y += ( - mouseY - camera.position.y ) * .05;
		camera.lookAt( scene.position );
		var positions = particles.geometry.attributes.position.array;
		var scales = particles.geometry.attributes.scale.array;
		var i = 0, j = 0;
		for ( var ix = 0; ix < AMOUNTX; ix ++ ) {
			for ( var iy = 0; iy < AMOUNTY; iy ++ ) {
				positions[ i + 1 ] = ( Math.sin( ( ix + count ) * 0.3 ) * 50 ) +
								( Math.sin( ( iy + count ) * 0.5 ) * 50 );
				scales[ j ] = ( Math.sin( ( ix + count ) * 0.3 ) + 1 ) * 8 +
								( Math.sin( ( iy + count ) * 0.5 ) + 1 ) * 8;
				i += 3;
				j ++;
			}
		}
		particles.geometry.attributes.position.needsUpdate = true;
		particles.geometry.attributes.scale.needsUpdate = true;
		renderer.render( scene, camera );
		count += 0.1;
	}
</script>


<style>
	.login-container a{color:#0078de;}
	#canvascontainer{
		position: absolute;
		top: 0px;
	}
	
	.wz-input-group-prepend{
		background-color: #141a48;
		border: 1px solid #2d8cf0;
		border-right: none;
		color:  #2d8cf0;
	}
</style>


<style rel="stylesheet/scss" lang="scss">
    .tips{
      font-size: 14px;
      color: #fff;
      margin-bottom: 5px;
    } 
    .login-container {
		background-color: #2d3a4b
		input:-webkit-autofill {
			-webkit-box-shadow: 0 0 0px 1000px #293444 inset !important;
			-webkit-text-fill-color: #fff !important;
		}
		input {
				background: transparent;
				border: 1px solid #2d8cf0;
				-webkit-appearance: none;
				border-radius: 3px;
				padding: 12px 5px 12px 15px;
				color: #eeeeee;
				height: 47px;
			}
			.el-input {
				display: inline-block;
				height: 47px;
				width: 85%;
			}
			.svg-container {
				padding: 6px 5px 6px 15px;
				color: #889aa4;
			}
	
			.title {
				font-size: 26px;
				font-weight: 400;
				color: #eeeeee;
				margin: 0px auto 40px auto;
				text-align: center;
				font-weight: bold;
			}
	
			.login-form {
				position: absolute;
				left: 0;
				right: 0;
				width: 400px;
				padding: 35px 35px 15px 35px;
				margin: 120px auto;
			}
	
			.el-form-item {
				border: 1px solid rgba(255, 255, 255, 0.1);
				background: rgba(0, 0, 0, 0.1);
				border-radius: 5px;
				color: #454545;
			}
	
			.forget-pwd {
				color: #fff;
			}
    }
</style>