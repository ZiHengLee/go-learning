package binarytree

type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右字树
}
//层次遍历
func Level(head *TreeNode) {

	if head == nil {
		return
	}

	// 用切片模仿队列
	var queue []*TreeNode
	queue = append(queue, head)

	for len(queue) != 0 {
		// 队头弹出，再把队头切掉，模仿队列的poll操作
		cur := queue[0]
		queue = queue[1:]
		print(cur.Data)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

}