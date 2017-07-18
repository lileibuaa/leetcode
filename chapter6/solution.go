package chapter6

func Convert(s string, numRows int) string {
	strLen := len(s)
	if strLen == 0 {
		return ""
	} else if strLen == 1 || strLen == 2 {
		return s
	}
	if numRows == 1 {
		return s
	}
	sourceByte := []byte(s)
	targetByte := make([]byte, strLen)
	targetIndex := 0
	for i := 0; i < numRows; i++ {
		j := 0
		oddNum := true
		var preIndex int = i
		for {
			var nextIndex int = 0
			if i == 0 || i == numRows-1 {
				nextIndex = i + (numRows-1)*2*j
			} else {
				if j == 0 {
					nextIndex = i
				} else if oddNum {
					nextIndex = preIndex + i*2
				} else {
					nextIndex = preIndex + (numRows-i-1)*2
				}
			}
			oddNum = !oddNum
			if nextIndex >= strLen {
				break
			}
			preIndex = nextIndex
			targetByte[targetIndex] = sourceByte[nextIndex]
			targetIndex++
			j++
		}
	}
	return string(targetByte)
}
