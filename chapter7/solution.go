package chapter7

import "math"

func ReverseInteger(x int) int {
	var result int64 = 0
	target := x
	for target != 0 {
		result = result*10 + int64(target%10)
		target /= 10
	}
	if math.MinInt32 > result || math.MaxInt32 < result {
		return 0
	}
	return int(result)
}
