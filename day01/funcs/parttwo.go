package funcs

import "strings"

// -----PART TWO-----
func PartTwo(fileLines []string) []int {
	mixedDigits := map[string]string{
		"oneight":   "18",
		"threeight": "38",
		"fiveight":  "58",
		"nineight":  "98",
		"sevenine":  "79",
		"eightwo":   "82",
		"eighthree": "83",
		"twone":     "21",
	}

	wordDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var digits []int
	var resArr []int
	var twoDigit int

	for _, line := range fileLines {
		for oldSeq, newSeq := range mixedDigits {
			str := strings.Replace(line, oldSeq, newSeq, -1)
			line = str
		}
		for oldSeq, newSeq := range wordDigits {
			str := strings.Replace(line, oldSeq, newSeq, -1)
			line = str
		}
		for _, char := range line {
			if char >= '1' && char <= '9' {
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
