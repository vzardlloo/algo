/*
 * Author:   Cailei Lu
 * Email:    lucailei@foxmail.com
 * Time:     2020/8/16
 */
package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
