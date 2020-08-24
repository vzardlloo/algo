package btree

type BSTIterator struct {
	List     []int
	CurIndex int
}

func Constructor(root *TreeNode) BSTIterator {
	itor := &BSTIterator{
		List:     make([]int, 0),
		CurIndex: -1,
	}
	if root == nil {
		return *itor
	}
	ConstructorHelper(root, itor)
	return *itor
}

func ConstructorHelper(root *TreeNode, iter *BSTIterator) *BSTIterator {
	if root.Left != nil {
		iter = ConstructorHelper(root.Left, iter)
	}
	iter.List = append(iter.List, root.Val)
	if root.Right != nil {
		iter = ConstructorHelper(root.Right, iter)
	}

	return iter
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if len(this.List) <= 0 {
		return 0
	}
	this.CurIndex++
	value := this.List[this.CurIndex]
	return value
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.CurIndex >= len(this.List)-1 {
		return false
	}
	return true
}
