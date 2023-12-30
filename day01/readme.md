# Day 01

## Overview

This code involves processing a text file (`data.txt`, not included) containing lines of alphanumeric characters.

## Functions

1. **PartOne**

```go
func PartOne(fileLines []string) []int
```

- **Description**: Finds the first and last digit (1 - 9) in each line and converts them into a two-digit number. Returns a slice of two-digit integers.

2. **PartTwo**

```go
func PartTwo(fileLines []string) []int
```

- **Description**: Finds digits represented as words (one - nine) and replaces them with their numeric counterparts (1 - 9). Then proceeds like `PartOne`. Returns a slice of two-digit integers.

3. **SumNumbers**

```go
func SumNumbers(digits []int) int
```

- **Description**: Calculates the sum of all integers in the provided slice and returns it as an integer.
