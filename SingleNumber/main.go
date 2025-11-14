package main

// singleNumber 使用 XOR 位元運算找出只出現一次的數字
// XOR 運算的關鍵性質：
// 1. a ^ a = 0（相同的數字 XOR 結果為 0）
// 2. a ^ 0 = a（任何數字與 0 XOR 結果為自己）
// 3. XOR 滿足交換律和結合律：a ^ b ^ a = (a ^ a) ^ b = 0 ^ b = b
//
// 因此，將陣列中所有數字進行 XOR，成對的數字會相互抵消，
// 最後剩下的就是只出現一次的數字。
//
// 時間複雜度: O(n)，其中 n 是陣列長度
// 空間複雜度: O(1)，只使用一個變數
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// singleNumberWithMap 使用雜湊表的替代解法（用於對比）
// 時間複雜度: O(n)
// 空間複雜度: O(n)
func singleNumberWithMap(nums []int) int {
	count := make(map[int]int)

	// 統計每個數字出現的次數
	for _, num := range nums {
		count[num]++
	}

	// 找出只出現一次的數字
	for num, cnt := range count {
		if cnt == 1 {
			return num
		}
	}

	return 0
}

// singleNumberWithSet 使用集合的數學解法
// 利用 2*(a+b+c) - (a+a+b+b+c) = c 的原理
// 時間複雜度: O(n)
// 空間複雜度: O(n)
func singleNumberWithSet(nums []int) int {
	set := make(map[int]bool)
	sumUnique := 0
	sumAll := 0

	for _, num := range nums {
		sumAll += num
		if !set[num] {
			set[num] = true
			sumUnique += num
		}
	}

	return 2*sumUnique - sumAll
}

func main() {
	// 範例 1: [2,2,1]
	nums1 := []int{2, 2, 1}
	println("範例 1: [2,2,1]")
	println("結果:", singleNumber(nums1)) // 1

	// 範例 2: [4,1,2,1,2]
	nums2 := []int{4, 1, 2, 1, 2}
	println("\n範例 2: [4,1,2,1,2]")
	println("結果:", singleNumber(nums2)) // 4

	// 範例 3: [1]
	nums3 := []int{1}
	println("\n範例 3: [1]")
	println("結果:", singleNumber(nums3)) // 1

	// XOR 運算演示
	println("\nXOR 運算演示:")
	println("2 ^ 2 =", 2^2)     // 0
	println("1 ^ 0 =", 1^0)     // 1
	println("4 ^ 1 ^ 2 ^ 1 ^ 2 =", 4^1^2^1^2) // 4
}
