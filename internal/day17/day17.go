package day17

import (
	"strconv"
	"time"
)

func Solve(input string, verbose bool) (string, int64, string, int64) {
	startTime1 := time.Now()
	part1Result := part1(input, verbose)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(input, verbose)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(input string, verbose bool) int {
	return 0
}

func part2(input string, verbose bool) int {
	return 0
}
