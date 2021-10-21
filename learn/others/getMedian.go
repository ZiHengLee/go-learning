package others

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	len1 := len(nums1)
//	len2 := len(nums2)
//	allLength := len1 + len2
//	var nums3 []int
//	if len1 > len2{
//		nums3 = nums1
//		nums1 = nums2
//		nums2 = nums3
//	}
//	len1 = len(nums1)
//	len2 = len(nums2)
//	if len1 == 0{
//		if len2 % 2 ==0{
//			return float64(nums2[len2/2]+nums2[(len2/2)-1])/2
//		}else{
//			return float64(nums2[len2/2])
//		}
//	}
//	halfLen := allLength/2
//	left := 0
//	right := len1
//	if allLength%2 == 0 {
//		if len1 == 1 && len2 == 1{
//			return float64(nums1[0]+nums2[0])/2
//		}else if len1 == 1 && len2 > 1{
//			leftNums2 := halfLen-1
//			if nums1[0] < nums2[leftNums2-1] || nums1[0] > nums2[leftNums2+1]{
//				if  nums1[0] > nums2[leftNums2+1]{
//					return float64(nums2[leftNums2+1]+nums2[leftNums2])/2
//				}else{
//					return float64(nums2[leftNums2-1]+nums2[leftNums2])/2
//				}
//			}else{
//				return float64(nums2[leftNums2]+nums1[0])/2
//			}
//		}
//		for left <= right{
//			mid := (left+right)/2
//			if mid == len1 || mid == 0 {
//				break
//			}
//			leftNums2 := nums2[halfLen-mid-1]
//			rightNums2 := nums2[halfLen-mid]
//
//			leftNums1 := nums1[mid-1]
//			rightNums1 := nums1[mid]
//
//			if rightNums2 >= leftNums1 && rightNums1 >= leftNums2 {
//				return (math.Max(float64(leftNums1), float64(leftNums2))+math.Min(float64(rightNums1), float64(rightNums2)))/2
//			}else if leftNums2 > rightNums1{
//				left = mid+1
//			}else{
//				right = mid
//			}
//		}
//		if right == len1{
//			if len1 < halfLen{
//				return float64(float64(nums2[halfLen-len1]) + math.Max(float64(nums2[halfLen-len1-1]),float64(nums1[len1-1])))/2
//			}else{
//				return float64(nums2[0] + nums1[halfLen-1])/2
//			}
//		}else if left == 0{
//			//fmt.Println(halfLen, len1)
//			if len1 < halfLen{
//				return float64(float64(nums2[halfLen-1]) + math.Min(float64(nums2[halfLen]),float64(nums1[0])))/2
//			}else{
//				return float64(nums2[len2-1] + nums1[0])/2
//			}
//		}
//	}else{
//		half := len2/2
//		if len1 == 1 {
//			if nums1[0] < nums2[half-1]{
//				return float64(nums2[half-1])
//			}else if nums1[0] > nums2[half]{
//				return float64(nums2[half])
//			}else{
//				return float64(nums1[0])
//			}
//		}
//		for left <= right{
//			mid := (left+right)/2
//			if mid == len1 || mid == 0 {
//				break
//			}
//			leftNums2 := nums2[halfLen-mid-1]
//			rightNums2 := nums2[halfLen-mid]
//
//			leftNums1 := nums1[mid-1]
//			rightNums1 := nums1[mid]
//
//			if rightNums2 >= leftNums1 && rightNums1 >= leftNums2 {
//				return math.Max(math.Max(float64(leftNums1), float64(leftNums2)),math.Min(float64(rightNums1), float64(rightNums2)))
//			}else if leftNums2 > rightNums1{
//				left = mid+1
//			}else{
//				right = mid
//			}
//		}
//		if right == len1{
//			return float64(nums2[halfLen-len1])
//		}else if left == 0{
//			return math.Min(float64(nums1[0]),float64(nums2[halfLen]))
//		}
//	}
//	return 0
//}

//思路
//时间复杂度要求O(log2(m+n))
//归并排序再找到中位数不满足时间复杂度
//问题可以理解为两个排好序的数组找切割位置，找切割位置的过程可以理解为二分查找
//  2 3 | 4 5
//2 3 4 | 5 6 7
//然后通过比较切割位置两边 数的大小找到中位数
//额外需要考虑两个问题，1.奇数个数和偶数个数中位数的求法不通 2.极端位置的处理，切割位置在最左边和最右边
//对短数组二分查找，m<n
//时间复杂度O(log2(m+1))

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	len1 := len(nums1)
//	len2 := len(nums2)
//	if len1 > len2{
//		nums3 := nums1
//		nums1 = nums2
//		nums2 = nums3
//		len1 = len(nums1)
//		len2 = len(nums2)
//	}
//	allLength := len1 + len2
//	halfLen := allLength/2
//	if len1 == 1 && len2 == 1{
//		return float64(float64(nums1[0])+float64(nums2[0]))/2
//	}else if len1 == 1{
//		if allLength % 2==0{
//			numsMid := nums2[len2/2]
//			numsLeft := nums2[len2/2-1]
//			numsRight := nums2[len2/2+1]
//			//fmt.Println(nums1[0],numsLeft,)
//			if nums1[0]>numsLeft && nums1[0]<=numsMid{
//				return float64(numsMid+nums1[0])/2
//			}else if nums1[0]<numsRight && nums1[0]>=numsMid{
//				return float64(numsMid+nums1[0])/2
//			}else if nums1[0] <= numsLeft{
//				return float64(numsMid+numsLeft)/2
//			}else {
//				return float64(numsMid+numsRight)/2
//			}
//		}else{
//			numsLeft := nums2[len2/2-1]
//			numsRight := nums2[len2/2]
//			if nums1[0] <= numsLeft{
//				return float64(numsLeft)
//			}else if nums1[0] >= numsRight{
//				return float64(numsRight)
//			}else{
//				return float64(nums1[0])
//			}
//		}
//	}
//	if len1 == 0{
//		if len2 % 2 ==0{
//			return float64(nums2[len2/2]+nums2[(len2/2)-1])/2
//		}else{
//			return float64(nums2[len2/2])
//		}
//	}
//	left := 0
//	right := len1
//
//	mid := (left+right)/2
//
//	for left <= right {
//		mid = (left+right)/2
//		//fmt.Println(left,right,mid,len1,halfLen)
//		if mid == 0 || mid == len1{
//			if mid == 0{
//				if allLength % 2==0{
//					if halfLen > len1 {
//						rightNums2 := math.Min(float64(nums2[halfLen-mid]), float64(nums1[0]))
//						leftNums1 := float64(nums2[halfLen-mid-1])
//						return float64(leftNums1+rightNums2) / 2
//					}else{
//						leftNums1 := nums1[0]
//						rightNums2 := float64(nums2[halfLen-mid-1])
//						return (float64(leftNums1)+float64(rightNums2))/2
//					}
//				}else{
//					leftNums2 := nums2[halfLen-mid-1]
//					rightNums2 := nums2[halfLen-mid]
//					Nums1 := nums1[mid]
//					return math.Max(math.Min(float64(Nums1), float64(rightNums2)),float64(leftNums2))
//				}
//			}else if mid == len1{
//				if allLength % 2==0{
//					if halfLen > mid{
//						leftNums2 := math.Max(float64(nums2[halfLen-mid-1]),float64(nums1[len1-1]))
//						rightNums2 := nums2[halfLen-mid]
//						return (float64(leftNums2)+float64(rightNums2))/2
//					}else{
//						leftNums1 := nums1[len1-1]
//						rightNums2 := nums2[0]
//						return (float64(leftNums1)+float64(rightNums2))/2
//					}
//				}else{
//					leftNums2 := nums2[halfLen-mid]
//					rightNums2 := nums2[halfLen-mid+1]
//					Nums1 := nums1[mid-1]
//					return math.Min(math.Max(float64(Nums1), float64(leftNums2)),float64(rightNums2))
//				}
//			}
//		}
//		leftNums2 := nums2[halfLen-mid-1]
//		rightNums2 := nums2[halfLen-mid]
//
//		leftNums1 := nums1[mid-1]
//		rightNums1 := nums1[mid]
//		if rightNums2 >= leftNums1 && rightNums1 >= leftNums2 {
//			if allLength % 2==0{
//				return (math.Max(float64(leftNums1), float64(leftNums2))+math.Min(float64(rightNums1), float64(rightNums2)))/2
//			}else{
//				return math.Max(math.Max(float64(leftNums1), float64(leftNums2)),math.Min(float64(rightNums1), float64(rightNums2)))
//			}
//		}else if leftNums2 > rightNums1{
//			left = mid+1
//		}else{
//			right = mid
//		}
//	}
//	return 0
//}

//把问题简化为寻找第k小数的问题，每次把两段数组从中间分开，比较边界位置数的大小，
//对与小于第k小数的那些数据（假设m个）可以删除，那么问题就变成了寻找第k-m个数的问题

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	left,right := (len1+len2+1)/2,(len1+len2+2)/2
	return float64(findKthElement(nums1,nums2,left)+findKthElement(nums1,nums2,right))/2

}

func findKthElement(nums1, nums2 []int, k int) int {
	if len(nums1) > len(nums2) {
		return findKthElement(nums2, nums1, k)
	}
	if len(nums1) == 0{
		return nums2[k-1]
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}
	len1, len2 := len(nums1), len(nums2)
	i, j := minInt(k/2, len1)-1, minInt(k/2, len2)-1
	if nums1[i] > nums2[j]{
		return findKthElement(nums1, nums2[j+1:], k-j-1)
	}else{
		return findKthElement(nums1[i+1:], nums2, k-i-1)
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
