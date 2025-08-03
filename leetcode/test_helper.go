package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

// TestCase 定义测试用例结构
type TestCase struct {
	Name     string
	Nums     []int
	Target   int
	Expected []int
}

// RunTwoSumTests 专门用于TwoSum类型题目的测试
func RunTwoSumTests(t *testing.T, testCases []TestCase, fn func([]int, int) []int) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := fn(tc.Nums, tc.Target)
			if !reflect.DeepEqual(actual, tc.Expected) {
				t.Errorf("测试失败\n输入: nums=%v, target=%d\n期望: %v\n实际: %v", 
					tc.Nums, tc.Target, tc.Expected, actual)
			} else {
				fmt.Printf("✓ %s: 通过\n", tc.Name)
			}
		})
	}
}

// BenchmarkTwoSumSolution 性能基准测试
func BenchmarkTwoSumSolution(b *testing.B, nums []int, target int, fn func([]int, int) []int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(nums, target)
	}
}

// ============ 通用测试框架 ============

// TestCaseGeneric 通用测试用例结构
type TestCaseGeneric struct {
	Name     string
	Input    interface{}
	Expected interface{}
}

// RunGenericTests 通用测试运行器，使用类型断言避免反射
func RunGenericTests[T any](t *testing.T, testCases []TestCaseGeneric, fn func(T) interface{}) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			input, ok := tc.Input.(T)
			if !ok {
				t.Fatalf("输入类型错误: 期望 %T, 实际 %T", *new(T), tc.Input)
			}
			
			actual := fn(input)
			if !reflect.DeepEqual(actual, tc.Expected) {
				t.Errorf("测试失败\n输入: %v\n期望: %v\n实际: %v", tc.Input, tc.Expected, actual)
			} else {
				fmt.Printf("✓ %s: 通过\n", tc.Name)
			}
		})
	}
}

// CreateProblemTemplate 创建题目模板
func CreateProblemTemplate(problemNum int, problemName string, difficulty string, tags []string) string {
	return fmt.Sprintf(`/*
LeetCode %d. %s
难度: %s
标签: %s

题目描述:
[在这里添加题目描述]

示例:
[在这里添加示例]
*/

package leetcode

// %sPractice 练习方法 - 在这里实现你的解法
// 时间复杂度: O(?)
// 空间复杂度: O(?)
func %sPractice(/* 参数列表 */) /* 返回类型 */ {
	// TODO: 在这里实现你的解法
	panic("请实现这个方法")
}

// %sSolution1 解法1 - 暴力解法
// 时间复杂度: O(?)
// 空间复杂度: O(?)
func %sSolution1(/* 参数列表 */) /* 返回类型 */ {
	// TODO: 实现暴力解法
	panic("待实现")
}

// %sSolution2 解法2 - 优化解法
// 时间复杂度: O(?)
// 空间复杂度: O(?)
func %sSolution2(/* 参数列表 */) /* 返回类型 */ {
	// TODO: 实现优化解法
	panic("待实现")
}
`, problemNum, problemName, difficulty, fmt.Sprintf("%v", tags), 
	problemName, problemName, problemName, problemName, problemName, problemName)
}

// CreateTestTemplate 创建测试模板
func CreateTestTemplate(problemName string) string {
	return fmt.Sprintf(`package leetcode

import (
	"testing"
)

func Test%s(t *testing.T) {
	// TODO: 添加测试用例
	testCases := []TestCaseGeneric{
		{
			Name:     "示例1",
			Input:    nil, // TODO: 设置输入
			Expected: nil, // TODO: 设置期望输出
		},
	}

	t.Run("练习方法", func(t *testing.T) {
		// TODO: 实现测试逻辑
		// RunGenericTests(t, testCases, %sPractice)
	})

	t.Run("解法1", func(t *testing.T) {
		// TODO: 实现测试逻辑
		// RunGenericTests(t, testCases, %sSolution1)
	})

	t.Run("解法2", func(t *testing.T) {
		// TODO: 实现测试逻辑
		// RunGenericTests(t, testCases, %sSolution2)
	})
}

func Benchmark%s(b *testing.B) {
	// TODO: 添加性能测试
	// input := // 设置测试输入
	
	// b.Run("练习方法", func(b *testing.B) {
	//     for i := 0; i < b.N; i++ {
	//         %sPractice(input)
	//     }
	// })
}
`, problemName, problemName, problemName, problemName, problemName, problemName)
}