package numparse

import (
	"strconv"
	"strings"
)

// SNumParse parses the string in num based on any prefix I can think of, and returns an int64 representation
func SNumParse(num string) (int64, error) {
	strip, base := BasePrefix(num)
	return strconv.ParseInt(strip, base, 64)
}

// UNumParse parses the string in num based on any prefix I can think of, and returns a uint64 representation
func UNumParse(num string) (uint64, error) {
	strip, base := BasePrefix(num)
	return strconv.ParseUint(strip, base, 64)
}

// BasePrefix returns the stripped number and numerical base implied by the prefix
func BasePrefix(num string) (string, int) {
	if strings.HasPrefix(num, "0x") {
		return strings.TrimPrefix(num, "0x"), 16
	} else if strings.HasPrefix(num, "0o") {
		return strings.TrimPrefix(num, "0o"), 8
	} else if strings.HasPrefix(num, "0b") {
		return strings.TrimPrefix(num, "0b"), 2
	} else if strings.HasPrefix(num, "&h") {
		return strings.TrimPrefix(num, "&h"), 16
	} else if strings.HasPrefix(num, "$") {
		return strings.TrimPrefix(num, "$"), 16
	} else if strings.HasPrefix(num, "#") {
		return strings.TrimPrefix(num, "#"), 16
	} else if strings.HasPrefix(num, "&") {
		return strings.TrimPrefix(num, "&"), 8
	} else if strings.HasPrefix(num, "0") {
		return strings.TrimPrefix(num, "0"), 8
	} else if strings.HasPrefix(num, "%") {
		return strings.TrimPrefix(num, "%"), 2
	}
	return num, 10
}
