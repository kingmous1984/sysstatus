package status

import (
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

// CPUInfo CPU信息
func (s *StatusOutput) CPUInfo() {
	s.Logger.Infoln("=============================================CPU=============================================")
	s.Logger.Infof("虚拟内核数：%d", runtime.NumCPU())
	s.Logger.Infof("架构：%s", runtime.GOARCH)
	res, err := cpu.Times(false)
	if err != nil {
		s.Logger.Errorln("CPU ERROR: ", err)
		return
	}
	s.Logger.Infof("Total:%f, User:%f, System:%f, Idle:%f, UsedPercent:%.2f%%", res[0].Total(), res[0].User, res[0].System, res[0].Idle, (res[0].User+res[0].System)*100/res[0].Total())
	// fmt.Printf("使用率：%f", 100-res[0].Idle*100/res[0].Total())
	info, _ := cpu.Info() //总体信息
	//[{"cpu":0,"vendorId":"GenuineIntel","family":"205","model":"","stepping":0,"physicalId":"BFEBFBFF000906EA","coreId":"","cores":6,
	// "modelName":"Intel(R) Core(TM) i5-9500 CPU @ 3.00GHz","mhz":3000,"cacheSize":0,"flags":[],"microcode":""}]
	s.Logger.Infof("物理核心数：%d", len(info))
	for _, c := range info {
		s.Logger.Infof("index:%d, cores:%d, modelName:%s, mhz:%f", c.CPU+1, c.Cores, c.ModelName, c.Mhz)
	}
	cpuLoad, err := load.Avg()
	if err != nil {
		s.Logger.Errorln("LOAD ERROR: ", err)
		return
	}
	s.Logger.Infof("Load1:%f, Load5:%f, Load15:%f", cpuLoad.Load1, cpuLoad.Load5, cpuLoad.Load15)
}
