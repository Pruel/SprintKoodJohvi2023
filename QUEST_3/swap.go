package piscine

func Swap(a *int, b *int) {
	swap2 := *a
	swap1 := *b
	*a = swap1
	*b = swap2
}
