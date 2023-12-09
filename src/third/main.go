package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Index struct {
	row    int
	column int
}

type numStruct struct {
	index Index
	value int
}

func addIndex(a, b Index) Index {
	return Index{a.row + b.row, a.column + b.column}
}

func checkForSymbols(numIndex []numStruct, symbolIndexes []Index) bool {
	directions := []Index{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, n := range numIndex {
		for _, v := range directions {
			if slices.Contains(symbolIndexes, addIndex(n.index, v)) {
				return true
			}
		}
	}

	return false
}

func splitNumbers(numIndexes []numStruct) [][]numStruct {
	var prevV numStruct
	var nums [][]numStruct
	var currentNums []numStruct
	for _, v := range numIndexes {
		if v.index.column-prevV.index.column == 1 {
			currentNums = append(currentNums, v)
			prevV = v
		} else {
			nums = append(nums, currentNums)
			currentNums = []numStruct{v}
			prevV = v
		}
	}

	return nums
}

func checkForEngineParts(numIndexes []numStruct, symbolIndexes []Index) int {
	nums := splitNumbers(numIndexes)
	returnNum := 0

	for _, v := range nums {
		if checkForSymbols(v, symbolIndexes) {
			sum := ""
			for _, vv := range v {
				sum = fmt.Sprintf("%v%v", sum, vv.value)
			}
			sumNum, _ := strconv.Atoi(sum)
			returnNum += sumNum
		}
	}

	return returnNum
}

func main() {
	var numIndexes []numStruct
	var symbolIndexes []Index

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rowNum int = 0

	for scanner.Scan() {
		var currentLine string = scanner.Text()
		currentArr := strings.Split(currentLine, "")

		for i, v := range currentArr {
			num, err := strconv.Atoi(v)
			if err == nil {
				numIndexes = append(numIndexes, numStruct{Index{rowNum, i}, num})
			}
			if err != nil && v == "*" {
				symbolIndexes = append(symbolIndexes, Index{rowNum, i})
			}
		}
		rowNum++
	}

	fmt.Println(checkForEngineParts(numIndexes, symbolIndexes))
}
