package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sum struct {
	Red   int
	Green int
	Blue  int
}

func calcCount(arr []string) Sum {
	var red int = 0
	var green int = 0
	var blue int = 0

	for _, v := range arr {
		if strings.Contains(v, "red") {
			count, _ := strconv.Atoi(strings.Split(v, " ")[0])
			if red < count {
				red = count
			}
		}
		if strings.Contains(v, "green") {
			count, _ := strconv.Atoi(strings.Split(v, " ")[0])
			if green < count {
				green = count
			}
		}
		if strings.Contains(v, "blue") {
			count, _ := strconv.Atoi(strings.Split(v, " ")[0])
			if blue < count {
				blue = count
			}
		}
	}

	return Sum{Red: red, Green: green, Blue: blue}
}

func findPower(count Sum) int {
	return count.Red * count.Green * count.Blue
}

func main() {
	limit := Sum{Red: 12, Green: 13, Blue: 14}
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sums []Sum

	for scanner.Scan() {
		var currentLine string = scanner.Text()
		currentLine = strings.Split(currentLine, ": ")[1]
		firstArr := strings.Split(currentLine, "; ")
		var currentArr []string
		for _, v := range firstArr {
			currentArr = append(currentArr, strings.Split(v, ", ")...)
		}
		sums = append(sums, calcCount(currentArr))
	}

	finalCount := 0
	finalSum := 0

	for i, v := range sums {
		finalSum += findPower(v)

		if v.Red <= limit.Red && v.Green <= limit.Green && v.Blue <= limit.Blue {
			finalCount += i + 1
		}
	}
	fmt.Printf("Count: %v, Sum: %v", finalCount, finalSum)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
