package binarytree

//层次便利变通，一下从左出，一下从右出
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int,0)
	if root == nil {
		return ret
	}
	var queue []*TreeNode
	queue = append(queue, root)

	dir := true
	for len(queue) != 0 {
		level := make([]int,0)
		tempQueue := make([]*TreeNode,len(queue))
		copy(tempQueue,queue)
		queue = make([]*TreeNode,0)
		lenTemp := len(tempQueue)
		for i:=0;i<lenTemp;i++{
			curNode := tempQueue[0]
			level = append(level, curNode.Val)
			tempQueue = tempQueue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		if !dir{
			temp1Queue := make([]int,0)
			for i:=0;i<len(level);i++{
				temp1Queue = append(temp1Queue, level[len(level)-i-1])
			}
			ret = append(ret, temp1Queue)
		}else{
			ret = append(ret, level)
		}
		dir=!dir
	}
	return ret
}