package binarytree

//又是层次便利？每次记录最后一个节点
func rightSideView(root *TreeNode) []int {
	ret := make([]int,0)
	myQueue := make([]*TreeNode,0)
	if root == nil{
		return ret
	}
	myQueue = append(myQueue, root)
	for len(myQueue) > 0{
		var tempQueue []*TreeNode
		ret = append(ret, myQueue[len(myQueue)-1].Val)
		for len(myQueue) > 0{
			cur := myQueue[0]
			myQueue=myQueue[1:]
			if cur.Left !=nil{
				tempQueue=append(tempQueue,cur.Left)
			}
			if cur.Right !=nil{
				tempQueue=append(tempQueue,cur.Right)
			}
		}
		for _,a := range tempQueue{
			myQueue = append(myQueue,a)
		}
	}
	return ret
}
