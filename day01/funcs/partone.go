package funcs

// -----PART ONE-----
func SumNumbers(digits []int) int {
	var result int
	for _, num := range digits {
		result += num
	}
	return result
}

func PartOne(fileLines []string) []int {
	var digits []int
	var resArr []int
	var twoDigit int
	for _, line := range fileLines {
		for _, char := range line {
			if char >= '0' && char <= '9' {
				digits = append(digits, int(char-'0'))
			}
		}
		if len(digits) == 1 {
			twoDigit = digits[0]*10 + digits[0]
		} else if len(digits) > 1 {
			twoDigit = digits[0]*10 + digits[len(digits)-1]
		}

		resArr = append(resArr, twoDigit)
		digits = nil
	}
	return resArr
}
