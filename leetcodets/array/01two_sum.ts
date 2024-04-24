function twoSum(nums: number[], target: number): number[] | null {
  const tmp: { [key: number]: number } = {};
  for (let i = 0; i < nums.length; i++) {
    const tmp1 = target - nums[i];
    if (tmp1 in tmp) {
      const tmp2 = tmp[tmp1];
      return [tmp2, i];
    }
    tmp[nums[i]] = i;
  }
  return null;
}

// 示例
const nums1 = [2, 7, 11, 15];
const target1 = 9;
console.log(twoSum(nums1, target1));
