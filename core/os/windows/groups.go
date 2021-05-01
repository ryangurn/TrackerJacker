package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
	"strings"
)

func GroupExist(grp string) (retBool bool) {
	retBool = false

	groups, err := wapi.ListLocalGroups()
	if err != nil {
		fmt.Printf("Error fetching group list, %s.\r\n", err.Error())
		return
	}

	for _, g := range groups {
		if strings.TrimSpace(strings.ToLower(grp)) == strings.TrimSpace(strings.ToLower(g.Name)) {
			retBool = true
			return
		}
	}

	return
}

func GroupMeta(grp string, key string, value interface{}) (retBool bool) {
	retBool = false

	groups, err := wapi.ListLocalGroups()
	if err != nil {
		fmt.Printf("Error fetching group list, %s.\r\n", err.Error())
		return
	}

	for _, g := range groups {
		if strings.TrimSpace(strings.ToLower(grp)) == strings.TrimSpace(strings.ToLower(g.Name)) {
			if key == "Comment" {
				if g.Comment == value.(string) {
					retBool = true
					return
				} else {
					return
				}
			}
		}
	}

	return
}