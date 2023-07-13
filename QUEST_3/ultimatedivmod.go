package piscine

func UltimateDivMod(a *int, b *int) {
	del := *a / *b
	ost := *a % *b
	*a = del
	*b = ost
}
