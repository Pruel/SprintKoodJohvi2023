package piscine

func StrLen(s string) int {
	var len int = len([]rune(s))
	return len
}
