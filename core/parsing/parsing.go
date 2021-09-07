package parsing

import (
	"encoding/json"
	"fmt"
	"time"
)

type PayloadType []struct {
	ID         int    				`json:"id"`
	Image      int    				`json:"image"`
	Rule       int   				`json:"rule"`
	RuleAction string 				`json:"rule_action"`
	Parameters interface{} 			`json:"parameters,omitempty"`
	Score     float64   			`json:"score"`
	Comment   string    			`json:"comment"`
	CreatedAt time.Time 			`json:"created_at"`
	UpdatedAt time.Time 			`json:"updated_at"`
	GetRule   struct {
		ID         int      		`json:"id"`
		Name       string   		`json:"name"`
		Space      string   		`json:"space"`
		Actions    []string 		`json:"actions"`
		Parameters struct {
			Exists struct {
				Path string 		`json:"path"`
			} 						`json:"exists"`
			DoesNotExist struct {
				Path string 		`json:"path"`
			} 						`json:"does_not_exist"`
			Hash struct {
				Path     string 	`json:"path"`
				Expected string 	`json:"expected"`
			} 						`json:"hash"`
		} 							`json:"parameters"`
		Description string    		`json:"description"`
		CreatedAt   time.Time 		`json:"created_at"`
		UpdatedAt   time.Time 		`json:"updated_at"`
	} 								`json:"get_rule"`
}

func (payload PayloadType) GetSpace(id int) string {
	return payload[id].GetRule.Space
}

func (payload PayloadType) GetAction(id int) string {
	return payload[id].RuleAction
}

func (payload PayloadType) GetParameter(id int, key string) string {
	val, ok := payload[id].Parameters.(map[string]interface{})[key]
	if ok {
		return val.(string)
	}
	return ""
}

func (payload PayloadType) DebugPrint(id int, result bool, data string) {
	fmt.Printf("Space: %s | Action: %s | ID: %d | Output: %t | Data %s\n", payload.GetSpace(id), payload.GetAction(id), payload[id].ID, result, data)
}

func ParsePayload(str []byte, payload *PayloadType) *PayloadType {
	json.Unmarshal([]byte(str), &payload)
	return payload
}

type InitialAuthorization struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     struct {
		ID              int         `json:"id"`
		OperatingSystem int         `json:"operating_system"`
		Name            string      `json:"name"`
		Creator         interface{} `json:"creator"`
		Status          int         `json:"status"`
		Method          int         `json:"method"`
		CreatedAt       time.Time   `json:"created_at"`
		UpdatedAt       time.Time   `json:"updated_at"`
	} `json:"image"`
}

func (auth InitialAuthorization) GetID() int {
	return auth.Image.ID
}

func (auth InitialAuthorization) GetMethod() int {
	return auth.Image.Method
}