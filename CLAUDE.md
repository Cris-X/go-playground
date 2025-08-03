# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go playground project containing basic Go code examples and a dedicated LeetCode problem-solving environment. The project includes a structured approach to solving LeetCode Hot 100 problems with multiple solution methods and comprehensive testing.

## Common Commands

### Build and Run
```bash
# Run the main program
go run main.go

# Build the binary
go build

# Run the built binary
./go-playground
```

### Development
```bash
# Format code
go fmt

# Check for common mistakes
go vet

# Run tests (when test files are added)
go test

# Get dependencies
go mod tidy
```

### LeetCode Development
```bash
# Run all LeetCode tests
go test ./leetcode/...

# Run specific problem tests
go test ./leetcode/ -run TestTwoSum

# Run performance benchmarks
go test ./leetcode/ -bench=.

# Run tests with verbose output
go test ./leetcode/... -v

# Test a specific solution method
go test ./leetcode/ -run TestTwoSum/暴力解法
```

## Project Structure

- `main.go` - Single entry point with basic Go examples and GoLand IDE tips
- `go.mod` - Go module file (Go 1.22)
- `leetcode/` - LeetCode Hot 100 问题解决方案目录
  - `test_helper.go` - 通用测试框架和工具函数
  - `README.md` - LeetCode解题规范和模板说明
  - `{题号}_{题目名}.go` - 具体题目的多种解法实现
  - `{题号}_{题目名}_test.go` - 对应的测试文件

## LeetCode Development Guidelines

- 每个题目包含多种解法（暴力解法到最优解法）
- 详细的时间复杂度和空间复杂度分析
- 完整的测试用例覆盖，包括边界情况
- 支持性能基准测试对比不同解法
- 使用统一的文件命名和代码结构规范