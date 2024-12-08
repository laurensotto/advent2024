package day06

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/laurensotto/advent2024/pkg/gridutil"
	"github.com/laurensotto/advent2024/pkg/sliceutil"
)

func Solve(input string, verbose bool) (string, int64, string, int64) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := gridutil.CreateGrid(lines, "")

	copiedGrid := gridutil.DeepCopyGrid(grid)

	startTime1 := time.Now()
	part1Result := part1(grid)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(copiedGrid)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(grid [][]string) int {
	currentX, currentY := findStart(grid)
	direction := "up"
	guardOnGrid := true

	totalVisitedFields := 0

	for guardOnGrid {
		if grid[currentY][currentX] != "X" {
			totalVisitedFields++
			grid[currentY][currentX] = "X"
		}

		currentX, currentY = getNextPosition(currentX, currentY, direction)

		if gridutil.IsOffGrid(currentX, currentY, grid) {
			guardOnGrid = false
			continue
		}

		if grid[currentY][currentX] == "#" {
			currentX, currentY = getPreviousPosition(currentX, currentY, direction)
			direction = rotate(direction)
		}
	}

	return totalVisitedFields
}

func part2(grid [][]string) int {
	xCoordinates, yCoordinates := getRelevantCoordinates(gridutil.DeepCopyGrid(grid))

	var wg sync.WaitGroup
	resultChan := make(chan int, len(grid)*len(grid[0]))

	for i := 0; i < len(xCoordinates); i++ {
		wg.Add(1)
		deepCopiedGrid := gridutil.DeepCopyGrid(grid)
		deepCopiedGrid[yCoordinates[i]][xCoordinates[i]] = "#"

		go func() {
			defer wg.Done()
			resultChan <- solvePart2(deepCopiedGrid)
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalOptions := 0
	for result := range resultChan {
		totalOptions += result
	}

	return totalOptions
}

func getRelevantCoordinates(grid [][]string) ([]int, []int) {
	currentX, currentY := findStart(grid)
	direction := "up"
	guardOnGrid := true

	var xCoordinates []int
	var yCoordinates []int

	for guardOnGrid {
		if grid[currentY][currentX] != "X" {
			xCoordinates = append(xCoordinates, currentX)
			yCoordinates = append(yCoordinates, currentY)
			grid[currentY][currentX] = "X"
		}

		currentX, currentY = getNextPosition(currentX, currentY, direction)

		if gridutil.IsOffGrid(currentX, currentY, grid) {
			guardOnGrid = false
			continue
		}

		if grid[currentY][currentX] == "#" {
			currentX, currentY = getPreviousPosition(currentX, currentY, direction)
			direction = rotate(direction)
		}
	}

	return xCoordinates, yCoordinates
}

func solvePart2(grid [][]string) int {
	currentX, currentY := findStart(grid)
	direction := "up"
	guardOnGrid := true

	visitedCoordinates := make(map[string][]string)

	for guardOnGrid {
		currentCoordinate := strconv.Itoa(currentX) + "," + strconv.Itoa(currentY)

		if _, ok := visitedCoordinates[currentCoordinate]; ok {
			if sliceutil.Contains(visitedCoordinates[currentCoordinate], direction) {
				return 1
			}
		} else {
			visitedCoordinates[currentCoordinate] = []string{}
		}

		visitedCoordinates[currentCoordinate] = append(visitedCoordinates[currentCoordinate], direction)

		currentX, currentY = getNextPosition(currentX, currentY, direction)

		if gridutil.IsOffGrid(currentX, currentY, grid) {
			guardOnGrid = false
			continue
		}

		if grid[currentY][currentX] == "#" {
			currentX, currentY = getPreviousPosition(currentX, currentY, direction)
			direction = rotate(direction)

		}
	}

	return 0
}

func findStart(grid [][]string) (int, int) {
	for y, row := range grid {
		for x, value := range row {
			if value == "^" {
				return x, y
			}
		}
	}
	return 0, 0
}

func getNextPosition(currentX, currentY int, direction string) (int, int) {
	switch direction {
	case "up":
		return currentX, currentY - 1
	case "down":
		return currentX, currentY + 1
	case "left":
		return currentX - 1, currentY
	default:
		return currentX + 1, currentY
	}
}

func getPreviousPosition(currentX, currentY int, direction string) (int, int) {
	switch direction {
	case "up":
		return currentX, currentY + 1
	case "down":
		return currentX, currentY - 1
	case "left":
		return currentX + 1, currentY
	default:
		return currentX - 1, currentY
	}
}

func rotate(direction string) string {
	switch direction {
	case "up":
		return "right"
	case "down":
		return "left"
	case "left":
		return "up"
	default:
		return "down"
	}
}
