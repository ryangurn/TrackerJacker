package windows

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"strconv"
	"strings"
)

func checkString(k registry.Key, path string, key string, value interface{}) (retBool bool, err error) {
	retBool = false
	err = nil

	policy, err := registry.OpenKey(k, path, registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer policy.Close()

	s, _, err := policy.GetStringValue(key)
	if err != nil {
		return
	}

	if strings.Compare(s, value.(string)) == 0{
		retBool = true
		return
	}

	return
}

func checkStrings(k registry.Key, path string, key string, value interface{}) (retBool bool, err error) {
	retBool = false
	err = nil

	policy, err := registry.OpenKey(k, path, registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer policy.Close()

	s, _, err := policy.GetStringsValue(key)
	if err != nil {
		return
	}

	check := true
	for i := 0; i < len(s); i++ {
		if strings.Compare(value.([]string)[i], s[i]) != 0 {
			check = false
			break
		}
	}
	retBool = check

	return
}

func checkInteger64(k registry.Key, path string, key string, value interface{}) (retBool bool, err error) {
	retBool = false
	err = nil

	policy, err := registry.OpenKey(k, path, registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer policy.Close()

	i, _, err := policy.GetIntegerValue(key)
	if err != nil {
		return
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

func checkBinary(k registry.Key, path string, key string, value interface{}) (retBool bool, err error) {
	retBool = false
	err = nil

	policy, err := registry.OpenKey(k, path, registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer policy.Close()

	b, _, err := policy.GetBinaryValue(key)
	if err != nil {
		return
	}

	if string(b) == value.(string) {
		retBool = true
		return
	}

	return
}

func checkMUI(k registry.Key, path string, key string, value interface{}) (retBool bool, err error) {
	retBool = false
	err = nil

	policy, err := registry.OpenKey(k, path, registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer policy.Close()

	m, err := policy.GetMUIStringValue(key)
	if err != nil {
		return
	}

	if m == value.(string) {
		retBool = true
		return
	}

	return
}

func PolicyValue(k registry.Key, path string, key string, value interface{}) (retBool bool) {
	retBool = false

	// check string type
	// string
	strBool, err := checkString(k, path, key, value)
	if err != nil {

		strsBool, err := checkStrings(k, path, key, value)
		if err != nil {
			int64Bool, err := checkInteger64(k, path, key, value)
			if err != nil {
				binBool, err := checkBinary(k, path, key, value)
				if err != nil {
					MUIBool, err := checkMUI(k, path, key, value)
					if err != nil {
						return
					} else {
						return MUIBool
					}
				} else {
					return binBool
				}
			} else {
				return int64Bool
			}
		} else {
			return strsBool
		}
	} else {
		return strBool
	}
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
