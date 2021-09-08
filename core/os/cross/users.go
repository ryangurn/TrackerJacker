package cross

import (
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	"os"
	"os/user"
)

// TODO: Add support for meta and logged in checks

func UserExist(usr string) (retBool bool, retData string) {
	retBool = false
	u, err := user.Lookup(usr)

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

	if u != nil {
		retBool = true
		if out, err := json.Marshal(u); err == nil {
			retData = string(out)
		}
	}

	return
}