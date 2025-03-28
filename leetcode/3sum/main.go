package threesum

import "sort"

func threeSum(nums []int) [][]int {
	// 先對數組進行排序，便於處理重複元素和使用雙指針
	sort.Ints(nums)

	result := [][]int{}
	n := len(nums)

	// 遍歷數組，固定第一個元素
	for i := 0; i < n-2; i++ {
		// 如果當前元素與前一個元素相同，跳過以避免重複
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 如果最小的三個數之和大於0，後面不可能有和為0的三元組
		if nums[i] > 0 {
			break
		}

		// 使用雙指針尋找剩下的兩個數
		left, right := i+1, n-1
		target := -nums[i] // 目標是找到兩個數的和等於-nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum < target {
				// 如果和小於目標值，增加左指針
				left++
			} else if sum > target {
				// 如果和大於目標值，減小右指針
				right--
			} else {
				// 找到一個解，記錄結果
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 跳過重複的元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 移動指針尋找新的解
				left++
				right--
			}
		}
	}

	return result
}
