package utils

// Snake convert camel case to snake case
func Snake(name string) string {
	var result string
	for i, v := range name {
		if v >= 'A' && v <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}
	return result
}
