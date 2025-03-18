package threesum

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	// 處理邊界情況
	n := len(nums)
	if n < 3 {
		return result
	}

	// 排序數組
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// 跳過重複的第一個數
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 使用雙指針尋找剩下的兩個數
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				// 找到一組解
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 移動指針
				left++
				right--

				// 跳過左側重複元素
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// 跳過右側重複元素
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}
