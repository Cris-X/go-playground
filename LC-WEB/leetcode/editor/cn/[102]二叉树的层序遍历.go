//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
// 
//
// 示例 2： 
//
// 
//输入：root = [1]
//输出：[[1]]
// 
//
// 示例 3： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 2000] 内 
// -1000 <= Node.val <= 1000 
// 
//
// Related Topics 树 广度优先搜索 二叉树 👍 2176 👎 0

package cn

// levelOrder 实现二叉树的层序遍历，返回每一层节点值的二维数组
func levelOrder01(root *TreeNode) (ans [][]int) {
	// 边界条件：如果根节点为空，直接返回空结果
	if root == nil {
		return
	}

	// cur 存储当前层的所有节点，初始时只有根节点
	cur := []*TreeNode{root}

	// 当当前层还有节点时，继续遍历
	for len(cur) > 0 {
		// nxt 用于存储下一层的所有节点
		nxt := []*TreeNode{}

		// vals 存储当前层所有节点的值，长度等于当前层节点数量
		vals := make([]int, len(cur))

		// 遍历当前层的每一个节点
		for i, node := range cur {
			// 将当前节点的值存储到 vals 数组中
			vals[i] = node.Val

			// 如果左子节点不为空，加入下一层节点列表
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}

			// 如果右子节点不为空，加入下一层节点列表
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}

		// 更新 cur 为下一层节点，准备处理下一层
		cur = nxt

		// 将当前层的所有节点值添加到最终结果中
		ans = append(ans, vals)
	}

	// 返回层序遍历结果
	return
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// levelOrder 实现二叉树的层序遍历，返回每一层节点值的二维数组
func levelOrder(root *TreeNode) (ans [][]int) {
	// 边界条件：如果根节点为空，直接返回空结果
	if root == nil {
		return
	}

	// q 作为队列存储待处理的节点，初始时只放入根节点
	q := []*TreeNode{root}

	// 当队列不为空时，继续遍历
	for len(q) > 0 {
		// n 记录当前层的节点数量（队列当前长度）
		n := len(q)

		// vals 存储当前层所有节点的值，预分配长度为 n
		vals := make([]int, n)

		// 处理当前层的 n 个节点
		for i := range vals {
			// 从队列头部取出一个节点（出队操作）
			node := q[0]
			q = q[1:]

			// 将节点值存入当前层结果数组
			vals[i] = node.Val

			// 如果左子节点存在，将其加入队列尾部（入队操作）
			if node.Left != nil {
				q = append(q, node.Left)
			}

			// 如果右子节点存在，将其加入队列尾部（入队操作）
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		// 将当前层的所有节点值添加到最终结果中
		ans = append(ans, vals)
	}

	// 返回层序遍历结果
	return
}
//leetcode submit region end(Prohibit modification and deletion)
