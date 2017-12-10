package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func generateCaptcha(digits string) (sum int) {
	var first, last rune = -1, -1

	// because why not! lets be professional about our input - certain utf8 sequences might trick it if we did straight
	// byte looping. It's also a function I've never used before
	strings.Map(func(r rune) rune {
		// protect yo-self
		if r < 48 || r > 57 {
			panic(fmt.Sprintf("unexpected rune '%v' in input", r))
		}
		if first < 0 {
			first = r
		} else if r == last {
			sum += int(r - 48)
		}
		last = r
		return r // this return could be anything we don't actually care
	}, digits)

	// cyclic input
	if first != -1 && first == last {
		sum += int(first - 48)
	}
	return sum
}

// this version is way less fancy (no utf8 and crazy stuff like v1 :) )
func generateCaptcha2(digits string) (sum int) {
	ld := len(digits)
	if ld%2 != 0 {
		panic("expected even length list")
	}
	for i, ii := 0, ld/2; i < ld; i, ii = i+1, (i+1+ld/2)%ld {
		if digits[i] == digits[ii] {
			sum += int(digits[i]) - 48
		}
	}
	return
}

func main() {
	rawInput, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	digits := strings.TrimSpace(string(rawInput))
	fmt.Println(generateCaptcha(digits))
	fmt.Println(generateCaptcha2(digits))
}
