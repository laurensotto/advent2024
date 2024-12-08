package day04

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/laurensotto/advent2024/pkg/gridutil"
)

func Solve(input string) (string, int64, string, int64) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := gridutil.CreateGrid(lines, "")

	startTime1 := time.Now()
	part1Result := part1(grid)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(grid)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(grid [][]string) int {
	var wg sync.WaitGroup
	resultChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, cell := range row {
			if cell == "X" {
				wg.Add(1)
				findXmas(x, y, grid, &wg, resultChan)
			}
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalXmas := 0
	for result := range resultChan {
		totalXmas += result
	}

	return totalXmas
}

func part2(grid [][]string) int {
	var wg sync.WaitGroup
	resultChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, cell := range row {
			if cell == "A" {
				wg.Add(1)
				findCrossmas(x, y, grid, &wg, resultChan)
			}
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalXmas := 0
	for result := range resultChan {
		totalXmas += result
	}

	return totalXmas
}

func findXmas(startX int, startY int, grid [][]string, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	needToCheckLeft := startX > 2
	needToCheckRight := startX < len(grid[0])-3
	needToCheckTop := startY > 2
	needToCheckBottom := startY < len(grid)-3

	if needToCheckTop {
		wg.Add(1)
		go findCharacter(startX, startY, 0, -1, grid, "M", wg, resultChan)
	}

	if needToCheckBottom {
		wg.Add(1)
		go findCharacter(startX, startY, 0, 1, grid, "M", wg, resultChan)
	}

	if needToCheckLeft {
		wg.Add(1)
		go findCharacter(startX, startY, -1, 0, grid, "M", wg, resultChan)

		if needToCheckTop {
			wg.Add(1)
			go findCharacter(startX, startY, -1, -1, grid, "M", wg, resultChan)
		}

		if needToCheckBottom {
			wg.Add(1)
			go findCharacter(startX, startY, -1, 1, grid, "M", wg, resultChan)
		}
	}

	if needToCheckRight {
		wg.Add(1)
		go findCharacter(startX, startY, 1, 0, grid, "M", wg, resultChan)

		if needToCheckTop {
			wg.Add(1)
			go findCharacter(startX, startY, 1, -1, grid, "M", wg, resultChan)
		}

		if needToCheckBottom {
			wg.Add(1)
			go findCharacter(startX, startY, 1, 1, grid, "M", wg, resultChan)
		}
	}
}

func findCrossmas(startX int, startY int, grid [][]string, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	canCheckLeft := startX > 0
	canCheckRight := startX < len(grid[0])-1
	canCheckTop := startY > 0
	canCheckBottom := startY < len(grid)-1

	if !(canCheckLeft && canCheckRight && canCheckTop && canCheckBottom) {
		return
	}

	if grid[startY-1][startX-1] == "A" ||
		grid[startY-1][startX+1] == "A" ||
		grid[startY-1][startX-1] == "X" ||
		grid[startY-1][startX+1] == "X" {
		return
	}

	if grid[startY-1][startX-1] == "M" {
		if grid[startY+1][startX+1] != "S" {
			return
		}
	}

	if grid[startY-1][startX-1] == "S" {
		if grid[startY+1][startX+1] != "M" {
			return
		}
	}

	if grid[startY-1][startX+1] == "M" {
		if grid[startY+1][startX-1] != "S" {
			return
		}
	}

	if grid[startY-1][startX+1] == "S" {
		if grid[startY+1][startX-1] != "M" {
			return
		}
	}

	resultChan <- 1
}

func findCharacter(
	startX int,
	startY int,
	directionX int,
	directionY int,
	grid [][]string,
	character string,
	wg *sync.WaitGroup,
	resultChan chan int,
) {
	defer wg.Done()

	if grid[startY+directionY][startX+directionX] == character {
		switch character {
		case "M":
			wg.Add(1)
			go findCharacter(
				startX+directionX,
				startY+directionY,
				directionX,
				directionY,
				grid,
				"A",
				wg,
				resultChan,
			)
		case "A":
			wg.Add(1)
			go findCharacter(
				startX+directionX,
				startY+directionY,
				directionX,
				directionY,
				grid,
				"S",
				wg,
				resultChan,
			)
		case "S":
			resultChan <- 1
		}
	}
}
