package cross

import (
	"github.com/goodhosts/hostsfile"
)

func HostCheckExist(desiredHost string) (retBool bool) {
	retBool = false
	hosts, _ := hostsfile.NewHosts()

	for _, line := range hosts.Lines {
		if line.Hosts != nil && line.IP != "" {
			for _, host := range line.Hosts {
				if host == desiredHost {
					retBool = true
				}
			}
		}
	}
	return
}

func IpCheckExist(desiredIp string) (retBool bool) {
	retBool = false
	hosts, _ := hostsfile.NewHosts()

	for _, line := range hosts.Lines {
		if line.Hosts != nil && line.IP != "" {
			if line.IP == desiredIp {
				retBool = true
			}
		}
	}
	return
}
