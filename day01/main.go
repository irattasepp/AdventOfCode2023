package main

import (
	"AdventOfCode2023/day01/funcs"
	"AdventOfCode2023/helpers"
	"fmt"
)

func main() {
	inputFile, err := helpers.ReadFileLines("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// -----PART ONE-----
	digitsOne := funcs.PartOne(inputFile)
	resultOne := funcs.SumNumbers(digitsOne)
	fmt.Printf("Answer to Part One: %d\n", resultOne)

	// -----PART TWO-----
	digitsTwo := funcs.PartTwo(inputFile)
	resultTwo := funcs.SumNumbers(digitsTwo)
	fmt.Printf("Answer to Part Two: %d\n", resultTwo)
}
