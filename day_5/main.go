package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func escapeTheMaze(offsets []int) (steps int64) {
	for p := 0; p >= 0 && p < len(offsets); steps++ {
		j := offsets[p]
		if j >= 3 {
			offsets[p]--
		} else {
			offsets[p]++
		}
		p += j
	}
	return
}

func main() {
	rawData, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(rawData)), "\n")
	offsets := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		v, err := strconv.ParseInt(lines[i], 0, 64)
		if err != nil {
			panic(err)
		}
		offsets[i] = int(v)
	}
	fmt.Println(escapeTheMaze(offsets))
}
