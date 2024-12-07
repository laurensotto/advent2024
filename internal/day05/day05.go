package day05

import (
	"errors"
	"github.com/laurensotto/advent2024/pkg/sliceutil"
	"strconv"
	"strings"
	"time"
)

func Solve(input string) (string, int64, string, int64) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var manualPages [][]int
	manualRules := make(map[int][]int)

	foundBreakPoint := false
	for i := range lines {
		if lines[i] == "" {
			foundBreakPoint = true
			continue
		}

		if !foundBreakPoint {
			values := strings.Split(lines[i], "|")

			beforeValue, _ := strconv.Atoi(values[0])
			afterValue, _ := strconv.Atoi(values[1])

			if _, ok := manualRules[beforeValue]; !ok {
				manualRules[beforeValue] = []int{}
			}

			manualRules[beforeValue] = append(manualRules[beforeValue], afterValue)

			continue
		}

		values := strings.Split(lines[i], ",")
		intValues := make([]int, len(values))
		for i := range values {
			intValue, _ := strconv.Atoi(values[i])

			intValues[i] = intValue
		}

		manualPages = append(manualPages, intValues)
	}

	startTime1 := time.Now()
	part1Result := part1(manualRules, manualPages)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(manualRules, manualPages)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(manualRules map[int][]int, manualPages [][]int) int {
	totalMiddlePages := 0

	for i := range manualPages {
		if isPageValid(manualRules, manualPages[i]) {
			totalMiddlePages += manualPages[i][len(manualPages[i])/2]
		}
	}

	return totalMiddlePages
}

func part2(manualRules map[int][]int, manualPages [][]int) int {
	totalMiddlePages := 0

	for i := range manualPages {
		if !isPageValid(manualRules, manualPages[i]) {
			sortedManualPage := sortManualPage(manualRules, manualPages[i])
			totalMiddlePages += sortedManualPage[len(sortedManualPage)/2]
		}
	}

	return totalMiddlePages
}

func isPageValid(manualRules map[int][]int, manualPage []int) bool {
	for i := 1; i < len(manualPage); i++ {
		if rule, ok := manualRules[manualPage[i]]; ok {

			for j, value := range manualPage {
				if j < i && sliceutil.Contains(rule, value) {

					return false
				}

				if j > i && !sliceutil.Contains(rule, value) {
					return false
				}
			}
		}
	}

	return true
}

func sortManualPage(manualRules map[int][]int, manualPage []int) []int {
	sortedManualPage := []int{manualPage[0]}

	for i := 1; i < len(manualPage); i++ {
		valueToPlace := manualPage[i]

		if rule, ok := manualRules[valueToPlace]; ok {
			var err error

			sortedManualPage, err = tryInsertForward(rule, sortedManualPage, valueToPlace)

			if err == nil {
				continue
			}
		}

		for j := len(sortedManualPage) - 1; j >= 0; j-- {
			if rule, ok := manualRules[manualPage[j]]; ok {
				if sliceutil.Contains(rule, valueToPlace) {
					if j == 0 {
						insert, _ := sliceutil.InsertBetween(sortedManualPage, 0, 1, valueToPlace)
						sortedManualPage = insert
						break
					}

					if j == len(sortedManualPage)-1 {
						sortedManualPage = append(sortedManualPage, valueToPlace)
						break
					}

					insert, _ := sliceutil.InsertBetween(sortedManualPage, i-1, i, valueToPlace)
					sortedManualPage = insert
					break
				}
			}
		}

		if !sliceutil.Contains(sortedManualPage, valueToPlace) {
			sortedManualPage = append(sortedManualPage, valueToPlace)
		}
	}
	return sortedManualPage
}

func tryInsertForward(rule []int, sortedManualPage []int, valueToPlace int) ([]int, error) {
	for j, value := range sortedManualPage {
		if sliceutil.Contains(rule, value) {
			if j == 0 {
				sortedManualPage = append([]int{valueToPlace}, sortedManualPage...)
				return sortedManualPage, nil
			}

			if j == len(sortedManualPage)-1 {
				insert, _ := sliceutil.InsertBetween(sortedManualPage, len(sortedManualPage)-2, len(sortedManualPage)-1, valueToPlace)
				sortedManualPage = insert
				return sortedManualPage, nil
			}

			insert, _ := sliceutil.InsertBetween(sortedManualPage, j-1, j, valueToPlace)
			sortedManualPage = insert
			return sortedManualPage, nil
		}
	}

	return sortedManualPage, errors.New("not able to place value")
}
