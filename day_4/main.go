package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func main() {
	rawInput, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(rawInput)), "\n")
	validCount := 0
	for _, l := range lines {
		b := isValid(l)
		log.Printf("%v : %v", b, l)
		if b {
			validCount++
		}
	}
	fmt.Println(validCount)
}
