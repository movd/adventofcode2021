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

func castNumberStringToBinary(s string) int {
	output, _ := strconv.ParseInt(s, 2, 64)
	return int(output)
}

func positionGammaRate(b []int) int {
	// Get gamma rate for nth position
	oneCount := 0
	zeroCount := 0
	gammaRate := 0
	for _, v := range b {
		if v == 1 {
			zeroCount++
		} else {
			oneCount++
		}
	}
	if oneCount < zeroCount {
		gammaRate = 1
	}
	// fmt.Println("zeroes=", zeroCount, "ones=", oneCount)
	return gammaRate
}

func positionEpsilonRate(b []int) int {
	// Get epsilon rate for nth position
	oneCount := 0
	zeroCount := 0
	epsilonRate := 0
	for _, v := range b {
		if v == 1 {
			zeroCount++
		} else {
			oneCount++
		}
	}
	if oneCount > zeroCount {
		epsilonRate = 1
	}
	// fmt.Println("zeroes=", zeroCount, "ones=", oneCount)
	return epsilonRate
}

func concatSliceOfNumbers(s []int) string {
	valuesText := []string{}
	for i := range s {
		number := s[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	concatedString := strings.Join(valuesText, "")
	return concatedString
}

func main() {
	inputSlice, err := readLineInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Create a slice of the positions to compare
	sliceOfBinaries := [][]int{}
	for i := 0; i < len(inputSlice[0]); i++ {
		binaries := []int{}
		for _, v := range inputSlice {
			binary, _ := strconv.Atoi(v[i : i+1])
			binaries = append(binaries, binary)
		}
		sliceOfBinaries = append(sliceOfBinaries, binaries)
	}

	// Loop over slice of slices and get corresponding gamma rate
	gammaRates := []int{}
	epsilonRates := []int{}
	for _, s := range sliceOfBinaries {
		gammaRate := positionGammaRate(s)
		gammaRates = append(gammaRates, gammaRate)
		// fmt.Println("the gamma rate for", s, " is ", gammaRate)

		epsiolonRate := positionEpsilonRate(s)
		epsilonRates = append(epsilonRates, epsiolonRate)
		// fmt.Println("the epsilon rate for", s, " is ", epsiolonRate)

	}
	// fmt.Println(gammaRates)
	// fmt.Println(epsilonRates)

	binaryStringGamma := concatSliceOfNumbers(gammaRates)
	decimalGammaRate := castNumberStringToBinary(binaryStringGamma)
	fmt.Println("Gamma Rate =", binaryStringGamma, "(binary)", decimalGammaRate, "(decimal)")

	binaryStringEpsilon := concatSliceOfNumbers(epsilonRates)
	decimalEpsilonRate := castNumberStringToBinary(binaryStringEpsilon)
	fmt.Println("Epsilon Rate =", binaryStringEpsilon, "(binary)", decimalEpsilonRate, "(decimal)")

	powerConsumption := decimalGammaRate * decimalEpsilonRate
	fmt.Println("Power Consumption =", powerConsumption)

}
