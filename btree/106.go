/*
 * Author:   Cailei Lu
 * Email:    lucailei@foxmail.com
 * Time:     2020/8/21
 */
package btree

// [106] 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 || len(postorder) <= 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1], Left: nil, Right: nil}
	index := 0
	for k, v := range inorder {
		if v == postorder[len(postorder)-1] {
			index = k
		}
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}
