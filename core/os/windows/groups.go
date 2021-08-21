package windows

import (
	"encoding/json"
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
)

func GroupExist(group string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	groups, err := wapi.ListLocalGroups()
	if err != nil {
		return
	}

	for _, g := range groups {
		if group == g.Name {
			retBool = true
			if out, err := json.Marshal(g); err == nil {
				return retBool, string(out)
			}
			return
		}
	}

	return
}

func GroupComment(group string, comment string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	groups, err := wapi.ListLocalGroups()
	if err != nil {
		fmt.Printf("Error fetching group list, %s.\r\n", err.Error())
		return
	}

	for _, g := range groups {
		if group == g.Name {
			if g.Comment == comment {
				retBool = true
				if out, err := json.Marshal(g); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}