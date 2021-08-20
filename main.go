package main

import (
	"TrackerJacker/core/enc"
	"TrackerJacker/core/os/cross"
	"TrackerJacker/core/os/windows"
	"TrackerJacker/core/parsing"
	"TrackerJacker/core/submission"
	"fmt"
	"github.com/bugsnag/bugsnag-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
)

const encKey = "Password"
const payload = "payload.txt"
const baseURL = "http://cste.test"

func generatePayload(data []byte) {
	//data, _ := ioutil.ReadFile(inFile)
	enc.EncryptFile(payload, data, encKey)
}

func parsePayload() parsing.PayloadType {
	str := enc.DecryptFile(payload, encKey)

	var payload parsing.PayloadType
	parsing.ParsePayload(str, &payload)
	return payload
}

// todo: work on verbose a bit more
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file")
		err = submission.Initialize(baseURL)
		if err != nil {
			return
		}
	}

	// setup bugsnag
	bugsnagKey := os.Getenv("BUGSNAG_KEY")
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          bugsnagKey,
		ReleaseStage:    "alpha",
		AppVersion: 	"0.0.1",
		ProjectPackages: []string{"main"},
	})

	// get check data
	checks, err := submission.GetPayload()
	if err != nil {
		return
	}

	// print payload information
	//fmt.Println(string(checks))

	// generate the payload
	generatePayload(checks)
	// get the payload
	payload := parsePayload()

	// set batch value
	batch, err := uuid.NewUUID()
	if err != nil {
		return
	}

	// loop through payload items
	for i := 0; i < len(payload); i++ {

		if payload.GetSpace(i) == "files" {
			// files rule implementation
			if payload.GetAction(i) == "exists" {
				// exists
				result, data := cross.FileExists(payload.GetParameter(i, "path"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "does_not_exist" {
				// negate exists
				result, data := cross.FileExists(payload.GetParameter(i, "path"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "hash" {
				str, err := cross.FileHash(payload.GetParameter(i, "path"))
				if err != nil {
					// hashing error
					payload.DebugPrint(i, false)
				} else {
					// no error
					result := str == payload.GetParameter(i, "hash")
					payload.DebugPrint(i, result) // debug print
					submission.Send(str, result, payload[i].ID, batch) // send score
				}
			}
			// end files
		} else if payload.GetSpace(i) == "hosts" {
			// hosts rule implementation
			if payload.GetAction(i) == "ip_exists" {
				// ip address exists
				result, data := cross.HostIpExist(payload.GetParameter(i, "ip"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "ip_does_not_exist" {
				// ip address does not exist
				result, data := cross.HostIpExist(payload.GetParameter(i, "ip"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "host_exist" {
				// host exists
				result, data := cross.HostExist(payload.GetParameter(i, "host"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "host_does_not_exist" {
				// host does not exist
				result, data := cross.HostExist(payload.GetParameter(i, "host"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end hosts
		} else if payload.GetSpace(i) == "users" {
			// users rule implementation
			if payload.GetAction(i) == "exists" {
				result, data := cross.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "does_not_exist" {
				result, data := cross.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end users
		} else if payload.GetSpace(i) == "bitlocker" {
			// bitlocker rule implementation
			if payload.GetAction(i) == "drive_locked" {
				result, data := windows.BitlockerDriveLocked(payload.GetParameter(i, "drive"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "drive_unlocked" {
				result, data := windows.BitlockerDriveLocked(payload.GetParameter(i, "drive"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "device_id_is" {
				result, data := windows.BitlockerDeviceID(payload.GetParameter(i, "drive"), payload.GetParameter(i, "device"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "device_id_is_not" {
				result, data := windows.BitlockerDeviceID(payload.GetParameter(i, "drive"), payload.GetParameter(i, "device"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "persistent_volume_id_is" {
				result, data := windows.BitlockerPersistentVolumeID(payload.GetParameter(i, "drive"), payload.GetParameter(i, "volume"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "persistent_volume_id_is_not" {
				result, data := windows.BitlockerPersistentVolumeID(payload.GetParameter(i, "drive"), payload.GetParameter(i, "volume"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "conversion_status_is" {
				result, data := windows.BitlockerConversionStatus(payload.GetParameter(i, "drive"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "conversion_status_is_not" {
				result, data := windows.BitlockerConversionStatus(payload.GetParameter(i, "drive"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "protection_status_is" {
				result, data := windows.BitlockerProtectionStatus(payload.GetParameter(i, "drive"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "protection_status_is_not" {
				result, data := windows.BitlockerProtectionStatus(payload.GetParameter(i, "drive"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end bitlocker
		} else if payload.GetSpace(i) == "firewall" {
			// firewall rule implementation
			if payload.GetAction(i) == "firewall_enabled" {
				result, data := windows.FirewallEnabled(payload.GetParameter(i, "firewall"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "firewall_not_enabled" {
				result, data := windows.FirewallEnabled(payload.GetParameter(i, "firewall"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "profile_is" {
				result, data := windows.FirewallProfile(payload.GetParameter(i, "profile"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "profile_is_not" {
				result, data := windows.FirewallProfile(payload.GetParameter(i, "profile"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "rule_exists" {
				result, data := windows.FirewallRuleExists(payload.GetParameter(i, "rule"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "rule_does_not_exist" {
				result, data := windows.FirewallRuleExists(payload.GetParameter(i, "rule"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "application_is" {
				result, data := windows.FirewallApplication(payload.GetParameter(i, "rule"), payload.GetParameter(i, "application"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "application_is_not" {
				result, data := windows.FirewallApplication(payload.GetParameter(i, "rule"), payload.GetParameter(i, "application"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "service_is" {
				result, data := windows.FirewallService(payload.GetParameter(i, "rule"), payload.GetParameter(i, "service"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "service_is_not" {
				result, data := windows.FirewallService(payload.GetParameter(i, "rule"), payload.GetParameter(i, "service"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "local_port_is" {
				result, data := windows.FirewallLocalPort(payload.GetParameter(i, "rule"), payload.GetParameter(i, "port"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "local_port_is_not" {
				result, data := windows.FirewallLocalPort(payload.GetParameter(i, "rule"), payload.GetParameter(i, "port"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "remote_port_is" {
				result, data := windows.FirewallRemotePort(payload.GetParameter(i, "rule"), payload.GetParameter(i, "port"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "remote_port_is_not" {
				result, data := windows.FirewallRemotePort(payload.GetParameter(i, "rule"), payload.GetParameter(i, "port"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "local_address_is" {
				result, data := windows.FirewallLocalAddress(payload.GetParameter(i, "rule"), payload.GetParameter(i, "address"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "local_address_is_not" {
				result, data := windows.FirewallLocalAddress(payload.GetParameter(i, "rule"), payload.GetParameter(i, "address"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "remote_address_is" {
				result, data := windows.FirewallRemoteAddress(payload.GetParameter(i, "rule"), payload.GetParameter(i, "address"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "remote_address_is_not" {
				result, data := windows.FirewallRemoteAddress(payload.GetParameter(i, "rule"), payload.GetParameter(i, "address"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "icmp_types_and_codes_is" {
				result, data := windows.FirewallICMP(payload.GetParameter(i, "rule"), payload.GetParameter(i, "value"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "icmp_types_and_codes_is_not" {
				result, data := windows.FirewallICMP(payload.GetParameter(i, "rule"), payload.GetParameter(i, "value"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "grouping_is" {
				result, data := windows.FirewallGrouping(payload.GetParameter(i, "rule"), payload.GetParameter(i, "group"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "grouping_is_not" {
				result, data := windows.FirewallGrouping(payload.GetParameter(i, "rule"), payload.GetParameter(i, "group"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "interface_type_is" {
				result, data := windows.FirewallInterface(payload.GetParameter(i, "rule"), payload.GetParameter(i, "interface"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "interface_type_is_not" {
				result, data := windows.FirewallInterface(payload.GetParameter(i, "rule"), payload.GetParameter(i, "interface"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "protocol_is" {
				result, data := windows.FirewallProtocol(payload.GetParameter(i, "rule"), payload.GetParameter(i, "protocol"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "protocol_is_not" {
				result, data := windows.FirewallProtocol(payload.GetParameter(i, "rule"), payload.GetParameter(i, "protocol"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "direction_is" {
				result, data := windows.FirewallDirection(payload.GetParameter(i, "rule"), payload.GetParameter(i, "direction"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "direction_is_not" {
				result, data := windows.FirewallDirection(payload.GetParameter(i, "rule"), payload.GetParameter(i, "direction"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "action_is" {
				result, data := windows.FirewallAction(payload.GetParameter(i, "rule"), payload.GetParameter(i, "action"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "action_is_not" {
				result, data := windows.FirewallAction(payload.GetParameter(i, "rule"), payload.GetParameter(i, "action"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "rule_profile_is" {
				result, data := windows.FirewallRuleProfile(payload.GetParameter(i, "rule"), payload.GetParameter(i, "profile"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "rule_profile_is_not" {
				result, data := windows.FirewallRuleProfile(payload.GetParameter(i, "rule"), payload.GetParameter(i, "profile"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "enabled_is" {
				result, data := windows.FirewallRuleEnabled(payload.GetParameter(i, "rule"), payload.GetParameter(i, "value"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "enabled_is_not" {
				result, data := windows.FirewallRuleEnabled(payload.GetParameter(i, "rule"), payload.GetParameter(i, "value"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "edge_traversal_is" {
				result, data := windows.FirewallEdge(payload.GetParameter(i, "rule"), payload.GetParameter(i, "traversal"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "edge_traversal_is_not" {
				result, data := windows.FirewallEdge(payload.GetParameter(i, "rule"), payload.GetParameter(i, "traversal"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end firewall
		}
	}
	batch, err = uuid.NewUUID()
}