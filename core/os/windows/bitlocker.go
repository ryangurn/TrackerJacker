package windows

import (
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
)

func BitlockerDriveExist(drive string) (retBool bool) {
	retBool = false

	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
	if err != nil {
		return
	}

	if info.DriveLetter == drive {
		retBool = true
		return
	}

	return
}

func BitlockerDriveMeta(drive string, key string, value interface{}) (retBool bool) {
	retBool = false

	info, err := wapi.GetBitLockerRecoveryInfoForDrive(drive)
	if err != nil {
		 return
	}

	if info.DriveLetter == drive {
		if key == "DeviceID" {
			if info.DeviceID == value.(string) {
				retBool = true
				return
			}
		} else if key == "PersistentVolumeID" {
			if info.PersistentVolumeID == value.(string) {
				retBool = true
				return
			}
		} else if key == "ConversionStatus" {
			us, err := strconv.ParseInt(value.(string), 10, 32)
			if err != nil {
				return
			}
			if info.ConversionStatus == int32(us) {
				retBool = true
				return
			}
		} else if key == "ProtectionStatus" {
			us, err := strconv.ParseInt(value.(string), 10, 32)
			if err != nil {
				return
			}

			if info.ProtectionStatus == int32(us) {
				retBool = true
				return
			}
		}
	}

	return
}