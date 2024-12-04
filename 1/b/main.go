package main

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
)

func main() {

	fileName := "../input.txt"
	data, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(data)
	reader.Comma = ' '
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()

	var a []int
	var b []int
	var c int = 0

	for _, row := range records {
		aa, err := strconv.Atoi(row[0])
		bb, err := strconv.Atoi(row[1])
		a = append(a, aa)
		b = append(b, bb)

		if err != nil {
			panic(err)
		}
	}

	for _, outer := range a {
		for _, inner := range b {
			if outer == inner {
				c += outer
			}

		}
	}

	fmt.Println(c)
}
