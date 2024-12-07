package day02

import (
	"strconv"
	"strings"
	"time"
)

func Solve(input string) (string, int64, string, int64) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var reports [][]int

	for _, line := range lines {
		values := strings.Split(line, " ")
		report := make([]int, len(values))
		for i, value := range values {
			intValue, _ := strconv.Atoi(value)
			report[i] = intValue
		}
		reports = append(reports, report)
	}

	startTime1 := time.Now()
	part1Result := part1(reports)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(reports)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(reports [][]int) int {
	var validReports = 0
	for _, report := range reports {

		if checkReportValidity(report) {
			validReports++
		}
	}

	return validReports
}

func part2(reports [][]int) int {
	var validReports = 0
	for _, report := range reports {
		if checkReportValidity(report) {
			validReports++
			continue
		}

		for i := 0; i < len(report); i++ {
			var newReport []int

			for j := 0; j < len(report); j++ {
				if j != i {
					newReport = append(newReport, report[j])
				}
			}

			if checkReportValidity(newReport) {
				validReports++
				break
			}
		}
	}

	return validReports
}

func isNextIntGraduallyHigher(int1 int, int2 int) bool {
	if int1 <= int2 {
		return false
	}

	if int1-int2 > 3 {
		return false
	}

	return true
}

func checkReportValidity(report []int) bool {
	var isReportIncreasing = report[0] < report[1]

	for i, level := range report {
		if i == len(report)-1 {
			break
		}

		if isReportIncreasing && !isNextIntGraduallyHigher(report[i+1], level) {
			return false
		}

		if !isReportIncreasing && !isNextIntGraduallyHigher(level, report[i+1]) {
			return false
		}
	}

	return true
}
