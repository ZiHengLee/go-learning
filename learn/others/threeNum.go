package others

//思路：排好序的数组[n1,n2,,,nm]，从头n1开始便利，分别设置前后指针指向n2和nm
//计算和与0比较，小于0，右移左指针，大于0，左移右指针，直至左右指针相遇
//关键点就是去重，重复的值一定是相邻的

//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
func threeSum(nums []int) [][]int {
	nums = sort(nums)
	ret := make([][]int, 0)
	for i:=0;i< len(nums)-2;i++{
		if i >= 1 && nums[i] == nums[i-1]{
			continue
		}
		tempArray := make([]int,0)
		if nums[i] > 0{
			break
		}
		tempArray = append(tempArray, nums[i])
		left := i+1
		right := len(nums)-1
		for ;left<right;{
			if left > i+1 && nums[left] == nums[left-1]{
				left+=1
				continue
			}else if right < len(nums)-1 && nums[right] == nums[right+1]{
				right-=1
				continue
			}
			//fmt.Println(nums[i],nums[left],nums[right])
			if nums[i] + nums[left] +nums[right] == 0 {
				helpArray := tempArray
				helpArray = append(helpArray, nums[left], nums[right])
				ret = append(ret, helpArray)
				left+=1
				right-=1
			}else if nums[i] + nums[left] +nums[right] < 0{
				left+=1
			}else{
				right-=1
			}
		}
	}
	return ret
}
func sort(nums []int) []int {
	for i:=1 ; i< len(nums) ; i++{
		temp := nums[i]
		j := i-1
		for ;j>=0&&temp<nums[j];{
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
	return nums
}