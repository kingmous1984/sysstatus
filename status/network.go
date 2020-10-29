package status

import (
	"github.com/shirou/gopsutil/net"
)

// NetInfo 网络信息
func (s *StatusOutput) NetInfo() {
	s.Logger.Infoln("=============================================网络=============================================")
	ioCounters, err := net.IOCounters(true)
	if err != nil {
		s.Logger.Errorln("IOCOUNTERS ERROR: ", err)
	}
	// fmt.Println(ioCounters)
	all, err := net.IOCounters(false)
	if err != nil {
		s.Logger.Errorln("IOCOUNTERS ERROR: ", err)
	}
	// fmt.Println(all)
	// [{"name":"all","bytesSent":85144263,"bytesRecv":461580914,"packetsSent":384940,"packetsRecv":319568,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0}]

	// [{"name":"VirtualBox Host-Only Network","bytesSent":0,"bytesRecv":0,"packetsSent":0,"packetsRecv":0,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0}
	// {"name":"以太网","bytesSent":85144263,"bytesRecv":461580914,"packetsSent":384940,"packetsRecv":319568,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0}
	// {"name":"Loopback Pseudo-Interface 1","bytesSent":0,"bytesRecv":0,"packetsSent":0,"packetsRecv":0,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0}]
	asf, asu := memberNumber(all[0].BytesSent)
	arf, aru := memberNumber(all[0].BytesRecv)
	s.Logger.Infof("NET Name:%s BytesSent:%.2f%s, BytesRecv:%.2f%s, PacketsSent:%d, PacketsRecv:%d, Errin:%d, Errout:%d, Dropin:%d, Dropout:%d, Fifoin:%d, Fifoout:%d",
		all[0].Name, asf, asu, arf, aru, all[0].PacketsSent, all[0].PacketsRecv, all[0].Errin, all[0].Errout, all[0].Dropin, all[0].Dropout, all[0].Fifoin, all[0].Fifoout)
	for _, ioc := range ioCounters {
		iocsf, iocsu := memberNumber(ioc.BytesSent)
		iocrf, iocru := memberNumber(ioc.BytesRecv)
		s.Logger.Infof("NET Name:%s BytesSent:%.2f%s, BytesRecv:%.2f%s, PacketsSent:%d, PacketsRecv:%d, Errin:%d, Errout:%d, Dropin:%d, Dropout:%d, Fifoin:%d, Fifoout:%d",
			ioc.Name, iocsf, iocsu, iocrf, iocru, ioc.PacketsSent, ioc.PacketsRecv, ioc.Errin, ioc.Errout, ioc.Dropin, ioc.Dropout, ioc.Fifoin, ioc.Fifoout)
	}
}
