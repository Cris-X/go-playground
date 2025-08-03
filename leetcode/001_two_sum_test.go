package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []TestCase{
		{
			Name:     "示例1",
			Nums:     []int{2, 7, 11, 15},
			Target:   9,
			Expected: []int{0, 1},
		},
		{
			Name:     "示例2",
			Nums:     []int{3, 2, 4},
			Target:   6,
			Expected: []int{1, 2},
		},
		{
			Name:     "示例3",
			Nums:     []int{3, 3},
			Target:   6,
			Expected: []int{0, 1},
		},
		{
			Name:     "负数测试",
			Nums:     []int{-1, -2, -3, -4, -5},
			Target:   -8,
			Expected: []int{2, 4},
		},
		{
			Name:     "零值测试",
			Nums:     []int{0, 4, 3, 0},
			Target:   0,
			Expected: []int{0, 3},
		},
	}

	t.Run("练习方法", func(t *testing.T) {
		// 启用测试，验证你的解法
		RunTwoSumTests(t, testCases, TwoSumPractice)
	})

	t.Run("暴力解法", func(t *testing.T) {
		RunTwoSumTests(t, testCases, TwoSumBruteForce)
	})

	t.Run("哈希表一次遍历", func(t *testing.T) {
		RunTwoSumTests(t, testCases, TwoSumHashMap)
	})

	t.Run("哈希表两次遍历", func(t *testing.T) {
		RunTwoSumTests(t, testCases, TwoSumHashMapTwoPass)
	})
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	
	b.Run("练习方法", func(b *testing.B) {
		// 启用性能测试，测试你的解法
		BenchmarkTwoSumSolution(b, nums, target, TwoSumPractice)
	})
	
	b.Run("暴力解法", func(b *testing.B) {
		BenchmarkTwoSumSolution(b, nums, target, TwoSumBruteForce)
	})
	
	b.Run("哈希表一次遍历", func(b *testing.B) {
		BenchmarkTwoSumSolution(b, nums, target, TwoSumHashMap)
	})
	
	b.Run("哈希表两次遍历", func(b *testing.B) {
		BenchmarkTwoSumSolution(b, nums, target, TwoSumHashMapTwoPass)
	})
}