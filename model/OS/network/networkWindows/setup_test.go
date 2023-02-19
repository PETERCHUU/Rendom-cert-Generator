package networkWindows_test

import (
	"bytes"
	"fmt"
	"linuxWebUI/model/OS/network/networkWindows"
	"os/exec"
	"testing"
)

func TestReadgateway(t *testing.T) {
	windowsnetset := new(networkWindows.Windowsset)
	windowsnetset.ReadGateway()
	print(windowsnetset.Network.Gateway)
}

func TestRuncmd(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("netsh", "int", "ip", "set", "address", "13", "static", "10.120.8.101", "255.255.255.0", "10.120.8.254")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	cmdout := stdout.String()
	cmderr, _, _ := stderr.ReadRune()
	fmt.Printf("1\n%s\n2\n", cmdout)
	fmt.Printf("%c\n3\n", cmderr)
	print(err.Error(), "\n4\n")
}

// ifaces, err := net.Interfaces()
// if err != nil {
// 	println(err)
// }

// for _, i := range ifaces {
// 	if i.Flags&net.FlagBroadcast != 0 && i.Flags&net.FlagMulticast != 0 && i.Flags&net.FlagLoopback == 0 {

// 		print(i.Name, "\n")
// 		print(i.Index, "\n\n")
// 		addrs, err := i.Addrs()
// 		if err != nil {
// 			println(err)
// 		}
// 		for _, addr := range addrs {
// 			var ip net.IP
// 			var prefix net.IPMask
// 			switch v := addr.(type) {
// 			case *net.IPNet:
// 				ip = v.IP
// 				prefix = v.Mask
// 			default:
// 				continue
// 			}
// 			// lastind := strings.LastIndex(ip, ":")
// 			// mask := ip[:lastind]
// 			// ipforchack := ip[lastind+1:]
// 			// checkip := net.ParseIP(ipforchack)
// 			// ipv4 := checkip.To4()
// 			// if ipv4 != nil {
// 			// 	prefix, err := strconv.Atoi(mask)
// 			// 	if err != nil {
// 			// 		print(err)
// 			// 	} else {
// 			// 		netcard.Prefix = uint8(prefix)
// 			// 		netcard.Ip = ipv4
// 			// 		netcard.Gateway = readgateway(ip)
// 			// 		break
// 			// 	}
// 			// } else {
// 			// 	print("")
// 			ipstring := ip.String()
// 			maskones, _ := prefix.Size()
// 			print(ipstring, "\n")
// 			print(maskones, "\n")
// 		}
// 		print("\n\n")
// 	}
// }
