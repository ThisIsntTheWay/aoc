package day1part2

func Solve(puzzleInput [][]int) int {
	numsLeft, numsRight := puzzleInput[0], puzzleInput[1]
	var similarityScore int
	for _, v := range numsLeft {
		var occurrences int
		for _, w := range numsRight {
			if v == w {
				occurrences++
			}
		}

		similarityScore = similarityScore + (v * occurrences)
	}

	return similarityScore
}
