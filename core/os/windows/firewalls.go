package windows

import (
	"TrackerJacker/core"
	"encoding/json"
	"fmt"
	"strconv"
)

type Firewall struct {
	CimClass struct {
		CimSuperClassName   string `json:"CimSuperClassName"`
		CimSuperClass       string `json:"CimSuperClass"`
		CimClassProperties  string `json:"CimClassProperties"`
		CimClassQualifiers  string `json:"CimClassQualifiers"`
		CimClassMethods     string `json:"CimClassMethods"`
		CimSystemProperties string `json:"CimSystemProperties"`
	} `json:"CimClass"`
	CimInstanceProperties []string `json:"CimInstanceProperties"`
	CimSystemProperties   struct {
		Namespace  string      `json:"Namespace"`
		ServerName string      `json:"ServerName"`
		ClassName  string      `json:"ClassName"`
		Path       interface{} `json:"Path"`
	} `json:"CimSystemProperties"`
	Profile                         string      `json:"Profile"`
	Enabled                         int         `json:"Enabled"`
	DefaultInboundAction            int         `json:"DefaultInboundAction"`
	DefaultOutboundAction           int         `json:"DefaultOutboundAction"`
	AllowInboundRules               int         `json:"AllowInboundRules"`
	AllowLocalFirewallRules         int         `json:"AllowLocalFirewallRules"`
	AllowLocalIPsecRules            int         `json:"AllowLocalIPsecRules"`
	AllowUserApps                   int         `json:"AllowUserApps"`
	AllowUserPorts                  int         `json:"AllowUserPorts"`
	AllowUnicastResponseToMulticast int         `json:"AllowUnicastResponseToMulticast"`
	NotifyOnListen                  int         `json:"NotifyOnListen"`
	EnableStealthModeForIPsec       int         `json:"EnableStealthModeForIPsec"`
	LogMaxSizeKilobytes             int         `json:"LogMaxSizeKilobytes"`
	LogAllowed                      int         `json:"LogAllowed"`
	LogBlocked                      int         `json:"LogBlocked"`
	LogIgnored                      int         `json:"LogIgnored"`
	Caption                         interface{} `json:"Caption"`
	Description                     interface{} `json:"Description"`
	ElementName                     string      `json:"ElementName"`
	InstanceID                      string      `json:"InstanceID"`
	DisabledInterfaceAliases        []string    `json:"DisabledInterfaceAliases"`
	LogFileName                     string      `json:"LogFileName"`
	Name                            string      `json:"Name"`
	PSComputerName                  interface{} `json:"PSComputerName"`
}

func FirewallEnabled(firewall string) (retBool bool) {
	retBool = false

	var firewalls []Firewall
	cmd := core.Command("Get-NetFirewallProfile -All | convertto-json")
	json.Unmarshal([]byte(cmd), &firewalls)

	for _, v := range firewalls {
		if v.Name == firewall {
			retBool = true
			return
		}
	}

	return
}

func FirewallMeta(firewall string, key string, value interface{}) (retBool bool) {
	retBool = false

	var firewalls []Firewall
	cmd := core.Command("Get-NetFirewallProfile -All | convertto-json")
	json.Unmarshal([]byte(cmd), &firewalls)

	// find and store the firewall we are dealing with
	var tmp Firewall
	for _, v := range firewalls {
		if v.Name == firewall {
			tmp = v
		}
	}

	// error if tmp is null
	if tmp.Name == "" {
		return
	}

	if key == "Profile" {
		if tmp.Profile == value.(string) {
			retBool = true
			return
		}
	} else if key == "Enabled" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.Enabled == int(ui) {
			retBool = true
			return
		}
	} else if key == "DefaultInboundAction" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.DefaultInboundAction == int(ui) {
			retBool = true
			return
		}
	} else if key == "DefaultOutboundAction" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.DefaultOutboundAction == int(ui) {
			retBool = true
			return
		}
	} else if key == "AllowInboundRules" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.AllowInboundRules == int(ui) {
			retBool = true
			return
		}
	} else if key == "AllowLocalFirewallRules" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.AllowLocalFirewallRules == int(ui) {
			retBool = true
			return
		}
	} else if key == "AllowLocalIPsecRules" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.AllowLocalIPsecRules == int(ui) {
			retBool = true
			return
		}
	} else if key == "AllowUserApps" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.AllowUserApps == int(ui) {
			retBool = true
			return
		}
	} else if key == "AllowUserPorts" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.AllowUserPorts == int(ui) {
			retBool = true
			return
		}
	} else if key == "AllowUnicastResponseToMulticast" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.AllowUnicastResponseToMulticast == int(ui) {
			retBool = true
			return
		}
	} else if key == "NotifyOnListen" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.NotifyOnListen == int(ui) {
			retBool = true
			return
		}
	} else if key == "EnableStealthModeForIPsec" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.Enabled == int(ui) {
			retBool = true
			return
		}
	} else if key == "LogMaxSizeKilobytes" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.LogMaxSizeKilobytes == int(ui) {
			retBool = true
			return
		}
	} else if key == "LogAllowed" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.LogAllowed == int(ui) {
			retBool = true
			return
		}
	} else if key == "LogBlocked" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.LogBlocked == int(ui) {
			retBool = true
			return
		}
	} else if key == "LogIgnored" {
		ui, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return
		}

		if tmp.LogIgnored == int(ui) {
			retBool = true
			return
		}
	} else if key == "Caption" {
		if tmp.Caption == value {
			retBool = true
			return
		}
	} else if key == "Description" {
		if tmp.Description == value {
			retBool = true
			return
		}
	} else if key == "ElementName" {
		if tmp.ElementName == value.(string) {
			retBool = true
			return
		}
	} else if key == "InstanceID" {
		if tmp.InstanceID == value.(string) {
			retBool = true
			return
		}
	} else if key == "LogFileName" {
		if tmp.LogFileName == value.(string) {
			retBool = true
			return
		}
	} else if key == "PSComputerName" {
		if tmp.PSComputerName == value {
			retBool = true
			return
		}
	}

	return
}

func FirewallParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) < 2 {
		return
	}

	if args[0] == "enabled" {
		if len(args) != 2 {
			return
		}

		if FirewallEnabled(args[1]) == result {
			retBool = true
			return
		}
	} else if args[0] == "meta" {
		if len(args) != 4 {
			return
		}

		if FirewallMeta(args[1], args[2], args[3]) == result {
			retBool = true
			return
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}
