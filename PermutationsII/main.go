package PermutationsII

func permuteUnique(nums []int) [][]int {
	// 使用 map 記錄每個元素的出現次數
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	result := [][]int{}
	path := []int{}

	// 提取不同的數字和它們的出現次數
	var uniqueNums []int
	var counts []int
	for num, count := range counter {
		uniqueNums = append(uniqueNums, num)
		counts = append(counts, count)
	}

	backtrackWithCounter(uniqueNums, counts, len(nums), 0, path, &result)

	return result
}

func backtrackWithCounter(nums []int, counts []int, n, pos int, path []int, result *[][]int) {
	if pos == n {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}

	for i := 0; i < len(nums); i++ {
		if counts[i] == 0 {
			continue
		}

		counts[i]--
		path = append(path, nums[i])

		backtrackWithCounter(nums, counts, n, pos+1, path, result)

		path = path[:len(path)-1]
		counts[i]++
	}
}
