/*
LeetCode 1. Two Sum
难度: Easy
标签: Array, Hash Table

题目描述:
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

示例:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
解释: 因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/

package leetcode

// TwoSumPractice 练习方法 - 在这里实现你的解法
// 时间复杂度: O(?)
// 空间复杂度: O(?)
func TwoSumPractice(nums []int, target int) []int {
	// TODO: 在这里实现你的解法
	//panic("请实现这个方法")
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if _, ok := m[complement]; ok {
			return []int{m[complement], i}
		}
		m[num] = i
	}

	return nil
}

// TwoSumBruteForce 解法1 - 暴力解法
// 时间复杂度: O(n²)
// 空间复杂度: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSumHashMap 解法2 - 哈希表一次遍历
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func TwoSumHashMap(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := hashMap[complement]; exists {
			return []int{j, i}
		}
		// 执行过程：
		// i=0: num=3, 先存入{3:0}, complement=3, 找到了3！
		// 返回 [0, 0] ❌ 错误！使用了同一个元素两次
		// 所以这里不允许在遍历前就存入
		hashMap[num] = i
	}
	return nil
}

// TwoSumHashMapTwoPass 解法3 - 哈希表两次遍历
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func TwoSumHashMapTwoPass(nums []int, target int) []int {
	hashMap := make(map[int]int)

	// 第一次遍历：建立哈希表
	for i, num := range nums {
		hashMap[num] = i
	}

	// 第二次遍历：查找complement
	for i, num := range nums {
		complement := target - num
		if j, exists := hashMap[complement]; exists && j != i {
			return []int{i, j}
		}
	}
	return nil
}
