package day09

import (
	"fmt"
	"github.com/laurensotto/advent2024/pkg/sliceutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func Solve(input string, verbose bool) (string, int64, string, int64) {
	splitString := strings.Split(input, "")

	var sliceLength int
	for _, value := range splitString {
		intValue, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		sliceLength += intValue
	}

	diskSlice := createDiskSlice(make([]string, sliceLength), splitString)

	copyDiskSlice := sliceutil.DeepCopySlice(diskSlice)

	if verbose {
		fmt.Println(diskSlice)
	}

	startTime1 := time.Now()
	part1Result := part1(diskSlice, verbose)
	time1Result := time.Since(startTime1).Milliseconds()

	startTime2 := time.Now()
	part2Result := part2(copyDiskSlice, verbose)
	time2Result := time.Since(startTime2).Milliseconds()

	return strconv.Itoa(part1Result), time1Result, strconv.Itoa(part2Result), time2Result
}

func part1(diskSlice []string, verbose bool) int {
	var emptyIndexes []int

	for i, diskSpace := range diskSlice {
		if diskSpace == "." {
			emptyIndexes = append(emptyIndexes, i)
		}
	}

	diskSlice = sortDisk(diskSlice, emptyIndexes)

	if verbose {
		fmt.Println(diskSlice)
	}

	return calculateChecksum(diskSlice)
}

func part2(diskSlice []string, verbose bool) int {
	backwardPivot := len(diskSlice) - 1

	endDiskObject := -1
	var currentCharacter string
	for backwardPivot >= 0 {
		if diskSlice[backwardPivot] != "." && endDiskObject == -1 {
			currentCharacter = diskSlice[backwardPivot]
			endDiskObject = backwardPivot
		}

		if (backwardPivot == 0 || diskSlice[backwardPivot-1] != currentCharacter) && endDiskObject != -1 {
			diskSlice = moveToAvailableSpace(diskSlice, backwardPivot, endDiskObject, verbose)

			endDiskObject = -1
		}

		backwardPivot--
	}

	return calculateChecksum(diskSlice)
}

func createDiskSlice(diskSlice []string, splitString []string) []string {

	var pivot int
	var fileIndex int
	for i, value := range splitString {
		intValue, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		fileIndexString := "."

		if i%2 == 0 {
			fileIndexString = strconv.Itoa(fileIndex)
			fileIndex++
		}

		for diskIndex := 0; diskIndex < intValue; diskIndex++ {
			diskSlice[pivot+diskIndex] = fileIndexString
		}

		pivot += intValue
	}

	return diskSlice
}

func moveToAvailableSpace(diskSlice []string, startIndex int, endIndex int, verbose bool) []string {
	requiredLength := endIndex - startIndex + 1

	startEmptyIndex := -1
	for i, value := range diskSlice {
		if i >= startIndex {
			break
		}

		if value == "." && startEmptyIndex == -1 {
			startEmptyIndex = i
		}

		if (i == len(diskSlice)-1 || diskSlice[i+1] != ".") && startEmptyIndex != -1 {

			if i-startEmptyIndex+1 < requiredLength {
				startEmptyIndex = -1
				continue
			}

			for j := startEmptyIndex; j < startEmptyIndex+requiredLength; j++ {
				diskSlice[j] = diskSlice[startIndex]
			}

			for j := startIndex; j <= endIndex; j++ {
				diskSlice[j] = "."
			}

			break
		}
	}

	return diskSlice
}

func sortDisk(diskSlice []string, emptyIndexes []int) []string {
	pivot := len(diskSlice) - 1

	stopIndex := len(diskSlice) - len(emptyIndexes)

	for _, index := range emptyIndexes {
		if index >= stopIndex {
			break
		}

		diskSlice[index] = diskSlice[pivot]
		diskSlice[pivot] = "."

		for diskSlice[pivot] == "." {
			pivot--
		}
	}

	return diskSlice
}

func calculateChecksum(diskSlice []string) int {
	var checksum int

	for i, diskValue := range diskSlice {
		if diskValue == "." {
			continue
		}

		diskIntValue, err := strconv.Atoi(diskValue)

		if err != nil {
			log.Fatal(err)
		}

		checksum += i * diskIntValue
	}

	return checksum
}
