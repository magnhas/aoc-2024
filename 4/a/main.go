package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rows := len(grid)
	cols := len(grid[0])
	word := "XMAS"
	wordLen := len(word)
	count := 0

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}

	checkWord := func(r, c, dr, dc int) bool {
		for i := 0; i < wordLen; i++ {
			nr, nc := r+dr*i, c+dc*i
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != word[i] {
				return false
			}
		}
		return true
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if checkWord(r, c, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	fmt.Printf("%d\n", count)
}

