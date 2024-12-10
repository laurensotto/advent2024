package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/laurensotto/advent2024/internal/day01"
	"github.com/laurensotto/advent2024/internal/day02"
	"github.com/laurensotto/advent2024/internal/day03"
	"github.com/laurensotto/advent2024/internal/day04"
	"github.com/laurensotto/advent2024/internal/day05"
	"github.com/laurensotto/advent2024/internal/day06"
	"github.com/laurensotto/advent2024/internal/day07"
	"github.com/laurensotto/advent2024/internal/day08"
	"github.com/laurensotto/advent2024/internal/day09"
	"github.com/laurensotto/advent2024/internal/day10"
	"github.com/laurensotto/advent2024/internal/day11"
	"github.com/laurensotto/advent2024/internal/day12"
	"github.com/laurensotto/advent2024/internal/day13"
	"github.com/laurensotto/advent2024/internal/day14"
	"github.com/laurensotto/advent2024/internal/day15"
	"github.com/laurensotto/advent2024/internal/day16"
	"github.com/laurensotto/advent2024/internal/day17"
	"github.com/laurensotto/advent2024/internal/day18"
	"github.com/laurensotto/advent2024/internal/day19"
	"github.com/laurensotto/advent2024/internal/day20"
	"github.com/laurensotto/advent2024/internal/day21"
	"github.com/laurensotto/advent2024/internal/day22"
	"github.com/laurensotto/advent2024/internal/day23"
	"github.com/laurensotto/advent2024/internal/day24"
	"github.com/laurensotto/advent2024/internal/day25"
)

const DayCount = 10

func main() {
	var day int
	var runExample bool
	var runVerbose bool

	flag.IntVar(&day, "d", 0, "Run a specific day's code")
	flag.BoolVar(&runExample, "e", false, "Run the example exercise for a specific day")
	flag.BoolVar(&runVerbose, "v", false, "Run the exercise verbose")
	flag.Parse()

	if day != 0 {
		if day < 1 || day > DayCount {
			log.Fatalf("day must be between 1 and %d", DayCount)
		}
		runDay(day, runExample, runVerbose)
	} else {
		runEverything(runExample, runVerbose)
	}
}

func runDay(day int, runExample bool, runVerbose bool) {
	data, err := os.ReadFile(getDayPath(day, runExample))

	if err != nil {
		log.Fatal(err)
	}

	dataString := string(data)
	var part1, part2 string
	var executionTime1, executionTime2 int64

	switch day {
	case 1:
		part1, executionTime1, part2, executionTime2 = day01.Solve(dataString, runVerbose)
	case 2:
		part1, executionTime1, part2, executionTime2 = day02.Solve(dataString, runVerbose)
	case 3:
		part1, executionTime1, part2, executionTime2 = day03.Solve(dataString, runVerbose)
	case 4:
		part1, executionTime1, part2, executionTime2 = day04.Solve(dataString, runVerbose)
	case 5:
		part1, executionTime1, part2, executionTime2 = day05.Solve(dataString, runVerbose)
	case 6:
		part1, executionTime1, part2, executionTime2 = day06.Solve(dataString, runVerbose)
	case 7:
		part1, executionTime1, part2, executionTime2 = day07.Solve(dataString, runVerbose)
	case 8:
		part1, executionTime1, part2, executionTime2 = day08.Solve(dataString, runVerbose)
	case 9:
		part1, executionTime1, part2, executionTime2 = day09.Solve(dataString, runVerbose)
	case 10:
		part1, executionTime1, part2, executionTime2 = day10.Solve(dataString, runVerbose)
	case 11:
		part1, executionTime1, part2, executionTime2 = day11.Solve(dataString, runVerbose)
	case 12:
		part1, executionTime1, part2, executionTime2 = day12.Solve(dataString, runVerbose)
	case 13:
		part1, executionTime1, part2, executionTime2 = day13.Solve(dataString, runVerbose)
	case 14:
		part1, executionTime1, part2, executionTime2 = day14.Solve(dataString, runVerbose)
	case 15:
		part1, executionTime1, part2, executionTime2 = day15.Solve(dataString, runVerbose)
	case 16:
		part1, executionTime1, part2, executionTime2 = day16.Solve(dataString, runVerbose)
	case 17:
		part1, executionTime1, part2, executionTime2 = day17.Solve(dataString, runVerbose)
	case 18:
		part1, executionTime1, part2, executionTime2 = day18.Solve(dataString, runVerbose)
	case 19:
		part1, executionTime1, part2, executionTime2 = day19.Solve(dataString, runVerbose)
	case 20:
		part1, executionTime1, part2, executionTime2 = day20.Solve(dataString, runVerbose)
	case 21:
		part1, executionTime1, part2, executionTime2 = day21.Solve(dataString, runVerbose)
	case 22:
		part1, executionTime1, part2, executionTime2 = day22.Solve(dataString, runVerbose)
	case 23:
		part1, executionTime1, part2, executionTime2 = day23.Solve(dataString, runVerbose)
	case 24:
		part1, executionTime1, part2, executionTime2 = day24.Solve(dataString, runVerbose)
	case 25:
		part1, executionTime1, part2, executionTime2 = day25.Solve(dataString, runVerbose)
	}

	fmt.Printf("Part 1: %s (%d ms)\n", part1, executionTime1)
	fmt.Printf("Part 2: %s (%d ms)\n", part2, executionTime2)
}

func runEverything(runExample bool, runVerbose bool) {
	for day := 1; day <= DayCount; day++ {
		fmt.Printf("Day %d:\n", day)
		runDay(day, runExample, runVerbose)
		fmt.Print("\n")
	}
}

func getDayPath(day int, runExample bool) string {
	var file string
	if runExample {
		file = "example.txt"
	} else {
		file = "challenge.txt"
	}

	return "input/day_" + strconv.Itoa(day) + "_" + file
}
