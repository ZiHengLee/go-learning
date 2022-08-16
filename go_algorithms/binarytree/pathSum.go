package binarytree

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil{
		return nil
	}
	ret := make([][]int,0)
	cur := make([]int,0)
	if root.Left == nil && root.Right == nil && targetSum==root.Val{
		cur = append(cur, root.Val)
		ret = append(ret, cur)
		return ret
	}
	cur = append(cur, root.Val)
	if root.Left != nil{
		dfsPath(root.Left,targetSum-root.Val,cur,&ret)
	}
	if root.Right != nil{
		dfsPath(root.Right,targetSum-root.Val,cur,&ret)
	}
	return ret
}

func dfsPath(node *TreeNode,targetSum int,curArry []int,ret *[][]int){
	if node.Val == targetSum && node.Left == nil && node.Right==nil{
		curArry = append(curArry, node.Val)
		temp := make([]int,len(curArry))
		copy(temp,curArry)
		*ret = append(*ret, temp)
		return
	}
	newLeft,newRight := make([]int,len(curArry)), make([]int,len(curArry))
	copy(newLeft,curArry)
	copy(newRight,curArry)

	if node.Left != nil{
		newLeft = append(newLeft, node.Val)
		dfsPath(node.Left,targetSum-node.Val,newLeft,ret)
	}
	if node.Right != nil{
		newRight = append(newRight, node.Val)
		dfsPath(node.Right,targetSum-node.Val,newRight,ret)
	}
	return
}
