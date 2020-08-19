/*
 * Author:   Cailei Lu
 * Email:    lucailei@foxmail.com
 * Time:     2020/8/16
 */
package btree

// [96] 不同的二叉搜索树
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

// [95] 不同的二叉搜索树 ‖
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := make([]*TreeNode, 0)

	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currentTree := &TreeNode{
					Val:   i,
					Left:  nil,
					Right: nil,
				}
				currentTree.Left = left
				currentTree.Right = right

				allTrees = append(allTrees, currentTree)
			}
		}
	}
	return allTrees
}

//[98] 验证二叉搜索树
func isValidBST(root *TreeNode) bool {

	return check(root, nil, nil)
}

func check(node, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= min.Val {
		return false
	}
	if max != nil && node.Val >= max.Val {
		return false
	}

	return check(node.Left, min, node) && check(node.Right, node, max)
}

//[103] 二叉树的锯齿形层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	result = zigzagLevelOrderDFS(root, 0, result)
	return result
}

func zigzagLevelOrderDFS(node *TreeNode, level int, result [][]int) [][]int {
	// if a new level
	if level >= len(result) {
		list := make([]int, 0)
		list = append(list, node.Val)
		result = append(result, list)
	} else {
		if level%2 == 0 {
			result[level] = append(result[level], node.Val)
		} else {
			temp := make([]int, len(result[level])+1)
			for i := len(result[level]); i > 0; i-- {
				temp[i] = result[level][i-1]
			}
			temp[0] = node.Val
			result[level] = temp
		}
	}

	if node.Left != nil {
		result = zigzagLevelOrderDFS(node.Left, level+1, result)
	}
	if node.Right != nil {
		result = zigzagLevelOrderDFS(node.Right, level+1, result)
	}
	return result
}
