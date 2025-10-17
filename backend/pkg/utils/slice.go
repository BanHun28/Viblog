package utils

// ContainsString checks if a string slice contains a specific string
func ContainsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ContainsUint checks if a uint slice contains a specific uint
func ContainsUint(slice []uint, item uint) bool {
	for _, u := range slice {
		if u == item {
			return true
		}
	}
	return false
}

// UniqueStrings removes duplicate strings from a slice
func UniqueStrings(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	
	return result
}

// UniqueUints removes duplicate uints from a slice
func UniqueUints(slice []uint) []uint {
	seen := make(map[uint]bool)
	result := []uint{}
	
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	
	return result
}

// RemoveString removes a string from a slice
func RemoveString(slice []string, item string) []string {
	result := []string{}
	for _, s := range slice {
		if s != item {
			result = append(result, s)
		}
	}
	return result
}

// RemoveUint removes a uint from a slice
func RemoveUint(slice []uint, item uint) []uint {
	result := []uint{}
	for _, u := range slice {
		if u != item {
			result = append(result, u)
		}
	}
	return result
}

// ChunkStrings splits a string slice into chunks of specified size
func ChunkStrings(slice []string, chunkSize int) [][]string {
	if chunkSize <= 0 {
		return [][]string{slice}
	}
	
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	
	return chunks
}

// FilterStrings filters strings based on a predicate function
func FilterStrings(slice []string, predicate func(string) bool) []string {
	result := []string{}
	for _, s := range slice {
		if predicate(s) {
			result = append(result, s)
		}
	}
	return result
}

// MapStrings transforms strings based on a mapper function
func MapStrings(slice []string, mapper func(string) string) []string {
	result := make([]string, len(slice))
	for i, s := range slice {
		result[i] = mapper(s)
	}
	return result
}
