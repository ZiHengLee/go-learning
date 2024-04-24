

def threeSum(nums):
    nums.sort()
    lenNums = len(nums)
    ret = []
    for i in range(lenNums-2):
        if i>=1 and nums[i]==nums[i-1]:
            continue
        left,right = i+1,lenNums-1
        while left < right:
            if left > i+1 and nums[left]==nums[left+1]:
                left+=1
                continue
            if right < lenNums-1 and nums[right]==nums[right+1]:
                right-=1
                continue

            if nums[left]+nums[right]+nums[i]==0:
                ret.append([nums[left],nums[right],nums[i]])
                # 这里是可以继续往下走的
                left+=1
                right-=1
                continue
            elif nums[left]+nums[right]+nums[i]<0:
                left+=1
                continue
            else:
                right-=1
                continue
    return ret

nums = [-1, 0, 1, 2, -1, -4];
print(threeSum(nums));