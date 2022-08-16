package binarytree

//主要在于利用完全二叉数的性质

//左右高度相同，左边一定是满的
//左右不同，右边一定是满的
//这里比较难理解的是左移运算符，主要是因为需要包括head节点
func countNodes(root *TreeNode) int {
	if root == nil{
		return 0
	}
	left, right := getDepth(root.Left),getDepth(root.Right)
	if left == right{
		return (1 << left) + countNodes(root.Right)
	}
	return (1 << right) + countNodes(root.Left)
}

func getDepth(root *TreeNode) int{
	depth := 0
	for root != nil{
		root=root.Left
		depth+=1
	}
	return depth
}
