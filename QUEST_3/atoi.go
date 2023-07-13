package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0 // Return 0 for an empty string
	}

	result := 0
	sign := 1
	startIndex := 0

	if s[0] == '-' {
		sign = -1
		startIndex = 1
	} else if s[0] == '+' {
		startIndex = 1
	}

	for i := startIndex; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0 // Return 0 if non-digit character is found
		}
		digit := int(ch - '0')
		result = result*10 + digit
	}

	return sign * result
}
