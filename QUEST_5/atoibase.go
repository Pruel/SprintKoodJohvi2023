package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	result := 0
	baseLen := len(base)

	for i := len(s) - 1; i >= 0; i-- {
		validDigit := false
		for j, b := range base {
			if rune(s[i]) == b {
				result += j * pow(baseLen, len(s)-1-i)
				validDigit = true
				break
			}
		}
		if !validDigit {
			return 0
		}
	}

	return result
}

func pow(base int, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '-' || ch == '+' {
			return false
		}

		if seen[ch] {
			return false
		}

		seen[ch] = true
	}

	return true
}
