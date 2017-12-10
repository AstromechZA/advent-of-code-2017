package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const _ = `
17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...
`

func getCycleStart(n int) int {
	return 2*n*(n+1) - 3
}

func getSpiralPosition(cycle int, d int) int {
	return getCycleStart(cycle) + d
}

func getBufferPosition(cycle int, d int) int {
	return getSpiralPosition(cycle, d) - 1
}

func FindValueUpper(target int, squareSize int) int {
	if target < 1 {
		panic("v cannot be less than 1")
	}
	if target == 1 {
		return 2
	}

	side := squareSize*2 + 1
	buffer := make([]int, side*side)
	x, y := squareSize, squareSize
	buffer[y*side+x] = 1

	for cycle := 1; true; cycle++ {
		log.Printf("cycle %d", cycle)
		cycleWidth := cycle * 2

		x++
		y--

		for i := 0; i < cycleWidth; i++ {
			y++
			for j := -1; j < 2; j++ {
				for k := -1; k < 2; k++ {
					if k != 0 || j != 0 {
						v := buffer[(y+j)*side+x+k]
						if v > 0 {
							buffer[y*side+x] += v
						}
					}
				}
			}
			if buffer[y*side+x] > target {
				return buffer[y*side+x]
			}
		}
		for i := 0; i < cycleWidth; i++ {
			x--
			for j := -1; j < 2; j++ {
				for k := -1; k < 2; k++ {
					if k != 0 || j != 0 {
						v := buffer[(y+j)*side+x+k]
						if v > 0 {
							buffer[y*side+x] += v
						}
					}
				}
			}
			if buffer[y*side+x] > target {
				return buffer[y*side+x]
			}
		}
		for i := 0; i < cycleWidth; i++ {
			y--
			for j := -1; j < 2; j++ {
				for k := -1; k < 2; k++ {
					if k != 0 || j != 0 {
						v := buffer[(y+j)*side+x+k]
						if v > 0 {
							buffer[y*side+x] += v
						}
					}
				}
			}
			if buffer[y*side+x] > target {
				return buffer[y*side+x]
			}
		}
		for i := 0; i < cycleWidth; i++ {
			x++
			for j := -1; j < 2; j++ {
				for k := -1; k < 2; k++ {
					if k != 0 || j != 0 {
						v := buffer[(y+j)*side+x+k]
						if v > 0 {
							buffer[y*side+x] += v
						}
					}
				}
			}
			if buffer[y*side+x] > target {
				return buffer[y*side+x]
			}
		}

		for y := 0; y < side; y++ {
			for x := 0; x < side; x++ {
				fmt.Printf("%6d ", buffer[y*side+x])
			}
			fmt.Println()
		}

	}
	return 0
}

func main() {
	r := bufio.NewScanner(bufio.NewReader(os.Stdin))
	if !r.Scan() {
		panic("input error: " + r.Err().Error())
	}
	x, err := strconv.ParseInt(r.Text(), 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(FindValueUpper(int(x), 5))
}
