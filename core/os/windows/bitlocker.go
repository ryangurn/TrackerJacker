package windows

import (
	"encoding/json"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
)

func BitlockerDriveLocked(drive string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
	if err != nil {
		return
	}

	if info.ProtectionStatus == 2 {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func BitlockerDeviceID(drive string, id string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
	if err != nil {
		return
	}

	if info.DeviceID == id {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func BitlockerPersistentVolumeID(drive string, id string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
	if err != nil {
		return
	}

	if info.PersistentVolumeID == id {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}

	return
}

func BitlockerConversionStatus(drive string, status string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
	if err != nil {
		return
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

	return
}

func BitlockerProtectionStatus(drive string, status string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
	if err != nil {
		return
	}

	val, err := strconv.ParseUint(status, 10, 32)
	if err != nil {
		return
	}

	if info.ProtectionStatus == uint32(val) {
		retBool = true
		if out, err := json.Marshal(info); err == nil {
			return retBool, string(out)
		}
	}
	return
}