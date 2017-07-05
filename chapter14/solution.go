package chapter14

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length <= 0 {
		return ""
	}
	for i := 0; ; i++ {
		var targetChar byte
		for j := 0; j < length; j++ {
			if len(strs[j]) <= i {
				return strs[j]
			}
			if j == 0 {
				targetChar = strs[j][i]
			} else if strs[j][i] != targetChar {
				return strs[0][:i]
			}
		}
	}
}
