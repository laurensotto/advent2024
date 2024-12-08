package sliceutil

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

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

func DeepCopySlice[T any](slice []T) []T {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)

	return copiedSlice
}

func CreateIntSliceFromString(string string, separator string) []int {
	separatedString := strings.Split(string, separator)
	intSlice := make([]int, len(separatedString))

	for i, str := range separatedString {
		intValue, err := strconv.Atoi(str)

		if err != nil {
			log.Fatalf("can not convert string %s to int", str)
		}

		intSlice[i] = intValue
	}

	return intSlice
}
