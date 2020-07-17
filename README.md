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

[![数组](https://img.shields.io/badge/数组-301-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/array/)
[![动态规划](https://img.shields.io/badge/动态规划-228-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/dynamic-programming/)
[![字符串](https://img.shields.io/badge/字符串-204-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/string/)
[![数学](https://img.shields.io/badge/数学-199-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/math/)
[![树](https://img.shields.io/badge/树-164-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/tree/)
[![深度优先搜索](https://img.shields.io/badge/深度优先搜索-141-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/depth-first-search/)
[![哈希表](https://img.shields.io/badge/哈希表-137-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/hash-table/)
[![二分查找](https://img.shields.io/badge/二分查找-93-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/binary-search/)
[![贪心算法](https://img.shields.io/badge/贪心算法-82-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/greedy/)
[![广度优先搜索](https://img.shields.io/badge/广度优先搜索-76-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/breadth-first-search/)
[![双指针](https://img.shields.io/badge/双指针-70-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/two-pointers/)
[![栈](https://img.shields.io/badge/栈-64-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/stack/)
[![位运算](https://img.shields.io/badge/位运算-63-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/bit-manipulation/)
[![排序](https://img.shields.io/badge/排序-63-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/sort/)
[![回溯算法](https://img.shields.io/badge/回溯算法-62-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/backtracking/)
[![设计](https://img.shields.io/badge/设计-61-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/design/)
[![链表](https://img.shields.io/badge/链表-53-orange.svg?style=flat)](https://leetcode-cn.com/tag/linked-list/)
[![图](https://img.shields.io/badge/图-45-orange.svg?style=flat)](https://leetcode-cn.com/tag/graph/)
[![堆](https://img.shields.io/badge/堆-40-blue.svg?style=flat)](https://leetcode-cn.com/tag/heap/)
[![并查集](https://img.shields.io/badge/并查集-31-blue.svg?style=flat)](https://leetcode-cn.com/tag/union-find/)
[![SlidingWindow](https://img.shields.io/badge/SlidingWindow-29-blue.svg?style=flat)](https://leetcode-cn.com/tag/sliding-window/)
[![分治算法](https://img.shields.io/badge/分治算法-28-blue.svg?style=flat)](https://leetcode-cn.com/tag/divide-and-conquer/)
[![递归](https://img.shields.io/badge/递归-23-blue.svg?style=flat)](https://leetcode-cn.com/tag/recursion/)
[![字典树](https://img.shields.io/badge/字典树-18-red.svg?style=flat)](https://leetcode-cn.com/tag/trie/)
[![线段树](https://img.shields.io/badge/线段树-12-red.svg?style=flat)](https://leetcode-cn.com/tag/segment-tree/)
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

|#|标题|难度|提交次数|通过率|
|:-:|:-|:-: | :-: | :-: |
|[1](https://leetcode-cn.com/problems/two-sum/)|[Two Sum](./src/0001-two-sum/)|简单|2479611|49%|
|[2](https://leetcode-cn.com/problems/add-two-numbers/)|[Add Two Numbers](./src/0002-add-two-numbers/)|中等|1274407|37%|
|[3](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)|[Longest Substring Without Repeating Characters](./src/0003-longest-substring-without-repeating-characters/)|中等|1625831|35%|
|[4](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)|[Median of Two Sorted Arrays](./src/0004-median-of-two-sorted-arrays/)|困难|586539|38%|
|[5](https://leetcode-cn.com/problems/longest-palindromic-substring/)|[Longest Palindromic Substring](./src/0005-longest-palindromic-substring/)|中等|1025116|31%|
|[6](https://leetcode-cn.com/problems/zigzag-conversion/)|[ZigZag Conversion](./src/0006-zigzag-conversion/)|中等|306527|48%|
|[7](https://leetcode-cn.com/problems/reverse-integer/)|[Reverse Integer](./src/0007-reverse-integer/)|简单|1177565|34%|
|[8](https://leetcode-cn.com/problems/string-to-integer-atoi/)|[String to Integer (atoi)](./src/0008-string-to-integer-atoi/)|中等|880165|20%|
|[9](https://leetcode-cn.com/problems/palindrome-number/)|[Palindrome Number](./src/0009-palindrome-number/)|简单|675965|58%|
|[10](https://leetcode-cn.com/problems/regular-expression-matching/)|[Regular Expression Matching](./src/0010-regular-expression-matching/)|困难|334550|29%|
|[11](https://leetcode-cn.com/problems/container-with-most-water/)|[Container With Most Water](./src/0011-container-with-most-water/)|中等|381626|63%|
|[12](https://leetcode-cn.com/problems/integer-to-roman/)|[Integer to Roman](./src/0012-integer-to-roman/)|中等|157361|63%|
|[13](https://leetcode-cn.com/problems/roman-to-integer/)|[Roman to Integer](./src/0013-roman-to-integer/)|简单|362717|61%|
|[14](https://leetcode-cn.com/problems/longest-common-prefix/)|[Longest Common Prefix](./src/0014-longest-common-prefix/)|简单|812322|38%|
|[15](https://leetcode-cn.com/problems/3sum/)|[3Sum](./src/0015-3sum/)|中等|973530|28%|
|[16](https://leetcode-cn.com/problems/3sum-closest/)|[3Sum Closest](./src/0016-3sum-closest/)|中等|296112|45%|
|[17](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)|[Letter Combinations of a Phone Number](./src/0017-letter-combinations-of-a-phone-number/)|中等|245231|53%|
|[18](https://leetcode-cn.com/problems/4sum/)|[4Sum](./src/0018-4sum/)|中等|240890|38%|
|[19](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)|[Remove Nth Node From End of List](./src/0019-remove-nth-node-from-end-of-list/)|中等|507416|39%|
|[20](https://leetcode-cn.com/problems/valid-parentheses/)|[Valid Parentheses](./src/0020-valid-parentheses/)|简单|781788|41%|
|[21](https://leetcode-cn.com/problems/merge-two-sorted-lists/)|[Merge Two Sorted Lists](./src/0021-merge-two-sorted-lists/)|简单|498417|63%|
|[22](https://leetcode-cn.com/problems/generate-parentheses/)|[Generate Parentheses](./src/0022-generate-parentheses/)|中等|197743|75%|
|[23](https://leetcode-cn.com/problems/merge-k-sorted-lists/)|[Merge k Sorted Lists](./src/0023-merge-k-sorted-lists/)|困难|275214|52%|
|[24](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)|[Swap Nodes in Pairs](./src/0024-swap-nodes-in-pairs/)|中等|189429|66%|
|[25](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)|[Reverse Nodes in k-Group](./src/0025-reverse-nodes-in-k-group/)|困难|134797|61%|
|[26](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)|[Remove Duplicates from Sorted Array](./src/0026-remove-duplicates-from-sorted-array/)|简单|718248|50%|
|[27](https://leetcode-cn.com/problems/remove-element/)|[Remove Element](./src/0027-remove-element/)|简单|346907|58%|
|[28](https://leetcode-cn.com/problems/implement-strstr/)|[Implement strStr()](./src/0028-implement-strstr/)|简单|499079|39%|
|[29](https://leetcode-cn.com/problems/divide-two-integers/)|[Divide Two Integers](./src/0029-divide-two-integers/)|中等|272736|19%|
|[30](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)|[Substring with Concatenation of All Words](./src/0030-substring-with-concatenation-of-all-words/)|困难|122598|31%|
|[31](https://leetcode-cn.com/problems/next-permutation/)|[Next Permutation](./src/0031-next-permutation/)|中等|217388|34%|
|[32](https://leetcode-cn.com/problems/longest-valid-parentheses/)|[Longest Valid Parentheses](./src/0032-longest-valid-parentheses/)|困难|261489|33%|
|[33](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)|[Search in Rotated Sorted Array](./src/0033-search-in-rotated-sorted-array/)|中等|378281|38%|
|[34](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)|[Find First and Last Position of Element in Sorted Array](./src/0034-find-first-and-last-position-of-element-in-sorted-array/)|中等|274358|39%|
|[35](https://leetcode-cn.com/problems/search-insert-position/)|[Search Insert Position](./src/0035-search-insert-position/)|简单|435634|46%|
|[36](https://leetcode-cn.com/problems/valid-sudoku/)|[Valid Sudoku](./src/0036-valid-sudoku/)|中等|137052|60%|
|[37](https://leetcode-cn.com/problems/sudoku-solver/)|[Sudoku Solver](./src/0037-sudoku-solver/)|困难|54594|62%|
|[38](https://leetcode-cn.com/problems/count-and-say/)|[Count and Say](./src/0038-count-and-say/)|简单|198779|55%|
|[39](https://leetcode-cn.com/problems/combination-sum/)|[Combination Sum](./src/0039-combination-sum/)|中等|159938|69%|
|[40](https://leetcode-cn.com/problems/combination-sum-ii/)|[Combination Sum II](./src/0040-combination-sum-ii/)|中等|109186|62%|
|[41](https://leetcode-cn.com/problems/first-missing-positive/)|[First Missing Positive](./src/0041-first-missing-positive/)|困难|190998|40%|
|[42](https://leetcode-cn.com/problems/trapping-rain-water/)|[Trapping Rain Water](./src/0042-trapping-rain-water/)|困难|237098|51%|
|[43](https://leetcode-cn.com/problems/multiply-strings/)|[Multiply Strings](./src/0043-multiply-strings/)|中等|172802|42%|
|[44](https://leetcode-cn.com/problems/wildcard-matching/)|Wildcard Matching|困难|152139|31%|
|[45](https://leetcode-cn.com/problems/jump-game-ii/)|Jump Game II|困难|190948|36%|
|[46](https://leetcode-cn.com/problems/permutations/)|Permutations|中等|201379|76%|
|[47](https://leetcode-cn.com/problems/permutations-ii/)|Permutations II|中等|120200|59%|
|[48](https://leetcode-cn.com/problems/rotate-image/)|Rotate Image|中等|121777|68%|
|[49](https://leetcode-cn.com/problems/group-anagrams/)|Group Anagrams|中等|140917|62%|
|[50](https://leetcode-cn.com/problems/powx-n/)|Pow(x, n)|中等|300157|36%|
|[51](https://leetcode-cn.com/problems/n-queens/)|N-Queens|困难|71018|70%|
|[52](https://leetcode-cn.com/problems/n-queens-ii/)|N-Queens II|困难|34425|78%|
|[53](https://leetcode-cn.com/problems/maximum-subarray/)|[Maximum Subarray](./src/0053-maximum-subarray/)|简单|539985|51%|
|[54](https://leetcode-cn.com/problems/spiral-matrix/)|Spiral Matrix|中等|169341|40%|
|[55](https://leetcode-cn.com/problems/jump-game/)|Jump Game|中等|320823|40%|
|[56](https://leetcode-cn.com/problems/merge-intervals/)|Merge Intervals|中等|274478|42%|
|[57](https://leetcode-cn.com/problems/insert-interval/)|[Insert Interval](./src/0057-insert-interval/)|困难|68965|37%|
|[58](https://leetcode-cn.com/problems/length-of-last-word/)|[Length of Last Word](./src/0058-length-of-last-word/)|简单|316400|33%|
|[59](https://leetcode-cn.com/problems/spiral-matrix-ii/)|Spiral Matrix II|中等|51199|77%|
|[60](https://leetcode-cn.com/problems/permutation-sequence/)|Permutation Sequence|中等|81070|48%|
|[61](https://leetcode-cn.com/problems/rotate-list/)|Rotate List|中等|178417|40%|
|[62](https://leetcode-cn.com/problems/unique-paths/)|Unique Paths|中等|199900|61%|
|[63](https://leetcode-cn.com/problems/unique-paths-ii/)|Unique Paths II|中等|243046|36%|
|[64](https://leetcode-cn.com/problems/minimum-path-sum/)|Minimum Path Sum|中等|152648|66%|
|[65](https://leetcode-cn.com/problems/valid-number/)|Valid Number|困难|79143|20%|
|[66](https://leetcode-cn.com/problems/plus-one/)|[Plus One](./src/0066-plus-one/)|简单|393367|44%|
|[67](https://leetcode-cn.com/problems/add-binary/)|[Add Binary](./src/0067-add-binary/)|简单|209116|54%|
|[68](https://leetcode-cn.com/problems/text-justification/)|Text Justification|困难|22372|43%|
|[69](https://leetcode-cn.com/problems/sqrtx/)|[Sqrt(x)](./src/0069-sqrtx/)|简单|452785|38%|
|[70](https://leetcode-cn.com/problems/climbing-stairs/)|[Climbing Stairs](./src/0070-climbing-stairs/)|简单|484834|50%|
|[71](https://leetcode-cn.com/problems/simplify-path/)|Simplify Path|中等|117192|40%|
|[72](https://leetcode-cn.com/problems/edit-distance/)|Edit Distance|困难|116490|59%|
|[73](https://leetcode-cn.com/problems/set-matrix-zeroes/)|Set Matrix Zeroes|中等|82746|55%|
|[74](https://leetcode-cn.com/problems/search-a-2d-matrix/)|Search a 2D Matrix|中等|134435|38%|
|[75](https://leetcode-cn.com/problems/sort-colors/)|Sort Colors|中等|175361|55%|
|[76](https://leetcode-cn.com/problems/minimum-window-substring/)|Minimum Window Substring|困难|165708|38%|
|[77](https://leetcode-cn.com/problems/combinations/)|Combinations|中等|82509|74%|
|[78](https://leetcode-cn.com/problems/subsets/)|Subsets|中等|138181|77%|
|[79](https://leetcode-cn.com/problems/word-search/)|Word Search|中等|166696|42%|
|[80](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/)|Remove Duplicates from Sorted Array II|中等|90293|55%|
|[81](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)|Search in Rotated Sorted Array II|中等|92084|35%|
|[82](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)|[Remove Duplicates from Sorted List II](./src/0082-remove-duplicates-from-sorted-list-ii/)|中等|113014|48%|
|[83](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)|[Remove Duplicates from Sorted List](./src/0083-remove-duplicates-from-sorted-list/)|简单|233374|50%|
|[84](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)|Largest Rectangle in Histogram|困难|174867|41%|
|[85](https://leetcode-cn.com/problems/maximal-rectangle/)|Maximal Rectangle|困难|73240|47%|
|[86](https://leetcode-cn.com/problems/partition-list/)|Partition List|中等|74480|58%|
|[87](https://leetcode-cn.com/problems/scramble-string/)|Scramble String|困难|24451|46%|
|[88](https://leetcode-cn.com/problems/merge-sorted-array/)|[Merge Sorted Array](./src/0088-merge-sorted-array/)|简单|366794|48%|
|[89](https://leetcode-cn.com/problems/gray-code/)|Gray Code|中等|42351|68%|
|[90](https://leetcode-cn.com/problems/subsets-ii/)|Subsets II|中等|65580|60%|
|[91](https://leetcode-cn.com/problems/decode-ways/)|Decode Ways|中等|234355|23%|
|[92](https://leetcode-cn.com/problems/reverse-linked-list-ii/)|Reverse Linked List II|中等|120296|50%|
|[93](https://leetcode-cn.com/problems/restore-ip-addresses/)|Restore IP Addresses|中等|110988|47%|
|[94](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)|Binary Tree Inorder Traversal|中等|267708|72%|
|[95](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/)|Unique Binary Search Trees II|中等|58338|63%|
|[96](https://leetcode-cn.com/problems/unique-binary-search-trees/)|Unique Binary Search Trees|中等|105234|69%|
|[97](https://leetcode-cn.com/problems/interleaving-string/)|Interleaving String|困难|39760|40%|
|[98](https://leetcode-cn.com/problems/validate-binary-search-tree/)|Validate Binary Search Tree|中等|447385|31%|
|[99](https://leetcode-cn.com/problems/recover-binary-search-tree/)|Recover Binary Search Tree|困难|36924|57%|
|[100](https://leetcode-cn.com/problems/same-tree/)|[Same Tree](./src/0100-same-tree/)|简单|175226|58%|
|[101](https://leetcode-cn.com/problems/symmetric-tree/)|[Symmetric Tree](./src/0101-symmetric-tree/)|简单|335865|52%|
|[102](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)|Binary Tree Level Order Traversal|中等|255890|63%|
|[103](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)|Binary Tree Zigzag Level Order Traversal|中等|106762|54%|
|[104](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)|[Maximum Depth of Binary Tree](./src/0104-maximum-depth-of-binary-tree/)|简单|282576|73%|
|[105](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)|Construct Binary Tree from Preorder and Inorder Traversal|中等|141104|67%|
|[106](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)|Construct Binary Tree from Inorder and Postorder Traversal|中等|62035|68%|
|[107](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)|[Binary Tree Level Order Traversal II](./src/0107-binary-tree-level-order-traversal-ii/)|简单|104235|66%|
|[108](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)|[Convert Sorted Array to Binary Search Tree](./src/0108-convert-sorted-array-to-binary-search-tree/)|简单|133957|73%|
|[109](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/)|Convert Sorted List to Binary Search Tree|中等|46399|72%|
|[110](https://leetcode-cn.com/problems/balanced-binary-tree/)|[Balanced Binary Tree](./src/0110-balanced-binary-tree/)|简单|172613|52%|
|[111](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)|[Minimum Depth of Binary Tree](./src/0111-minimum-depth-of-binary-tree/)|简单|213240|42%|
|[112](https://leetcode-cn.com/problems/path-sum/)|[Path Sum](./src/0112-path-sum/)|简单|227686|50%|
|[113](https://leetcode-cn.com/problems/path-sum-ii/)|Path Sum II|中等|98675|60%|
|[114](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)|Flatten Binary Tree to Linked List|中等|73001|69%|
|[115](https://leetcode-cn.com/problems/distinct-subsequences/)|Distinct Subsequences|困难|32221|48%|
|[116](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/)|Populating Next Right Pointers in Each Node|中等|65166|61%|
|[117](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)|Populating Next Right Pointers in Each Node II|中等|49372|50%|
|[118](https://leetcode-cn.com/problems/pascals-triangle/)|[Pascal&#39;s Triangle](./src/0118-pascals-triangle/)|简单|135560|66%|
|[119](https://leetcode-cn.com/problems/pascals-triangle-ii/)|[Pascal&#39;s Triangle II](./src/0119-pascals-triangle-ii/)|简单|97619|61%|
|[120](https://leetcode-cn.com/problems/triangle/)|Triangle|中等|139478|66%|
|[121](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|[Best Time to Buy and Sell Stock](./src/0121-best-time-to-buy-and-sell-stock/)|简单|446658|54%|
|[122](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|[Best Time to Buy and Sell Stock II](./src/0122-best-time-to-buy-and-sell-stock-ii/)|简单|301581|61%|
|[123](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|Best Time to Buy and Sell Stock III|困难|101576|43%|
|[124](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/)|Binary Tree Maximum Path Sum|困难|143464|42%|
|[125](https://leetcode-cn.com/problems/valid-palindrome/)|[Valid Palindrome](./src/0125-valid-palindrome/)|简单|309250|45%|
|[126](https://leetcode-cn.com/problems/word-ladder-ii/)|Word Ladder II|困难|54505|38%|
|[127](https://leetcode-cn.com/problems/word-ladder/)|Word Ladder|中等|116018|42%|
|[128](https://leetcode-cn.com/problems/longest-consecutive-sequence/)|Longest Consecutive Sequence|困难|130130|51%|
|[129](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/)|Sum Root to Leaf Numbers|中等|50461|64%|
|[130](https://leetcode-cn.com/problems/surrounded-regions/)|Surrounded Regions|中等|105157|40%|
|[131](https://leetcode-cn.com/problems/palindrome-partitioning/)|Palindrome Partitioning|中等|54990|68%|
|[132](https://leetcode-cn.com/problems/palindrome-partitioning-ii/)|Palindrome Partitioning II|困难|29722|43%|
|[133](https://leetcode-cn.com/problems/clone-graph/)|Clone Graph|中等|41881|57%|
|[134](https://leetcode-cn.com/problems/gas-station/)|Gas Station|中等|73884|53%|
|[135](https://leetcode-cn.com/problems/candy/)|Candy|困难|55776|44%|
|[136](https://leetcode-cn.com/problems/single-number/)|[Single Number](./src/0136-single-number/)|简单|343709|69%|
|[137](https://leetcode-cn.com/problems/single-number-ii/)|Single Number II|中等|52532|67%|
|[138](https://leetcode-cn.com/problems/copy-list-with-random-pointer/)|Copy List with Random Pointer|中等|69438|54%|
|[139](https://leetcode-cn.com/problems/word-break/)|Word Break|中等|164460|46%|
|[140](https://leetcode-cn.com/problems/word-break-ii/)|Word Break II|困难|47401|38%|
|[141](https://leetcode-cn.com/problems/linked-list-cycle/)|[Linked List Cycle](./src/0141-linked-list-cycle/)|简单|388175|48%|
|[142](https://leetcode-cn.com/problems/linked-list-cycle-ii/)|Linked List Cycle II|中等|186379|50%|
|[143](https://leetcode-cn.com/problems/reorder-list/)|Reorder List|中等|54848|56%|
|[144](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)|Binary Tree Preorder Traversal|中等|200828|66%|
|[145](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)|Binary Tree Postorder Traversal|困难|133823|71%|
|[146](https://leetcode-cn.com/problems/lru-cache/)|[LRU Cache](./src/0146-lru-cache/)|中等|161969|49%|
|[147](https://leetcode-cn.com/problems/insertion-sort-list/)|Insertion Sort List|中等|53684|64%|
|[148](https://leetcode-cn.com/problems/sort-list/)|Sort List|中等|112735|66%|
|[149](https://leetcode-cn.com/problems/max-points-on-a-line/)|Max Points on a Line|困难|61261|22%|
|[150](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)|Evaluate Reverse Polish Notation|中等|90633|50%|
|[151](https://leetcode-cn.com/problems/reverse-words-in-a-string/)|Reverse Words in a String|中等|189842|42%|
|[152](https://leetcode-cn.com/problems/maximum-product-subarray/)|Maximum Product Subarray|中等|200686|39%|
|[153](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)|Find Minimum in Rotated Sorted Array|中等|118582|51%|
|[154](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)|Find Minimum in Rotated Sorted Array II|困难|56006|48%|
|[155](https://leetcode-cn.com/problems/min-stack/)|[Min Stack](./src/0155-min-stack/)|简单|255301|54%|
|[156](https://leetcode-cn.com/problems/binary-tree-upside-down/)|Binary Tree Upside Down|中等|3671|73%|
|[157](https://leetcode-cn.com/problems/read-n-characters-given-read4/)|Read N Characters Given Read4|简单|4179|51%|
|[158](https://leetcode-cn.com/problems/read-n-characters-given-read4-ii-call-multiple-times/)|Read N Characters Given Read4 II - Call multiple times|困难|1975|54%|
|[159](https://leetcode-cn.com/problems/longest-substring-with-at-most-two-distinct-characters/)|Longest Substring with At Most Two Distinct Characters|中等|10790|51%|
|[160](https://leetcode-cn.com/problems/intersection-of-two-linked-lists/)|Intersection of Two Linked Lists|简单|227833|55%|
|[161](https://leetcode-cn.com/problems/one-edit-distance/)|One Edit Distance|中等|10514|32%|
|[162](https://leetcode-cn.com/problems/find-peak-element/)|Find Peak Element|中等|104735|46%|
|[163](https://leetcode-cn.com/problems/missing-ranges/)|Missing Ranges|中等|15023|25%|
|[164](https://leetcode-cn.com/problems/maximum-gap/)|Maximum Gap|困难|29619|55%|
|[165](https://leetcode-cn.com/problems/compare-version-numbers/)|Compare Version Numbers|中等|48783|42%|
|[166](https://leetcode-cn.com/problems/fraction-to-recurring-decimal/)|Fraction to Recurring Decimal|中等|48774|26%|
|[167](https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/)|Two Sum II - Input array is sorted|简单|197674|54%|
|[168](https://leetcode-cn.com/problems/excel-sheet-column-title/)|Excel Sheet Column Title|简单|78796|38%|
|[169](https://leetcode-cn.com/problems/majority-element/)|Majority Element|简单|293555|63%|
|[170](https://leetcode-cn.com/problems/two-sum-iii-data-structure-design/)|Two Sum III - Data structure design|简单|8230|40%|
|[171](https://leetcode-cn.com/problems/excel-sheet-column-number/)|Excel Sheet Column Number|简单|65759|67%|
|[172](https://leetcode-cn.com/problems/factorial-trailing-zeroes/)|Factorial Trailing Zeroes|简单|103806|39%|
|[173](https://leetcode-cn.com/problems/binary-search-tree-iterator/)|Binary Search Tree Iterator|中等|32669|73%|
|[174](https://leetcode-cn.com/problems/dungeon-game/)|Dungeon Game|困难|49714|47%|
|[175](https://leetcode-cn.com/problems/combine-two-tables/)|[Combine Two Tables](./src/0175-combine-two-tables/)|简单|203990|73%|
|[176](https://leetcode-cn.com/problems/second-highest-salary/)|[Second Highest Salary](./src/0176-second-highest-salary/)|简单|329666|35%|
|[177](https://leetcode-cn.com/problems/nth-highest-salary/)|[Nth Highest Salary](./src/0177-nth-highest-salary/)|中等|119216|45%|
|[178](https://leetcode-cn.com/problems/rank-scores/)|[Rank Scores](./src/0178-rank-scores/)|中等|88613|58%|
|[179](https://leetcode-cn.com/problems/largest-number/)|Largest Number|中等|96056|36%|
|[180](https://leetcode-cn.com/problems/consecutive-numbers/)|[Consecutive Numbers](./src/0180-consecutive-numbers/)|中等|76366|49%|
|[181](https://leetcode-cn.com/problems/employees-earning-more-than-their-managers/)|[Employees Earning More Than Their Managers](./src/0181-employees-earning-more-than-their-managers/)|简单|109337|69%|
|[182](https://leetcode-cn.com/problems/duplicate-emails/)|[Duplicate Emails](./src/0182-duplicate-emails/)|简单|114024|79%|
|[183](https://leetcode-cn.com/problems/customers-who-never-order/)|[Customers Who Never Order](./src/0183-customers-who-never-order/)|简单|106353|67%|
|[184](https://leetcode-cn.com/problems/department-highest-salary/)|[Department Highest Salary](./src/0184-department-highest-salary/)|中等|96661|45%|
|[185](https://leetcode-cn.com/problems/department-top-three-salaries/)|[Department Top Three Salaries](./src/0185-department-top-three-salaries/)|困难|66986|44%|
|[186](https://leetcode-cn.com/problems/reverse-words-in-a-string-ii/)|Reverse Words in a String II|中等|5054|73%|
|[187](https://leetcode-cn.com/problems/repeated-dna-sequences/)|Repeated DNA Sequences|中等|43568|44%|
|[188](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)|Best Time to Buy and Sell Stock IV|困难|82054|29%|
|[189](https://leetcode-cn.com/problems/rotate-array/)|Rotate Array|简单|336636|42%|
|[190](https://leetcode-cn.com/problems/reverse-bits/)|Reverse Bits|简单|78371|59%|
|[191](https://leetcode-cn.com/problems/number-of-1-bits/)|Number of 1 Bits|简单|104642|68%|
|[192](https://leetcode-cn.com/problems/word-frequency/)|[Word Frequency](./src/0192-word-frequency/)|中等|34371|34%|
|[193](https://leetcode-cn.com/problems/valid-phone-numbers/)|[Valid Phone Numbers](./src/0193-valid-phone-numbers/)|简单|46404|29%|
|[194](https://leetcode-cn.com/problems/transpose-file/)|[Transpose File](./src/0194-transpose-file/)|中等|16756|33%|
|[195](https://leetcode-cn.com/problems/tenth-line/)|[Tenth Line](./src/0195-tenth-line/)|简单|45741|43%|
|[196](https://leetcode-cn.com/problems/delete-duplicate-emails/)|[Delete Duplicate Emails](./src/0196-delete-duplicate-emails/)|简单|73715|63%|
|[197](https://leetcode-cn.com/problems/rising-temperature/)|[Rising Temperature](./src/0197-rising-temperature/)|简单|83548|51%|
|[198](https://leetcode-cn.com/problems/house-robber/)|[House Robber](./src/0198-house-robber/)|简单|332128|45%|
|[199](https://leetcode-cn.com/problems/binary-tree-right-side-view/)|Binary Tree Right Side View|中等|86835|63%|
|[200](https://leetcode-cn.com/problems/number-of-islands/)|Number of Islands|中等|259154|49%|
|[201](https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/)|Bitwise AND of Numbers Range|中等|30685|46%|
|[202](https://leetcode-cn.com/problems/happy-number/)|Happy Number|简单|144955|60%|
|[203](https://leetcode-cn.com/problems/remove-linked-list-elements/)|Remove Linked List Elements|简单|194161|45%|
|[204](https://leetcode-cn.com/problems/count-primes/)|Count Primes|简单|190966|34%|
|[205](https://leetcode-cn.com/problems/isomorphic-strings/)|Isomorphic Strings|简单|93096|47%|
|[206](https://leetcode-cn.com/problems/reverse-linked-list/)|Reverse Linked List|简单|403539|69%|
|[207](https://leetcode-cn.com/problems/course-schedule/)|Course Schedule|中等|88524|51%|
|[208](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)|Implement Trie (Prefix Tree)|中等|65841|67%|
|[209](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)|Minimum Size Subarray Sum|中等|169704|44%|
|[210](https://leetcode-cn.com/problems/course-schedule-ii/)|Course Schedule II|中等|78640|50%|
|[211](https://leetcode-cn.com/problems/add-and-search-word-data-structure-design/)|Add and Search Word - Data structure design|中等|28537|45%|
|[212](https://leetcode-cn.com/problems/word-search-ii/)|Word Search II|困难|40892|41%|
|[213](https://leetcode-cn.com/problems/house-robber-ii/)|House Robber II|中等|111232|38%|
|[214](https://leetcode-cn.com/problems/shortest-palindrome/)|Shortest Palindrome|困难|30076|31%|
|[215](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)|Kth Largest Element in an Array|中等|263499|64%|
|[216](https://leetcode-cn.com/problems/combination-sum-iii/)|Combination Sum III|中等|34274|71%|
|[217](https://leetcode-cn.com/problems/contains-duplicate/)|Contains Duplicate|简单|284753|52%|
|[218](https://leetcode-cn.com/problems/the-skyline-problem/)|The Skyline Problem|困难|19073|42%|
|[219](https://leetcode-cn.com/problems/contains-duplicate-ii/)|Contains Duplicate II|简单|130516|39%|
|[220](https://leetcode-cn.com/problems/contains-duplicate-iii/)|Contains Duplicate III|中等|77460|26%|
|[221](https://leetcode-cn.com/problems/maximal-square/)|Maximal Square|中等|143984|42%|
|[222](https://leetcode-cn.com/problems/count-complete-tree-nodes/)|Count Complete Tree Nodes|中等|36658|70%|
|[223](https://leetcode-cn.com/problems/rectangle-area/)|Rectangle Area|中等|23853|43%|
|[224](https://leetcode-cn.com/problems/basic-calculator/)|Basic Calculator|困难|42510|38%|
|[225](https://leetcode-cn.com/problems/implement-stack-using-queues/)|Implement Stack using Queues|简单|94853|65%|
|[226](https://leetcode-cn.com/problems/invert-binary-tree/)|Invert Binary Tree|简单|127012|75%|
|[227](https://leetcode-cn.com/problems/basic-calculator-ii/)|Basic Calculator II|中等|54106|36%|
|[228](https://leetcode-cn.com/problems/summary-ranges/)|Summary Ranges|中等|21023|53%|
|[229](https://leetcode-cn.com/problems/majority-element-ii/)|Majority Element II|中等|41133|43%|
|[230](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)|Kth Smallest Element in a BST|中等|84154|71%|
|[231](https://leetcode-cn.com/problems/power-of-two/)|Power of Two|简单|140362|48%|
|[232](https://leetcode-cn.com/problems/implement-queue-using-stacks/)|Implement Queue using Stacks|简单|82436|64%|
|[233](https://leetcode-cn.com/problems/number-of-digit-one/)|Number of Digit One|困难|24627|36%|
|[234](https://leetcode-cn.com/problems/palindrome-linked-list/)|Palindrome Linked List|简单|247534|42%|
|[235](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)|Lowest Common Ancestor of a Binary Search Tree|简单|101767|64%|
|[236](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)|Lowest Common Ancestor of a Binary Tree|中等|155584|64%|
|[237](https://leetcode-cn.com/problems/delete-node-in-a-linked-list/)|Delete Node in a Linked List|简单|135276|82%|
|[238](https://leetcode-cn.com/problems/product-of-array-except-self/)|Product of Array Except Self|中等|99286|70%|
|[239](https://leetcode-cn.com/problems/sliding-window-maximum/)|Sliding Window Maximum|困难|126114|48%|
|[240](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)|Search a 2D Matrix II|中等|161083|41%|
|[241](https://leetcode-cn.com/problems/different-ways-to-add-parentheses/)|Different Ways to Add Parentheses|中等|17439|71%|
|[242](https://leetcode-cn.com/problems/valid-anagram/)|Valid Anagram|简单|192658|60%|
|[243](https://leetcode-cn.com/problems/shortest-word-distance/)|Shortest Word Distance|简单|6108|64%|
|[244](https://leetcode-cn.com/problems/shortest-word-distance-ii/)|Shortest Word Distance II|中等|4256|51%|
|[245](https://leetcode-cn.com/problems/shortest-word-distance-iii/)|Shortest Word Distance III|中等|3334|60%|
|[246](https://leetcode-cn.com/problems/strobogrammatic-number/)|Strobogrammatic Number|简单|5478|46%|
|[247](https://leetcode-cn.com/problems/strobogrammatic-number-ii/)|Strobogrammatic Number II|中等|5171|50%|
|[248](https://leetcode-cn.com/problems/strobogrammatic-number-iii/)|Strobogrammatic Number III|困难|2614|42%|
|[249](https://leetcode-cn.com/problems/group-shifted-strings/)|Group Shifted Strings|中等|3854|61%|
|[250](https://leetcode-cn.com/problems/count-univalue-subtrees/)|Count Univalue Subtrees|中等|3354|63%|
|[251](https://leetcode-cn.com/problems/flatten-2d-vector/)|Flatten 2D Vector|中等|2718|53%|
|[252](https://leetcode-cn.com/problems/meeting-rooms/)|Meeting Rooms|简单|8640|53%|
|[253](https://leetcode-cn.com/problems/meeting-rooms-ii/)|Meeting Rooms II|中等|25578|44%|
|[254](https://leetcode-cn.com/problems/factor-combinations/)|Factor Combinations|中等|3440|56%|
|[255](https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree/)|Verify Preorder Sequence in Binary Search Tree|中等|4967|46%|
|[256](https://leetcode-cn.com/problems/paint-house/)|Paint House|简单|8291|57%|
|[257](https://leetcode-cn.com/problems/binary-tree-paths/)|Binary Tree Paths|简单|69231|64%|
|[258](https://leetcode-cn.com/problems/add-digits/)|Add Digits|简单|64590|67%|
|[259](https://leetcode-cn.com/problems/3sum-smaller/)|3Sum Smaller|中等|5798|55%|
|[260](https://leetcode-cn.com/problems/single-number-iii/)|Single Number III|中等|34693|73%|
|[261](https://leetcode-cn.com/problems/graph-valid-tree/)|Graph Valid Tree|中等|7422|46%|
|[262](https://leetcode-cn.com/problems/trips-and-users/)|[Trips and Users](./src/0262-trips-and-users/)|困难|35663|46%|
|[263](https://leetcode-cn.com/problems/ugly-number/)|Ugly Number|简单|74188|49%|
|[264](https://leetcode-cn.com/problems/ugly-number-ii/)|Ugly Number II|中等|55221|53%|
|[265](https://leetcode-cn.com/problems/paint-house-ii/)|Paint House II|困难|4053|53%|
|[266](https://leetcode-cn.com/problems/palindrome-permutation/)|Palindrome Permutation|简单|4508|64%|
|[267](https://leetcode-cn.com/problems/palindrome-permutation-ii/)|Palindrome Permutation II|中等|3488|41%|
|[268](https://leetcode-cn.com/problems/missing-number/)|Missing Number|简单|132000|56%|
|[269](https://leetcode-cn.com/problems/alien-dictionary/)|Alien Dictionary|困难|8286|33%|
|[270](https://leetcode-cn.com/problems/closest-binary-search-tree-value/)|Closest Binary Search Tree Value|简单|7966|51%|
|[271](https://leetcode-cn.com/problems/encode-and-decode-strings/)|Encode and Decode Strings|中等|2405|53%|
|[272](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii/)|Closest Binary Search Tree Value II|困难|2420|62%|
|[273](https://leetcode-cn.com/problems/integer-to-english-words/)|Integer to English Words|困难|22012|27%|
|[274](https://leetcode-cn.com/problems/h-index/)|H-Index|中等|34674|38%|
|[275](https://leetcode-cn.com/problems/h-index-ii/)|H-Index II|中等|24633|41%|
|[276](https://leetcode-cn.com/problems/paint-fence/)|Paint Fence|简单|6967|44%|
|[277](https://leetcode-cn.com/problems/find-the-celebrity/)|Find the Celebrity|中等|5849|56%|
|[278](https://leetcode-cn.com/problems/first-bad-version/)|[First Bad Version](./src/0278-first-bad-version/)|简单|138591|39%|
|[279](https://leetcode-cn.com/problems/perfect-squares/)|Perfect Squares|中等|125452|56%|
|[280](https://leetcode-cn.com/problems/wiggle-sort/)|Wiggle Sort|中等|4117|67%|
|[281](https://leetcode-cn.com/problems/zigzag-iterator/)|Zigzag Iterator|中等|1480|74%|
|[282](https://leetcode-cn.com/problems/expression-add-operators/)|Expression Add Operators|困难|9930|33%|
|[283](https://leetcode-cn.com/problems/move-zeroes/)|Move Zeroes|简单|284444|61%|
|[284](https://leetcode-cn.com/problems/peeking-iterator/)|Peeking Iterator|中等|6647|70%|
|[285](https://leetcode-cn.com/problems/inorder-successor-in-bst/)|Inorder Successor in BST|中等|5439|61%|
|[286](https://leetcode-cn.com/problems/walls-and-gates/)|Walls and Gates|中等|12408|46%|
|[287](https://leetcode-cn.com/problems/find-the-duplicate-number/)|Find the Duplicate Number|中等|129349|65%|
|[288](https://leetcode-cn.com/problems/unique-word-abbreviation/)|Unique Word Abbreviation|中等|4593|32%|
|[289](https://leetcode-cn.com/problems/game-of-life/)|Game of Life|中等|48289|74%|
|[290](https://leetcode-cn.com/problems/word-pattern/)|Word Pattern|简单|63050|43%|
|[291](https://leetcode-cn.com/problems/word-pattern-ii/)|Word Pattern II|困难|1943|52%|
|[292](https://leetcode-cn.com/problems/nim-game/)|[Nim Game](./src/0292-nim-game/)|简单|74467|69%|
|[293](https://leetcode-cn.com/problems/flip-game/)|Flip Game|简单|3643|70%|
|[294](https://leetcode-cn.com/problems/flip-game-ii/)|Flip Game II|中等|2981|57%|
|[295](https://leetcode-cn.com/problems/find-median-from-data-stream/)|Find Median from Data Stream|困难|39856|46%|
|[296](https://leetcode-cn.com/problems/best-meeting-point/)|Best Meeting Point|困难|1776|59%|
|[297](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)|Serialize and Deserialize Binary Tree|困难|81903|51%|
|[298](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence/)|Binary Tree Longest Consecutive Sequence|中等|3817|56%|
|[299](https://leetcode-cn.com/problems/bulls-and-cows/)|Bulls and Cows|简单|33283|47%|
|[300](https://leetcode-cn.com/problems/longest-increasing-subsequence/)|Longest Increasing Subsequence|中等|261250|44%|
|[301](https://leetcode-cn.com/problems/remove-invalid-parentheses/)|Remove Invalid Parentheses|困难|25372|47%|
|[302](https://leetcode-cn.com/problems/smallest-rectangle-enclosing-black-pixels/)|Smallest Rectangle Enclosing Black Pixels|困难|1331|69%|
|[303](https://leetcode-cn.com/problems/range-sum-query-immutable/)|[Range Sum Query - Immutable](./src/0303-range-sum-query-immutable/)|简单|73208|62%|
|[304](https://leetcode-cn.com/problems/range-sum-query-2d-immutable/)|Range Sum Query 2D - Immutable|中等|22132|44%|
|[305](https://leetcode-cn.com/problems/number-of-islands-ii/)|Number of Islands II|困难|4896|33%|
|[306](https://leetcode-cn.com/problems/additive-number/)|Additive Number|中等|22266|32%|
|[307](https://leetcode-cn.com/problems/range-sum-query-mutable/)|Range Sum Query - Mutable|中等|21135|56%|
|[308](https://leetcode-cn.com/problems/range-sum-query-2d-mutable/)|Range Sum Query 2D - Mutable|困难|1763|58%|
|[309](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)|Best Time to Buy and Sell Stock with Cooldown|中等|84684|56%|
|[310](https://leetcode-cn.com/problems/minimum-height-trees/)|Minimum Height Trees|中等|27374|34%|
|[311](https://leetcode-cn.com/problems/sparse-matrix-multiplication/)|Sparse Matrix Multiplication|中等|1747|74%|
|[312](https://leetcode-cn.com/problems/burst-balloons/)|Burst Balloons|困难|24614|61%|
|[313](https://leetcode-cn.com/problems/super-ugly-number/)|Super Ugly Number|中等|14294|63%|
|[314](https://leetcode-cn.com/problems/binary-tree-vertical-order-traversal/)|Binary Tree Vertical Order Traversal|中等|3245|54%|
|[315](https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/)|Count of Smaller Numbers After Self|困难|77176|41%|
|[316](https://leetcode-cn.com/problems/remove-duplicate-letters/)|Remove Duplicate Letters|困难|38563|39%|
|[317](https://leetcode-cn.com/problems/shortest-distance-from-all-buildings/)|Shortest Distance from All Buildings|困难|4006|47%|
|[318](https://leetcode-cn.com/problems/maximum-product-of-word-lengths/)|Maximum Product of Word Lengths|中等|13275|64%|
|[319](https://leetcode-cn.com/problems/bulb-switcher/)|Bulb Switcher|中等|25735|46%|
|[320](https://leetcode-cn.com/problems/generalized-abbreviation/)|Generalized Abbreviation|中等|2212|65%|
|[321](https://leetcode-cn.com/problems/create-maximum-number/)|Create Maximum Number|困难|12510|30%|
|[322](https://leetcode-cn.com/problems/coin-change/)|Coin Change|中等|260464|40%|
|[323](https://leetcode-cn.com/problems/number-of-connected-components-in-an-undirected-graph/)|Number of Connected Components in an Undirected Graph|中等|5970|57%|
|[324](https://leetcode-cn.com/problems/wiggle-sort-ii/)|Wiggle Sort II|中等|35601|35%|
|[325](https://leetcode-cn.com/problems/maximum-size-subarray-sum-equals-k/)|Maximum Size Subarray Sum Equals k|中等|7799|49%|
|[326](https://leetcode-cn.com/problems/power-of-three/)|Power of Three|简单|108536|46%|
|[327](https://leetcode-cn.com/problems/count-of-range-sum/)|Count of Range Sum|困难|12256|34%|
|[328](https://leetcode-cn.com/problems/odd-even-linked-list/)|Odd Even Linked List|中等|73684|62%|
|[329](https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/)|Longest Increasing Path in a Matrix|困难|38979|41%|
|[330](https://leetcode-cn.com/problems/patching-array/)|Patching Array|困难|6150|41%|
|[331](https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/)|Verify Preorder Serialization of a Binary Tree|中等|15966|45%|
|[332](https://leetcode-cn.com/problems/reconstruct-itinerary/)|Reconstruct Itinerary|中等|24064|38%|
|[333](https://leetcode-cn.com/problems/largest-bst-subtree/)|Largest BST Subtree|中等|5100|43%|
|[334](https://leetcode-cn.com/problems/increasing-triplet-subsequence/)|Increasing Triplet Subsequence|中等|52458|38%|
|[335](https://leetcode-cn.com/problems/self-crossing/)|Self Crossing|困难|4212|33%|
|[336](https://leetcode-cn.com/problems/palindrome-pairs/)|Palindrome Pairs|困难|13481|34%|
|[337](https://leetcode-cn.com/problems/house-robber-iii/)|House Robber III|中等|64347|57%|
|[338](https://leetcode-cn.com/problems/counting-bits/)|Counting Bits|中等|60587|75%|
|[339](https://leetcode-cn.com/problems/nested-list-weight-sum/)|Nested List Weight Sum|简单|2485|78%|
|[340](https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters/)|Longest Substring with At Most K Distinct Characters|困难|7186|47%|
|[341](https://leetcode-cn.com/problems/flatten-nested-list-iterator/)|Flatten Nested List Iterator|中等|14975|63%|
|[342](https://leetcode-cn.com/problems/power-of-four/)|Power of Four|简单|56309|48%|
|[343](https://leetcode-cn.com/problems/integer-break/)|Integer Break|中等|56373|56%|
|[344](https://leetcode-cn.com/problems/reverse-string/)|Reverse String|简单|224814|70%|
|[345](https://leetcode-cn.com/problems/reverse-vowels-of-a-string/)|Reverse Vowels of a String|简单|80554|50%|
|[346](https://leetcode-cn.com/problems/moving-average-from-data-stream/)|Moving Average from Data Stream|简单|8099|68%|
|[347](https://leetcode-cn.com/problems/top-k-frequent-elements/)|Top K Frequent Elements|中等|108729|60%|
|[348](https://leetcode-cn.com/problems/design-tic-tac-toe/)|Design Tic-Tac-Toe|中等|3603|57%|
|[349](https://leetcode-cn.com/problems/intersection-of-two-arrays/)|Intersection of Two Arrays|简单|113900|69%|
|[350](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)|Intersection of Two Arrays II|简单|239385|51%|
|[351](https://leetcode-cn.com/problems/android-unlock-patterns/)|Android Unlock Patterns|中等|4961|57%|
|[352](https://leetcode-cn.com/problems/data-stream-as-disjoint-intervals/)|Data Stream as Disjoint Intervals|困难|3932|54%|
|[353](https://leetcode-cn.com/problems/design-snake-game/)|Design Snake Game|中等|2491|40%|
|[354](https://leetcode-cn.com/problems/russian-doll-envelopes/)|Russian Doll Envelopes|困难|36610|36%|
|[355](https://leetcode-cn.com/problems/design-twitter/)|Design Twitter|中等|39783|41%|
|[356](https://leetcode-cn.com/problems/line-reflection/)|Line Reflection|中等|2567|29%|
|[357](https://leetcode-cn.com/problems/count-numbers-with-unique-digits/)|Count Numbers with Unique Digits|中等|21178|51%|
|[358](https://leetcode-cn.com/problems/rearrange-string-k-distance-apart/)|Rearrange String k Distance Apart|困难|6331|34%|
|[359](https://leetcode-cn.com/problems/logger-rate-limiter/)|Logger Rate Limiter|简单|3412|70%|
|[360](https://leetcode-cn.com/problems/sort-transformed-array/)|Sort Transformed Array|中等|2455|60%|
|[361](https://leetcode-cn.com/problems/bomb-enemy/)|Bomb Enemy|中等|2989|53%|
|[362](https://leetcode-cn.com/problems/design-hit-counter/)|Design Hit Counter|中等|1867|67%|
|[363](https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/)|Max Sum of Rectangle No Larger Than K|困难|10524|37%|
|[364](https://leetcode-cn.com/problems/nested-list-weight-sum-ii/)|Nested List Weight Sum II|中等|1253|70%|
|[365](https://leetcode-cn.com/problems/water-and-jug-problem/)|Water and Jug Problem|中等|70711|35%|
|[366](https://leetcode-cn.com/problems/find-leaves-of-binary-tree/)|Find Leaves of Binary Tree|中等|2284|74%|
|[367](https://leetcode-cn.com/problems/valid-perfect-square/)|Valid Perfect Square|简单|85904|43%|
|[368](https://leetcode-cn.com/problems/largest-divisible-subset/)|Largest Divisible Subset|中等|20743|38%|
|[369](https://leetcode-cn.com/problems/plus-one-linked-list/)|Plus One Linked List|中等|3252|61%|
|[370](https://leetcode-cn.com/problems/range-addition/)|Range Addition|中等|1704|66%|
|[371](https://leetcode-cn.com/problems/sum-of-two-integers/)|Sum of Two Integers|简单|58021|55%|
|[372](https://leetcode-cn.com/problems/super-pow/)|Super Pow|中等|13141|42%|
|[373](https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/)|Find K Pairs with Smallest Sums|中等|20845|42%|
|[374](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)|[Guess Number Higher or Lower](./src/0374-guess-number-higher-or-lower/)|简单|58864|44%|
|[375](https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/)|Guess Number Higher or Lower II|中等|17818|39%|
|[376](https://leetcode-cn.com/problems/wiggle-subsequence/)|Wiggle Subsequence|中等|50634|41%|
|[377](https://leetcode-cn.com/problems/combination-sum-iv/)|Combination Sum IV|中等|31558|42%|
|[378](https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/)|Kth Smallest Element in a Sorted Matrix|中等|75942|62%|
|[379](https://leetcode-cn.com/problems/design-phone-directory/)|Design Phone Directory|中等|1944|66%|
|[380](https://leetcode-cn.com/problems/insert-delete-getrandom-o1/)|Insert Delete GetRandom O(1)|中等|29256|48%|
|[381](https://leetcode-cn.com/problems/insert-delete-getrandom-o1-duplicates-allowed/)|Insert Delete GetRandom O(1) - Duplicates allowed|困难|8279|37%|
|[382](https://leetcode-cn.com/problems/linked-list-random-node/)|Linked List Random Node|中等|10392|57%|
|[383](https://leetcode-cn.com/problems/ransom-note/)|[Ransom Note](./src/0383-ransom-note/)|简单|46563|53%|
|[384](https://leetcode-cn.com/problems/shuffle-an-array/)|Shuffle an Array|中等|42660|52%|
|[385](https://leetcode-cn.com/problems/mini-parser/)|Mini Parser|中等|9550|39%|
|[386](https://leetcode-cn.com/problems/lexicographical-numbers/)|Lexicographical Numbers|中等|14137|70%|
|[387](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)|[First Unique Character in a String](./src/0387-first-unique-character-in-a-string/)|简单|188608|45%|
|[388](https://leetcode-cn.com/problems/longest-absolute-file-path/)|Longest Absolute File Path|中等|5415|47%|
|[389](https://leetcode-cn.com/problems/find-the-difference/)|Find the Difference|简单|49175|62%|
|[390](https://leetcode-cn.com/problems/elimination-game/)|Elimination Game|中等|7896|44%|
|[391](https://leetcode-cn.com/problems/perfect-rectangle/)|Perfect Rectangle|困难|5993|26%|
|[392](https://leetcode-cn.com/problems/is-subsequence/)|Is Subsequence|简单|100686|48%|
|[393](https://leetcode-cn.com/problems/utf-8-validation/)|UTF-8 Validation|中等|16125|38%|
|[394](https://leetcode-cn.com/problems/decode-string/)|Decode String|中等|97177|52%|
|[395](https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/)|Longest Substring with At Least K Repeating Characters|中等|26285|43%|
|[396](https://leetcode-cn.com/problems/rotate-function/)|Rotate Function|中等|8670|40%|
|[397](https://leetcode-cn.com/problems/integer-replacement/)|Integer Replacement|中等|21181|35%|
|[398](https://leetcode-cn.com/problems/random-pick-index/)|Random Pick Index|中等|9606|62%|
|[399](https://leetcode-cn.com/problems/evaluate-division/)|Evaluate Division|中等|17356|54%|
|[400](https://leetcode-cn.com/problems/nth-digit/)|Nth Digit|中等|25664|37%|
|[401](https://leetcode-cn.com/problems/binary-watch/)|Binary Watch|简单|30462|52%|
|[402](https://leetcode-cn.com/problems/remove-k-digits/)|Remove K Digits|中等|71232|29%|
|[403](https://leetcode-cn.com/problems/frog-jump/)|Frog Jump|困难|20146|36%|
|[404](https://leetcode-cn.com/problems/sum-of-left-leaves/)|Sum of Left Leaves|简单|53570|55%|
|[405](https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/)|Convert a Number to Hexadecimal|简单|25347|50%|
|[406](https://leetcode-cn.com/problems/queue-reconstruction-by-height/)|Queue Reconstruction by Height|中等|52753|65%|
|[407](https://leetcode-cn.com/problems/trapping-rain-water-ii/)|Trapping Rain Water II|困难|10342|41%|
|[408](https://leetcode-cn.com/problems/valid-word-abbreviation/)|Valid Word Abbreviation|简单|5673|30%|
|[409](https://leetcode-cn.com/problems/longest-palindrome/)|Longest Palindrome|简单|97128|55%|
|[410](https://leetcode-cn.com/problems/split-array-largest-sum/)|Split Array Largest Sum|困难|17839|43%|
|[411](https://leetcode-cn.com/problems/minimum-unique-word-abbreviation/)|Minimum Unique Word Abbreviation|困难|1036|47%|
|[412](https://leetcode-cn.com/problems/fizz-buzz/)|Fizz Buzz|简单|60811|64%|
|[413](https://leetcode-cn.com/problems/arithmetic-slices/)|Arithmetic Slices|中等|24986|62%|
|[414](https://leetcode-cn.com/problems/third-maximum-number/)|Third Maximum Number|简单|82856|35%|
|[415](https://leetcode-cn.com/problems/add-strings/)|[Add Strings](./src/0415-add-strings/)|简单|81184|50%|
|[416](https://leetcode-cn.com/problems/partition-equal-subset-sum/)|Partition Equal Subset Sum|中等|83530|48%|
|[417](https://leetcode-cn.com/problems/pacific-atlantic-water-flow/)|Pacific Atlantic Water Flow|中等|22895|42%|
|[418](https://leetcode-cn.com/problems/sentence-screen-fitting/)|Sentence Screen Fitting|中等|3103|34%|
|[419](https://leetcode-cn.com/problems/battleships-in-a-board/)|Battleships in a Board|中等|6605|73%|
|[420](https://leetcode-cn.com/problems/strong-password-checker/)|Strong Password Checker|困难|8653|16%|
|[421](https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/)|[Maximum XOR of Two Numbers in an Array](./src/0421-maximum-xor-of-two-numbers-in-an-array/)|中等|8958|59%|
|[422](https://leetcode-cn.com/problems/valid-word-square/)|Valid Word Square|简单|2291|40%|
|[423](https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/)|Reconstruct Original Digits from English|中等|7661|53%|
|[424](https://leetcode-cn.com/problems/longest-repeating-character-replacement/)|Longest Repeating Character Replacement|中等|17879|47%|
|[425](https://leetcode-cn.com/problems/word-squares/)|Word Squares|困难|1048|56%|
|[432](https://leetcode-cn.com/problems/all-oone-data-structure/)|All O`one Data Structure|困难|10584|34%|
|[433](https://leetcode-cn.com/problems/minimum-genetic-mutation/)|Minimum Genetic Mutation|中等|11798|51%|
|[434](https://leetcode-cn.com/problems/number-of-segments-in-a-string/)|Number of Segments in a String|简单|53098|35%|
|[435](https://leetcode-cn.com/problems/non-overlapping-intervals/)|Non-overlapping Intervals|中等|41107|45%|
|[436](https://leetcode-cn.com/problems/find-right-interval/)|Find Right Interval|中等|7873|46%|
|[437](https://leetcode-cn.com/problems/path-sum-iii/)|Path Sum III|简单|77634|55%|
|[438](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)|Find All Anagrams in a String|中等|70246|45%|
|[439](https://leetcode-cn.com/problems/ternary-expression-parser/)|Ternary Expression Parser|中等|2953|57%|
|[440](https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/)|K-th Smallest in Lexicographical Order|困难|17451|34%|
|[441](https://leetcode-cn.com/problems/arranging-coins/)|Arranging Coins|简单|58775|41%|
|[442](https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/)|Find All Duplicates in an Array|中等|29581|66%|
|[443](https://leetcode-cn.com/problems/string-compression/)|String Compression|简单|45474|40%|
|[444](https://leetcode-cn.com/problems/sequence-reconstruction/)|Sequence Reconstruction|中等|6142|21%|
|[445](https://leetcode-cn.com/problems/add-two-numbers-ii/)|Add Two Numbers II|中等|77226|57%|
|[446](https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence/)|Arithmetic Slices II - Subsequence|困难|6140|34%|
|[447](https://leetcode-cn.com/problems/number-of-boomerangs/)|Number of Boomerangs|简单|22944|57%|
|[448](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/)|[Find All Numbers Disappeared in an Array](./src/0448-find-all-numbers-disappeared-in-an-array/)|简单|74893|59%|
|[449](https://leetcode-cn.com/problems/serialize-and-deserialize-bst/)|Serialize and Deserialize BST|中等|12280|52%|
|[450](https://leetcode-cn.com/problems/delete-node-in-a-bst/)|Delete Node in a BST|中等|39597|42%|
|[451](https://leetcode-cn.com/problems/sort-characters-by-frequency/)|Sort Characters By Frequency|中等|38136|64%|
|[452](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/)|Minimum Number of Arrows to Burst Balloons|中等|36416|50%|
|[453](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/)|Minimum Moves to Equal Array Elements|简单|22067|54%|
|[454](https://leetcode-cn.com/problems/4sum-ii/)|4Sum II|中等|42662|55%|
|[455](https://leetcode-cn.com/problems/assign-cookies/)|Assign Cookies|简单|80254|54%|
|[456](https://leetcode-cn.com/problems/132-pattern/)|132 Pattern|中等|36352|27%|
|[457](https://leetcode-cn.com/problems/circular-array-loop/)|Circular Array Loop|中等|13166|34%|
|[458](https://leetcode-cn.com/problems/poor-pigs/)|Poor Pigs|困难|6737|56%|
|[459](https://leetcode-cn.com/problems/repeated-substring-pattern/)|Repeated Substring Pattern|简单|44046|47%|
|[460](https://leetcode-cn.com/problems/lfu-cache/)|LFU Cache|困难|37119|42%|
|[461](https://leetcode-cn.com/problems/hamming-distance/)|Hamming Distance|简单|81171|76%|
|[462](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/)|Minimum Moves to Equal Array Elements II|中等|10982|58%|
|[463](https://leetcode-cn.com/problems/island-perimeter/)|Island Perimeter|简单|29126|67%|
|[464](https://leetcode-cn.com/problems/can-i-win/)|Can I Win|中等|14942|33%|
|[465](https://leetcode-cn.com/problems/optimal-account-balancing/)|Optimal Account Balancing|困难|1396|50%|
|[466](https://leetcode-cn.com/problems/count-the-repetitions/)|Count The Repetitions|困难|25687|37%|
|[467](https://leetcode-cn.com/problems/unique-substrings-in-wraparound-string/)|Unique Substrings in Wraparound String|中等|7683|40%|
|[468](https://leetcode-cn.com/problems/validate-ip-address/)|Validate IP Address|中等|39596|22%|
|[469](https://leetcode-cn.com/problems/convex-polygon/)|Convex Polygon|中等|1517|45%|
|[471](https://leetcode-cn.com/problems/encode-string-with-shortest-length/)|Encode String with Shortest Length|困难|1028|57%|
|[472](https://leetcode-cn.com/problems/concatenated-words/)|Concatenated Words|困难|7420|44%|
|[473](https://leetcode-cn.com/problems/matchsticks-to-square/)|Matchsticks to Square|中等|24357|39%|
|[474](https://leetcode-cn.com/problems/ones-and-zeroes/)|Ones and Zeroes|中等|20577|54%|
|[475](https://leetcode-cn.com/problems/heaters/)|Heaters|简单|37523|30%|
|[476](https://leetcode-cn.com/problems/number-complement/)|Number Complement|简单|33101|68%|
|[477](https://leetcode-cn.com/problems/total-hamming-distance/)|[Total Hamming Distance](./src/0477-total-hamming-distance/)|中等|12897|50%|
|[479](https://leetcode-cn.com/problems/largest-palindrome-product/)|Largest Palindrome Product|困难|4733|34%|
|[480](https://leetcode-cn.com/problems/sliding-window-median/)|Sliding Window Median|困难|12760|35%|
|[481](https://leetcode-cn.com/problems/magical-string/)|Magical String|中等|5073|50%|
|[482](https://leetcode-cn.com/problems/license-key-formatting/)|License Key Formatting|简单|23847|39%|
|[483](https://leetcode-cn.com/problems/smallest-good-base/)|Smallest Good Base|困难|3291|39%|
|[484](https://leetcode-cn.com/problems/find-permutation/)|Find Permutation|中等|1144|65%|
|[485](https://leetcode-cn.com/problems/max-consecutive-ones/)|Max Consecutive Ones|简单|77319|56%|
|[486](https://leetcode-cn.com/problems/predict-the-winner/)|Predict the Winner|中等|15419|52%|
|[487](https://leetcode-cn.com/problems/max-consecutive-ones-ii/)|Max Consecutive Ones II|中等|3011|55%|
|[488](https://leetcode-cn.com/problems/zuma-game/)|Zuma Game|困难|4388|41%|
|[490](https://leetcode-cn.com/problems/the-maze/)|The Maze|中等|7345|45%|
|[491](https://leetcode-cn.com/problems/increasing-subsequences/)|Increasing Subsequences|中等|15949|48%|
|[492](https://leetcode-cn.com/problems/construct-the-rectangle/)|Construct the Rectangle|简单|19256|52%|
|[493](https://leetcode-cn.com/problems/reverse-pairs/)|Reverse Pairs|困难|20866|26%|
|[494](https://leetcode-cn.com/problems/target-sum/)|Target Sum|中等|78523|44%|
|[495](https://leetcode-cn.com/problems/teemo-attacking/)|Teemo Attacking|中等|19877|53%|
|[496](https://leetcode-cn.com/problems/next-greater-element-i/)|Next Greater Element I|简单|57787|65%|
|[498](https://leetcode-cn.com/problems/diagonal-traverse/)|Diagonal Traverse|中等|47287|41%|
|[499](https://leetcode-cn.com/problems/the-maze-iii/)|The Maze III|困难|4389|32%|
|[500](https://leetcode-cn.com/problems/keyboard-row/)|Keyboard Row|简单|26347|69%|
|[501](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)|Find Mode in Binary Search Tree|简单|35273|45%|
|[502](https://leetcode-cn.com/problems/ipo/)|IPO|困难|8023|39%|
|[503](https://leetcode-cn.com/problems/next-greater-element-ii/)|[Next Greater Element II](./src/0503-next-greater-element-ii/)|中等|43423|56%|
|[504](https://leetcode-cn.com/problems/base-7/)|Base 7|简单|27833|49%|
|[505](https://leetcode-cn.com/problems/the-maze-ii/)|The Maze II|中等|5950|44%|
|[506](https://leetcode-cn.com/problems/relative-ranks/)|Relative Ranks|简单|18331|54%|
|[507](https://leetcode-cn.com/problems/perfect-number/)|Perfect Number|简单|40938|38%|
|[508](https://leetcode-cn.com/problems/most-frequent-subtree-sum/)|Most Frequent Subtree Sum|中等|10760|64%|
|[510](https://leetcode-cn.com/problems/inorder-successor-in-bst-ii/)|Inorder Successor in BST II|中等|2188|57%|
|[1059](https://leetcode-cn.com/problems/all-paths-from-source-lead-to-destination/)|All Paths from Source Lead to Destination|中等|3483|31%|
|[513](https://leetcode-cn.com/problems/find-bottom-left-tree-value/)|Find Bottom Left Tree Value|中等|24305|70%|
|[514](https://leetcode-cn.com/problems/freedom-trail/)|Freedom Trail|困难|4883|39%|
|[515](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)|Find Largest Value in Each Tree Row|中等|23602|60%|
|[516](https://leetcode-cn.com/problems/longest-palindromic-subsequence/)|Longest Palindromic Subsequence|中等|41281|56%|
|[517](https://leetcode-cn.com/problems/super-washing-machines/)|Super Washing Machines|困难|4335|41%|
|[518](https://leetcode-cn.com/problems/coin-change-2/)|Coin Change 2|中等|35798|53%|
|[520](https://leetcode-cn.com/problems/detect-capital/)|Detect Capital|简单|41584|55%|
|[521](https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/)|Longest Uncommon Subsequence I |简单|20517|68%|
|[522](https://leetcode-cn.com/problems/longest-uncommon-subsequence-ii/)|Longest Uncommon Subsequence II|中等|11164|34%|
|[523](https://leetcode-cn.com/problems/continuous-subarray-sum/)|Continuous Subarray Sum|中等|77375|22%|
|[524](https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/)|Longest Word in Dictionary through Deleting|中等|33977|45%|
|[525](https://leetcode-cn.com/problems/contiguous-array/)|Contiguous Array|中等|13536|43%|
|[526](https://leetcode-cn.com/problems/beautiful-arrangement/)|Beautiful Arrangement|中等|9945|62%|
|[527](https://leetcode-cn.com/problems/word-abbreviation/)|Word Abbreviation|困难|2122|53%|
|[529](https://leetcode-cn.com/problems/minesweeper/)|Minesweeper|中等|15449|60%|
|[530](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)|Minimum Absolute Difference in BST|简单|30598|57%|
|[531](https://leetcode-cn.com/problems/lonely-pixel-i/)|Lonely Pixel I|中等|2764|66%|
|[532](https://leetcode-cn.com/problems/k-diff-pairs-in-an-array/)|K-diff Pairs in an Array|简单|49294|34%|
|[533](https://leetcode-cn.com/problems/lonely-pixel-ii/)|Lonely Pixel II|中等|1625|48%|
|[535](https://leetcode-cn.com/problems/encode-and-decode-tinyurl/)|Encode and Decode TinyURL|中等|11203|82%|
|[536](https://leetcode-cn.com/problems/construct-binary-tree-from-string/)|Construct Binary Tree from String|中等|3161|51%|
|[537](https://leetcode-cn.com/problems/complex-number-multiplication/)|Complex Number Multiplication|中等|9339|69%|
|[538](https://leetcode-cn.com/problems/convert-bst-to-greater-tree/)|Convert BST to Greater Tree|简单|50402|62%|
|[539](https://leetcode-cn.com/problems/minimum-time-difference/)|Minimum Time Difference|中等|12284|55%|
|[540](https://leetcode-cn.com/problems/single-element-in-a-sorted-array/)|Single Element in a Sorted Array|中等|20764|60%|
|[541](https://leetcode-cn.com/problems/reverse-string-ii/)|Reverse String II|简单|32698|54%|
|[542](https://leetcode-cn.com/problems/01-matrix/)|01 Matrix|中等|79326|44%|
|[543](https://leetcode-cn.com/problems/diameter-of-binary-tree/)|Diameter of Binary Tree|简单|115555|50%|
|[544](https://leetcode-cn.com/problems/output-contest-matches/)|Output Contest Matches|中等|2088|69%|
|[545](https://leetcode-cn.com/problems/boundary-of-binary-tree/)|Boundary of Binary Tree|中等|3573|38%|
|[546](https://leetcode-cn.com/problems/remove-boxes/)|Remove Boxes|困难|6448|51%|
|[547](https://leetcode-cn.com/problems/friend-circles/)|Friend Circles|中等|91327|57%|
|[548](https://leetcode-cn.com/problems/split-array-with-equal-sum/)|Split Array with Equal Sum|中等|2861|32%|
|[549](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence-ii/)|Binary Tree Longest Consecutive Sequence II|中等|1440|47%|
|[551](https://leetcode-cn.com/problems/student-attendance-record-i/)|Student Attendance Record I|简单|31753|51%|
|[552](https://leetcode-cn.com/problems/student-attendance-record-ii/)|Student Attendance Record II|困难|6889|40%|
|[553](https://leetcode-cn.com/problems/optimal-division/)|Optimal Division|中等|4841|56%|
|[554](https://leetcode-cn.com/problems/brick-wall/)|Brick Wall|中等|33537|37%|
|[555](https://leetcode-cn.com/problems/split-concatenated-strings/)|Split Concatenated Strings|中等|2034|32%|
|[556](https://leetcode-cn.com/problems/next-greater-element-iii/)|Next Greater Element III|中等|17555|31%|
|[557](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)|Reverse Words in a String III|简单|90754|71%|
|[560](https://leetcode-cn.com/problems/subarray-sum-equals-k/)|Subarray Sum Equals K|中等|133682|44%|
|[561](https://leetcode-cn.com/problems/array-partition-i/)|Array Partition I|简单|57926|71%|
|[562](https://leetcode-cn.com/problems/longest-line-of-consecutive-one-in-matrix/)|Longest Line of Consecutive One in Matrix|中等|5040|42%|
|[563](https://leetcode-cn.com/problems/binary-tree-tilt/)|Binary Tree Tilt|简单|25109|55%|
|[564](https://leetcode-cn.com/problems/find-the-closest-palindrome/)|Find the Closest Palindrome|困难|14409|15%|
|[565](https://leetcode-cn.com/problems/array-nesting/)|Array Nesting|中等|11563|58%|
|[566](https://leetcode-cn.com/problems/reshape-the-matrix/)|Reshape the Matrix|简单|29467|64%|
|[567](https://leetcode-cn.com/problems/permutation-in-string/)|Permutation in String|中等|94470|36%|
|[568](https://leetcode-cn.com/problems/maximum-vacation-days/)|Maximum Vacation Days|困难|1293|48%|
|[569](https://leetcode-cn.com/problems/median-employee-salary/)|Median Employee Salary|困难|4095|52%|
|[570](https://leetcode-cn.com/problems/managers-with-at-least-5-direct-reports/)|Managers with at Least 5 Direct Reports|中等|5151|67%|
|[571](https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers/)|Find Median Given Frequency of Numbers|困难|2570|49%|
|[572](https://leetcode-cn.com/problems/subtree-of-another-tree/)|Subtree of Another Tree|简单|94506|46%|
|[573](https://leetcode-cn.com/problems/squirrel-simulation/)|Squirrel Simulation|中等|712|63%|
|[574](https://leetcode-cn.com/problems/winning-candidate/)|Winning Candidate|中等|5326|59%|
|[575](https://leetcode-cn.com/problems/distribute-candies/)|Distribute Candies|简单|34259|67%|
|[576](https://leetcode-cn.com/problems/out-of-boundary-paths/)|Out of Boundary Paths|中等|12448|36%|
|[577](https://leetcode-cn.com/problems/employee-bonus/)|Employee Bonus|简单|6289|70%|
|[578](https://leetcode-cn.com/problems/get-highest-answer-rate-question/)|Get Highest Answer Rate Question|中等|5240|49%|
|[579](https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee/)|Find Cumulative Salary of an Employee|困难|4633|42%|
|[580](https://leetcode-cn.com/problems/count-student-number-in-departments/)|Count Student Number in Departments|中等|5278|51%|
|[581](https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/)|Shortest Unsorted Continuous Subarray|简单|90929|34%|
|[582](https://leetcode-cn.com/problems/kill-process/)|Kill Process|中等|15987|36%|
|[583](https://leetcode-cn.com/problems/delete-operation-for-two-strings/)|Delete Operation for Two Strings|中等|17389|49%|
|[584](https://leetcode-cn.com/problems/find-customer-referee/)|Find Customer Referee|简单|5758|76%|
|[585](https://leetcode-cn.com/problems/investments-in-2016/)|Investments in 2016|中等|4048|57%|
|[586](https://leetcode-cn.com/problems/customer-placing-the-largest-number-of-orders/)|Customer Placing the Largest Number of Orders|简单|5914|75%|
|[587](https://leetcode-cn.com/problems/erect-the-fence/)|Erect the Fence|困难|2504|31%|
|[588](https://leetcode-cn.com/problems/design-in-memory-file-system/)|Design In-Memory File System|困难|1753|38%|
|[591](https://leetcode-cn.com/problems/tag-validator/)|Tag Validator|困难|2309|30%|
|[592](https://leetcode-cn.com/problems/fraction-addition-and-subtraction/)|Fraction Addition and Subtraction|中等|4192|49%|
|[593](https://leetcode-cn.com/problems/valid-square/)|Valid Square|中等|10350|43%|
|[594](https://leetcode-cn.com/problems/longest-harmonious-subsequence/)|Longest Harmonious Subsequence|简单|27386|47%|
|[595](https://leetcode-cn.com/problems/big-countries/)|[Big Countries](./src/0595-big-countries/)|简单|70980|78%|
|[596](https://leetcode-cn.com/problems/classes-more-than-5-students/)|[Classes More Than 5 Students](./src/0596-classes-more-than-5-students/)|简单|96130|40%|
|[597](https://leetcode-cn.com/problems/friend-requests-i-overall-acceptance-rate/)|Friend Requests I: Overall Acceptance Rate|简单|7023|44%|
|[598](https://leetcode-cn.com/problems/range-addition-ii/)|Range Addition II|简单|14823|50%|
|[599](https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/)|Minimum Index Sum of Two Lists|简单|29735|51%|
|[600](https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/)|[Non-negative Integers without Consecutive Ones](./src/0600-non-negative-integers-without-consecutive-ones/)|困难|6738|31%|
|[601](https://leetcode-cn.com/problems/human-traffic-of-stadium/)|[Human Traffic of Stadium](./src/0601-human-traffic-of-stadium/)|困难|31145|45%|
|[602](https://leetcode-cn.com/problems/friend-requests-ii-who-has-the-most-friends/)|Friend Requests II: Who Has the Most Friends|中等|3578|61%|
|[603](https://leetcode-cn.com/problems/consecutive-available-seats/)|Consecutive Available Seats|简单|5790|68%|
|[604](https://leetcode-cn.com/problems/design-compressed-string-iterator/)|Design Compressed String Iterator|简单|2997|34%|
|[605](https://leetcode-cn.com/problems/can-place-flowers/)|Can Place Flowers|简单|90234|32%|
|[606](https://leetcode-cn.com/problems/construct-string-from-binary-tree/)|Construct String from Binary Tree|简单|25556|53%|
|[607](https://leetcode-cn.com/problems/sales-person/)|Sales Person|简单|5575|66%|
|[608](https://leetcode-cn.com/problems/tree-node/)|Tree Node|中等|3486|68%|
|[609](https://leetcode-cn.com/problems/find-duplicate-file-in-system/)|Find Duplicate File in System|中等|6757|46%|
|[610](https://leetcode-cn.com/problems/triangle-judgement/)|Triangle Judgement|简单|4915|65%|
|[611](https://leetcode-cn.com/problems/valid-triangle-number/)|Valid Triangle Number|中等|15103|48%|
|[612](https://leetcode-cn.com/problems/shortest-distance-in-a-plane/)|Shortest Distance in a Plane|中等|3193|64%|
|[613](https://leetcode-cn.com/problems/shortest-distance-in-a-line/)|Shortest Distance in a Line|简单|4953|82%|
|[614](https://leetcode-cn.com/problems/second-degree-follower/)|Second Degree Follower|中等|6478|32%|
|[615](https://leetcode-cn.com/problems/average-salary-departments-vs-company/)|Average Salary: Departments VS Company|困难|3221|41%|
|[616](https://leetcode-cn.com/problems/add-bold-tag-in-string/)|Add Bold Tag in String|中等|3664|44%|
|[617](https://leetcode-cn.com/problems/merge-two-binary-trees/)|Merge Two Binary Trees|简单|77853|76%|
|[618](https://leetcode-cn.com/problems/students-report-by-geography/)|Students Report By Geography|困难|1911|55%|
|[619](https://leetcode-cn.com/problems/biggest-single-number/)|Biggest Single Number|简单|7930|48%|
|[620](https://leetcode-cn.com/problems/not-boring-movies/)|[Not Boring Movies](./src/0620-not-boring-movies/)|简单|63342|76%|
|[621](https://leetcode-cn.com/problems/task-scheduler/)|Task Scheduler|中等|53003|49%|
|[623](https://leetcode-cn.com/problems/add-one-row-to-tree/)|Add One Row to Tree|中等|9729|52%|
|[624](https://leetcode-cn.com/problems/maximum-distance-in-arrays/)|Maximum Distance in Arrays|简单|3373|40%|
|[625](https://leetcode-cn.com/problems/minimum-factorization/)|Minimum Factorization|中等|4404|32%|
|[626](https://leetcode-cn.com/problems/exchange-seats/)|[Exchange Seats](./src/0626-exchange-seats/)|中等|32227|66%|
|[627](https://leetcode-cn.com/problems/swap-salary/)|[Swap Salary](./src/0627-swap-salary/)|简单|53713|77%|
|[628](https://leetcode-cn.com/problems/maximum-product-of-three-numbers/)|Maximum Product of Three Numbers|简单|43740|50%|
|[629](https://leetcode-cn.com/problems/k-inverse-pairs-array/)|K Inverse Pairs Array|困难|3984|37%|
|[630](https://leetcode-cn.com/problems/course-schedule-iii/)|Course Schedule III|困难|9094|31%|
|[631](https://leetcode-cn.com/problems/design-excel-sum-formula/)|Design Excel Sum Formula|困难|2656|26%|
|[632](https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/)|Smallest Range Covering Elements from K Lists|困难|7510|38%|
|[633](https://leetcode-cn.com/problems/sum-of-square-numbers/)|Sum of Square Numbers|简单|73432|33%|
|[634](https://leetcode-cn.com/problems/find-the-derangement-of-an-array/)|Find the Derangement of An Array|中等|1093|41%|
|[635](https://leetcode-cn.com/problems/design-log-storage-system/)|Design Log Storage System|中等|2892|55%|
|[636](https://leetcode-cn.com/problems/exclusive-time-of-functions/)|Exclusive Time of Functions|中等|8236|50%|
|[637](https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/)|Average of Levels in Binary Tree|简单|34209|64%|
|[638](https://leetcode-cn.com/problems/shopping-offers/)|Shopping Offers|中等|10178|56%|
|[639](https://leetcode-cn.com/problems/decode-ways-ii/)|Decode Ways II|困难|7757|27%|
|[640](https://leetcode-cn.com/problems/solve-the-equation/)|Solve the Equation|中等|14842|40%|
|[642](https://leetcode-cn.com/problems/design-search-autocomplete-system/)|Design Search Autocomplete System|困难|2270|45%|
|[643](https://leetcode-cn.com/problems/maximum-average-subarray-i/)|Maximum Average Subarray I|简单|41882|38%|
|[644](https://leetcode-cn.com/problems/maximum-average-subarray-ii/)|Maximum Average Subarray II|困难|1504|33%|
|[645](https://leetcode-cn.com/problems/set-mismatch/)|Set Mismatch|简单|39738|42%|
|[646](https://leetcode-cn.com/problems/maximum-length-of-pair-chain/)|Maximum Length of Pair Chain|中等|16605|55%|
|[647](https://leetcode-cn.com/problems/palindromic-substrings/)|Palindromic Substrings|中等|52185|61%|
|[648](https://leetcode-cn.com/problems/replace-words/)|Replace Words|中等|15805|53%|
|[649](https://leetcode-cn.com/problems/dota2-senate/)|Dota2 Senate|中等|9681|38%|
|[650](https://leetcode-cn.com/problems/2-keys-keyboard/)|2 Keys Keyboard|中等|21052|50%|
|[651](https://leetcode-cn.com/problems/4-keys-keyboard/)|4 Keys Keyboard|中等|2782|57%|
|[652](https://leetcode-cn.com/problems/find-duplicate-subtrees/)|Find Duplicate Subtrees|中等|17169|53%|
|[653](https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/)|Two Sum IV - Input is a BST|简单|32474|55%|
|[654](https://leetcode-cn.com/problems/maximum-binary-tree/)|Maximum Binary Tree|中等|20684|81%|
|[655](https://leetcode-cn.com/problems/print-binary-tree/)|Print Binary Tree|中等|7785|56%|
|[656](https://leetcode-cn.com/problems/coin-path/)|Coin Path|困难|1602|29%|
|[657](https://leetcode-cn.com/problems/robot-return-to-origin/)|Robot Return to Origin|简单|46412|74%|
|[658](https://leetcode-cn.com/problems/find-k-closest-elements/)|Find K Closest Elements|中等|25500|43%|
|[659](https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/)|Split Array into Consecutive Subsequences|中等|8913|41%|
|[660](https://leetcode-cn.com/problems/remove-9/)|Remove 9|困难|528|61%|
|[661](https://leetcode-cn.com/problems/image-smoother/)|Image Smoother|简单|16132|53%|
|[662](https://leetcode-cn.com/problems/maximum-width-of-binary-tree/)|Maximum Width of Binary Tree|中等|23766|38%|
|[663](https://leetcode-cn.com/problems/equal-tree-partition/)|Equal Tree Partition|中等|1818|45%|
|[664](https://leetcode-cn.com/problems/strange-printer/)|Strange Printer|困难|4987|46%|
|[665](https://leetcode-cn.com/problems/non-decreasing-array/)|Non-decreasing Array|简单|98503|22%|
|[666](https://leetcode-cn.com/problems/path-sum-iv/)|Path Sum IV|中等|1637|59%|
|[667](https://leetcode-cn.com/problems/beautiful-arrangement-ii/)|Beautiful Arrangement II|中等|6543|60%|
|[668](https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table/)|Kth Smallest Number in Multiplication Table|困难|6611|47%|
|[669](https://leetcode-cn.com/problems/trim-a-binary-search-tree/)|Trim a Binary Search Tree|简单|22036|66%|
|[670](https://leetcode-cn.com/problems/maximum-swap/)|Maximum Swap|中等|17396|40%|
|[671](https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/)|Second Minimum Node In a Binary Tree|简单|28781|46%|
|[672](https://leetcode-cn.com/problems/bulb-switcher-ii/)|Bulb Switcher II|中等|3386|54%|
|[673](https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/)|Number of Longest Increasing Subsequence|中等|27814|35%|
|[674](https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/)|Longest Continuous Increasing Subsequence|简单|69440|45%|
|[675](https://leetcode-cn.com/problems/cut-off-trees-for-golf-event/)|Cut Off Trees for Golf Event|困难|3675|36%|
|[676](https://leetcode-cn.com/problems/implement-magic-dictionary/)|Implement Magic Dictionary|中等|5135|56%|
|[677](https://leetcode-cn.com/problems/map-sum-pairs/)|Map Sum Pairs|中等|12119|61%|
|[678](https://leetcode-cn.com/problems/valid-parenthesis-string/)|Valid Parenthesis String|中等|20262|32%|
|[679](https://leetcode-cn.com/problems/24-game/)|24 Game|困难|11942|44%|
|[680](https://leetcode-cn.com/problems/valid-palindrome-ii/)|Valid Palindrome II|简单|116339|39%|
|[681](https://leetcode-cn.com/problems/next-closest-time/)|Next Closest Time|中等|2358|49%|
|[682](https://leetcode-cn.com/problems/baseball-game/)|Baseball Game|简单|35087|67%|
|[683](https://leetcode-cn.com/problems/k-empty-slots/)|K Empty Slots|困难|2573|39%|
|[684](https://leetcode-cn.com/problems/redundant-connection/)|Redundant Connection|中等|24888|59%|
|[685](https://leetcode-cn.com/problems/redundant-connection-ii/)|Redundant Connection II|困难|10203|35%|
|[686](https://leetcode-cn.com/problems/repeated-string-match/)|Repeated String Match|简单|32419|33%|
|[687](https://leetcode-cn.com/problems/longest-univalue-path/)|Longest Univalue Path|简单|47879|40%|
|[688](https://leetcode-cn.com/problems/knight-probability-in-chessboard/)|Knight Probability in Chessboard|中等|8199|47%|
|[689](https://leetcode-cn.com/problems/maximum-sum-of-3-non-overlapping-subarrays/)|Maximum Sum of 3 Non-Overlapping Subarrays|困难|3381|45%|
|[690](https://leetcode-cn.com/problems/employee-importance/)|Employee Importance|简单|25111|58%|
|[691](https://leetcode-cn.com/problems/stickers-to-spell-word/)|Stickers to Spell Word|困难|3380|43%|
|[692](https://leetcode-cn.com/problems/top-k-frequent-words/)|Top K Frequent Words|中等|29146|51%|
|[693](https://leetcode-cn.com/problems/binary-number-with-alternating-bits/)|Binary Number with Alternating Bits|简单|24311|60%|
|[694](https://leetcode-cn.com/problems/number-of-distinct-islands/)|Number of Distinct Islands|中等|3453|47%|
|[695](https://leetcode-cn.com/problems/max-area-of-island/)|Max Area of Island|中等|83666|63%|
|[696](https://leetcode-cn.com/problems/count-binary-substrings/)|Count Binary Substrings|简单|22165|53%|
|[697](https://leetcode-cn.com/problems/degree-of-an-array/)|Degree of an Array|简单|34598|53%|
|[698](https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/)|Partition to K Equal Sum Subsets|中等|29357|42%|
|[699](https://leetcode-cn.com/problems/falling-squares/)|Falling Squares|困难|2524|40%|
|[711](https://leetcode-cn.com/problems/number-of-distinct-islands-ii/)|Number of Distinct Islands II|困难|696|54%|
|[712](https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/)|Minimum ASCII Delete Sum for Two Strings|中等|8221|63%|
|[713](https://leetcode-cn.com/problems/subarray-product-less-than-k/)|Subarray Product Less Than K|中等|21831|36%|
|[714](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)|Best Time to Buy and Sell Stock with Transaction Fee|中等|36019|66%|
|[715](https://leetcode-cn.com/problems/range-module/)|Range Module|困难|3226|33%|
|[716](https://leetcode-cn.com/problems/max-stack/)|Max Stack|简单|3735|45%|
|[717](https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/)|1-bit and 2-bit Characters|简单|37853|48%|
|[718](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/)|Maximum Length of Repeated Subarray|中等|72026|53%|
|[719](https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance/)|Find K-th Smallest Pair Distance|困难|13155|33%|
|[720](https://leetcode-cn.com/problems/longest-word-in-dictionary/)|Longest Word in Dictionary|简单|19910|46%|
|[721](https://leetcode-cn.com/problems/accounts-merge/)|Accounts Merge|中等|13230|34%|
|[722](https://leetcode-cn.com/problems/remove-comments/)|Remove Comments|中等|8897|28%|
|[723](https://leetcode-cn.com/problems/candy-crush/)|Candy Crush|中等|934|68%|
|[724](https://leetcode-cn.com/problems/find-pivot-index/)|Find Pivot Index|简单|125239|37%|
|[725](https://leetcode-cn.com/problems/split-linked-list-in-parts/)|Split Linked List in Parts|中等|19799|54%|
|[726](https://leetcode-cn.com/problems/number-of-atoms/)|Number of Atoms|困难|7223|45%|
|[727](https://leetcode-cn.com/problems/minimum-window-subsequence/)|Minimum Window Subsequence|困难|3092|37%|
|[728](https://leetcode-cn.com/problems/self-dividing-numbers/)|Self Dividing Numbers|简单|30024|73%|
|[729](https://leetcode-cn.com/problems/my-calendar-i/)|My Calendar I|中等|7532|48%|
|[730](https://leetcode-cn.com/problems/count-different-palindromic-subsequences/)|Count Different Palindromic Subsequences|困难|2718|47%|
|[731](https://leetcode-cn.com/problems/my-calendar-ii/)|My Calendar II|中等|4031|48%|
|[732](https://leetcode-cn.com/problems/my-calendar-iii/)|My Calendar III|困难|2456|56%|
|[733](https://leetcode-cn.com/problems/flood-fill/)|Flood Fill|简单|29663|54%|
|[734](https://leetcode-cn.com/problems/sentence-similarity/)|Sentence Similarity|简单|1981|47%|
|[735](https://leetcode-cn.com/problems/asteroid-collision/)|Asteroid Collision|中等|21778|39%|
|[736](https://leetcode-cn.com/problems/parse-lisp-expression/)|Parse Lisp Expression|困难|1864|41%|
|[737](https://leetcode-cn.com/problems/sentence-similarity-ii/)|Sentence Similarity II|中等|4113|44%|
|[738](https://leetcode-cn.com/problems/monotone-increasing-digits/)|Monotone Increasing Digits|中等|11860|43%|
|[739](https://leetcode-cn.com/problems/daily-temperatures/)|[Daily Temperatures](./src/0739-daily-temperatures/)|中等|138518|64%|
|[740](https://leetcode-cn.com/problems/delete-and-earn/)|Delete and Earn|中等|12273|51%|
|[741](https://leetcode-cn.com/problems/cherry-pickup/)|Cherry Pickup|困难|5644|32%|
|[709](https://leetcode-cn.com/problems/to-lower-case/)|To Lower Case|简单|63185|75%|
|[742](https://leetcode-cn.com/problems/closest-leaf-in-a-binary-tree/)|Closest Leaf in a Binary Tree|中等|1557|46%|
|[743](https://leetcode-cn.com/problems/network-delay-time/)|Network Delay Time|中等|30463|45%|
|[744](https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/)|Find Smallest Letter Greater Than Target|简单|44733|44%|
|[745](https://leetcode-cn.com/problems/prefix-and-suffix-search/)|Prefix and Suffix Search|困难|4021|33%|
|[746](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)|[Min Cost Climbing Stairs](./src/0746-min-cost-climbing-stairs/)|简单|76734|49%|
|[747](https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/)|Largest Number At Least Twice of Others|简单|83933|39%|
|[748](https://leetcode-cn.com/problems/shortest-completing-word/)|Shortest Completing Word|简单|12354|58%|
|[749](https://leetcode-cn.com/problems/contain-virus/)|Contain Virus|困难|1367|46%|
|[750](https://leetcode-cn.com/problems/number-of-corner-rectangles/)|Number Of Corner Rectangles|中等|1520|71%|
|[751](https://leetcode-cn.com/problems/ip-to-cidr/)|IP to CIDR|简单|1318|66%|
|[752](https://leetcode-cn.com/problems/open-the-lock/)|Open the Lock|中等|37236|49%|
|[753](https://leetcode-cn.com/problems/cracking-the-safe/)|Cracking the Safe|困难|2119|53%|
|[754](https://leetcode-cn.com/problems/reach-a-number/)|Reach a Number|中等|13005|40%|
|[755](https://leetcode-cn.com/problems/pour-water/)|Pour Water|中等|1415|48%|
|[756](https://leetcode-cn.com/problems/pyramid-transition-matrix/)|Pyramid Transition Matrix|中等|4734|55%|
|[426](https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/)|Convert Binary Search Tree to Sorted Doubly Linked List|中等|4550|63%|
|[757](https://leetcode-cn.com/problems/set-intersection-size-at-least-two/)|Set Intersection Size At Least Two|困难|3254|37%|
|[758](https://leetcode-cn.com/problems/bold-words-in-string/)|Bold Words in String|简单|4893|44%|
|[759](https://leetcode-cn.com/problems/employee-free-time/)|Employee Free Time|困难|1119|58%|
|[760](https://leetcode-cn.com/problems/find-anagram-mappings/)|Find Anagram Mappings|简单|2508|83%|
|[761](https://leetcode-cn.com/problems/special-binary-string/)|Special Binary String|困难|1761|56%|
|[429](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)|N-ary Tree Level Order Traversal|中等|37059|65%|
|[428](https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree/)|Serialize and Deserialize N-ary Tree|困难|1642|61%|
|[430](https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/)|Flatten a Multilevel Doubly Linked List|中等|20570|49%|
|[762](https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/)|Prime Number of Set Bits in Binary Representation|简单|15944|67%|
|[763](https://leetcode-cn.com/problems/partition-labels/)|Partition Labels|中等|23048|71%|
|[764](https://leetcode-cn.com/problems/largest-plus-sign/)|Largest Plus Sign|中等|5013|48%|
|[765](https://leetcode-cn.com/problems/couples-holding-hands/)|Couples Holding Hands|困难|12039|58%|
|[431](https://leetcode-cn.com/problems/encode-n-ary-tree-to-binary-tree/)|Encode N-ary Tree to Binary Tree|困难|541|68%|
|[427](https://leetcode-cn.com/problems/construct-quad-tree/)|Construct Quad Tree|中等|4331|57%|
|[558](https://leetcode-cn.com/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/)|Logical OR of Two Binary Grids Represented as Quad-Trees|中等|4424|47%|
|[559](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/)|Maximum Depth of N-ary Tree|简单|38758|69%|
|[589](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)|N-ary Tree Preorder Traversal|简单|46574|73%|
|[590](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)|N-ary Tree Postorder Traversal|简单|38412|74%|
|[766](https://leetcode-cn.com/problems/toeplitz-matrix/)|Toeplitz Matrix|简单|22409|66%|
|[767](https://leetcode-cn.com/problems/reorganize-string/)|Reorganize String|中等|20873|41%|
|[768](https://leetcode-cn.com/problems/max-chunks-to-make-sorted-ii/)|Max Chunks To Make Sorted II|困难|5089|46%|
|[769](https://leetcode-cn.com/problems/max-chunks-to-make-sorted/)|Max Chunks To Make Sorted|中等|8494|55%|
|[770](https://leetcode-cn.com/problems/basic-calculator-iv/)|Basic Calculator IV|困难|790|49%|
|[771](https://leetcode-cn.com/problems/jewels-and-stones/)|Jewels and Stones|简单|104531|83%|
|[700](https://leetcode-cn.com/problems/search-in-a-binary-search-tree/)|Search in a Binary Search Tree|简单|39331|73%|
|[701](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)|Insert into a Binary Search Tree|中等|20863|72%|
|[772](https://leetcode-cn.com/problems/basic-calculator-iii/)|Basic Calculator III|困难|4280|37%|
|[702](https://leetcode-cn.com/problems/search-in-a-sorted-array-of-unknown-size/)|Search in a Sorted Array of Unknown Size|中等|2505|68%|
|[773](https://leetcode-cn.com/problems/sliding-puzzle/)|Sliding Puzzle|困难|5806|60%|
|[774](https://leetcode-cn.com/problems/minimize-max-distance-to-gas-station/)|Minimize Max Distance to Gas Station|困难|908|57%|
|[703](https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/)|Kth Largest Element in a Stream|简单|51301|44%|
|[775](https://leetcode-cn.com/problems/global-and-local-inversions/)|Global and Local Inversions|中等|6371|44%|
|[776](https://leetcode-cn.com/problems/split-bst/)|Split BST|中等|1769|57%|
|[704](https://leetcode-cn.com/problems/binary-search/)|Binary Search|简单|99962|54%|
|[777](https://leetcode-cn.com/problems/swap-adjacent-in-lr-string/)|Swap Adjacent in LR String|中等|9924|32%|
|[778](https://leetcode-cn.com/problems/swim-in-rising-water/)|Swim in Rising Water|困难|8550|46%|
|[779](https://leetcode-cn.com/problems/k-th-symbol-in-grammar/)|K-th Symbol in Grammar|中等|25992|42%|
|[780](https://leetcode-cn.com/problems/reaching-points/)|Reaching Points|困难|5867|27%|
|[781](https://leetcode-cn.com/problems/rabbits-in-forest/)|Rabbits in Forest|中等|8203|55%|
|[782](https://leetcode-cn.com/problems/transform-to-chessboard/)|Transform to Chessboard|困难|1823|40%|
|[783](https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/)|Minimum Distance Between BST Nodes|简单|27440|53%|
|[784](https://leetcode-cn.com/problems/letter-case-permutation/)|Letter Case Permutation|简单|34727|64%|
|[785](https://leetcode-cn.com/problems/is-graph-bipartite/)|Is Graph Bipartite?|中等|49329|49%|
|[786](https://leetcode-cn.com/problems/k-th-smallest-prime-fraction/)|K-th Smallest Prime Fraction|困难|4426|39%|
|[787](https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/)|Cheapest Flights Within K Stops|中等|20967|35%|
|[788](https://leetcode-cn.com/problems/rotated-digits/)|Rotated Digits|简单|19325|59%|
|[789](https://leetcode-cn.com/problems/escape-the-ghosts/)|Escape The Ghosts|中等|4436|59%|
|[790](https://leetcode-cn.com/problems/domino-and-tromino-tiling/)|Domino and Tromino Tiling|中等|4491|42%|
|[791](https://leetcode-cn.com/problems/custom-sort-string/)|Custom Sort String|中等|8772|67%|
|[792](https://leetcode-cn.com/problems/number-of-matching-subsequences/)|Number of Matching Subsequences|中等|10731|41%|
|[793](https://leetcode-cn.com/problems/preimage-size-of-factorial-zeroes-function/)|Preimage Size of Factorial Zeroes Function|困难|4915|33%|
|[794](https://leetcode-cn.com/problems/valid-tic-tac-toe-state/)|Valid Tic-Tac-Toe State|中等|12858|32%|
|[795](https://leetcode-cn.com/problems/number-of-subarrays-with-bounded-maximum/)|Number of Subarrays with Bounded Maximum|中等|7276|51%|
|[796](https://leetcode-cn.com/problems/rotate-string/)|Rotate String|简单|21997|50%|
|[797](https://leetcode-cn.com/problems/all-paths-from-source-to-target/)|All Paths From Source to Target|中等|5588|76%|
|[798](https://leetcode-cn.com/problems/smallest-rotation-with-highest-score/)|Smallest Rotation with Highest Score|困难|1536|41%|
|[799](https://leetcode-cn.com/problems/champagne-tower/)|Champagne Tower|中等|6280|36%|
|[705](https://leetcode-cn.com/problems/design-hashset/)|Design HashSet|简单|27694|56%|
|[706](https://leetcode-cn.com/problems/design-hashmap/)|Design HashMap|简单|24419|57%|
|[800](https://leetcode-cn.com/problems/similar-rgb-color/)|Similar RGB Color|简单|1068|66%|
|[801](https://leetcode-cn.com/problems/minimum-swaps-to-make-sequences-increasing/)|Minimum Swaps To Make Sequences Increasing|中等|8287|39%|
|[802](https://leetcode-cn.com/problems/find-eventual-safe-states/)|Find Eventual Safe States|中等|9102|46%|
|[803](https://leetcode-cn.com/problems/bricks-falling-when-hit/)|Bricks Falling When Hit|困难|3608|25%|
|[804](https://leetcode-cn.com/problems/unique-morse-code-words/)|Unique Morse Code Words|简单|36619|75%|
|[805](https://leetcode-cn.com/problems/split-array-with-same-average/)|Split Array With Same Average|困难|5099|27%|
|[806](https://leetcode-cn.com/problems/number-of-lines-to-write-string/)|Number of Lines To Write String|简单|12732|65%|
|[807](https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/)|Max Increase to Keep City Skyline|中等|12089|81%|
|[808](https://leetcode-cn.com/problems/soup-servings/)|Soup Servings|中等|5407|45%|
|[809](https://leetcode-cn.com/problems/expressive-words/)|Expressive Words|中等|5733|41%|
|[810](https://leetcode-cn.com/problems/chalkboard-xor-game/)|Chalkboard XOR Game|困难|2018|54%|
|[811](https://leetcode-cn.com/problems/subdomain-visit-count/)|Subdomain Visit Count|简单|14158|67%|
|[812](https://leetcode-cn.com/problems/largest-triangle-area/)|Largest Triangle Area|简单|10868|60%|
|[813](https://leetcode-cn.com/problems/largest-sum-of-averages/)|Largest Sum of Averages|中等|5784|50%|
|[814](https://leetcode-cn.com/problems/binary-tree-pruning/)|Binary Tree Pruning|中等|12704|74%|
|[815](https://leetcode-cn.com/problems/bus-routes/)|Bus Routes|困难|12562|30%|
|[816](https://leetcode-cn.com/problems/ambiguous-coordinates/)|Ambiguous Coordinates|中等|3088|49%|
|[817](https://leetcode-cn.com/problems/linked-list-components/)|Linked List Components|中等|12438|57%|
|[818](https://leetcode-cn.com/problems/race-car/)|Race Car|困难|4238|36%|
|[819](https://leetcode-cn.com/problems/most-common-word/)|Most Common Word|简单|27890|40%|
|[707](https://leetcode-cn.com/problems/design-linked-list/)|Design Linked List|中等|87850|27%|
|[820](https://leetcode-cn.com/problems/short-encoding-of-words/)|Short Encoding of Words|中等|78747|49%|
|[821](https://leetcode-cn.com/problems/shortest-distance-to-a-character/)|Shortest Distance to a Character|简单|19862|67%|
|[822](https://leetcode-cn.com/problems/card-flipping-game/)|Card Flipping Game|中等|3058|47%|
|[823](https://leetcode-cn.com/problems/binary-trees-with-factors/)|Binary Trees With Factors|中等|3469|40%|
|[708](https://leetcode-cn.com/problems/insert-into-a-sorted-circular-linked-list/)|Insert into a Sorted Circular Linked List|中等|2999|30%|
|[824](https://leetcode-cn.com/problems/goat-latin/)|Goat Latin|简单|17212|59%|
|[825](https://leetcode-cn.com/problems/friends-of-appropriate-ages/)|Friends Of Appropriate Ages|中等|9640|35%|
|[826](https://leetcode-cn.com/problems/most-profit-assigning-work/)|Most Profit Assigning Work|中等|11347|35%|
|[827](https://leetcode-cn.com/problems/making-a-large-island/)|Making A Large Island|困难|6209|40%|
|[828](https://leetcode-cn.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/)|Count Unique Characters of All Substrings of a Given String|困难|3590|41%|
|[829](https://leetcode-cn.com/problems/consecutive-numbers-sum/)|Consecutive Numbers Sum|困难|14399|32%|
|[830](https://leetcode-cn.com/problems/positions-of-large-groups/)|Positions of Large Groups|简单|22450|47%|
|[831](https://leetcode-cn.com/problems/masking-personal-information/)|Masking Personal Information|中等|12397|38%|
|[641](https://leetcode-cn.com/problems/design-circular-deque/)|Design Circular Deque|中等|17909|52%|
|[622](https://leetcode-cn.com/problems/design-circular-queue/)|[Design Circular Queue](./src/0622-design-circular-queue/)|中等|84236|40%|
|[832](https://leetcode-cn.com/problems/flipping-an-image/)|Flipping an Image|简单|44620|74%|
|[833](https://leetcode-cn.com/problems/find-and-replace-in-string/)|Find And Replace in String|中等|6902|39%|
|[834](https://leetcode-cn.com/problems/sum-of-distances-in-tree/)|Sum of Distances in Tree|困难|4829|33%|
|[835](https://leetcode-cn.com/problems/image-overlap/)|Image Overlap|中等|4096|58%|
|[489](https://leetcode-cn.com/problems/robot-room-cleaner/)|Robot Room Cleaner|困难|1105|67%|
|[836](https://leetcode-cn.com/problems/rectangle-overlap/)|Rectangle Overlap|简单|58622|50%|
|[837](https://leetcode-cn.com/problems/new-21-game/)|New 21 Game|中等|37587|41%|
|[838](https://leetcode-cn.com/problems/push-dominoes/)|Push Dominoes|中等|7977|45%|
|[839](https://leetcode-cn.com/problems/similar-string-groups/)|Similar String Groups|困难|7974|34%|
|[840](https://leetcode-cn.com/problems/magic-squares-in-grid/)|Magic Squares In Grid|简单|17885|34%|
|[841](https://leetcode-cn.com/problems/keys-and-rooms/)|Keys and Rooms|中等|24844|61%|
|[842](https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/)|Split Array into Fibonacci Sequence|中等|12763|37%|
|[843](https://leetcode-cn.com/problems/guess-the-word/)|Guess the Word|困难|3287|35%|
|[844](https://leetcode-cn.com/problems/backspace-string-compare/)|Backspace String Compare|简单|44751|50%|
|[845](https://leetcode-cn.com/problems/longest-mountain-in-array/)|Longest Mountain in Array|中等|16118|35%|
|[846](https://leetcode-cn.com/problems/hand-of-straights/)|Hand of Straights|中等|9770|48%|
|[847](https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/)|Shortest Path Visiting All Nodes|困难|5047|52%|
|[848](https://leetcode-cn.com/problems/shifting-letters/)|Shifting Letters|中等|11886|40%|
|[849](https://leetcode-cn.com/problems/maximize-distance-to-closest-person/)|Maximize Distance to Closest Person|简单|26162|40%|
|[850](https://leetcode-cn.com/problems/rectangle-area-ii/)|Rectangle Area II|困难|2549|41%|
|[851](https://leetcode-cn.com/problems/loud-and-rich/)|Loud and Rich|中等|4597|43%|
|[852](https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/)|Peak Index in a Mountain Array|简单|37570|70%|
|[853](https://leetcode-cn.com/problems/car-fleet/)|Car Fleet|中等|13885|35%|
|[854](https://leetcode-cn.com/problems/k-similar-strings/)|K-Similar Strings|困难|7027|31%|
|[855](https://leetcode-cn.com/problems/exam-room/)|Exam Room|中等|4789|35%|
|[856](https://leetcode-cn.com/problems/score-of-parentheses/)|Score of Parentheses|中等|12841|59%|
|[857](https://leetcode-cn.com/problems/minimum-cost-to-hire-k-workers/)|Minimum Cost to Hire K Workers|困难|3155|44%|
|[858](https://leetcode-cn.com/problems/mirror-reflection/)|Mirror Reflection|中等|3008|53%|
|[859](https://leetcode-cn.com/problems/buddy-strings/)|Buddy Strings|简单|50786|29%|
|[860](https://leetcode-cn.com/problems/lemonade-change/)|Lemonade Change|简单|46740|55%|
|[861](https://leetcode-cn.com/problems/score-after-flipping-matrix/)|Score After Flipping Matrix|中等|6976|74%|
|[862](https://leetcode-cn.com/problems/shortest-subarray-with-sum-at-least-k/)|Shortest Subarray with Sum at Least K|困难|53271|15%|
|[863](https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/)|All Nodes Distance K in Binary Tree|中等|11970|50%|
|[710](https://leetcode-cn.com/problems/random-pick-with-blacklist/)|Random Pick with Blacklist|困难|4781|30%|
|[864](https://leetcode-cn.com/problems/shortest-path-to-get-all-keys/)|Shortest Path to Get All Keys|困难|3340|40%|
|[865](https://leetcode-cn.com/problems/smallest-subtree-with-all-the-deepest-nodes/)|Smallest Subtree with all the Deepest Nodes|中等|5470|56%|
|[866](https://leetcode-cn.com/problems/prime-palindrome/)|Prime Palindrome|中等|19632|20%|
|[867](https://leetcode-cn.com/problems/transpose-matrix/)|Transpose Matrix|简单|30078|68%|
|[868](https://leetcode-cn.com/problems/binary-gap/)|Binary Gap|简单|16315|60%|
|[869](https://leetcode-cn.com/problems/reordered-power-of-2/)|Reordered Power of 2|中等|6671|51%|
|[870](https://leetcode-cn.com/problems/advantage-shuffle/)|Advantage Shuffle|中等|19122|38%|
|[871](https://leetcode-cn.com/problems/minimum-number-of-refueling-stops/)|Minimum Number of Refueling Stops|困难|13503|30%|
|[470](https://leetcode-cn.com/problems/implement-rand10-using-rand7/)|Implement Rand10() Using Rand7()|中等|16830|49%|
|[872](https://leetcode-cn.com/problems/leaf-similar-trees/)|Leaf-Similar Trees|简单|22096|62%|
|[873](https://leetcode-cn.com/problems/length-of-longest-fibonacci-subsequence/)|Length of Longest Fibonacci Subsequence|中等|14066|46%|
|[874](https://leetcode-cn.com/problems/walking-robot-simulation/)|Walking Robot Simulation|简单|25806|36%|
|[875](https://leetcode-cn.com/problems/koko-eating-bananas/)|Koko Eating Bananas|中等|34997|43%|
|[876](https://leetcode-cn.com/problems/middle-of-the-linked-list/)|Middle of the Linked List|简单|97181|69%|
|[877](https://leetcode-cn.com/problems/stone-game/)|Stone Game|中等|23947|69%|
|[878](https://leetcode-cn.com/problems/nth-magical-number/)|Nth Magical Number|困难|9671|23%|
|[879](https://leetcode-cn.com/problems/profitable-schemes/)|Profitable Schemes|困难|3183|35%|
|[528](https://leetcode-cn.com/problems/random-pick-with-weight/)|Random Pick with Weight|中等|6637|44%|
|[519](https://leetcode-cn.com/problems/random-flip-matrix/)|Random Flip Matrix|中等|3327|36%|
|[497](https://leetcode-cn.com/problems/random-point-in-non-overlapping-rectangles/)|Random Point in Non-overlapping Rectangles|中等|3373|39%|
|[478](https://leetcode-cn.com/problems/generate-random-point-in-a-circle/)|Generate Random Point in a Circle|中等|6891|41%|
|[880](https://leetcode-cn.com/problems/decoded-string-at-index/)|Decoded String at Index|中等|14138|23%|
|[881](https://leetcode-cn.com/problems/boats-to-save-people/)|Boats to Save People|中等|20620|43%|
|[882](https://leetcode-cn.com/problems/reachable-nodes-in-subdivided-graph/)|Reachable Nodes In Subdivided Graph|困难|1801|41%|
|[883](https://leetcode-cn.com/problems/projection-area-of-3d-shapes/)|Projection Area of 3D Shapes|简单|10357|65%|
|[884](https://leetcode-cn.com/problems/uncommon-words-from-two-sentences/)|Uncommon Words from Two Sentences|简单|17496|62%|
|[885](https://leetcode-cn.com/problems/spiral-matrix-iii/)|Spiral Matrix III|中等|3473|66%|
|[886](https://leetcode-cn.com/problems/possible-bipartition/)|Possible Bipartition|中等|8626|38%|
|[887](https://leetcode-cn.com/problems/super-egg-drop/)|Super Egg Drop|困难|93371|28%|
|[888](https://leetcode-cn.com/problems/fair-candy-swap/)|Fair Candy Swap|简单|23045|54%|
|[889](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/)|Construct Binary Tree from Preorder and Postorder Traversal|中等|8614|65%|
|[890](https://leetcode-cn.com/problems/find-and-replace-pattern/)|Find and Replace Pattern|中等|7442|70%|
|[891](https://leetcode-cn.com/problems/sum-of-subsequence-widths/)|Sum of Subsequence Widths|困难|5219|29%|
|[892](https://leetcode-cn.com/problems/surface-area-of-3d-shapes/)|Surface Area of 3D Shapes|简单|44481|64%|
|[893](https://leetcode-cn.com/problems/groups-of-special-equivalent-strings/)|Groups of Special-Equivalent Strings|简单|8867|69%|
|[894](https://leetcode-cn.com/problems/all-possible-full-binary-trees/)|All Possible Full Binary Trees|中等|9315|75%|
|[895](https://leetcode-cn.com/problems/maximum-frequency-stack/)|Maximum Frequency Stack|困难|5389|46%|
|[896](https://leetcode-cn.com/problems/monotonic-array/)|Monotonic Array|简单|31249|52%|
|[897](https://leetcode-cn.com/problems/increasing-order-search-tree/)|Increasing Order Search Tree|简单|17218|69%|
|[898](https://leetcode-cn.com/problems/bitwise-ors-of-subarrays/)|Bitwise ORs of Subarrays|中等|10247|28%|
|[899](https://leetcode-cn.com/problems/orderly-queue/)|Orderly Queue|困难|4900|51%|
|[900](https://leetcode-cn.com/problems/rle-iterator/)|RLE Iterator|中等|3734|45%|
|[901](https://leetcode-cn.com/problems/online-stock-span/)|Online Stock Span|中等|13669|45%|
|[902](https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set/)|Numbers At Most N Given Digit Set|困难|6029|29%|
|[903](https://leetcode-cn.com/problems/valid-permutations-for-di-sequence/)|Valid Permutations for DI Sequence|困难|3719|50%|
|[904](https://leetcode-cn.com/problems/fruit-into-baskets/)|Fruit Into Baskets|中等|14504|41%|
|[905](https://leetcode-cn.com/problems/sort-array-by-parity/)|Sort Array By Parity|简单|54880|69%|
|[906](https://leetcode-cn.com/problems/super-palindromes/)|Super Palindromes|困难|5068|23%|
|[907](https://leetcode-cn.com/problems/sum-of-subarray-minimums/)|Sum of Subarray Minimums|中等|16765|29%|
|[908](https://leetcode-cn.com/problems/smallest-range-i/)|Smallest Range I|简单|17907|68%|
|[909](https://leetcode-cn.com/problems/snakes-and-ladders/)|Snakes and Ladders|中等|5966|31%|
|[910](https://leetcode-cn.com/problems/smallest-range-ii/)|Smallest Range II|中等|12573|29%|
|[911](https://leetcode-cn.com/problems/online-election/)|Online Election|中等|4609|41%|
|[912](https://leetcode-cn.com/problems/sort-an-array/)|Sort an Array|中等|121460|59%|
|[913](https://leetcode-cn.com/problems/cat-and-mouse/)|Cat and Mouse|困难|2399|34%|
|[914](https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/)|X of a Kind in a Deck of Cards|简单|100255|39%|
|[915](https://leetcode-cn.com/problems/partition-array-into-disjoint-intervals/)|Partition Array into Disjoint Intervals|中等|11371|44%|
|[916](https://leetcode-cn.com/problems/word-subsets/)|Word Subsets|中等|9241|41%|
|[917](https://leetcode-cn.com/problems/reverse-only-letters/)|Reverse Only Letters|简单|25990|55%|
|[918](https://leetcode-cn.com/problems/maximum-sum-circular-subarray/)|Maximum Sum Circular Subarray|中等|12381|32%|
|[919](https://leetcode-cn.com/problems/complete-binary-tree-inserter/)|Complete Binary Tree Inserter|中等|4089|60%|
|[920](https://leetcode-cn.com/problems/number-of-music-playlists/)|Number of Music Playlists|困难|2142|43%|
|[921](https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/)|Minimum Add to Make Parentheses Valid|中等|13699|72%|
|[922](https://leetcode-cn.com/problems/sort-array-by-parity-ii/)|Sort Array By Parity II|简单|48044|67%|
|[923](https://leetcode-cn.com/problems/3sum-with-multiplicity/)|3Sum With Multiplicity|中等|9847|31%|
|[924](https://leetcode-cn.com/problems/minimize-malware-spread/)|Minimize Malware Spread|困难|7578|36%|
|[925](https://leetcode-cn.com/problems/long-pressed-name/)|Long Pressed Name|简单|24425|41%|
|[926](https://leetcode-cn.com/problems/flip-string-to-monotone-increasing/)|Flip String to Monotone Increasing|中等|6957|47%|
|[927](https://leetcode-cn.com/problems/three-equal-parts/)|[Three Equal Parts](./src/0927-three-equal-parts/)|困难|5326|30%|
|[928](https://leetcode-cn.com/problems/minimize-malware-spread-ii/)|Minimize Malware Spread II|困难|2870|39%|
|[929](https://leetcode-cn.com/problems/unique-email-addresses/)|Unique Email Addresses|简单|28061|63%|
|[930](https://leetcode-cn.com/problems/binary-subarrays-with-sum/)|Binary Subarrays With Sum|中等|9598|36%|
|[931](https://leetcode-cn.com/problems/minimum-falling-path-sum/)|Minimum Falling Path Sum|中等|10643|60%|
|[932](https://leetcode-cn.com/problems/beautiful-array/)|Beautiful Array|中等|7313|59%|
|[933](https://leetcode-cn.com/problems/number-of-recent-calls/)|Number of Recent Calls|简单|21004|71%|
|[934](https://leetcode-cn.com/problems/shortest-bridge/)|Shortest Bridge|中等|16795|44%|
|[935](https://leetcode-cn.com/problems/knight-dialer/)|Knight Dialer|中等|6575|45%|
|[936](https://leetcode-cn.com/problems/stamping-the-sequence/)|Stamping The Sequence|困难|2533|34%|
|[937](https://leetcode-cn.com/problems/reorder-data-in-log-files/)|Reorder Data in Log Files|简单|10485|56%|
|[938](https://leetcode-cn.com/problems/range-sum-of-bst/)|Range Sum of BST|简单|36029|76%|
|[939](https://leetcode-cn.com/problems/minimum-area-rectangle/)|Minimum Area Rectangle|中等|7041|41%|
|[940](https://leetcode-cn.com/problems/distinct-subsequences-ii/)|Distinct Subsequences II|困难|3752|41%|
|[941](https://leetcode-cn.com/problems/valid-mountain-array/)|Valid Mountain Array|简单|31295|35%|
|[942](https://leetcode-cn.com/problems/di-string-match/)|DI String Match|简单|21127|71%|
|[943](https://leetcode-cn.com/problems/find-the-shortest-superstring/)|Find the Shortest Superstring|困难|2099|42%|
|[944](https://leetcode-cn.com/problems/delete-columns-to-make-sorted/)|Delete Columns to Make Sorted|简单|18247|69%|
|[945](https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/)|Minimum Increment to Make Array Unique|中等|60597|47%|
|[946](https://leetcode-cn.com/problems/validate-stack-sequences/)|Validate Stack Sequences|中等|20448|58%|
|[947](https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column/)|Most Stones Removed with Same Row or Column|中等|5445|51%|
|[948](https://leetcode-cn.com/problems/bag-of-tokens/)|Bag of Tokens|中等|9842|37%|
|[949](https://leetcode-cn.com/problems/largest-time-for-given-digits/)|Largest Time for Given Digits|简单|15261|35%|
|[950](https://leetcode-cn.com/problems/reveal-cards-in-increasing-order/)|Reveal Cards In Increasing Order|中等|8488|76%|
|[951](https://leetcode-cn.com/problems/flip-equivalent-binary-trees/)|Flip Equivalent Binary Trees|中等|7265|64%|
|[952](https://leetcode-cn.com/problems/largest-component-size-by-common-factor/)|Largest Component Size by Common Factor|困难|4364|30%|
|[953](https://leetcode-cn.com/problems/verifying-an-alien-dictionary/)|Verifying an Alien Dictionary|简单|11899|56%|
|[954](https://leetcode-cn.com/problems/array-of-doubled-pairs/)|Array of Doubled Pairs|中等|12315|27%|
|[955](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-ii/)|Delete Columns to Make Sorted II|中等|5791|29%|
|[956](https://leetcode-cn.com/problems/tallest-billboard/)|Tallest Billboard|困难|3970|38%|
|[957](https://leetcode-cn.com/problems/prison-cells-after-n-days/)|Prison Cells After N Days|中等|26218|34%|
|[958](https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/)|Check Completeness of a Binary Tree|中等|13575|49%|
|[959](https://leetcode-cn.com/problems/regions-cut-by-slashes/)|Regions Cut By Slashes|中等|4481|66%|
|[960](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-iii/)|Delete Columns to Make Sorted III|困难|2241|53%|
|[961](https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array/)|N-Repeated Element in Size 2N Array|简单|35843|66%|
|[962](https://leetcode-cn.com/problems/maximum-width-ramp/)|Maximum Width Ramp|中等|14726|38%|
|[963](https://leetcode-cn.com/problems/minimum-area-rectangle-ii/)|Minimum Area Rectangle II|中等|3116|47%|
|[964](https://leetcode-cn.com/problems/least-operators-to-express-number/)|Least Operators to Express Number|困难|1777|39%|
|[965](https://leetcode-cn.com/problems/univalued-binary-tree/)|Univalued Binary Tree|简单|24943|67%|
|[966](https://leetcode-cn.com/problems/vowel-spellchecker/)|Vowel Spellchecker|中等|4537|38%|
|[967](https://leetcode-cn.com/problems/numbers-with-same-consecutive-differences/)|Numbers With Same Consecutive Differences|中等|8229|36%|
|[968](https://leetcode-cn.com/problems/binary-tree-cameras/)|Binary Tree Cameras|困难|7563|38%|
|[969](https://leetcode-cn.com/problems/pancake-sorting/)|Pancake Sorting|中等|9955|64%|
|[970](https://leetcode-cn.com/problems/powerful-integers/)|Powerful Integers|简单|17443|39%|
|[971](https://leetcode-cn.com/problems/flip-binary-tree-to-match-preorder-traversal/)|Flip Binary Tree To Match Preorder Traversal|中等|4721|45%|
|[972](https://leetcode-cn.com/problems/equal-rational-numbers/)|Equal Rational Numbers|困难|1901|40%|
|[509](https://leetcode-cn.com/problems/fibonacci-number/)|Fibonacci Number|简单|96819|66%|
|[973](https://leetcode-cn.com/problems/k-closest-points-to-origin/)|K Closest Points to Origin|中等|30840|58%|
|[974](https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/)|Subarray Sums Divisible by K|中等|49269|45%|
|[975](https://leetcode-cn.com/problems/odd-even-jump/)|Odd Even Jump|困难|3603|41%|
|[976](https://leetcode-cn.com/problems/largest-perimeter-triangle/)|[Largest Perimeter Triangle](./src/0976-largest-perimeter-triangle/)|简单|31885|55%|
|[977](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)|Squares of a Sorted Array|简单|56043|71%|
|[978](https://leetcode-cn.com/problems/longest-turbulent-subarray/)|[Longest Turbulent Subarray](./src/0978-longest-turbulent-subarray/)|中等|13090|41%|
|[979](https://leetcode-cn.com/problems/distribute-coins-in-binary-tree/)|Distribute Coins in Binary Tree|中等|5756|67%|
|[980](https://leetcode-cn.com/problems/unique-paths-iii/)|Unique Paths III|困难|9362|71%|
|[981](https://leetcode-cn.com/problems/time-based-key-value-store/)|[Time Based Key-Value Store](./src/0981-time-based-key-value-store/)|中等|4963|38%|
|[982](https://leetcode-cn.com/problems/triples-with-bitwise-and-equal-to-zero/)|[Triples with Bitwise AND Equal To Zero](./src/0982-triples-with-bitwise-and-equal-to-zero/)|困难|2707|49%|
|[983](https://leetcode-cn.com/problems/minimum-cost-for-tickets/)|[Minimum Cost For Tickets](./src/0983-minimum-cost-for-tickets/)|中等|38126|63%|
|[984](https://leetcode-cn.com/problems/string-without-aaa-or-bbb/)|String Without AAA or BBB|中等|12812|38%|
|[985](https://leetcode-cn.com/problems/sum-of-even-numbers-after-queries/)|Sum of Even Numbers After Queries|简单|17068|59%|
|[986](https://leetcode-cn.com/problems/interval-list-intersections/)|Interval List Intersections|中等|11235|63%|
|[987](https://leetcode-cn.com/problems/vertical-order-traversal-of-a-binary-tree/)|Vertical Order Traversal of a Binary Tree|中等|8787|40%|
|[988](https://leetcode-cn.com/problems/smallest-string-starting-from-leaf/)|Smallest String Starting From Leaf|中等|9526|46%|
|[989](https://leetcode-cn.com/problems/add-to-array-form-of-integer/)|Add to Array-Form of Integer|简单|32028|44%|
|[990](https://leetcode-cn.com/problems/satisfiability-of-equality-equations/)|Satisfiability of Equality Equations|中等|43522|45%|
|[991](https://leetcode-cn.com/problems/broken-calculator/)|Broken Calculator|中等|10208|50%|
|[992](https://leetcode-cn.com/problems/subarrays-with-k-different-integers/)|Subarrays with K Different Integers|困难|14769|29%|
|[993](https://leetcode-cn.com/problems/cousins-in-binary-tree/)|Cousins in Binary Tree|简单|19943|51%|
|[994](https://leetcode-cn.com/problems/rotting-oranges/)|Rotting Oranges|中等|60444|50%|
|[995](https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/)|Minimum Number of K Consecutive Bit Flips|困难|4247|41%|
|[996](https://leetcode-cn.com/problems/number-of-squareful-arrays/)|Number of Squareful Arrays|困难|4377|45%|
|[997](https://leetcode-cn.com/problems/find-the-town-judge/)|Find the Town Judge|简单|36169|50%|
|[998](https://leetcode-cn.com/problems/maximum-binary-tree-ii/)|Maximum Binary Tree II|中等|4537|60%|
|[999](https://leetcode-cn.com/problems/available-captures-for-rook/)|Available Captures for Rook|简单|36654|69%|
|[1000](https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/)|Minimum Cost to Merge Stones|困难|4483|33%|
|[1001](https://leetcode-cn.com/problems/grid-illumination/)|Grid Illumination|困难|4114|28%|
|[1002](https://leetcode-cn.com/problems/find-common-characters/)|[Find Common Characters](./src/1002-find-common-characters/)|简单|20219|68%|
|[1003](https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions/)|Check If Word Is Valid After Substitutions|中等|10711|57%|
|[1004](https://leetcode-cn.com/problems/max-consecutive-ones-iii/)|Max Consecutive Ones III|中等|19215|52%|
|[1005](https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/)|Maximize Sum Of Array After K Negations|简单|17951|51%|
|[1006](https://leetcode-cn.com/problems/clumsy-factorial/)|Clumsy Factorial|中等|6501|50%|
|[1007](https://leetcode-cn.com/problems/minimum-domino-rotations-for-equal-row/)|Minimum Domino Rotations For Equal Row|中等|8050|43%|
|[1008](https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal/)|Construct Binary Search Tree from Preorder Traversal|中等|10174|71%|
|[1055](https://leetcode-cn.com/problems/shortest-way-to-form-string/)|Shortest Way to Form String|中等|3035|58%|
|[1057](https://leetcode-cn.com/problems/campus-bikes/)|Campus Bikes|中等|4544|40%|
|[1058](https://leetcode-cn.com/problems/minimize-rounding-error-to-meet-target/)|Minimize Rounding Error to Meet Target|中等|2755|30%|
|[1009](https://leetcode-cn.com/problems/complement-of-base-10-integer/)|[Complement of Base 10 Integer](./src/1009-complement-of-base-10-integer/)|简单|14139|59%|
|[1010](https://leetcode-cn.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/)|Pairs of Songs With Total Durations Divisible by 60|简单|30201|43%|
|[1011](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/)|[Capacity To Ship Packages Within D Days](./src/1011-capacity-to-ship-packages-within-d-days/)|中等|17278|53%|
|[1012](https://leetcode-cn.com/problems/numbers-with-repeated-digits/)|Numbers With Repeated Digits|困难|6384|30%|
|[1061](https://leetcode-cn.com/problems/lexicographically-smallest-equivalent-string/)|Lexicographically Smallest Equivalent String|中等|2785|54%|
|[1060](https://leetcode-cn.com/problems/missing-element-in-sorted-array/)|Missing Element in Sorted Array|中等|4005|49%|
|[1062](https://leetcode-cn.com/problems/longest-repeating-substring/)|Longest Repeating Substring|中等|3792|47%|
|[1063](https://leetcode-cn.com/problems/number-of-valid-subarrays/)|Number of Valid Subarrays|困难|1571|65%|
|[1013](https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/)|[Partition Array Into Three Parts With Equal Sum](./src/1013-partition-array-into-three-parts-with-equal-sum/)|简单|88730|40%|
|[1014](https://leetcode-cn.com/problems/best-sightseeing-pair/)|[Best Sightseeing Pair](./src/1014-best-sightseeing-pair/)|中等|49894|53%|
|[1015](https://leetcode-cn.com/problems/smallest-integer-divisible-by-k/)|Smallest Integer Divisible by K|中等|8870|32%|
|[1016](https://leetcode-cn.com/problems/binary-string-with-substrings-representing-1-to-n/)|Binary String With Substrings Representing 1 To N|中等|5613|57%|
|[1064](https://leetcode-cn.com/problems/fixed-point/)|Fixed Point|简单|4586|66%|
|[1066](https://leetcode-cn.com/problems/campus-bikes-ii/)|Campus Bikes II|中等|2483|42%|
|[1067](https://leetcode-cn.com/problems/digit-count-in-range/)|Digit Count in Range|困难|1196|38%|
|[1056](https://leetcode-cn.com/problems/confusing-number/)|Confusing Number|简单|6000|36%|
|[1017](https://leetcode-cn.com/problems/convert-to-base-2/)|[Convert to Base -2](./src/1017-convert-to-base-2/)|中等|3913|53%|
|[1018](https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/)|[Binary Prefix Divisible By 5](./src/1018-binary-prefix-divisible-by-5/)|简单|20775|46%|
|[1019](https://leetcode-cn.com/problems/next-greater-node-in-linked-list/)|[Next Greater Node In Linked List](./src/1019-next-greater-node-in-linked-list/)|中等|17956|55%|
|[1020](https://leetcode-cn.com/problems/number-of-enclaves/)|[Number of Enclaves](./src/1020-number-of-enclaves/)|中等|9024|51%|
|[1086](https://leetcode-cn.com/problems/high-five/)|High Five|简单|4477|70%|
|[1065](https://leetcode-cn.com/problems/index-pairs-of-a-string/)|Index Pairs of a String|简单|2497|51%|
|[1087](https://leetcode-cn.com/problems/brace-expansion/)|Brace Expansion|中等|2149|48%|
|[1088](https://leetcode-cn.com/problems/confusing-number-ii/)|Confusing Number II|困难|1386|38%|
|[1021](https://leetcode-cn.com/problems/remove-outermost-parentheses/)|[Remove Outermost Parentheses](./src/1021-remove-outermost-parentheses/)|简单|40548|77%|
|[1022](https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/)|[Sum of Root To Leaf Binary Numbers](./src/1022-sum-of-root-to-leaf-binary-numbers/)|简单|12826|64%|
|[1023](https://leetcode-cn.com/problems/camelcase-matching/)|Camelcase Matching|中等|6840|53%|
|[1024](https://leetcode-cn.com/problems/video-stitching/)|Video Stitching|中等|10170|47%|
|[1085](https://leetcode-cn.com/problems/sum-of-digits-in-the-minimum-number/)|Sum of Digits in the Minimum Number|简单|2893|75%|
|[1099](https://leetcode-cn.com/problems/two-sum-less-than-k/)|Two Sum Less Than K|简单|4006|56%|
|[1100](https://leetcode-cn.com/problems/find-k-length-substrings-with-no-repeated-characters/)|Find K-Length Substrings With No Repeated Characters|中等|3438|69%|
|[1101](https://leetcode-cn.com/problems/the-earliest-moment-when-everyone-become-friends/)|The Earliest Moment When Everyone Become Friends|中等|2176|65%|
|[1025](https://leetcode-cn.com/problems/divisor-game/)|Divisor Game|简单|37574|67%|
|[1027](https://leetcode-cn.com/problems/longest-arithmetic-sequence/)|Longest Arithmetic Sequence|中等|14030|44%|
|[1118](https://leetcode-cn.com/problems/number-of-days-in-a-month/)|Number of Days in a Month|简单|2012|62%|
|[1119](https://leetcode-cn.com/problems/remove-vowels-from-a-string/)|Remove Vowels from a String|简单|5165|86%|
|[1134](https://leetcode-cn.com/problems/armstrong-number/)|Armstrong Number|简单|2481|75%|
|[1120](https://leetcode-cn.com/problems/maximum-average-subtree/)|Maximum Average Subtree|中等|2250|57%|
|[1026](https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor/)|Maximum Difference Between Node and Ancestor|中等|8039|62%|
|[1028](https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/)|Recover a Tree From Preorder Traversal|困难|21654|74%|
|[1030](https://leetcode-cn.com/problems/matrix-cells-in-distance-order/)|Matrix Cells in Distance Order|简单|12649|64%|
|[1029](https://leetcode-cn.com/problems/two-city-scheduling/)|Two City Scheduling|简单|16861|62%|
|[1031](https://leetcode-cn.com/problems/maximum-sum-of-two-non-overlapping-subarrays/)|Maximum Sum of Two Non-Overlapping Subarrays|中等|4624|52%|
|[1032](https://leetcode-cn.com/problems/stream-of-characters/)|Stream of Characters|困难|3724|34%|
|[1133](https://leetcode-cn.com/problems/largest-unique-number/)|Largest Unique Number|简单|2597|62%|
|[1102](https://leetcode-cn.com/problems/path-with-maximum-minimum-value/)|Path With Maximum Minimum Value|中等|7071|31%|
|[1135](https://leetcode-cn.com/problems/connecting-cities-with-minimum-cost/)|Connecting Cities With Minimum Cost|中等|5168|47%|
|[1136](https://leetcode-cn.com/problems/parallel-courses/)|Parallel Courses|困难|1615|53%|
|[1150](https://leetcode-cn.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array/)|Check If a Number Is Majority Element in a Sorted Array|简单|3966|63%|
|[1033](https://leetcode-cn.com/problems/moving-stones-until-consecutive/)|Moving Stones Until Consecutive|简单|15961|37%|
|[1034](https://leetcode-cn.com/problems/coloring-a-border/)|Coloring A Border|中等|4983|42%|
|[1035](https://leetcode-cn.com/problems/uncrossed-lines/)|Uncrossed Lines|中等|7552|52%|
|[1036](https://leetcode-cn.com/problems/escape-a-large-maze/)|Escape a Large Maze|困难|6370|27%|
|[1151](https://leetcode-cn.com/problems/minimum-swaps-to-group-all-1s-together/)|Minimum Swaps to Group All 1&#39;s Together|中等|4455|46%|
|[1152](https://leetcode-cn.com/problems/analyze-user-website-visit-pattern/)|Analyze User Website Visit Pattern|中等|2031|34%|
|[1039](https://leetcode-cn.com/problems/minimum-score-triangulation-of-polygon/)|Minimum Score Triangulation of Polygon|中等|4537|52%|
|[1160](https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/)|Find Words That Can Be Formed by Characters|简单|64074|68%|
|[1040](https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/)|Moving Stones Until Consecutive II|中等|3114|48%|
|[1038](https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/)|Binary Search Tree to Greater Sum Tree|中等|13547|75%|
|[1037](https://leetcode-cn.com/problems/valid-boomerang/)|Valid Boomerang|简单|12110|42%|
|[1161](https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree/)|Maximum Level Sum of a Binary Tree|中等|7941|68%|
|[1162](https://leetcode-cn.com/problems/as-far-from-land-as-possible/)|As Far from Land as Possible|中等|57904|46%|
|[1121](https://leetcode-cn.com/problems/divide-array-into-increasing-sequences/)|Divide Array Into Increasing Sequences|困难|1379|53%|
|[1041](https://leetcode-cn.com/problems/robot-bounded-in-circle/)|Robot Bounded In Circle|中等|8522|44%|
|[1042](https://leetcode-cn.com/problems/flower-planting-with-no-adjacent/)|Flower Planting With No Adjacent|简单|12437|50%|
|[1043](https://leetcode-cn.com/problems/partition-array-for-maximum-sum/)|Partition Array for Maximum Sum|中等|4564|64%|
|[1044](https://leetcode-cn.com/problems/longest-duplicate-substring/)|Longest Duplicate Substring|困难|11740|19%|
|[1165](https://leetcode-cn.com/problems/single-row-keyboard/)|Single-Row Keyboard|简单|5821|83%|
|[1153](https://leetcode-cn.com/problems/string-transforms-into-another-string/)|String Transforms Into Another String|困难|2845|33%|
|[1166](https://leetcode-cn.com/problems/design-file-system/)|Design File System|中等|2521|48%|
|[1167](https://leetcode-cn.com/problems/minimum-cost-to-connect-sticks/)|Minimum Cost to Connect Sticks|中等|3583|35%|
|[1046](https://leetcode-cn.com/problems/last-stone-weight/)|Last Stone Weight|简单|29117|60%|
|[1047](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/)|Remove All Adjacent Duplicates In String|简单|30931|67%|
|[1048](https://leetcode-cn.com/problems/longest-string-chain/)|Longest String Chain|中等|10876|42%|
|[1049](https://leetcode-cn.com/problems/last-stone-weight-ii/)|Last Stone Weight II|中等|9101|45%|
|[1180](https://leetcode-cn.com/problems/count-substrings-with-only-one-distinct-letter/)|Count Substrings with Only One Distinct Letter|简单|2408|74%|
|[1181](https://leetcode-cn.com/problems/before-and-after-puzzle/)|Before and After Puzzle|中等|3178|35%|
|[1163](https://leetcode-cn.com/problems/last-substring-in-lexicographical-order/)|Last Substring in Lexicographical Order|困难|13447|24%|
|[1182](https://leetcode-cn.com/problems/shortest-distance-to-target-color/)|Shortest Distance to Target Color|中等|3381|34%|
|[1045](https://leetcode-cn.com/problems/customers-who-bought-all-products/)|Customers Who Bought All Products|中等|4074|65%|
|[1050](https://leetcode-cn.com/problems/actors-and-directors-who-cooperated-at-least-three-times/)|Actors and Directors Who Cooperated At Least Three Times|简单|5220|74%|
|[1051](https://leetcode-cn.com/problems/height-checker/)|Height Checker|简单|26746|73%|
|[1052](https://leetcode-cn.com/problems/grumpy-bookstore-owner/)|Grumpy Bookstore Owner|中等|11434|49%|
|[1053](https://leetcode-cn.com/problems/previous-permutation-with-one-swap/)|Previous Permutation With One Swap|中等|8144|43%|
|[1054](https://leetcode-cn.com/problems/distant-barcodes/)|Distant Barcodes|中等|12940|34%|
|[1196](https://leetcode-cn.com/problems/how-many-apples-can-you-put-into-the-basket/)|How Many Apples Can You Put into the Basket|简单|3203|69%|
|[1197](https://leetcode-cn.com/problems/minimum-knight-moves/)|Minimum Knight Moves|中等|6905|32%|
|[1198](https://leetcode-cn.com/problems/find-smallest-common-element-in-all-rows/)|Find Smallest Common Element in All Rows|中等|3065|73%|
|[1168](https://leetcode-cn.com/problems/optimize-water-distribution-in-a-village/)|Optimize Water Distribution in a Village|困难|1622|46%|
|[1074](https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/)|Number of Submatrices That Sum to Target|困难|2992|48%|
|[1071](https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/)|Greatest Common Divisor of Strings|简单|46970|58%|
|[1072](https://leetcode-cn.com/problems/flip-columns-for-maximum-number-of-equal-rows/)|Flip Columns For Maximum Number of Equal Rows|中等|3524|52%|
|[1073](https://leetcode-cn.com/problems/adding-two-negabinary-numbers/)|Adding Two Negabinary Numbers|中等|4432|31%|
|[1213](https://leetcode-cn.com/problems/intersection-of-three-sorted-arrays/)|Intersection of Three Sorted Arrays|简单|4328|75%|
|[1214](https://leetcode-cn.com/problems/two-sum-bsts/)|Two Sum BSTs|中等|3273|59%|
|[1215](https://leetcode-cn.com/problems/stepping-numbers/)|Stepping Numbers|中等|3451|37%|
|[1183](https://leetcode-cn.com/problems/maximum-number-of-ones/)|Maximum Number of Ones|困难|722|52%|
|[1068](https://leetcode-cn.com/problems/product-sales-analysis-i/)|Product Sales Analysis I|简单|4965|84%|
|[1069](https://leetcode-cn.com/problems/product-sales-analysis-ii/)|Product Sales Analysis II|简单|4597|82%|
|[1070](https://leetcode-cn.com/problems/product-sales-analysis-iii/)|Product Sales Analysis III|中等|4612|47%|
|[1078](https://leetcode-cn.com/problems/occurrences-after-bigram/)|Occurrences After Bigram|简单|11297|61%|
|[1080](https://leetcode-cn.com/problems/insufficient-nodes-in-root-to-leaf-paths/)|Insufficient Nodes in Root to Leaf Paths|中等|4909|45%|
|[1081](https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters/)|Smallest Subsequence of Distinct Characters|中等|6438|49%|
|[1079](https://leetcode-cn.com/problems/letter-tile-possibilities/)|Letter Tile Possibilities|中等|7538|72%|
|[1075](https://leetcode-cn.com/problems/project-employees-i/)|Project Employees I|简单|4353|64%|
|[1076](https://leetcode-cn.com/problems/project-employees-ii/)|Project Employees II|简单|6025|50%|
|[1077](https://leetcode-cn.com/problems/project-employees-iii/)|Project Employees III|中等|3219|69%|
|[1228](https://leetcode-cn.com/problems/missing-number-in-arithmetic-progression/)|Missing Number In Arithmetic Progression|简单|3758|54%|
|[1229](https://leetcode-cn.com/problems/meeting-scheduler/)|Meeting Scheduler|中等|3466|35%|
|[1230](https://leetcode-cn.com/problems/toss-strange-coins/)|Toss Strange Coins|中等|2861|42%|
|[1199](https://leetcode-cn.com/problems/minimum-time-to-build-blocks/)|Minimum Time to Build Blocks|困难|1306|39%|
|[1089](https://leetcode-cn.com/problems/duplicate-zeros/)|Duplicate Zeros|简单|18798|57%|
|[1090](https://leetcode-cn.com/problems/largest-values-from-labels/)|Largest Values From Labels|中等|4331|51%|
|[1092](https://leetcode-cn.com/problems/shortest-common-supersequence/)|Shortest Common Supersequence |困难|3031|43%|
|[1091](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)|Shortest Path in Binary Matrix|中等|26801|33%|
|[1082](https://leetcode-cn.com/problems/sales-analysis-i/)|Sales Analysis I|简单|4926|74%|
|[1083](https://leetcode-cn.com/problems/sales-analysis-ii/)|Sales Analysis II|简单|6005|51%|
|[1084](https://leetcode-cn.com/problems/sales-analysis-iii/)|Sales Analysis III|简单|6364|53%|
|[1243](https://leetcode-cn.com/problems/array-transformation/)|Array Transformation|简单|2798|52%|
|[1244](https://leetcode-cn.com/problems/design-a-leaderboard/)|Design A Leaderboard|中等|2179|56%|
|[1245](https://leetcode-cn.com/problems/tree-diameter/)|Tree Diameter|中等|3975|45%|
|[1216](https://leetcode-cn.com/problems/valid-palindrome-iii/)|Valid Palindrome III|困难|1747|49%|
|[511](https://leetcode-cn.com/problems/game-play-analysis-i/)|Game Play Analysis I|简单|6806|72%|
|[512](https://leetcode-cn.com/problems/game-play-analysis-ii/)|Game Play Analysis II|简单|9502|52%|
|[534](https://leetcode-cn.com/problems/game-play-analysis-iii/)|Game Play Analysis III|中等|5729|65%|
|[550](https://leetcode-cn.com/problems/game-play-analysis-iv/)|Game Play Analysis IV|中等|7388|43%|
|[1093](https://leetcode-cn.com/problems/statistics-from-a-large-sample/)|Statistics from a Large Sample|中等|5092|41%|
|[1094](https://leetcode-cn.com/problems/car-pooling/)|Car Pooling|中等|15820|58%|
|[1095](https://leetcode-cn.com/problems/find-in-mountain-array/)|Find in Mountain Array|困难|37074|36%|
|[1117](https://leetcode-cn.com/problems/building-h2o/)|Building H2O|中等|15811|50%|
|[1115](https://leetcode-cn.com/problems/print-foobar-alternately/)|Print FooBar Alternately|中等|34427|55%|
|[1096](https://leetcode-cn.com/problems/brace-expansion-ii/)|Brace Expansion II|困难|1807|51%|
|[1256](https://leetcode-cn.com/problems/encode-number/)|Encode Number|中等|1757|62%|
|[1257](https://leetcode-cn.com/problems/smallest-common-region/)|Smallest Common Region|中等|2192|52%|
|[1258](https://leetcode-cn.com/problems/synonymous-sentences/)|Synonymous Sentences|中等|1383|56%|
|[1231](https://leetcode-cn.com/problems/divide-chocolate/)|Divide Chocolate|困难|2405|55%|
|[1097](https://leetcode-cn.com/problems/game-play-analysis-v/)|Game Play Analysis V|困难|2359|48%|
|[1104](https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/)|Path In Zigzag Labelled Binary Tree|中等|6987|69%|
|[1103](https://leetcode-cn.com/problems/distribute-candies-to-people/)|Distribute Candies to People|简单|43764|62%|
|[1105](https://leetcode-cn.com/problems/filling-bookcase-shelves/)|Filling Bookcase Shelves|中等|5362|51%|
|[1106](https://leetcode-cn.com/problems/parsing-a-boolean-expression/)|Parsing A Boolean Expression|困难|4472|54%|
|[1098](https://leetcode-cn.com/problems/unpopular-books/)|Unpopular Books|中等|4814|45%|
|[1271](https://leetcode-cn.com/problems/hexspeak/)|Hexspeak|简单|3411|48%|
|[1272](https://leetcode-cn.com/problems/remove-interval/)|Remove Interval|中等|2150|48%|
|[1273](https://leetcode-cn.com/problems/delete-tree-nodes/)|Delete Tree Nodes|中等|1890|53%|
|[1246](https://leetcode-cn.com/problems/palindrome-removal/)|Palindrome Removal|困难|3461|46%|
|[1114](https://leetcode-cn.com/problems/print-in-order/)|Print in Order|简单|55012|63%|
|[1107](https://leetcode-cn.com/problems/new-users-daily-count/)|New Users Daily Count|中等|3990|40%|
|[1108](https://leetcode-cn.com/problems/defanging-an-ip-address/)|Defanging an IP Address|简单|47378|82%|
|[1109](https://leetcode-cn.com/problems/corporate-flight-bookings/)|Corporate Flight Bookings|中等|25045|45%|
|[1110](https://leetcode-cn.com/problems/delete-nodes-and-return-forest/)|Delete Nodes And Return Forest|中等|6519|59%|
|[1111](https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/)|Maximum Nesting Depth of Two Valid Parentheses Strings|中等|25940|77%|
|[1188](https://leetcode-cn.com/problems/design-bounded-blocking-queue/)|Design Bounded Blocking Queue|中等|1668|69%|
|[1286](https://leetcode-cn.com/problems/iterator-for-combination/)|Iterator for Combination|中等|3069|59%|
|[1291](https://leetcode-cn.com/problems/sequential-digits/)|Sequential Digits|中等|7454|49%|
|[1259](https://leetcode-cn.com/problems/handshakes-that-dont-cross/)|Handshakes That Don&#39;t Cross|困难|2367|46%|
|[1112](https://leetcode-cn.com/problems/highest-grade-for-each-student/)|Highest Grade For Each Student|中等|2995|59%|
|[1113](https://leetcode-cn.com/problems/reported-posts/)|Reported Posts|简单|3555|54%|
|[1116](https://leetcode-cn.com/problems/print-zero-even-odd/)|Print Zero Even Odd|中等|20411|48%|
|[1122](https://leetcode-cn.com/problems/relative-sort-array/)|Relative Sort Array|简单|28306|66%|
|[1123](https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves/)|Lowest Common Ancestor of Deepest Leaves|中等|4777|66%|
|[1124](https://leetcode-cn.com/problems/longest-well-performing-interval/)|Longest Well-Performing Interval|中等|20938|28%|
|[1125](https://leetcode-cn.com/problems/smallest-sufficient-team/)|Smallest Sufficient Team|困难|3767|44%|
|[1287](https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/)|[Element Appearing More Than 25% In Sorted Array](./src/1287-element-appearing-more-than-25-in-sorted-array/)|简单|12206|61%|
|[1288](https://leetcode-cn.com/problems/remove-covered-intervals/)|Remove Covered Intervals|中等|4996|55%|
|[1289](https://leetcode-cn.com/problems/minimum-falling-path-sum-ii/)|Minimum Falling Path Sum II|困难|3460|60%|
|[1126](https://leetcode-cn.com/problems/active-businesses/)|Active Businesses|中等|2216|66%|
|[1127](https://leetcode-cn.com/problems/user-purchase-platform/)|User Purchase Platform|困难|1628|45%|
|[1128](https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/)|Number of Equivalent Domino Pairs|简单|14012|45%|
|[1130](https://leetcode-cn.com/problems/minimum-cost-tree-from-leaf-values/)|Minimum Cost Tree From Leaf Values|中等|4336|60%|
|[1129](https://leetcode-cn.com/problems/shortest-path-with-alternating-colors/)|Shortest Path with Alternating Colors|中等|10368|34%|
|[1131](https://leetcode-cn.com/problems/maximum-of-absolute-value-expression/)|Maximum of Absolute Value Expression|中等|4140|42%|
|[1299](https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/)|Replace Elements with Greatest Element on Right Side|简单|18153|77%|
|[1300](https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target/)|Sum of Mutated Array Closest to Target|中等|39706|47%|
|[1274](https://leetcode-cn.com/problems/number-of-ships-in-a-rectangle/)|Number of Ships in a Rectangle|困难|1111|54%|
|[1301](https://leetcode-cn.com/problems/number-of-paths-with-max-score/)|Number of Paths with Max Score|困难|3935|32%|
|[1137](https://leetcode-cn.com/problems/n-th-tribonacci-number/)|N-th Tribonacci Number|简单|29633|52%|
|[1132](https://leetcode-cn.com/problems/reported-posts-ii/)|Reported Posts II|中等|2910|39%|
|[1138](https://leetcode-cn.com/problems/alphabet-board-path/)|Alphabet Board Path|中等|8908|40%|
|[1139](https://leetcode-cn.com/problems/largest-1-bordered-square/)|Largest 1-Bordered Square|中等|10663|45%|
|[1140](https://leetcode-cn.com/problems/stone-game-ii/)|Stone Game II|中等|4341|62%|
|[1313](https://leetcode-cn.com/problems/decompress-run-length-encoded-list/)|Decompress Run-Length Encoded List|简单|20362|82%|
|[1314](https://leetcode-cn.com/problems/matrix-block-sum/)|Matrix Block Sum|中等|4516|70%|
|[1315](https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent/)|Sum of Nodes with Even-Valued Grandparent|中等|7014|81%|
|[1316](https://leetcode-cn.com/problems/distinct-echo-substrings/)|Distinct Echo Substrings|困难|3588|41%|
|[1141](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-i/)|User Activity for the Past 30 Days I|简单|4465|52%|
|[1142](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-ii/)|User Activity for the Past 30 Days II|简单|5240|37%|
|[1144](https://leetcode-cn.com/problems/decrease-elements-to-make-array-zigzag/)|Decrease Elements To Make Array Zigzag|中等|14832|42%|
|[1145](https://leetcode-cn.com/problems/binary-tree-coloring-game/)|Binary Tree Coloring Game|中等|8477|43%|
|[1146](https://leetcode-cn.com/problems/snapshot-array/)|Snapshot Array|中等|8866|25%|
|[1143](https://leetcode-cn.com/problems/longest-common-subsequence/)|Longest Common Subsequence|中等|51061|60%|
|[1147](https://leetcode-cn.com/problems/longest-chunked-palindrome-decomposition/)|Longest Chunked Palindrome Decomposition|困难|3587|53%|
|[1328](https://leetcode-cn.com/problems/break-a-palindrome/)|Break a Palindrome|中等|4347|43%|
|[1329](https://leetcode-cn.com/problems/sort-the-matrix-diagonally/)|Sort the Matrix Diagonally|中等|3978|76%|
|[1302](https://leetcode-cn.com/problems/deepest-leaves-sum/)|Deepest Leaves Sum|中等|9746|80%|
|[1330](https://leetcode-cn.com/problems/reverse-subarray-to-maximize-array-value/)|Reverse Subarray To Maximize Array Value|困难|1797|35%|
|[1331](https://leetcode-cn.com/problems/rank-transform-of-an-array/)|Rank Transform of an Array|简单|7603|52%|
|[1148](https://leetcode-cn.com/problems/article-views-i/)|Article Views I|简单|3244|74%|
|[1149](https://leetcode-cn.com/problems/article-views-ii/)|Article Views II|中等|3087|43%|
|[1154](https://leetcode-cn.com/problems/day-of-the-year/)|Day of the Year|简单|14323|54%|
|[1156](https://leetcode-cn.com/problems/swap-for-longest-repeated-character-substring/)|Swap For Longest Repeated Character Substring|中等|6499|43%|
|[1157](https://leetcode-cn.com/problems/online-majority-element-in-subarray/)|Online Majority Element In Subarray|困难|6194|26%|
|[1155](https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/)|Number of Dice Rolls With Target Sum|中等|9776|44%|
|[1171](https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/)|Remove Zero Sum Consecutive Nodes from Linked List|中等|15144|42%|
|[1158](https://leetcode-cn.com/problems/market-analysis-i/)|Market Analysis I|中等|2667|55%|
|[1159](https://leetcode-cn.com/problems/market-analysis-ii/)|Market Analysis II|困难|1705|47%|
|[1172](https://leetcode-cn.com/problems/dinner-plate-stacks/)|Dinner Plate Stacks|困难|5382|24%|
|[1236](https://leetcode-cn.com/problems/web-crawler/)|Web Crawler|中等|1106|53%|
|[1169](https://leetcode-cn.com/problems/invalid-transactions/)|Invalid Transactions|中等|13996|31%|
|[1170](https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character/)|Compare Strings by Frequency of the Smallest Character|简单|12840|60%|
|[1360](https://leetcode-cn.com/problems/number-of-days-between-two-dates/)|Number of Days Between Two Dates|简单|9908|47%|
|[1361](https://leetcode-cn.com/problems/validate-binary-tree-nodes/)|Validate Binary Tree Nodes|中等|9667|45%|
|[1362](https://leetcode-cn.com/problems/closest-divisors/)|Closest Divisors|中等|6929|49%|
|[1363](https://leetcode-cn.com/problems/largest-multiple-of-three/)|Largest Multiple of Three|困难|7662|32%|
|[1164](https://leetcode-cn.com/problems/product-price-at-a-given-date/)|Product Price at a Given Date|中等|2503|62%|
|[1175](https://leetcode-cn.com/problems/prime-arrangements/)|Prime Arrangements|简单|8689|47%|
|[1176](https://leetcode-cn.com/problems/diet-plan-performance/)|Diet Plan Performance|简单|7386|41%|
|[1177](https://leetcode-cn.com/problems/can-make-palindrome-from-substring/)|Can Make Palindrome from Substring|中等|16154|23%|
|[1178](https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle/)|Number of Valid Words for Each Puzzle|困难|4993|28%|
|[1507](https://leetcode-cn.com/problems/reformat-date/)|Reformat Date|简单|4464|58%|
|[1390](https://leetcode-cn.com/problems/four-divisors/)|Four Divisors|中等|13570|33%|
|[1382](https://leetcode-cn.com/problems/balance-a-binary-search-tree/)|Balance a Binary Search Tree|中等|6674|66%|
|[1425](https://leetcode-cn.com/problems/constrained-subsequence-sum/)|Constrained Subsequence Sum|困难|4468|40%|
|[1184](https://leetcode-cn.com/problems/distance-between-bus-stops/)|Distance Between Bus Stops|简单|13349|57%|
|[1186](https://leetcode-cn.com/problems/maximum-subarray-sum-with-one-deletion/)|Maximum Subarray Sum with One Deletion|中等|10032|34%|
|[1185](https://leetcode-cn.com/problems/day-of-the-week/)|Day of the Week|简单|10620|60%|
|[1187](https://leetcode-cn.com/problems/make-array-strictly-increasing/)|Make Array Strictly Increasing|困难|2521|39%|
|[1173](https://leetcode-cn.com/problems/immediate-food-delivery-i/)|Immediate Food Delivery I|简单|3271|75%|
|[1174](https://leetcode-cn.com/problems/immediate-food-delivery-ii/)|Immediate Food Delivery II|中等|2339|55%|
|[1483](https://leetcode-cn.com/problems/kth-ancestor-of-a-tree-node/)|Kth Ancestor of a Tree Node|困难|6942|25%|
|[1189](https://leetcode-cn.com/problems/maximum-number-of-balloons/)|Maximum Number of Balloons|简单|18335|63%|
|[1190](https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/)|Reverse Substrings Between Each Pair of Parentheses|中等|11117|54%|
|[1191](https://leetcode-cn.com/problems/k-concatenation-maximum-sum/)|K-Concatenation Maximum Sum|中等|11373|25%|
|[1192](https://leetcode-cn.com/problems/critical-connections-in-a-network/)|Critical Connections in a Network|困难|4050|41%|
|[1179](https://leetcode-cn.com/problems/reformat-department-table/)|Reformat Department Table|简单|16764|61%|
|[1405](https://leetcode-cn.com/problems/longest-happy-string/)|Longest Happy String|中等|6584|44%|
|[1200](https://leetcode-cn.com/problems/minimum-absolute-difference/)|Minimum Absolute Difference|简单|14418|66%|
|[1201](https://leetcode-cn.com/problems/ugly-number-iii/)|Ugly Number III|中等|13588|22%|
|[1202](https://leetcode-cn.com/problems/smallest-string-with-swaps/)|Smallest String With Swaps|中等|10625|32%|
|[1203](https://leetcode-cn.com/problems/sort-items-by-groups-respecting-dependencies/)|Sort Items by Groups Respecting Dependencies|困难|2704|38%|
|[1195](https://leetcode-cn.com/problems/fizz-buzz-multithreaded/)|Fizz Buzz Multithreaded|中等|8910|62%|
|[1193](https://leetcode-cn.com/problems/monthly-transactions-i/)|Monthly Transactions I|中等|2715|61%|
|[1194](https://leetcode-cn.com/problems/tournament-winners/)|Tournament Winners|困难|1683|46%|
|[1207](https://leetcode-cn.com/problems/unique-number-of-occurrences/)|Unique Number of Occurrences|简单|18439|69%|
|[1209](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii/)|Remove All Adjacent Duplicates in String II|中等|10222|52%|
|[1208](https://leetcode-cn.com/problems/get-equal-substrings-within-budget/)|Get Equal Substrings Within Budget|中等|19647|38%|
|[1210](https://leetcode-cn.com/problems/minimum-moves-to-reach-target-with-rotations/)|Minimum Moves to Reach Target with Rotations|困难|4153|42%|
|[1514](https://leetcode-cn.com/problems/path-with-maximum-probability/)|Path with Maximum Probability|中等|11913|23%|
|[1204](https://leetcode-cn.com/problems/last-person-to-fit-in-the-elevator/)|Last Person to Fit in the Elevator|中等|1821|69%|
|[1205](https://leetcode-cn.com/problems/monthly-transactions-ii/)|Monthly Transactions II|中等|3493|47%|
|[1217](https://leetcode-cn.com/problems/play-with-chips/)|Play with Chips|简单|14761|68%|
|[1218](https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/)|Longest Arithmetic Subsequence of Given Difference|中等|11456|39%|
|[1219](https://leetcode-cn.com/problems/path-with-maximum-gold/)|Path with Maximum Gold|中等|9342|61%|
|[1220](https://leetcode-cn.com/problems/count-vowels-permutation/)|Count Vowels Permutation|困难|4645|51%|
|[1206](https://leetcode-cn.com/problems/design-skiplist/)|Design Skiplist|困难|2642|58%|
|[1211](https://leetcode-cn.com/problems/queries-quality-and-percentage/)|Queries Quality and Percentage|简单|2947|63%|
|[1212](https://leetcode-cn.com/problems/team-scores-in-football-tournament/)|Team Scores in Football Tournament|中等|2316|50%|
|[1226](https://leetcode-cn.com/problems/the-dining-philosophers/)|The Dining Philosophers|中等|7648|59%|
|[1221](https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/)|Split a String in Balanced Strings|简单|26972|78%|
|[1222](https://leetcode-cn.com/problems/queens-that-can-attack-the-king/)|Queens That Can Attack the King|中等|5332|65%|
|[1223](https://leetcode-cn.com/problems/dice-roll-simulation/)|Dice Roll Simulation|中等|4467|45%|
|[1224](https://leetcode-cn.com/problems/maximum-equal-frequency/)|Maximum Equal Frequency|困难|9827|29%|
|[1427](https://leetcode-cn.com/problems/perform-string-shifts/)|Perform String Shifts|简单|329|67%|
|[1232](https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/)|Check If It Is a Straight Line|简单|14938|49%|
|[1233](https://leetcode-cn.com/problems/remove-sub-folders-from-the-filesystem/)|Remove Sub-Folders from the Filesystem|中等|9828|45%|
|[1234](https://leetcode-cn.com/problems/replace-the-substring-for-balanced-string/)|Replace the Substring for Balanced String|中等|10245|30%|
|[1235](https://leetcode-cn.com/problems/maximum-profit-in-job-scheduling/)|Maximum Profit in Job Scheduling|困难|5783|42%|
|[1225](https://leetcode-cn.com/problems/report-contiguous-dates/)|Report Contiguous Dates|困难|1243|53%|
|[1237](https://leetcode-cn.com/problems/find-positive-integer-solution-for-a-given-equation/)|Find Positive Integer Solution for a Given Equation|简单|8835|69%|
|[1238](https://leetcode-cn.com/problems/circular-permutation-in-binary-representation/)|Circular Permutation in Binary Representation|中等|3108|62%|
|[1239](https://leetcode-cn.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/)|Maximum Length of a Concatenated String with Unique Characters|中等|26528|39%|
|[1240](https://leetcode-cn.com/problems/tiling-a-rectangle-with-the-fewest-squares/)|Tiling a Rectangle with the Fewest Squares|困难|2960|46%|
|[1227](https://leetcode-cn.com/problems/airplane-seat-assignment-probability/)|Airplane Seat Assignment Probability|中等|7596|63%|
|[1429](https://leetcode-cn.com/problems/first-unique-number/)|First Unique Number|中等|253|58%|
|[1242](https://leetcode-cn.com/problems/web-crawler-multithreaded/)|Web Crawler Multithreaded|中等|944|40%|
|[1247](https://leetcode-cn.com/problems/minimum-swaps-to-make-strings-equal/)|Minimum Swaps to Make Strings Equal|中等|7516|57%|
|[1248](https://leetcode-cn.com/problems/count-number-of-nice-subarrays/)|Count Number of Nice Subarrays|中等|43994|54%|
|[1249](https://leetcode-cn.com/problems/minimum-remove-to-make-valid-parentheses/)|Minimum Remove to Make Valid Parentheses|中等|14458|56%|
|[1250](https://leetcode-cn.com/problems/check-if-it-is-a-good-array/)|Check If It Is a Good Array|困难|2960|51%|
|[1428](https://leetcode-cn.com/problems/leftmost-column-with-at-least-a-one/)|Leftmost Column with at Least a One|中等|226|66%|
|[1241](https://leetcode-cn.com/problems/number-of-comments-per-post/)|Number of Comments per Post|简单|3237|59%|
|[1252](https://leetcode-cn.com/problems/cells-with-odd-values-in-a-matrix/)|Cells with Odd Values in a Matrix|简单|13715|75%|
|[1253](https://leetcode-cn.com/problems/reconstruct-a-2-row-binary-matrix/)|Reconstruct a 2-Row Binary Matrix|中等|9095|37%|
|[1254](https://leetcode-cn.com/problems/number-of-closed-islands/)|Number of Closed Islands|中等|10419|57%|
|[1255](https://leetcode-cn.com/problems/maximum-score-words-formed-by-letters/)|Maximum Score Words Formed by Letters|困难|2680|64%|
|[1260](https://leetcode-cn.com/problems/shift-2d-grid/)|Shift 2D Grid|简单|9990|59%|
|[1261](https://leetcode-cn.com/problems/find-elements-in-a-contaminated-binary-tree/)|Find Elements in a Contaminated Binary Tree|中等|5837|71%|
|[1262](https://leetcode-cn.com/problems/greatest-sum-divisible-by-three/)|Greatest Sum Divisible by Three|中等|9649|47%|
|[1263](https://leetcode-cn.com/problems/minimum-moves-to-move-a-box-to-their-target-location/)|Minimum Moves to Move a Box to Their Target Location|困难|3917|40%|
|[1251](https://leetcode-cn.com/problems/average-selling-price/)|Average Selling Price|简单|3573|78%|
|[1426](https://leetcode-cn.com/problems/counting-elements/)|Counting Elements|简单|400|64%|
|[1266](https://leetcode-cn.com/problems/minimum-time-visiting-all-points/)|Minimum Time Visiting All Points|简单|18736|81%|
|[1267](https://leetcode-cn.com/problems/count-servers-that-communicate/)|Count Servers that Communicate|中等|7452|59%|
|[1268](https://leetcode-cn.com/problems/search-suggestions-system/)|Search Suggestions System|中等|7845|55%|
|[1269](https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/)|Number of Ways to Stay in the Same Place After Some Steps|困难|5784|39%|
|[1264](https://leetcode-cn.com/problems/page-recommendations/)|Page Recommendations|中等|1927|59%|
|[1275](https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game/)|Find Winner on a Tic Tac Toe Game|简单|7393|52%|
|[1276](https://leetcode-cn.com/problems/number-of-burgers-with-no-waste-of-ingredients/)|Number of Burgers with No Waste of Ingredients|中等|7779|49%|
|[1277](https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/)|Count Square Submatrices with All Ones|中等|12643|70%|
|[1278](https://leetcode-cn.com/problems/palindrome-partitioning-iii/)|Palindrome Partitioning III|困难|2852|58%|
|[1265](https://leetcode-cn.com/problems/print-immutable-linked-list-in-reverse/)|Print Immutable Linked List in Reverse|中等|1096|92%|
|[1270](https://leetcode-cn.com/problems/all-people-report-to-the-given-manager/)|All People Report to the Given Manager|中等|1924|81%|
|[1281](https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/)|Subtract the Product and Sum of Digits of an Integer|简单|29922|82%|
|[1282](https://leetcode-cn.com/problems/group-the-people-given-the-group-size-they-belong-to/)|Group the People Given the Group Size They Belong To|中等|8314|80%|
|[1283](https://leetcode-cn.com/problems/find-the-smallest-divisor-given-a-threshold/)|Find the Smallest Divisor Given a Threshold|中等|10449|37%|
|[1284](https://leetcode-cn.com/problems/minimum-number-of-flips-to-convert-binary-matrix-to-zero-matrix/)|Minimum Number of Flips to Convert Binary Matrix to Zero Matrix|困难|2765|63%|
|[1279](https://leetcode-cn.com/problems/traffic-light-controlled-intersection/)|Traffic Light Controlled Intersection|简单|746|58%|
|[1290](https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/)|[Convert Binary Number in a Linked List to Integer](./src/1290-convert-binary-number-in-a-linked-list-to-integer/)|简单|29209|80%|
|[1292](https://leetcode-cn.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/)|Maximum Side Length of a Square with Sum Less than or Equal to Threshold|中等|7059|42%|
|[1293](https://leetcode-cn.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/)|Shortest Path in a Grid with Obstacles Elimination|困难|16574|36%|
|[1280](https://leetcode-cn.com/problems/students-and-examinations/)|Students and Examinations|简单|3572|53%|
|[1285](https://leetcode-cn.com/problems/find-the-start-and-end-number-of-continuous-ranges/)|Find the Start and End Number of Continuous Ranges|中等|1449|77%|
|[1295](https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/)|Find Numbers with Even Number of Digits|简单|33348|82%|
|[1296](https://leetcode-cn.com/problems/divide-array-in-sets-of-k-consecutive-numbers/)|Divide Array in Sets of K Consecutive Numbers|中等|10241|42%|
|[1297](https://leetcode-cn.com/problems/maximum-number-of-occurrences-of-a-substring/)|Maximum Number of Occurrences of a Substring|中等|7873|43%|
|[1298](https://leetcode-cn.com/problems/maximum-candies-you-can-get-from-boxes/)|Maximum Candies You Can Get from Boxes|困难|2857|58%|
|[1294](https://leetcode-cn.com/problems/weather-type-in-each-country/)|Weather Type in Each Country|简单|2889|65%|
|[1304](https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero/)|Find N Unique Integers Sum up to Zero|简单|15445|74%|
|[1305](https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/)|All Elements in Two Binary Search Trees|中等|9542|73%|
|[1306](https://leetcode-cn.com/problems/jump-game-iii/)|Jump Game III|中等|12336|58%|
|[1307](https://leetcode-cn.com/problems/verbal-arithmetic-puzzle/)|Verbal Arithmetic Puzzle|困难|3586|33%|
|[1430](https://leetcode-cn.com/problems/check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree/)|Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree|中等|221|58%|
|[1309](https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping/)|Decrypt String from Alphabet to Integer Mapping|简单|11456|74%|
|[1310](https://leetcode-cn.com/problems/xor-queries-of-a-subarray/)|XOR Queries of a Subarray|中等|5176|64%|
|[1311](https://leetcode-cn.com/problems/get-watched-videos-by-your-friends/)|Get Watched Videos by Your Friends|中等|9657|35%|
|[1312](https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/)|Minimum Insertion Steps to Make a String Palindrome|困难|3588|57%|
|[1303](https://leetcode-cn.com/problems/find-the-team-size/)|Find the Team Size|简单|3205|82%|
|[1308](https://leetcode-cn.com/problems/running-total-for-different-genders/)|Running Total for Different Genders|中等|1721|73%|
|[1317](https://leetcode-cn.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/)|Convert Integer to the Sum of Two No-Zero Integers|简单|8831|61%|
|[1318](https://leetcode-cn.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/)|Minimum Flips to Make a OR b Equal to c|中等|4592|61%|
|[1319](https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected/)|Number of Operations to Make Network Connected|中等|11019|50%|
|[1320](https://leetcode-cn.com/problems/minimum-distance-to-type-a-word-using-two-fingers/)|Minimum Distance to Type a Word Using Two Fingers|困难|2642|57%|
|[1342](https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/)|Number of Steps to Reduce a Number to Zero|简单|27090|82%|
|[1343](https://leetcode-cn.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/)|Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold|中等|13993|53%|
|[1344](https://leetcode-cn.com/problems/angle-between-hands-of-a-clock/)|Angle Between Hands of a Clock|中等|4796|58%|
|[1345](https://leetcode-cn.com/problems/jump-game-iv/)|Jump Game IV|困难|7805|32%|
|[1323](https://leetcode-cn.com/problems/maximum-69-number/)|Maximum 69 Number|简单|15363|75%|
|[1324](https://leetcode-cn.com/problems/print-words-vertically/)|Print Words Vertically|中等|5719|56%|
|[1325](https://leetcode-cn.com/problems/delete-leaves-with-a-given-value/)|Delete Leaves With a Given Value|中等|5623|70%|
|[1326](https://leetcode-cn.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/)|Minimum Number of Taps to Open to Water a Garden|困难|6134|45%|
|[1321](https://leetcode-cn.com/problems/restaurant-growth/)|Restaurant Growth|中等|1594|56%|
|[1322](https://leetcode-cn.com/problems/ads-performance/)|Ads Performance|简单|2484|58%|
|[1332](https://leetcode-cn.com/problems/remove-palindromic-subsequences/)|Remove Palindromic Subsequences|简单|6427|64%|
|[1333](https://leetcode-cn.com/problems/filter-restaurants-by-vegan-friendly-price-and-distance/)|Filter Restaurants by Vegan-Friendly, Price and Distance|中等|5554|50%|
|[1334](https://leetcode-cn.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/)|Find the City With the Smallest Number of Neighbors at a Threshold Distance|中等|5243|45%|
|[1335](https://leetcode-cn.com/problems/minimum-difficulty-of-a-job-schedule/)|Minimum Difficulty of a Job Schedule|困难|2935|57%|
|[1356](https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/)|Sort Integers by The Number of 1 Bits|简单|9415|67%|
|[1357](https://leetcode-cn.com/problems/apply-discount-every-n-orders/)|Apply Discount Every n Orders|中等|3805|47%|
|[1358](https://leetcode-cn.com/problems/number-of-substrings-containing-all-three-characters/)|Number of Substrings Containing All Three Characters|中等|6184|42%|
|[1359](https://leetcode-cn.com/problems/count-all-valid-pickup-and-delivery-options/)|Count All Valid Pickup and Delivery Options|困难|2900|54%|
|[1327](https://leetcode-cn.com/problems/list-the-products-ordered-in-a-period/)|List the Products Ordered in a Period|简单|2282|70%|
|[1337](https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/)|The K Weakest Rows in a Matrix|简单|9912|63%|
|[1338](https://leetcode-cn.com/problems/reduce-array-size-to-the-half/)|Reduce Array Size to The Half|中等|7215|61%|
|[1339](https://leetcode-cn.com/problems/maximum-product-of-splitted-binary-tree/)|Maximum Product of Splitted Binary Tree|中等|9779|37%|
|[1340](https://leetcode-cn.com/problems/jump-game-v/)|Jump Game V|困难|4379|53%|
|[1336](https://leetcode-cn.com/problems/number-of-transactions-per-visit/)|Number of Transactions per Visit|困难|901|42%|
|[1346](https://leetcode-cn.com/problems/check-if-n-and-its-double-exist/)|Check If N and Its Double Exist|简单|16689|43%|
|[1347](https://leetcode-cn.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/)|Minimum Number of Steps to Make Two Strings Anagram|中等|5976|70%|
|[1348](https://leetcode-cn.com/problems/tweet-counts-per-frequency/)|Tweet Counts Per Frequency|中等|7208|26%|
|[1349](https://leetcode-cn.com/problems/maximum-students-taking-exam/)|Maximum Students Taking Exam|困难|3777|43%|
|[1370](https://leetcode-cn.com/problems/increasing-decreasing-string/)|Increasing Decreasing String|简单|9587|74%|
|[1371](https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/)|Find the Longest Substring Containing Vowels in Even Counts|中等|27179|58%|
|[1372](https://leetcode-cn.com/problems/longest-zigzag-path-in-a-binary-tree/)|Longest ZigZag Path in a Binary Tree|中等|6874|47%|
|[1373](https://leetcode-cn.com/problems/maximum-sum-bst-in-binary-tree/)|Maximum Sum BST in Binary Tree|困难|5018|41%|
|[1351](https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/)|Count Negative Numbers in a Sorted Matrix|简单|17975|76%|
|[1352](https://leetcode-cn.com/problems/product-of-the-last-k-numbers/)|Product of the Last K Numbers|中等|8866|41%|
|[1353](https://leetcode-cn.com/problems/maximum-number-of-events-that-can-be-attended/)|Maximum Number of Events That Can Be Attended|中等|15616|27%|
|[1354](https://leetcode-cn.com/problems/construct-target-array-with-multiple-sums/)|Construct Target Array With Multiple Sums|困难|6419|28%|
|[1341](https://leetcode-cn.com/problems/movie-rating/)|Movie Rating|中等|1708|46%|
|[1350](https://leetcode-cn.com/problems/students-with-invalid-departments/)|Students With Invalid Departments|简单|2711|85%|
|[1365](https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/)|How Many Numbers Are Smaller Than the Current Number|简单|25551|82%|
|[1366](https://leetcode-cn.com/problems/rank-teams-by-votes/)|Rank Teams by Votes|中等|7640|49%|
|[1367](https://leetcode-cn.com/problems/linked-list-in-binary-tree/)|Linked List in Binary Tree|中等|18750|39%|
|[1368](https://leetcode-cn.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/)|Minimum Cost to Make at Least One Valid Path in a Grid|困难|4254|48%|
|[1385](https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays/)|Find the Distance Value Between Two Arrays|简单|8023|71%|
|[1386](https://leetcode-cn.com/problems/cinema-seat-allocation/)|Cinema Seat Allocation|中等|12534|28%|
|[1387](https://leetcode-cn.com/problems/sort-integers-by-the-power-value/)|Sort Integers by The Power Value|中等|6055|68%|
|[1388](https://leetcode-cn.com/problems/pizza-with-3n-slices/)|Pizza With 3n Slices|困难|2331|43%|
|[1374](https://leetcode-cn.com/problems/generate-a-string-with-characters-that-have-odd-counts/)|Generate a String With Characters That Have Odd Counts|简单|10832|72%|
|[1375](https://leetcode-cn.com/problems/bulb-switcher-iii/)|Bulb Switcher III|中等|10278|51%|
|[1376](https://leetcode-cn.com/problems/time-needed-to-inform-all-employees/)|Time Needed to Inform All Employees|中等|9295|47%|
|[1377](https://leetcode-cn.com/problems/frog-position-after-t-seconds/)|Frog Position After T Seconds|困难|8584|29%|
|[1355](https://leetcode-cn.com/problems/activity-participants/)|Activity Participants|中等|1444|62%|
|[1364](https://leetcode-cn.com/problems/number-of-trusted-contacts-of-a-customer/)|Number of Trusted Contacts of a Customer|中等|1147|68%|
|[1380](https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/)|Lucky Numbers in a Matrix|简单|11231|70%|
|[1381](https://leetcode-cn.com/problems/design-a-stack-with-increment-operation/)|Design a Stack With Increment Operation|中等|7612|72%|
|[1379](https://leetcode-cn.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/)|Find a Corresponding Node of a Binary Tree in a Clone of That Tree|中等|3912|82%|
|[1383](https://leetcode-cn.com/problems/maximum-performance-of-a-team/)|Maximum Performance of a Team|困难|7773|29%|
|[1399](https://leetcode-cn.com/problems/count-largest-group/)|Count Largest Group|简单|6298|65%|
|[1401](https://leetcode-cn.com/problems/circle-and-rectangle-overlapping/)|Circle and Rectangle Overlapping|中等|5266|36%|
|[1400](https://leetcode-cn.com/problems/construct-k-palindrome-strings/)|Construct K Palindrome Strings|中等|4813|57%|
|[1402](https://leetcode-cn.com/problems/reducing-dishes/)|Reducing Dishes|困难|5512|73%|
|[1369](https://leetcode-cn.com/problems/get-the-second-most-recent-activity/)|Get the Second Most Recent Activity|困难|970|52%|
|[1389](https://leetcode-cn.com/problems/create-target-array-in-the-given-order/)|Create Target Array in the Given Order|简单|14788|82%|
|[1391](https://leetcode-cn.com/problems/check-if-there-is-a-valid-path-in-a-grid/)|Check if There is a Valid Path in a Grid|中等|10088|38%|
|[1392](https://leetcode-cn.com/problems/longest-happy-prefix/)|Longest Happy Prefix|困难|10247|35%|
|[1378](https://leetcode-cn.com/problems/replace-employee-id-with-the-unique-identifier/)|Replace Employee ID With The Unique Identifier|简单|2269|85%|
|[1394](https://leetcode-cn.com/problems/find-lucky-integer-in-an-array/)|Find Lucky Integer in an Array|简单|12662|67%|
|[1395](https://leetcode-cn.com/problems/count-number-of-teams/)|Count Number of Teams|中等|9501|82%|
|[1396](https://leetcode-cn.com/problems/design-underground-system/)|Design Underground System|中等|8053|44%|
|[1397](https://leetcode-cn.com/problems/find-all-good-strings/)|Find All Good Strings|困难|1922|32%|
|[1413](https://leetcode-cn.com/problems/minimum-value-to-get-positive-step-by-step-sum/)|Minimum Value to Get Positive Step by Step Sum|简单|6682|68%|
|[1414](https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/)|Find the Minimum Number of Fibonacci Numbers Whose Sum Is K|中等|6564|55%|
|[1415](https://leetcode-cn.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/)|The k-th Lexicographical String of All Happy Strings of Length n|中等|3789|69%|
|[1416](https://leetcode-cn.com/problems/restore-the-array/)|Restore The Array|困难|4145|35%|
|[1384](https://leetcode-cn.com/problems/total-sales-amount-by-year/)|Total Sales Amount by Year|困难|680|52%|
|[1403](https://leetcode-cn.com/problems/minimum-subsequence-in-non-increasing-order/)|Minimum Subsequence in Non-Increasing Order|简单|12169|69%|
|[1404](https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/)|Number of Steps to Reduce a Number in Binary Representation to One|中等|10056|46%|
|[1406](https://leetcode-cn.com/problems/stone-game-iii/)|Stone Game III|困难|3957|56%|
|[1393](https://leetcode-cn.com/problems/capital-gainloss/)|Capital Gain/Loss|中等|1019|87%|
|[1408](https://leetcode-cn.com/problems/string-matching-in-an-array/)|String Matching in an Array|简单|11213|59%|
|[1409](https://leetcode-cn.com/problems/queries-on-a-permutation-with-key/)|Queries on a Permutation With Key|中等|6018|79%|
|[1410](https://leetcode-cn.com/problems/html-entity-parser/)|HTML Entity Parser|中等|9170|49%|
|[1411](https://leetcode-cn.com/problems/number-of-ways-to-paint-n-x-3-grid/)|Number of Ways to Paint N × 3 Grid|困难|4773|57%|
|[1431](https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/)|Kids With the Greatest Number of Candies|简单|46180|86%|
|[1432](https://leetcode-cn.com/problems/max-difference-you-can-get-from-changing-an-integer/)|Max Difference You Can Get From Changing an Integer|中等|6453|38%|
|[1433](https://leetcode-cn.com/problems/check-if-a-string-can-break-another-string/)|Check If a String Can Break Another String|中等|4721|61%|
|[1434](https://leetcode-cn.com/problems/number-of-ways-to-wear-different-hats-to-each-other/)|Number of Ways to Wear Different Hats to Each Other|困难|3115|40%|
|[1417](https://leetcode-cn.com/problems/reformat-the-string/)|Reformat The String|简单|15080|50%|
|[1418](https://leetcode-cn.com/problems/display-table-of-food-orders-in-a-restaurant/)|Display Table of Food Orders in a Restaurant|中等|7326|57%|
|[1419](https://leetcode-cn.com/problems/minimum-number-of-frogs-croaking/)|Minimum Number of Frogs Croaking|中等|12549|38%|
|[1420](https://leetcode-cn.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/)|Build Array Where You Can Find The Maximum Exactly K Comparisons|困难|2716|57%|
|[1398](https://leetcode-cn.com/problems/customers-who-bought-products-a-and-b-but-not-c/)|Customers Who Bought Products A and B but Not C|中等|1081|80%|
|[1422](https://leetcode-cn.com/problems/maximum-score-after-splitting-a-string/)|Maximum Score After Splitting a String|简单|9740|52%|
|[1423](https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards/)|Maximum Points You Can Obtain from Cards|中等|10452|41%|
|[1424](https://leetcode-cn.com/problems/diagonal-traverse-ii/)|Diagonal Traverse II|中等|9037|36%|
|[1407](https://leetcode-cn.com/problems/top-travellers/)|Top Travellers|简单|1544|66%|
|[1446](https://leetcode-cn.com/problems/consecutive-characters/)|Consecutive Characters|简单|7486|58%|
|[1447](https://leetcode-cn.com/problems/simplified-fractions/)|Simplified Fractions|中等|4705|60%|
|[1448](https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree/)|Count Good Nodes in Binary Tree|中等|4269|70%|
|[1449](https://leetcode-cn.com/problems/form-largest-integer-with-digits-that-add-up-to-target/)|Form Largest Integer With Digits That Add up to Target|困难|4367|37%|
|[1412](https://leetcode-cn.com/problems/find-the-quiet-students-in-all-exams/)|Find the Quiet Students in All Exams|困难|619|58%|
|[1436](https://leetcode-cn.com/problems/destination-city/)|Destination City|简单|11092|78%|
|[1437](https://leetcode-cn.com/problems/check-if-all-1s-are-at-least-length-k-places-away/)|Check If All 1&#39;s Are at Least Length K Places Away|中等|7611|60%|
|[1438](https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)|Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit|中等|11428|38%|
|[1439](https://leetcode-cn.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/)|Find the Kth Smallest Sum of a Matrix With Sorted Rows|困难|3629|54%|
|[1421](https://leetcode-cn.com/problems/npv-queries/)|NPV Queries|中等|638|73%|
|[1441](https://leetcode-cn.com/problems/build-an-array-with-stack-operations/)|Build an Array With Stack Operations|简单|10520|64%|
|[1442](https://leetcode-cn.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/)|Count Triplets That Can Form Two Arrays of Equal XOR|中等|6370|63%|
|[1443](https://leetcode-cn.com/problems/minimum-time-to-collect-all-apples-in-a-tree/)|Minimum Time to Collect All Apples in a Tree|中等|7383|43%|
|[1444](https://leetcode-cn.com/problems/number-of-ways-of-cutting-a-pizza/)|Number of Ways of Cutting a Pizza|困难|3081|49%|
|[1460](https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/)|Make Two Arrays Equal by Reversing Sub-arrays|简单|7579|75%|
|[1461](https://leetcode-cn.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/)|Check If a String Contains All Binary Codes of Size K|中等|6250|46%|
|[1462](https://leetcode-cn.com/problems/course-schedule-iv/)|Course Schedule IV|中等|5989|38%|
|[1463](https://leetcode-cn.com/problems/cherry-pickup-ii/)|Cherry Pickup II|困难|2446|56%|
|[1450](https://leetcode-cn.com/problems/number-of-students-doing-homework-at-a-given-time/)|Number of Students Doing Homework at a Given Time|简单|12473|81%|
|[1451](https://leetcode-cn.com/problems/rearrange-words-in-a-sentence/)|Rearrange Words in a Sentence|中等|11053|48%|
|[1452](https://leetcode-cn.com/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list/)|People Whose List of Favorite Companies Is Not a Subset of Another List|中等|7971|46%|
|[1453](https://leetcode-cn.com/problems/maximum-number-of-darts-inside-of-a-circular-dartboard/)|Maximum Number of Darts Inside of a Circular Dartboard|困难|3353|34%|
|[1435](https://leetcode-cn.com/problems/create-a-session-bar-chart/)|Create a Session Bar Chart|简单|833|63%|
|[1440](https://leetcode-cn.com/problems/evaluate-boolean-expression/)|Evaluate Boolean Expression|中等|547|72%|
|[1455](https://leetcode-cn.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/)|Check If a Word Occurs As a Prefix of Any Word in a Sentence|简单|8508|64%|
|[1456](https://leetcode-cn.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/)|Maximum Number of Vowels in a Substring of Given Length|中等|10010|46%|
|[1457](https://leetcode-cn.com/problems/pseudo-palindromic-paths-in-a-binary-tree/)|Pseudo-Palindromic Paths in a Binary Tree|中等|6571|63%|
|[1458](https://leetcode-cn.com/problems/max-dot-product-of-two-subsequences/)|Max Dot Product of Two Subsequences|困难|6501|41%|
|[1475](https://leetcode-cn.com/problems/final-prices-with-a-special-discount-in-a-shop/)|Final Prices With a Special Discount in a Shop|简单|5837|74%|
|[1478](https://leetcode-cn.com/problems/allocate-mailboxes/)|Allocate Mailboxes|困难|1665|54%|
|[1476](https://leetcode-cn.com/problems/subrectangle-queries/)|Subrectangle Queries|中等|3166|91%|
|[1477](https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/)|Find Two Non-overlapping Sub-arrays Each With Target Sum|中等|7439|25%|
|[1464](https://leetcode-cn.com/problems/maximum-product-of-two-elements-in-an-array/)|Maximum Product of Two Elements in an Array|简单|11666|79%|
|[1465](https://leetcode-cn.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/)|Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts|中等|14365|28%|
|[1466](https://leetcode-cn.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/)|Reorder Routes to Make All Paths Lead to the City Zero|中等|6864|49%|
|[1467](https://leetcode-cn.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/)|Probability of a Two Boxes Having The Same Number of Distinct Balls|困难|1507|59%|
|[1445](https://leetcode-cn.com/problems/apples-oranges/)|Apples &amp; Oranges|中等|678|88%|
|[1454](https://leetcode-cn.com/problems/active-users/)|Active Users|中等|1389|42%|
|[1470](https://leetcode-cn.com/problems/shuffle-the-array/)|Shuffle the Array|简单|16032|87%|
|[1471](https://leetcode-cn.com/problems/the-k-strongest-values-in-an-array/)|The k Strongest Values in an Array|中等|8224|52%|
|[1472](https://leetcode-cn.com/problems/design-browser-history/)|Design Browser History|中等|5815|57%|
|[1473](https://leetcode-cn.com/problems/paint-house-iii/)|Paint House III|困难|3518|47%|
|[1491](https://leetcode-cn.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/)|Average Salary Excluding the Minimum and Maximum Salary|简单|6575|69%|
|[1492](https://leetcode-cn.com/problems/the-kth-factor-of-n/)|The kth Factor of n|中等|3880|70%|
|[1493](https://leetcode-cn.com/problems/longest-subarray-of-1s-after-deleting-one-element/)|Longest Subarray of 1&#39;s After Deleting One Element|中等|4869|54%|
|[1494](https://leetcode-cn.com/problems/parallel-courses-ii/)|Parallel Courses II|困难|3338|33%|
|[1480](https://leetcode-cn.com/problems/running-sum-of-1d-array/)|Running Sum of 1d Array|简单|19289|89%|
|[1481](https://leetcode-cn.com/problems/least-number-of-unique-integers-after-k-removals/)|Least Number of Unique Integers after K Removals|中等|9473|51%|
|[1482](https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/)|Minimum Number of Days to Make m Bouquets|中等|6874|41%|
|[1459](https://leetcode-cn.com/problems/rectangles-area/)|Rectangles Area|中等|476|48%|
|[1468](https://leetcode-cn.com/problems/calculate-salaries/)|Calculate Salaries|中等|398|75%|
|[1469](https://leetcode-cn.com/problems/find-all-the-lonely-nodes/)|Find All The Lonely Nodes|简单|448|80%|
|[1486](https://leetcode-cn.com/problems/xor-operation-in-an-array/)|XOR Operation in an Array|简单|10886|86%|
|[1487](https://leetcode-cn.com/problems/making-file-names-unique/)|Making File Names Unique|中等|14116|27%|
|[1488](https://leetcode-cn.com/problems/avoid-flood-in-the-city/)|Avoid Flood in The City|中等|13144|20%|
|[1489](https://leetcode-cn.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/)|Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree|困难|1364|47%|
|[1508](https://leetcode-cn.com/problems/range-sum-of-sorted-subarray-sums/)|Range Sum of Sorted Subarray Sums|中等|3041|74%|
|[1509](https://leetcode-cn.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/)|Minimum Difference Between Largest and Smallest Value in Three Moves|中等|3678|51%|
|[1510](https://leetcode-cn.com/problems/stone-game-iv/)|Stone Game IV|困难|3626|47%|
|[1474](https://leetcode-cn.com/problems/delete-n-nodes-after-m-nodes-of-a-linked-list/)|Delete N Nodes After M Nodes of a Linked List|简单|315|69%|
|[1496](https://leetcode-cn.com/problems/path-crossing/)|Path Crossing|简单|8280|53%|
|[1497](https://leetcode-cn.com/problems/check-if-array-pairs-are-divisible-by-k/)|Check If Array Pairs Are Divisible by k|中等|8305|39%|
|[1498](https://leetcode-cn.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/)|Number of Subsequences That Satisfy the Given Sum Condition|中等|6662|31%|
|[1499](https://leetcode-cn.com/problems/max-value-of-equation/)|Max Value of Equation|困难|3126|37%|
|[1479](https://leetcode-cn.com/problems/sales-by-day-of-the-week/)|Sales by Day of the Week|困难|205|61%|
|[1484](https://leetcode-cn.com/problems/clone-binary-tree-with-random-pointer/)|Clone Binary Tree With Random Pointer|中等|127|77%|
|[1485](https://leetcode-cn.com/problems/group-sold-products-by-the-date/)|Group Sold Products By The Date|简单|475|68%|
|[1502](https://leetcode-cn.com/problems/can-make-arithmetic-progression-from-sequence/)|[Can Make Arithmetic Progression From Sequence](./src/1502-can-make-arithmetic-progression-from-sequence/)|简单|10161|79%|
|[1503](https://leetcode-cn.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/)|[Last Moment Before All Ants Fall Out of a Plank](./src/1503-last-moment-before-all-ants-fall-out-of-a-plank/)|中等|10031|47%|
|[1504](https://leetcode-cn.com/problems/count-submatrices-with-all-ones/)|[Count Submatrices With All Ones](./src/1504-count-submatrices-with-all-ones/)|中等|6534|53%|
|[1505](https://leetcode-cn.com/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits/)|Minimum Possible Integer After at Most K Adjacent Swaps On Digits|困难|5627|33%|
|[1490](https://leetcode-cn.com/problems/clone-n-ary-tree/)|Clone N-ary Tree|中等|189|89%|
|[1512](https://leetcode-cn.com/problems/number-of-good-pairs/)|Number of Good Pairs|简单|9558|88%|
|[1513](https://leetcode-cn.com/problems/number-of-substrings-with-only-1s/)|Number of Substrings With Only 1s|中等|15337|35%|
|[1515](https://leetcode-cn.com/problems/best-position-for-a-service-centre/)|Best Position for a Service Centre|困难|5616|29%|
|[1495](https://leetcode-cn.com/problems/friendly-movies-streamed-last-month/)|Friendly Movies Streamed Last Month|简单|444|55%|
|[1500](https://leetcode-cn.com/problems/design-a-file-sharing-system/)|Design a File Sharing System|中等|69|52%|
|[1501](https://leetcode-cn.com/problems/countries-you-can-safely-invest-in/)|Countries You Can Safely Invest In|中等|192|68%|
|[1506](https://leetcode-cn.com/problems/find-root-of-n-ary-tree/)|Find Root of N-Ary Tree|中等|37|81%|
|[1511](https://leetcode-cn.com/problems/customer-order-frequency/)|Customer Order Frequency|简单|160|66%|
|[1516](https://leetcode-cn.com/problems/move-sub-tree-of-n-ary-tree/)|Move Sub-Tree of N-Ary Tree|困难|9|77%|
|[LCP 02](https://leetcode-cn.com/problems/deep-dark-fraction/)|Deep Dark Fraction|简单|12180|63%|
|[LCP 04](https://leetcode-cn.com/problems/broken-board-dominoes/)|Broken Board Dominoes|困难|4017|33%|
|[LCP 05](https://leetcode-cn.com/problems/coin-bonus/)|Coin Bonus|困难|6547|16%|
|[LCP 03](https://leetcode-cn.com/problems/programmable-robot/)|Programmable Robot|中等|32008|20%|
|[LCP 01](https://leetcode-cn.com/problems/guess-numbers/)|Guess Numbers|简单|60222|83%|
|[面试题 01.01](https://leetcode-cn.com/problems/is-unique-lcci/)|Is Unique LCCI|简单|35386|72%|
|[面试题 01.02](https://leetcode-cn.com/problems/check-permutation-lcci/)|Check Permutation LCCI|简单|26365|66%|
|[面试题 01.03](https://leetcode-cn.com/problems/string-to-url-lcci/)|String to URL LCCI|简单|22093|58%|
|[面试题 01.06](https://leetcode-cn.com/problems/compress-string-lcci/)|Compress String LCCI|简单|73181|48%|
|[面试题 01.09](https://leetcode-cn.com/problems/string-rotation-lcci/)|String Rotation LCCI|简单|17861|55%|
|[面试题 02.01](https://leetcode-cn.com/problems/remove-duplicate-node-lcci/)|Remove Duplicate Node LCCI|简单|41607|70%|
|[面试题 02.06](https://leetcode-cn.com/problems/palindrome-linked-list-lcci/)|Palindrome Linked List LCCI|简单|21463|46%|
|[面试题 02.07](https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/)|Intersection of Two Linked Lists LCCI|简单|16739|68%|
|[面试题 02.08](https://leetcode-cn.com/problems/linked-list-cycle-lcci/)|Linked List Cycle LCCI|中等|13918|50%|
|[面试题 03.02](https://leetcode-cn.com/problems/min-stack-lcci/)|Min Stack LCCI|简单|12733|59%|
|[面试题 03.04](https://leetcode-cn.com/problems/implement-queue-using-stacks-lcci/)|Implement Queue using Stacks LCCI|简单|9756|71%|
|[面试题 04.01](https://leetcode-cn.com/problems/route-between-nodes-lcci/)|Route Between Nodes LCCI|中等|8692|53%|
|[面试题 03.01](https://leetcode-cn.com/problems/three-in-one-lcci/)|Three in One LCCI|简单|5433|55%|
|[面试题 03.05](https://leetcode-cn.com/problems/sort-of-stacks-lcci/)|Sort of Stacks LCCI|中等|6664|53%|
|[面试题 04.02](https://leetcode-cn.com/problems/minimum-height-tree-lcci/)|Minimum Height Tree LCCI|简单|13550|79%|
|[面试题 04.03](https://leetcode-cn.com/problems/list-of-depth-lcci/)|List of Depth LCCI|中等|6562|80%|
|[面试题 04.04](https://leetcode-cn.com/problems/check-balance-lcci/)|Check Balance LCCI|简单|20640|55%|
|[面试题 04.05](https://leetcode-cn.com/problems/legal-binary-search-tree-lcci/)|Legal Binary Search Tree LCCI|中等|22326|33%|
|[面试题 04.06](https://leetcode-cn.com/problems/successor-lcci/)|Successor LCCI|中等|11386|57%|
|[面试题 04.08](https://leetcode-cn.com/problems/first-common-ancestor-lcci/)|First Common Ancestor LCCI|中等|7342|69%|
|[面试题 05.01](https://leetcode-cn.com/problems/insert-into-bits-lcci/)|Insert Into Bits LCCI|简单|4879|49%|
|[面试题 05.06](https://leetcode-cn.com/problems/convert-integer-lcci/)|Convert Integer LCCI|简单|6690|50%|
|[面试题 05.07](https://leetcode-cn.com/problems/exchange-lcci/)|Exchange LCCI|简单|4852|68%|
|[面试题 05.04](https://leetcode-cn.com/problems/closed-number-lcci/)|Closed Number LCCI|中等|3689|42%|
|[面试题 01.04](https://leetcode-cn.com/problems/palindrome-permutation-lcci/)|Palindrome Permutation LCCI|简单|21721|54%|
|[面试题 01.07](https://leetcode-cn.com/problems/rotate-matrix-lcci/)|Rotate Matrix LCCI|中等|36323|77%|
|[面试题 01.08](https://leetcode-cn.com/problems/zero-matrix-lcci/)|Zero Matrix LCCI|中等|15189|63%|
|[面试题 02.03](https://leetcode-cn.com/problems/delete-middle-node-lcci/)|Delete Middle Node LCCI|简单|23329|84%|
|[面试题 02.05](https://leetcode-cn.com/problems/sum-lists-lcci/)|Sum Lists LCCI|中等|20251|45%|
|[面试题 03.03](https://leetcode-cn.com/problems/stack-of-plates-lcci/)|Stack of Plates LCCI|中等|5716|37%|
|[面试题 05.08](https://leetcode-cn.com/problems/draw-line-lcci/)|Draw Line LCCI|中等|1934|54%|
|[面试题 08.01](https://leetcode-cn.com/problems/three-steps-problem-lcci/)|Three Steps Problem LCCI|简单|36134|34%|
|[面试题 08.04](https://leetcode-cn.com/problems/power-set-lcci/)|Power Set LCCI|中等|6995|80%|
|[面试题 08.05](https://leetcode-cn.com/problems/recursive-mulitply-lcci/)|Recursive Mulitply LCCI|中等|9011|67%|
|[面试题 08.09](https://leetcode-cn.com/problems/bracket-lcci/)|Bracket LCCI|中等|7356|81%|
|[面试题 08.10](https://leetcode-cn.com/problems/color-fill-lcci/)|Color Fill LCCI|简单|7682|54%|
|[面试题 08.13](https://leetcode-cn.com/problems/pile-box-lcci/)|Pile Box LCCI|困难|4018|48%|
|[面试题 05.02](https://leetcode-cn.com/problems/bianry-number-to-string-lcci/)|Bianry Number to String LCCI|中等|3655|60%|
|[面试题 03.06](https://leetcode-cn.com/problems/animal-shelter-lcci/)|Animal Shelter LCCI|简单|5053|59%|
|[面试题 04.10](https://leetcode-cn.com/problems/check-subtree-lcci/)|Check SubTree LCCI|中等|7369|70%|
|[面试题 05.03](https://leetcode-cn.com/problems/reverse-bits-lcci/)|Reverse Bits LCCI|简单|5891|49%|
|[面试题 08.11](https://leetcode-cn.com/problems/coin-lcci/)|Coin LCCI|中等|41815|49%|
|[面试题 10.03](https://leetcode-cn.com/problems/search-rotate-array-lcci/)|Search Rotate Array LCCI|中等|8595|38%|
|[面试题 08.12](https://leetcode-cn.com/problems/eight-queens-lcci/)|Eight Queens LCCI|困难|5234|74%|
|[面试题 08.03](https://leetcode-cn.com/problems/magic-index-lcci/)|Magic Index LCCI|简单|10619|64%|
|[面试题 08.07](https://leetcode-cn.com/problems/permutation-i-lcci/)|Permutation I LCCI|中等|8122|81%|
|[面试题 10.05](https://leetcode-cn.com/problems/sparse-array-search-lcci/)|Sparse Array Search LCCI|简单|7713|58%|
|[面试题 16.01](https://leetcode-cn.com/problems/swap-numbers-lcci/)|Swap Numbers LCCI|中等|9804|82%|
|[面试题 16.02](https://leetcode-cn.com/problems/words-frequency-lcci/)|Words Frequency LCCI|中等|5728|76%|
|[面试题 16.03](https://leetcode-cn.com/problems/intersection-lcci/)|Intersection LCCI|困难|18161|45%|
|[面试题 16.04](https://leetcode-cn.com/problems/tic-tac-toe-lcci/)|Tic-Tac-Toe LCCI|中等|6172|42%|
|[面试题 16.06](https://leetcode-cn.com/problems/smallest-difference-lcci/)|Smallest Difference LCCI|中等|9417|40%|
|[剑指 Offer 09](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)|用两个栈实现队列 LCOF|简单|86886|74%|
|[剑指 Offer 10- I](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/)|斐波那契数列  LCOF|简单|136609|33%|
|[剑指 Offer 03](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)|数组中重复的数字 LCOF|简单|125690|67%|
|[剑指 Offer 04](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)|二维数组中的查找 LCOF|简单|135881|40%|
|[剑指 Offer 10- II](https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/)|青蛙跳台阶问题  LCOF|简单|92605|41%|
|[剑指 Offer 11](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)|旋转数组的最小数字  LCOF|简单|95796|47%|
|[剑指 Offer 12](https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/)|矩阵中的路径  LCOF|中等|74751|44%|
|[剑指 Offer 05](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)|替换空格 LCOF|简单|73611|76%|
|[剑指 Offer 13](https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)|机器人的运动范围  LCOF|中等|97741|49%|
|[剑指 Offer 06](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)|从尾到头打印链表 LCOF|简单|77010|76%|
|[剑指 Offer 07](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)|重建二叉树 LCOF|中等|60123|68%|
|[剑指 Offer 14- I](https://leetcode-cn.com/problems/jian-sheng-zi-lcof/)|剪绳子  LCOF|中等|62679|54%|
|[剑指 Offer 14- II](https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/)|剪绳子 II LCOF|中等|58547|30%|
|[剑指 Offer 25](https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/)|合并两个排序的链表  LCOF|简单|47284|74%|
|[剑指 Offer 26](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/)|树的子结构  LCOF|中等|59276|46%|
|[剑指 Offer 27](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/)|二叉树的镜像  LCOF|简单|48459|78%|
|[剑指 Offer 28](https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/)|对称的二叉树  LCOF|简单|49735|57%|
|[剑指 Offer 20](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)|表示数值的字符串 LCOF|中等|61605|20%|
|[剑指 Offer 21](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/)|调整数组顺序使奇数位于偶数前面 LCOF|简单|59786|64%|
|[剑指 Offer 15](https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/)|二进制中1的个数 LCOF|简单|52645|73%|
|[剑指 Offer 29](https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/)|顺时针打印矩阵  LCOF|简单|94565|45%|
|[剑指 Offer 22](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)|链表中倒数第k个节点 LCOF|简单|58870|78%|
|[剑指 Offer 16](https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/)|数值的整数次方 LCOF|中等|81308|32%|
|[剑指 Offer 17](https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)|打印从1到最大的n位数 LCOF|简单|44878|78%|
|[剑指 Offer 19](https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/)|正则表达式匹配 LCOF|困难|42515|35%|
|[剑指 Offer 24](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)|反转链表 LCOF|简单|72627|75%|
|[剑指 Offer 18](https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/)|删除链表的节点 LCOF|简单|63808|58%|
|[剑指 Offer 35](https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/)|复杂链表的复制  LCOF|中等|29219|71%|
|[剑指 Offer 40](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)|最小的k个数  LCOF|简单|121707|57%|
|[剑指 Offer 30](https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/)|包含min函数的栈 LCOF|简单|43266|57%|
|[剑指 Offer 41](https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/)|数据流中的中位数  LCOF|困难|22905|55%|
|[剑指 Offer 42](https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/)|连续子数组的最大和  LCOF|简单|66983|59%|
|[剑指 Offer 36](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/)|二叉搜索树与双向链表  LCOF|中等|29110|63%|
|[剑指 Offer 31](https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/)|栈的压入、弹出序列 LCOF|中等|39005|58%|
|[剑指 Offer 37](https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof/)|序列化二叉树  LCOF|困难|24883|52%|
|[剑指 Offer 38](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)|字符串的排列  LCOF|中等|46217|54%|
|[剑指 Offer 43](https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/)|1～n整数中1出现的次数  LCOF|中等|26644|45%|
|[剑指 Offer 39](https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/)|数组中出现次数超过一半的数字  LCOF|简单|51285|67%|
|[剑指 Offer 32 - I](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/)|从上到下打印二叉树 LCOF|中等|40826|64%|
|[剑指 Offer 32 - II](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/)|从上到下打印二叉树 II LCOF|简单|39481|68%|
|[剑指 Offer 44](https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/)|数字序列中某一位的数字  LCOF|中等|29388|38%|
|[剑指 Offer 32 - III](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/)|从上到下打印二叉树 III LCOF|中等|40385|58%|
|[剑指 Offer 33](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/)|二叉搜索树的后序遍历序列 LCOF|中等|41643|51%|
|[剑指 Offer 50](https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/)|第一个只出现一次的字符  LCOF|简单|49582|60%|
|[剑指 Offer 34](https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/)|二叉树中和为某一值的路径 LCOF|中等|46842|56%|
|[剑指 Offer 51](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)|数组中的逆序对  LCOF|困难|63868|45%|
|[剑指 Offer 55 - I](https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/)|二叉树的深度 LCOF|简单|46546|78%|
|[剑指 Offer 56 - I](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/)|数组中数字出现的次数 LCOF|中等|51183|71%|
|[剑指 Offer 56 - II](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/)|数组中数字出现的次数 II LCOF|中等|23576|79%|
|[剑指 Offer 57](https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/)|和为s的两个数字 LCOF|简单|41728|65%|
|[剑指 Offer 45](https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/)|把数组排成最小的数 LCOF|中等|35687|55%|
|[剑指 Offer 57 - II](https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/)|和为s的连续正数序列 LCOF|简单|62828|68%|
|[剑指 Offer 46](https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/)|把数字翻译成字符串 LCOF|中等|64879|53%|
|[剑指 Offer 52](https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/)|两个链表的第一个公共节点  LCOF|简单|44771|63%|
|[剑指 Offer 47](https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/)|礼物的最大价值 LCOF|中等|35392|67%|
|[剑指 Offer 58 - I](https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/)|翻转单词顺序 LCOF|简单|59428|42%|
|[剑指 Offer 53 - I](https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/)|在排序数组中查找数字  LCOF|简单|55562|53%|
|[剑指 Offer 58 - II](https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/)|左旋转字符串 LCOF|简单|53844|85%|
|[剑指 Offer 53 - II](https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/)|缺失的数字  LCOF|简单|78438|44%|
|[剑指 Offer 48](https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/)|最长不含重复字符的子字符串 LCOF|中等|55363|45%|
|[剑指 Offer 54](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/)|二叉搜索树的第k大节点  LCOF|简单|41610|73%|
|[剑指 Offer 49](https://leetcode-cn.com/problems/chou-shu-lcof/)|丑数 LCOF|中等|28173|64%|
|[剑指 Offer 65](https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/)|不用加减乘除做加法 LCOF|简单|25889|55%|
|[剑指 Offer 59 - I](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)|滑动窗口的最大值 LCOF|简单|64823|44%|
|[剑指 Offer 59 - II](https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/)|队列的最大值 LCOF|中等|62712|48%|
|[剑指 Offer 66](https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/)|构建乘积数组 LCOF|简单|27518|58%|
|[剑指 Offer 60](https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/)|n个骰子的点数  LCOF|简单|26643|53%|
|[剑指 Offer 67](https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/)|把字符串转换成整数 LCOF|中等|49199|27%|
|[剑指 Offer 61](https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/)|扑克牌中的顺子  LCOF|简单|44395|44%|
|[剑指 Offer 55 - II](https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/)|平衡二叉树 LCOF|简单|45881|58%|
|[剑指 Offer 62](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/)|圆圈中最后剩下的数字 LCOF|简单|60839|61%|
|[剑指 Offer 63](https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/)|股票的最大利润  LCOF|中等|40303|64%|
|[剑指 Offer 64](https://leetcode-cn.com/problems/qiu-12n-lcof/)|求1&#43;2&#43;…&#43;n LCOF|中等|62190|85%|
|[面试题68 - I](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/)|二叉搜索树的最近公共祖先 LCOF|简单|34488|68%|
|[面试题68 - II](https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/)|二叉树的最近公共祖先 LCOF|简单|32826|68%|
|[面试题 08.08](https://leetcode-cn.com/problems/permutation-ii-lcci/)|Permutation II LCCI|中等|6890|70%|
|[面试题 16.07](https://leetcode-cn.com/problems/maximum-lcci/)|Maximum LCCI|简单|8532|72%|
|[面试题 16.09](https://leetcode-cn.com/problems/operations-lcci/)|Operations LCCI|中等|1316|48%|
|[面试题 16.10](https://leetcode-cn.com/problems/living-people-lcci/)|Living People LCCI|中等|4801|67%|
|[面试题 16.11](https://leetcode-cn.com/problems/diving-board-lcci/)|Diving Board LCCI|简单|78454|44%|
|[面试题 16.13](https://leetcode-cn.com/problems/bisect-squares-lcci/)|Bisect Squares LCCI|中等|1414|42%|
|[面试题 16.14](https://leetcode-cn.com/problems/best-line-lcci/)|Best Line LCCI|中等|1618|50%|
|[面试题 16.15](https://leetcode-cn.com/problems/master-mind-lcci/)|Master Mind LCCI|简单|4779|49%|
|[面试题 16.16](https://leetcode-cn.com/problems/sub-sort-lcci/)|Sub Sort LCCI|中等|9297|43%|
|[面试题 16.17](https://leetcode-cn.com/problems/contiguous-sequence-lcci/)|Contiguous Sequence LCCI|简单|17834|59%|
|[面试题 16.18](https://leetcode-cn.com/problems/pattern-matching-lcci/)|Pattern Matching LCCI|中等|34545|34%|
|[面试题 16.19](https://leetcode-cn.com/problems/pond-sizes-lcci/)|Pond Sizes LCCI|中等|7684|59%|
|[面试题 01.05](https://leetcode-cn.com/problems/one-away-lcci/)|One Away LCCI|中等|29180|32%|
|[面试题 02.02](https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/)|Kth Node From End of List LCCI|简单|23643|79%|
|[面试题 02.04](https://leetcode-cn.com/problems/partition-list-lcci/)|Partition List LCCI|中等|12000|63%|
|[面试题 04.12](https://leetcode-cn.com/problems/paths-with-sum-lcci/)|Paths with Sum LCCI|中等|9012|48%|
|[面试题 04.09](https://leetcode-cn.com/problems/bst-sequences-lcci/)|BST Sequences LCCI|困难|4326|44%|
|[面试题 08.02](https://leetcode-cn.com/problems/robot-in-a-grid-lcci/)|Robot in a Grid LCCI|中等|9791|34%|
|[面试题 10.01](https://leetcode-cn.com/problems/sorted-merge-lcci/)|Sorted Merge LCCI|简单|51439|54%|
|[面试题 10.11](https://leetcode-cn.com/problems/peaks-and-valleys-lcci/)|Peaks and Valleys LCCI|中等|3280|65%|
|[面试题 08.06](https://leetcode-cn.com/problems/hanota-lcci/)|Hanota LCCI|简单|9602|64%|
|[面试题 10.09](https://leetcode-cn.com/problems/sorted-matrix-search-lcci/)|Sorted Matrix Search LCCI|中等|8060|45%|
|[面试题 16.05](https://leetcode-cn.com/problems/factorial-zeros-lcci/)|Factorial Zeros LCCI|简单|6783|46%|
|[面试题 16.08](https://leetcode-cn.com/problems/english-int-lcci/)|English Int LCCI|困难|2034|36%|
|[面试题 17.12](https://leetcode-cn.com/problems/binode-lcci/)|BiNode LCCI|简单|9520|59%|
|[面试题 17.13](https://leetcode-cn.com/problems/re-space-lcci/)|Re-Space LCCI|中等|34079|56%|
|[面试题 17.14](https://leetcode-cn.com/problems/smallest-k-lcci/)|Smallest K LCCI|中等|15445|55%|
|[面试题 17.15](https://leetcode-cn.com/problems/longest-word-lcci/)|Longest Word LCCI|中等|4441|36%|
|[面试题 17.16](https://leetcode-cn.com/problems/the-masseuse-lcci/)|The Masseuse LCCI|简单|58138|53%|
|[面试题 17.17](https://leetcode-cn.com/problems/multi-search-lcci/)|Multi Search LCCI|中等|5579|42%|
|[面试题 17.01](https://leetcode-cn.com/problems/add-without-plus-lcci/)|Add Without Plus LCCI|简单|6901|54%|
|[面试题 16.25](https://leetcode-cn.com/problems/lru-cache-lcci/)|LRU Cache LCCI|中等|8199|54%|
|[面试题 16.26](https://leetcode-cn.com/problems/calculator-lcci/)|Calculator LCCI|中等|6149|37%|
|[面试题 17.20](https://leetcode-cn.com/problems/continuous-median-lcci/)|Continuous Median LCCI|困难|1996|54%|
|[面试题 17.21](https://leetcode-cn.com/problems/volume-of-histogram-lcci/)|Volume of Histogram LCCI|困难|4417|64%|
|[面试题 17.22](https://leetcode-cn.com/problems/word-transformer-lcci/)|Word Transformer LCCI|中等|7720|33%|
|[面试题 08.14](https://leetcode-cn.com/problems/boolean-evaluation-lcci/)|Boolean Evaluation LCCI|中等|2380|49%|
|[面试题 17.04](https://leetcode-cn.com/problems/missing-number-lcci/)|Missing Number LCCI|简单|13079|64%|
|[面试题 17.05](https://leetcode-cn.com/problems/find-longest-subarray-lcci/)|Find Longest Subarray LCCI|中等|4535|36%|
|[面试题 17.06](https://leetcode-cn.com/problems/number-of-2s-in-range-lcci/)|Number Of 2s In Range LCCI|中等|3922|41%|
|[面试题 17.07](https://leetcode-cn.com/problems/baby-names-lcci/)|Baby Names LCCI|中等|5737|38%|
|[面试题 17.08](https://leetcode-cn.com/problems/circus-tower-lcci/)|Circus Tower LCCI|中等|12702|24%|
|[面试题 17.09](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)|Get Kth Magic Number LCCI|中等|7345|53%|
|[面试题 17.10](https://leetcode-cn.com/problems/find-majority-element-lcci/)|Find Majority Element LCCI|简单|12336|64%|
|[面试题 17.11](https://leetcode-cn.com/problems/find-closest-lcci/)|Find Closest LCCI|中等|7273|67%|
|[面试题 10.02](https://leetcode-cn.com/problems/group-anagrams-lcci/)|Group Anagrams LCCI|中等|4932|65%|
|[面试题 10.10](https://leetcode-cn.com/problems/rank-from-stream-lcci/)|Rank from Stream LCCI|中等|2132|60%|
|[面试题 16.24](https://leetcode-cn.com/problems/pairs-with-sum-lcci/)|Pairs With Sum LCCI|中等|7815|46%|
|[面试题 17.18](https://leetcode-cn.com/problems/shortest-supersequence-lcci/)|Shortest Supersequence LCCI|中等|4308|43%|
|[面试题 17.19](https://leetcode-cn.com/problems/missing-two-lcci/)|Missing Two LCCI|困难|4524|55%|
|[面试题 17.23](https://leetcode-cn.com/problems/max-black-square-lcci/)|Max Black Square LCCI|中等|3380|35%|
|[面试题 17.24](https://leetcode-cn.com/problems/max-submatrix-lcci/)|Max Submatrix LCCI|困难|2466|51%|
|[面试题 16.20](https://leetcode-cn.com/problems/t9-lcci/)|T9 LCCI|中等|4129|71%|
|[面试题 16.21](https://leetcode-cn.com/problems/sum-swap-lcci/)|Sum Swap LCCI|中等|6356|44%|
|[面试题 17.25](https://leetcode-cn.com/problems/word-rectangle-lcci/)|Word Rectangle LCCI|困难|812|44%|
|[面试题 16.22](https://leetcode-cn.com/problems/langtons-ant-lcci/)|Langtons Ant LCCI|中等|1366|57%|
|[面试题 17.26](https://leetcode-cn.com/problems/sparse-similarity-lcci/)|Sparse Similarity LCCI|困难|3088|31%|
|[LCP 06](https://leetcode-cn.com/problems/na-ying-bi/)|拿硬币|简单|15343|82%|
|[LCP 08](https://leetcode-cn.com/problems/ju-qing-hong-fa-shi-jian/)|剧情触发时间|中等|12531|22%|
|[LCP 14](https://leetcode-cn.com/problems/qie-fen-shu-zu/)|切分数组|困难|6712|14%|
|[LCP 10](https://leetcode-cn.com/problems/er-cha-shu-ren-wu-diao-du/)|二叉树任务调度|困难|1360|50%|
|[LCP 07](https://leetcode-cn.com/problems/chuan-di-xin-xi/)|传递信息|简单|8047|57%|
|[LCP 16](https://leetcode-cn.com/problems/you-le-yuan-de-you-lan-ji-hua/)|游乐园的游览计划|困难|909|22%|
|[LCP 11](https://leetcode-cn.com/problems/qi-wang-ge-shu-tong-ji/)|期望个数统计|简单|4660|58%|
|[LCP 12](https://leetcode-cn.com/problems/xiao-zhang-shua-ti-ji-hua/)|小张刷题计划|中等|6014|29%|
|[LCP 15](https://leetcode-cn.com/problems/you-le-yuan-de-mi-gong/)|游乐园的迷宫|困难|703|49%|
|[LCP 09](https://leetcode-cn.com/problems/zui-xiao-tiao-yue-ci-shu/)|最小跳跃次数|困难|13937|20%|
|[LCP 13](https://leetcode-cn.com/problems/xun-bao/)|寻宝|困难|2165|20%|


---

#### 相似题型

|#|标题|
|:-:|:-|
|[92](https://leetcode-cn.com/problems/reverse-linked-list-ii/)|Reverse Linked List II|
|[206](https://leetcode-cn.com/problems/reverse-linked-list/)|Reverse Linked List|
|||
|[153](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)|Find Minimum in Rotated Sorted Array|
|[154](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)|Find Minimum in Rotated Sorted Array II|
|||
|[151](https://leetcode-cn.com/problems/reverse-words-in-a-string/)|Reverse Words in a String|
|[186](https://leetcode-cn.com/problems/reverse-words-in-a-string-ii/)|Reverse Words in a String II|
|[557](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)|Reverse Words in a String III|
|||
|[252](https://leetcode-cn.com/problems/meeting-rooms/)|Meeting Rooms|
|[253](https://leetcode-cn.com/problems/meeting-rooms-ii/)|Meeting Rooms II|
|||
|[496](https://leetcode-cn.com/problems/next-greater-element-i/)|Next Greater Element I|
|[503](https://leetcode-cn.com/problems/next-greater-element-ii/)|[Next Greater Element II](./src/0503-next-greater-element-ii/)|
|[556](https://leetcode-cn.com/problems/next-greater-element-iii/)|Next Greater Element III|
|||
|[521](https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/)|Longest Uncommon Subsequence I |
|[522](https://leetcode-cn.com/problems/longest-uncommon-subsequence-ii/)|Longest Uncommon Subsequence II|
|||
|[263](https://leetcode-cn.com/problems/ugly-number/)|Ugly Number|
|[264](https://leetcode-cn.com/problems/ugly-number-ii/)|Ugly Number II|
|[1201](https://leetcode-cn.com/problems/ugly-number-iii/)|Ugly Number III|
|||
|[2](https://leetcode-cn.com/problems/add-two-numbers/)|[Add Two Numbers](./src/0002-add-two-numbers/)|
|[445](https://leetcode-cn.com/problems/add-two-numbers-ii/)|Add Two Numbers II|
|||
|[125](https://leetcode-cn.com/problems/valid-palindrome/)|[Valid Palindrome](./src/0125-valid-palindrome/)|
|[680](https://leetcode-cn.com/problems/valid-palindrome-ii/)|Valid Palindrome II|
|[1216](https://leetcode-cn.com/problems/valid-palindrome-iii/)|Valid Palindrome III|
|||
|[905](https://leetcode-cn.com/problems/sort-array-by-parity/)|Sort Array By Parity|
|[922](https://leetcode-cn.com/problems/sort-array-by-parity-ii/)|Sort Array By Parity II|
|||
|[1113](https://leetcode-cn.com/problems/reported-posts/)|Reported Posts|
|[1132](https://leetcode-cn.com/problems/reported-posts-ii/)|Reported Posts II|
|||
|[498](https://leetcode-cn.com/problems/diagonal-traverse/)|Diagonal Traverse|
|[1424](https://leetcode-cn.com/problems/diagonal-traverse-ii/)|Diagonal Traverse II|
|||
|[223](https://leetcode-cn.com/problems/rectangle-area/)|Rectangle Area|
|[850](https://leetcode-cn.com/problems/rectangle-area-ii/)|Rectangle Area II|
|||
|[1046](https://leetcode-cn.com/problems/last-stone-weight/)|Last Stone Weight|
|[1049](https://leetcode-cn.com/problems/last-stone-weight-ii/)|Last Stone Weight II|
|||
|[112](https://leetcode-cn.com/problems/path-sum/)|[Path Sum](./src/0112-path-sum/)|
|[113](https://leetcode-cn.com/problems/path-sum-ii/)|Path Sum II|
|[437](https://leetcode-cn.com/problems/path-sum-iii/)|Path Sum III|
|[666](https://leetcode-cn.com/problems/path-sum-iv/)|Path Sum IV|
|||
|[116](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/)|Populating Next Right Pointers in Each Node|
|[117](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)|Populating Next Right Pointers in Each Node II|
|||
|[217](https://leetcode-cn.com/problems/contains-duplicate/)|Contains Duplicate|
|[219](https://leetcode-cn.com/problems/contains-duplicate-ii/)|Contains Duplicate II|
|[220](https://leetcode-cn.com/problems/contains-duplicate-iii/)|Contains Duplicate III|
|||
|[200](https://leetcode-cn.com/problems/number-of-islands/)|Number of Islands|
|[305](https://leetcode-cn.com/problems/number-of-islands-ii/)|Number of Islands II|
|||
|[285](https://leetcode-cn.com/problems/inorder-successor-in-bst/)|Inorder Successor in BST|
|[510](https://leetcode-cn.com/problems/inorder-successor-in-bst-ii/)|Inorder Successor in BST II|
|||
|[551](https://leetcode-cn.com/problems/student-attendance-record-i/)|Student Attendance Record I|
|[552](https://leetcode-cn.com/problems/student-attendance-record-ii/)|Student Attendance Record II|
|||
|[694](https://leetcode-cn.com/problems/number-of-distinct-islands/)|Number of Distinct Islands|
|[711](https://leetcode-cn.com/problems/number-of-distinct-islands-ii/)|Number of Distinct Islands II|
|||
|[654](https://leetcode-cn.com/problems/maximum-binary-tree/)|Maximum Binary Tree|
|[998](https://leetcode-cn.com/problems/maximum-binary-tree-ii/)|Maximum Binary Tree II|
|||
|[1148](https://leetcode-cn.com/problems/article-views-i/)|Article Views I|
|[1149](https://leetcode-cn.com/problems/article-views-ii/)|Article Views II|
|||
|[54](https://leetcode-cn.com/problems/spiral-matrix/)|Spiral Matrix|
|[59](https://leetcode-cn.com/problems/spiral-matrix-ii/)|Spiral Matrix II|
|[885](https://leetcode-cn.com/problems/spiral-matrix-iii/)|Spiral Matrix III|
|||
|[274](https://leetcode-cn.com/problems/h-index/)|H-Index|
|[275](https://leetcode-cn.com/problems/h-index-ii/)|H-Index II|
|||
|[370](https://leetcode-cn.com/problems/range-addition/)|Range Addition|
|[598](https://leetcode-cn.com/problems/range-addition-ii/)|Range Addition II|
|||
|[643](https://leetcode-cn.com/problems/maximum-average-subarray-i/)|Maximum Average Subarray I|
|[644](https://leetcode-cn.com/problems/maximum-average-subarray-ii/)|Maximum Average Subarray II|
|||
|[1173](https://leetcode-cn.com/problems/immediate-food-delivery-i/)|Immediate Food Delivery I|
|[1174](https://leetcode-cn.com/problems/immediate-food-delivery-ii/)|Immediate Food Delivery II|
|||
|[374](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)|[Guess Number Higher or Lower](./src/0374-guess-number-higher-or-lower/)|
|[375](https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/)|Guess Number Higher or Lower II|
|||
|[684](https://leetcode-cn.com/problems/redundant-connection/)|Redundant Connection|
|[685](https://leetcode-cn.com/problems/redundant-connection-ii/)|Redundant Connection II|
|||
|[39](https://leetcode-cn.com/problems/combination-sum/)|[Combination Sum](./src/0039-combination-sum/)|
|[40](https://leetcode-cn.com/problems/combination-sum-ii/)|[Combination Sum II](./src/0040-combination-sum-ii/)|
|[216](https://leetcode-cn.com/problems/combination-sum-iii/)|Combination Sum III|
|[377](https://leetcode-cn.com/problems/combination-sum-iv/)|Combination Sum IV|
|||
|[33](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)|[Search in Rotated Sorted Array](./src/0033-search-in-rotated-sorted-array/)|
|[81](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)|Search in Rotated Sorted Array II|
|||
|[126](https://leetcode-cn.com/problems/word-ladder-ii/)|Word Ladder II|
|[127](https://leetcode-cn.com/problems/word-ladder/)|Word Ladder|
|||
|[141](https://leetcode-cn.com/problems/linked-list-cycle/)|[Linked List Cycle](./src/0141-linked-list-cycle/)|
|[142](https://leetcode-cn.com/problems/linked-list-cycle-ii/)|Linked List Cycle II|
|||
|[256](https://leetcode-cn.com/problems/paint-house/)|Paint House|
|[265](https://leetcode-cn.com/problems/paint-house-ii/)|Paint House II|
|[1473](https://leetcode-cn.com/problems/paint-house-iii/)|Paint House III|
|||
|[266](https://leetcode-cn.com/problems/palindrome-permutation/)|Palindrome Permutation|
|[267](https://leetcode-cn.com/problems/palindrome-permutation-ii/)|Palindrome Permutation II|
|||
|[734](https://leetcode-cn.com/problems/sentence-similarity/)|Sentence Similarity|
|[737](https://leetcode-cn.com/problems/sentence-similarity-ii/)|Sentence Similarity II|
|||
|[939](https://leetcode-cn.com/problems/minimum-area-rectangle/)|Minimum Area Rectangle|
|[963](https://leetcode-cn.com/problems/minimum-area-rectangle-ii/)|Minimum Area Rectangle II|
|||
|[1075](https://leetcode-cn.com/problems/project-employees-i/)|Project Employees I|
|[1076](https://leetcode-cn.com/problems/project-employees-ii/)|Project Employees II|
|[1077](https://leetcode-cn.com/problems/project-employees-iii/)|Project Employees III|
|||
|[62](https://leetcode-cn.com/problems/unique-paths/)|Unique Paths|
|[63](https://leetcode-cn.com/problems/unique-paths-ii/)|Unique Paths II|
|[980](https://leetcode-cn.com/problems/unique-paths-iii/)|Unique Paths III|
|||
|[26](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)|[Remove Duplicates from Sorted Array](./src/0026-remove-duplicates-from-sorted-array/)|
|[80](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/)|Remove Duplicates from Sorted Array II|
|||
|[290](https://leetcode-cn.com/problems/word-pattern/)|Word Pattern|
|[291](https://leetcode-cn.com/problems/word-pattern-ii/)|Word Pattern II|
|||
|[18](https://leetcode-cn.com/problems/4sum/)|[4Sum](./src/0018-4sum/)|
|[454](https://leetcode-cn.com/problems/4sum-ii/)|4Sum II|
|||
|[908](https://leetcode-cn.com/problems/smallest-range-i/)|Smallest Range I|
|[910](https://leetcode-cn.com/problems/smallest-range-ii/)|Smallest Range II|
|||
|[924](https://leetcode-cn.com/problems/minimize-malware-spread/)|Minimize Malware Spread|
|[928](https://leetcode-cn.com/problems/minimize-malware-spread-ii/)|Minimize Malware Spread II|
|||
|[136](https://leetcode-cn.com/problems/single-number/)|[Single Number](./src/0136-single-number/)|
|[137](https://leetcode-cn.com/problems/single-number-ii/)|Single Number II|
|[260](https://leetcode-cn.com/problems/single-number-iii/)|Single Number III|
|||
|[224](https://leetcode-cn.com/problems/basic-calculator/)|Basic Calculator|
|[227](https://leetcode-cn.com/problems/basic-calculator-ii/)|Basic Calculator II|
|[770](https://leetcode-cn.com/problems/basic-calculator-iv/)|Basic Calculator IV|
|[772](https://leetcode-cn.com/problems/basic-calculator-iii/)|Basic Calculator III|
|||
|[270](https://leetcode-cn.com/problems/closest-binary-search-tree-value/)|Closest Binary Search Tree Value|
|[272](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii/)|Closest Binary Search Tree Value II|
|||
|[42](https://leetcode-cn.com/problems/trapping-rain-water/)|[Trapping Rain Water](./src/0042-trapping-rain-water/)|
|[407](https://leetcode-cn.com/problems/trapping-rain-water-ii/)|Trapping Rain Water II|
|||
|[453](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/)|Minimum Moves to Equal Array Elements|
|[462](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/)|Minimum Moves to Equal Array Elements II|
|||
|[1056](https://leetcode-cn.com/problems/confusing-number/)|Confusing Number|
|[1088](https://leetcode-cn.com/problems/confusing-number-ii/)|Confusing Number II|
|||
|[1047](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/)|Remove All Adjacent Duplicates In String|
|[1209](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii/)|Remove All Adjacent Duplicates in String II|
|||
|[45](https://leetcode-cn.com/problems/jump-game-ii/)|Jump Game II|
|[55](https://leetcode-cn.com/problems/jump-game/)|Jump Game|
|[1306](https://leetcode-cn.com/problems/jump-game-iii/)|Jump Game III|
|[1345](https://leetcode-cn.com/problems/jump-game-iv/)|Jump Game IV|
|[1340](https://leetcode-cn.com/problems/jump-game-v/)|Jump Game V|
|||
|[102](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)|Binary Tree Level Order Traversal|
|[107](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)|[Binary Tree Level Order Traversal II](./src/0107-binary-tree-level-order-traversal-ii/)|
|||
|[198](https://leetcode-cn.com/problems/house-robber/)|[House Robber](./src/0198-house-robber/)|
|[213](https://leetcode-cn.com/problems/house-robber-ii/)|House Robber II|
|[337](https://leetcode-cn.com/problems/house-robber-iii/)|House Robber III|
|||
|[246](https://leetcode-cn.com/problems/strobogrammatic-number/)|Strobogrammatic Number|
|[247](https://leetcode-cn.com/problems/strobogrammatic-number-ii/)|Strobogrammatic Number II|
|[248](https://leetcode-cn.com/problems/strobogrammatic-number-iii/)|Strobogrammatic Number III|
|||
|[349](https://leetcode-cn.com/problems/intersection-of-two-arrays/)|Intersection of Two Arrays|
|[350](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)|Intersection of Two Arrays II|
|||
|[1033](https://leetcode-cn.com/problems/moving-stones-until-consecutive/)|Moving Stones Until Consecutive|
|[1040](https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/)|Moving Stones Until Consecutive II|
|||
|[95](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/)|Unique Binary Search Trees II|
|[96](https://leetcode-cn.com/problems/unique-binary-search-trees/)|Unique Binary Search Trees|
|||
|[741](https://leetcode-cn.com/problems/cherry-pickup/)|Cherry Pickup|
|[1463](https://leetcode-cn.com/problems/cherry-pickup-ii/)|Cherry Pickup II|
|||
|[46](https://leetcode-cn.com/problems/permutations/)|Permutations|
|[47](https://leetcode-cn.com/problems/permutations-ii/)|Permutations II|
|||
|[169](https://leetcode-cn.com/problems/majority-element/)|Majority Element|
|[229](https://leetcode-cn.com/problems/majority-element-ii/)|Majority Element II|
|||
|[293](https://leetcode-cn.com/problems/flip-game/)|Flip Game|
|[294](https://leetcode-cn.com/problems/flip-game-ii/)|Flip Game II|
|||
|[339](https://leetcode-cn.com/problems/nested-list-weight-sum/)|Nested List Weight Sum|
|[364](https://leetcode-cn.com/problems/nested-list-weight-sum-ii/)|Nested List Weight Sum II|
|||
|[531](https://leetcode-cn.com/problems/lonely-pixel-i/)|Lonely Pixel I|
|[533](https://leetcode-cn.com/problems/lonely-pixel-ii/)|Lonely Pixel II|
|||
|[115](https://leetcode-cn.com/problems/distinct-subsequences/)|Distinct Subsequences|
|[940](https://leetcode-cn.com/problems/distinct-subsequences-ii/)|Distinct Subsequences II|
|||
|[78](https://leetcode-cn.com/problems/subsets/)|Subsets|
|[90](https://leetcode-cn.com/problems/subsets-ii/)|Subsets II|
|||
|[118](https://leetcode-cn.com/problems/pascals-triangle/)|[Pascal&#39;s Triangle](./src/0118-pascals-triangle/)|
|[119](https://leetcode-cn.com/problems/pascals-triangle-ii/)|[Pascal&#39;s Triangle II](./src/0119-pascals-triangle-ii/)|
|||
|[490](https://leetcode-cn.com/problems/the-maze/)|The Maze|
|[499](https://leetcode-cn.com/problems/the-maze-iii/)|The Maze III|
|[505](https://leetcode-cn.com/problems/the-maze-ii/)|The Maze II|
|||
|[298](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence/)|Binary Tree Longest Consecutive Sequence|
|[549](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence-ii/)|Binary Tree Longest Consecutive Sequence II|
|||
|[319](https://leetcode-cn.com/problems/bulb-switcher/)|Bulb Switcher|
|[672](https://leetcode-cn.com/problems/bulb-switcher-ii/)|Bulb Switcher II|
|[1375](https://leetcode-cn.com/problems/bulb-switcher-iii/)|Bulb Switcher III|
|||
|[729](https://leetcode-cn.com/problems/my-calendar-i/)|My Calendar I|
|[731](https://leetcode-cn.com/problems/my-calendar-ii/)|My Calendar II|
|[732](https://leetcode-cn.com/problems/my-calendar-iii/)|My Calendar III|
|||
|[1158](https://leetcode-cn.com/problems/market-analysis-i/)|Market Analysis I|
|[1159](https://leetcode-cn.com/problems/market-analysis-ii/)|Market Analysis II|
|||
|[1193](https://leetcode-cn.com/problems/monthly-transactions-i/)|Monthly Transactions I|
|[1205](https://leetcode-cn.com/problems/monthly-transactions-ii/)|Monthly Transactions II|
|||
|[82](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)|[Remove Duplicates from Sorted List II](./src/0082-remove-duplicates-from-sorted-list-ii/)|
|[83](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)|[Remove Duplicates from Sorted List](./src/0083-remove-duplicates-from-sorted-list/)|
|||
|[121](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|[Best Time to Buy and Sell Stock](./src/0121-best-time-to-buy-and-sell-stock/)|
|[122](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|[Best Time to Buy and Sell Stock II](./src/0122-best-time-to-buy-and-sell-stock-ii/)|
|[123](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|Best Time to Buy and Sell Stock III|
|[188](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)|Best Time to Buy and Sell Stock IV|
|||
|[131](https://leetcode-cn.com/problems/palindrome-partitioning/)|Palindrome Partitioning|
|[132](https://leetcode-cn.com/problems/palindrome-partitioning-ii/)|Palindrome Partitioning II|
|[1278](https://leetcode-cn.com/problems/palindrome-partitioning-iii/)|Palindrome Partitioning III|
|||
|[768](https://leetcode-cn.com/problems/max-chunks-to-make-sorted-ii/)|Max Chunks To Make Sorted II|
|[769](https://leetcode-cn.com/problems/max-chunks-to-make-sorted/)|Max Chunks To Make Sorted|
|||
|[1057](https://leetcode-cn.com/problems/campus-bikes/)|Campus Bikes|
|[1066](https://leetcode-cn.com/problems/campus-bikes-ii/)|Campus Bikes II|
|||
|[1087](https://leetcode-cn.com/problems/brace-expansion/)|Brace Expansion|
|[1096](https://leetcode-cn.com/problems/brace-expansion-ii/)|Brace Expansion II|
|||
|[1136](https://leetcode-cn.com/problems/parallel-courses/)|Parallel Courses|
|[1494](https://leetcode-cn.com/problems/parallel-courses-ii/)|Parallel Courses II|
|||
|[207](https://leetcode-cn.com/problems/course-schedule/)|Course Schedule|
|[210](https://leetcode-cn.com/problems/course-schedule-ii/)|Course Schedule II|
|[630](https://leetcode-cn.com/problems/course-schedule-iii/)|Course Schedule III|
|[1462](https://leetcode-cn.com/problems/course-schedule-iv/)|Course Schedule IV|
|||
|[344](https://leetcode-cn.com/problems/reverse-string/)|Reverse String|
|[541](https://leetcode-cn.com/problems/reverse-string-ii/)|Reverse String II|
|||
|[526](https://leetcode-cn.com/problems/beautiful-arrangement/)|Beautiful Arrangement|
|[667](https://leetcode-cn.com/problems/beautiful-arrangement-ii/)|Beautiful Arrangement II|
|||
|[944](https://leetcode-cn.com/problems/delete-columns-to-make-sorted/)|Delete Columns to Make Sorted|
|[955](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-ii/)|Delete Columns to Make Sorted II|
|[960](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-iii/)|Delete Columns to Make Sorted III|
|||
|[511](https://leetcode-cn.com/problems/game-play-analysis-i/)|Game Play Analysis I|
|[512](https://leetcode-cn.com/problems/game-play-analysis-ii/)|Game Play Analysis II|
|[534](https://leetcode-cn.com/problems/game-play-analysis-iii/)|Game Play Analysis III|
|[550](https://leetcode-cn.com/problems/game-play-analysis-iv/)|Game Play Analysis IV|
|[1097](https://leetcode-cn.com/problems/game-play-analysis-v/)|Game Play Analysis V|
|||
|[51](https://leetcode-cn.com/problems/n-queens/)|N-Queens|
|[52](https://leetcode-cn.com/problems/n-queens-ii/)|N-Queens II|
|||
|[74](https://leetcode-cn.com/problems/search-a-2d-matrix/)|Search a 2D Matrix|
|[240](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)|Search a 2D Matrix II|
|||
|[485](https://leetcode-cn.com/problems/max-consecutive-ones/)|Max Consecutive Ones|
|[487](https://leetcode-cn.com/problems/max-consecutive-ones-ii/)|Max Consecutive Ones II|
|[1004](https://leetcode-cn.com/problems/max-consecutive-ones-iii/)|Max Consecutive Ones III|
|||
|[91](https://leetcode-cn.com/problems/decode-ways/)|Decode Ways|
|[639](https://leetcode-cn.com/problems/decode-ways-ii/)|Decode Ways II|
|||
|[1082](https://leetcode-cn.com/problems/sales-analysis-i/)|Sales Analysis I|
|[1083](https://leetcode-cn.com/problems/sales-analysis-ii/)|Sales Analysis II|
|[1084](https://leetcode-cn.com/problems/sales-analysis-iii/)|Sales Analysis III|
|||
|[1141](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-i/)|User Activity for the Past 30 Days I|
|[1142](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-ii/)|User Activity for the Past 30 Days II|
|||
|[877](https://leetcode-cn.com/problems/stone-game/)|Stone Game|
|[1140](https://leetcode-cn.com/problems/stone-game-ii/)|Stone Game II|
|[1406](https://leetcode-cn.com/problems/stone-game-iii/)|Stone Game III|
|[1510](https://leetcode-cn.com/problems/stone-game-iv/)|Stone Game IV|
|||
|[139](https://leetcode-cn.com/problems/word-break/)|Word Break|
|[140](https://leetcode-cn.com/problems/word-break-ii/)|Word Break II|
|||
|[79](https://leetcode-cn.com/problems/word-search/)|Word Search|
|[212](https://leetcode-cn.com/problems/word-search-ii/)|Word Search II|
|||
|[243](https://leetcode-cn.com/problems/shortest-word-distance/)|Shortest Word Distance|
|[244](https://leetcode-cn.com/problems/shortest-word-distance-ii/)|Shortest Word Distance II|
|[245](https://leetcode-cn.com/problems/shortest-word-distance-iii/)|Shortest Word Distance III|
|||
|[280](https://leetcode-cn.com/problems/wiggle-sort/)|Wiggle Sort|
|[324](https://leetcode-cn.com/problems/wiggle-sort-ii/)|Wiggle Sort II|
|||
|[1068](https://leetcode-cn.com/problems/product-sales-analysis-i/)|Product Sales Analysis I|
|[1069](https://leetcode-cn.com/problems/product-sales-analysis-ii/)|Product Sales Analysis II|
|[1070](https://leetcode-cn.com/problems/product-sales-analysis-iii/)|Product Sales Analysis III|
|||
|[931](https://leetcode-cn.com/problems/minimum-falling-path-sum/)|Minimum Falling Path Sum|
|[1289](https://leetcode-cn.com/problems/minimum-falling-path-sum-ii/)|Minimum Falling Path Sum II|
|||


---

[⬆️Top](#leetcli)
