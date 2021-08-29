package windows

import (
	"encoding/json"
	"fmt"
	"github.com/ceshihao/windowsupdate"
	"github.com/go-ole/go-ole"
	wapi "github.com/iamacarpet/go-win64api"
	"github.com/scjalliance/comshim"
)

func getUpdateHistory() (data []*windowsupdate.IUpdateHistoryEntry) {
	comshim.Add(1)
	defer comshim.Done()

	// ole.CoInitialize(0)
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	defer ole.CoUninitialize()

	var err error

	session, err := windowsupdate.NewUpdateSession()
	if err != nil {
		panic(err)
	}

	// Query Update History
	fmt.Println("Step 1: Query Update History")
	searcher, err := session.CreateUpdateSearcher()
	if err != nil {
		panic(err)
	}

	result, err := searcher.QueryHistoryAll()
	if err != nil {
		panic(err)
	}

	return result
}

func UpdateCompleted(update string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	history := getUpdateHistory()
	for _, h := range history {
		if h.Title == update {
			retBool = true
			if out, err := json.Marshal(h); err == nil {
				return retBool, string(out)
			}
		}
	}
	return
}

func UpdatePending(update string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	history, err := wapi.UpdatesPending()
	if err != nil {
		return
	}

	for _, h := range history.UpdateHistory {
		if h.UpdateName == update {
			retBool = true
			if out, err := json.Marshal(h); err == nil {
				return retBool, string(out)
			}
		}
	}
	return
}

func UpdateHistory(update string, status string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	history, err := wapi.UpdatesPending()
	if err != nil {
		return
	}

	for _, h := range history.UpdateHistory {
		if h.UpdateName == update {
			if h.Status == status {
				retBool = true
				if out, err := json.Marshal(h); err == nil {
					return retBool, string(out)
				}
			}
		}
	}
	return
}