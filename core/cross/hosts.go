package cross

import (
	"fmt"
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

func HostParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) != 3 {
		return
	}

	// exists is the only secondary cmd at this point
	if args[1] == "exist" {
		if args[0] == "host" {
			if HostExist(args[2]) == result {
				retBool = true
			}
		} else if args[0] == "ip" {
			if HostIpExist(args[2]) == result {
				retBool = true
			}
		} else {
			fmt.Printf("Unrecognized Command: %s\n", args[0])
		}
	} else {
		fmt.Printf("Unrecognized Subcommand: %s\n", args[1])
	}

	return
}