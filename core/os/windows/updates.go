package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
)

func UpdateHistoryExists(update string, status string) (retBool bool) {
	retBool = false

	updates, err := wapi.UpdatesPending()
	if err != nil {
		return
	}

	for _, u := range updates.UpdateHistory {
		if u.UpdateName == update && u.Status == status {
			retBool = true
			return
		}
	}

	return
}

func Updated(completed bool) (retBool bool) {
	retBool = false

	updates, err := wapi.UpdatesPending()
	if err != nil {
		return
	}

	fmt.Println(!updates.UpdatesReq, completed)
	if !updates.UpdatesReq == completed {
		retBool = true
		return
	}

	return
}

func UpdateParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) < 2 {
		return
	}

	if args[0] == "history_exist" {
		if len(args) != 3 {
			return
		}

		if UpdateHistoryExists(args[1], args[2]) == result {
			retBool = true
			return
		}
	} else if args[0] == "updated" {
		if len(args) != 2 {
			return
		}

		b, err := strconv.ParseBool(args[1])
		if err != nil {
			return
		}

		if Updated(b) == result {
			 retBool = true
			 return
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}