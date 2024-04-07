package linkcode

//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6

//思路最小堆的问题
func mergeKLists(lists []*ListNode) *ListNode {
	nums := make([]*ListNode, 0)
	for _, list := range lists {
		if list != nil {
			nums = append(nums, list)
		}
	}
	total := len(nums)

	retList := &ListNode{-1, nil}
	curRet := retList
	for i := total/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, total)
	}
	curMin := nums[0]
	for curMin != nil {
		newNode := &ListNode{nums[0].Val, nil}
		curRet.Next = newNode
		curRet = curRet.Next
		if curMin.Next != nil {
			curMin = curMin.Next
		} else {

		}
		nums[0] = curMin
		adjustHeap(nums, 0, total)
	}
	for j := total - 1; j >= 0; j-- {
		newNode := &ListNode{nums[0].Val, nil}
		curRet.Next = newNode
		curRet = curRet.Next
		temp := nums[j]
		nums[j] = nums[0]
		nums[0] = temp
		adjustHeap(nums, 0, j)
	}
	return retList.Next
}

func adjustHeap(nums []*ListNode, start, len int) {
	i, j := start, start*2+1
	temp := nums[i]
	for j < len {
		if j+1 < len && nums[j+1].Val < nums[j].Val {
			j += 1
		}
		if nums[j].Val < temp.Val {
			nums[i] = nums[j]
			i = j
			j = i*2 + 1
		} else {
			break
		}
	}
	nums[i] = temp
	return
}

//用归并的思路
func mergeKListsV1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge2List(lists, 0, len(lists)-1)
}

func merge2List(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	mid := (start + end) / 2
	left := merge2List(lists, start, mid)
	right := merge2List(lists, mid+1, end)
	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	ret := &ListNode{Val: -1, Next: nil}
	cur := ret
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			cur = left
			left = left.Next
		} else {
			cur.Next = right
			cur = right
			right = right.Next
		}
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return ret.Next
}
