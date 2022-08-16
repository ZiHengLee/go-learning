package binarytree

//计算二叉树的最大宽度，基本思路还是层次便利，只不过需要记录每个节点的val值
//left = root*2
//right = root*2+1

func widthOfBinaryTree(root *TreeNode) int {
	var ret int
	root.Val = 0
	myQueue := make([]*TreeNode,0)
	myQueue = append(myQueue, root)
	for len(myQueue) > 0{
		levelLen := myQueue[len(myQueue)-1].Val-myQueue[0].Val+1
		if levelLen > ret{
			ret = levelLen
		}
		var tempQueue []*TreeNode
		for len(myQueue) > 0{
			cur := myQueue[0]
			myQueue=myQueue[1:]
			if cur.Left !=nil{
				cur.Left.Val = cur.Val*2
				tempQueue=append(tempQueue,cur.Left)
			}
			if cur.Right !=nil{
				cur.Right.Val = cur.Val*2+1
				tempQueue=append(tempQueue,cur.Right)
			}
		}
		for _,a := range tempQueue{
			myQueue = append(myQueue,a)
		}
	}
	return ret
}
