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

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

var numberLookup = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
var result []int

func main() {
	// Files
	data, err := os.Open("../input2")
	check(err)
	defer data.Close()
	scanner := bufio.NewScanner(data)

	// Regexes
	//regxp := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	regxp := regexp.MustCompile(`\b(one|two|three|four|five|six|seven|eight|nine|\d)\b`)

	// Iterate lines
	for scanner.Scan() {
		thisString := scanner.Text()
		fmt.Println(thisString)

		// Skip empty lines
		if len(thisString) == 0 {
			continue
		}

		match := regxp.FindAllString(thisString, -1)
		fmt.Println(match)

		firstDigit := ""
		lastDigit := ""
		for i := 0; i < len(match); i++ {
			thisFinalDigitPresentable := ""

			// Convert digit
			if isDigit(match[i]) {
				thisFinalDigitPresentable = match[i]
			} else {
				// Get from numberLookup
				thisFinalDigitPresentable = fmt.Sprintf("%d", numberLookup[match[i]])
			}

			// Store digits
			if i == 0 {
				firstDigit = thisFinalDigitPresentable
			} else if i == len(match)-1 {
				lastDigit = thisFinalDigitPresentable
			}
		}

		// Combine digits
		num, err := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		check(err)

		fmt.Printf("> %d\n", num)
		fmt.Println()

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
