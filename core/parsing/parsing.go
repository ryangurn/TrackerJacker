package parsing

import "encoding/json"

type PayloadType []struct {
	Namespace string   `json:"namespace"`
	Arguments []string `json:"arguments"`
	Result    interface{}     `json:"result"`
}

func ParsePayload(str []byte, payload *PayloadType) *PayloadType {
	json.Unmarshal([]byte(str), &payload)
	return payload
}