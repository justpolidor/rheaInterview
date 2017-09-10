package milestones

const HUNDRED = 7
const HUNDREDAND = 10
const ONETHOUSAND = 11

func CharsCount(and bool) int {
	sum := 0
	for i := 1; i < 1000; i++ {
		sum += computeNumber(i, and)
	}
	return sum + ONETHOUSAND
}

func computeNumber(number int, and bool) int {

	// case : the number we are operating on - return : the length of that number (number of letters)
	switch number {
	case 0:
		return 0
	case 1,2,6,10:
		return 3
	case 3, 7, 8, 40, 50, 60:
		return 5
	case 4, 5, 9:
		return 4
	case 11, 12, 20, 30, 80, 90:
		return 6
	case 13, 14, 18, 19:
		return 8
	case 15, 16, 70:
		return 7
	case 17:
		return 9
	}
	if number < 100 {
		//calculate the dozens part + the units part
		return computeNumber(number-(number%10), and) + computeNumber(number%10, and)
	}
	// 100, 200, 300... and so on
	if number%100 == 0 {
		return computeNumber(number/100, and) + HUNDRED
	}
	// we get the hundred part + the dozens part (calling the first "if" to get it)
	return computeNumber(number/100, and) + separator(and) + computeNumber(number%100, and)
}

func separator(and bool) int {
	if and {
		return HUNDREDAND
	}
	return HUNDRED
}