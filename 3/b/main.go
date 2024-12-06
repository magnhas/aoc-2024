package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mul(numbers string) int {
	re := regexp.MustCompile(`\d+`)
	numberArray := re.FindAllString(numbers, -1)
	x, err := strconv.Atoi(string(numberArray[0]))
	y, err := strconv.Atoi(string(numberArray[1]))
	check(err)
	return x * y
}

func main() {

	buffer, err := os.ReadFile("../data.txt")
	check(err)

	var sum int = 0
	var active bool = true

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(string(buffer), -1)

	for _, match := range matches {

		switch match {
		case "do()":
			active = true
		case "don't()":
			active = false
		default:
			if active {
				sum += mul(match)
			}
		}
	}
	fmt.Printf("%d\n", sum)

}
