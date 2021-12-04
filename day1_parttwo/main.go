package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
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

	ints := make([]int, len(inputSlice))

	for i, s := range inputSlice {
		ints[i], _ = strconv.Atoi(s)
	}

	// slice for results to check for increase or decrease
	slidingWindowResults := []int{}

	for i := range ints {
		// if i == 0 {
		// 	fmt.Printf("%d (N/A - no previous measurement)\n", num)
		// }
		// three numbers each
		slidingWindow := []int{}
		for j := 0; j < 3; j++ {
			// prevent out of bounds
			if i+j < len(ints) {
				slidingWindow = append(slidingWindow, ints[i+j])
				if len(slidingWindow) == 3 {
					slidingWindowResults = append(slidingWindowResults, sum(slidingWindow))
					fmt.Printf("Result: %d \n", sum(slidingWindow))
				}
			}
		}
	}
	fmt.Println(slidingWindowResults)

	// same as part one
	var increaseCount int

	for i, num := range slidingWindowResults {
		if i == 0 {
			fmt.Printf("%d (N/A - no previous measurement)\n", num)
		}

		if i >= 1 {
			if num > slidingWindowResults[i-1] {
				fmt.Printf("%d **(increased)**\n", num)
				increaseCount++
			} else {
				fmt.Printf("%d (decreased)\n", num)
			}
		}

	}
	fmt.Printf("increased %d times\n", increaseCount)

}
