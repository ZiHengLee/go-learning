function threeSum(nums: number[]): number[][] | null {
  const lenNums = nums.length;
  nums.sort((a, b) => a - b);
  let ret: number[][] = [];
  for (let i = 0; i < lenNums - 2; i++) {
    if (i >= 1 && nums[i] == nums[i - 1]) {
      continue;
    }
    let left = i + 1,
      right = lenNums - 1;
    while (left < right) {
      if (left > i + 1 && nums[left] == nums[left + 1]) {
        left++;
        continue;
      }

      if (right < lenNums - 1 && nums[right] == nums[right - 1]) {
        right--;
        continue;
      }

      if (nums[right] + nums[left] + nums[i] == 0) {
        ret.push([nums[left], nums[right], nums[i]]);
        left++;
        right--;
        continue;
      } else if (nums[right] + nums[left] + nums[i] < 0) {
        left++;
        continue;
      } else {
        right--;
        continue;
      }
    }
  }
  return ret;
}

const nums = [-1, 0, 1, 2, -1, -4];
console.log(threeSum(nums));
