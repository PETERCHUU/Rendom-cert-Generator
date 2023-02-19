package networkWindows

import (
	"bytes"
	"io"
	_config "linuxWebUI/model/config"
	"net"
	"os/exec"
	"strconv"
)

//usecase :os network setting

type Windowsset struct {
	Network *_config.OS
}

var Windowsseting Windowsset

func (m *Windowsset) ReadGateway() {
	route := exec.Command("route", "print", "-4", "0.0.0.0")
	cmd := exec.Command("findstr", "0.0.0.0")
	r, w := io.Pipe()
	route.Stdout = w
	cmd.Stdin = r
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	route.Start()
	cmd.Start()
	route.Wait()
	w.Close()
	cmd.Wait()
	if stderr.Bytes() != nil {
		print(stderr.String())
	}
	r.Close()
	listOfIp := bytes.Fields(stdout.Bytes())
	m.Network.Gateway = net.ParseIP(string(listOfIp[2]))
}

func (m *Windowsset) SetDNS(i int, DNS string) string {
	cmd, cmderr, err := runcmd("netsh", "int", "ip", "set", "dns", strconv.Itoa(i), DNS)
	if cmderr != "" || err != nil {
		return cmderr + err.Error()
	}

	return cmd
}

// func ScanNetwork(network *_config.OS) {
// 	ifaces, err := net.Interfaces()
// 	if err != nil {
// 		println(err)
// 	}
// 	for _, i := range ifaces {
// 		if i.Flags&net.FlagBroadcast != 0 && i.Flags&net.FlagMulticast != 0 && i.Flags&net.FlagLoopback == 0 {
// 			netcard := _config.Network{}
// 			netcard.Name = i.Name
// 			//netcard.Index = i.Index
// 			addrs, err := i.Addrs()

// 			if err != nil {
// 				println(err)
// 			}

// 			for _, addr := range addrs {

// 				var ip net.IP
// 				var prefix net.IPMask
// 				switch v := addr.(type) {
// 				case *net.IPNet:
// 					ip = v.IP
// 					prefix = v.Mask

// 				default:
// 					continue
// 				}
// 				netcard.Ip = ip
// 				prefixInt, _ := prefix.Size()
// 				netcard.Prefix = uint8(prefixInt)
// 			}
// 			//network.InterFaces = append(network.InterFaces, netcard)

// 		}
// 	}
// }

// func checkip(ip string) error {
// 	//check if ip avialable
// 	forCheck := strings.Split(ip, ".")
// 	for i := range forCheck {
// 		e, err := strconv.Atoi(forCheck[i])
// 		if err != nil || e < 255 || e > 0 {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func readnetwork(index int) {
// 		readip := fmt.Sprintf(`(Get-NetIPConfiguration | Where-Object {$_.IPv4DefaultGateway -ne $null -and $_.NetAdapter.Status -ne "Disconnected"}).IPv4Address.IPAddress`)
// }

// func (m *Windowsset) SetNTP(i int, address string) {
// 	ntpshell := fmt.Sprintf(`Set-ItemProperty -Path
// 	"HKLM:\SYSTEM\CurrentControlSet\Services\w32time\Parameters" -Name
// 	"NtpServer" -Value %s ,0x9`, address)
// 	command(ntpshell)
// 	command("Restart-Service w32Time")
// }

// func getNetworkIndex() []int {
// 	indoxNo, errout, err := runcmd(`Get-NetAdapter | Where-Object -Property Status -EQ Up | Select-Object -Property ifIndex`)
// 	if err != nil {
// 		log.Printf("error: %v\n", err)
// 		fmt.Print(errout)
// 	}
// 	s := strings.Split(indoxNo, "\n")
// 	indexList := []int{}
// 	for i := 3; i <= len(s); i++ {
// 		n := strings.TrimSpace(s[i])
// 		inte, err := strconv.Atoi(n)
// 		if err != nil {
// 			fmt.Printf("error by getting index %s", err.Error())
// 			continue
// 		}
// 		indexList = append(indexList, inte)
// 	}

// 	return indexList
// }

// func shellOutputWithDir(command, dir string) (string, string, error) {
// 	var stdout bytes.Buffer
// 	var stderr bytes.Buffer

// 	cmd := exec.Command("powershell.exe", command)
// 	cmd.Dir = dir
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr

// 	err := cmd.Run()

// 	return stdout.String(), stderr.String(), err
// }
