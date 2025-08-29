//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。 
//
// 子数组是数组中元素的连续非空序列。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,1], k = 2
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3], k = 3
//输出：2
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2 * 10⁴ 
// -1000 <= nums[i] <= 1000 
// -10⁷ <= k <= 10⁷ 
// 
//
// Related Topics 数组 哈希表 前缀和 👍 2856 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	// 步骤1: 构建前缀和数组
	// s[i] 表示 nums[0] + nums[1] + ... + nums[i-1] 的和
	// s[0] = 0 (空数组的和)
	s := make([]int, len(nums) + 1)
	for i, v := range nums {
		s[i + 1] = s[i] + v
	}

	// 步骤2: 创建哈希表记录前缀和出现次数
	ans := 0
	cnt := make(map[int]int, len(s))

	// 步骤3: 遍历前缀和数组
	for _, sj := range s {
		// 关键思路: 如果存在前缀和 si，使得 sj - si = k
		// 那么从位置 i+1 到位置 j 的子数组和就是 k
		// 即: sj - si = k  =>  si = sj - k
		ans += cnt[sj - k]

		// 将当前前缀和加入哈希表
		cnt[sj]++
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
