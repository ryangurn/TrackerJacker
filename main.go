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

func getParameters(parameters interface{}, key string) interface{} {
	val, ok := parameters.(map[string]interface{})[key]
	if ok {
		return val
	}
	return nil
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

		if payload[i].GetRule.Space == "files" {
			// files rule implementation
			if payload[i].RuleAction == "exists" {
				// exists
				result := cross.FileExists(getParameters(payload[i].Parameters, "path").(string))
				printDebug(payload[i].GetRule.Space, payload[i].ID, result)
			} else if payload[i].RuleAction == "does_not_exist" {
				// negate exists
				result := !cross.FileExists(getParameters(payload[i].Parameters, "path").(string))
				printDebug(payload[i].GetRule.Space, payload[i].ID, result)
			} else if payload[i].RuleAction == "hash" {
				str, err := cross.FileHash(getParameters(payload[i].Parameters, "path").(string))
				if err != nil {
					// hashing error
					printDebug(payload[i].GetRule.Space, payload[i].ID, false)
				} else {
					// no error
					result := str == getParameters(payload[i].Parameters, "hash").(string)
					printDebug(payload[i].GetRule.Space, payload[i].ID, result)
				}
			}
			// end files
		} else if payload[i].GetRule.Space == "hosts" {
			// hosts rule implementation
			if payload[i].RuleAction == "ip_exists" {
				// ip address exists
				result := cross.HostIpExist(getParameters(payload[i].Parameters, "ip").(string))
				printDebug(payload[i].GetRule.Space, payload[i].ID, result)
			} else if payload[i].RuleAction == "ip_does_not_exist" {
				// ip address does not exist
				result := !cross.HostIpExist(getParameters(payload[i].Parameters, "ip").(string))
				printDebug(payload[i].GetRule.Space, payload[i].ID, result)
			} else if payload[i].RuleAction == "host_exist" {
				// host exists
				result := cross.HostExist(getParameters(payload[i].Parameters, "ip").(string))
				printDebug(payload[i].GetRule.Space, payload[i].ID, result)
			} else if payload[i].RuleAction == "host_does_not_exist" {
				// host does not exist
				result := !cross.HostExist(getParameters(payload[i].Parameters, "ip").(string))
				printDebug(payload[i].GetRule.Space, payload[i].ID, result)
			}
			// end hosts
		} else if payload[i].GetRule.Space == "users" {
			// users rule implementation
			if payload[i].RuleAction == "exists" {
				result := cross.UserExist(getParameters(payload[i].Parameters, "username").(string))
				printDebug(payload[i].GetRule.Space, payload[i].ID, result)
			} else if payload[i].RuleAction == "does_not_exist" {
				result := !cross.UserExist(getParameters(payload[i].Parameters, "username").(string))
				printDebug(payload[i].GetRule.Space, payload[i].ID, result)
			}
			// end users
		}
	}
}