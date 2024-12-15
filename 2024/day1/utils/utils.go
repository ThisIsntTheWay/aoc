package day1utils

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

// Parses day 1 puzzle input and returns a 2d array of type []int:
// [[]numsRight, []numsLeft]
func ParsePuzzleInput(puzzleInput *bufio.Scanner) [][]int {
	var numsLeft []int
	var numsRight []int

	for puzzleInput.Scan() {
		// Puzzle input will always be separated by 3 whitespaces
		split := strings.Split(puzzleInput.Text(), "   ")

		n1, e := strconv.Atoi(split[0])
		if e != nil {
			panic(e)
		}
		numsLeft = append(numsLeft, n1)

		n2, e := strconv.Atoi(split[1])
		if e != nil {
			panic(e)
		}
		numsRight = append(numsRight, n2)
	}
	if e := puzzleInput.Err(); e != nil {
		panic(e)
	}

	// Sort arrs
	sort.Slice(numsLeft, func(i, j int) bool {
		return numsLeft[i] < numsLeft[j]
	})
	sort.Slice(numsRight, func(i, j int) bool {
		return numsRight[i] < numsRight[j]
	})

	return [][]int{numsLeft, numsRight}
}
