package cross

import (
	"encoding/json"
	"github.com/goodhosts/hostsfile"
)

func HostExist(desiredHost string) (retBool bool, retData string) {
	retBool = false
	hosts, _ := hostsfile.NewHosts()

	for _, line := range hosts.Lines {
		if line.Hosts != nil && line.IP != "" {
			for _, host := range line.Hosts {
				if host == desiredHost {
					retBool = true
					if out, err := json.Marshal(host); err == nil {
						retData = string(out)
					}
				}
			}
		}
	}
	return
}

func HostIpExist(desiredIp string) (retBool bool, retData string) {
	retBool = false
	hosts, _ := hostsfile.NewHosts()

	for _, line := range hosts.Lines {
		if line.Hosts != nil && line.IP != "" {
			if line.IP == desiredIp {
				retBool = true
				if out, err := json.Marshal(line); err == nil {
					retData = string(out)
				}
			}
		}
	}
	return
}