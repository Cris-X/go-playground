package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// createLeetCodeProblem 创建新的LeetCode题目文件
func createLeetCodeProblem(problemNum int, problemName, difficulty string, tags []string) error {
	// 创建文件名
	fileName := fmt.Sprintf("%03d_%s", problemNum, strings.ToLower(strings.ReplaceAll(problemName, " ", "_")))
	
	// 创建题目文件
	problemFile := filepath.Join("leetcode", fileName+".go")
	testFile := filepath.Join("leetcode", fileName+"_test.go")
	
	// 生成题目模板
	problemContent := createProblemTemplate(problemNum, problemName, difficulty, tags)
	testContent := createTestTemplate(problemName)
	
	// 写入文件
	if err := os.WriteFile(problemFile, []byte(problemContent), 0644); err != nil {
		return fmt.Errorf("创建题目文件失败: %w", err)
	}
	
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		return fmt.Errorf("创建测试文件失败: %w", err)
	}
	
	fmt.Printf("✅ 成功创建题目文件:\n")
	fmt.Printf("   - %s\n", problemFile)
	fmt.Printf("   - %s\n", testFile)
	
	return nil
}

func createProblemTemplate(problemNum int, problemName, difficulty string, tags []string) string {
	// 处理函数名，确保符合Go命名规范
	funcName := strings.ReplaceAll(problemName, " ", "")
	funcName = strings.ReplaceAll(funcName, "-", "")
	// 如果函数名以数字开头，添加前缀
	if len(funcName) > 0 && funcName[0] >= '0' && funcName[0] <= '9' {
		funcName = "Problem" + funcName
	}
	
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
`, problemNum, problemName, difficulty, strings.Join(tags, ", "), 
	funcName, funcName, funcName, funcName, funcName, funcName)
}

func createTestTemplate(problemName string) string {
	// 处理函数名，确保符合Go命名规范
	funcName := strings.ReplaceAll(problemName, " ", "")
	funcName = strings.ReplaceAll(funcName, "-", "")
	// 如果函数名以数字开头，添加前缀
	if len(funcName) > 0 && funcName[0] >= '0' && funcName[0] <= '9' {
		funcName = "Problem" + funcName
	}
	
	return fmt.Sprintf(`package leetcode

import (
	"testing"
)

func Test%s(t *testing.T) {
	// TODO: 添加测试用例
	// testCases := []TestCaseGeneric{
	// 	{
	// 		Name:     "示例1",
	// 		Input:    nil, // TODO: 设置输入
	// 		Expected: nil, // TODO: 设置期望输出
	// 	},
	// }

	t.Run("练习方法", func(t *testing.T) {
		// TODO: 实现测试逻辑
		// RunGenericTests(t, testCases, %sPractice)
		t.Skip("请先实现测试用例")
	})

	t.Run("解法1", func(t *testing.T) {
		// TODO: 实现测试逻辑
		// RunGenericTests(t, testCases, %sSolution1)
		t.Skip("请先实现测试用例")
	})

	t.Run("解法2", func(t *testing.T) {
		// TODO: 实现测试逻辑
		// RunGenericTests(t, testCases, %sSolution2)
		t.Skip("请先实现测试用例")
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
	b.Skip("请先实现性能测试")
}
`, funcName, funcName, funcName, funcName, funcName, funcName)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("用法: go run create_problem.go <题号> <题目名> <难度> [标签1,标签2,...]")
		fmt.Println("示例: go run create_problem.go 15 \"3Sum\" Medium Array,TwoPointers")
		os.Exit(1)
	}
	
	problemNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("错误: 题号必须是数字: %v\n", err)
		os.Exit(1)
	}
	
	problemName := os.Args[2]
	difficulty := os.Args[3]
	
	var tags []string
	if len(os.Args) > 4 {
		tags = strings.Split(os.Args[4], ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
	} else {
		tags = []string{"TODO"}
	}
	
	if err := createLeetCodeProblem(problemNum, problemName, difficulty, tags); err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}
}