package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	outFile, err := os.Create("output")
	check(err)
	defer outFile.Close()

	// Regexes
	firstDigitRegx := regexp.MustCompile(`\d`)
	lastDigitRegx := regexp.MustCompile(`.*(\d)`)

	var result []int

	// Iterate lines
	for scanner.Scan() {
		thisString := scanner.Text()
		fmt.Println(thisString)

		// Skip empty lines
		if len(thisString) == 0 {
			continue
		}

		// Iterate through string
		firstDigit := firstDigitRegx.FindString(thisString)
		lastDigit := lastDigitRegx.FindStringSubmatch(thisString)[1]
		// Using FindString would match all up to last digit.
		// Therefore, FindStringSubmatch must be used, which returns an array)

		// Add to our output
		num, err := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		check(err)

		result = append(result, num)
	}

	// Calculate all nums
	var resultCalculated int

	for i := 0; i < len(result); i++ {
		resultCalculated = resultCalculated + result[i]
	}

	fmt.Println("---------------")
	fmt.Printf("Result: %d\n", resultCalculated)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
