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

		// Parse parts
		// ["<num> <color>", ...]
		isGameValid := true
		for _, part := range parts {
			if !isGameValid {
				continue
			}

			splitPart := strings.Split(strings.TrimSpace(part), " ")

			color := splitPart[1]
			quantity, err := strconv.Atoi(splitPart[0])
			check(err)

			//fmt.Printf("%s %s -> %s %d\n", splitPart[1], splitPart[0], color, quantity)

			if color == "red" {
				isGameValid = quantity <= 12
			} else if color == "green" {
				isGameValid = quantity <= 13
			} else { // Blue
				isGameValid = quantity <= 14
			}
		}

		if isGameValid {
			result += gameID
		}
	}

	fmt.Println()
	fmt.Printf("Result: %d", result)
}
