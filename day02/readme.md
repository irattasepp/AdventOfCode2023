# Day 02

## Overview

This code involves processing a text file (`data.txt`, not included), each line representing a game with it's rounds.

## Functions

1. **PartOne**

```go
func PartOneD2(fileLines []string) int
```

- **Description**: Finds the IDs of the games that are possible with specified amount of colored cubes. Returns the sum of IDs.

2. **PartTwo**

```go
func PartTwoD2(fileLines []string) int
```

- **Description**: Finds the minimum amount of cubes for each color needed for each game. Multiplies together the resulting numbers. Returns the sum of all multiplications.
