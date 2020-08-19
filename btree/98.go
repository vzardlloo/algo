/*
 * Author:   Cailei Lu
 * Email:    lucailei@foxmail.com
 * Time:     2020/8/19
 */
package btree

//[98] 验证二叉搜索树
func isValidBST(root *TreeNode) bool {

	return isValidBSTCheck(root, nil, nil)
}

func isValidBSTCheck(node, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= min.Val {
		return false
	}
	if max != nil && node.Val >= max.Val {
		return false
	}

	return isValidBSTCheck(node.Left, min, node) && isValidBSTCheck(node.Right, node, max)
}
