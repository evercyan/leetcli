# coding:utf-8

# 你调用一个预先定义好的接口 guess(int num)，它会返回 3 个可能的结果（-1，1 或 0）：
# -1 : 我的数字比较小
# 1 : 我的数字比较大
# 0 : 恭喜！你猜对了！


# 二分二分二分


class Solution(object):

    def guessNumber(self, n):
        low = 0
        high = n
        while low <= high:
            mid = (low + high) // 2
            guessret = guess(mid)
            if guessret == 0:
                return mid
            elif guessret == -1:
                high = mid - 1
            else:
                low = mid + 1
        return high
