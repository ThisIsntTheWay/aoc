package day1part1

func Solve(puzzleInput [][]int) int {
	numsLeft, numsRight := puzzleInput[0], puzzleInput[1]

	var totalDistance int
	for i := range numsLeft {
		delta := numsLeft[i] - numsRight[i]
		if delta < 0 {
			delta = -delta
		}

		totalDistance = totalDistance + delta
	}

	return totalDistance
}
