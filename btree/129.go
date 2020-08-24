package btree

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	path := make([]int, 0)
	result := make([][]int, 0)
	sum := 0
	result = sumNumbersDFS(root, path, result)
	for _, list := range result {
		v := genIntValue(list)
		sum += v
	}
	return sum
}

func sumNumbersDFS(root *TreeNode, path []int, result [][]int) [][]int {
	if root.Left == nil && root.Right == nil {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, append(temp, root.Val))
		return result
	}
	path = append(path, root.Val)
	if root.Left != nil {
		result = sumNumbersDFS(root.Left, path, result)
	}
	if root.Right != nil {
		result = sumNumbersDFS(root.Right, path, result)
	}

	return result
}

func genIntValue(list []int) int {
	result := 0
	y := 1
	for i := len(list) - 1; i >= 0; i-- {
		value := list[i] * y
		result += value
		y = y * 10
	}

	return result
}
