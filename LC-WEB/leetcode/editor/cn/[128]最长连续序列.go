//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。 
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。 
//
// 示例 2： 
//
// 
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
// 
//
// 示例 3： 
//
// 
//输入：nums = [1,0,1,2]
//输出：3
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 10⁵ 
// -10⁹ <= nums[i] <= 10⁹ 
// 
//
// Related Topics 并查集 数组 哈希表 👍 2583 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	// 图解在vsc中查看
	has := map[int]bool{}
	for _, num := range nums {
		has[num] = true
	}

	res := 0
	for x := range has {
		// 该数字非序列起点
		if has[x - 1] {
			continue
		}
		// x 是序列起点
		y := x + 1
		for has[y] { // 不断查找下一个数是否在集合中
			y ++
		}
		res = max(res, y - x) // 最终序列为x...y - 1, 因为上一步y是不存在hash中的
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
