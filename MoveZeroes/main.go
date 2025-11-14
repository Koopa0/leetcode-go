package main

// moveZeroes 使用雙指標將所有零移到陣列末尾，同時保持非零元素的相對順序
// 使用兩個指標：
// - left: 指向下一個非零元素應該放置的位置
// - right: 遍歷陣列，尋找非零元素
// 時間複雜度: O(n)，其中 n 是陣列長度
// 空間複雜度: O(1)，原地修改
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// left 指標：下一個非零元素要放置的位置
	left := 0

	// right 指標：遍歷陣列
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			// 找到非零元素，移動到 left 位置
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	// 此時 left 之前都是非零元素（按原順序）
	// left 之後的元素已經在交換過程中變成了 0
}

// moveZeroesOptimized 優化版本：只在必要時交換
// 避免不必要的自我交換（當 left == right 時）
// 時間複雜度: O(n)
// 空間複雜度: O(1)
func moveZeroesOptimized(nums []int) {
	if len(nums) <= 1 {
		return
	}

	left := 0

	// 第一遍：將所有非零元素移到前面
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}

	// 第二遍：將剩餘位置填充為 0
	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}

// moveZeroesBruteForce 暴力法（用於對比）
// 創建新陣列，先放非零元素，再填充零
// 時間複雜度: O(n)
// 空間複雜度: O(n)
func moveZeroesBruteForce(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// 計算零的數量
	zeroCount := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		}
	}

	// 創建臨時陣列存儲非零元素
	temp := make([]int, 0, len(nums)-zeroCount)
	for _, num := range nums {
		if num != 0 {
			temp = append(temp, num)
		}
	}

	// 複製回原陣列
	copy(nums, temp)

	// 填充剩餘位置為 0
	for i := len(temp); i < len(nums); i++ {
		nums[i] = 0
	}
}

// moveZeroesSnowball 雪球法（有趣的視角）
// 把所有的零想像成一個雪球，不斷往後滾
// 時間複雜度: O(n)
// 空間複雜度: O(1)
func moveZeroesSnowball(nums []int) {
	if len(nums) <= 1 {
		return
	}

	snowballSize := 0 // 雪球大小（零的數量）

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// 遇到零，雪球變大
			snowballSize++
		} else if snowballSize > 0 {
			// 遇到非零元素且有雪球，交換
			// 相當於非零元素"穿過"雪球
			nums[i-snowballSize] = nums[i]
			nums[i] = 0
		}
	}
}

func main() {
	// 範例 1: [0,1,0,3,12]
	nums1 := []int{0, 1, 0, 3, 12}
	println("範例 1: [0,1,0,3,12]")
	print("處理前: ")
	for _, v := range nums1 {
		print(v, " ")
	}
	println()

	moveZeroes(nums1)
	print("處理後: ")
	for _, v := range nums1 {
		print(v, " ")
	}
	println() // 預期: [1,3,12,0,0]

	// 範例 2: [0]
	nums2 := []int{0}
	println("\n範例 2: [0]")
	print("處理前: ")
	for _, v := range nums2 {
		print(v, " ")
	}
	println()

	moveZeroes(nums2)
	print("處理後: ")
	for _, v := range nums2 {
		print(v, " ")
	}
	println() // 預期: [0]

	// 範例 3: [1,2,3]（沒有零）
	nums3 := []int{1, 2, 3}
	println("\n範例 3: [1,2,3]")
	print("處理前: ")
	for _, v := range nums3 {
		print(v, " ")
	}
	println()

	moveZeroes(nums3)
	print("處理後: ")
	for _, v := range nums3 {
		print(v, " ")
	}
	println() // 預期: [1,2,3]

	// 範例 4: [0,0,1]
	nums4 := []int{0, 0, 1}
	println("\n範例 4: [0,0,1]")
	print("處理前: ")
	for _, v := range nums4 {
		print(v, " ")
	}
	println()

	moveZeroes(nums4)
	print("處理後: ")
	for _, v := range nums4 {
		print(v, " ")
	}
	println() // 預期: [1,0,0]

	// 雙指標演示
	println("\n\n雙指標演示（範例 1: [0,1,0,3,12]）:")
	nums := []int{0, 1, 0, 3, 12}
	left := 0

	print("初始: ")
	for _, v := range nums {
		print(v, " ")
	}
	println()

	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			print("right=", right, ", left=", left, ", 交換後: ")
			for _, v := range nums {
				print(v, " ")
			}
			println()
			left++
		}
	}

	print("最終: ")
	for _, v := range nums {
		print(v, " ")
	}
	println()
}
