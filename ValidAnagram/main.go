package main

import (
	"sort"
)

// isAnagram 使用陣列計數法判斷是否為字母異位詞（最優解法）
// 字母異位詞：兩個字串包含相同的字元，只是順序不同
// 使用固定大小的陣列（26個字母）統計每個字母出現次數
// 時間複雜度: O(n)，其中 n 是字串長度
// 空間複雜度: O(1)，陣列大小固定為 26
func isAnagram(s string, t string) bool {
	// 長度不同，肯定不是字母異位詞
	if len(s) != len(t) {
		return false
	}

	// 使用陣列統計每個字母出現的次數
	// count[0] 對應 'a'，count[1] 對應 'b'，以此類推
	var count [26]int

	// 遍歷兩個字串
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++ // s 中的字元計數加 1
		count[t[i]-'a']-- // t 中的字元計數減 1
	}

	// 檢查所有計數是否都為 0
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}

// isAnagramHashMap 使用雜湊表計數法
// 適用於 Unicode 字元（不限於小寫英文字母）
// 時間複雜度: O(n)
// 空間複雜度: O(k)，其中 k 是字元集大小
func isAnagramHashMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	// 統計 s 中每個字元的出現次數
	for _, ch := range s {
		count[ch]++
	}

	// 減去 t 中每個字元的出現次數
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	// 檢查所有計數是否都為 0
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}

// isAnagramSorting 使用排序法
// 將兩個字串排序後比較是否相等
// 時間複雜度: O(n log n)
// 空間複雜度: O(n)，排序需要額外空間
func isAnagramSorting(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 將字串轉換為 byte slice 並排序
	sBytes := []byte(s)
	tBytes := []byte(t)

	sort.Slice(sBytes, func(i, j int) bool {
		return sBytes[i] < sBytes[j]
	})

	sort.Slice(tBytes, func(i, j int) bool {
		return tBytes[i] < tBytes[j]
	})

	// 比較排序後的字串
	return string(sBytes) == string(tBytes)
}

// isAnagramTwoMaps 使用兩個雜湊表分別統計
// 時間複雜度: O(n)
// 空間複雜度: O(k)
func isAnagramTwoMaps(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make(map[rune]int)
	countT := make(map[rune]int)

	// 統計 s 中每個字元的出現次數
	for _, ch := range s {
		countS[ch]++
	}

	// 統計 t 中每個字元的出現次數
	for _, ch := range t {
		countT[ch]++
	}

	// 比較兩個 map
	if len(countS) != len(countT) {
		return false
	}

	for ch, count := range countS {
		if countT[ch] != count {
			return false
		}
	}

	return true
}

// sortString 輔助函數：對字串排序
func sortString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

// isAnagramSortingSimplified 使用排序法（簡化版）
func isAnagramSortingSimplified(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return sortString(s) == sortString(t)
}

func main() {
	// 範例 1: "anagram" 和 "nagaram"
	s1 := "anagram"
	t1 := "nagaram"
	println("範例 1:")
	println("s =", s1, ", t =", t1)
	println("是否為字母異位詞:", isAnagram(s1, t1)) // true

	// 範例 2: "rat" 和 "car"
	s2 := "rat"
	t2 := "car"
	println("\n範例 2:")
	println("s =", s2, ", t =", t2)
	println("是否為字母異位詞:", isAnagram(s2, t2)) // false

	// 範例 3: 空字串
	s3 := ""
	t3 := ""
	println("\n範例 3:")
	println("s = \"\", t = \"\"")
	println("是否為字母異位詞:", isAnagram(s3, t3)) // true

	// 範例 4: 不同長度
	s4 := "a"
	t4 := "ab"
	println("\n範例 4:")
	println("s =", s4, ", t =", t4)
	println("是否為字母異位詞:", isAnagram(s4, t4)) // false

	// 字母計數演示
	println("\n\n字母計數演示（\"anagram\" vs \"nagaram\"）:")
	s := "anagram"
	t := "nagaram"
	var count [26]int

	println("處理字串 s:")
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		count[idx]++
		print("  ", string(s[i]), ": count[", idx, "] = ", count[idx], "\n")
	}

	println("\n處理字串 t:")
	for i := 0; i < len(t); i++ {
		idx := t[i] - 'a'
		count[idx]--
		print("  ", string(t[i]), ": count[", idx, "] = ", count[idx], "\n")
	}

	println("\n最終計數:")
	letters := "abcdefghijklmnopqrstuvwxyz"
	allZero := true
	for i, c := range count {
		if c != 0 {
			print("  ", string(letters[i]), ": ", c, "\n")
			allZero = false
		}
	}
	if allZero {
		println("  所有計數都為 0 → 是字母異位詞")
	}
}
