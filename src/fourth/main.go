package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func splitCards(s string) ([]int, []int) {
	ownedCard := strings.Split(s, " | ")[0]
	ownedCard = strings.Split(ownedCard, ": ")[1]
	ownedCardArr := strings.Split(ownedCard, " ")

	var ownedCardNums []int
	for _, val := range ownedCardArr {
		num, err := strconv.Atoi(val)
		if err == nil {
			ownedCardNums = append(ownedCardNums, num)
		}
	}

	winningCard := strings.Split(s, " | ")[1]

	var winningCardNums []int
	winningCardArr := strings.Split(winningCard, " ")
	for _, val := range winningCardArr {
		num, err := strconv.Atoi(val)
		if err == nil {
			winningCardNums = append(winningCardNums, num)
		}
	}
	return ownedCardNums, winningCardNums
}

func findPoints(ownedCards [][]int, winningCards [][]int) int {
	points := 0

	for i, ownedCard := range ownedCards {
		returnVal := 0
		for _, num := range ownedCard {
			if slices.Contains(winningCards[i], num) {
				if returnVal == 0 {
					returnVal = 1
				} else {
					returnVal = returnVal * 2
				}
			}
		}
		points += returnVal
	}
	return points
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// var sum int = 0

	ownedCards := [][]int{}
	winningCards := [][]int{}

	for scanner.Scan() {
		var currentLine string = scanner.Text()

		ownedArr, winningArr := splitCards(currentLine)
		ownedCards = append(ownedCards, ownedArr)
		winningCards = append(winningCards, winningArr)
	}

	fmt.Println(findPoints(ownedCards, winningCards))
}
