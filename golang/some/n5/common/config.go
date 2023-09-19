package common

import "math"

type Config struct {
	MAX_FILTER_LIMIT int
}

var Instance = Config{
	MAX_FILTER_LIMIT: 5,
}

func Abs(value int) int {
	return int(
		math.Abs(
			float64(value)))
}
