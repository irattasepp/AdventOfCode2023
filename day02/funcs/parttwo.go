package funcs

import (
	"strconv"
	"strings"
)

func PartTwoD2(fileLines []string) int {
	var output int

	cubeMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, game := range fileLines {
		lineParts := strings.Split(game, ": ")
		rounds := strings.Split(lineParts[1], "; ")
		for _, selected := range rounds {
			cubes := strings.Split(selected, ", ")
			for _, colors := range cubes {
				eachCol := strings.Split(colors, " ")
				amountEachCol, _ := strconv.Atoi(eachCol[0])
				if amountEachCol > cubeMap[eachCol[1]] {
					cubeMap[eachCol[1]] = amountEachCol
				}
			}
		}
		power := cubeMap["red"] * cubeMap["green"] * cubeMap["blue"]
		output += power
		for key := range cubeMap {
			cubeMap[key] = 0
		}
	}
	return output
}
