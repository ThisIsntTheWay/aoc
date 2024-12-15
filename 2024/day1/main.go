package main

import (
	"bufio"
	"fmt"
	"os"

	day1part1 "github.com/ThisIsntTheWay/aoc/day1/part1"
	day1part2 "github.com/ThisIsntTheWay/aoc/day1/part2"
	day1utils "github.com/ThisIsntTheWay/aoc/day1/utils"
)

func main() {
	const inputFile = "./input"
	input, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("unable to open %s: %v", inputFile, err))
	}

	puzzleInput := day1utils.ParsePuzzleInput(bufio.NewScanner(input))

	fmt.Println("Day 1")
	fmt.Printf("Total distance: %d\n", day1part1.Solve(puzzleInput))
	fmt.Printf("Similarity score: %d\n", day1part2.Solve(puzzleInput))
}
