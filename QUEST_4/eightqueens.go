package piscine

import "github.com/01-edu/z01"

const n = 8

var board = [n]int{}

func EightQueens() {
	solve(0)
}

func underAttack(x, y int) bool {
	for i := 0; i < y; i++ {
		if board[i] == x || board[i]-x == i-y || board[i]-x == y-i {
			return true
		}
	}
	return false
}

func solve(y int) {
	if y == n {
		printBoard()
		return
	}
	for x := 0; x < n; x++ {
		if !underAttack(x, y) {
			board[y] = x
			solve(y + 1)
			board[y] = 0
		}
	}
}

func printBoard() {
	for _, pos := range board {
		z01.PrintRune(rune(pos + 49))
	}
	z01.PrintRune('\n')
}
