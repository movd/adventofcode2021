package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

	var increaseCount int

	for i, num := range ints {
		if i == 0 {
			fmt.Printf("%d (N/A - no previous measurement)\n", num)
		}

		if i >= 1 {
			if num > ints[i-1] {
				fmt.Printf("%d (increased)\n", num)
				increaseCount++
			} else {
				fmt.Printf("%d (decreased)\n", num)
			}
		}

	}
	fmt.Printf("increased %d times\n", increaseCount)
}
