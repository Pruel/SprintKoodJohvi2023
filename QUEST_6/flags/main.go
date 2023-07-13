package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func bubbleSort(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		printHelp()
		return
	}

	insertFlag := "--insert="
	iFlag := "-i="

	var (
		strToInsert string
		strToSort   string
		order       bool
	)

	for _, arg := range args {
		if len(arg) >= len(insertFlag) && arg[:len(insertFlag)] == insertFlag {
			strToInsert = arg[len(insertFlag):]
		} else if len(arg) >= len(iFlag) && arg[:len(iFlag)] == iFlag {
			strToInsert = arg[len(iFlag):]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			strToSort = arg
		}
	}

	strToSort += strToInsert

	if order {
		strToSort = bubbleSort(strToSort)
	}

	fmt.Println(strToSort)
}
