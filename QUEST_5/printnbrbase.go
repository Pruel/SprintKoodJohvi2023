package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		PrintStr("NV")
		return
	}

	isNegative := nbr < 0
	if isNegative {
		z01.PrintRune('-')
		nbr = -nbr
	}

	length := len(base)
	digits := make([]int, 0)

	if nbr == 0 {
		digits = append(digits, 0)
	}

	for nbr > 0 {
		digit := nbr % length
		digits = append(digits, digit)
		nbr /= length
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(base[digits[i]]))
	}

	// For the special case of -9223372036854775808 in base 10
	if isNegative && len(digits) == 0 {
		PrintStr("9223372036854775808")
	}
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
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
