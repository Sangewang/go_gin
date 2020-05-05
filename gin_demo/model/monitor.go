package model

type OsSys struct {
	Arch string `json: "arch"`  // 服务器架构
	Operate string `json: "operate"`  // 操作系统
}

type CpuSys struct {
	Cpunum int `json: "cpunum"` // cpu数量
}

type MemSys struct {
	MemTotal float32 `json: "total"`  // 总内存
	MemUsed float32 `json: "used"`   // 已使用内存
	MemUsage float32 `json: "usage"`
}

type GoSys struct {
	GoVersion string `json "version"` // go的版本号
	NumGoroutine int `json "numgoroutine"` // Goroutine
}

type DiskSys struct {
	DiskTotal int `json: "total"`
	DiskFree int `json "free"`
}

type MonitorSys struct {
	Name string `json: "name"`
	OsSys
	CpuSys
	MemSys
	GoSys
	DiskSys
 }