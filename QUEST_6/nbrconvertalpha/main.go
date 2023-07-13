package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	charg := 'a'
	startIndex := 1
	args := os.Args

	// Проверяем, нужно ли печатать заглавные буквы
	if len(args) > 1 && args[1] == "--upper" {
		charg = 'A'
		startIndex++
	}

	// Если нет аргументов, то завершаем программу
	if len(args) <= startIndex {
		return
	}

	for _, arg := range args[startIndex:] {
		num := strToInt(arg)
		if num < 1 || num > 26 {
			// Если аргумент не является допустимым числом, печатаем пробел
			z01.PrintRune(' ')
		} else {
			// Конвертируем число в соответствующую букву алфавита и печатаем её
			z01.PrintRune(rune(charg) + rune(num-1))
		}
	}

	// Выводим новую строку после всех символов
	z01.PrintRune('\n')
}

// функция преобразования строки в число
func strToInt(str string) (num int) {
	for _, ch := range str {
		if ch < '0' || ch > '9' {
			return 0 // Вернуть 0, если символ не является числом
		}
		num = num*10 + int(ch-'0')
	}
	return num
}
