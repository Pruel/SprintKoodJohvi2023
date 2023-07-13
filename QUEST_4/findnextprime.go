package piscine

func IsPrime2(nb int) bool {
	if nb <= 1 {
		return false
	}

	for i := 2; i*i <= nb; i++ {
		result := nb % i
		if result == 0 {
			return false
		}

	}
	return true
}

func FindNextPrime(nb int) int {
	if nb == 2 {
		return 2
	}

	for i := nb; ; i++ {
		if IsPrime2(i) {
			return i
		}
	}
}

// Напиши функцию которая вернет prime число которое равно или превосходит значение  int переданного в качестве аргумента.
// prime - Это число которое делится на 1 и на себя.
