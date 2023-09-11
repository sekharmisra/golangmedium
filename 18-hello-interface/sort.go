package main

import (
	"fmt"
	"sort"
)

type CricketLegends []string

func doSort() {

	legends := CricketLegends{
		"Kapil",
		"Gavaskar",
		"Sachin",
		"Dhoni",
		"Dravid",
		"Kohli",
	}

	sort.Slice(legends, func(i, j int) bool {
		return legends[i] < legends[j]
	})

	fmt.Println(legends)

	var result []string
	for i := len(legends) - 1; i >= 0; i-- {
		result = append(result, legends[i])
	}

	fmt.Println(result)
}
