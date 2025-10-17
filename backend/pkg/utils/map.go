package utils

// MergeStringMaps merges multiple string maps (later values override earlier ones)
func MergeStringMaps(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	
	return result
}

// GetMapKeys returns all keys from a string map
func GetMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GetMapValues returns all values from a map
func GetMapValues(m map[string]interface{}) []interface{} {
	values := make([]interface{}, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// MapHasKey checks if a map has a specific key
func MapHasKey(m map[string]interface{}, key string) bool {
	_, exists := m[key]
	return exists
}

// FilterMap filters map entries based on a predicate function
func FilterMap(m map[string]interface{}, predicate func(string, interface{}) bool) map[string]interface{} {
	result := make(map[string]interface{})
	
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	
	return result
}
