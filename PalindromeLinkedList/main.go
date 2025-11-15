package main

// ListNode 定義鏈結串列節點
type ListNode struct {
	Val  int
	Next *ListNode
}

// isPalindrome 使用快慢指標和反轉鏈結串列判斷是否為回文（最優解法）
// 步驟：
// 1. 使用快慢指標找到鏈結串列中點
// 2. 反轉後半部分鏈結串列
// 3. 比較前半部分和反轉後的後半部分
// 4. （可選）恢復鏈結串列
// 時間複雜度: O(n)
// 空間複雜度: O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 步驟 1：使用快慢指標找到中點
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 步驟 2：反轉後半部分
	// slow 現在指向後半部分的起始（或中點的下一個）
	secondHalf := reverseList(slow)

	// 步驟 3：比較前半部分和反轉後的後半部分
	firstHalf := head
	result := true
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			result = false
			break
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	// 步驟 4：（可選）恢復鏈結串列
	// reverseList(secondHalfStart)

	return result
}

// reverseList 反轉鏈結串列
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// isPalindromeArray 使用陣列法（空間換時間）
// 將鏈結串列值複製到陣列，然後使用雙指標判斷
// 時間複雜度: O(n)
// 空間複雜度: O(n)
func isPalindromeArray(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 複製到陣列
	values := []int{}
	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}

	// 雙指標判斷回文
	left, right := 0, len(values)-1
	for left < right {
		if values[left] != values[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// isPalindromeRecursive 使用遞迴法
// 時間複雜度: O(n)
// 空間複雜度: O(n)，遞迴堆疊
func isPalindromeRecursive(head *ListNode) bool {
	frontPointer := head
	var recursiveCheck func(*ListNode) bool

	recursiveCheck = func(currentNode *ListNode) bool {
		if currentNode != nil {
			// 遞迴到鏈結串列末尾
			if !recursiveCheck(currentNode.Next) {
				return false
			}
			// 回溯時比較
			if currentNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}

	return recursiveCheck(head)
}

// 輔助函數：從陣列創建鏈結串列
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	curr := head

	for i := 1; i < len(values); i++ {
		curr.Next = &ListNode{Val: values[i]}
		curr = curr.Next
	}

	return head
}

// 輔助函數：將鏈結串列轉換為陣列（用於輸出）
func listToArray(head *ListNode) []int {
	result := []int{}
	curr := head

	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}

	return result
}

func main() {
	// 範例 1: [1,2,2,1]
	list1 := createList([]int{1, 2, 2, 1})
	print("範例 1: [1,2,2,1]\n")
	print("是否為回文: ", isPalindrome(list1), "\n") // true

	// 範例 2: [1,2]
	list2 := createList([]int{1, 2})
	print("\n範例 2: [1,2]\n")
	print("是否為回文: ", isPalindrome(list2), "\n") // false

	// 範例 3: [1]
	list3 := createList([]int{1})
	print("\n範例 3: [1]\n")
	print("是否為回文: ", isPalindrome(list3), "\n") // true

	// 範例 4: [1,2,3,2,1]（奇數長度）
	list4 := createList([]int{1, 2, 3, 2, 1})
	print("\n範例 4: [1,2,3,2,1]\n")
	print("是否為回文: ", isPalindrome(list4), "\n") // true

	// 快慢指標演示
	println("\n\n快慢指標找中點演示（[1,2,3,4,5]）:")
	list := createList([]int{1, 2, 3, 4, 5})
	slow, fast := list, list
	step := 0

	print("初始: slow=1, fast=1\n")
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		step++
		print("步驟 ", step, ": slow=", slow.Val)
		if fast != nil {
			print(", fast=", fast.Val)
		} else {
			print(", fast=nil")
		}
		print("\n")
	}
	print("中點: ", slow.Val, "\n")
}
