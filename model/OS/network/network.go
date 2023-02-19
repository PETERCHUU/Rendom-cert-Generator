package network

import (
	_windows "linuxWebUI/model/OS/network/networkWindows"
	_config "linuxWebUI/model/config"
	"net"
	"strings"
)

type OsSetNet interface {
	Setipv4(i int, ip string, prefix string, gateway string) string
	SetDHCP(i int, DHCP bool) string
	SetDNS(i int, DNS string) string
	ReadGateway()
}

func NewWindowSetup(network *_config.OS) OsSetNet {
	ScanNetwork(network)
	netinterface := &_windows.Windowsset{
		Network: network,
	}
	netinterface.ReadGateway()
	return netinterface
}

func ScanNetwork(network *_config.OS) {
	network.Interfaces = make(map[int]*_config.Network)
	ifaces, err := net.Interfaces()
	if err != nil {
		println(err)
	}
	for _, i := range ifaces {
		if i.Flags&net.FlagBroadcast+i.Flags&net.FlagMulticast != 0 && i.Flags&net.FlagLoopback == 0 {
			netcard := &_config.Network{}
			if strings.Contains(i.Flags.String(), "up") {
				netcard.IsOnline = true
			} else {
				netcard.IsOnline = false
			}
			netcard.Name = i.Name
			//netcard.Index = i.Index
			addrs, err := i.Addrs()

			if err != nil {
				println(err)
			}

			for _, addr := range addrs {

				switch v := addr.(type) {
				case *net.IPNet:
					netcard.Ip.IP = v.IP
					netcard.Ip.Mask = v.Mask

				default:
					continue
				}
			}
			//network.InterFaces = append(network.InterFaces, netcard)
			network.Interfaces[i.Index] = netcard
		}
	}
}

/*net stop w32time
w32tm /config /syncfromflags:manual /manualpeerlist:"0.it.pool.ntp.org 1.it.pool.ntp.org 2.it.pool.ntp.org 3.it.pool.ntp.org"
net start w32time
w32tm /config /update
w32tm /resync /rediscover*/
