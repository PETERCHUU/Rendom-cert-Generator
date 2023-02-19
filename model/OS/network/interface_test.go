package network_test

import (
	_network "linuxWebUI/model/OS/network"
	_config "linuxWebUI/model/config"
	"net"
	"testing"
	"unsafe"
)

func TestNewWindowSetup(t *testing.T) {
	config := _config.OS{}
	_network.NewWindowSetup(&config)
	if config.Interfaces == nil {
		t.Fail()
	}
	for k, v := range config.Interfaces {
		print("interface index:", k, "\n")
		print("interface name:", v.Name, "\n")
		print("interface ip:", v.Ip.IP.String(), "\n")
		print("interface ip:", (*net.IP)(unsafe.Pointer(&v.Ip.Mask)).String(), "\n")
		print("interface ip:", v.Ip.IP.String(), "\n")
		print("interface isDHCP:", v.IsDHCP, "\n")
		print("interface isOnline:", v.IsOnline, "\n")
	}
	print(config.Gateway.String(), "\n")
}
