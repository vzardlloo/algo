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

func preorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		if node != nil {
			stack = stack[:len(stack)-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}

	}

	return res
}
