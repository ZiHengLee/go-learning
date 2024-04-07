package binarytree

import (
	"math"
)

//二叉树的最大距离
//递归的方式计算
func maxPathSum(root *TreeNode) int {
	var final = math.MinInt64
	maxPathSumDep(root,&final)
	return final
}

func maxPathSumDep(root *TreeNode,final *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if root.Val > *final{
			*final = root.Val
		}
		return root.Val
	}
	left := maxPathSumDep(root.Left,final)
	right := maxPathSumDep(root.Right,final)
	all := root.Val + myMax(0, left) + myMax(0, right)
	ret := root.Val + myMax(0, myMax(left, right))
	if myMax(root.Val, myMax(all, ret)) > *final{
		*final = myMax(root.Val, myMax(all, ret))
	}
	//注意这里只返回单边最大值，因为在连接的时候只能是单边
	return ret
}

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
