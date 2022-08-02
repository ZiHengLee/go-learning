package binarytree

//最近公共祖先
//思路：
//递归思路，可以把问题简化为最简单的二叉树
//2个节点和3个节点


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	if p==root || q== root{
		return root
	}
	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)
	if left != nil && right != nil{
		return root
	}else if left != nil{
		return left
	}else if right != nil{
		return right
	}
	return nil
}