function twoSum(nums, target) {
  var tmp = {};
  for (var i = 0; i < nums.length; i++) {
    var tmp1 = target - nums[i];
    if (tmp1 in tmp) {
      var tmp2 = tmp[tmp1];
      return [tmp2, i];
    }
    tmp[nums[i]] = i;
  }
  return null;
}
var nums1 = [2, 7, 11, 15];
var target1 = 9;
console.log(twoSum(nums1, target1));
