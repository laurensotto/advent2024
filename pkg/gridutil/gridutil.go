package gridutil

import (
	"strings"

	"github.com/laurensotto/advent2024/pkg/sliceutil"
)

func CreateIntGrid(rows []string, separator string) [][]int {
	grid := make([][]int, len(rows))

	for y, row := range rows {
		intRow := sliceutil.CreateIntSliceFromString(row, separator)

		grid[y] = intRow
	}

	return grid
}

func CreateGrid(rows []string, separator string) [][]string {
	grid := make([][]string, len(rows))

	for y, row := range rows {
		grid[y] = strings.Split(row, separator)
	}

	return grid
}

func DeepCopyGrid[T any](grid [][]T) [][]T {
	copiedGrid := make([][]T, len(grid))
	for i := range grid {
		copiedGrid[i] = sliceutil.DeepCopySlice(grid[i])
	}
	return copiedGrid
}
