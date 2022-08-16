package binarytree

//二叉搜索树的第k大节点
//https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/

//再次温故中序便利
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	cur := root
	//start := 0
	myStack := make([]*TreeNode, 0)
	myArray := make([]int,0)
	for len(myStack) > 0 || cur != nil {
		for cur != nil {
			myStack = append(myStack, cur)
			cur = cur.Left
		}
		if len(myStack) > 0 {
			nowCur := myStack[len(myStack)-1]
			myArray = append(myArray, nowCur.Val)
			myStack = myStack[:len(myStack)-1]
			//start += 1
			//if start == k {
			//	return nowCur.Val
			//}
			cur = nowCur.Right
		}
	}
	//fmt.Println(myArray)
	if len(myArray) >= k{
		return myArray[len(myArray)-k]
	}
	return -1
}
