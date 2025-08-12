//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != 
//k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
// 
//
// 示例 3： 
//
// 
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
// 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 3000 
// -10⁵ <= nums[i] <= 10⁵ 
// 
//
// Related Topics 数组 双指针 排序 👍 7593 👎 0

package cn

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) (res [][]int) {
	n := len(nums)
	if n < 3 {
		return nil
	}

	sort.Ints(nums) // 排序
	for i := 0; i < n - 2; i ++ { // 固定第一个数
		if nums[i] > 0 { // 优化：第一个数大于0，不可能和为0，因为该数组排序过
			break
		}

		// 去重第一个数
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		left, right := i + 1, n - 1 // 设置双指针
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]
			
			if sum == target { // 找到一组解
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// 内部双指针去重记得用for，因为不像第一个数那样continue
				// 左指针去重
				for left < right && nums[left] == nums[left + 1] {
					left ++
				}
				// 有指针去重
				for left < right && nums[right] == nums[right - 1] {
					right --
				}

				left ++
				right --
			} else if sum < target { // 比目标小，增大
				left ++
			} else { // 比目标大，减小
				right --
			}
		}
	}

	return
}
//leetcode submit region end(Prohibit modification and deletion)
