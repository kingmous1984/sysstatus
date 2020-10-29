package status

import (
	"github.com/shirou/gopsutil/disk"
)

// DiskInfo 磁盘信息
func (s *StatusOutput) DiskInfo() {
	s.Logger.Infoln("=============================================硬盘=============================================")
	ioCounters, err := disk.IOCounters("")
	if err != nil {
		s.Logger.Errorln("IOCOUNTERS ERROR: ", err)
	}
	// fmt.Println(ioCounters)
	// map[C::{"readCount":257409,"mergedReadCount":0,"writeCount":272518,"mergedWriteCount":0,"readBytes":9112945664,"writeBytes":5878823936,"readTime":222,"writeTime":56,"iopsInProgress":0,"ioTime":0,"weightedIO":0,"name":"C:","serialNumber":"","label":""}
	// 	E::{"readCount":82,"mergedReadCount":0,"writeCount":23,"mergedWriteCount":0,"readBytes":18553344,"writeBytes":94208,"readTime":1,"writeTime":0,"iopsInProgress":0,"ioTime":0,"weightedIO":0,"name":"E:","serialNumber":"","label":""}
	// 	F::{"readCount":66,"mergedReadCount":0,"writeCount":17,"mergedWriteCount":0,"readBytes":13564416,"writeBytes":69632,"readTime":0,"writeTime":0,"iopsInProgress":0,"ioTime":0,"weightedIO":0,"name":"F:","serialNumber":"","label":""}]
	partitions, err := disk.Partitions(true) //所有分区
	if err != nil {
		s.Logger.Errorln("PATITION ERROR: ", err)
	}
	for _, p := range partitions {
		// fmt.Println(p.String())
		// {"device":"C:","mountpoint":"C:","fstype":"NTFS","opts":"rw.compress"}
		s.Logger.Infof("Device:%s, Mountpoint:%s, Fstype:%s, Opts:%s", p.Device, p.Mountpoint, p.Fstype, p.Opts)

		u, err := disk.Usage(p.Mountpoint)
		if err != nil {
			s.Logger.Errorln("Device:%s USAGE ERROR: ", p.Device, err)
		}
		// {"path":"C:","fstype":"","total":511478624256,"free":434537721856,"used":76940902400,"usedPercent":15.042838302757758,
		// "inodesTotal":0,"inodesUsed":0,"inodesFree":0,"inodesUsedPercent":0}
		// fmt.Println(u)
		tf, tu := memberNumber(u.Total)
		uf, uu := memberNumber(u.Used)
		ff, fu := memberNumber(u.Free)
		s.Logger.Infof("USAGE Total:%.2f%s, Used:%.2f%s, Free:%.2f%s, UsedPercent:%.2f%%, InodesTotal:%d, InodesUsed:%d, InodesFree:%d, InodesUsedPercent:%.2f%%",
			tf, tu, uf, uu, ff, fu, u.UsedPercent, u.InodesTotal, u.InodesUsed, u.InodesFree, u.InodesUsedPercent)
		// {"readCount":106,"mergedReadCount":0,"writeCount":1657,"mergedWriteCount":0,"readBytes":14486528,"writeBytes":63649280,
		// "readTime":1,"writeTime":3,"iopsInProgress":0,"ioTime":0,"weightedIO":0,"name":"F:","serialNumber":"","label":""}
		ic := ioCounters[p.Device]
		rf, ru := memberNumber(ic.ReadBytes)
		wf, wu := memberNumber(ic.WriteBytes)
		s.Logger.Infof("ReadCount:%d, MergedReadCount:%d, WriteCount:%d, MergedWriteCount:%d, ReadBytes:%.2f%s, WriteBytes:%.2f%s,ReadTime:%d, WriteTime:%d, IopsInProgress:%d, IoTime:%d, WeightedIO:%d,Name:%s, SerialNumber:%s, Label:%s",
			ic.ReadCount, ic.MergedReadCount, ic.WriteCount, ic.MergedWriteCount, rf, ru, wf, wu, ic.ReadTime, ic.WriteTime, ic.IopsInProgress, ic.IoTime, ic.WeightedIO, ic.Name, ic.SerialNumber, ic.Label)
	}
}
