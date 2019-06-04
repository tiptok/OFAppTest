package leetcode

import (
	"log"
	"testing"
)
//二叉树 中序遍历
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Test_inorderTraversal(t *testing.T){
	input :=[][]int{
		{-10,3,0,5,9},
		{1,2,3,4,5},
	}
	for i:=range input{
		inBST :=SortedArrayToBST(input[i])
		out :=inorderTraversal(inBST)
		log.Println("Input:",input[i]," output:",out)
	}
}
//使用堆栈的方式
func inorderTraversal(root *TreeNode) []int {
	stack :=make([]*TreeNode,0)
	ret :=make([]int,0)
	var (
		cur *TreeNode = root
	)
	pop :=func(stack []*TreeNode)([]*TreeNode,*TreeNode){
		if len(stack)==0 || stack==nil{
			return stack,nil
		}
		if len(stack)==1{
			return nil,stack[0]
		}
		p :=stack[len(stack)-1]
		stack =stack[:len(stack)-1]
		return stack,p
	}
	push:=func(stack []*TreeNode,cur *TreeNode)[]*TreeNode{
		if stack==nil{
			stack=make([]*TreeNode,0)
		}
		stack=append(stack,cur)
		return stack
	}
	for ;cur!=nil || (len(stack)!=0 && stack!=nil);{
		if cur!=nil{
			stack=push(stack,cur)
			cur = cur.Left
		}else{
			stack,cur =pop(stack)
			ret=append(ret,cur.Val)
			cur = cur.Right
		}
	}
	return ret
}