package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	start := 0

	for i, char := range s {
		if char == '-' && i == start {
			sign = -1
		} else if char == '+' && i == start {
			sign = 1
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else if i == start {
			start++
		}
	}

	return result * sign
}
