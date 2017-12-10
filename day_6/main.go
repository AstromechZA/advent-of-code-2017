package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cycle(buckets []int) {
	biggest, biggestI := 0, 0
	for i, b := range buckets {
		if b > biggest {
			biggest = b
			biggestI = i
		}
	}
	remaining := biggest
	buckets[biggestI] = 0
	for i, p := 0, biggestI+1; i < remaining; i, p = i+1, p+1 {
		if p >= len(buckets) {
			p = p % len(buckets)
		}
		buckets[p]++
	}
}

func Sign(buckets []int) string {
	return fmt.Sprintf("%v", buckets)
}

func CountCycles(buckets []int) (count int, diff int) {
	seensigs := make(map[string]int)
	seensigs[Sign(buckets)] = 0
	for i := 0; true; i++ {
		Cycle(buckets)
		count++
		s := Sign(buckets)
		seenat, ok := seensigs[s]
		if ok {
			return i, i - seenat
		}
		seensigs[s] = i
	}
	return
}

func main() {
	r := bufio.NewScanner(bufio.NewReader(os.Stdin))
	if !r.Scan() {
		panic("input error: " + r.Err().Error())
	}
	parts := strings.Split(strings.TrimSpace(r.Text()), "\t")
	buckets := make([]int, len(parts))
	for i, p := range parts {
		v, _ := strconv.ParseInt(strings.TrimSpace(p), 0, 64)
		buckets[i] = int(v)
	}
	fmt.Println(CountCycles(buckets))
}
