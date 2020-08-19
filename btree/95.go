/*
 * Author:   Cailei Lu
 * Email:    lucailei@foxmail.com
 * Time:     2020/8/19
 */
package btree

// [95] 不同的二叉搜索树 ‖
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesHelper(1, n)
}

func generateTreesHelper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := make([]*TreeNode, 0)

	for i := start; i <= end; i++ {
		leftTrees := generateTreesHelper(start, i-1)
		rightTrees := generateTreesHelper(i+1, end)

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
