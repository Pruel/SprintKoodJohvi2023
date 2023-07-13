package piscine

func AlphaCount(s string) int {
	count := 0

	for _, elem := range s {
		if (elem >= 'a' && 'z' >= elem) || (elem >= 'A' && 'Z' >= elem) {
			count++
		}
	}
	return count
}
