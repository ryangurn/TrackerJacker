package windows

import (
	"encoding/json"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
)

func FirewallEnabled(profile string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	val, err := strconv.ParseInt(profile, 10, 32)
	if err != nil {
		return
	}

	res, err := wapi.FirewallIsEnabled(int32(val))
	if err != nil {
		return
	}

	retBool = res
	if out, err := json.Marshal(res); err == nil {
		return retBool, string(out)
	}

	return
}

func FirewallProfile(firewall string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	profiles, err := wapi.FirewallCurrentProfiles()
	if err != nil {
		return
	}

	switch firewall {
	case "Public":
		retBool = profiles.Public
		if out, err := json.Marshal(profiles); err == nil {
			return retBool, string(out)
		}
		break
	case "Domain":
		retBool = profiles.Domain
		if out, err := json.Marshal(profiles); err == nil {
			return retBool, string(out)
		}
		break
	case "Private":
		retBool = profiles.Private
		if out, err := json.Marshal(profiles); err == nil {
			return retBool, string(out)
		}
		break
	}

	return
}

func FirewallRuleExists(rule string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.Name == rule {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallApplication(rule string, application string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.ApplicationName == application {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallService(rule string, service string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.ServiceName == service {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallLocalPort(rule string, port string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.LocalPorts == port {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallRemotePort(rule string, port string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.RemotePorts == port {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallLocalAddress(rule string, port string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.LocalAddresses == port {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallRemoteAddress(rule string, port string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.RemoteAddresses == port {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallICMP(rule string, value string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.ICMPTypesAndCodes == value {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallGrouping(rule string, group string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.Grouping == group {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallInterface(rule string, interfacee string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	if info.InterfaceTypes == interfacee {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallProtocol(rule string, protocol string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	val, err := strconv.ParseInt(protocol, 10, 32)
	if err != nil {
		return
	}
	if info.Protocol == int32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallDirection(rule string, direction string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	val, err := strconv.ParseInt(direction, 10, 32)
	if err != nil {
		return
	}
	if info.Direction == int32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallAction(rule string, action string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	val, err := strconv.ParseInt(action, 10, 32)
	if err != nil {
		return
	}
	if info.Action == int32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallRuleProfile(rule string, profile string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	val, err := strconv.ParseInt(profile, 10, 32)
	if err != nil {
		return
	}
	if info.Profiles == int32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallRuleEnabled(rule string, value string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	val, err := strconv.ParseBool(value)
	if err != nil {
		return
	}
	if info.Enabled == val {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func FirewallEdge(rule string, traversal string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		return
	}

	val, err := strconv.ParseBool(traversal)
	if err != nil {
		return
	}
	if info.EdgeTraversal == val {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}