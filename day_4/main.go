package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func isValid(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	seen := make(map[string]bool)
	for _, w := range words {
		if seen[w] {
			return false
		}
		seen[w] = true
	}
	return true
}

func isValidStrict(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	seen := make(map[string]bool)
	for _, w := range words {
		b := []byte(w)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		w = string(b)
		if seen[w] {
			return false
		}
		seen[w] = true
	}
	return true
}

func main() {
	rawInput, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(rawInput)), "\n")
	validCount := 0
	for _, l := range lines {
		b := isValidStrict(l)
		log.Printf("%v : %v", b, l)
		if b {
			validCount++
		}
	}
	fmt.Println(validCount)
}
