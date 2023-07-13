package piscine

func NRune(s string, n int) rune {
	if 1 > n || n > len(s) {
		return 0
	}

	return []rune(s)[n-1]
}
