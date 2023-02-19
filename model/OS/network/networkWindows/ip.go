package networkWindows

import (
	"net"
	"strconv"
	"unsafe"
)

func checkprefix(prefix string) (net.IPMask, bool) {
	if prefix == "" {
		return nil, false
	}
	if index, ok := strconv.Atoi(prefix); ok != nil {
		return net.CIDRMask(index, 32), true
	}
	mask := net.IPMask(net.ParseIP(prefix).To4())
	if i, _ := mask.Size(); i == 0 {
		return nil, false
	}
	return net.IPMask(net.ParseIP(prefix).To4()), true
}

// input sould be int(interface index), string(ip that can be XXX.XXX.XXX.XXX or XXX.XXX.XXX.XXX/X), string("" if no need, prefixlength or subnet mask will be find).
// out if false , Setip false.
func (m *Windowsset) Setipv4(i int, ip string, prefix string, gateway string) string { //input used to have gateway string
	// cmd :netsh int ip set address <index> static 192.168.0.10 255.255.255.0 192.168.0.1
	if m.Network.Interfaces[i] != nil {
		return "interface index error"
	}

	//check if ip avialable
	_, ipv4Net, err := net.ParseCIDR(ip)

	mask, ok := checkprefix(prefix)
	if (err != nil) && ok {
		ipv4Net = &net.IPNet{IP: net.ParseIP(ip), Mask: mask}
	}
	if (err != nil) && !ok {
		return err.Error()
	}

	cmd, cmderr, err := runcmd("netsh", "int", "ip", "set", "address", strconv.Itoa(i), "static", ipv4Net.IP.String(), (*net.IP)(unsafe.Pointer(&ipv4Net.Mask)).String())
	if cmderr != "" || err != nil {
		return cmderr + err.Error()
	}

	return cmd
}
