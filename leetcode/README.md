# LeetCode Hot 100 题解模板

## 快速开始

### 创建新题目
使用项目根目录的题目生成器快速创建新题目：

```bash
# 基本用法
go run create_problem.go <题号> <题目名> <难度> [标签1,标签2,...]

# 示例：创建3Sum题目
go run create_problem.go 15 "3Sum" Medium Array,TwoPointers

# 示例：创建Valid Parentheses题目  
go run create_problem.go 20 "Valid Parentheses" Easy String,Stack
```

这将自动生成：
- `leetcode/015_3sum.go` - 包含练习方法和参考解法
- `leetcode/015_3sum_test.go` - 对应的测试文件

### 解题流程

1. **选择练习方法**: 每个题目都有一个 `{题目名}Practice` 方法供你练习
2. **实现解法**: 在Practice方法中实现你的解法
3. **运行测试**: 取消注释测试代码，验证你的解法
4. **对比优化**: 查看参考解法，学习优化思路
5. **性能测试**: 运行基准测试对比不同解法的性能

## 文件结构说明

### 解题文件结构
```go
// {题目名}Practice - 你的练习方法
func TwoSumPractice(nums []int, target int) []int {
    // TODO: 在这里实现你的解法
    panic("请实现这个方法")
}

// {题目名}Solution1 - 参考解法1（通常是暴力解法）
func TwoSumSolution1(nums []int, target int) []int {
    // 暴力解法实现
}

// {题目名}Solution2 - 参考解法2（优化解法）
func TwoSumSolution2(nums []int, target int) []int {
    // 优化解法实现  
}
```

### 测试文件结构
```go
func TestTwoSum(t *testing.T) {
    // 练习方法测试（初始为注释状态）
    t.Run("练习方法", func(t *testing.T) {
        // RunTwoSumTests(t, testCases, TwoSumPractice)
        t.Skip("练习方法待实现...")
    })
    
    // 参考解法测试
    t.Run("解法1", func(t *testing.T) {
        RunTwoSumTests(t, testCases, TwoSumSolution1)
    })
}
```

## 测试命令

### 基本测试
```bash
# 测试所有LeetCode题目
go test ./leetcode/...

# 测试特定题目
go test ./leetcode/ -run TestTwoSum

# 详细测试输出
go test ./leetcode/... -v

# 测试特定解法
go test ./leetcode/ -run TestTwoSum/暴力解法
```

### 性能测试
```bash
# 运行所有性能测试
go test ./leetcode/ -bench=.

# 运行特定题目的性能测试
go test ./leetcode/ -bench=BenchmarkTwoSum

# 运行特定解法的性能测试  
go test ./leetcode/ -bench=BenchmarkTwoSum/哈希表
```

## 使用示例

### 步骤1: 创建新题目
```bash
go run create_problem.go 15 "3Sum" Medium Array,TwoPointers
```

### 步骤2: 实现你的解法
编辑 `leetcode/015_3sum.go` 中的 `ThreeSumPractice` 方法

### 步骤3: 启用测试
编辑 `leetcode/015_3sum_test.go`，取消注释练习方法的测试代码

### 步骤4: 运行测试
```bash
go test ./leetcode/ -run TestThreeSum/练习方法 -v
```

### 步骤5: 对比参考解法
查看同文件中的 `ThreeSumSolution1` 和 `ThreeSumSolution2` 方法

## 文件命名规范
- 题目文件: `{题号}_{题目英文名}.go`
- 测试文件: `{题号}_{题目英文名}_test.go`

例如: `001_two_sum.go` 和 `001_two_sum_test.go`

## 特性

✅ **练习导向**: 每题都有专门的练习方法供你实现  
✅ **多解法对比**: 从暴力解法到最优解法的完整实现  
✅ **完整测试**: 包含边界用例和性能基准测试  
✅ **一键生成**: 自动生成题目模板和测试框架  
✅ **详细分析**: 每种解法都有时间/空间复杂度分析  

## 常用数据结构

需要时可以在 `test_helper.go` 中添加链表、二叉树等通用数据结构定义。