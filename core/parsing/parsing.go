package parsing

import (
	"encoding/json"
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

func ParsePayload(str []byte, payload *PayloadType) *PayloadType {
	json.Unmarshal([]byte(str), &payload)
	return payload
}