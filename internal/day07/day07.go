package day07

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Solve(input string) (string, int64, string, int64) {
	rows := strings.Split(input, "\n")

	answersSlice := make([]int, len(rows))
	valuesSlice := make([][]int, len(rows))

	for i, row := range rows {
		equationSplit := strings.Split(row, ": ")

		answer, _ := strconv.Atoi(equationSplit[0])
		answersSlice[i] = answer

		var values []int
		for _, value := range strings.Split(equationSplit[1], " ") {
			integerValue, _ := strconv.Atoi(value)

			values = append(values, integerValue)
		}

		valuesSlice[i] = values
	}

	startTime1 := time.Now()
	part1Result := part1(answersSlice, valuesSlice)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(answersSlice, valuesSlice)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(answersSlice []int, valuesSlice [][]int) int {
	var sumOfAnswers int
	for i, answer := range answersSlice {
		values := valuesSlice[i]

		binaryValue := getCharacterStringOfLength(len(values)-1, "0")
		maxBinaryValue := getCharacterStringOfLength(len(values)-1, "1")
		incrementedMaxBinaryValue := incrementBinary(maxBinaryValue)

		for binaryValue != incrementedMaxBinaryValue {
			foundAnswer := values[0]
			splitBinaryValue := strings.Split(binaryValue, "")
			for j, binaryNumber := range splitBinaryValue {
				if binaryNumber == "0" {
					foundAnswer = foundAnswer + values[j+1]
				} else {
					foundAnswer = foundAnswer * values[j+1]
				}
			}

			if foundAnswer == answer {
				sumOfAnswers += answer
				break
			}

			binaryValue = incrementBinary(binaryValue)
		}
	}

	return sumOfAnswers
}

func part2(answersSlice []int, valuesSlice [][]int) int {
	var sumOfAnswers int
	for i, answer := range answersSlice {
		values := valuesSlice[i]

		trinaryValue := getCharacterStringOfLength(len(values)-1, "0")
		maxTrinaryValue := getCharacterStringOfLength(len(values)-1, "2")
		incrementedMaxTrinaryValue := incrementTrinary(maxTrinaryValue)

		for trinaryValue != incrementedMaxTrinaryValue {
			foundAnswer := values[0]
			splitBinaryValue := strings.Split(trinaryValue, "")
			for j, binaryNumber := range splitBinaryValue {
				switch binaryNumber {
				case "0":
					foundAnswer = foundAnswer + values[j+1]
				case "1":
					foundAnswer = foundAnswer * values[j+1]
				case "2":
					result, _ := strconv.Atoi(strconv.Itoa(foundAnswer) + strconv.Itoa(values[j+1]))
					foundAnswer = result
				default:
					log.Fatal("a non trinary number was used")
				}
			}

			if foundAnswer == answer {
				sumOfAnswers += answer
				break
			}

			trinaryValue = incrementTrinary(trinaryValue)
		}
	}

	return sumOfAnswers
}

func incrementBinary(binaryStr string) string {
	num, _ := strconv.ParseInt(binaryStr, 2, 64)
	num++
	incremented := strconv.FormatInt(num, 2)
	padded := fmt.Sprintf("%0*s", len(binaryStr), incremented)
	return padded
}

func incrementTrinary(binaryStr string) string {
	num, _ := strconv.ParseInt(binaryStr, 3, 64)
	num++
	incremented := strconv.FormatInt(num, 3)
	padded := fmt.Sprintf("%0*s", len(binaryStr), incremented)
	return padded
}

func getCharacterStringOfLength(length int, character string) string {
	characterString := ""
	for i := 0; i < length; i++ {
		characterString += character
	}

	return characterString
}
