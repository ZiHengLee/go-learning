package binarytree

//type TreeNode struct {
//	Data  string    // 节点用来存放数据
//	Left  *TreeNode // 左子树
//	Right *TreeNode // 右字树
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
		//print(cur.Data)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}

//https://leetcode.cn/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	var myQueue []*TreeNode
	myQueue = append(myQueue, root)
	level := 0
	for len(myQueue) > 0 {
		var tempQueue []*TreeNode
		arrLevel := make([]int, 0)
		for len(myQueue) > 0 {
			cur := myQueue[0]
			myQueue = myQueue[1:]
			if cur.Left != nil {
				tempQueue = append(tempQueue, cur.Left)
			}
			if cur.Right != nil {
				tempQueue = append(tempQueue, cur.Right)
			}
			arrLevel = append(arrLevel, cur.Val)
		}
		for _, a := range tempQueue {
			myQueue = append(myQueue, a)
		}
		ret = append(ret, arrLevel)
		level++
	}
	return ret
}

func BiSearch(list []int64, num int64) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] > num {
			high = mid - 1
		} else if list[mid] < num {
			low = mid + 1
		} else {
			for list[mid] == list[mid+1] {
				mid = mid + 1
			}
			return mid + 1
		}
	}
	print("out")
	//没有找到
	return low
}

// PreOrder 二叉树先序非递归
func PreOrder(head *TreeNode) {
	if head == nil {
		return
	}
	var myStack []*TreeNode
	myStack = append(myStack, head)
	for len(myStack) > 0 {
		cur := myStack[len(myStack)-1]
		myStack = myStack[:len(myStack)-1]
		//print(cur.Data+"-")
		if cur.Right != nil {
			myStack = append(myStack, cur.Right)
		}
		if cur.Left != nil {
			myStack = append(myStack, cur.Left)
		}
	}
	return
}

// MidOrder 二叉树中序非递归
func MidOrder(head *TreeNode) {
	if head == nil {
		return
	}
	var myStack []*TreeNode
	//myStack = append(myStack, head)
	temp := head
	for len(myStack) > 0 || temp != nil {
		for temp != nil {
			myStack = append(myStack, temp)
			temp = temp.Left
		}
		if len(myStack) > 0 {
			temp = myStack[len(myStack)-1]
			//print(temp.Data+"-")
			myStack = myStack[:len(myStack)-1]
			temp = temp.Right
		}
	}
	return
}

// PostOrder 二叉树后序非递归
func PostOrder(head *TreeNode) {
	if head == nil {
		return
	}
	var myStack []*TreeNode
	temp := head
	for {
		for temp != nil {
			myStack = append(myStack, temp)
			temp = temp.Left
		}
		flag := 1
		var p *TreeNode
		for len(myStack) > 0 && flag == 1 {
			temp = myStack[len(myStack)-1]
			if temp.Right == p {
				//print(temp.Data+"-")
				myStack = myStack[:len(myStack)-1]
				p = temp
			} else {
				flag = 0
				temp = temp.Right
			}
		}
		if len(myStack) == 0 {
			break
		}
	}
	return
}

func ReBuildTreeByPreMidOrder(pres, mids []int) (head *TreeNode) {
	if len(pres) == 0 || len(mids) == 0 {
		return
	}
	head = &TreeNode{Val: pres[0]}
	mid := 0

	for k, v := range mids {
		if v == pres[0] {
			mid = k
			break
		}
	}
	midLeft := mids[:mid]
	midRight := mids[mid+1:]

	preLeft := pres[1 : mid+1]
	preRight := pres[mid+1:]

	head.Left = ReBuildTreeByPreMidOrder(preLeft, midLeft)
	head.Right = ReBuildTreeByPreMidOrder(preRight, midRight)
	return
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	head := &TreeNode{Val: preorder[0]}
	mid := 0

	for k, v := range inorder {
		if v == preorder[0] {
			mid = k
			break
		}
	}
	midLeft := inorder[:mid]
	midRight := inorder[mid+1:]

	preLeft := preorder[1 : mid+1]
	preRight := preorder[mid+1:]

	head.Left = buildTree(preLeft, midLeft)
	head.Right = buildTree(preRight, midRight)
	return head
}
