package piscine

import "github.com/01-edu/z01"

var test bool = true

func PrintCombN(n int) {
	digits := make([]int, n)
	genCombinations(digits, 0, 0)
	z01.PrintRune('\n')
	test = true
}

func genCombinations(digits []int, currI int, currD int) {
	if currI == len(digits) {
		printOut(digits)
		return
	}

	for i := currD; i <= 9; i++ {
		digits[currI] = i
		genCombinations(digits, currI+1, i+1)
	}
}

func printOut(digits []int) {
	if test {
		test = false
	} else {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	for _, digit := range digits {
		z01.PrintRune(rune(int(digit) + '0'))
	}
}
