package util

import (
	"regexp"
)

func StringSliceToSet(slice []string) Set {
	if len(slice) == 0 {
		return nil
	}
	sliceSet := NewSet()
	for _, elem := range slice {
		sliceSet.Add(elem)
	}
	return sliceSet
}

// host必须满足基本规则 连续字符+.
func CheckHostValid(host string) bool {
	reg := regexp.MustCompile(`[\w-]+\.`)
	if reg.FindString(host) != "" {
		return true
	}
	return false
}
