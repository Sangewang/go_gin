# gin_demo(后端框架) & sanxiaozhi（前端框架）

## 🎁 简介
### => [gin_demo] 是一个用gin框架搭建的后端项目
### => [sanxiaozhi] 是一个使用element_ui + vue搭建的前端框架
### => 前后端分离的方式使得整个系统更加灵活，可移植性、可扩展性更好


## ✨ 功能
```
- 登录

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
![登录界面](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/login.png)

#### 👉 后端运行截图
![后端运行截图](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/go_gin.jpg)

#### 👉 VUE前端
![前端](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/vue-ui.png)

#### 👉 主要界面展示1
![主要界面展示1](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/show1.png)

#### 👉 主要界面展示2
![主要界面展示2](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/show2.png)

## gin_demo 目录结构
```shell
├── template					// 页面模板
│   ├── fileUpload.html			// 文件上传的页面demo
│   ├── vueDemo.html			// vue的页面demo
│   └── insertUser.html			// 插入数据的页面demo
├── controller					// 控制器 具体业务逻辑处理
│   ├── FileController.go		// 文件处理控制器
│   ├── VueShowDeom.go			// vue Demo 无用
│   ├── UserController.go		// 用户信息注册、查询的控制器
│   └── SystemController.go		// 系统信息处理控制器
├── go.mod						// go mod 模块
├── go.sum						// go的安装信息
├── main.go						// 主函数 路由信息规划
├── model						// 数据模块
│   ├── commonResponse.go		// 返回信息结构体
│   ├── user.go					// 用户信息结构体
│   └── sys.go					// 系统信息结构体
├── config						// 系统配置
│   └── db.conf					// 数据库配置创建、插入
├── database					// 数据库
    └── connectDB.go			// 数据库连接
```
## 
## License
Copyright (c) 2020 linwang



