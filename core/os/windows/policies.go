package windows

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"strconv"
)

func checkString() (retBool bool, err error) {
	retBool = false
	err = nil


	return
}

func checkStrings() (retBool bool, err error) {
	retBool = false
	err = nil


	return
}

func checkInteger32() (retBool bool, err error) {
	retBool = false
	err = nil


	return
}

func getInteger64() (retBool bool, err error) {
	retBool = false
	err = nil


	return
}

func checkBinary() (retBool bool, err error) {
	retBool = false
	err = nil


	return
}

func checkMUI() (retBool bool, err error) {
	retBool = false
	err = nil


	return
}

func PolicyValue(k registry.Key, path string, key string, value interface{}) (retBool bool) {
	retBool = false

	policy, err := registry.OpenKey(k, path, registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer k.Close()

	// check string type

	s, _, err := policy.GetStringValue(key)
	if err != nil {
		// check int type
		i, _, err := policy.GetIntegerValue(key)
		if err != nil {
			fmt.Println("err2", err)
		}

		ui, err := strconv.ParseUint(value.(string), 10, 32)
		if err != nil {
			return
		}

		if i == ui {
			retBool = true
			return
		}

		return
	}

	if s == value.(string) {
		retBool = true
		return
	}

	return
}

func PolicyParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) != 5 {
		return
	}

	if args[0] == "value" {
		var key registry.Key
		switch args[1] {
		case "CLASSES_ROOT":
			key = registry.CLASSES_ROOT
			break
		case "CURRENT_USER":
			key = registry.CURRENT_USER
			break
		case "LOCAL_MACHINE":
			key = registry.LOCAL_MACHINE
			break
		case "USERS":
			key = registry.USERS
			break
		case "CURRENT_CONFIG":
			key = registry.CURRENT_CONFIG
			break
		case "PERFORMANCE_DATA":
			key = registry.PERFORMANCE_DATA
			break
		default:
			return false
			break
		}
		
		if PolicyValue(key, args[2], args[3], args[4]) == result {
			retBool = true
			return
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}
	
	return
}
