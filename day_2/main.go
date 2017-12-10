package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func rowChecksum(input []int64) int64 {
	if len(input) == 0 {
		return 0
	}
	largest, smallest := input[0], input[0]
	for _, i := range input {
		if i > largest {
			largest = i
		} else if i < smallest {
			smallest = i
		}
	}
	return largest - smallest
}

func tableChecksum(input [][]int64) (sum int64) {
	for _, row := range input {
		sum += rowChecksum(row)
	}
	return sum
}

func rowChecksum2(input []int64) int64 {
	li := len(input)
	if li == 0 {
		return 0
	}
	for i := 0; i < li-1; i++ {
		for j := i + 1; j < li; j++ {
			a, b := input[i], input[j]
			if a >= b && a%b == 0 {
				return a / b
			}
			if b > a && b%a == 0 {
				return b / a
			}
		}
	}
	return 0
}

func tableChecksum2(input [][]int64) (sum int64) {
	for _, row := range input {
		sum += rowChecksum2(row)
	}
	return sum
}

func main() {
	input, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	rawRows := strings.Split(string(input), "\n")
	tableRows := make([][]int64, len(rawRows))
	for i := 0; i < len(rawRows); i++ {
		rawRows[i] = strings.TrimSpace(rawRows[i])
		if len(rawRows[i]) == 0 {
			// catch empty rows!
			continue
		}
		rawDigits := strings.Split(strings.TrimSpace(rawRows[i]), "\t")
		row := make([]int64, len(rawDigits))
		for j := 0; j < len(rawDigits); j++ {
			v, err := strconv.ParseInt(rawDigits[j], 0, 64)
			if err != nil {
				panic(err)
			}
			row[j] = v
		}
		tableRows[i] = row
	}
	fmt.Println(tableChecksum(tableRows))
	fmt.Println(tableChecksum2(tableRows))
}
