package main

import (
	"TrackerJacker/core/enc"
	"TrackerJacker/core/os/cross"
	"TrackerJacker/core/os/windows"
	"TrackerJacker/core/parsing"
	"TrackerJacker/core/submission"
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
		//fmt.Println("Error loading env file")
		err = submission.Initialize(baseURL)
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
	}

	// setup bugsnag
	bugsnagKey := os.Getenv("BUGSNAG_KEY")
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          bugsnagKey,
		ReleaseStage:    "alpha",
		AppVersion: 	"0.0.1",
		ProjectPackages: []string{"main"},
	})
	bugsnag.AutoNotify()

	// get check data
	checks, err := submission.GetPayload()
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

	// print payload information
	//fmt.Println(string(checks))

	// generate the payload
	generatePayload(checks)
	// get the payload
	payload := parsePayload()

	// set batch value
	batch, err := uuid.NewUUID()
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

	// loop through payload items
	for i := 0; i < len(payload); i++ {

		if payload.GetSpace(i) == "files" {
			// files rule implementation
			if payload.GetAction(i) == "exists" {
				// exists
				result, data := cross.FileExists(payload.GetParameter(i, "path"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "does_not_exist" {
				// negate exists
				result, data := cross.FileExists(payload.GetParameter(i, "path"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "hash" {
				str, err := cross.FileHash(payload.GetParameter(i, "path"))
				if err != nil {
					bugsnag.Notify(err, bugsnag.HandledState{
						SeverityReason:   bugsnag.SeverityReasonUnhandledError,
						OriginalSeverity: bugsnag.SeverityError,
						Unhandled:      false,
					}, bugsnag.MetaData{
						"ENV": {
							"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
							"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
							"IMAGE": os.Getenv("IMAGE"),
							"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
							"SERVER": os.Getenv("SERVER"),
						},
						"CHECK": {
							"ID": payload[i].ID,
							"IMAGE": payload[i].Image,
							"RULE": payload[i].GetRule,
							"RULE_ACTION": payload[i].RuleAction,
							"SCORE": payload[i].Score,
							"COMMENT": payload[i].Comment,
							"SPACE": payload.GetSpace(i),
							"ACTION": payload.GetAction(i),
							"PARAMETER": payload[i].Parameters,
						},
					}, bugsnag.Event{
						GroupingHash: payload.GetSpace(i)+"."+payload.GetAction(i),
					})
				} else {
					result := str == payload.GetParameter(i, "expected")
					payload.DebugPrint(i, result, str) // debug print
					submission.Send(str, result, payload[i].ID, batch) // send score
				}
			}
			// end files
		} else if payload.GetSpace(i) == "hosts" {
			// hosts rule implementation
			if payload.GetAction(i) == "ip_exists" {
				// ip address exists
				result, data := cross.HostIpExist(payload.GetParameter(i, "ip"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "ip_does_not_exist" {
				// ip address does not exist
				result, data := cross.HostIpExist(payload.GetParameter(i, "ip"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "host_exist" {
				// host exists
				result, data := cross.HostExist(payload.GetParameter(i, "host"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "host_does_not_exist" {
				// host does not exist
				result, data := cross.HostExist(payload.GetParameter(i, "host"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end hosts
		} else if payload.GetSpace(i) == "users" {
			// users rule implementation
			if payload.GetAction(i) == "exists" {
				result, data := cross.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "does_not_exist" {
				result, data := cross.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end users
		} else if payload.GetSpace(i) == "bitlocker" {
			// bitlocker rule implementation
			if payload.GetAction(i) == "drive_locked" {
				result, data := windows.BitlockerDriveLocked(payload.GetParameter(i, "drive"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "drive_unlocked" {
				result, data := windows.BitlockerDriveLocked(payload.GetParameter(i, "drive"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "device_id_is" {
				result, data := windows.BitlockerDeviceID(payload.GetParameter(i, "drive"), payload.GetParameter(i, "device"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "device_id_is_not" {
				result, data := windows.BitlockerDeviceID(payload.GetParameter(i, "drive"), payload.GetParameter(i, "device"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "persistent_volume_id_is" {
				result, data := windows.BitlockerPersistentVolumeID(payload.GetParameter(i, "drive"), payload.GetParameter(i, "volume"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "persistent_volume_id_is_not" {
				result, data := windows.BitlockerPersistentVolumeID(payload.GetParameter(i, "drive"), payload.GetParameter(i, "volume"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "conversion_status_is" {
				result, data := windows.BitlockerConversionStatus(payload.GetParameter(i, "drive"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "conversion_status_is_not" {
				result, data := windows.BitlockerConversionStatus(payload.GetParameter(i, "drive"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "protection_status_is" {
				result, data := windows.BitlockerProtectionStatus(payload.GetParameter(i, "drive"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "protection_status_is_not" {
				result, data := windows.BitlockerProtectionStatus(payload.GetParameter(i, "drive"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end bitlocker
		} else if payload.GetSpace(i) == "firewall" {
			// firewall rule implementation
			if payload.GetAction(i) == "firewall_enabled" {
				result, data := windows.FirewallEnabled(payload.GetParameter(i, "firewall"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "firewall_not_enabled" {
				result, data := windows.FirewallEnabled(payload.GetParameter(i, "firewall"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "profile_is" {
				result, data := windows.FirewallProfile(payload.GetParameter(i, "profile"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "profile_is_not" {
				result, data := windows.FirewallProfile(payload.GetParameter(i, "profile"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "rule_exists" {
				result, data := windows.FirewallRuleExists(payload.GetParameter(i, "rule"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "rule_does_not_exist" {
				result, data := windows.FirewallRuleExists(payload.GetParameter(i, "rule"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "application_is" {
				result, data := windows.FirewallApplication(payload.GetParameter(i, "rule"), payload.GetParameter(i, "application"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "application_is_not" {
				result, data := windows.FirewallApplication(payload.GetParameter(i, "rule"), payload.GetParameter(i, "application"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "service_is" {
				result, data := windows.FirewallService(payload.GetParameter(i, "rule"), payload.GetParameter(i, "service"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "service_is_not" {
				result, data := windows.FirewallService(payload.GetParameter(i, "rule"), payload.GetParameter(i, "service"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "local_port_is" {
				result, data := windows.FirewallLocalPort(payload.GetParameter(i, "rule"), payload.GetParameter(i, "port"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "local_port_is_not" {
				result, data := windows.FirewallLocalPort(payload.GetParameter(i, "rule"), payload.GetParameter(i, "port"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "remote_port_is" {
				result, data := windows.FirewallRemotePort(payload.GetParameter(i, "rule"), payload.GetParameter(i, "port"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "remote_port_is_not" {
				result, data := windows.FirewallRemotePort(payload.GetParameter(i, "rule"), payload.GetParameter(i, "port"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "local_address_is" {
				result, data := windows.FirewallLocalAddress(payload.GetParameter(i, "rule"), payload.GetParameter(i, "address"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "local_address_is_not" {
				result, data := windows.FirewallLocalAddress(payload.GetParameter(i, "rule"), payload.GetParameter(i, "address"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "remote_address_is" {
				result, data := windows.FirewallRemoteAddress(payload.GetParameter(i, "rule"), payload.GetParameter(i, "address"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "remote_address_is_not" {
				result, data := windows.FirewallRemoteAddress(payload.GetParameter(i, "rule"), payload.GetParameter(i, "address"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "icmp_types_and_codes_is" {
				result, data := windows.FirewallICMP(payload.GetParameter(i, "rule"), payload.GetParameter(i, "value"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "icmp_types_and_codes_is_not" {
				result, data := windows.FirewallICMP(payload.GetParameter(i, "rule"), payload.GetParameter(i, "value"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "grouping_is" {
				result, data := windows.FirewallGrouping(payload.GetParameter(i, "rule"), payload.GetParameter(i, "group"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "grouping_is_not" {
				result, data := windows.FirewallGrouping(payload.GetParameter(i, "rule"), payload.GetParameter(i, "group"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "interface_type_is" {
				result, data := windows.FirewallInterface(payload.GetParameter(i, "rule"), payload.GetParameter(i, "interface"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "interface_type_is_not" {
				result, data := windows.FirewallInterface(payload.GetParameter(i, "rule"), payload.GetParameter(i, "interface"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "protocol_is" {
				result, data := windows.FirewallProtocol(payload.GetParameter(i, "rule"), payload.GetParameter(i, "protocol"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "protocol_is_not" {
				result, data := windows.FirewallProtocol(payload.GetParameter(i, "rule"), payload.GetParameter(i, "protocol"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "direction_is" {
				result, data := windows.FirewallDirection(payload.GetParameter(i, "rule"), payload.GetParameter(i, "direction"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "direction_is_not" {
				result, data := windows.FirewallDirection(payload.GetParameter(i, "rule"), payload.GetParameter(i, "direction"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "action_is" {
				result, data := windows.FirewallAction(payload.GetParameter(i, "rule"), payload.GetParameter(i, "action"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "action_is_not" {
				result, data := windows.FirewallAction(payload.GetParameter(i, "rule"), payload.GetParameter(i, "action"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "rule_profile_is" {
				result, data := windows.FirewallRuleProfile(payload.GetParameter(i, "rule"), payload.GetParameter(i, "profile"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "rule_profile_is_not" {
				result, data := windows.FirewallRuleProfile(payload.GetParameter(i, "rule"), payload.GetParameter(i, "profile"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "enabled_is" {
				result, data := windows.FirewallRuleEnabled(payload.GetParameter(i, "rule"), payload.GetParameter(i, "value"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "enabled_is_not" {
				result, data := windows.FirewallRuleEnabled(payload.GetParameter(i, "rule"), payload.GetParameter(i, "value"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "edge_traversal_is" {
				result, data := windows.FirewallEdge(payload.GetParameter(i, "rule"), payload.GetParameter(i, "traversal"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "edge_traversal_is_not" {
				result, data := windows.FirewallEdge(payload.GetParameter(i, "rule"), payload.GetParameter(i, "traversal"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end firewall
		} else if payload.GetSpace(i) == "group" {
			// groups rule implementation
			if payload.GetAction(i) == "group_exists" {
				result, data := windows.GroupExist(payload.GetParameter(i, "group"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "group_does_not_exist" {
				result, data := windows.GroupExist(payload.GetParameter(i, "group"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "comment_is" {
				result, data := windows.GroupComment(payload.GetParameter(i, "group"), payload.GetParameter(i, "comment"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "comment_is_not" {
				result, data := windows.GroupComment(payload.GetParameter(i, "group"), payload.GetParameter(i, "comment"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end groups
		} else if payload.GetSpace(i) == "process" {
			// process rule implementation
			if payload.GetAction(i) == "process_exists" {
				result, data := windows.ProcessExist(payload.GetParameter(i, "executable"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "process_does_not_exist" {
				result, data := windows.ProcessExist(payload.GetParameter(i, "executable"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "pid_is" {
				result, data := windows.ProcessPID(payload.GetParameter(i, "executable"), payload.GetParameter(i, "pid"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "pid_is_not" {
				result, data := windows.ProcessPID(payload.GetParameter(i, "executable"), payload.GetParameter(i, "pid"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "ppid_is" {
				result, data := windows.ProcessPPID(payload.GetParameter(i, "executable"), payload.GetParameter(i, "ppid"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "ppid_is_not" {
				result, data := windows.ProcessPPID(payload.GetParameter(i, "executable"), payload.GetParameter(i, "ppid"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "username_is" {
				result, data := windows.ProcessUsername(payload.GetParameter(i, "executable"), payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "username_is_not" {
				result, data := windows.ProcessUsername(payload.GetParameter(i, "executable"), payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end process
		} else if payload.GetSpace(i) == "profile" {
			// profile rule implementation
			if payload.GetAction(i) == "user_directory_is" {
				result, data := windows.ProfileUserDirectory(payload.GetParameter(i, "directory"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "user_directory_is_not" {
				result, data := windows.ProfileUserDirectory(payload.GetParameter(i, "directory"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "directory_is" {
				result, data := windows.ProfileDirectory(payload.GetParameter(i, "directory"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "directory_is_not" {
				result, data := windows.ProfileDirectory(payload.GetParameter(i, "directory"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end profile
		} else if payload.GetSpace(i) == "service" {
			// service rule implementation
			if payload.GetAction(i) == "service_exists" {
				result, data := windows.ServiceExist(payload.GetParameter(i, "service"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "service_does_not_exist" {
				result, data := windows.ServiceExist(payload.GetParameter(i, "service"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "display_name_is" {
				result, data := windows.ServiceDisplayName(payload.GetParameter(i, "service"), payload.GetParameter(i, "name"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "display_name_is_not" {
				result, data := windows.ServiceDisplayName(payload.GetParameter(i, "service"), payload.GetParameter(i, "name"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "status_text_is" {
				result, data := windows.ServiceStatusText(payload.GetParameter(i, "service"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "status_text_is_not" {
				result, data := windows.ServiceStatusText(payload.GetParameter(i, "service"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "status_is" {
				result, data := windows.ServiceStatus(payload.GetParameter(i, "service"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "status_is_not" {
				result, data := windows.ServiceStatus(payload.GetParameter(i, "service"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "accept_stop_is" {
				result, data := windows.ServiceAcceptStop(payload.GetParameter(i, "service"), payload.GetParameter(i, "stop"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "accept_stop_is_not" {
				result, data := windows.ServiceAcceptStop(payload.GetParameter(i, "service"), payload.GetParameter(i, "stop"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "is_running" {
				result, data := windows.ServiceRunning(payload.GetParameter(i, "service"), payload.GetParameter(i, "running"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "is_not_running" {
				result, data := windows.ServiceRunning(payload.GetParameter(i, "service"), payload.GetParameter(i, "running"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "running_pid_is" {
				result, data := windows.ServiceRunningPid(payload.GetParameter(i, "service"), payload.GetParameter(i, "pid"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "running_pid_is_not" {
				result, data := windows.ServiceRunningPid(payload.GetParameter(i, "service"), payload.GetParameter(i, "pid"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "service_type_is" {
				result, data := windows.ServiceType(payload.GetParameter(i, "service"), payload.GetParameter(i, "type"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "service_type_is_not" {
				result, data := windows.ServiceType(payload.GetParameter(i, "service"), payload.GetParameter(i, "type"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end service
		} else if payload.GetSpace(i) == "share" {
			// share rule implementation
			if payload.GetAction(i) == "share_exists" {
				result, data := windows.ShareExist(payload.GetParameter(i, "share"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "share_does_not_exist" {
				result, data := windows.ShareExist(payload.GetParameter(i, "share"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "status_is" {
				result, data := windows.ShareStatus(payload.GetParameter(i, "share"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "status_is_not" {
				result, data := windows.ShareStatus(payload.GetParameter(i, "share"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "caption_is" {
				result, data := windows.ShareCaption(payload.GetParameter(i, "share"), payload.GetParameter(i, "caption"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "caption_is_not" {
				result, data := windows.ShareCaption(payload.GetParameter(i, "share"), payload.GetParameter(i, "caption"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "description_is" {
				result, data := windows.ShareDescription(payload.GetParameter(i, "share"), payload.GetParameter(i, "description"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "description_is_not" {
				result, data := windows.ShareDescription(payload.GetParameter(i, "share"), payload.GetParameter(i, "description"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "path_is" {
				result, data := windows.SharePath(payload.GetParameter(i, "share"), payload.GetParameter(i, "path"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "path_is_not" {
				result, data := windows.SharePath(payload.GetParameter(i, "share"), payload.GetParameter(i, "path"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "allow_maximum_is" {
				result, data := windows.ShareAllowMaximum(payload.GetParameter(i, "share"), payload.GetParameter(i, "allow"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "allow_maximum_is_not" {
				result, data := windows.ShareAllowMaximum(payload.GetParameter(i, "share"), payload.GetParameter(i, "allow"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "type_is" {
				result, data := windows.ShareType(payload.GetParameter(i, "share"), payload.GetParameter(i, "type"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "type_is_not" {
				result, data := windows.ShareType(payload.GetParameter(i, "share"), payload.GetParameter(i, "type"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end share
		} else if payload.GetSpace(i) == "software" {
			// software rule implementation
			if payload.GetAction(i) == "software_exists" {
				result, data := windows.SoftwareExist(payload.GetParameter(i, "software"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "software_does_not_exist" {
				result, data := windows.SoftwareExist(payload.GetParameter(i, "software"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "arch_is" {
				result, data := windows.SoftwareArch(payload.GetParameter(i, "software"), payload.GetParameter(i, "arch"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "arch_is_not" {
				result, data := windows.SoftwareArch(payload.GetParameter(i, "software"), payload.GetParameter(i, "arch"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "publisher_is" {
				result, data := windows.SoftwarePublisher(payload.GetParameter(i, "software"), payload.GetParameter(i, "publisher"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "publisher_is_not" {
				result, data := windows.SoftwarePublisher(payload.GetParameter(i, "software"), payload.GetParameter(i, "publisher"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "install_date_is" {
				result, data := windows.SoftwareInstall(payload.GetParameter(i, "software"), payload.GetParameter(i, "date"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "install_date_is_not" {
				result, data := windows.SoftwareInstall(payload.GetParameter(i, "software"), payload.GetParameter(i, "date"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "estimated_size_is" {
				result, data := windows.SoftwareEstimatedSize(payload.GetParameter(i, "software"), payload.GetParameter(i, "size"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "estimated_size_is_not" {
				result, data := windows.SoftwareEstimatedSize(payload.GetParameter(i, "software"), payload.GetParameter(i, "size"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "contact_is" {
				result, data := windows.SoftwareContact(payload.GetParameter(i, "software"), payload.GetParameter(i, "contact"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "contact_is_not" {
				result, data := windows.SoftwareContact(payload.GetParameter(i, "software"), payload.GetParameter(i, "contact"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "helplink_is" {
				result, data := windows.SoftwareHelplink(payload.GetParameter(i, "software"), payload.GetParameter(i, "link"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "helplink_is_not" {
				result, data := windows.SoftwareHelplink(payload.GetParameter(i, "software"), payload.GetParameter(i, "link"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "install_source_is" {
				result, data := windows.SoftwareInstallSource(payload.GetParameter(i, "software"), payload.GetParameter(i, "source"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "install_source_is_not" {
				result, data := windows.SoftwareInstallSource(payload.GetParameter(i, "software"), payload.GetParameter(i, "source"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "install_location_is" {
				result, data := windows.SoftwareInstallLocation(payload.GetParameter(i, "software"), payload.GetParameter(i, "location"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "install_location_is_not" {
				result, data := windows.SoftwareInstallLocation(payload.GetParameter(i, "software"), payload.GetParameter(i, "location"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "version_major_is" {
				result, data := windows.SoftwareMajorVersion(payload.GetParameter(i, "software"), payload.GetParameter(i, "major"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "version_major_is_not" {
				result, data := windows.SoftwareMajorVersion(payload.GetParameter(i, "software"), payload.GetParameter(i, "major"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "version_minor_is" {
				result, data := windows.SoftwareMinorVersion(payload.GetParameter(i, "software"), payload.GetParameter(i, "minor"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "version_minor_is_not" {
				result, data := windows.SoftwareMinorVersion(payload.GetParameter(i, "software"), payload.GetParameter(i, "minor"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end software
		} else if payload.GetSpace(i) == "update" {
			// update rule implementation
			if payload.GetAction(i) == "update_completed" {
				result, data := windows.UpdateCompleted(payload.GetParameter(i, "update"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "update_pending" {
				result, data := windows.UpdatePending(payload.GetParameter(i, "update"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "in_history" {
				result, data := windows.UpdateHistory(payload.GetParameter(i, "update"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "not_in_history" {
				result, data := windows.UpdateHistory(payload.GetParameter(i, "update"), payload.GetParameter(i, "status"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end update
		} else if payload.GetSpace(i) == "user_windows" {
			// windows user rule implementation
			if payload.GetAction(i) == "user_exists" {
				result, data := windows.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "user_does_not_exist" {
				result, data := windows.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "user_logged_in" {
				result, data := windows.UserLoggedIn(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "user_not_logged_in" {
				result, data := windows.UserLoggedIn(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "bad_password_count_is" {
				result, data := windows.UserBadPassword(payload.GetParameter(i, "username"), payload.GetParameter(i, "count"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "bad_password_count_is_not" {
				result, data := windows.UserBadPassword(payload.GetParameter(i, "username"), payload.GetParameter(i, "count"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "fullname_is" {
				result, data := windows.UserFullName(payload.GetParameter(i, "username"), payload.GetParameter(i, "name"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "fullname_is_not" {
				result, data := windows.UserFullName(payload.GetParameter(i, "username"), payload.GetParameter(i, "name"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "is_admin" {
				result, data := windows.UserAdmin(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "is_not_admin" {
				result, data := windows.UserAdmin(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "is_enabled" {
				result, data := windows.UserEnabled(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "is_not_enabled" {
				result, data := windows.UserEnabled(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "is_locked" {
				result, data := windows.UserLocked(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "is_not_locked" {
				result, data := windows.UserLocked(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "last_logon_is" {
				result, data := windows.UserLastLogon(payload.GetParameter(i, "username"), payload.GetParameter(i, "date"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "last_logon_is_not" {
				result, data := windows.UserLastLogon(payload.GetParameter(i, "username"), payload.GetParameter(i, "date"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "no_change_password" {
				result, data := windows.UserNoChangePassword(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "password_changeable" {
				result, data := windows.UserPasswordChangeable(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "number_of_logons_is" {
				result, data := windows.UserNoOfLogons(payload.GetParameter(i, "username"), payload.GetParameter(i, "logons"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "number_of_logons_is_not" {
				result, data := windows.UserNoOfLogons(payload.GetParameter(i, "username"), payload.GetParameter(i, "logons"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "password_age_is" {
				result, data := windows.UserPasswordAge(payload.GetParameter(i, "username"), payload.GetParameter(i, "duration"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "password_age_is_not" {
				result, data := windows.UserPasswordAge(payload.GetParameter(i, "username"), payload.GetParameter(i, "duration"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "password_never_expires" {
				result, data := windows.UserPasswordExpires(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result, data) // debug print
				submission.Send(data, result, payload[i].ID, batch) // send score
			} else if payload.GetAction(i) == "password_expires" {
				result, data := windows.UserPasswordExpires(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result, data) // debug print
				submission.Send(data, !result, payload[i].ID, batch) // send score
			}
			// end windows user
		}
	}
	batch, err = uuid.NewUUID()
}