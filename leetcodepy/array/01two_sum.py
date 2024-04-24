def two_sum(nums, target):
    tmp = {}
    for i, num in enumerate(nums):
        tmp1 = target - num
        if tmp1 in tmp:
            return [tmp[tmp1], i]
        tmp[num] = i
    return None

if __name__ == '__main__':
    nums = [1, 2, 3, 4, 5]
    target = 10
    print(two_sum(nums, target))
