package day01

import (
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/laurensotto/advent2024/pkg/intutil"
	"github.com/laurensotto/advent2024/pkg/sliceutil"
)

func Solve(input string) (string, int64, string, int64) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	listOne := make([]int, len(lines))
	listTwo := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		values := strings.Split(lines[i], "   ")

		listOne[i], _ = strconv.Atoi(values[0])
		listTwo[i], _ = strconv.Atoi(values[1])
	}

	copyListOne := sliceutil.DeepCopySlice(listOne)
	copyListTwo := sliceutil.DeepCopySlice(listTwo)

	startTime1 := time.Now()
	part1Result := part1(listOne, listTwo)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(copyListOne, copyListTwo)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(listOne []int, listTwo []int) int {
	slices.Sort(listOne)
	slices.Sort(listTwo)

	totalDifference := 0
	for i := 0; i < len(listOne); i++ {
		totalDifference += intutil.GetDifference(listOne[i], listTwo[i])
	}

	return totalDifference
}

func part2(listOne []int, listTwo []int) int {
	totalSimilarity := 0

	for _, value := range listOne {
		totalSimilarity += getSimilarityScore(value, listTwo)
	}

	return totalSimilarity
}

func getSimilarityScore(int int, values []int) int {
	occurrence := 0

	for _, value := range values {
		if int == value {
			occurrence++
		}
	}

	return int * occurrence
}
