package piscine

func ToUpper(s string) string {
	var n string
	for i := 0; i < len(s); i++ {
		a := Changer(s[i])
		n += a
	}
	return n
}

func Changer(ch byte) string {
	if ch >= 'a' && ch <= 'z' {
		return string(ch - 32)
	}
	return string(ch)
}
