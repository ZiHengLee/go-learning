package binarytree

//root = [1,2,3]
//25
//根节点到叶子节点的数的和

//思路还是递归的思想，其实看起来简单，有点小恶心
func sumNumbers(root *TreeNode) int {
	ret := 0
	sumHelp(root,0,&ret)
	return ret
}

func sumHelp(node *TreeNode,sum int,ret *int) {
	sum = sum*10
	sum = sum + node.Val
	if node.Left == nil && node.Right == nil{
		*ret = sum+ *ret
		return
	}
	if node.Left != nil{
		sumHelp(node.Left,sum,ret)
	}
	if node.Right != nil{
		sumHelp(node.Right,sum,ret)
	}
}
