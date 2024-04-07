package binarytree

//判断一棵树是不是完全二叉树

//层次便利变种，把节点为nil也加入队列，当访问到第一个nil节点，其后如果能再访问到一个不是nil节点，就退出
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var myQueue []*TreeNode
	myQueue = append(myQueue, root)
	for len(myQueue) > 0 {
		cur := myQueue[0]
		myQueue = myQueue[1:]
		if cur == nil {
			break
		}
		myQueue = append(myQueue, cur.Left)
		myQueue = append(myQueue, cur.Right)
	}
	res := true
	for i := 0; i < len(myQueue); i++ {
		if myQueue[i] != nil {
			res = false
			break
		}
	}
	return res
}
