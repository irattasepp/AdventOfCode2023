package funcs

import (
	"strconv"
	"strings"
)

func PartOneD2(fileLines []string) int {
	var output int

	// which are possible with only 12 red cubes, 13 green cubes, and 14 blue cubes
	cubeMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, game := range fileLines {
		lineParts := strings.Split(game, ": ")
		partsId := strings.Split(lineParts[0], " ")
		gameId, _ := strconv.Atoi(partsId[1])
		enoughCubes := true
		rounds := strings.Split(lineParts[1], "; ")
		for _, selected := range rounds {
			cubes := strings.Split(selected, ", ")
			for _, colors := range cubes {
				eachCol := strings.Split(colors, " ")
				amountEachCol, _ := strconv.Atoi(eachCol[0])
				if amountEachCol > cubeMap[eachCol[1]] {
					enoughCubes = false
					break
				}
			}
		}
		if enoughCubes {
			output += gameId
		}
	}
	return output
}
