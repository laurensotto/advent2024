package sliceutil

import "errors"

func Contains[T int | string](slice []T, valueToFind T) bool {
	for _, value := range slice {
		if value == valueToFind {
			return true
		}
	}

	return false
}

func InsertBetween[T any](array []T, index1 int, index2 int, value T) ([]T, error) {
	if index2 != index1+1 {
		return array, errors.New("index1 and index2 are not neighbours")
	}

	newArray := make([]T, 0, len(array)+1)
	newArray = append(newArray, array[:index2]...)
	newArray = append(newArray, value)
	newArray = append(newArray, array[index2:]...)

	return newArray, nil
}

func DeepCopyGrid[T any](grid [][]T) [][]T {
	copiedGrid := make([][]T, len(grid))
	for i := range grid {
		copiedGrid[i] = DeepCopySlice(grid[i])
	}
	return copiedGrid
}

func DeepCopySlice[T any](slice []T) []T {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)

	return copiedSlice
}
