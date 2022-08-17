package arrary

//https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/

func TwoSum(numbers []int, target int) []int {
	leftIndex := 0
	rightIndex := len(numbers) - 1
	for leftIndex < rightIndex {
		if numbers[leftIndex]+numbers[rightIndex] == target {
			return []int{leftIndex + 1, rightIndex + 1}
		}
		if numbers[leftIndex]+numbers[rightIndex] > target {
			rightIndex--
		} else {
			leftIndex++
		}
	}

	return []int{}
}
