package main

import (
	"fmt"
	"strings"
	"strconv"
)

type IPAddr [4]byte

func (host IPAddr) String() string {
	ints := make([]string, len(host))
	for index := range host {
		ints[index] = strconv.Itoa(int(host[index]))
	}
	return strings.Join(ints, ".")
}

func main() {
	fmt.Println(IPAddr{127, 0, 0, 1})
}
