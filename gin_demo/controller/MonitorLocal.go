package controller

import (
	"fmt"
	"log"
	"io/ioutil"
	"reflect"
	"github.com/gin-gonic/gin"
	"gin_demo/model"
)

func MonitorLocal(context *gin.Context) {
	log.Println(">>>> func MonitorLocal start <<<<")

	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	fmt.Println("reflect(data) = ", reflect.TypeOf(string(data)))

	type TMonitor model.MonitorSys
	var RetMonitor TMonitor

	RetMonitor.Name = "linwangdeMacBook"
	RetMonitor.OsSys.Arch = "amd64"
	RetMonitor.OsSys.Operate = "MACOS"
	RetMonitor.CpuSys.Cpunum = 1
	RetMonitor.MemSys.MemTotal = 16
	RetMonitor.MemSys.MemUsed = 4
	RetMonitor.MemSys.MemUsage = RetMonitor.MemSys.MemUsed / RetMonitor.MemSys.MemTotal * 100
	RetMonitor.GoSys.GoVersion = "go1.12.7"
	RetMonitor.GoSys.NumGoroutine = 20
	RetMonitor.DiskSys.DiskTotal = 100
	RetMonitor.DiskSys.DiskFree = 50

	fmt.Println("RetMonitor = ", RetMonitor)

	flag := true
	if flag == false {
		context.JSON(200, gin.H{
			"code" : -1,
			"msg"  : "验证失败",
			"data" : "",
			"old_data": "",
		})
	} else {
		// 输出序列化后的结果
		// fmt.Printf("序列化后 = %v\n", string(SysData))
		context.JSON(200, gin.H{
			"code" : 0,
			"msg"  : "验证成功",
			"data"  : RetMonitor,
			"old_data": RetMonitor,
		})
	}
}

