package fourSum

import "sort"

func fourSum(nums []int, target int) [][]int {
	// 排序數組
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}

	// 特殊情況：數組長度不足4
	if n < 4 {
		return result
	}

	// 使用兩層循環固定前兩個元素
	for i := 0; i < n-3; i++ {
		// 跳過重複的第一個元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 最小可能和大於目標值，後面的元素更大，不可能找到解
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		// 當前元素與最大的三個元素和小於目標值，當前元素太小，嘗試下一個
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// 跳過重複的第二個元素
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 剪枝：判斷最小可能和
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}

			// 剪枝：判斷最大可能和
			if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}

			// 使用雙指針查找剩餘兩個元素
			left, right := j+1, n-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					// 找到一個四元組
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					// 跳過重複元素
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					// 移動雙指針
					left++
					right--
				}
			}
		}
	}

	return result
}
