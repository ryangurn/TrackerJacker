package cross

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	"io"
	"os"
)


func FileExists(path string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	if data, err := os.Stat(path); err == nil {
		retBool = true
		// marshal data
		if out, err := json.Marshal(data); err == nil {
			return retBool, string(out)
		}
	}
	return
}

func FileHash(path string) (string, error) {
	var h string

	file, err := os.Open(path)
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
		return h, err
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
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
		return h, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	h = hex.EncodeToString(hashInBytes)

	return h, nil
}