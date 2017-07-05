package chapter1

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for index, value := range nums {
		if mapValue, ok := numMap[target-value]; ok {
			return []int{index, mapValue}
		}
		numMap[value] = index
	}
	return nil
}
