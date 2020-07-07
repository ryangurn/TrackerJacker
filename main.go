package main

import (
	"TrackerJacker/core/enc"
	"TrackerJacker/core/os/cross"
	"TrackerJacker/core/os/windows"
	"TrackerJacker/core/parsing"
	"fmt"
	"io/ioutil"
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
	generatePayload("input.json")
	payload := parsePayload()
	for i := 0; i < len(payload); i++ {
		if payload[i].Namespace == "files" {
			out := cross.FileParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "hosts" {
			out := cross.HostParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else if payload[i].Namespace == "users" {
			out := windows.UserParse(payload[i].Arguments, payload[i].Result)
			fmt.Printf("Namespace %s | Command %s | Output: %t\n", payload[i].Namespace, payload[i].Arguments, out)
		} else {
			fmt.Printf("Unrecognized Namespace: %s\n", payload[i].Namespace)
		}
	}
}