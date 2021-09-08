package windows

import (
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	wapi "github.com/iamacarpet/go-win64api"
	"os"
)

func ProfileUserDirectory(dir string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	profile, err := wapi.GetDefaultUserProfileDirectory()
	if err != nil{
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	if profile == dir {
		retBool = true
		if out, err := json.Marshal(profile); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(profile)
	return retBool, string(out)
}

func ProfileDirectory(dir string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	profile, err := wapi.GetProfilesDirectory()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	if profile == dir {
		retBool = true
		if out, err := json.Marshal(profile); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(profile)
	return retBool, string(out)
}
