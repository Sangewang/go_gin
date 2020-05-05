# gin_demo(后端框架) & sanxiaozhi（前端框架）

## 🎁 简介
### => [gin_demo] 是一个用gin框架搭建的后端项目
### => [sanxiaozhi] 是一个使用element_ui + vue搭建的前端框架
### => 前后端分离的方式使得整个系统更加灵活，可移植性、可扩展性更好


## ✨ 功能
```
- 登录
- 定位分析平台

	- 基础信息

	- 系统工具
		- 工具模板
			- 模板创建办
			- 模板操作
			
	- 系统日历

	- 系统管理
		- 系统配置
			- 任务创建
			- 图表操作
		- 日志管理
			- 登录日志
			- 操作日志

- 地图分析平台

	- 系统简介

	- 数据上传

	- 任务创建

	- 任务查询

	- 设备状态

	- 评测结果

	- 备用功能
		- 图片
			- 图片列表

		- 地图
			- 百度地图
			- 高德地图
```


## ⚙ 前端开发
```bash

# 克隆项目
git clone https://github.com/Sangewang/go_gin.git

# 环境准备
vue相关准备和学习 => https://cn.vuejs.org/v2/guide/
element-ui相关准备和学习 => https://element.eleme.cn/#/zh-CN/component/installation
Hbuilder加载项目

# 进入项目目录
cd sanxizozhi

# 安装依赖
npm install

# 启动服务（可以在Hbuilder中启动和安装相关包）
npm run dev

```

## 📨 后端开发
```bash
# 克隆项目
git clone https://github.com/Sangewang/go_gin.git

# 环境准备
go的相关安装和入门学习 => https://www.kancloud.cn/kancloud/web-application-with-golang/44105
gin的gorm相关学习指导 => http://gorm.book.jasperxu.com/models.html#md

# 进入项目目录
cd gin_demo

# 依赖安装 & 运行
1、相关的包依赖使用go mod模式
2、具体步骤
	2.1、创建gin_demo文件夹
	2.2、执行go mod init => 在gin_test下生成go.mod、go.sum
	2.3、执行 go build 运行代码会发现 go mod 会自动查找依赖自动下载, 下载的依赖包存在pkg的mod下面
	2.4、生成的二进制文件gin_demo会在gin_demo下面，直接运行即可启动web服务

# 后端数据准备
注意connectDB.go 文件中的数据库五元组要改成自己的配置，并执行conf/下面的db.conf 插入测试数据

```

## 📦 常用命令
```bash
1、批量杀掉进程
ps -ef | grep gin_demo | grep -v grep | awk '{print $2}' | xargs kill -9
2、git 缓存清理
git rm -r --cached filepath
```

## 👍 Preview 效果图

#### 👉 登录界面
![登录界面](https://img-blog.csdnimg.cn/20200504213706535.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE1NTkyMzY=,size_16,color_FFFFFF,t_70)

#### 👉 后端运行截图
![后端运行截图](https://img-blog.csdnimg.cn/20200504214621343.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE1NTkyMzY=,size_16,color_FFFFFF,t_70)

#### 👉 VUE前端 主要包括 定位分析平台、地图分析平台、百度知道、我的工作台(github、csdn等)
![前端](https://img-blog.csdnimg.cn/20200504213736599.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE1NTkyMzY=,size_16,color_FFFFFF,t_70)

#### 👉 地图分析平台
![地图分析平台](https://img-blog.csdnimg.cn/20200504213803362.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE1NTkyMzY=,size_16,color_FFFFFF,t_70)

#### 👉 定位分析平台
![定位分析平台](https://img-blog.csdnimg.cn/20200504213824709.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE1NTkyMzY=,size_16,color_FFFFFF,t_70)

#### 👉 百度地图
![百度地图](https://img-blog.csdnimg.cn/20200504213902764.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE1NTkyMzY=,size_16,color_FFFFFF,t_70)

#### 👉 高德地图
![高德地图](https://img-blog.csdnimg.cn/2020050421391623.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE1NTkyMzY=,size_16,color_FFFFFF,t_70)

#### 👉 定位分析平台 => 系统监控
![系统监控](https://img-blog.csdnimg.cn/20200505124940371.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE1NTkyMzY=,size_16,color_FFFFFF,t_70)

## gin_demo 目录结构
```shell
gin_demo
├── config						// 系统配置
│   └── db.conf					// 数据库配置创建、插入
├── controller					// 控制器 具体业务逻辑处理
│   ├── FileController.go		// 文件处理控制器
│   ├── MonitorLocal.go			// 系统监控控制器
│   ├── PicController.go		// 图片上传、下载控制器
│   ├── SystemController.go		// 系统信息处理控制器 
│   ├── UserController.go		// 用户信息注册、查询的控制器
│   └── VueShowDeom.go			// vue Demo 暂无使用
├── database					// 数据库
│   └── connectDB.go			// 数据库连接
├── go.mod						// go mod 模块
├── go.sum						// go的安装信息
├── main.go						// 主函数 路由信息规划
├── model						// 数据模块
│   ├── commonResponse.go		// 返回信息结构体
│   ├── file.go					// 文件信息结构体
│   ├── monitor.go				// 监控信息结构体
│   ├── pic.go					// 文件信息结构体
│   ├── sys.go					// 系统信息结构体
│   └── user.go					// 用户信息结构体
└── template					// 页面模板
    ├── fileUpload.html			// 文件上传的页面demo
    ├── insertUser.html			// 插入数据的页面demo
    └── vueDemo.html			// vue的页面demo
```
## 

## sanxiaozhi 目录结构
```shell
.
├── README.md
├── babel.config.js
├── package-lock.json
├── package.json				// package.json
├── public
│   ├── favicon.ico
│   └── index.html
└── src							// 源代码
    ├── App.vue					// 入口页面
    ├── api						// 所有请求
    │   ├── http.js				// http请求
    │   ├── message.js			// 信息展示
    │   ├── monitor				// 后端请求服务数据
    │   │   └── server.js
    │   ├── request.js			// axios请求
    │   └── system				// 系统字典
    │       └── dict
    │           └── type.js
    ├── assets
    │   ├── login.png			// 登陆页面图片
    │   └── logo.png			// logo图片
    ├── components
    │   ├── Footer.vue				// 页脚
    │   ├── Header.vue				// 页眉
    │   ├── HelloWorld.vue			// 未使用
    │   ├── Pagination				// 翻页
    │   │   └── index.vue			// 翻页模板 
    │   ├── calendar.vue			// 日历
    │   ├── chart.vue				// 图表
    │   ├── containers				// 自适应布局组合
    │   │   └── Full.vue			// 主界面
    │   ├── home.vue				// 定位分析平台页面
    │   ├── login.vue				// 登陆组件
    │   ├── map						// 地图定位分析组件
    │   │   ├── bdmap.vue			// 百度地图
    │   │   ├── createRoute.vue		// 路线创建
    │   │   ├── gdmap.vue			// 高德地图
    │   │   ├── introduction.vue	// 简介
    │   │   ├── map.vue				// 地图定位分析平台
    │   │   ├── picshow.vue			// 图片展示
    │   │   ├── result.vue			// 结果展示
    │   │   └── upload.vue			// 文件上传
    │   ├── monitor					// 服务器监控
    │   │   └── localserver.vue		// 服务监控页面展示
    │   ├── page1.vue				// 主页面1
    │   ├── page2.vue				// 主页面2
    │   ├── sysdemocreate.vue		// 模板创建页面
    │   ├── sysdemoperate.vue		// 暂时未使用
    │   └── table.vue				// 表格展示
    ├── main.js						// 全局方法定义
    ├── router						// 路由
    │   └── index.js				// 路由添加
    └── util						// 全局公用方法
        ├── form.js					// 表格方法
        ├── request.js				// http请求方法
        └── scroll-to.js			// 页面滑动方法
```
## 


## License
Copyright (c) 2020 linwang



