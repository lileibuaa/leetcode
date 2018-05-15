package chapter8

import "math"

func MyAtoi(str string) int {
	srcByte := []byte(str)
	factor := int64(1)
	baseByte, upperByte := int64('0'), int64('9')
	addSymbol, minusSymbol := int64('+'), int64('-')
	validStart := false //为true表示已经有有效的数字
	spaceSymbol := int64(' ')
	var sum int64 = 0
	for _, tmpByte := range srcByte {
		byteValue := int64(tmpByte)
		if byteValue >= baseByte && byteValue <= upperByte {
			sum = sum*10 + (byteValue - baseByte)
			validStart = true
		} else if byteValue == minusSymbol {
			if validStart {
				return int64ToInt(sum * factor)
			} else {
				validStart = true
				factor = -1
			}
		} else if byteValue == addSymbol {
			if validStart {
				return int64ToInt(sum * factor)
			} else {
				validStart = true
			}
		} else if !validStart && spaceSymbol == byteValue {
			continue
		} else {
			return int64ToInt(sum * factor)
		}
		if sum > math.MaxInt32 || sum < math.MinInt32 {
			return int64ToInt(sum * factor)
		}
	}
	return int64ToInt(sum * factor)
}

func int64ToInt(value int64) int {
	if value > math.MaxInt32 {
		return math.MaxInt32
	} else if value < math.MinInt32 {
		return math.MinInt32
	} else {
		return int(value)
	}
}
