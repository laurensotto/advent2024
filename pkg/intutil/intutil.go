package intutil

func GetDifference(int1 int, int2 int) int {
	if int1 > int2 {
		return int1 - int2
	}

	return int2 - int1
}
