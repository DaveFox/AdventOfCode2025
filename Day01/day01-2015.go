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

	// testcase1 = 58
	fmt.Println(workOutPaper("2x3x4"))
	// testcase2 = 43
	fmt.Println(workOutPaper("1x1x10"))
	//fmt.Println(workOutPaper("10x1x1"))
	//fmt.Println(workOutPaper("1x10x1"))

	//fmt.Println(workOutPaper("20x8x20"))

	total := 0
	for i := 0; i < len(lines); i++ {
		total += workOutPaper(lines[i])
	}
	fmt.Println("")
	fmt.Println(total)
}

func fileToString() string {
	file, err := os.Open("./input-2015.txt")
	if err != nil {
		panic(err)
	}

	builder := new(strings.Builder)
	io.Copy(builder, file)
	return builder.String()
}

// 2*l*w + 2*w*h + 2*h*l
func workOutPaper(inputStr string) int {
	parts := strings.Split(inputStr, "x")
	length, errL := strconv.Atoi(parts[0])
	width, errW := strconv.Atoi(parts[1])
	height, errH := strconv.Atoi(parts[2])
	if errL != nil || errW != nil || errH != nil {
		panic("Invalid input")
	}

	// 2*l*w + 2*w*h + 2*h*l
	side1 := 2 * length * width
	side2 := 2 * width * height
	side3 := 2 * height * length

	smallest := 0
	if side1 < side2 && side1 < side3 {
		smallest = side1
	} else if side2 < side3 {
		smallest = side2
	} else {
		smallest = side3
	}

	total := side1 + side2 + side3 + (smallest / 2)
	//fmt.Println("slack", smallest/2)
	return total
}
