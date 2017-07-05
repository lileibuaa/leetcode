package chapter20

func ValidParentheses(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	byteA1, byteA2 := byte('('), byte(')')
	byteB1, byteB2 := byte('['), byte(']')
	byteC1, byteC2 := byte('{'), byte('}')
	var byteStack []byte
	for i := 0; i < length; i++{
		switch s[i] {
		case byteA1:
			byteStack = append(byteStack, s[i])
		case byteB1:
			byteStack = append(byteStack, s[i])
		case byteC1:
			byteStack = append(byteStack, s[i])
		case byteA2:
			if len(byteStack) <= 0 {
				return false
			} else if byteStack[len(byteStack)-1] == byteA1 {
				byteStack = byteStack[:len(byteStack)-1]
			} else {
				return false
			}
		case byteB2:
			if len(byteStack) <= 0 {
				return false
			} else if byteStack[len(byteStack)-1] == byteB1 {
				byteStack = byteStack[:len(byteStack)-1]
			} else {
				return false
			}
		case byteC2:
			if len(byteStack) <= 0 {
				return false
			} else if byteStack[len(byteStack)-1] == byteC1 {
				byteStack = byteStack[:len(byteStack)-1]
			} else {
				return false
			}
		}
	}
	return len(byteStack) == 0
}
