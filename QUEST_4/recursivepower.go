package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	result := 1

	switch power {
	case 1:
		result *= nb
	case 2:
		result *= nb * nb
	case 3:
		result *= nb * nb * nb
	case 4:
		result *= nb * nb * nb * nb
	case 5:
		result *= nb * nb * nb * nb * nb
	case 6:
		result *= nb * nb * nb * nb * nb * nb
	case 7:
		result *= nb * nb * nb * nb * nb * nb * nb
	case 8:
		result *= nb * nb * nb * nb * nb * nb * nb * nb
	case 9:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 10:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 11:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 12:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 13:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 14:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 15:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 16:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 17:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 18:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 19:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	case 20:
		result *= nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb * nb
	}

	return result
}
