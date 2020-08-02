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

func GroupParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) < 2 {
		return
	}

	if args[0] == "exist" {
		if GroupExist(args[1]) == result {
			retBool = true
		}
	} else if args[0] == "meta" {
		if len(args) < 4 {
			return
		}

		if GroupMeta(args[1], args[2], args[3]) == result {
			retBool = true
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}