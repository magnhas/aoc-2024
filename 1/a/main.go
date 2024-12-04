package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


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

	for _, row := range records {
		aa, err := strconv.Atoi(row[0])
		bb, err := strconv.Atoi(row[1])

		a = append(a, aa)
		b = append(b, bb)

		if err != nil {
			panic(err)
		}
	}

	sort.Ints(a)
	sort.Ints(b)

	var c int = 0

	for i := range a {
		c += Abs(a[i]-b[i])
	}

	fmt.Println(c)
}
