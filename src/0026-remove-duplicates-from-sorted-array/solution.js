/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    if (nums.length <= 1) {
        return nums.length;
    }
    var num = nums[0];
    for (var i = 1; i < nums.length; i++) {
        if (nums[i] != num) {
            num = nums[i];
            continue;
        }
        nums.splice(i, 1);
        i--;
    }
    return nums.length;
};
