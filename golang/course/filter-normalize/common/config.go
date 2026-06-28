package common

import "math"

type Config struct {
	MaxFilterLimit int
}

var Instance = Config{
	MaxFilterLimit: 5,
}

func Abs(value int) int {
	return int(math.Abs(float64(value)))
}
