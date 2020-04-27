一、Web通用的后端, 采用go_gin框架进行开发
1、相关的包依赖使用go mod模式
2、具体步骤
	2.1、创建gin_demo文件夹
	2.2、执行go mod init => 在gin_test下生成go.mod、go.sum
	2.3、执行 go build 运行代码会发现 go mod 会自动查找依赖自动下载, 下载的依赖包存在pkg的mod下面
	2.4、生成的二进制文件gin_demo会在gin_demo下面，直接运行即可启动web服务
3、注意connectDB.go 文件中的数据库五元组要改成自己的配置，并执行conf/下面的db.conf 插入测试数据

二、sanxiaozhi是采用vue + element_ui的方式开发的前段，目前留了比较通用的接口，只需要增加对应的后端接口即可
VUE的相关知识可以参考：https://cn.vuejs.org/v2/guide/
Element—UI相关知识参考：https://element.eleme.cn/#/zh-CN/component/installation

三、杀掉进程
ps -ef | grep gin_demo | grep -v grep | awk '{print $2}' | xargs kill -9

四、下面是一些运行起来的截图
4.1 后端运行截图
![Image text](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/go_gin.jpg)

4.2 VUE前端
![Image text](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/vue-ui.png)

4.3 主要界面展示1
![Image text](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/show1.png)

4.4 主要界面展示2
![Image text](https://raw.githubusercontent.com/Sangewang/go_gin/master/pic/show2.png)