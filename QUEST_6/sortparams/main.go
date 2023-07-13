package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := 0; i < len(arg)-1; i++ {
		for j := 1; j < len(arg)-i-1; j++ {
			if arg[j] > arg[j+1] {
				arg[j], arg[j+1] = arg[j+1], arg[j]
			}
		}
	}
	for index := 1; index < len(arg); index++ {
		name := arg[index]
		for _, str := range name {
			z01.PrintRune(str)
		}
		z01.PrintRune('\n')
	}
}
