//设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。 
//
// 示例： 
//
// 输入： arr = [1,3,5,7,2,4,6,8], k = 4
//输出： [1,2,3,4]
// 
//
// 提示： 
//
// 
// 0 <= len(arr) <= 100000 
// 0 <= k <= min(100000, len(arr)) 
// 
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 246 👎 0

package cn

import (
	"container/heap"
	"sort"
)


//leetcode submit region begin(Prohibit modification and deletion)
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }
//leetcode submit region end(Prohibit modification and deletion)

// 主要方法：快排 -> 快速选择
func quickSelect(arr []int, left, right, k int) {
	if left >= right {
		return
	}

	// 分区
	pivotIdx := partition(arr, left, right)

	if pivotIdx == k {
		return // 找到了第k个位置
	} else if pivotIdx < k {
		// 第k个元素在右边
		quickSelect(arr, pivotIdx+1, right, k)
	} else {
		// 第k个元素在左边
		quickSelect(arr, left, pivotIdx-1, k)
	}
}

func partition(arr []int, left, right int) int {
	// 选择最右边的元素作为基准
	pivot := arr[right]
	i := left // i指向小于pivot的区域的右边界

	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// 将基准放到正确位置
	arr[i], arr[right] = arr[right], arr[i]
	return i
}


