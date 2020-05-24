import {
	loginByEmail,
	logout,
	getInfo
} from '../../api/login.js';
import Cookies from 'js-cookie';

const user = {
	// store的结构可以参考 https://vuex.vuejs.org/zh/guide/modules.html
	/*
	Vuex 使用单一状态树——是的，用一个对象就包含了全部的应用层级状态。
	至此它便作为一个“唯一数据源 (SSOT)”而存在。这也意味着,每个应用将仅仅包含一个 store 实例。
	单一状态树让我们能够直接地定位任一特定的状态片段，在调试的过程中也能轻易地取得整个当前应用状态的快照。
	*/
	state: {
		user: '',
		status: '',
		email: '',
		code: '',
		uid: undefined,
		auth_type: '',
		token: Cookies.get('Admin-Token'),
		// token: "Cookies.get('Admin-Token')",
		name: '',
		avatar: '',
		introduction: '',
		roles: [],
		setting: {
			articlePlatform: []
		}
	},

	/* 
	更改 Vuex 的 store 中的状态的唯一方法是提交 mutation。
	Vuex 中的 mutation 非常类似于事件：每个 mutation 都有一个字符串的 事件类型 (type) 和 一个 回调函数 (handler)。
	这个回调函数就是我们实际进行状态更改的地方，并且它会接受 state 作为第一个参数
	*/
	mutations: {
		SET_AUTH_TYPE: (state, type) => {
			state.auth_type = type;
		},
		SET_CODE: (state, code) => {
			state.code = code;
		},
		SET_TOKEN: (state, token) => {
			state.token = token;
		},
		SET_UID: (state, uid) => {
			state.uid = uid;
		},
		SET_EMAIL: (state, email) => {
			state.email = email;
		},
		SET_INTRODUCTION: (state, introduction) => {
			state.introduction = introduction;
		},
		SET_SETTING: (state, setting) => {
			state.setting = setting;
		},
		SET_STATUS: (state, status) => {
			state.status = status;
		},
		SET_NAME: (state, name) => {
			state.name = name;
		},
		SET_AVATAR: (state, avatar) => {
			state.avatar = avatar;
		},
		SET_ROLES: (state, roles) => {
			state.roles = roles;
		},
		LOGIN_SUCCESS: () => {
			console.log('login success')
		},
		LOGOUT_USER: state => {
			state.user = '';
		}
	},

	/*
	Action 类似于 mutation，不同在于：
	Action 提交的是 mutation，而不是直接变更状态。
	Action 可以包含任意异步操作。
	*/
	actions: {
		// 邮箱登录
		LoginByEmail({
			commit
		}, userInfo) {
			const email = userInfo.email.trim();
			// alert(email);
			return new Promise((resolve, reject) => {
				loginByEmail(email, userInfo.password).then(response => {
					const data = response.data;
					console.log("store response.data = ", response.data);
					let currData = data.data;
					Cookies.set('Admin-Token', currData.token);
					commit('SET_TOKEN', currData.token);
					commit('SET_EMAIL', email);
					resolve();
				}).catch(error => {
					reject(error);
				});
			});
		},

		// 获取用户信息
		GetInfo({
			commit,
			state
		}) {
			return new Promise((resolve, reject) => {
				getInfo(state.token).then(response => {
					const data = response.data;
					console.log("GetInfo data = ", data)
					let retData = data.data
					commit('SET_ROLES', retData.role);
					commit('SET_NAME', retData.name);
					commit('SET_AVATAR', retData.avatar);
					commit('SET_UID', retData.uid);
					commit('SET_INTRODUCTION', retData.introduction);
					resolve(response);
				}).catch(error => {
					reject(error);
				});
			});
		},

		// 登出
		LogOut({
			commit,
			state
		}) {
			return new Promise((resolve, reject) => {
				logout(state.token).then(() => {
					commit('SET_TOKEN', '');
					commit('SET_ROLES', []);
					Cookies.remove('Admin-Token');
					resolve();
				}).catch(error => {
					reject(error);
				});
			});
		},
	}
};

export default user;
