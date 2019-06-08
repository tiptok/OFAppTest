package leetcode

//二叉树 中序遍历
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//BTreeToArray TreeNode->[]int
func (*TreeNode)BTreeToArray(tree *TreeNode)[]int{
	return inorderTraversal(tree)
}
//SortArrayToBTree  []int->TreeNode
func (*TreeNode)SortedArrayToBTree(array []int)*TreeNode{
	return sortedArrayToBST(array)
}
