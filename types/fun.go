package types

// MergeMapsString 多个map合并
func MergeMapsString(maps ...MapStrings) MapStrings {
	result := make(MapStrings)
	for _, m := range maps {
		for k, v := range m {
			result.Set(k, v)
		}
	}
	return result
}
