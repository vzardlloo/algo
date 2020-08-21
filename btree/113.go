/*
 * Author:   Cailei Lu
 * Email:    lucailei@foxmail.com
 * Time:     2020/8/21
 */
package btree

// [113] 路径总和 ‖
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	path := make([]int, 0)
	return pathSumHelper(root, sum, path, result)
}

func pathSumHelper(root *TreeNode, sum int, path []int, result [][]int) [][]int {
	if root == nil {
		return result
	}

	//path = append(path,root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		temp := make([]int, len(path)+1)
		copy(temp, append(path, root.Val))
		result = append(result, temp)
		return result

	}
	path = append(path, root.Val)
	if root.Left != nil {
		result = pathSumHelper(root.Left, sum-root.Val, path, result)
	}
	if root.Right != nil {
		result = pathSumHelper(root.Right, sum-root.Val, path, result)
	}

	return result
}
