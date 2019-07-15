package wolf

import "regexp"

func validateWith(pattern, arg string) bool {
	match, err := regexp.Match(pattern, []byte(arg))
	if err != nil {
		return false
	}
	return match
}

func IsValidHex(hex string) bool {
	return validateWith("^[0-9a-f]+$", hex)
}

func IsValidWolf(wolf string) bool {
	return validateWith("^([a-z]{5,8}\\-)*([a-z]{5,8})$", wolf)
}
