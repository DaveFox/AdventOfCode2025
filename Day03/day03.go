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
	fmt.Println("AoC Day 3: Part 1")

	fileStr := fileToString()

	//// Testcases
	testInput := "987654321111111\n811111111111119\n234234234234278\n818181911112111"
	fmt.Println(part1(testInput), "expect 357")

	fmt.Println("Part 1 result:", part1(fileStr))
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

func part1(testInput string) int {
	batteryLines := strings.Split(testInput, "\n")
	total := 0

	for i := 0; i < len(batteryLines); i++ {
		max1 := int32(0)
		max1Pos := 0
		for pos, jolt := range batteryLines[i] {
			if pos == len(batteryLines[i])-1 {
				continue
			}
			if jolt > max1 {
				max1 = jolt
				max1Pos = pos
			}
		}
		secondHalf := batteryLines[i][max1Pos+1:]
		max2 := int32(0)
		for _, jolt := range secondHalf {
			if jolt > max2 {
				max2 = jolt
			}
		}
		//fmt.Println("max1", string(max1), "max2", string(max2))
		combined, _ := strconv.Atoi(string(max1) + string(max2))
		//fmt.Println(combined)
		total += combined
	}
	return total
}
