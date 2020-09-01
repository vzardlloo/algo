package btree

// 定义f(x)为以x节点为根结点时,路径数量
// 定义g(x)为以x为节点的树里面所有拥有的路径总数
// g(x) = f(x) + g(x.Left) + g(x.Right)
func pathSumV3(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := countPath(root, sum)

	return result + pathSumV3(root.Left, sum) + pathSumV3(root.Right, sum)
}

func countPath(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	var result int
	if sum-root.Val == 0 {
		result = 1
	} else {
		result = 0
	}
	sum = sum - root.Val

	return result + countPath(root.Left, sum) + countPath(root.Right, sum)
}
