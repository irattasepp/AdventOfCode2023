package main

import (
	"AdventOfCode2023/day02/funcs"
	"AdventOfCode2023/helpers"
	"fmt"
)

func main() {
	inputFile, err := helpers.ReadFileLines("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// -------PART ONE-------
	resultOne := funcs.PartOneD2(inputFile)
	fmt.Printf("Answer to Part One: %d\n", resultOne)

	// -------PART ONE-------
	resultTwo := funcs.PartTwoD2(inputFile)
	fmt.Printf("Answer to Part Two: %d\n", resultTwo)
}
