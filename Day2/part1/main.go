package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Files
	data, err := os.Open("../input")
	check(err)
	defer data.Close()
	scanner := bufio.NewScanner(data)

	var result int

	// Iterate lines
	for scanner.Scan() {
		thisString := scanner.Text()

		// Game 83: 14 red, 2 green; 3 blue, 16 red, 2 green; 4 green, 13 red, 1 blue
		// Cull "Game x: "
		culledString := thisString[strings.Index(thisString, ":")+2:]
		gameIDString := regexp.MustCompile(`\b(\d+)\s*`).FindStringSubmatch(thisString)[0]
		gameID, err := strconv.Atoi(gameIDString)
		check(err)
		//fmt.Println(culledString)

		// Required for multi-delimiter splitting action
		delimiterSplitter := func(c rune) bool {
			return c == ';' || c == ','
		}
		parts := strings.FieldsFunc(culledString, delimiterSplitter)
		//fmt.Println(parts)

		// Parse parts
		// ["<num> <color>", ...]
		isGameValid := true
		for _, part := range parts {
			splitPart := strings.Split(strings.TrimSpace(part), " ")

			color := splitPart[1]
			quantity, err := strconv.Atoi(splitPart[0])
			check(err)

			if color == "red" {
				isGameValid = quantity > 12
			} else if color == "green" {
				isGameValid = quantity > 13
			} else { // Blue
				isGameValid = quantity > 13
			}
		}

		fmt.Printf("%d [%t] - %s\n", gameID, isGameValid, culledString)

		// Abort loop if winning conditions are not met, add to result otherwise

		if isGameValid {
			result = result + gameID
		}

		/*
			// Parse each set of the games
			sets := strings.Split(culledString, ";")

			var cubeSet [3]int
			gameIsValid := true

			for _, set := range sets {
				pairs := strings.Split(strings.TrimSpace(set), ",")

				for _, pair := range pairs {
					var number int
					var color string
					_, err := fmt.Sscanf(pair, "%d %s", &number, &color)
					if err != nil {
						fmt.Printf("Error parsing pair: %s\n", pair)
						continue
					}

					var cubeIndex int
					if color == "red" {
						cubeIndex = 0
					} else if color == "green" {
						cubeIndex = 1
					} else {
						cubeIndex = 2
					}

					cubeSet[cubeIndex] = cubeSet[cubeIndex] + number
				}

				// Pair verdict
				for i := 0; i < len(cubeSet); i++ {
					if cubeSet[i] > (12 + i) {
						gameIsValid = false
					}
				}
			}

			if gameIsValid {
				fmt.Printf("%d - %s\n", gameID, culledString)
				result = result + gameID
			}
		*/
	}

	fmt.Printf("Result: %d", result)
}
