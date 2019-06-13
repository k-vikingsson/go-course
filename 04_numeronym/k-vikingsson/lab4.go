package main

import (
	"fmt"
	"strconv"
)

func numeronyms(vals ...string) []string {
	var res []string
	for index, val := range vals {
		if (len(val) <= 3) {
			res = append(res, val)
		} else {
			res = append(res, string(val[0])+strconv.Itoa(len(val-2))+string(val[len(val)-1]))
		}
	}
	return res
}

func main() {
	fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))
}