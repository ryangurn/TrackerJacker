package cross

import (
	"github.com/goodhosts/hostsfile"
)

func HostExist(desiredHost string) (retBool bool) {
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

func HostIpExist(desiredIp string) (retBool bool) {
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