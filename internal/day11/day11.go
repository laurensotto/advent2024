package day11

import (
	"fmt"
	"github.com/laurensotto/advent2024/pkg/sliceutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func Solve(input string, verbose bool) (string, int64, string, int64) {
	splitString := strings.Split(input, " ")
	copiedString := sliceutil.DeepCopySlice(splitString)

	startTime1 := time.Now()
	part1Result := part1(splitString, verbose)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(copiedString, verbose)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(input []string, verbose bool) int {
	for i := 0; i < 25; i++ {
		input = blink(input, verbose)
	}

	return len(input)
}

func blink(input []string, verbose bool) []string {
	if verbose {
		fmt.Println(input)
	}

	currentIndexes := len(input) - 1

	for i := 0; i <= currentIndexes; i++ {
		if input[i] == "0" {
			input[i] = "1"

			continue
		}

		if len(input[i])%2 == 0 {
			splitString := strings.Split(input[i], "")

			var firstHalf string
			var secondHalf string

			for j := 0; j < len(splitString); j++ {
				if j < len(splitString)/2 {
					firstHalf += splitString[j]
				} else {
					secondHalf += splitString[j]
				}
			}

			input[i] = firstHalf

			secondHalfInt, err := strconv.Atoi(secondHalf)

			if err != nil {
				log.Fatal(err)
			}

			input = append(input, strconv.Itoa(secondHalfInt))

			continue
		}

		intValue, err := strconv.Atoi(input[i])

		if err != nil {
			log.Fatal(err)
		}

		input[i] = strconv.Itoa(intValue * 2024)
	}

	return input
}

func part2(input []string, verbose bool) int {
	stones := make(map[int]int)

	for _, value := range input {
		intValue, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		stones[intValue] += 1
	}

	for i := 0; i < 75; i++ {
		stones = blinkOptimized(stones, verbose)
	}

	var totalStones int

	for _, countOfStones := range stones {
		totalStones += countOfStones
	}

	return totalStones
}

func blinkOptimized(stones map[int]int, verbose bool) map[int]int {
	if verbose {
		fmt.Println(stones)
	}

	newStones := make(map[int]int)

	for stoneValue, stoneCount := range stones {
		if stoneValue == 0 {
			newStones[1] += stoneCount
			continue
		}

		stringStoneValue := strconv.Itoa(stoneValue)

		if len(stringStoneValue)%2 == 0 {
			mid := len(stringStoneValue) / 2
			firstHalf, firstErr := strconv.Atoi(stringStoneValue[:mid])
			secondHalf, secondErr := strconv.Atoi(stringStoneValue[mid:])

			if firstErr != nil {
				log.Fatal(firstErr)
			}

			if secondErr != nil {
				log.Fatal(secondErr)
			}

			newStones[firstHalf] += stoneCount
			newStones[secondHalf] += stoneCount
		} else {
			newStones[stoneValue*2024] += stoneCount
		}
	}

	return newStones
}
