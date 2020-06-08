package controller

import (
	"fmt"
	"log"
	"io/ioutil"
	"reflect"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"gin_demo/model"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	/*
    
	"github.com/shirou/gopsutil/net"
	*/
)

// 单位换算成G
const currG = 1024 * 1024 * 1024

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// Host情况
func HostConditon() (string, string, string, uint64) {
	HostInfo, _ := host.Info()
	fmt.Println("HostInfo = ", HostInfo)
	// 本机名称
	HostName := HostInfo.Hostname
	// 内核
	kernelArch := HostInfo.KernelArch
	// 系统运行事件
	uptime := HostInfo.Uptime / (60 * 60 * 24)
	// 系统开机时间
	boottime, _ := host.BootTime()
	btime := time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")
	
	fmt.Println("HostName = ", HostName)
	fmt.Println("kernelArch = ", kernelArch)
	fmt.Println("uptime = ", uptime)
	fmt.Println("btime = ", btime)
	return HostName, kernelArch, btime, uptime
}

// 内存mem查询
func MemCondition() (uint64, uint64, float64) {
	curMem, _ := mem.VirtualMemory()
	curMemTotal := curMem.Total / currG
	curMemUsed := curMem.Used / currG
	curUsedPercent := Decimal(curMem.UsedPercent)
	// almost every return value is a struct
    fmt.Printf("Total: %v, Used:%v, UsedPercent:%f%%\n", curMemTotal, curMemUsed, curUsedPercent)
    // convert to JSON. String() is also implemented
	fmt.Println("全部内存情况 = ", curMem)
	return curMemTotal, curMemUsed, curUsedPercent
}

// CPU使用情况
func CpuCondition() (string, int32, float64) {
	curCpu, _ := cpu.Info()
	fmt.Println("CPU情况 = ", curCpu)
	CpuPercentList, _ := cpu.Percent(time.Second, false)
	fmt.Println("CpuPercentList = ", CpuPercentList)
	curCpuPercent := Decimal(CpuPercentList[0])
	fmt.Println("CPU使用率情况 = ", curCpuPercent)
	var ModelName string
	var Cores int32
	if len(curCpu) > 1 {
        for _, sub_cpu := range curCpu {
            ModelName = sub_cpu.ModelName
            Cores = sub_cpu.Cores
            fmt.Printf("CPU : %v , %v cores \n", ModelName, Cores)
        }
    } else {
        sub_cpu := curCpu[0]
        ModelName = sub_cpu.ModelName
        Cores = sub_cpu.Cores
        fmt.Printf("CPU : %v , %v cores \n", ModelName, Cores)
	}
	return ModelName, Cores, curCpuPercent
}

// 硬盘disk情况
func DiskConditon() (string, uint64, uint64, float64) {
	currDisk, _ := disk.Usage("/Users/bytedance")
	diskPath := currDisk.Path
	diskTotal := currDisk.Total / currG
	diskFree := currDisk.Free / currG
	diskUsed := currDisk.Used / currG
	diskPercent := Decimal(currDisk.UsedPercent)
	fmt.Printf("Path: %v  HD: %v GB  Free: %v GB Used: %v GB Usage:%f%%\n", diskPath, diskTotal, diskFree, diskUsed, diskPercent)
	fmt.Println("硬盘情况 = ", currDisk)
	return diskPath, diskTotal, diskUsed, diskPercent
}

func MonitorLocal(context *gin.Context) {
	log.Println(">>>> func MonitorLocal start <<<<")

	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	fmt.Println("reflect(data) = ", reflect.TypeOf(string(data)))

	// 内存mem查询
	curMemTotal, curMemUsed, curUsedPercent := MemCondition()
	// disk查询
	diskPath, diskTotal, diskUsed, diskPercent := DiskConditon()
	// cpu 使用情况
	cpuName, cpuCores, cpuPercent := CpuCondition()
	// Host 情况
	hostName, kernelArch, btime, uptime := HostConditon()

	type TMonitor model.MonitorSys
	var RetMonitor TMonitor

	RetMonitor.HostSys.Hostname = hostName
	RetMonitor.HostSys.KernelArch = kernelArch
	RetMonitor.HostSys.BootTime = btime
	RetMonitor.HostSys.Uptime = uptime

	RetMonitor.CpuSys.CpuName = cpuName
	RetMonitor.CpuSys.CpuCores = cpuCores
	RetMonitor.CpuSys.CpuUsage = cpuPercent

	RetMonitor.MemSys.MemTotal = curMemTotal
	RetMonitor.MemSys.MemUsed = curMemUsed
	RetMonitor.MemSys.MemUsage = curUsedPercent

	RetMonitor.DiskSys.DiskPath = diskPath
	RetMonitor.DiskSys.DiskTotal = diskTotal
	RetMonitor.DiskSys.DiskUsed = diskUsed
	RetMonitor.DiskSys.DiskUsage = diskPercent

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

