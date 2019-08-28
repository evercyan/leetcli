# coding:utf-8

# 给定 n = 5，并且 version = 4 是第一个错误的版本。

# 调用 isBadVersion(3) -> false
# 调用 isBadVersion(5) -> true
# 调用 isBadVersion(4) -> true

# 所以，4 是第一个错误的版本。


class Solution:

    def firstBadVersion(self, n):
        low = 0
        high = n
        while low < high:
            mid = (low + high) // 2
            if not isBadVersion(mid):
                low = mid + 1
            else:
                high = mid
        return high
