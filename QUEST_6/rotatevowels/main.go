package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	vowels := []rune{}
	for _, arg := range args {
		for _, r := range arg {
			if isVowel(r) {
				vowels = append(vowels, r)
			}
		}
	}

	j := len(vowels) - 1
	for i, arg := range args {
		for _, r := range arg {
			if isVowel(r) {
				z01.PrintRune(vowels[j])
				j--
			} else {
				z01.PrintRune(r)
			}
		}
		if i != len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
