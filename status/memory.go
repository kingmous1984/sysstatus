package status

import (
	"github.com/shirou/gopsutil/mem"
)

// MemInfo 内存信息
func (s *StatusOutput) MemInfo() {
	s.Logger.Infoln("=============================================内存=============================================")
	v, err := mem.VirtualMemory()

	//{"total":16989884416,"available":10382397440,"used":6607486976,"usedPercent":38,"free":10382397440,
	// "active":0,"inactive":0,"wired":0,"laundry":0,"buffers":0,"cached":0,"writeback":0,"dirty":0,
	// "writebacktmp":0,"shared":0,"slab":0,"sreclaimable":0,"sunreclaim":0,"pagetables":0,"swapcached":0,
	// "commitlimit":0,"committedas":0,"hightotal":0,"highfree":0,"lowtotal":0,"lowfree":0,"swaptotal":0,
	// "swapfree":0,"mapped":0,"vmalloctotal":0,"vmallocused":0,"vmallocchunk":0,"hugepagestotal":0,
	// "hugepagesfree":0,"hugepagesize":0}
	// fmt.Println(v)
	if err == nil {
		tf, tu := memberNumber(v.Total)
		af, au := memberNumber(v.Available)
		uf, uu := memberNumber(v.Used)
		ff, fu := memberNumber(v.Free)
		s.Logger.Infof("Memory Total:%.2f%s, Available:%.2f%s, Used:%.2f%s, Free:%.2f%s, UsedPercent:%.2f%%",
			tf, tu, af, au, uf, uu, ff, fu, v.UsedPercent)
		bf, bu := memberNumber(v.Buffers)
		cf, cu := memberNumber(v.Cached)
		stf, stu := memberNumber(v.SwapTotal)
		scf, scu := memberNumber(v.SwapCached)
		sff, sfu := memberNumber(v.SwapFree)
		s.Logger.Infof("Memory Buffers:%.2f%s, Cached:%.2f%s, SwapTotal:%.2f%s, SwapCached:%.2f%s, SwapFree:%.2f%s",
			bf, bu, cf, cu, stf, stu, scf, scu, sff, sfu)
	}

	v2, err := mem.SwapMemory()
	// {"total":19540021248,"used":8305897472,"free":11234123776,"usedPercent":42.50710562993959,
	// "sin":0,"sout":0,"pgin":0,"pgout":0,"pgfault":0,"pgmajfault":0}
	// fmt.Println(v2)
	if err == nil {
		tf2, tu2 := memberNumber(v2.Total)
		uf2, uu2 := memberNumber(v2.Used)
		ff2, fu2 := memberNumber(v2.Free)
		s.Logger.Infof("Swap Total:%.2f%s, Used:%.2f%s, Free:%.2f%s, UsedPercent:%.2f%%",
			tf2, tu2, uf2, uu2, ff2, fu2, v.UsedPercent)
	}
}

// memberNumber 可识别数字
func memberNumber(n uint64) (float64, string) {
	if n >= 1024*1024*1024 {
		ret := float64(n) / (1024 * 1024 * 1024)
		return ret, "gb"
	}
	if n >= 1024*1024 {
		ret := float64(n) / (1024 * 1024)
		return ret, "mb"
	}
	if n >= 1024 {
		ret := float64(n) / (1024)
		return ret, "kb"
	}
	return float64(n), "byte"
}

// func Decimal(value float64) float64 {
// 	return math.Trunc(value*1e2+0.5) * 1e-2
// }

// Decimal value保留m位小数
// func Decimal(value float64, m uint) float64 {
// 	f := fmt.Sprintf("%%.%df", m)
// 	value, _ = strconv.ParseFloat(fmt.Sprintf(f, value), 64)
// 	return value
// }
