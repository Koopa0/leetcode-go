package RemoveElement

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] == val {
			// 將左指針指向的元素與右指針指向的元素交換
			nums[left] = nums[right]
			right--
		} else {
			// 左指針指向的元素不等於 val，向右移動
			left++
		}
	}

	return left
}
