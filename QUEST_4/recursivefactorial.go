package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}

	result := 1
	switch nb {
	case 1:
		result *= 1 * 1
	case 2:
		result *= 1 * 2
	case 3:
		result *= 1 * 2 * 3
	case 4:
		result *= 1 * 2 * 3 * 4
	case 5:
		result *= 1 * 2 * 3 * 4 * 5
	case 6:
		result *= 1 * 2 * 3 * 4 * 5 * 6
	case 7:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7
	case 8:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8
	case 9:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9
	case 10:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10
	case 11:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11
	case 12:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12
	case 13:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13
	case 14:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14
	case 15:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15
	case 16:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16
	case 17:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16 * 17
	case 18:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18
	case 19:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18 * 19
	case 20:
		result *= 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18 * 19 * 20

	}
	return result
}
