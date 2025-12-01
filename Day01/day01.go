package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 1: Part 1")

	fileStr := fileToString()
	lines := strings.Split(fileStr, "\n") // \r\n

	//// Testcases
	//rotateRight(11, "R8") // 19
	//rotateLeft(19, "L19") // 0
	//rotateRight(99, "R1") // 0
	//rotateLeft(0, "L1")   // 99
	//rotateLeft(5, "L10")  // 95
	//rotateRight(95, "R5") // 0

	result1 := part1(lines)
	fmt.Println("Part 1:", result1)

	fmt.Println("AoC Day 1: Part 2")
	// Testcases
	fmt.Println(part2([]string{"R1000"}), "expect 10")
	fmt.Println(part2([]string{"L1000"}), "expect 10")
	fmt.Println(part2([]string{"R49"}), "expect 0")
	fmt.Println(part2([]string{"R50"}), "expect 1")
	fmt.Println(part2([]string{"R1"}), "expect 0")
	fmt.Println(part2([]string{"L68"}), "expect 1")
	fmt.Println(part2([]string{"L50"}), "expect 1")
	fmt.Println(part2([]string{"L494"}), "expect 5") // 50 -> -444
	fmt.Println(part2([]string{"L50", "L144", "L300"}), "expect 5")
	fmt.Println(part2([]string{"L94", "L400"}), "expect 5")
	fmt.Println(part2([]string{"L144"}), "expect 1")
	fmt.Println(part2([]string{"L300"}), "expect 3")
	fmt.Println(part2([]string{"R444"}), "expect 4")
	fmt.Println(part2([]string{"L50", "L1"}), "expect 1")
	fmt.Println(part2([]string{"R49", "R1", "R1"}), "expect 1")

	result2 := part2(lines)
	fmt.Println("Part 2:", result2)
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

func part1(inputLines []string) int {
	currValue := 50
	numberOfZeros := 0
	for i := 0; i < len(inputLines); i++ {
		if inputLines[i][0] == 'R' {
			currValue = rotateRight(currValue, inputLines[i])
			if currValue == 0 {
				numberOfZeros++
			}
		}
		if inputLines[i][0] == 'L' {
			currValue = rotateLeft(currValue, inputLines[i])
			if currValue == 0 {
				numberOfZeros++
			}
		}
	}

	return numberOfZeros
}

func rotateRight(startValue int, instruction string) int {
	increaseBy, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(err)
	}

	returnVal := startValue + increaseBy
	if returnVal > 99 {
		returnVal = returnVal % 100
	}

	//fmt.Println("Start", startValue, "change", instruction, "result", returnVal)
	return returnVal
}

func rotateLeft(startValue int, instruction string) int {
	increaseBy, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(err)
	}

	returnVal := startValue - increaseBy
	if returnVal < 0 {
		returnVal = ((returnVal % 100) + 100) % 100
	}

	//fmt.Println("Start", startValue, "change", instruction, "result", returnVal)
	return returnVal
}

func part2(inputLines []string) int {
	currValue := 50
	numberOfZeros := 0
	for i := 0; i < len(inputLines); i++ {
		if inputLines[i][0] == 'R' {
			increaseBy, err := strconv.Atoi(inputLines[i][1:])
			if err != nil {
				panic(err)
			}

			returnVal := currValue + increaseBy
			numberOfZeros += (returnVal / 100)
			if returnVal > 99 {
				returnVal = returnVal % 100
			}
			currValue = returnVal
		}
		if inputLines[i][0] == 'L' {
			decreaseBy, err := strconv.Atoi(inputLines[i][1:])
			if err != nil {
				panic(err)
			}

			returnVal := currValue - decreaseBy
			if returnVal < 0 {
				returnVal = ((returnVal % 100) + 100) % 100
			}

			if (currValue - decreaseBy) <= 0 {
				if currValue == 0 {
					numberOfZeros += decreaseBy / 100
				} else if decreaseBy >= currValue {
					numberOfZeros += 1 + (decreaseBy-currValue)/100
				}
			}
			currValue = returnVal
		}
	}

	return numberOfZeros
}
