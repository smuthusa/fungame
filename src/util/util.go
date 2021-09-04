package util

import "errors"

var EmptyArrayError = errors.New("empty array")

func Min(indexes []int) (int, error) {
	return minMax(indexes, min)
}

func Max(indexes []int) (int, error) {
	return minMax(indexes, max)
}

func minMax(indexes []int, minMaxFn func(val1 int, val2 int) int) (int, error) {
	if len(indexes) == 0 {
		return 0, EmptyArrayError
	}
	var minMax int
	first := true
	for _, i := range indexes {
		if first {
			first = false
			minMax = i
		} else {
			minMax = minMaxFn(i, minMax)
		}
	}
	return minMax, nil
}

func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}

func min(val1 int, val2 int) int {
	if val1 < val2 {
		return val1
	} else {
		return val2
	}
}
