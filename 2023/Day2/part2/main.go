package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type set struct {
	red   []int
	green []int
	blue  []int
}

func main() {
	// Files
	data, err := os.Open("../input")
	check(err)
	defer data.Close()
	scanner := bufio.NewScanner(data)

	var result int
	for scanner.Scan() {
		thisString := scanner.Text()
		culledString := thisString[strings.Index(thisString, ":")+2:]
		check(err)

		delimiterSplitter := func(c rune) bool {
			return c == ';' || c == ','
		}

		set := set{}
		parts := strings.FieldsFunc(culledString, delimiterSplitter)
		for _, part := range parts {
			splitPart := strings.Split(strings.TrimSpace(part), " ")
			color := splitPart[1]
			quantity, _ := strconv.Atoi(splitPart[0])

			switch color {
			case "red":
				set.red = append(set.red, quantity)
			case "green":
				set.green = append(set.green, quantity)
			case "blue":
				set.blue = append(set.blue, quantity)
			}
		}

		// Determine powers
		sort.Ints(set.red)
		sort.Ints(set.green)
		sort.Ints(set.blue)
		mostRed := set.red[len(set.red)-1]
		mostGreen := set.green[len(set.green)-1]
		mostBlue := set.blue[len(set.blue)-1]

		fmt.Printf("%s\n> Red: %d, Green: %d, Blue: %d\n\n", culledString, mostRed, mostGreen, mostBlue)

		result += mostRed * mostGreen * mostBlue
	}

	fmt.Println()
	fmt.Printf("Result: %d", result)
}
