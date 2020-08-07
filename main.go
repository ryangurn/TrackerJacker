package main

import (
	"TrackerJacker/core/enc"
	"TrackerJacker/core/os/cross"
	"TrackerJacker/core/os/windows"
	"TrackerJacker/core/parsing"
	"TrackerJacker/core/web"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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

func main() {
	// setup the web directories
	if web.Setup() != true {
		return
	}

	// generate the payload
	generatePayload("input.json")
	// get the payload
	payload := parsePayload()

	// loop through payload items
	for i := 0; i < len(payload); i++ {
		if payload[i].Namespace == "files" {
			out := cross.FileParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "firewalls" {
			out := windows.FirewallParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "hosts" {
			out := cross.HostParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "users" {
			out := windows.UserParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "groups" {
			out := windows.GroupParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "softwares" {
			out := windows.SoftwareParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "services" {
			out := windows.ServiceParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "shares" {
			out := windows.ShareParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "processes" {
			out := windows.ProcessParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "policies" {
			out := windows.PolicyParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else {
			fmt.Printf("Unrecognized Namespace: %s\n", payload[i].Namespace)
		}
	}

	web.Clean()
}