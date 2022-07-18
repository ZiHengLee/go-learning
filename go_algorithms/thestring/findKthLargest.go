package thestring

func heapSort(nums []int, k int) int{
	k-=1
	lenN := len(nums)
	//从第一个非叶子节点开始
	for i := lenN/2-1;i>=0;i--{
		adjustHeap(nums,i,lenN)
	}
	start := 0
	for j := lenN-1;j>=0;j--{
		temp := nums[0]
		nums[0] = nums[j]
		nums[j] = temp
		//if start == k{
		//	return temp
		//}
		adjustHeap(nums,0,j)
		start+=1
	}
	print(nums)
	return -1
}

func adjustHeap(nums []int,start,len int){
	temp := nums[start]
	for k := start*2+1;k<len;k=k*2+1{
		if k+1< len && nums[k+1]> nums[k] {
			k+=1
		}
		if nums[k] > nums[start]{
			nums[start] = nums[k]
			start = k
		}else{
			break
		}
	}
	nums[start] = temp
	return
}

