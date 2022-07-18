package others

//https://leetcode.cn/problems/search-in-rotated-sorted-array/
//搜索旋转排序数组
//[4,5,6,7,0,1,2,3] 0 -> 4
//思路

//向左
//1.nums[0]<=nums[mid] && nums[0]<=target<=nums[mid]
//2.nums[0]>nums[mid] && (nums[mid] < nums[0] <= target || target <= nums[mid] < nums[0])
//
//其他向右
//抽象出来
//只有一个为真的时候向右
//nums[0] > nums[mid]
//nums[0] > target
//target > nums[mid]

func search(nums []int, target int) int {
	left,right := 0,len(nums)-1
	for left < right{
		mid := (left+right)/2
		con1,con2,con3:=0,0,0
		if nums[0]>nums[mid]{
			con1 = 1
		}
		if nums[0]>target{
			con2=1
		}
		if target>nums[mid]{
			con3=1
		}
		if con1^con2^con3==1{
			left = mid+1
		}else{
			right = mid
		}
	}
	if nums[left] == target{
		return left
	}
	return -1
}