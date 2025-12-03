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
	//testInput := "987654321111111\n811111111111119\n234234234234278\n818181911112111"
	//fmt.Println(part1(testInput), "expect 357")

	fmt.Println("Part 1 result:", part1(fileStr))

	//// Testcases
	//testInput := "987654321111111\n811111111111119\n234234234234278\n818181911112111"
	//fmt.Println(maxValInAllowedSpace("287654321911111", 2))
	//fmt.Println(part2("287654321911111\n811111111111119"))
	//fmt.Println(part2(testInput)) // expect 987654321111,811111111119,434234234278, 888911112111 = 3121910778619

	fmt.Println("Part 2 result:", part2(fileStr))
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

		combined, _ := strconv.Atoi(string(max1) + string(max2))
		total += combined
	}
	return total
}

func part2(testInput string) int {
	battaryLines := strings.Split(testInput, "\n")
	total := 0
	for i := 0; i < len(battaryLines); i++ {
		remainingSpace := 12
		currentBattery := battaryLines[i]
		maxString := ""
		for remainingSpace > 0 {
			maxVal, newBattery := maxValInAllowedSpace(currentBattery, remainingSpace)
			maxString += maxVal
			currentBattery = newBattery
			remainingSpace--
		}
		// fmt.Println(maxString)
		asNumber, _ := strconv.Atoi(maxString)
		total += asNumber
	}
	return total
}

func maxValInAllowedSpace(batteryLine string, remainingSpace int) (string, string) {
	localMax := int32(0)
	maxPos := 0
	for pos, jolt := range batteryLine {
		if pos > len(batteryLine)-remainingSpace {
			break
		}
		if jolt > localMax {
			localMax = jolt
			maxPos = pos
		}
	}
	return string(localMax), batteryLine[maxPos+1:]
}
