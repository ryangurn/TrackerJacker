package main

import (
	"TrackerJacker/core/enc"
	"TrackerJacker/core/os/cross"
	"TrackerJacker/core/parsing"
	"TrackerJacker/core/submission"
	"fmt"
	"github.com/bugsnag/bugsnag-go"
	"github.com/joho/godotenv"
	"os"
)

const encKey = "Password"
const payload = "payload.txt"

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
	// todo: wait for the error to be resolved before starting
	// (ideally this will be when another program executes and creates the .env file)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file")
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
	checks := submission.GetPayload()
	// generate the payload
	generatePayload(checks)
	// get the payload
	payload := parsePayload()

	// loop through payload items
	for i := 0; i < len(payload); i++ {

		if payload.GetSpace(i) == "files" {
			// files rule implementation
			if payload.GetAction(i) == "exists" {
				// exists
				result, data := cross.FileExists(payload.GetParameter(i, "path"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID) // send score
			} else if payload.GetAction(i) == "does_not_exist" {
				// negate exists
				result, data := cross.FileExists(payload.GetParameter(i, "path"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID) // send score
			} else if payload.GetAction(i) == "hash" {
				str, err := cross.FileHash(payload.GetParameter(i, "path"))
				if err != nil {
					// hashing error
					payload.DebugPrint(i, false)
				} else {
					// no error
					result := str == payload.GetParameter(i, "hash")
					payload.DebugPrint(i, result) // debug print
					submission.Send(str, result, payload[i].ID) // send score
				}
			}
			// end files
		} else if payload.GetSpace(i) == "hosts" {
			// hosts rule implementation
			if payload.GetAction(i) == "ip_exists" {
				// ip address exists
				result, data := cross.HostIpExist(payload.GetParameter(i, "ip"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID) // send score
			} else if payload.GetAction(i) == "ip_does_not_exist" {
				// ip address does not exist
				result, data := cross.HostIpExist(payload.GetParameter(i, "ip"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID) // send score
			} else if payload.GetAction(i) == "host_exist" {
				// host exists
				result, data := cross.HostExist(payload.GetParameter(i, "host"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID) // send score
			} else if payload.GetAction(i) == "host_does_not_exist" {
				// host does not exist
				result, data := cross.HostExist(payload.GetParameter(i, "host"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID) // send score
			}
			// end hosts
		} else if payload.GetSpace(i) == "users" {
			// users rule implementation
			if payload.GetAction(i) == "exists" {
				result, data := cross.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result) // debug print
				submission.Send(data, result, payload[i].ID) // send score
			} else if payload.GetAction(i) == "does_not_exist" {
				result, data := cross.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, !result) // debug print
				submission.Send(data, !result, payload[i].ID) // send score
			}
			// end users
		}
	}
}