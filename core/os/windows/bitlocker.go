package windows

import (
	"fmt"
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

func BitlockerParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) < 2 {
		return
	}

	if args[0] == "drive_exist" {
		if BitlockerDriveExist(args[1]) == result {
			retBool = true
			return
		}
	} else if args[0] == "meta" {
		if len(args) != 4 {
			return
		}

		if BitlockerDriveMeta(args[1], args[2], args[3]) == result {
			retBool = true
			return
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}