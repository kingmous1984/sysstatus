package status

import (
	"fmt"
	"os/user"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/host"
)

// var kernel = syscall.NewLazyDLL("Kernel32.dll")

// SysInfo 系统信息
func (s *StatusOutput) SysInfo() {
	info, err := host.Info()
	if err != nil {
		s.Logger.Errorln("HOSTINFO ERROR: ", err)
	}
	//{"hostname":"DESKTOP-UFLTFSG","uptime":10825,"bootTime":1603781112,"procs":196,"os":"windows","platform":"Microsoft Windows 10 Home China",
	// "platformFamily":"Standalone Workstation","platformVersion":"10.0.18363 Build 18363","kernelVersion":"10.0.18363 Build 18363",
	// "kernelArch":"x86_64","virtualizationSystem":"","virtualizationRole":"","hostid":"04bc39b4-74dc-40c5-9639-5ae338af66bc"}
	s.Logger.Infoln("============================================SYS=============================================")
	s.Logger.Infof("golang版本：%s", runtime.Version())
	s.Logger.Infof("操作系统：%s", runtime.GOOS)
	s.Logger.Infof("Platform：%s %s %s", info.Platform, info.PlatformFamily, info.PlatformVersion)
	s.Logger.Infof("KernelArch：%s", info.KernelArch)
	s.Logger.Infof("HostID：%v", info.HostID)
	s.Logger.Infof("Hostname：%s", info.Hostname)
	s.Logger.Infof("开机时间：%s", time.Unix(int64(info.BootTime), 0).Format("2006/01/02 15:04:05"))
	s.Logger.Infof("开机时长：%s", GetStartTime(info.Uptime))
	// s.Logger.Infof("开机时长2：%s", GetStartTime2())

	users, err := host.Users()
	if err != nil {
		usr, err := user.Current()
		if err != nil {
			s.Logger.Errorln("Current USER ERROR: ", err)
		}
		s.Logger.Infof("当前用户：%s", usr.Username)
	} else {
		for _, u := range users {
			s.Logger.Infof("当前用户 %s@%s(%s) Started:%s", u.User, u.Host, u.Terminal, GetStartTime(uint64(u.Started)))
		}
	}
	// fmt.Println(user)

	// sysVersion, err := GetSystemVersion()
	// if err != nil {
	// 	s.Logger.Errorln("SYSTEM VERSION ERROR: ", err)
	// }
	// s.Logger.Infof("系统版本：%s", sysVersion)
	// s.Logger.Infoln(info)
	// 设备温度
	temperatures, err := host.SensorsTemperatures()
	if err != nil {
		s.Logger.Errorln("Temperatures ERROR: ", err)
	} else {
		for _, st := range temperatures {
			s.Logger.Infof("设备：%s 温度：%.2f", st.SensorKey, st.Temperature)
		}
	}
}

// GetStartTime 开机时长
func GetStartTime(uptime uint64) string {
	var d, h, m, s uint64
	var result string
	if uptime >= 60*60*24 {
		d = uptime / (60 * 60 * 24)
		result += fmt.Sprintf("%d天", d)
	}
	if uptime >= 60*60 {
		h = (uptime - d*60*60*24) / (60 * 60)
		result += fmt.Sprintf("%d小时", h)
	}
	if uptime >= 60 {
		m = (uptime - d*60*60*24 - h*60*60) / 60
		result += fmt.Sprintf("%d分钟", m)
	}
	s = uptime % 60
	result += fmt.Sprintf("%d秒", s)
	return result
}

// // GetHostName 获取主机名
// func GetHostName() string {
// 	hostName, err := os.Hostname()
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	return hostName
// }

// //GetStartTime2 开机时长
// func GetStartTime2() string {
// 	GetTickCount := kernel.NewProc("GetTickCount")
// 	r, _, _ := GetTickCount.Call()
// 	if r == 0 {
// 		return ""
// 	}
// 	ms := time.Duration(r * 1000 * 1000)
// 	return ms.String()
// }

//GetUserName 当前用户名

// func GetUserName() (string, error) {
// 	var size uint32 = 128
// 	var buffer = make([]uint16, size)
// 	user := syscall.StringToUTF16Ptr("USERNAME")
// 	domain := syscall.StringToUTF16Ptr("USERDOMAIN")
// 	r, err := syscall.GetEnvironmentVariable(user, &buffer[0], size)
// 	if err != nil {
// 		return "", err
// 	}
// 	buffer[r] = '@'
// 	old := r + 1
// 	if old >= size {
// 		return syscall.UTF16ToString(buffer[:r]), nil
// 	}
// 	r, err = syscall.GetEnvironmentVariable(domain, &buffer[old], size-old)
// 	return syscall.UTF16ToString(buffer[:old+r]), err
// }

// //GetSystemVersion 系统版本
// func GetSystemVersion() (string, error) {
// 	version, err := syscall.GetVersion()
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16), nil
// }
