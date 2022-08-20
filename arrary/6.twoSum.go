package arrary

//https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {
	leftIndex, rightIndex := 0, len(numbers)-1
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

	return []int{-1, -1}
}

//二分法
func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		n := target - numbers[i]
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := low + (high-low)/2
			if n == numbers[mid] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] < n {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return []int{-1, -1}
}
