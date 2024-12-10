package day10

import (
	"github.com/laurensotto/advent2024/pkg/gridutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Solve(input string, verbose bool) (string, int64, string, int64) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := gridutil.CreateIntGrid(lines, "")
	copiedGrid := gridutil.DeepCopyGrid(grid)

	startTime1 := time.Now()
	part1Result := part1(grid, verbose)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(copiedGrid, verbose)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(grid [][]int, verbose bool) int {
	var wg sync.WaitGroup
	resultChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, value := range row {
			if value == 0 {
				wg.Add(1)

				go func() {
					defer wg.Done()
					var foundNines []Location
					result, _ := findNextOrFinish(grid, x, y, foundNines)
					resultChan <- result
				}()
			}
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var totalTrailScore int

	for result := range resultChan {
		totalTrailScore += result
	}

	return totalTrailScore
}

func findNextOrFinish(
	grid [][]int,
	xCoordinate,
	yCoordinate int,
	foundNines []Location,
) (int, []Location) {
	currentValue := grid[yCoordinate][xCoordinate]

	if currentValue == 9 {
		for _, nine := range foundNines {
			if xCoordinate == nine.x && yCoordinate == nine.y {
				return 0, foundNines
			}
		}

		foundNines = append(foundNines, Location{xCoordinate, yCoordinate})

		return 1, foundNines
	}

	canCheckLeft := xCoordinate != 0
	canCheckRight := xCoordinate != len(grid[0])-1
	canCheckTop := yCoordinate != 0
	canCheckDown := yCoordinate != len(grid)-1

	var totalResult int
	if canCheckLeft && grid[yCoordinate][xCoordinate-1] == currentValue+1 {
		result, newFoundNines := findNextOrFinish(grid, xCoordinate-1, yCoordinate, foundNines)
		totalResult += result
		foundNines = newFoundNines
	}

	if canCheckRight && grid[yCoordinate][xCoordinate+1] == currentValue+1 {
		result, newFoundNines := findNextOrFinish(grid, xCoordinate+1, yCoordinate, foundNines)
		totalResult += result
		foundNines = newFoundNines
	}

	if canCheckTop && grid[yCoordinate-1][xCoordinate] == currentValue+1 {
		result, newFoundNines := findNextOrFinish(grid, xCoordinate, yCoordinate-1, foundNines)
		totalResult += result
		foundNines = newFoundNines
	}

	if canCheckDown && grid[yCoordinate+1][xCoordinate] == currentValue+1 {
		result, newFoundNines := findNextOrFinish(grid, xCoordinate, yCoordinate+1, foundNines)
		totalResult += result
		foundNines = newFoundNines
	}

	return totalResult, foundNines
}

func part2(grid [][]int, verbose bool) int {
	var wg sync.WaitGroup
	resultChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, value := range row {
			if value == 0 {
				wg.Add(1)

				go func() {
					defer wg.Done()
					resultChan <- findNextOrFinishDistinctRoutes(grid, x, y)
				}()
			}
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var totalTrailScore int

	for result := range resultChan {
		totalTrailScore += result
	}

	return totalTrailScore
}

func findNextOrFinishDistinctRoutes(
	grid [][]int,
	xCoordinate,
	yCoordinate int,
) int {
	currentValue := grid[yCoordinate][xCoordinate]

	if currentValue == 9 {
		return 1
	}

	canCheckLeft := xCoordinate != 0
	canCheckRight := xCoordinate != len(grid[0])-1
	canCheckTop := yCoordinate != 0
	canCheckDown := yCoordinate != len(grid)-1

	var totalResult int
	if canCheckLeft && grid[yCoordinate][xCoordinate-1] == currentValue+1 {
		totalResult += findNextOrFinishDistinctRoutes(grid, xCoordinate-1, yCoordinate)
	}

	if canCheckRight && grid[yCoordinate][xCoordinate+1] == currentValue+1 {
		totalResult += findNextOrFinishDistinctRoutes(grid, xCoordinate+1, yCoordinate)
	}

	if canCheckTop && grid[yCoordinate-1][xCoordinate] == currentValue+1 {
		totalResult += findNextOrFinishDistinctRoutes(grid, xCoordinate, yCoordinate-1)
	}

	if canCheckDown && grid[yCoordinate+1][xCoordinate] == currentValue+1 {
		totalResult += findNextOrFinishDistinctRoutes(grid, xCoordinate, yCoordinate+1)
	}

	return totalResult
}

type Location struct {
	x int
	y int
}
