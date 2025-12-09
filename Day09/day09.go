package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fileStr := fileToString()
	fileLines := strings.Split(fileStr, "\n")
	fmt.Println(fileLines[0])

	testInput := "7,1\n11,1\n11,7\n9,7\n9,5\n2,5\n2,3\n7,3"
	lines := strings.Split(testInput, "\n")

	maxX, maxY := getMaxXY(lines)
	grid := buildGrid(maxX, maxY)
	printGrid(grid)

	addCorners(lines)

	// Testcase
	//fmt.Println(part1(grid, 4)) // 13

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

func getMaxXY(stringArray []string) (int, int) {
	var allX []int
	var allY []int
	for _, line := range stringArray {
		x := strings.Split(line, ",")[0]
		y := strings.Split(line, ",")[1]
		xInt, _ := strconv.Atoi(x)
		yInt, _ := strconv.Atoi(y)
		allX = append(allX, xInt)
		allY = append(allY, yInt)
	}
	sort.Ints(allX)
	sort.Ints(allY)
	maxX := allX[len(allX)-1]
	maxY := allY[len(allY)-1]

	fmt.Println(maxX, maxY)
	return maxX, maxY
}

func addCorners(lines []string) {

}

func buildGrid(maxX int, maxY int) [][]string {
	biggest := max(maxX, maxY)
	grid := make([][]string, biggest)
	for g := range grid {
		grid[g] = make([]string, biggest)
	}
	for i := 0; i < biggest; i++ {
		for j := 0; j < biggest; j++ {
			grid[i][j] = "."
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
	return 0
}
