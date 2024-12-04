package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fileName := "../data.txt"
	data, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(data)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	var safeCount int = 0
	var e1 int
	var e2 int

	for _, row := range records {

		var increasing bool = true
		var decreasing bool = true
		var difference bool = true

		for i := 0; i < len(row)-1; i++ {
			e1, err = strconv.Atoi(row[i])
			e2, err = strconv.Atoi(row[i+1])

			if err != nil {
				panic(err)

			}

			if e1 > e2 {
				increasing = false
			}

			if e1 < e2 {
				decreasing = false
			}

			if e1 == e2 {
				difference = false
			}

			if Abs(e1-e2) > 3 {
				difference = false
			}
		}

		if (increasing || decreasing) && difference {
			safeCount += 1
		}
	}
	fmt.Println(safeCount)
}
