package windows

import (
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	wapi "github.com/iamacarpet/go-win64api"
	"os"
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
		out, _ := json.Marshal(profiles)
		return retBool, string(out)
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

	out, _ := json.Marshal(profiles)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
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

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func FirewallProtocol(rule string, protocol string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	val, err := strconv.ParseInt(protocol, 10, 32)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}
	if info.Protocol == int32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func FirewallDirection(rule string, direction string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	val, err := strconv.ParseInt(direction, 10, 32)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}
	if info.Direction == int32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func FirewallAction(rule string, action string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	val, err := strconv.ParseInt(action, 10, 32)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}
	if info.Action == int32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func FirewallRuleProfile(rule string, profile string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	val, err := strconv.ParseInt(profile, 10, 32)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}
	if info.Profiles == int32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func FirewallRuleEnabled(rule string, value string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	val, err := strconv.ParseBool(value)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}
	if info.Enabled == val {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func FirewallEdge(rule string, traversal string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	info, err := wapi.FirewallRuleGet(rule)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	val, err := strconv.ParseBool(traversal)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}
	if info.EdgeTraversal == val {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}