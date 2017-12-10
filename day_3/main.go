package main

import (
	"fmt"
)

func DataDistance(square int64) int64 {
	if square < 1 {
		panic("square cannot be less than 1")
	}
	if square == 1 {
		return 0
	}

	pos := int64(1)

	// a more efficient solution would have a constant time calculation for cycle? (right?)
	// - one could use 2n^2 + 2n - 3 in a binary search (slighty faster < O(N/4) )

	for cycle := int64(1); true; cycle++ {
		cycleWidth := cycle * 2
		cycleLength := cycleWidth * 4

		// is square in our current cycle?
		if pos+cycleLength >= square {
			distanceIntoCycle := square - pos
			x := distanceIntoCycle % cycleWidth
			y := x - cycle
			if y < 0 {
				y = -y
			}
			return y + cycle
		}
		pos = pos + cycleLength
	}
	return 0
}

func main() {
	fmt.Println(DataDistance(361527))
}
