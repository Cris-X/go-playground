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

import "container/heap"


//leetcode submit region begin(Prohibit modification and deletion)
type MaxHeap []int
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 大顶堆
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func smallestK(arr []int, k int) (res []int) {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}

	h := &MaxHeap{}
	heap.Init(h)

	for _, num := range arr {
		if h.Len() < k {
			// 堆未满，直接加入
			heap.Push(h, num)
		} else if num < (*h)[0] {
			// 当前元素比堆顶小，替换堆顶
			heap.Pop(h)
			heap.Push(h, num)
		}
	}

	return *h
}
//leetcode submit region end(Prohibit modification and deletion)

// 主要方法：快排 -> 快速选择
func quickSelect(arr []int, left, right, k int) {
	if left >= right {
		return
	}

	// 分区
	pivot := partition(arr, left, right)

	if pivot == k {
		return // 找到了第k个位置
	} else if pivot < k {
		// 第k个元素在右边
		quickSelect(arr, pivot+1, right, k)
	} else {
		// 第k个元素在左边
		quickSelect(arr, left, pivot-1, k)
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


