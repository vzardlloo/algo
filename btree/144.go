package btree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return preorderTraversalDFS(root, []int{})
}

func preorderTraversalDFS(root *TreeNode, list []int) []int {
	list = append(list, root.Val)
	if root.Left != nil {
		list = preorderTraversalDFS(root.Left, list)
	}
	if root.Right != nil {
		list = preorderTraversalDFS(root.Right, list)
	}

	return list
}
