package mysort

func insertSort(nums []int) []int {
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

func insertSortByHalf(nums []int) []int {
	for i:=1 ; i< len(nums) ; i++{
		temp := nums[i]
		j := i-1
		left,right := 0, i-1
		for ;left<=right; {
			mid := (left+right)/2
			if nums[mid] < temp{
				left = mid + 1
			}else {
				right = mid - 1
			}
		}
		start := j-1
		for ;start>=right+1;start--{
			nums[start+1] = nums[start]
		}
		nums[start+1] = temp
	}
	return nums
}

func insertSortByShell(nums []int) []int {
	n := len(nums)
	d := n/2
	for ;d>0;d/=2{
		for i:=d;i<n;i++{
			temp := nums[i]
			start := i-d
			for ;start>=0 && temp<nums[start];{
				nums[start+d] = nums[start]
				start-=d
			}
			nums[start+d] = temp
		}
	}
	return nums
}

func selectSortBySimple(nums []int) []int {
	n := len(nums)
	for i := 0;i< n-1;i++{
		temp := nums[i]
		right := -1
		for j:=i+1;j<n;j++{
			if nums[j] < temp{
				right = j
			}
		}
		if right != -1{
			nums[i] = nums[right]
			nums[right] = temp
		}
	}
	return nums
}

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0;i< n-1;i++{
		exchange := 0
		for j:=n-1;j>i;j--{
			if nums[j] < nums[j-1]{
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
				exchange = 1
			}
		}
		if exchange == 0{
			break
		}
	}
	return nums
}

func quickSort(nums []int,s, e int){
	i,j := s,e
	if s < e{
		temp := nums[s]
		for ;i < j; {
			for ; i < j && nums[j] >= temp; {
				j--
			}
			nums[i] = nums[j]
			for ; i < j && nums[i] <= temp; {
				i++
			}
			nums[j] = nums[i]
		}
		nums[i] = temp
		quickSort(nums,s,i-1)
		quickSort(nums,i+1,e)
	}
}