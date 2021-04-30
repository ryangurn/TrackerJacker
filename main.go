package main

import (
	"TrackerJacker/core/enc"
	"TrackerJacker/core/os/cross"
	"TrackerJacker/core/parsing"
	"fmt"
	"github.com/bugsnag/bugsnag-go"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
)

const encKey = "Password"
const payload = "payload.txt"

func generatePayload(inFile string) {
	data, _ := ioutil.ReadFile(inFile)
	enc.EncryptFile(payload, data, encKey)
}

func parsePayload() parsing.PayloadType {
	str := enc.DecryptFile(payload, encKey)

	var payload parsing.PayloadType
	parsing.ParsePayload(str, &payload)
	return payload
}

func printDebug(space string, id int, result bool) {
	fmt.Printf("Space: %s | ID: %d | Output: %t\n", space, id, result)
}

func main() {
	// todo: wait for the error to be resolved before starting
	// (ideally this will be when another program executes and creates the .env file)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file")
	}

	bugsnagKey := os.Getenv("BUGSNAG_KEY")

	// setup bugsnag
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          bugsnagKey,
		ReleaseStage:    "alpha",
		AppVersion: 	"0.0.1",
		ProjectPackages: []string{"main"},
	})

	// generate the payload
	generatePayload("input.json")
	// get the payload
	payload := parsePayload()

	// loop through payload items
	for i := 0; i < len(payload); i++ {

		if payload.GetSpace(i) == "files" {
			// files rule implementation
			if payload.GetAction(i) == "exists" {
				// exists
				result := cross.FileExists(payload.GetParameter(i, "path"))
				payload.DebugPrint(i, result)
			} else if payload.GetAction(i) == "does_not_exist" {
				// negate exists
				result := !cross.FileExists(payload.GetParameter(i, "path"))
				payload.DebugPrint(i, result)
			} else if payload.GetAction(i) == "hash" {
				str, err := cross.FileHash(payload.GetParameter(i, "path"))
				if err != nil {
					// hashing error
					payload.DebugPrint(i, false)
				} else {
					// no error
					result := str == payload.GetParameter(i, "hash")
					payload.DebugPrint(i, result)
				}
			}
			// end files
		} else if payload.GetSpace(i) == "hosts" {
			// hosts rule implementation
			if payload.GetAction(i) == "ip_exists" {
				// ip address exists
				result := cross.HostIpExist(payload.GetParameter(i, "ip"))
				payload.DebugPrint(i, result)
			} else if payload.GetAction(i) == "ip_does_not_exist" {
				// ip address does not exist
				result := !cross.HostIpExist(payload.GetParameter(i, "ip"))
				payload.DebugPrint(i, result)
			} else if payload.GetAction(i) == "host_exist" {
				// host exists
				result := cross.HostExist(payload.GetParameter(i, "host"))
				payload.DebugPrint(i, result)
			} else if payload.GetAction(i) == "host_does_not_exist" {
				// host does not exist
				result := !cross.HostExist(payload.GetParameter(i, "host"))
				payload.DebugPrint(i, result)
			}
			// end hosts
		} else if payload.GetSpace(i) == "users" {
			// users rule implementation
			if payload.GetAction(i) == "exists" {
				result := cross.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result)
			} else if payload.GetAction(i) == "does_not_exist" {
				result := !cross.UserExist(payload.GetParameter(i, "username"))
				payload.DebugPrint(i, result)
			}
			// end users
		}
	}
}