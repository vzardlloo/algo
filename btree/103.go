/*
 * Author:   Cailei Lu
 * Email:    lucailei@foxmail.com
 * Time:     2020/8/19
 */
package btree

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
