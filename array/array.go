package array

import "log"

// ContainString determine whether the string element is in the array
func ContainString(array []string, item string) bool {
	if len(array) == 0 {
		return false
	}

	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

// HasDuplicateItem check has duplicate items
func HasDuplicateItem(array []string) bool {
	recordMap := make(map[string]bool)

	for _, item := range array {
		exists := recordMap[item]
		if exists {
			log.Printf("exists: %v", item)
			return true
		}
		recordMap[item] = true
	}
	return false
}
