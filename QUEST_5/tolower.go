package piscine

func ToLower(s string) string {
	var n string
	for i := 0; i < len(s); i++ {
		a := ChangerLower(s[i])
		n += a
	}
	return n
}

func ChangerLower(ch byte) string {
	if ch >= 'A' && ch <= 'Z' {
		return string(ch + 32)
	}
	return string(ch)
}
