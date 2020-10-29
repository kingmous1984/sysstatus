package main

import (
	"flag"
	"fmt"
	"sysstatus/status"
)

func main() {
	var allStatus, sysStatus, cpuStatus, memStatus, diskStatus, netStatus, processStatus bool
	var output int
	flag.BoolVar(&sysStatus, "s", false, "系统状态,默认为False")
	flag.BoolVar(&cpuStatus, "c", false, "CPU状态,默认为False")
	flag.BoolVar(&memStatus, "m", false, "内存状态,默认为False")
	flag.BoolVar(&diskStatus, "d", false, "磁盘状态,默认为False")
	flag.BoolVar(&netStatus, "n", false, "网络状态,默认为False")
	flag.BoolVar(&processStatus, "p", false, "进程状态,默认为False")
	flag.BoolVar(&allStatus, "a", false, "全部状态,默认为False")
	flag.IntVar(&output, "o", 1, "输出类型（1：屏幕打印；2：日志文件）")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	// fmt.Printf("output=%d; allStatus=%v; sysStatus=%v; cpuStatus=%v; memStatus=%v; diskStatus=%v; netStatus=%v;\n",
	// 	output, allStatus, sysStatus, cpuStatus, memStatus, diskStatus, netStatus)

	outputStatus, err := status.InitStatusOutput(output)
	if err != nil {
		fmt.Println(err)
	}
	if allStatus || sysStatus {
		outputStatus.SysInfo()
	}
	if allStatus || cpuStatus {
		outputStatus.CPUInfo()
	}
	if allStatus || memStatus {
		outputStatus.MemInfo()
	}
	if allStatus || diskStatus {
		outputStatus.DiskInfo()
	}
	if allStatus || netStatus {
		outputStatus.NetInfo()
	}
	if allStatus || processStatus {
		outputStatus.ProcessInfo()
	}
	fmt.Println("Work is over!")
}
