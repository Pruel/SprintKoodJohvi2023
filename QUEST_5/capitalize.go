package piscine

func toUpper(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		return ch - 'a' + 'A'
	}
	return ch
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch - 'A' + 'a'
	}
	return ch
}

func Capitalize(s string) string {
	words := make([]byte, len(s))
	wordStart := true
	for i := 0; i < len(s); i++ {
		if wordStart && isAlphanumeric(s[i]) {
			words[i] = toUpper(s[i])
			wordStart = false
		} else {
			words[i] = toLower(s[i])
		}
		if !isAlphanumeric(s[i]) {
			wordStart = true
		}
	}
	return string(words)
}
