package SortColors

func sortColors(nums []int) {
	low, curr, high := 0, 0, len(nums)-1

	// 當 curr 指針超過 high 指針時，排序完成
	for curr <= high {
		switch nums[curr] {
		case 0:
			// 遇到0，與low位置交換，low和curr都右移
			nums[low], nums[curr] = nums[curr], nums[low]
			low++
			curr++
		case 1:
			// 遇到1，保持位置不變，curr右移
			curr++
		case 2:
			// 遇到2，與high位置交換，high左移
			// 注意：此時curr不右移，因為交換後的值還未處理
			nums[curr], nums[high] = nums[high], nums[curr]
			high--
		}
	}
}
