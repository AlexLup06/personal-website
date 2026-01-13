package helpers

func Intersect[T comparable](arr1, arr2 []T) []T {
	set := make(map[T]bool)
	var result []T

	// Store elements of arr1 in a set
	for _, num := range arr1 {
		set[num] = true
	}

	// Check for common elements in arr2
	for _, num := range arr2 {
		if set[num] {
			result = append(result, num)
			// Remove to avoid duplicates
			delete(set, num)
		}
	}

	return result
}
