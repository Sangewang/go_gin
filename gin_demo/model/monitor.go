package model

type HostSys struct {
	Hostname    string `json:"hostname"`	// 服务器名称
    Uptime      uint64 `json:"uptime"`		// 系统运行时间
	BootTime    string `json:"bootTime"`	// 系统启动事件
	KernelArch  string `json:"kernelArch"`	// 服务器架构
}

type CpuSys struct {
	CpuName string `json: "modelname"` // cpu名称
	CpuCores int32 `json: "cpucores"` // 内核数量
	CpuUsage float64 `json: "cpuusage"` // cpu使用率
}

type MemSys struct {
	MemTotal uint64 `json: "total"`  // 总内存
	MemUsed uint64 `json: "used"`   // 已使用内存
	MemUsage float64 `json: "memusage"`
}

type GoSys struct {
	GoVersion string `json "version"` // go的版本号
	NumGoroutine int `json "numgoroutine"` // Goroutine
}

type DiskSys struct {
	DiskPath string `json: "diskpath"`
	DiskTotal uint64 `json: "total"`
	DiskUsed uint64 `json "used"`
	DiskUsage float64 `json "diskusage"`
}

type MonitorSys struct {
	Name string `json: "name"`
	HostSys
	CpuSys
	MemSys
	GoSys
	DiskSys
 }