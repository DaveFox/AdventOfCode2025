package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fileStr := fileToString()

	testInput := "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@."
	lines := strings.Split(testInput, "\n")
	grid := buildGrid(lines)
	printGrid(grid)

	// Testcase
	fmt.Println(part1(grid, 4)) // 13

	lines2 := strings.Split(fileStr, "\n")
	grid2 := buildGrid(lines2)
	fmt.Println(part1(grid2, 4))
	fmt.Println("Took: ", time.Since(start))
}

func fileToString() string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	builder := new(strings.Builder)
	io.Copy(builder, file)
	return builder.String()
}

func buildGrid(lines []string) [][]string {
	size := len(lines)
	grid := make([][]string, size)
	for g := range grid {
		grid[g] = make([]string, size)
	}
	for i, line := range lines {
		chars := strings.Split(line, "")
		for j, char := range chars {
			grid[i][j] = char
		}
	}
	return grid
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(" ")
}

func part1(grid [][]string, maxNeighbors int) int {
	goodPositions := 0
	for rowIndex, row := range grid {
		for i := 0; i < len(row); i++ {
			count := 0
			if row[i] == "@" {
				if i-1 >= 0 && row[i-1] == "@" {
					//fmt.Println("one to left")
					count++
				}
				if i-1 >= 0 && rowIndex > 0 && grid[rowIndex-1][i-1] == "@" {
					//fmt.Println("one above left")
					count++
				}
				if rowIndex > 0 && grid[rowIndex-1][i] == "@" {
					//fmt.Println("one above")
					count++
				}
				if rowIndex > 0 && i+1 < len(row) && grid[rowIndex-1][i+1] == "@" {
					//fmt.Println("one above right")
					count++
				}
				if i+1 < len(row) && grid[rowIndex][i+1] == "@" {
					//fmt.Println("one right")
					count++
				}
				if i+1 < len(row) && rowIndex+1 < len(grid) && grid[rowIndex+1][i+1] == "@" {
					//fmt.Println("one right bellow")
					count++
				}
				if rowIndex+1 < len(grid) && grid[rowIndex+1][i] == "@" {
					//fmt.Println("one bellow")
					count++
				}
				if i-1 >= 0 && rowIndex+1 < len(grid) && grid[rowIndex+1][i-1] == "@" {
					//fmt.Println("one left bellow")
					count++
				}
				//fmt.Println(count)
				if count < maxNeighbors {
					//fmt.Println("Good pos:", rowIndex, i, count)
					goodPositions++
				}
			}
		}
	}
	return goodPositions
}

// @ @ .
//. @ @
//. @ .

// if rowIndex == 5 && i == 1 {
//	 fmt.Println(grid[rowIndex-1][i-1], grid[rowIndex-1][i], grid[rowIndex-1][i+1])
//	 fmt.Println(grid[rowIndex][i-1], grid[rowIndex][i], grid[rowIndex][i+1])
//   fmt.Println(grid[rowIndex+1][i-1], grid[rowIndex+1][i], grid[rowIndex+1][i+1])
// }
