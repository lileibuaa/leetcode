package chapter11

func MaxArea(height []int) int {
	sum := 0
	for i, j := 0, len(height)-1; i < j; {
		iValue, jValue := height[i], height[j]
		minValue := iValue
		if jValue < minValue {
			minValue = jValue
		}
		tmpSum := (j - i) * minValue
		if tmpSum > sum {
			sum = tmpSum
		}
		if iValue > jValue {
			j--
		} else {
			i++
		}
	}
	return sum
}
