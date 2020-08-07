package windows

func PolicyValue(policy string, key string, value interface{}) (retBool bool) {
	retBool = false

	return
}

func PolicyMeta(policy string, key string, value interface{}) (retBool bool) {
	retBool = false

	return
}

func PolicyParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) != 4 {
		return
	}

	if args[0] == "value" {
		if PolicyValue(args[1], args[2], args[3]) == result {
			retBool = true
			return
		}
	} else if args[0] == "meta" {
		if PolicyMeta(args[1], args[2], args[3]) == result {
			retBool = true
			return
		}
	}
	
	return
}
