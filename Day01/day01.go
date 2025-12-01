package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 1")

	fileStr := fileToString()
	lines := strings.Split(fileStr, "\n") // \r\n
	testcase := lines[0]
	fmt.Println(testcase)

	//// Testcases
	//rotateRight(11, "R8") // 19
	//rotateLeft(19, "L19") // 0
	//rotateRight(99, "R1") // 0
	//rotateLeft(0, "L1")   // 99
	//rotateLeft(5, "L10")  // 95
	//rotateRight(95, "R5") // 0
	//
	//// My test cases
	//rotateRight(0, "R99")  // 99
	//rotateRight(0, "R301") // 11
	//rotateRight(0, "R400") // 0
	//rotateLeft(0, "L301")  // 99
	//rotateLeft(0, "L400")  // 0
	//rotateLeft(36, "L46")  // 90

	result1 := part1(lines)
	fmt.Println("Part 1:", result1)
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
