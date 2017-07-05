package chapter9

func PalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	var tmpValue = x
	var newValue = 0
	for tmpValue > 0 {
		newValue = newValue*10 + tmpValue%10
		tmpValue = tmpValue / 10
	}
	return newValue == x
}
