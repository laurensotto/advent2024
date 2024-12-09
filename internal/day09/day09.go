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
	var emptyDiskSpace []DiskSpace

	for i := 0; i < len(diskSlice); i++ {
		if diskSlice[i] == "." {
			startEmptySpace := i

			for j := i + 1; j < len(diskSlice); j++ {
				if diskSlice[j] != "." {
					endEmptySpace := j

					emptyDiskSpace = append(emptyDiskSpace, DiskSpace{
						startEmptySpace,
						endEmptySpace - startEmptySpace,
					})

					i = j - 1 // Make sure to double-check this afterwards.
					break
				}
			}
		}
	}

	var filledDiskSpace []DiskSpace

	for i := 0; i < len(diskSlice); i++ {
		foundCharacter := diskSlice[i]
		if foundCharacter != "." {
			startEmptySpace := i

			for j := i + 1; j <= len(diskSlice); j++ {
				if j == len(diskSlice) || diskSlice[j] != foundCharacter {
					endEmptySpace := j

					filledDiskSpace = append(filledDiskSpace, DiskSpace{
						startEmptySpace,
						endEmptySpace - startEmptySpace,
					})

					i = j - 1 // Make sure to double-check this afterwards.
					break
				}
			}
		}
	}

	for filledIndex := len(filledDiskSpace) - 1; filledIndex >= 0; filledIndex-- {
		requiredLength := filledDiskSpace[filledIndex].blocks
		startIndex := filledDiskSpace[filledIndex].startIndex
		endIndex := startIndex + requiredLength - 1
		character := diskSlice[startIndex]

		fmt.Println(character)

		for i, value := range emptyDiskSpace {
			availableLength := value.blocks
			emptyStartIndex := value.startIndex

			if availableLength >= requiredLength {

				endRequiredIndex := emptyStartIndex + requiredLength
				for emptyStartIndex < endRequiredIndex {
					diskSlice[emptyStartIndex] = character
					emptyStartIndex++
				}

				for startIndex <= endIndex {
					diskSlice[startIndex] = "."
					startIndex++
				}

				emptyDiskSpace[i] = DiskSpace{
					endRequiredIndex,
					value.blocks + value.startIndex - endRequiredIndex,
				}

				break
			}
		}
	}

	if verbose {
		fmt.Println(diskSlice)
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
