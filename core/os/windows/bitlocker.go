package windows

import (
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	wapi "github.com/iamacarpet/go-win64api"
	"os"
	"strconv"
)

func BitlockerDriveLocked(drive string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
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
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	if info.ProtectionStatus == 2 {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func BitlockerDeviceID(drive string, id string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
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
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	if info.DeviceID == id {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func BitlockerPersistentVolumeID(drive string, id string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
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
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	if info.PersistentVolumeID == id {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func BitlockerConversionStatus(drive string, status string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
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
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	val, err := strconv.ParseUint(status, 10, 32)
	if err != nil {
		return
	}

	if info.ConversionStatus == uint32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}

func BitlockerProtectionStatus(drive string, status string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
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
		out, _ := json.Marshal(info)
		return retBool, string(out)
	}

	val, err := strconv.ParseUint(status, 10, 32)
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

	if info.ProtectionStatus == uint32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(info)
	return retBool, string(out)
}