package chapter127

func LadderLength(beginWord string, endWord string, wordList []string) int {
	init := true
	startTarget := []string{}
	endTarget := []string{}
	alternativeTarget := make([]string, len(wordList))
	length := 1
	for {
		length++
		if init {
			containsEnd := false
			for i, value := range wordList {
				if endWord == value {
					containsEnd = true
				}
				alternativeTarget[i] = value
			}
			if !containsEnd {
				return 0
			}
			init = false
			startTarget = append(startTarget, beginWord)
		}
		if len(startTarget) <= 0 || len(alternativeTarget) <= 0 {
			return 0
		}
		alternativeTarget2 := []string{}
	outTag:
		for m := range alternativeTarget {
			for n := range startTarget {
				if hasPath(alternativeTarget[m], startTarget[n]) {
					if alternativeTarget[m] == endWord {
						return length
					}
					endTarget = append(endTarget, alternativeTarget[m])
					continue outTag
				}
			}
			alternativeTarget2 = append(alternativeTarget2, alternativeTarget[m])
		}
		startTarget = endTarget
		endTarget = []string{}
		alternativeTarget = alternativeTarget2
	}
	return 0
}

func hasPath(start string, end string) bool {
	diffCount := 0
	for index := range start {
		if start[index] == end[index] {
			continue
		} else {
			diffCount++
		}
		if diffCount > 1 {
			return false
		}
	}
	return true
}
