package binarytree

import "fmt"

//计算二叉树直径和,求边数，注意
func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	dfs(root,&ret)
	return ret
}

func dfs(root *TreeNode,sum *int) int{
	if root == nil{
		return 0
	}
	leftSize,rightSize := 0,0
	if root.Left != nil{
		leftSize = dfs(root.Left,sum)+1
	}
	if root.Right != nil{
		rightSize = dfs(root.Right,sum)+1
	}
	//因为是边数，所以这个地方是不用-1的
	if *sum < leftSize+rightSize{
		*sum = leftSize+rightSize
	}
	fmt.Println(leftSize,rightSize)
	if leftSize > rightSize{
		return leftSize
	}
	return rightSize
}