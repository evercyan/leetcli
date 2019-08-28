# coding:utf-8


class Solution:

    def twoSum(self, nums, target):
        result = []
        count = len(nums)
        for i in range(count - 1):
            for j in range(i + 1, count):
                if nums[i] + nums[j] == target:
                    result = [i, j]
                    break
        return result


solution = Solution()
nums = [2, 4, 3]
target = 6
ret = solution.twoSum(nums, target)
print(ret)
