package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isAlphanumeric(s[i]) {
			return false
		}
	}
	return true
}

func isAlphanumeric(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z')
}
