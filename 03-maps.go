// https://tour.golang.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	
	stringMap := make(map[string]int)
	for _, v := range fields {
		stringMap[v] += 1
	}
	return stringMap
}

func main() {
	wc.Test(WordCount)
}
