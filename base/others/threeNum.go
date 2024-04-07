package others

//思路：排好序的数组[n1,n2,,,nm]，从头n1开始便利，分别设置前后指针指向n2和nm
//计算和与0比较，小于0，右移左指针，大于0，左移右指针，直至左右指针相遇
//关键点就是去重，重复的值一定是相邻的

//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//func threeSum(nums []int) [][]int {
//	nums = sort(nums)
//	ret := make([][]int, 0)
//	for i := 0; i < len(nums)-2; i++ {
//		if i >= 1 && nums[i] == nums[i-1] {
//			continue
//		}
//		tempArray := make([]int, 0)
//		if nums[i] > 0 {
//			break
//		}
//		tempArray = append(tempArray, nums[i])
//		left := i + 1
//		right := len(nums) - 1
//		for left < right {
//			if left > i+1 && nums[left] == nums[left-1] {
//				left += 1
//				continue
//			} else if right < len(nums)-1 && nums[right] == nums[right+1] {
//				right -= 1
//				continue
//			}
//			if nums[i]+nums[left]+nums[right] == 0 {
//				helpArray := tempArray
//				helpArray = append(helpArray, nums[left], nums[right])
//				ret = append(ret, helpArray)
//				left += 1
//				right -= 1
//			} else if nums[i]+nums[left]+nums[right] < 0 {
//				left += 1
//			} else {
//				right -= 1
//			}
//		}
//	}
//	return ret
//}
//func sort(nums []int) []int {
//	for i := 1; i < len(nums); i++ {
//		temp := nums[i]
//		j := i - 1
//		for j >= 0 && temp < nums[j] {
//			nums[j+1] = nums[j]
//			j--
//		}
//		nums[j+1] = temp
//	}
//	return nums
//}


//v2版本20220716
//1.数组需要赋值
//2.重复的数字的解决办法

func threeSum(nums []int) [][]int {
	lenNums := len(nums)
	quickSort(nums,0,lenNums-1)
	ret := make([][]int,0)
	for i:=0;i<lenNums-2;i++{
		if i>=1&&nums[i]==nums[i-1]{
			continue
		}
		tempArr := make([]int,0)
		left,right := i+1,lenNums-1
		tempArr = append(tempArr,nums[i])
		for left < right{
			//边界条件也很关键
			if left>i+1 && nums[left]==nums[left-1]{
				left++
				continue
			}else if right+1<lenNums&&nums[right]==nums[right+1]{
				right--
				continue
			}
			if nums[i]+nums[left]+nums[right] == 0{
				helpArray := tempArr
				helpArray = append(helpArray,nums[left],nums[right])
				ret = append(ret,helpArray)
				left++
				right--
			}else if nums[i]+nums[left]+nums[right] > 0{
				right--
			}else{
				left++
			}
		}
	}
	return ret
}

func quickSort(nums []int, s, e int) {
	i, j := s, e
	if s < e {
		swap(nums,s,(s+e)/2)
		temp := nums[s]
		for i < j {
			for i < j && nums[j] >= temp {
				j--
			}
			nums[i] = nums[j]
			for i < j && nums[i] <= temp {
				i++
			}
			nums[j] = nums[i]
		}
		nums[i] = temp
		quickSort(nums, s, i-1)
		quickSort(nums, i+1, e)
	}
}

func swap(nums []int, left, right int){
	temp := nums[left]
	nums[left] = nums[right]
	nums[right] = temp
}
