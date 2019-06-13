package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func letters(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range s {
		if freq[char] == 0 {
			freq[char] = 1
		} else {
			freq[char]++
		}
	}
	return freq
}

func sortLetters(m map[rune]int) []string {
	var s []string
	for letter, freq := range m {
		s = append(s, string(letter)+string(':')+strconv.Itoa(freq))
	}
	sort.Strings(s)
	return s
}

func main() {
	fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))
}
