package piscine

func IsUpper(s string) bool {
	for _, elem := range s {
		if elem < 'A' || elem > 'Z' {
			return false
		}
	}
	return true
}
