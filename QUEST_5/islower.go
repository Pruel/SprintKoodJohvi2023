package piscine

func IsLower(s string) bool {
	for _, elem := range s {
		if elem < 'a' || elem > 'z' {
			return false
		}
	}
	return true
}
