package networkWindows

import (
	"strconv"
	"strings"
)

func (m *Windowsset) SetDHCP(i int, DHCP bool) string {

	if DHCP {
		cmd, cmderr, err := runcmd("netsh", "int", "ip", "set", "address", strconv.Itoa(i), "dhcp")
		dnscmd, dnscmderr, dnserr := runcmd("netsh", "int", "ip", "set", "dns", strconv.Itoa(i), "dhcp")
		if strings.TrimSpace(cmderr+dnscmderr) != "" || err != nil || dnserr != nil {
			return cmderr + err.Error()
		}

		return cmd + dnscmd
	} else {
		return "don't need to set DHCP"
	}

}
