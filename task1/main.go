package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("=== 只出现一次的数字 ===")
	testCases := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
	}
	for i, nums := range testCases {
		result := singleNumber(nums)
		fmt.Printf("输入 %d: %v -> 输出: %d\n", i+1, nums, result)
	}

	fmt.Println("\n=== 回文数 ===")
	palindromeTests := []int{121, -121, 10}
	for i, x := range palindromeTests {
		result := isPalindrome(x)
		fmt.Printf("输入 %d: %d -> 输出: %t\n", i+1, x, result)
	}

	fmt.Println("\n=== 有效的括号 ===")
	parenthesesTests := []string{"()", "()[]{}", "(]"}
	for i, s := range parenthesesTests {
		result := isValid(s)
		fmt.Printf("输入 %d: \"%s\" -> 输出: %t\n", i+1, s, result)
	}

	fmt.Println("\n=== 最长公共前缀 ===")
	prefixTests := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"", "b"},
	}
	for i, strs := range prefixTests {
		result := longestCommonPrefix(strs)
		fmt.Printf("输入 %d: %v -> 输出: \"%s\"\n", i+1, strs, result)
	}

	fmt.Println("\n=== 加一 ===")
	plusOneTests := [][]int{
		{1, 2, 3},
		{9},
		{9, 9, 9},
	}
	for i, digits := range plusOneTests {
		result := plusOne(digits)
		fmt.Printf("输入 %d: %v -> 输出: %v\n", i+1, digits, result)
	}

	fmt.Println("\n=== 删除重复项 ===")
	removeDuplicatesTests := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{1, 2, 3, 4, 5},
	}
	for i, nums := range removeDuplicatesTests {
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		k := removeDuplicates(testNums)
		fmt.Printf("输入 %d: %v -> 输出: %v\n", i+1, nums, testNums[:k])
	}

	fmt.Println("\n=== 合并区间 ===")
	mergeIntervalsTests := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {5, 6}},
	}
	for i, intervals := range mergeIntervalsTests {
		result := mergeIntervals(intervals)
		fmt.Printf("输入 %d: %v -> 输出: %v\n", i+1, intervals, result)
	}

	fmt.Println("\n=== 两数之和 ===")
	twoSumTests := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}
	for i, test := range twoSumTests {
		result := twoSum(test.nums, test.target)
		fmt.Printf("输入 %d: nums=%v, target=%d -> 输出: %v\n", i+1, test.nums, test.target, result)
	}
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x != 0 && x%10 == 0 {
		return false
	}

	reversed := 0

	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}

// 有效的括号
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]rune, 0)

	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if left, isRight := pairs[char]; isRight {
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

// 加一
func plusOne(digits []int) []int {
	result := make([]int, len(digits))
	copy(result, digits)

	for i := len(result) - 1; i >= 0; i-- {
		if result[i] < 9 {
			result[i]++
			return result
		}
		result[i] = 0
	}

	return append([]int{1}, result...)
}

// 删除重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// 合并区间
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			result = append(result, current)
		}
	}

	return result
}

// 两数之和
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, exists := numMap[complement]; exists {
			return []int{index, i}
		}

		numMap[num] = i
	}

	return []int{}
}
