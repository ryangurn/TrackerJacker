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
		}
	}
	batch, err = uuid.NewUUID()
}