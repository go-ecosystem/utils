package array

// Contains determine whether the element is in the array
func Contains[V string | int64 | int32 | int | float64 | float32](array []V, item V) bool {
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
func HasDuplicateItem[V string | int64 | int32 | int | float64 | float32](values []V) (bool, V) {
	hash := make(map[V]bool)
	for _, v := range values {
		if !hash[v] {
			hash[v] = true
		} else {
			return true, v
		}
	}
	var v V
	return false, v
}

// RemoveDuplicateElements remove duplicate elements
func RemoveDuplicateElements[T comparable](values []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range values {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// Map map
func Map[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))
	for _, e := range data {
		res = append(res, f(e))
	}
	return res
}
