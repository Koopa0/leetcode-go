package TwoSum

// TwoSum 使用雜湊表找出陣列中兩數之和等於目標值的索引
// 時間複雜度: O(n)，其中 n 是陣列長度
// 空間複雜度: O(n)，用於儲存雜湊表
func TwoSum(nums []int, target int) []int {
	// 建立雜湊表，key 為數值，value 為索引
	m := make(map[int]int, len(nums))

	// 遍歷陣列
	for i, num := range nums {
		// 檢查目標值減去當前值是否存在於雜湊表中
		if j, ok := m[target-num]; ok {
			// 找到符合條件的兩數，返回它們的索引
			return []int{j, i}
		}
		// 將當前數值和索引存入雜湊表
		m[num] = i
	}
	// 未找到符合條件的兩數
	return nil
}
