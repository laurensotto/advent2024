package day03

import (
	"regexp"
	"strconv"
	"time"
)

func Solve(input string, verbose bool) (string, int64, string, int64) {
	startTime1 := time.Now()
	part1Result := part1(input)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(input)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(input string) int {
	var calculationTotal = 0

	pattern := `mul\([0-9]{1,3},[0-9]{1,3}\)`

	re, _ := regexp.Compile(pattern)

	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		pattern := `[0-9]{1,3}`

		re, _ := regexp.Compile(pattern)

		stringFormattedInts := re.FindAllString(match, -1)

		valueOne, _ := strconv.Atoi(stringFormattedInts[0])
		valueTwo, _ := strconv.Atoi(stringFormattedInts[1])

		calculationTotal += valueOne * valueTwo
	}

	return calculationTotal
}

func part2(input string) int {
	var calculationTotal = 0

	pattern := `mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`

	re, _ := regexp.Compile(pattern)

	matches := re.FindAllString(input, -1)

	do := true
	for _, match := range matches {
		if match == "do()" {
			do = true
			continue
		}

		if match == "don't()" {
			do = false
			continue
		}

		if !do {
			continue
		}

		pattern := `[0-9]{1,3}`

		re, _ := regexp.Compile(pattern)

		stringFormattedInts := re.FindAllString(match, -1)

		valueOne, _ := strconv.Atoi(stringFormattedInts[0])
		valueTwo, _ := strconv.Atoi(stringFormattedInts[1])

		calculationTotal += valueOne * valueTwo
	}

	return calculationTotal
}
