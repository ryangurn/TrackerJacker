package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
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