# leetcli

[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
[![Build Status](https://www.travis-ci.org/evercyan/leetcli.svg?branch=master)](https://www.travis-ci.org/evercyan/leetcli)
[![codecov](https://codecov.io/gh/evercyan/leetcli/branch/master/graph/badge.svg?token=RbJTUtAlvl)](https://codecov.io/gh/evercyan/leetcli)

> leetcode 刷题小助手, 帮助生成题目 readme, 答题文件, 答题测试文件等.

---

#### 帮助

```
NAME:
	leetcli - A cli tool for leetcode

USAGE:
	helper [global options] command [command options] [arguments...]

VERSION:
	1.0.0

COMMANDS:
	readme   生成 README.md
	problem  生成答题文件相关 [eg: problem 101]
	quit     退出程序
	help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
	--help, -h     show help
	--version, -v  print the version
```

---

#### 题目标签

[![数组](https://img.shields.io/badge/数组-304-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/array/)
[![动态规划](https://img.shields.io/badge/动态规划-230-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/dynamic-programming/)
[![字符串](https://img.shields.io/badge/字符串-207-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/string/)
[![数学](https://img.shields.io/badge/数学-201-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/math/)
[![树](https://img.shields.io/badge/树-165-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/tree/)
[![深度优先搜索](https://img.shields.io/badge/深度优先搜索-143-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/depth-first-search/)
[![哈希表](https://img.shields.io/badge/哈希表-137-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/hash-table/)
[![二分查找](https://img.shields.io/badge/二分查找-95-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/binary-search/)
[![贪心算法](https://img.shields.io/badge/贪心算法-85-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/greedy/)
[![广度优先搜索](https://img.shields.io/badge/广度优先搜索-77-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/breadth-first-search/)
[![双指针](https://img.shields.io/badge/双指针-70-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/two-pointers/)
[![位运算](https://img.shields.io/badge/位运算-65-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/bit-manipulation/)
[![排序](https://img.shields.io/badge/排序-64-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/sort/)
[![栈](https://img.shields.io/badge/栈-64-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/stack/)
[![回溯算法](https://img.shields.io/badge/回溯算法-62-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/backtracking/)
[![设计](https://img.shields.io/badge/设计-61-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/design/)
[![链表](https://img.shields.io/badge/链表-53-orange.svg?style=flat)](https://leetcode-cn.com/tag/linked-list/)
[![图](https://img.shields.io/badge/图-46-orange.svg?style=flat)](https://leetcode-cn.com/tag/graph/)
[![堆](https://img.shields.io/badge/堆-40-blue.svg?style=flat)](https://leetcode-cn.com/tag/heap/)
[![并查集](https://img.shields.io/badge/并查集-31-blue.svg?style=flat)](https://leetcode-cn.com/tag/union-find/)
[![SlidingWindow](https://img.shields.io/badge/SlidingWindow-29-blue.svg?style=flat)](https://leetcode-cn.com/tag/sliding-window/)
[![分治算法](https://img.shields.io/badge/分治算法-28-blue.svg?style=flat)](https://leetcode-cn.com/tag/divide-and-conquer/)
[![递归](https://img.shields.io/badge/递归-23-blue.svg?style=flat)](https://leetcode-cn.com/tag/recursion/)
[![字典树](https://img.shields.io/badge/字典树-18-red.svg?style=flat)](https://leetcode-cn.com/tag/trie/)
[![线段树](https://img.shields.io/badge/线段树-14-red.svg?style=flat)](https://leetcode-cn.com/tag/segment-tree/)
[![几何](https://img.shields.io/badge/几何-11-red.svg?style=flat)](https://leetcode-cn.com/tag/geometry/)
[![队列](https://img.shields.io/badge/队列-11-red.svg?style=flat)](https://leetcode-cn.com/tag/queue/)
[![OrderedMap](https://img.shields.io/badge/OrderedMap-10-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/ordered-map/)
[![极小化极大](https://img.shields.io/badge/极小化极大-8-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/minimax/)
[![树状数组](https://img.shields.io/badge/树状数组-6-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/binary-indexed-tree/)
[![脑筋急转弯](https://img.shields.io/badge/脑筋急转弯-6-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/brainteaser/)
[![LineSweep](https://img.shields.io/badge/LineSweep-6-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/line-sweep/)
[![Random](https://img.shields.io/badge/Random-6-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/random/)
[![拓扑排序](https://img.shields.io/badge/拓扑排序-6-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/topological-sort/)
[![二叉搜索树](https://img.shields.io/badge/二叉搜索树-5-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/binary-search-tree/)
[![记忆化](https://img.shields.io/badge/记忆化-3-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/memoization/)
[![RejectionSampling](https://img.shields.io/badge/RejectionSampling-2-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/rejection-sampling/)
[![蓄水池抽样](https://img.shields.io/badge/蓄水池抽样-2-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/reservoir-sampling/)


---

#### 题目列表

|#|标题|难度|
|:-:|:-|:-:|
|[1](https://leetcode-cn.com/problems/two-sum/)|[Two Sum](./src/0001-two-sum/)|简单|
|[2](https://leetcode-cn.com/problems/add-two-numbers/)|[Add Two Numbers](./src/0002-add-two-numbers/)|中等|
|[3](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)|[Longest Substring Without Repeating Characters](./src/0003-longest-substring-without-repeating-characters/)|中等|
|[4](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)|[Median of Two Sorted Arrays](./src/0004-median-of-two-sorted-arrays/)|困难|
|[5](https://leetcode-cn.com/problems/longest-palindromic-substring/)|[Longest Palindromic Substring](./src/0005-longest-palindromic-substring/)|中等|
|[6](https://leetcode-cn.com/problems/zigzag-conversion/)|[ZigZag Conversion](./src/0006-zigzag-conversion/)|中等|
|[7](https://leetcode-cn.com/problems/reverse-integer/)|[Reverse Integer](./src/0007-reverse-integer/)|简单|
|[8](https://leetcode-cn.com/problems/string-to-integer-atoi/)|[String to Integer (atoi)](./src/0008-string-to-integer-atoi/)|中等|
|[9](https://leetcode-cn.com/problems/palindrome-number/)|[Palindrome Number](./src/0009-palindrome-number/)|简单|
|[10](https://leetcode-cn.com/problems/regular-expression-matching/)|[Regular Expression Matching](./src/0010-regular-expression-matching/)|困难|
|[11](https://leetcode-cn.com/problems/container-with-most-water/)|[Container With Most Water](./src/0011-container-with-most-water/)|中等|
|[12](https://leetcode-cn.com/problems/integer-to-roman/)|[Integer to Roman](./src/0012-integer-to-roman/)|中等|
|[13](https://leetcode-cn.com/problems/roman-to-integer/)|[Roman to Integer](./src/0013-roman-to-integer/)|简单|
|[14](https://leetcode-cn.com/problems/longest-common-prefix/)|[Longest Common Prefix](./src/0014-longest-common-prefix/)|简单|
|[15](https://leetcode-cn.com/problems/3sum/)|[3Sum](./src/0015-3sum/)|中等|
|[16](https://leetcode-cn.com/problems/3sum-closest/)|[3Sum Closest](./src/0016-3sum-closest/)|中等|
|[17](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)|[Letter Combinations of a Phone Number](./src/0017-letter-combinations-of-a-phone-number/)|中等|
|[18](https://leetcode-cn.com/problems/4sum/)|[4Sum](./src/0018-4sum/)|中等|
|[19](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)|[Remove Nth Node From End of List](./src/0019-remove-nth-node-from-end-of-list/)|中等|
|[20](https://leetcode-cn.com/problems/valid-parentheses/)|[Valid Parentheses](./src/0020-valid-parentheses/)|简单|
|[21](https://leetcode-cn.com/problems/merge-two-sorted-lists/)|[Merge Two Sorted Lists](./src/0021-merge-two-sorted-lists/)|简单|
|[22](https://leetcode-cn.com/problems/generate-parentheses/)|[Generate Parentheses](./src/0022-generate-parentheses/)|中等|
|[23](https://leetcode-cn.com/problems/merge-k-sorted-lists/)|[Merge k Sorted Lists](./src/0023-merge-k-sorted-lists/)|困难|
|[24](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)|[Swap Nodes in Pairs](./src/0024-swap-nodes-in-pairs/)|中等|
|[25](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)|[Reverse Nodes in k-Group](./src/0025-reverse-nodes-in-k-group/)|困难|
|[26](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)|[Remove Duplicates from Sorted Array](./src/0026-remove-duplicates-from-sorted-array/)|简单|
|[27](https://leetcode-cn.com/problems/remove-element/)|[Remove Element](./src/0027-remove-element/)|简单|
|[28](https://leetcode-cn.com/problems/implement-strstr/)|[Implement strStr()](./src/0028-implement-strstr/)|简单|
|[29](https://leetcode-cn.com/problems/divide-two-integers/)|[Divide Two Integers](./src/0029-divide-two-integers/)|中等|
|[30](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)|[Substring with Concatenation of All Words](./src/0030-substring-with-concatenation-of-all-words/)|困难|
|[31](https://leetcode-cn.com/problems/next-permutation/)|[Next Permutation](./src/0031-next-permutation/)|中等|
|[32](https://leetcode-cn.com/problems/longest-valid-parentheses/)|[Longest Valid Parentheses](./src/0032-longest-valid-parentheses/)|困难|
|[33](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)|[Search in Rotated Sorted Array](./src/0033-search-in-rotated-sorted-array/)|中等|
|[34](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)|[Find First and Last Position of Element in Sorted Array](./src/0034-find-first-and-last-position-of-element-in-sorted-array/)|中等|
|[35](https://leetcode-cn.com/problems/search-insert-position/)|[Search Insert Position](./src/0035-search-insert-position/)|简单|
|[36](https://leetcode-cn.com/problems/valid-sudoku/)|[Valid Sudoku](./src/0036-valid-sudoku/)|中等|
|[37](https://leetcode-cn.com/problems/sudoku-solver/)|[Sudoku Solver](./src/0037-sudoku-solver/)|困难|
|[38](https://leetcode-cn.com/problems/count-and-say/)|[Count and Say](./src/0038-count-and-say/)|简单|
|[39](https://leetcode-cn.com/problems/combination-sum/)|[Combination Sum](./src/0039-combination-sum/)|中等|
|[40](https://leetcode-cn.com/problems/combination-sum-ii/)|[Combination Sum II](./src/0040-combination-sum-ii/)|中等|
|[41](https://leetcode-cn.com/problems/first-missing-positive/)|[First Missing Positive](./src/0041-first-missing-positive/)|困难|
|[42](https://leetcode-cn.com/problems/trapping-rain-water/)|[Trapping Rain Water](./src/0042-trapping-rain-water/)|困难|
|[43](https://leetcode-cn.com/problems/multiply-strings/)|[Multiply Strings](./src/0043-multiply-strings/)|中等|
|[44](https://leetcode-cn.com/problems/wildcard-matching/)|[Wildcard Matching](./src/0044-wildcard-matching/)|困难|
|[53](https://leetcode-cn.com/problems/maximum-subarray/)|[Maximum Subarray](./src/0053-maximum-subarray/)|简单|
|[57](https://leetcode-cn.com/problems/insert-interval/)|[Insert Interval](./src/0057-insert-interval/)|困难|
|[58](https://leetcode-cn.com/problems/length-of-last-word/)|[Length of Last Word](./src/0058-length-of-last-word/)|简单|
|[66](https://leetcode-cn.com/problems/plus-one/)|[Plus One](./src/0066-plus-one/)|简单|
|[67](https://leetcode-cn.com/problems/add-binary/)|[Add Binary](./src/0067-add-binary/)|简单|
|[69](https://leetcode-cn.com/problems/sqrtx/)|[Sqrt(x)](./src/0069-sqrtx/)|简单|
|[70](https://leetcode-cn.com/problems/climbing-stairs/)|[Climbing Stairs](./src/0070-climbing-stairs/)|简单|
|[82](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)|[Remove Duplicates from Sorted List II](./src/0082-remove-duplicates-from-sorted-list-ii/)|中等|
|[83](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)|[Remove Duplicates from Sorted List](./src/0083-remove-duplicates-from-sorted-list/)|简单|
|[88](https://leetcode-cn.com/problems/merge-sorted-array/)|[Merge Sorted Array](./src/0088-merge-sorted-array/)|简单|
|[100](https://leetcode-cn.com/problems/same-tree/)|[Same Tree](./src/0100-same-tree/)|简单|
|[101](https://leetcode-cn.com/problems/symmetric-tree/)|[Symmetric Tree](./src/0101-symmetric-tree/)|简单|
|[104](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)|[Maximum Depth of Binary Tree](./src/0104-maximum-depth-of-binary-tree/)|简单|
|[107](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)|[Binary Tree Level Order Traversal II](./src/0107-binary-tree-level-order-traversal-ii/)|简单|
|[108](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)|[Convert Sorted Array to Binary Search Tree](./src/0108-convert-sorted-array-to-binary-search-tree/)|简单|
|[110](https://leetcode-cn.com/problems/balanced-binary-tree/)|[Balanced Binary Tree](./src/0110-balanced-binary-tree/)|简单|
|[111](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)|[Minimum Depth of Binary Tree](./src/0111-minimum-depth-of-binary-tree/)|简单|
|[112](https://leetcode-cn.com/problems/path-sum/)|[Path Sum](./src/0112-path-sum/)|简单|
|[118](https://leetcode-cn.com/problems/pascals-triangle/)|[Pascal&#39;s Triangle](./src/0118-pascals-triangle/)|简单|
|[119](https://leetcode-cn.com/problems/pascals-triangle-ii/)|[Pascal&#39;s Triangle II](./src/0119-pascals-triangle-ii/)|简单|
|[121](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|[Best Time to Buy and Sell Stock](./src/0121-best-time-to-buy-and-sell-stock/)|简单|
|[122](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|[Best Time to Buy and Sell Stock II](./src/0122-best-time-to-buy-and-sell-stock-ii/)|简单|
|[125](https://leetcode-cn.com/problems/valid-palindrome/)|[Valid Palindrome](./src/0125-valid-palindrome/)|简单|
|[136](https://leetcode-cn.com/problems/single-number/)|[Single Number](./src/0136-single-number/)|简单|
|[141](https://leetcode-cn.com/problems/linked-list-cycle/)|[Linked List Cycle](./src/0141-linked-list-cycle/)|简单|
|[146](https://leetcode-cn.com/problems/lru-cache/)|[LRU Cache](./src/0146-lru-cache/)|中等|
|[155](https://leetcode-cn.com/problems/min-stack/)|[Min Stack](./src/0155-min-stack/)|简单|
|[175](https://leetcode-cn.com/problems/combine-two-tables/)|[Combine Two Tables](./src/0175-combine-two-tables/)|简单|
|[176](https://leetcode-cn.com/problems/second-highest-salary/)|[Second Highest Salary](./src/0176-second-highest-salary/)|简单|
|[177](https://leetcode-cn.com/problems/nth-highest-salary/)|[Nth Highest Salary](./src/0177-nth-highest-salary/)|中等|
|[178](https://leetcode-cn.com/problems/rank-scores/)|[Rank Scores](./src/0178-rank-scores/)|中等|
|[180](https://leetcode-cn.com/problems/consecutive-numbers/)|[Consecutive Numbers](./src/0180-consecutive-numbers/)|中等|
|[181](https://leetcode-cn.com/problems/employees-earning-more-than-their-managers/)|[Employees Earning More Than Their Managers](./src/0181-employees-earning-more-than-their-managers/)|简单|
|[182](https://leetcode-cn.com/problems/duplicate-emails/)|[Duplicate Emails](./src/0182-duplicate-emails/)|简单|
|[183](https://leetcode-cn.com/problems/customers-who-never-order/)|[Customers Who Never Order](./src/0183-customers-who-never-order/)|简单|
|[184](https://leetcode-cn.com/problems/department-highest-salary/)|[Department Highest Salary](./src/0184-department-highest-salary/)|中等|
|[185](https://leetcode-cn.com/problems/department-top-three-salaries/)|[Department Top Three Salaries](./src/0185-department-top-three-salaries/)|困难|
|[192](https://leetcode-cn.com/problems/word-frequency/)|[Word Frequency](./src/0192-word-frequency/)|中等|
|[193](https://leetcode-cn.com/problems/valid-phone-numbers/)|[Valid Phone Numbers](./src/0193-valid-phone-numbers/)|简单|
|[194](https://leetcode-cn.com/problems/transpose-file/)|[Transpose File](./src/0194-transpose-file/)|中等|
|[195](https://leetcode-cn.com/problems/tenth-line/)|[Tenth Line](./src/0195-tenth-line/)|简单|
|[196](https://leetcode-cn.com/problems/delete-duplicate-emails/)|[Delete Duplicate Emails](./src/0196-delete-duplicate-emails/)|简单|
|[197](https://leetcode-cn.com/problems/rising-temperature/)|[Rising Temperature](./src/0197-rising-temperature/)|简单|
|[198](https://leetcode-cn.com/problems/house-robber/)|[House Robber](./src/0198-house-robber/)|简单|
|[262](https://leetcode-cn.com/problems/trips-and-users/)|[Trips and Users](./src/0262-trips-and-users/)|困难|
|[278](https://leetcode-cn.com/problems/first-bad-version/)|[First Bad Version](./src/0278-first-bad-version/)|简单|
|[292](https://leetcode-cn.com/problems/nim-game/)|[Nim Game](./src/0292-nim-game/)|简单|
|[303](https://leetcode-cn.com/problems/range-sum-query-immutable/)|[Range Sum Query - Immutable](./src/0303-range-sum-query-immutable/)|简单|
|[374](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)|[Guess Number Higher or Lower](./src/0374-guess-number-higher-or-lower/)|简单|
|[383](https://leetcode-cn.com/problems/ransom-note/)|[Ransom Note](./src/0383-ransom-note/)|简单|
|[387](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)|[First Unique Character in a String](./src/0387-first-unique-character-in-a-string/)|简单|
|[415](https://leetcode-cn.com/problems/add-strings/)|[Add Strings](./src/0415-add-strings/)|简单|
|[421](https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/)|[Maximum XOR of Two Numbers in an Array](./src/0421-maximum-xor-of-two-numbers-in-an-array/)|中等|
|[448](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/)|[Find All Numbers Disappeared in an Array](./src/0448-find-all-numbers-disappeared-in-an-array/)|简单|
|[477](https://leetcode-cn.com/problems/total-hamming-distance/)|[Total Hamming Distance](./src/0477-total-hamming-distance/)|中等|
|[503](https://leetcode-cn.com/problems/next-greater-element-ii/)|[Next Greater Element II](./src/0503-next-greater-element-ii/)|中等|
|[595](https://leetcode-cn.com/problems/big-countries/)|[Big Countries](./src/0595-big-countries/)|简单|
|[596](https://leetcode-cn.com/problems/classes-more-than-5-students/)|[Classes More Than 5 Students](./src/0596-classes-more-than-5-students/)|简单|
|[600](https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/)|[Non-negative Integers without Consecutive Ones](./src/0600-non-negative-integers-without-consecutive-ones/)|困难|
|[601](https://leetcode-cn.com/problems/human-traffic-of-stadium/)|[Human Traffic of Stadium](./src/0601-human-traffic-of-stadium/)|困难|
|[620](https://leetcode-cn.com/problems/not-boring-movies/)|[Not Boring Movies](./src/0620-not-boring-movies/)|简单|
|[626](https://leetcode-cn.com/problems/exchange-seats/)|[Exchange Seats](./src/0626-exchange-seats/)|中等|
|[627](https://leetcode-cn.com/problems/swap-salary/)|[Swap Salary](./src/0627-swap-salary/)|简单|
|[739](https://leetcode-cn.com/problems/daily-temperatures/)|[Daily Temperatures](./src/0739-daily-temperatures/)|中等|
|[746](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)|[Min Cost Climbing Stairs](./src/0746-min-cost-climbing-stairs/)|简单|
|[622](https://leetcode-cn.com/problems/design-circular-queue/)|[Design Circular Queue](./src/0622-design-circular-queue/)|中等|
|[927](https://leetcode-cn.com/problems/three-equal-parts/)|[Three Equal Parts](./src/0927-three-equal-parts/)|困难|
|[976](https://leetcode-cn.com/problems/largest-perimeter-triangle/)|[Largest Perimeter Triangle](./src/0976-largest-perimeter-triangle/)|简单|
|[978](https://leetcode-cn.com/problems/longest-turbulent-subarray/)|[Longest Turbulent Subarray](./src/0978-longest-turbulent-subarray/)|中等|
|[981](https://leetcode-cn.com/problems/time-based-key-value-store/)|[Time Based Key-Value Store](./src/0981-time-based-key-value-store/)|中等|
|[982](https://leetcode-cn.com/problems/triples-with-bitwise-and-equal-to-zero/)|[Triples with Bitwise AND Equal To Zero](./src/0982-triples-with-bitwise-and-equal-to-zero/)|困难|
|[983](https://leetcode-cn.com/problems/minimum-cost-for-tickets/)|[Minimum Cost For Tickets](./src/0983-minimum-cost-for-tickets/)|中等|
|[1002](https://leetcode-cn.com/problems/find-common-characters/)|[Find Common Characters](./src/1002-find-common-characters/)|简单|
|[1009](https://leetcode-cn.com/problems/complement-of-base-10-integer/)|[Complement of Base 10 Integer](./src/1009-complement-of-base-10-integer/)|简单|
|[1011](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/)|[Capacity To Ship Packages Within D Days](./src/1011-capacity-to-ship-packages-within-d-days/)|中等|
|[1013](https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/)|[Partition Array Into Three Parts With Equal Sum](./src/1013-partition-array-into-three-parts-with-equal-sum/)|简单|
|[1014](https://leetcode-cn.com/problems/best-sightseeing-pair/)|[Best Sightseeing Pair](./src/1014-best-sightseeing-pair/)|中等|
|[1017](https://leetcode-cn.com/problems/convert-to-base-2/)|[Convert to Base -2](./src/1017-convert-to-base-2/)|中等|
|[1018](https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/)|[Binary Prefix Divisible By 5](./src/1018-binary-prefix-divisible-by-5/)|简单|
|[1019](https://leetcode-cn.com/problems/next-greater-node-in-linked-list/)|[Next Greater Node In Linked List](./src/1019-next-greater-node-in-linked-list/)|中等|
|[1020](https://leetcode-cn.com/problems/number-of-enclaves/)|[Number of Enclaves](./src/1020-number-of-enclaves/)|中等|
|[1021](https://leetcode-cn.com/problems/remove-outermost-parentheses/)|[Remove Outermost Parentheses](./src/1021-remove-outermost-parentheses/)|简单|
|[1022](https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/)|[Sum of Root To Leaf Binary Numbers](./src/1022-sum-of-root-to-leaf-binary-numbers/)|简单|
|[1287](https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/)|[Element Appearing More Than 25% In Sorted Array](./src/1287-element-appearing-more-than-25-in-sorted-array/)|简单|
|[1179](https://leetcode-cn.com/problems/reformat-department-table/)|[Reformat Department Table](./src/1179-reformat-department-table/)|简单|
|[1290](https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/)|[Convert Binary Number in a Linked List to Integer](./src/1290-convert-binary-number-in-a-linked-list-to-integer/)|简单|
|[1502](https://leetcode-cn.com/problems/can-make-arithmetic-progression-from-sequence/)|[Can Make Arithmetic Progression From Sequence](./src/1502-can-make-arithmetic-progression-from-sequence/)|简单|
|[1503](https://leetcode-cn.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/)|[Last Moment Before All Ants Fall Out of a Plank](./src/1503-last-moment-before-all-ants-fall-out-of-a-plank/)|中等|
|[1504](https://leetcode-cn.com/problems/count-submatrices-with-all-ones/)|[Count Submatrices With All Ones](./src/1504-count-submatrices-with-all-ones/)|中等|
|[面试题 16.11](https://leetcode-cn.com/problems/diving-board-lcci/)|[Diving Board LCCI](./src/100352-diving-board-lcci/)|简单|
|[面试题 17.16](https://leetcode-cn.com/problems/the-masseuse-lcci/)|[The Masseuse LCCI](./src/1000023-the-masseuse-lcci/)|简单|


---

#### 相似题型

|#|标题|难度|
|:-:|:-|:-|
|[45](https://leetcode-cn.com/problems/jump-game-ii/)|Jump Game II|困难|
|[55](https://leetcode-cn.com/problems/jump-game/)|Jump Game|中等|
|[1306](https://leetcode-cn.com/problems/jump-game-iii/)|Jump Game III|中等|
|[1345](https://leetcode-cn.com/problems/jump-game-iv/)|Jump Game IV|困难|
|[1340](https://leetcode-cn.com/problems/jump-game-v/)|Jump Game V|困难|
|[82](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)|[Remove Duplicates from Sorted List II](./src/0082-remove-duplicates-from-sorted-list-ii/)|中等|
|[83](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)|[Remove Duplicates from Sorted List](./src/0083-remove-duplicates-from-sorted-list/)|简单|
|[92](https://leetcode-cn.com/problems/reverse-linked-list-ii/)|Reverse Linked List II|中等|
|[206](https://leetcode-cn.com/problems/reverse-linked-list/)|Reverse Linked List|简单|
|[121](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|[Best Time to Buy and Sell Stock](./src/0121-best-time-to-buy-and-sell-stock/)|简单|
|[122](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|[Best Time to Buy and Sell Stock II](./src/0122-best-time-to-buy-and-sell-stock-ii/)|简单|
|[123](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|Best Time to Buy and Sell Stock III|困难|
|[188](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)|Best Time to Buy and Sell Stock IV|困难|
|[453](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/)|Minimum Moves to Equal Array Elements|简单|
|[462](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/)|Minimum Moves to Equal Array Elements II|中等|
|[490](https://leetcode-cn.com/problems/the-maze/)|The Maze|中等|
|[499](https://leetcode-cn.com/problems/the-maze-iii/)|The Maze III|困难|
|[505](https://leetcode-cn.com/problems/the-maze-ii/)|The Maze II|中等|
|[91](https://leetcode-cn.com/problems/decode-ways/)|Decode Ways|中等|
|[639](https://leetcode-cn.com/problems/decode-ways-ii/)|Decode Ways II|困难|
|[1033](https://leetcode-cn.com/problems/moving-stones-until-consecutive/)|Moving Stones Until Consecutive|简单|
|[1040](https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/)|Moving Stones Until Consecutive II|中等|
|[1136](https://leetcode-cn.com/problems/parallel-courses/)|Parallel Courses|困难|
|[1494](https://leetcode-cn.com/problems/parallel-courses-ii/)|Parallel Courses II|困难|
|[26](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)|[Remove Duplicates from Sorted Array](./src/0026-remove-duplicates-from-sorted-array/)|简单|
|[80](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/)|Remove Duplicates from Sorted Array II|中等|
|[79](https://leetcode-cn.com/problems/word-search/)|Word Search|中等|
|[212](https://leetcode-cn.com/problems/word-search-ii/)|Word Search II|困难|
|[270](https://leetcode-cn.com/problems/closest-binary-search-tree-value/)|Closest Binary Search Tree Value|简单|
|[272](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii/)|Closest Binary Search Tree Value II|困难|
|[2](https://leetcode-cn.com/problems/add-two-numbers/)|[Add Two Numbers](./src/0002-add-two-numbers/)|中等|
|[445](https://leetcode-cn.com/problems/add-two-numbers-ii/)|Add Two Numbers II|中等|
|[1148](https://leetcode-cn.com/problems/article-views-i/)|Article Views I|简单|
|[1149](https://leetcode-cn.com/problems/article-views-ii/)|Article Views II|中等|
|[1047](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/)|Remove All Adjacent Duplicates In String|简单|
|[1209](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii/)|Remove All Adjacent Duplicates in String II|中等|
|[243](https://leetcode-cn.com/problems/shortest-word-distance/)|Shortest Word Distance|简单|
|[244](https://leetcode-cn.com/problems/shortest-word-distance-ii/)|Shortest Word Distance II|中等|
|[245](https://leetcode-cn.com/problems/shortest-word-distance-iii/)|Shortest Word Distance III|中等|
|[290](https://leetcode-cn.com/problems/word-pattern/)|Word Pattern|简单|
|[291](https://leetcode-cn.com/problems/word-pattern-ii/)|Word Pattern II|困难|
|[1057](https://leetcode-cn.com/problems/campus-bikes/)|Campus Bikes|中等|
|[1066](https://leetcode-cn.com/problems/campus-bikes-ii/)|Campus Bikes II|中等|
|[1193](https://leetcode-cn.com/problems/monthly-transactions-i/)|Monthly Transactions I|中等|
|[1205](https://leetcode-cn.com/problems/monthly-transactions-ii/)|Monthly Transactions II|中等|
|[274](https://leetcode-cn.com/problems/h-index/)|H-Index|中等|
|[275](https://leetcode-cn.com/problems/h-index-ii/)|H-Index II|中等|
|[496](https://leetcode-cn.com/problems/next-greater-element-i/)|Next Greater Element I|简单|
|[503](https://leetcode-cn.com/problems/next-greater-element-ii/)|[Next Greater Element II](./src/0503-next-greater-element-ii/)|中等|
|[556](https://leetcode-cn.com/problems/next-greater-element-iii/)|Next Greater Element III|中等|
|[125](https://leetcode-cn.com/problems/valid-palindrome/)|[Valid Palindrome](./src/0125-valid-palindrome/)|简单|
|[680](https://leetcode-cn.com/problems/valid-palindrome-ii/)|Valid Palindrome II|简单|
|[1216](https://leetcode-cn.com/problems/valid-palindrome-iii/)|Valid Palindrome III|困难|
|[684](https://leetcode-cn.com/problems/redundant-connection/)|Redundant Connection|中等|
|[685](https://leetcode-cn.com/problems/redundant-connection-ii/)|Redundant Connection II|困难|
|[877](https://leetcode-cn.com/problems/stone-game/)|Stone Game|中等|
|[1140](https://leetcode-cn.com/problems/stone-game-ii/)|Stone Game II|中等|
|[1406](https://leetcode-cn.com/problems/stone-game-iii/)|Stone Game III|困难|
|[1510](https://leetcode-cn.com/problems/stone-game-iv/)|Stone Game IV|困难|
|[153](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)|Find Minimum in Rotated Sorted Array|中等|
|[154](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)|Find Minimum in Rotated Sorted Array II|困难|
|[200](https://leetcode-cn.com/problems/number-of-islands/)|Number of Islands|中等|
|[305](https://leetcode-cn.com/problems/number-of-islands-ii/)|Number of Islands II|困难|
|[768](https://leetcode-cn.com/problems/max-chunks-to-make-sorted-ii/)|Max Chunks To Make Sorted II|困难|
|[769](https://leetcode-cn.com/problems/max-chunks-to-make-sorted/)|Max Chunks To Make Sorted|中等|
|[1173](https://leetcode-cn.com/problems/immediate-food-delivery-i/)|Immediate Food Delivery I|简单|
|[1174](https://leetcode-cn.com/problems/immediate-food-delivery-ii/)|Immediate Food Delivery II|中等|
|[443](https://leetcode-cn.com/problems/string-compression/)|String Compression|简单|
|[1531](https://leetcode-cn.com/problems/string-compression-ii/)|String Compression II|困难|
|[198](https://leetcode-cn.com/problems/house-robber/)|[House Robber](./src/0198-house-robber/)|简单|
|[213](https://leetcode-cn.com/problems/house-robber-ii/)|House Robber II|中等|
|[337](https://leetcode-cn.com/problems/house-robber-iii/)|House Robber III|中等|
|[531](https://leetcode-cn.com/problems/lonely-pixel-i/)|Lonely Pixel I|中等|
|[533](https://leetcode-cn.com/problems/lonely-pixel-ii/)|Lonely Pixel II|中等|
|[319](https://leetcode-cn.com/problems/bulb-switcher/)|Bulb Switcher|中等|
|[672](https://leetcode-cn.com/problems/bulb-switcher-ii/)|Bulb Switcher II|中等|
|[1375](https://leetcode-cn.com/problems/bulb-switcher-iii/)|Bulb Switcher III|中等|
|[1529](https://leetcode-cn.com/problems/bulb-switcher-iv/)|Bulb Switcher IV|中等|
|[939](https://leetcode-cn.com/problems/minimum-area-rectangle/)|Minimum Area Rectangle|中等|
|[963](https://leetcode-cn.com/problems/minimum-area-rectangle-ii/)|Minimum Area Rectangle II|中等|
|[112](https://leetcode-cn.com/problems/path-sum/)|[Path Sum](./src/0112-path-sum/)|简单|
|[113](https://leetcode-cn.com/problems/path-sum-ii/)|Path Sum II|中等|
|[437](https://leetcode-cn.com/problems/path-sum-iii/)|Path Sum III|中等|
|[666](https://leetcode-cn.com/problems/path-sum-iv/)|Path Sum IV|中等|
|[169](https://leetcode-cn.com/problems/majority-element/)|Majority Element|简单|
|[229](https://leetcode-cn.com/problems/majority-element-ii/)|Majority Element II|中等|
|[344](https://leetcode-cn.com/problems/reverse-string/)|Reverse String|简单|
|[541](https://leetcode-cn.com/problems/reverse-string-ii/)|Reverse String II|简单|
|[905](https://leetcode-cn.com/problems/sort-array-by-parity/)|Sort Array By Parity|简单|
|[922](https://leetcode-cn.com/problems/sort-array-by-parity-ii/)|Sort Array By Parity II|简单|
|[1056](https://leetcode-cn.com/problems/confusing-number/)|Confusing Number|简单|
|[1088](https://leetcode-cn.com/problems/confusing-number-ii/)|Confusing Number II|困难|
|[1113](https://leetcode-cn.com/problems/reported-posts/)|Reported Posts|简单|
|[1132](https://leetcode-cn.com/problems/reported-posts-ii/)|Reported Posts II|中等|
|[39](https://leetcode-cn.com/problems/combination-sum/)|[Combination Sum](./src/0039-combination-sum/)|中等|
|[40](https://leetcode-cn.com/problems/combination-sum-ii/)|[Combination Sum II](./src/0040-combination-sum-ii/)|中等|
|[216](https://leetcode-cn.com/problems/combination-sum-iii/)|Combination Sum III|中等|
|[377](https://leetcode-cn.com/problems/combination-sum-iv/)|Combination Sum IV|中等|
|[46](https://leetcode-cn.com/problems/permutations/)|Permutations|中等|
|[47](https://leetcode-cn.com/problems/permutations-ii/)|Permutations II|中等|
|[139](https://leetcode-cn.com/problems/word-break/)|Word Break|中等|
|[140](https://leetcode-cn.com/problems/word-break-ii/)|Word Break II|困难|
|[141](https://leetcode-cn.com/problems/linked-list-cycle/)|[Linked List Cycle](./src/0141-linked-list-cycle/)|简单|
|[142](https://leetcode-cn.com/problems/linked-list-cycle-ii/)|Linked List Cycle II|中等|
|[266](https://leetcode-cn.com/problems/palindrome-permutation/)|Palindrome Permutation|简单|
|[267](https://leetcode-cn.com/problems/palindrome-permutation-ii/)|Palindrome Permutation II|中等|
|[280](https://leetcode-cn.com/problems/wiggle-sort/)|Wiggle Sort|中等|
|[324](https://leetcode-cn.com/problems/wiggle-sort-ii/)|Wiggle Sort II|中等|
|[374](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)|[Guess Number Higher or Lower](./src/0374-guess-number-higher-or-lower/)|简单|
|[375](https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/)|Guess Number Higher or Lower II|中等|
|[734](https://leetcode-cn.com/problems/sentence-similarity/)|Sentence Similarity|简单|
|[737](https://leetcode-cn.com/problems/sentence-similarity-ii/)|Sentence Similarity II|中等|
|[223](https://leetcode-cn.com/problems/rectangle-area/)|Rectangle Area|中等|
|[850](https://leetcode-cn.com/problems/rectangle-area-ii/)|Rectangle Area II|困难|
|[1141](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-i/)|User Activity for the Past 30 Days I|简单|
|[1142](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-ii/)|User Activity for the Past 30 Days II|简单|
|[131](https://leetcode-cn.com/problems/palindrome-partitioning/)|Palindrome Partitioning|中等|
|[132](https://leetcode-cn.com/problems/palindrome-partitioning-ii/)|Palindrome Partitioning II|困难|
|[1278](https://leetcode-cn.com/problems/palindrome-partitioning-iii/)|Palindrome Partitioning III|困难|
|[252](https://leetcode-cn.com/problems/meeting-rooms/)|Meeting Rooms|简单|
|[253](https://leetcode-cn.com/problems/meeting-rooms-ii/)|Meeting Rooms II|中等|
|[285](https://leetcode-cn.com/problems/inorder-successor-in-bst/)|Inorder Successor in BST|中等|
|[510](https://leetcode-cn.com/problems/inorder-successor-in-bst-ii/)|Inorder Successor in BST II|中等|
|[551](https://leetcode-cn.com/problems/student-attendance-record-i/)|Student Attendance Record I|简单|
|[552](https://leetcode-cn.com/problems/student-attendance-record-ii/)|Student Attendance Record II|困难|
|[511](https://leetcode-cn.com/problems/game-play-analysis-i/)|Game Play Analysis I|简单|
|[512](https://leetcode-cn.com/problems/game-play-analysis-ii/)|Game Play Analysis II|简单|
|[534](https://leetcode-cn.com/problems/game-play-analysis-iii/)|Game Play Analysis III|中等|
|[550](https://leetcode-cn.com/problems/game-play-analysis-iv/)|Game Play Analysis IV|中等|
|[1097](https://leetcode-cn.com/problems/game-play-analysis-v/)|Game Play Analysis V|困难|
|[151](https://leetcode-cn.com/problems/reverse-words-in-a-string/)|Reverse Words in a String|中等|
|[186](https://leetcode-cn.com/problems/reverse-words-in-a-string-ii/)|Reverse Words in a String II|中等|
|[557](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)|Reverse Words in a String III|简单|
|[349](https://leetcode-cn.com/problems/intersection-of-two-arrays/)|Intersection of Two Arrays|简单|
|[350](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)|Intersection of Two Arrays II|简单|
|[339](https://leetcode-cn.com/problems/nested-list-weight-sum/)|Nested List Weight Sum|简单|
|[364](https://leetcode-cn.com/problems/nested-list-weight-sum-ii/)|Nested List Weight Sum II|中等|
|[924](https://leetcode-cn.com/problems/minimize-malware-spread/)|Minimize Malware Spread|困难|
|[928](https://leetcode-cn.com/problems/minimize-malware-spread-ii/)|Minimize Malware Spread II|困难|
|[115](https://leetcode-cn.com/problems/distinct-subsequences/)|Distinct Subsequences|困难|
|[940](https://leetcode-cn.com/problems/distinct-subsequences-ii/)|Distinct Subsequences II|困难|
|[78](https://leetcode-cn.com/problems/subsets/)|Subsets|中等|
|[90](https://leetcode-cn.com/problems/subsets-ii/)|Subsets II|中等|
|[102](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)|Binary Tree Level Order Traversal|中等|
|[107](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)|[Binary Tree Level Order Traversal II](./src/0107-binary-tree-level-order-traversal-ii/)|简单|
|[293](https://leetcode-cn.com/problems/flip-game/)|Flip Game|简单|
|[294](https://leetcode-cn.com/problems/flip-game-ii/)|Flip Game II|中等|
|[42](https://leetcode-cn.com/problems/trapping-rain-water/)|[Trapping Rain Water](./src/0042-trapping-rain-water/)|困难|
|[407](https://leetcode-cn.com/problems/trapping-rain-water-ii/)|Trapping Rain Water II|困难|
|[694](https://leetcode-cn.com/problems/number-of-distinct-islands/)|Number of Distinct Islands|中等|
|[711](https://leetcode-cn.com/problems/number-of-distinct-islands-ii/)|Number of Distinct Islands II|困难|
|[1075](https://leetcode-cn.com/problems/project-employees-i/)|Project Employees I|简单|
|[1076](https://leetcode-cn.com/problems/project-employees-ii/)|Project Employees II|简单|
|[1077](https://leetcode-cn.com/problems/project-employees-iii/)|Project Employees III|中等|
|[1158](https://leetcode-cn.com/problems/market-analysis-i/)|Market Analysis I|中等|
|[1159](https://leetcode-cn.com/problems/market-analysis-ii/)|Market Analysis II|困难|
|[95](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/)|Unique Binary Search Trees II|中等|
|[96](https://leetcode-cn.com/problems/unique-binary-search-trees/)|Unique Binary Search Trees|中等|
|[116](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/)|Populating Next Right Pointers in Each Node|中等|
|[117](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)|Populating Next Right Pointers in Each Node II|中等|
|[207](https://leetcode-cn.com/problems/course-schedule/)|Course Schedule|中等|
|[210](https://leetcode-cn.com/problems/course-schedule-ii/)|Course Schedule II|中等|
|[630](https://leetcode-cn.com/problems/course-schedule-iii/)|Course Schedule III|困难|
|[1462](https://leetcode-cn.com/problems/course-schedule-iv/)|Course Schedule IV|中等|
|[485](https://leetcode-cn.com/problems/max-consecutive-ones/)|Max Consecutive Ones|简单|
|[487](https://leetcode-cn.com/problems/max-consecutive-ones-ii/)|Max Consecutive Ones II|中等|
|[1004](https://leetcode-cn.com/problems/max-consecutive-ones-iii/)|Max Consecutive Ones III|中等|
|[521](https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/)|Longest Uncommon Subsequence I |简单|
|[522](https://leetcode-cn.com/problems/longest-uncommon-subsequence-ii/)|Longest Uncommon Subsequence II|中等|
|[526](https://leetcode-cn.com/problems/beautiful-arrangement/)|Beautiful Arrangement|中等|
|[667](https://leetcode-cn.com/problems/beautiful-arrangement-ii/)|Beautiful Arrangement II|中等|
|[1087](https://leetcode-cn.com/problems/brace-expansion/)|Brace Expansion|中等|
|[1096](https://leetcode-cn.com/problems/brace-expansion-ii/)|Brace Expansion II|困难|
|[931](https://leetcode-cn.com/problems/minimum-falling-path-sum/)|Minimum Falling Path Sum|中等|
|[1289](https://leetcode-cn.com/problems/minimum-falling-path-sum-ii/)|Minimum Falling Path Sum II|困难|
|[126](https://leetcode-cn.com/problems/word-ladder-ii/)|Word Ladder II|困难|
|[127](https://leetcode-cn.com/problems/word-ladder/)|Word Ladder|中等|
|[217](https://leetcode-cn.com/problems/contains-duplicate/)|Contains Duplicate|简单|
|[219](https://leetcode-cn.com/problems/contains-duplicate-ii/)|Contains Duplicate II|简单|
|[220](https://leetcode-cn.com/problems/contains-duplicate-iii/)|Contains Duplicate III|中等|
|[74](https://leetcode-cn.com/problems/search-a-2d-matrix/)|Search a 2D Matrix|中等|
|[240](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)|Search a 2D Matrix II|中等|
|[256](https://leetcode-cn.com/problems/paint-house/)|Paint House|简单|
|[265](https://leetcode-cn.com/problems/paint-house-ii/)|Paint House II|困难|
|[1473](https://leetcode-cn.com/problems/paint-house-iii/)|Paint House III|困难|
|[298](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence/)|Binary Tree Longest Consecutive Sequence|中等|
|[549](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence-ii/)|Binary Tree Longest Consecutive Sequence II|中等|
|[654](https://leetcode-cn.com/problems/maximum-binary-tree/)|Maximum Binary Tree|中等|
|[998](https://leetcode-cn.com/problems/maximum-binary-tree-ii/)|Maximum Binary Tree II|中等|
|[498](https://leetcode-cn.com/problems/diagonal-traverse/)|Diagonal Traverse|中等|
|[1424](https://leetcode-cn.com/problems/diagonal-traverse-ii/)|Diagonal Traverse II|中等|
|[62](https://leetcode-cn.com/problems/unique-paths/)|Unique Paths|中等|
|[63](https://leetcode-cn.com/problems/unique-paths-ii/)|Unique Paths II|中等|
|[980](https://leetcode-cn.com/problems/unique-paths-iii/)|Unique Paths III|困难|
|[33](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)|[Search in Rotated Sorted Array](./src/0033-search-in-rotated-sorted-array/)|中等|
|[81](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)|Search in Rotated Sorted Array II|中等|
|[246](https://leetcode-cn.com/problems/strobogrammatic-number/)|Strobogrammatic Number|简单|
|[247](https://leetcode-cn.com/problems/strobogrammatic-number-ii/)|Strobogrammatic Number II|中等|
|[248](https://leetcode-cn.com/problems/strobogrammatic-number-iii/)|Strobogrammatic Number III|困难|
|[643](https://leetcode-cn.com/problems/maximum-average-subarray-i/)|Maximum Average Subarray I|简单|
|[644](https://leetcode-cn.com/problems/maximum-average-subarray-ii/)|Maximum Average Subarray II|困难|
|[1068](https://leetcode-cn.com/problems/product-sales-analysis-i/)|Product Sales Analysis I|简单|
|[1069](https://leetcode-cn.com/problems/product-sales-analysis-ii/)|Product Sales Analysis II|简单|
|[1070](https://leetcode-cn.com/problems/product-sales-analysis-iii/)|Product Sales Analysis III|中等|
|[51](https://leetcode-cn.com/problems/n-queens/)|N-Queens|困难|
|[52](https://leetcode-cn.com/problems/n-queens-ii/)|N-Queens II|困难|
|[54](https://leetcode-cn.com/problems/spiral-matrix/)|Spiral Matrix|中等|
|[59](https://leetcode-cn.com/problems/spiral-matrix-ii/)|Spiral Matrix II|中等|
|[885](https://leetcode-cn.com/problems/spiral-matrix-iii/)|Spiral Matrix III|中等|
|[224](https://leetcode-cn.com/problems/basic-calculator/)|Basic Calculator|困难|
|[227](https://leetcode-cn.com/problems/basic-calculator-ii/)|Basic Calculator II|中等|
|[770](https://leetcode-cn.com/problems/basic-calculator-iv/)|Basic Calculator IV|困难|
|[772](https://leetcode-cn.com/problems/basic-calculator-iii/)|Basic Calculator III|困难|
|[18](https://leetcode-cn.com/problems/4sum/)|[4Sum](./src/0018-4sum/)|中等|
|[454](https://leetcode-cn.com/problems/4sum-ii/)|4Sum II|中等|
|[370](https://leetcode-cn.com/problems/range-addition/)|Range Addition|中等|
|[598](https://leetcode-cn.com/problems/range-addition-ii/)|Range Addition II|简单|
|[729](https://leetcode-cn.com/problems/my-calendar-i/)|My Calendar I|中等|
|[731](https://leetcode-cn.com/problems/my-calendar-ii/)|My Calendar II|中等|
|[732](https://leetcode-cn.com/problems/my-calendar-iii/)|My Calendar III|困难|
|[944](https://leetcode-cn.com/problems/delete-columns-to-make-sorted/)|Delete Columns to Make Sorted|简单|
|[955](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-ii/)|Delete Columns to Make Sorted II|中等|
|[960](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-iii/)|Delete Columns to Make Sorted III|困难|
|[1082](https://leetcode-cn.com/problems/sales-analysis-i/)|Sales Analysis I|简单|
|[1083](https://leetcode-cn.com/problems/sales-analysis-ii/)|Sales Analysis II|简单|
|[1084](https://leetcode-cn.com/problems/sales-analysis-iii/)|Sales Analysis III|简单|
|[741](https://leetcode-cn.com/problems/cherry-pickup/)|Cherry Pickup|困难|
|[1463](https://leetcode-cn.com/problems/cherry-pickup-ii/)|Cherry Pickup II|困难|
|[118](https://leetcode-cn.com/problems/pascals-triangle/)|[Pascal&#39;s Triangle](./src/0118-pascals-triangle/)|简单|
|[119](https://leetcode-cn.com/problems/pascals-triangle-ii/)|[Pascal&#39;s Triangle II](./src/0119-pascals-triangle-ii/)|简单|
|[136](https://leetcode-cn.com/problems/single-number/)|[Single Number](./src/0136-single-number/)|简单|
|[137](https://leetcode-cn.com/problems/single-number-ii/)|Single Number II|中等|
|[260](https://leetcode-cn.com/problems/single-number-iii/)|Single Number III|中等|
|[263](https://leetcode-cn.com/problems/ugly-number/)|Ugly Number|简单|
|[264](https://leetcode-cn.com/problems/ugly-number-ii/)|Ugly Number II|中等|
|[1201](https://leetcode-cn.com/problems/ugly-number-iii/)|Ugly Number III|中等|
|[908](https://leetcode-cn.com/problems/smallest-range-i/)|Smallest Range I|简单|
|[910](https://leetcode-cn.com/problems/smallest-range-ii/)|Smallest Range II|中等|
|[1046](https://leetcode-cn.com/problems/last-stone-weight/)|Last Stone Weight|简单|
|[1049](https://leetcode-cn.com/problems/last-stone-weight-ii/)|Last Stone Weight II|中等|


---

[⬆️Top](#leetcli)
