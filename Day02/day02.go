package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("AoC Day 2: Part 1")

	fileStr := fileToString()

	//// Testcases
	//testInput := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	//testRanges := strings.Split(testInput, ",")
	//fmt.Println(part1([]string{"11-22"}), "expect 11 & 22")
	//fmt.Println(part1([]string{"95-115"}), "expect 99")
	//fmt.Println(part1([]string{"998-1012"}), "expect 1010")
	//fmt.Println(part1(testRanges), "expect 1227775554")

	idRanges := strings.Split(fileStr, ",")
	result1 := part1(idRanges)
	fmt.Println("Part 1:", result1)

	//// Testcases
	//fmt.Println(part2(testRanges), "expect 4174379265")

	fmt.Println("Part 2:")
	elapsed := time.Since(start)
	fmt.Println("Took: ", elapsed)
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

func part1(idRanges []string) int {
	total := 0
	for i := 0; i < len(idRanges); i++ {
		minId := strings.Split(idRanges[i], "-")[0]
		minIdInt, _ := strconv.Atoi(minId)
		maxId := strings.Split(idRanges[i], "-")[1]
		maxIdInt, _ := strconv.Atoi(maxId)

		for j := minIdInt; j <= maxIdInt; j++ {
			currentVal := strconv.Itoa(j)
			length := len(currentVal)
			firstHalf := currentVal[0 : length/2]
			secondHalf := currentVal[length/2 : length]
			if firstHalf == secondHalf {
				//fmt.Println("Found invalid ID:", currentVal)
				total += j
			}
		}
	}
	// fmt.Println("Total summed invalid IDs:", total)
	return total
}
