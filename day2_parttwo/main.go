package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readLineInput() ([]string, error) {
	// Reads input file to slice of string
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	fileName := os.Args[1]
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inputSlice := strings.Split(string(fileBytes), "\n")
	return inputSlice, nil
}

// Returns a direction of travel and a number of the position
func converToPosition(i string) (string, int) {
	s := strings.Split(i, " ")
	position, err := strconv.Atoi(s[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return s[0], position
}

func main() {
	inputSlice, err := readLineInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, e := range inputSlice {
		direction, position := converToPosition(e)
		if direction == "forward" {
			horizontalPosition = horizontalPosition + position
			depth = depth + (aim * position)
		}

		if direction == "down" {
			aim = aim + position
		}

		if direction == "up" {
			aim = aim - position
		}
	}

	fmt.Println("horizontalPosition:", horizontalPosition)
	fmt.Println("depth:", depth)
	fmt.Println("aim:", aim)
	multiply_result := horizontalPosition * depth
	fmt.Println("result:", multiply_result)
}
