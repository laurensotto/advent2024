package day08

import (
	"fmt"
	"github.com/laurensotto/advent2024/pkg/gridutil"
	"github.com/laurensotto/advent2024/pkg/sliceutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Solve(input string, verbose bool) (string, int64, string, int64) {
	lines := strings.Split(input, "\n")

	grid := gridutil.CreateGrid(lines, "")

	startTime1 := time.Now()
	part1Result := part1(grid, verbose)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(grid, verbose)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(grid [][]string, verbose bool) int {
	var wg sync.WaitGroup
	var handledCharacters []string
	antiNodes := gridutil.DeepCopyGrid(grid)

	for y, row := range grid {
		for x, cell := range row {
			if cell != "." && !sliceutil.Contains(handledCharacters, cell) {
				wg.Add(1)
				handledCharacters = append(handledCharacters, cell)
				antenna := Antenna{
					cell,
					Location{x, y},
				}

				go handleCharacterPart1(grid, antiNodes, antenna, &wg)
			}
		}
	}

	wg.Wait()

	var foundAntiNodes int
	for _, row := range antiNodes {
		if verbose {
			fmt.Println(row)
		}
		for _, cell := range row {
			if cell == "#" {
				foundAntiNodes++
			}
		}
	}

	return foundAntiNodes
}

func part2(grid [][]string, verbose bool) int {
	var wg sync.WaitGroup
	var handledCharacters []string
	antiNodes := gridutil.DeepCopyGrid(grid)

	for y, row := range grid {
		for x, cell := range row {
			if cell != "." && !sliceutil.Contains(handledCharacters, cell) {
				wg.Add(1)
				handledCharacters = append(handledCharacters, cell)
				antenna := Antenna{
					cell,
					Location{x, y},
				}

				go handleCharacterPart2(grid, antiNodes, antenna, &wg)
			}
		}
	}

	wg.Wait()

	var foundAntiNodes int
	for _, row := range antiNodes {
		if verbose {
			fmt.Println(row)
		}
		for _, cell := range row {
			if cell != "." {
				foundAntiNodes++
			}
		}
	}

	return foundAntiNodes
}

func handleCharacterPart1(grid [][]string, antiNodes [][]string, initialAntenna Antenna, wg *sync.WaitGroup) {
	defer wg.Done()

	var sameTypeAntennas []Antenna

	for y, row := range grid {
		for x, cell := range row {
			if cell == initialAntenna.character {
				sameTypeAntennas = append(sameTypeAntennas, Antenna{
					initialAntenna.character,
					Location{x, y},
				})
			}
		}
	}

	for _, antenna := range sameTypeAntennas {
		for _, matchingAntenna := range sameTypeAntennas {
			if antenna.equals(matchingAntenna) {
				continue
			}

			antiNodeX := antenna.location.x() + antenna.location.x() - matchingAntenna.location.x()
			antiNodeY := antenna.location.y() + antenna.location.y() - matchingAntenna.location.y()

			if gridutil.IsOffGrid(antiNodeX, antiNodeY, antiNodes) {
				continue
			}

			antiNodes[antiNodeY][antiNodeX] = "#"
		}
	}
}

func handleCharacterPart2(grid [][]string, antiNodes [][]string, initialAntenna Antenna, wg *sync.WaitGroup) {
	defer wg.Done()

	var sameTypeAntennas []Antenna

	for y, row := range grid {
		for x, cell := range row {
			if cell == initialAntenna.character {
				sameTypeAntennas = append(sameTypeAntennas, Antenna{
					initialAntenna.character,
					Location{x, y},
				})
			}
		}
	}

	for _, antenna := range sameTypeAntennas {
		for _, matchingAntenna := range sameTypeAntennas {
			if antenna.equals(matchingAntenna) {
				continue
			}

			xDiff := antenna.location.x() - matchingAntenna.location.x()
			yDiff := antenna.location.y() - matchingAntenna.location.y()

			antiNodeX := antenna.location.x() + xDiff
			antiNodeY := antenna.location.y() + yDiff

			for !gridutil.IsOffGrid(antiNodeX, antiNodeY, antiNodes) {
				antiNodes[antiNodeY][antiNodeX] = "#"

				antiNodeX += xDiff
				antiNodeY += yDiff
			}
		}
	}
}
