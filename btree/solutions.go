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
