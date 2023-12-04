package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	finalNum := 0

	for scanner.Scan() {
		var currentLine string = scanner.Text()
		fmt.Println(currentLine)

		currentLine = strings.ReplaceAll(currentLine, "one", "o1e")
		currentLine = strings.ReplaceAll(currentLine, "two", "t2o")
		currentLine = strings.ReplaceAll(currentLine, "three", "t3e")
		currentLine = strings.ReplaceAll(currentLine, "four", "f4r")
		currentLine = strings.ReplaceAll(currentLine, "five", "f5e")
		currentLine = strings.ReplaceAll(currentLine, "six", "s6x")
		currentLine = strings.ReplaceAll(currentLine, "seven", "s7n")
		currentLine = strings.ReplaceAll(currentLine, "eight", "e8t")
		currentLine = strings.ReplaceAll(currentLine, "nine", "n9e")

		currentArr := strings.Split(currentLine, "")

		var numbers []string
		for i := 0; i < len(currentArr); i++ {
			_, err := strconv.Atoi(currentArr[i])
			if err == nil {
				numbers = append(numbers, currentArr[i])
			}
		}

		if len(numbers) != 0 {
			firstNum := numbers[0]
			lastNum := numbers[len(numbers)-1]

			i, _ := strconv.Atoi(fmt.Sprintf("%v%v", firstNum, lastNum))
			fmt.Println(currentLine, i)
			fmt.Println("----------")

			finalNum += i
		}
	}

	fmt.Println(finalNum)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
