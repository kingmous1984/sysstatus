package status

import (
	"sort"
	"time"

	"github.com/shirou/gopsutil/process"
)

type processAndPersent struct {
	process *process.Process
	persent float64
}

// ProcessInfo 进程信息
func (s *StatusOutput) ProcessInfo() {
	s.Logger.Infoln("=============================================进程=============================================")
	processes, err := process.Processes()
	if err != nil {
		s.Logger.Errorln("Processes ERROR: ", err)
	}

	// var cpuTom10, memTom10 []*processAndPersent
	cpuTom10 := make([]*processAndPersent, 10)
	memTom10 := make([]*processAndPersent, 10)
	// 筛选进程
	for _, p := range processes {
		cp, err := p.CPUPercent()
		if err == nil {
			CheckProcess(cp, p, cpuTom10)
		}
		mp, err := p.MemoryPercent()
		if err == nil {
			CheckProcess(float64(mp), p, memTom10)
		}
	}
	//  排序
	sortProcess(cpuTom10)
	sortProcess(memTom10)
	// 输出
	s.Logger.Infoln("----------------------------PROCESS CPU----------------------------")
	s.PrintProcess(cpuTom10)
	s.Logger.Infoln("----------------------------PROCESS MEMORY----------------------------")
	s.PrintProcess(memTom10)
}

// PrintProcess 打印输出
func (s *StatusOutput) PrintProcess(pList []*processAndPersent) {
	for index, p := range pList {
		pname, _ := p.process.Name()
		username, _ := p.process.Username()
		status, _ := p.process.Status()
		numThreads, _ := p.process.NumThreads()
		createTime, _ := p.process.CreateTime()
		s.Logger.Infof("index:%d, PID:%d, Name:%s, Persent:%.2f, Username:%s, Status:%s, NumThreads:%d, CreateTime:%s",
			index+1, p.process.Pid, pname, p.persent, username, status, numThreads, time.Unix(createTime, 0).Format("2006/01/02 15:04:05"))
		memoryInfo, err := p.process.MemoryInfo()
		if err == nil {
			s.Logger.Infof("Processe MemoryInfo:%s", memoryInfo.String())
		}
		ioCountersStat, err := p.process.IOCounters()
		if err == nil {
			s.Logger.Infof("Processe IOCounters:%s", ioCountersStat.String())
		}
	}
}

// CheckProcess 过滤进程
func CheckProcess(persent float64, p *process.Process, pList []*processAndPersent) {
	pap := processAndPersent{
		process: p,
		persent: persent,
	}
	if len(pList) < 10 {
		pList[len(pList)] = &pap
	} else {
		var index int = -1
		minPersent := persent
		for i, tmpP := range pList {
			if tmpP == nil {
				pList[i] = &pap
				return
			}
			if tmpP.persent < minPersent {
				minPersent = tmpP.persent
				index = i
			}
		}
		if index >= 0 {
			pList[index] = &pap
		}
	}
}

func sortProcess(pList []*processAndPersent) {
	sort.SliceStable(pList, func(i, j int) bool {
		if pList[i].persent > pList[j].persent {
			return true
		}
		return false
	})
}
