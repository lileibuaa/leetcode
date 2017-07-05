package chapter13

func RomanToInteger(s string) int {
	romanMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	resultValue := 0
	for index := range s {
		indexValue := s[index:index+1]
		if index == 0 {
			resultValue += romanMap[indexValue]
			continue
		}
		beforeIndexValue := s[index-1: index]
		if romanMap[beforeIndexValue] < romanMap[indexValue] {
			resultValue += romanMap[indexValue] - 2*romanMap[beforeIndexValue]
		} else {
			resultValue += romanMap[indexValue]
		}
	}
	return resultValue
}
