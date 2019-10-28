
# leetcli

[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
[![Build Status](https://www.travis-ci.org/evercyan/leetcli.svg?branch=master)](https://www.travis-ci.org/evercyan/leetcli)
[![codecov](https://codecov.io/gh/evercyan/leetcli/branch/master/graph/badge.svg?token=RbJTUtAlvl)](https://codecov.io/gh/evercyan/leetcli)

---

#### Emmmm..

```
刷题烧脑的过程是快乐的, 但准备工作是蛋疼的, 通常是像这样:

准备刷某题 > 创建答题目录 > 创建答题文件 > 准备测试用例 > 刷刷刷...

So, all For DRY (Don't Repeat Yourself)...

本程序使用 golang 编程, 基于终端命令式交互
通过 leetcode api 获取问题相关数据
处理生成答题文件, 测试文件, 和 readme.md 等
支持 golang 覆盖率测试等
```

---

#### Help

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

#### Problem Tag

[![数组](https://img.shields.io/badge/数组-201-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/array/)
[![动态规划](https://img.shields.io/badge/动态规划-166-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/dynamic-programming/)
[![数学](https://img.shields.io/badge/数学-155-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/math/)
[![字符串](https://img.shields.io/badge/字符串-146-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/string/)
[![树](https://img.shields.io/badge/树-120-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/tree/)
[![哈希表](https://img.shields.io/badge/哈希表-119-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/hash-table/)
[![深度优先搜索](https://img.shields.io/badge/深度优先搜索-109-brightgreen.svg?style=flat)](https://leetcode-cn.com/tag/depth-first-search/)
[![二分查找](https://img.shields.io/badge/二分查找-79-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/binary-search/)
[![贪心算法](https://img.shields.io/badge/贪心算法-64-yellowgreen.svg?style=flat)](https://leetcode-cn.com/tag/greedy/)
[![双指针](https://img.shields.io/badge/双指针-59-orange.svg?style=flat)](https://leetcode-cn.com/tag/two-pointers/)
[![广度优先搜索](https://img.shields.io/badge/广度优先搜索-54-orange.svg?style=flat)](https://leetcode-cn.com/tag/breadth-first-search/)
[![栈](https://img.shields.io/badge/栈-53-orange.svg?style=flat)](https://leetcode-cn.com/tag/stack/)
[![回溯算法](https://img.shields.io/badge/回溯算法-49-orange.svg?style=flat)](https://leetcode-cn.com/tag/backtracking/)
[![设计](https://img.shields.io/badge/设计-41-orange.svg?style=flat)](https://leetcode-cn.com/tag/design/)
[![排序](https://img.shields.io/badge/排序-37-blue.svg?style=flat)](https://leetcode-cn.com/tag/sort/)
[![位运算](https://img.shields.io/badge/位运算-36-blue.svg?style=flat)](https://leetcode-cn.com/tag/bit-manipulation/)
[![图](https://img.shields.io/badge/图-36-blue.svg?style=flat)](https://leetcode-cn.com/tag/graph/)
[![链表](https://img.shields.io/badge/链表-36-blue.svg?style=flat)](https://leetcode-cn.com/tag/linked-list/)
[![堆](https://img.shields.io/badge/堆-34-blue.svg?style=flat)](https://leetcode-cn.com/tag/heap/)
[![并查集](https://img.shields.io/badge/并查集-28-blue.svg?style=flat)](https://leetcode-cn.com/tag/union-find/)
[![SlidingWindow](https://img.shields.io/badge/SlidingWindow-20-red.svg?style=flat)](https://leetcode-cn.com/tag/sliding-window/)
[![分治算法](https://img.shields.io/badge/分治算法-18-red.svg?style=flat)](https://leetcode-cn.com/tag/divide-and-conquer/)
[![字典树](https://img.shields.io/badge/字典树-17-red.svg?style=flat)](https://leetcode-cn.com/tag/trie/)
[![递归](https://img.shields.io/badge/递归-15-red.svg?style=flat)](https://leetcode-cn.com/tag/recursion/)
[![线段树](https://img.shields.io/badge/线段树-11-red.svg?style=flat)](https://leetcode-cn.com/tag/segment-tree/)
[![OrderedMap](https://img.shields.io/badge/OrderedMap-10-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/ordered-map/)
[![队列](https://img.shields.io/badge/队列-9-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/queue/)
[![极小化极大](https://img.shields.io/badge/极小化极大-8-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/minimax/)
[![树状数组](https://img.shields.io/badge/树状数组-6-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/binary-indexed-tree/)
[![Random](https://img.shields.io/badge/Random-6-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/random/)
[![拓扑排序](https://img.shields.io/badge/拓扑排序-6-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/topological-sort/)
[![脑筋急转弯](https://img.shields.io/badge/脑筋急转弯-5-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/brainteaser/)
[![几何](https://img.shields.io/badge/几何-4-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/geometry/)
[![LineSweep](https://img.shields.io/badge/LineSweep-4-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/line-sweep/)
[![二叉搜索树](https://img.shields.io/badge/二叉搜索树-2-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/binary-search-tree/)
[![RejectionSampling](https://img.shields.io/badge/RejectionSampling-2-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/rejection-sampling/)
[![蓄水池抽样](https://img.shields.io/badge/蓄水池抽样-2-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/reservoir-sampling/)
[![记忆化](https://img.shields.io/badge/记忆化-1-lightgray.svg?style=flat)](https://leetcode-cn.com/tag/memoization/)

---

#### Problem List

|#|标题|难度|通过率|总提交次数|
|:-:|:-|:-: | :-: | :-: |
|[0000](https://leetcode-cn.com/problems/guess-numbers/)|Guess Numbers|简单|82%|13123|
|[0000](https://leetcode-cn.com/problems/programmable-robot/)|Programmable Robot|中等|14%|11765|
|[0000](https://leetcode-cn.com/problems/coin-bonus/)|Coin Bonus|困难|21%|1973|
|[0000](https://leetcode-cn.com/problems/broken-board-dominoes/)|Broken Board Dominoes|困难|22%|1648|
|[0000](https://leetcode-cn.com/problems/deep-dark-fraction/)|Deep Dark Fraction|简单|51%|4324|
|[1227](https://leetcode-cn.com/problems/airplane-seat-assignment-probability/)|等价多米诺骨牌对的数量|中等|64%|206|
|[5241](https://leetcode-cn.com/problems/tiling-a-rectangle-with-the-fewest-squares/)|Tiling a Rectangle with the Fewest Squares|困难|45%|441|
|[5240](https://leetcode-cn.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/)|Maximum Length of a Concatenated String with Unique Characters|中等|30%|1949|
|[5239](https://leetcode-cn.com/problems/circular-permutation-in-binary-representation/)|Circular Permutation in Binary Representation|中等|52%|1222|
|[5238](https://leetcode-cn.com/problems/find-positive-integer-solution-for-a-given-equation/)|Find Positive Integer Solution for a Given Equation|简单|62%|1710|
|[1225](https://leetcode-cn.com/problems/report-contiguous-dates/)|Report Contiguous Dates|困难|61%|100|
|[1235](https://leetcode-cn.com/problems/maximum-profit-in-job-scheduling/)|Maximum Profit in Job Scheduling|困难|41%|1397|
|[1234](https://leetcode-cn.com/problems/replace-the-substring-for-balanced-string/)|Replace the Substring for Balanced String|中等|25%|2832|
|[1233](https://leetcode-cn.com/problems/remove-sub-folders-from-the-filesystem/)|Remove Sub-Folders from the Filesystem|中等|37%|3149|
|[1232](https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/)|Check If It Is a Straight Line|简单|45%|4527|
|[1224](https://leetcode-cn.com/problems/maximum-equal-frequency/)|Maximum Equal Frequency|困难|26%|2015|
|[1223](https://leetcode-cn.com/problems/dice-roll-simulation/)|Dice Roll Simulation|中等|37%|1943|
|[1222](https://leetcode-cn.com/problems/queens-that-can-attack-the-king/)|Queens That Can Attack the King|中等|60%|2405|
|[1221](https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/)|Split a String in Balanced Strings|简单|76%|4772|
|[1226](https://leetcode-cn.com/problems/the-dining-philosophers/)|The Dining Philosophers|中等|65%|321|
|[1212](https://leetcode-cn.com/problems/team-scores-in-football-tournament/)|Team Scores in Football Tournament|中等|76%|46|
|[1211](https://leetcode-cn.com/problems/queries-quality-and-percentage/)|Queries Quality and Percentage|简单|62%|106|
|[1206](https://leetcode-cn.com/problems/design-skiplist/)|航班预订统计|困难|48%|290|
|[1220](https://leetcode-cn.com/problems/count-vowels-permutation/)|最小的必要团队|困难|52%|1961|
|[1219](https://leetcode-cn.com/problems/path-with-maximum-gold/)|表现良好的最长时间段|中等|59%|2147|
|[1218](https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/)|最深叶节点的最近公共祖先|中等|31%|4784|
|[1217](https://leetcode-cn.com/problems/play-with-chips/)|数组的相对排序|简单|64%|3848|
|[1205](https://leetcode-cn.com/problems/monthly-transactions-ii/)|IP 地址无效化|中等|33%|198|
|[1204](https://leetcode-cn.com/problems/last-person-to-fit-in-the-elevator/)|每日新用户统计|中等|50%|55|
|[1210](https://leetcode-cn.com/problems/minimum-moves-to-reach-target-with-rotations/)|Minimum Moves to Reach Target with Rotations|困难|43%|1060|
|[1208](https://leetcode-cn.com/problems/get-equal-substrings-within-budget/)|有效括号的嵌套深度|中等|26%|5344|
|[1209](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii/)|设计有限阻塞队列|中等|50%|2802|
|[1207](https://leetcode-cn.com/problems/unique-number-of-occurrences/)|删点成林|简单|66%|4538|
|[1194](https://leetcode-cn.com/problems/tournament-winners/)|二叉树寻路|困难|38%|122|
|[1193](https://leetcode-cn.com/problems/monthly-transactions-i/)|游戏玩法分析 V|中等|57%|276|
|[1195](https://leetcode-cn.com/problems/fizz-buzz-multithreaded/)|分糖果 II|中等|57%|1037|
|[1203](https://leetcode-cn.com/problems/sort-items-by-groups-respecting-dependencies/)|按序打印|困难|36%|529|
|[1202](https://leetcode-cn.com/problems/smallest-string-with-swaps/)|Smallest String With Swaps|中等|23%|2447|
|[1201](https://leetcode-cn.com/problems/ugly-number-iii/)|Ugly Number III|中等|16%|4982|
|[1200](https://leetcode-cn.com/problems/minimum-absolute-difference/)|Minimum Absolute Difference|简单|62%|3990|
|[1179](https://leetcode-cn.com/problems/reformat-department-table/)|游戏玩法分析 I|简单|66%|840|
|[1192](https://leetcode-cn.com/problems/critical-connections-in-a-network/)|分享巧克力|困难|28%|1637|
|[1191](https://leetcode-cn.com/problems/k-concatenation-maximum-sum/)|K-Concatenation Maximum Sum|中等|22%|5713|
|[1190](https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/)|Reverse Substrings Between Each Pair of Parentheses|中等|46%|3864|
|[1189](https://leetcode-cn.com/problems/maximum-number-of-balloons/)|Maximum Number of Balloons|简单|60%|6402|
|[1174](https://leetcode-cn.com/problems/immediate-food-delivery-ii/)|销售分析III|中等|57%|61|
|[1173](https://leetcode-cn.com/problems/immediate-food-delivery-i/)|销售分析 II|简单|74%|110|
|[1187](https://leetcode-cn.com/problems/make-array-strictly-increasing/)|交替打印FooBar|困难|33%|1251|
|[1185](https://leetcode-cn.com/problems/day-of-the-week/)|山脉数组中查找目标值|简单|56%|4075|
|[1186](https://leetcode-cn.com/problems/maximum-subarray-sum-with-one-deletion/)|H2O 生成|中等|24%|4597|
|[1184](https://leetcode-cn.com/problems/distance-between-bus-stops/)|拼车|简单|54%|4962|
|[1178](https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle/)|验证回文字符串 III|困难|23%|2774|
|[1177](https://leetcode-cn.com/problems/can-make-palindrome-from-substring/)|Can Make Palindrome from Substring|中等|19%|5020|
|[1176](https://leetcode-cn.com/problems/diet-plan-performance/)|Diet Plan Performance|简单|38%|5595|
|[1175](https://leetcode-cn.com/problems/prime-arrangements/)|数组变换|简单|44%|4185|
|[1164](https://leetcode-cn.com/problems/product-price-at-a-given-date/)|等差数列中缺失的数字|中等|58%|65|
|[1170](https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character/)|最短公共超序列|简单|57%|4471|
|[1169](https://leetcode-cn.com/problems/invalid-transactions/)|受标签影响的最大值|中等|25%|6986|
|[1236](https://leetcode-cn.com/problems/web-crawler/)|第 N 个泰波那契数|中等|41%|43|
|[1172](https://leetcode-cn.com/problems/dinner-plate-stacks/)|销售分析 I |困难|21%|2377|
|[1159](https://leetcode-cn.com/problems/market-analysis-ii/)|不同字符的最小子序列|困难|47%|36|
|[1158](https://leetcode-cn.com/problems/market-analysis-i/)|Market Analysis I|中等|56%|65|
|[1171](https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/)|二进制矩阵中的最短路径|中等|29%|4717|
|[1155](https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/)|产品销售分析 III|中等|41%|4414|
|[1157](https://leetcode-cn.com/problems/online-majority-element-in-subarray/)|根到叶路径上的不足节点|困难|20%|3373|
|[1156](https://leetcode-cn.com/problems/swap-for-longest-repeated-character-substring/)|Bigram 分词|中等|39%|2814|
|[1154](https://leetcode-cn.com/problems/day-of-the-year/)|产品销售分析 II|简单|50%|6745|
|[1149](https://leetcode-cn.com/problems/article-views-ii/)|三个有序数组的交集|中等|53%|90|
|[1148](https://leetcode-cn.com/problems/article-views-i/)|负二进制数相加|简单|80%|126|
|[1147](https://leetcode-cn.com/problems/longest-chunked-palindrome-decomposition/)|按列翻转得到最大值等行数|困难|50%|1998|
|[1143](https://leetcode-cn.com/problems/longest-common-subsequence/)|找出所有行中最小公共元素|中等|57%|2772|
|[1146](https://leetcode-cn.com/problems/snapshot-array/)|字符串的最大公因子|中等|19%|5052|
|[1145](https://leetcode-cn.com/problems/binary-tree-coloring-game/)|元素和为目标值的子矩阵数量|中等|35%|3724|
|[1144](https://leetcode-cn.com/problems/decrease-elements-to-make-array-zigzag/)|水资源分配优化|中等|37%|7059|
|[1142](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-ii/)|进击的骑士|简单|30%|205|
|[1141](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-i/)|最多可以买到的苹果数量|简单|56%|136|
|[1140](https://leetcode-cn.com/problems/stone-game-ii/)|距离相等的条形码|中等|53%|1279|
|[1139](https://leetcode-cn.com/problems/largest-1-bordered-square/)|交换一次的先前排列|中等|37%|2674|
|[1138](https://leetcode-cn.com/problems/alphabet-board-path/)|爱生气的书店老板|中等|35%|4265|
|[1132](https://leetcode-cn.com/problems/reported-posts-ii/)|前后拼接|中等|26%|150|
|[1137](https://leetcode-cn.com/problems/n-th-tribonacci-number/)|高度检查器|简单|53%|8605|
|[1131](https://leetcode-cn.com/problems/maximum-of-absolute-value-expression/)|统计只含单一字母的子串|中等|34%|1873|
|[1129](https://leetcode-cn.com/problems/shortest-path-with-alternating-colors/)|最长字符串链|中等|28%|2347|
|[1130](https://leetcode-cn.com/problems/minimum-cost-tree-from-leaf-values/)|最后一块石头的重量 II|中等|53%|1222|
|[1128](https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/)|删除字符串中的所有相邻重复项|简单|37%|6529|
|[1127](https://leetcode-cn.com/problems/user-purchase-platform/)|最后一块石头的重量|困难|45%|40|
|[1126](https://leetcode-cn.com/problems/active-businesses/)|连接棒材的最低费用|中等|73%|72|
|[1125](https://leetcode-cn.com/problems/smallest-sufficient-team/)|设计文件系统|困难|37%|1758|
|[1124](https://leetcode-cn.com/problems/longest-well-performing-interval/)|字符串转化|中等|19%|9463|
|[1123](https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves/)|单行键盘|中等|59%|1665|
|[1122](https://leetcode-cn.com/problems/relative-sort-array/)|最长重复子串|简单|63%|8953|
|[1116](https://leetcode-cn.com/problems/print-zero-even-odd/)|最大层内元素和|中等|45%|5993|
|[1113](https://leetcode-cn.com/problems/reported-posts/)|移动石子直到连续 II|简单|53%|137|
|[1112](https://leetcode-cn.com/problems/highest-grade-for-each-student/)|拼写单词|中等|62%|120|
|[1188](https://leetcode-cn.com/problems/design-bounded-blocking-queue/)|花括号展开 II|中等|70%|55|
|[1111](https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/)|多边形三角剖分的最低得分|中等|66%|1178|
|[1110](https://leetcode-cn.com/problems/delete-nodes-and-return-forest/)|公路流量|中等|53%|2095|
|[1109](https://leetcode-cn.com/problems/corporate-flight-bookings/)|Corporate Flight Bookings|中等|32%|6291|
|[1108](https://leetcode-cn.com/problems/defanging-an-ip-address/)|用户网站访问行为分析|简单|80%|19724|
|[1107](https://leetcode-cn.com/problems/new-users-daily-count/)|最少交换次数来组合所有的 1|中等|44%|177|
|[1114](https://leetcode-cn.com/problems/print-in-order/)|从二叉搜索树到更大和树|简单|56%|16381|
|[1098](https://leetcode-cn.com/problems/unpopular-books/)|最大唯一数|中等|36%|1171|
|[1106](https://leetcode-cn.com/problems/parsing-a-boolean-expression/)|逃离大迷宫|困难|54%|2111|
|[1105](https://leetcode-cn.com/problems/filling-bookcase-shelves/)|不相交的线|中等|45%|3279|
|[1103](https://leetcode-cn.com/problems/distribute-candies-to-people/)|移动石子直到连续|简单|60%|5853|
|[1104](https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/)|边框着色|中等|65%|2647|
|[1097](https://leetcode-cn.com/problems/game-play-analysis-v/)|字符流|困难|39%|457|
|[1231](https://leetcode-cn.com/problems/divide-chocolate/)|Divide Chocolate|困难|51%|601|
|[1096](https://leetcode-cn.com/problems/brace-expansion-ii/)|两个非重叠子数组的最大和|困难|41%|628|
|[1115](https://leetcode-cn.com/problems/print-foobar-alternately/)|有效的回旋镖|中等|55%|9456|
|[1117](https://leetcode-cn.com/problems/building-h2o/)|地图分析|中等|39%|5998|
|[1095](https://leetcode-cn.com/problems/find-in-mountain-array/)|两地调度|困难|27%|2368|
|[1094](https://leetcode-cn.com/problems/car-pooling/)|距离顺序排列矩阵单元格|中等|48%|3826|
|[1093](https://leetcode-cn.com/problems/statistics-from-a-large-sample/)|从先序遍历还原二叉树|中等|36%|3070|
|[0550](https://leetcode-cn.com/problems/game-play-analysis-iv/)|Game Play Analysis IV|中等|43%|1345|
|[0534](https://leetcode-cn.com/problems/game-play-analysis-iii/)|Game Play Analysis III|中等|62%|1208|
|[0512](https://leetcode-cn.com/problems/game-play-analysis-ii/)|Game Play Analysis II|简单|49%|2131|
|[0511](https://leetcode-cn.com/problems/game-play-analysis-i/)|从始点到终点的所有路径|简单|70%|1505|
|[1216](https://leetcode-cn.com/problems/valid-palindrome-iii/)|打印零与奇偶数|困难|43%|886|
|[1084](https://leetcode-cn.com/problems/sales-analysis-iii/)|长度为 K 的无重复字符子串|简单|44%|1470|
|[1083](https://leetcode-cn.com/problems/sales-analysis-ii/)|小于 K 的两数之和|简单|46%|1344|
|[1082](https://leetcode-cn.com/problems/sales-analysis-i/)|最小元素各数位之和|简单|71%|1210|
|[1091](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)|子树的最大平均值|中等|24%|4774|
|[1092](https://leetcode-cn.com/problems/shortest-common-supersequence/)|节点与其祖先之间的最大差值|困难|35%|1384|
|[1090](https://leetcode-cn.com/problems/largest-values-from-labels/)|阿姆斯特朗数|中等|45%|2125|
|[1089](https://leetcode-cn.com/problems/duplicate-zeros/)|删去字符串中的元音|简单|57%|7865|
|[1199](https://leetcode-cn.com/problems/minimum-time-to-build-blocks/)|Minimum Time to Build Blocks|困难|29%|773|
|[1230](https://leetcode-cn.com/problems/toss-strange-coins/)|绝对值表达式的最大值|中等|37%|1288|
|[1229](https://leetcode-cn.com/problems/meeting-scheduler/)|颜色交替的最短路径|中等|29%|2119|
|[1228](https://leetcode-cn.com/problems/missing-number-in-arithmetic-progression/)|叶值的最小代价生成树|简单|52%|1721|
|[1077](https://leetcode-cn.com/problems/project-employees-iii/)|易混淆数 II|中等|66%|759|
|[1076](https://leetcode-cn.com/problems/project-employees-ii/)|字母切换|简单|50%|1223|
|[1075](https://leetcode-cn.com/problems/project-employees-i/)|字符串的索引对|简单|38%|1153|
|[1079](https://leetcode-cn.com/problems/letter-tile-possibilities/)|从根到叶的二进制数之和|中等|71%|2424|
|[1081](https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters/)|视频拼接|中等|42%|1960|
|[1080](https://leetcode-cn.com/problems/insufficient-nodes-in-root-to-leaf-paths/)|驼峰式匹配|中等|39%|2296|
|[1078](https://leetcode-cn.com/problems/occurrences-after-bigram/)|删除最外层的括号|简单|58%|5470|
|[1070](https://leetcode-cn.com/problems/product-sales-analysis-iii/)|负二进制转换|中等|43%|1010|
|[1069](https://leetcode-cn.com/problems/product-sales-analysis-ii/)|易混淆数|简单|78%|1222|
|[1068](https://leetcode-cn.com/problems/product-sales-analysis-i/)|范围内的数字计数|简单|77%|1080|
|[1183](https://leetcode-cn.com/problems/maximum-number-of-ones/)|大样本统计|困难|36%|363|
|[1215](https://leetcode-cn.com/problems/stepping-numbers/)|Stepping Numbers|中等|31%|1656|
|[1214](https://leetcode-cn.com/problems/two-sum-bsts/)|Two Sum BSTs|中等|59%|1193|
|[1213](https://leetcode-cn.com/problems/intersection-of-three-sorted-arrays/)|Intersection of Three Sorted Arrays|简单|73%|1616|
|[1073](https://leetcode-cn.com/problems/adding-two-negabinary-numbers/)|飞地的数量|中等|28%|2503|
|[1072](https://leetcode-cn.com/problems/flip-columns-for-maximum-number-of-equal-rows/)|链表中的下一个更大节点|中等|41%|1724|
|[1071](https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/)|可被 5 整除的二进制前缀|简单|47%|5494|
|[1074](https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/)|前五科的均分|困难|34%|1219|
|[1168](https://leetcode-cn.com/problems/optimize-water-distribution-in-a-village/)|复写零|困难|35%|627|
|[1198](https://leetcode-cn.com/problems/find-smallest-common-element-in-all-rows/)|小众书籍|中等|68%|1130|
|[1197](https://leetcode-cn.com/problems/minimum-knight-moves/)|解析布尔表达式|中等|24%|2295|
|[1196](https://leetcode-cn.com/problems/how-many-apples-can-you-put-into-the-basket/)|填充书架|简单|67%|1436|
|[1054](https://leetcode-cn.com/problems/distant-barcodes/)|十进制整数的反码|中等|29%|5341|
|[1053](https://leetcode-cn.com/problems/previous-permutation-with-one-swap/)|最小化舍入误差以满足目标|中等|40%|3920|
|[1052](https://leetcode-cn.com/problems/grumpy-bookstore-owner/)|校园自行车分配|中等|43%|4809|
|[1051](https://leetcode-cn.com/problems/height-checker/)|形成字符串的最短路径|简单|71%|10652|
|[1050](https://leetcode-cn.com/problems/actors-and-directors-who-cooperated-at-least-three-times/)|先序遍历构造二叉树|简单|72%|1345|
|[1045](https://leetcode-cn.com/problems/customers-who-bought-all-products/)|检查替换后的词是否有效|中等|59%|1016|
|[1182](https://leetcode-cn.com/problems/shortest-distance-to-target-color/)|游戏玩法分析 IV|中等|24%|1410|
|[1163](https://leetcode-cn.com/problems/last-substring-in-lexicographical-order/)|项目员工 III|困难|21%|4811|
|[1181](https://leetcode-cn.com/problems/before-and-after-puzzle/)|游戏玩法分析 III|中等|33%|1066|
|[1180](https://leetcode-cn.com/problems/count-substrings-with-only-one-distinct-letter/)|游戏玩法分析 II|简单|71%|949|
|[1049](https://leetcode-cn.com/problems/last-stone-weight-ii/)|行相等的最少多米诺旋转|中等|36%|3162|
|[1048](https://leetcode-cn.com/problems/longest-string-chain/)|笨阶乘|中等|35%|4060|
|[1047](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/)|K 次取反后最大化的数组和|简单|62%|10591|
|[1046](https://leetcode-cn.com/problems/last-stone-weight/)|最大连续1的个数 III|简单|57%|10538|
|[1167](https://leetcode-cn.com/problems/minimum-cost-to-connect-sticks/)|建造街区的最短时间|中等|27%|1527|
|[1166](https://leetcode-cn.com/problems/design-file-system/)|抛掷硬币|中等|46%|804|
|[1153](https://leetcode-cn.com/problems/string-transforms-into-another-string/)|产品销售分析 I|困难|27%|1268|
|[1165](https://leetcode-cn.com/problems/single-row-keyboard/)|安排会议日程|简单|81%|1560|
|[1044](https://leetcode-cn.com/problems/longest-duplicate-substring/)|查找常用字符|困难|14%|3317|
|[1043](https://leetcode-cn.com/problems/partition-array-for-maximum-sum/)|网格照明|中等|59%|1626|
|[1042](https://leetcode-cn.com/problems/flower-planting-with-no-adjacent/)|合并石头的最低成本|简单|44%|4683|
|[1041](https://leetcode-cn.com/problems/robot-bounded-in-circle/)|车的可用捕获量|中等|40%|5457|
|[1121](https://leetcode-cn.com/problems/divide-array-into-increasing-sequences/)|分隔数组以得到最大和|困难|44%|782|
|[1162](https://leetcode-cn.com/problems/as-far-from-land-as-possible/)|项目员工II|中等|33%|8934|
|[1161](https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree/)|项目员工 I|中等|63%|2993|
|[1037](https://leetcode-cn.com/problems/valid-boomerang/)|K 连续位的最小翻转次数|简单|38%|6399|
|[1038](https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/)|正方形数组的数目|中等|69%|4324|
|[1040](https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/)|最大二叉树 II|中等|43%|1188|
|[1160](https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/)|活字印刷|简单|62%|8323|
|[1039](https://leetcode-cn.com/problems/minimum-score-triangulation-of-polygon/)|找到小镇的法官|中等|39%|1865|
|[1152](https://leetcode-cn.com/problems/analyze-user-website-visit-pattern/)|矩阵中 1 的最大数量|中等|27%|910|
|[1151](https://leetcode-cn.com/problems/minimum-swaps-to-group-all-1s-together/)|步进数|中等|35%|1481|
|[1036](https://leetcode-cn.com/problems/escape-a-large-maze/)|腐烂的橘子|困难|23%|1953|
|[1035](https://leetcode-cn.com/problems/uncrossed-lines/)|二叉树的堂兄弟节点|中等|47%|2043|
|[1034](https://leetcode-cn.com/problems/coloring-a-border/)|K 个不同整数的子数组|中等|37%|2041|
|[1033](https://leetcode-cn.com/problems/moving-stones-until-consecutive/)|坏了的计算器|简单|35%|7058|
|[1150](https://leetcode-cn.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array/)|查找两棵二叉搜索树之和|简单|63%|1257|
|[1136](https://leetcode-cn.com/problems/parallel-courses/)|合作过至少三次的演员和导演|困难|52%|568|
|[1135](https://leetcode-cn.com/problems/connecting-cities-with-minimum-cost/)|买下所有产品的客户|中等|31%|1416|
|[1102](https://leetcode-cn.com/problems/path-with-maximum-minimum-value/)|检查一个数是否在数组中占绝大多数|中等|22%|1308|
|[1133](https://leetcode-cn.com/problems/largest-unique-number/)|按字典序排在最后的子串|简单|58%|1196|
|[1032](https://leetcode-cn.com/problems/stream-of-characters/)|等式方程的可满足性|困难|24%|2133|
|[1031](https://leetcode-cn.com/problems/maximum-sum-of-two-non-overlapping-subarrays/)|数组形式的整数加法|中等|49%|2014|
|[1029](https://leetcode-cn.com/problems/two-city-scheduling/)|二叉树的垂序遍历|简单|55%|7046|
|[1030](https://leetcode-cn.com/problems/matrix-cells-in-distance-order/)|从叶结点开始的最小字符串|简单|58%|4830|
|[1028](https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/)|区间列表的交集|困难|58%|1623|
|[1026](https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor/)|不含 AAA 或 BBB 的字符串|中等|53%|3105|
|[1120](https://leetcode-cn.com/problems/maximum-average-subtree/)|不邻接植花|中等|55%|890|
|[1134](https://leetcode-cn.com/problems/armstrong-number/)|与目标颜色间的最短距离|简单|73%|1036|
|[1119](https://leetcode-cn.com/problems/remove-vowels-from-a-string/)|困于环中的机器人|简单|82%|1845|
|[1118](https://leetcode-cn.com/problems/number-of-days-in-a-month/)|将数组分成几个递增序列|简单|61%|1055|
|[1027](https://leetcode-cn.com/problems/longest-arithmetic-sequence/)|查询后的偶数和|中等|36%|6574|
|[1025](https://leetcode-cn.com/problems/divisor-game/)|最低票价|简单|68%|14445|
|[1101](https://leetcode-cn.com/problems/the-earliest-moment-when-everyone-become-friends/)|平行课程|中等|59%|697|
|[1100](https://leetcode-cn.com/problems/find-k-length-substrings-with-no-repeated-characters/)|最低成本联通所有城市|中等|59%|825|
|[1099](https://leetcode-cn.com/problems/two-sum-less-than-k/)|得分最高的路径|简单|50%|1451|
|[1085](https://leetcode-cn.com/problems/sum-of-digits-in-the-minimum-number/)|彼此熟识的最早时间|简单|72%|1116|
|[1024](https://leetcode-cn.com/problems/video-stitching/)|按位与为零的三元组|中等|45%|3466|
|[1023](https://leetcode-cn.com/problems/camelcase-matching/)|基于时间的键值存储|中等|50%|2521|
|[1022](https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/)|[不同路径 III](./src/1022-sum-of-root-to-leaf-binary-numbers/)|简单|50%|5365|
|[1021](https://leetcode-cn.com/problems/remove-outermost-parentheses/)|[在二叉树中分配硬币](./src/1021-remove-outermost-parentheses/)|简单|75%|17866|
|[1088](https://leetcode-cn.com/problems/confusing-number-ii/)|一月有多少天|困难|26%|613|
|[1087](https://leetcode-cn.com/problems/brace-expansion/)|最长等差数列|中等|39%|737|
|[1065](https://leetcode-cn.com/problems/index-pairs-of-a-string/)|子串能表示从 1 到 N 数字的二进制串|简单|44%|891|
|[1086](https://leetcode-cn.com/problems/high-five/)|除数博弈|简单|64%|1440|
|[1020](https://leetcode-cn.com/problems/number-of-enclaves/)|[最长湍流子数组](./src/1020-number-of-enclaves/)|中等|46%|3210|
|[1019](https://leetcode-cn.com/problems/next-greater-node-in-linked-list/)|[有序数组的平方](./src/1019-next-greater-node-in-linked-list/)|中等|46%|6441|
|[1018](https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/)|[三角形的最大周长](./src/1018-binary-prefix-divisible-by-5/)|简单|42%|10284|
|[1017](https://leetcode-cn.com/problems/convert-to-base-2/)|[奇偶跳](./src/1017-convert-to-base-2/)|中等|50%|2309|
|[1056](https://leetcode-cn.com/problems/confusing-number/)|在 D 天内送达包裹的能力|简单|31%|4112|
|[1067](https://leetcode-cn.com/problems/digit-count-in-range/)|校园自行车分配 II|困难|28%|512|
|[1066](https://leetcode-cn.com/problems/campus-bikes-ii/)|不动点|中等|26%|795|
|[1064](https://leetcode-cn.com/problems/fixed-point/)|可被 K 整除的最小整数|简单|74%|1674|
|[1016](https://leetcode-cn.com/problems/binary-string-with-substrings-representing-1-to-n/)|和可被 K 整除的子数组|中等|55%|2340|
|[1015](https://leetcode-cn.com/problems/smallest-integer-divisible-by-k/)|Smallest Integer Divisible by K|中等|25%|5151|
|[1014](https://leetcode-cn.com/problems/best-sightseeing-pair/)|[最接近原点的 K 个点](./src/1014-best-sightseeing-pair/)|中等|44%|4794|
|[1013](https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/)|[斐波那契数](./src/1013-partition-array-into-three-parts-with-equal-sum/)|简单|51%|8372|
|[1063](https://leetcode-cn.com/problems/number-of-valid-subarrays/)|最佳观光组合|困难|56%|729|
|[1062](https://leetcode-cn.com/problems/longest-repeating-substring/)|将数组分成和相等的三个部分|中等|36%|1345|
|[1060](https://leetcode-cn.com/problems/missing-element-in-sorted-array/)|最长重复子串|中等|44%|2089|
|[1061](https://leetcode-cn.com/problems/lexicographically-smallest-equivalent-string/)|有效子数组的数目|中等|41%|975|
|[1012](https://leetcode-cn.com/problems/numbers-with-repeated-digits/)|相等的有理数|困难|23%|2945|
|[1011](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/)|[翻转二叉树以匹配先序遍历](./src/1011-capacity-to-ship-packages-within-d-days/)|中等|46%|4082|
|[1010](https://leetcode-cn.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/)|强整数|简单|41%|11242|
|[1009](https://leetcode-cn.com/problems/complement-of-base-10-integer/)|[煎饼排序](./src/1009-complement-of-base-10-integer/)|简单|57%|7257|
|[1058](https://leetcode-cn.com/problems/minimize-rounding-error-to-meet-target/)|按字典序排列最小的等效字符串|中等|23%|1812|
|[1057](https://leetcode-cn.com/problems/campus-bikes/)|至少有 1 位重复的数字|中等|25%|1671|
|[1055](https://leetcode-cn.com/problems/shortest-way-to-form-string/)|总持续时间可被 60 整除的歌曲|中等|41%|546|
|[1008](https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal/)|监控二叉树|中等|71%|4288|
|[1007](https://leetcode-cn.com/problems/minimum-domino-rotations-for-equal-row/)|连续差相同的数字|中等|41%|3237|
|[1006](https://leetcode-cn.com/problems/clumsy-factorial/)|元音拼写检查器|中等|49%|3870|
|[1005](https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/)|单值二叉树|简单|49%|8450|
|[1004](https://leetcode-cn.com/problems/max-consecutive-ones-iii/)|表示数字的最少运算符|中等|44%|5030|
|[1003](https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions/)|最小面积矩形 II|中等|53%|4425|
|[1002](https://leetcode-cn.com/problems/find-common-characters/)|[最大宽度坡](./src/1002-find-common-characters/)|简单|64%|10052|
|[1001](https://leetcode-cn.com/problems/grid-illumination/)|重复 N 次的元素|困难|27%|1980|
|[1000](https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/)|删列造序 III|困难|26%|2141|
|[0999](https://leetcode-cn.com/problems/available-captures-for-rook/)|由斜杠划分区域|简单|64%|5796|
|[0998](https://leetcode-cn.com/problems/maximum-binary-tree-ii/)|二叉树的完全性检验|中等|55%|1889|
|[0997](https://leetcode-cn.com/problems/find-the-town-judge/)|Find the Town Judge|简单|46%|15085|
|[0996](https://leetcode-cn.com/problems/number-of-squareful-arrays/)|Number of Squareful Arrays|困难|40%|1703|
|[0995](https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/)|Minimum Number of K Consecutive Bit Flips|困难|36%|1811|
|[0994](https://leetcode-cn.com/problems/rotting-oranges/)|N 天后的牢房|简单|46%|8377|
|[0993](https://leetcode-cn.com/problems/cousins-in-binary-tree/)|最高的广告牌|简单|48%|7359|
|[0992](https://leetcode-cn.com/problems/subarrays-with-k-different-integers/)|删列造序 II|困难|22%|6831|
|[0991](https://leetcode-cn.com/problems/broken-calculator/)|二倍数对数组|中等|50%|4627|
|[0990](https://leetcode-cn.com/problems/satisfiability-of-equality-equations/)|验证外星语词典|中等|37%|3061|
|[0989](https://leetcode-cn.com/problems/add-to-array-form-of-integer/)|按公因数计算最大组件大小|简单|42%|14784|
|[0988](https://leetcode-cn.com/problems/smallest-string-starting-from-leaf/)|翻转等价二叉树|中等|41%|3482|
|[0987](https://leetcode-cn.com/problems/vertical-order-traversal-of-a-binary-tree/)|按递增顺序显示卡牌|中等|35%|2831|
|[0986](https://leetcode-cn.com/problems/interval-list-intersections/)|给定数字能组成的最大时间|中等|60%|3439|
|[0985](https://leetcode-cn.com/problems/sum-of-even-numbers-after-queries/)|令牌放置|简单|57%|10056|
|[0984](https://leetcode-cn.com/problems/string-without-aaa-or-bbb/)|移除最多的同行或同列石头|中等|33%|7039|
|[0983](https://leetcode-cn.com/problems/minimum-cost-for-tickets/)|[验证栈序列](./src/0983-minimum-cost-for-tickets/)|中等|53%|4898|
|[0982](https://leetcode-cn.com/problems/triples-with-bitwise-and-equal-to-zero/)|[使数组唯一的最小增量](./src/0982-triples-with-bitwise-and-equal-to-zero/)|困难|45%|1347|
|[0981](https://leetcode-cn.com/problems/time-based-key-value-store/)|[删列造序](./src/0981-time-based-key-value-store/)|中等|30%|2343|
|[0980](https://leetcode-cn.com/problems/unique-paths-iii/)|最短超级串|困难|67%|2643|
|[0979](https://leetcode-cn.com/problems/distribute-coins-in-binary-tree/)|增减字符串匹配|中等|63%|2391|
|[0978](https://leetcode-cn.com/problems/longest-turbulent-subarray/)|[有效的山脉数组](./src/0978-longest-turbulent-subarray/)|中等|36%|5399|
|[0977](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)|不同的子序列 II|简单|70%|32433|
|[0976](https://leetcode-cn.com/problems/largest-perimeter-triangle/)|[最小面积矩形](./src/0976-largest-perimeter-triangle/)|简单|54%|14713|
|[0975](https://leetcode-cn.com/problems/odd-even-jump/)|二叉搜索树的范围和|困难|40%|1536|
|[0974](https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/)|重新排列日志文件|中等|34%|4272|
|[0973](https://leetcode-cn.com/problems/k-closest-points-to-origin/)|戳印序列|中等|55%|10931|
|[0509](https://leetcode-cn.com/problems/fibonacci-number/)|二叉搜索树中的中序后继 II|简单|65%|36784|
|[0972](https://leetcode-cn.com/problems/equal-rational-numbers/)|骑士拨号器|困难|40%|1054|
|[0971](https://leetcode-cn.com/problems/flip-binary-tree-to-match-preorder-traversal/)|最短的桥|中等|40%|2072|
|[0970](https://leetcode-cn.com/problems/powerful-integers/)|Powerful Integers|简单|37%|9358|
|[0969](https://leetcode-cn.com/problems/pancake-sorting/)|最近的请求次数|中等|60%|3776|
|[0968](https://leetcode-cn.com/problems/binary-tree-cameras/)|漂亮数组|困难|34%|2651|
|[0967](https://leetcode-cn.com/problems/numbers-with-same-consecutive-differences/)|下降路径最小和|中等|32%|4477|
|[0966](https://leetcode-cn.com/problems/vowel-spellchecker/)|和相同的二元子数组|中等|34%|1972|
|[0965](https://leetcode-cn.com/problems/univalued-binary-tree/)|独特的电子邮件地址|简单|65%|12113|
|[0964](https://leetcode-cn.com/problems/least-operators-to-express-number/)|尽量减少恶意软件的传播 II|困难|28%|853|
|[0963](https://leetcode-cn.com/problems/minimum-area-rectangle-ii/)|三等分|中等|48%|1513|
|[0962](https://leetcode-cn.com/problems/maximum-width-ramp/)|将字符串翻转到单调递增|中等|32%|6134|
|[0961](https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array/)|长按键入|简单|65%|21973|
|[0960](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-iii/)|尽量减少恶意软件的传播|困难|45%|920|
|[0959](https://leetcode-cn.com/problems/regions-cut-by-slashes/)|三数之和的多种可能|中等|62%|1473|
|[0958](https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/)|按奇偶排序数组 II|中等|44%|4472|
|[0957](https://leetcode-cn.com/problems/prison-cells-after-n-days/)|使括号有效的最少添加|中等|32%|4684|
|[0956](https://leetcode-cn.com/problems/tallest-billboard/)|播放列表的数量|困难|34%|1666|
|[0955](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-ii/)|完全二叉树插入器|中等|26%|2798|
|[0954](https://leetcode-cn.com/problems/array-of-doubled-pairs/)|环形子数组的最大和|中等|25%|5448|
|[0953](https://leetcode-cn.com/problems/verifying-an-alien-dictionary/)|仅仅反转字母|简单|55%|5646|
|[0952](https://leetcode-cn.com/problems/largest-component-size-by-common-factor/)|单词子集|困难|25%|1993|
|[0951](https://leetcode-cn.com/problems/flip-equivalent-binary-trees/)|分割数组|中等|59%|2813|
|[0950](https://leetcode-cn.com/problems/reveal-cards-in-increasing-order/)|卡牌分组|中等|76%|5213|
|[0949](https://leetcode-cn.com/problems/largest-time-for-given-digits/)|猫和老鼠|简单|32%|7187|
|[0948](https://leetcode-cn.com/problems/bag-of-tokens/)|排序数组|中等|34%|4568|
|[0947](https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column/)|在线选举|中等|44%|1943|
|[0946](https://leetcode-cn.com/problems/validate-stack-sequences/)|最小差值 II|中等|51%|7711|
|[0945](https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/)|蛇梯棋|中等|39%|5900|
|[0944](https://leetcode-cn.com/problems/delete-columns-to-make-sorted/)|最小差值 I|简单|67%|9809|
|[0943](https://leetcode-cn.com/problems/find-the-shortest-superstring/)|子数组的最小值之和|困难|34%|918|
|[0942](https://leetcode-cn.com/problems/di-string-match/)|超级回文数|简单|70%|13173|
|[0941](https://leetcode-cn.com/problems/valid-mountain-array/)|按奇偶排序数组|简单|34%|15072|
|[0940](https://leetcode-cn.com/problems/distinct-subsequences-ii/)|水果成篮|困难|39%|1641|
|[0939](https://leetcode-cn.com/problems/minimum-area-rectangle/)|DI 序列的有效排列|中等|40%|2522|
|[0938](https://leetcode-cn.com/problems/range-sum-of-bst/)|最大为 N 的数字组合|简单|75%|15824|
|[0937](https://leetcode-cn.com/problems/reorder-data-in-log-files/)|股票价格跨度|简单|49%|5019|
|[0936](https://leetcode-cn.com/problems/stamping-the-sequence/)|RLE 迭代器|困难|27%|1184|
|[0935](https://leetcode-cn.com/problems/knight-dialer/)|有序队列|中等|36%|3041|
|[0934](https://leetcode-cn.com/problems/shortest-bridge/)|子数组按位或操作|中等|43%|3250|
|[0933](https://leetcode-cn.com/problems/number-of-recent-calls/)|递增顺序查找树|简单|68%|8387|
|[0932](https://leetcode-cn.com/problems/beautiful-array/)|单调数列|中等|51%|2661|
|[0931](https://leetcode-cn.com/problems/minimum-falling-path-sum/)|最大频率栈|中等|57%|5146|
|[0930](https://leetcode-cn.com/problems/binary-subarrays-with-sum/)|所有可能的满二叉树|中等|32%|4276|
|[0929](https://leetcode-cn.com/problems/unique-email-addresses/)|特殊等价字符串组|简单|64%|18941|
|[0928](https://leetcode-cn.com/problems/minimize-malware-spread-ii/)|三维形体的表面积|困难|38%|1125|
|[0927](https://leetcode-cn.com/problems/three-equal-parts/)|[子序列宽度之和](./src/0927-three-equal-parts/)|困难|27%|2336|
|[0926](https://leetcode-cn.com/problems/flip-string-to-monotone-increasing/)|查找和替换模式|中等|41%|3613|
|[0925](https://leetcode-cn.com/problems/long-pressed-name/)|根据前序和后序遍历构造二叉树|简单|43%|10858|
|[0924](https://leetcode-cn.com/problems/minimize-malware-spread/)|公平的糖果交换|困难|37%|1688|
|[0923](https://leetcode-cn.com/problems/3sum-with-multiplicity/)|鸡蛋掉落|中等|29%|4696|
|[0922](https://leetcode-cn.com/problems/sort-array-by-parity-ii/)|可能的二分法|简单|65%|27040|
|[0921](https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/)|螺旋矩阵 III|中等|67%|6328|
|[0920](https://leetcode-cn.com/problems/number-of-music-playlists/)|两句话中的不常见单词|困难|39%|989|
|[0919](https://leetcode-cn.com/problems/complete-binary-tree-inserter/)|三维形体投影面积|中等|51%|1307|
|[0918](https://leetcode-cn.com/problems/maximum-sum-circular-subarray/)|细分图中的可到达结点|中等|30%|4418|
|[0917](https://leetcode-cn.com/problems/reverse-only-letters/)|救生艇|简单|50%|11275|
|[0916](https://leetcode-cn.com/problems/word-subsets/)|索引处的解码字符串|中等|36%|3705|
|[0915](https://leetcode-cn.com/problems/partition-array-into-disjoint-intervals/)|在圆内随机生成点|中等|40%|5188|
|[0914](https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/)|非重叠矩形中的随机点|简单|31%|19920|
|[0913](https://leetcode-cn.com/problems/cat-and-mouse/)|随机翻转矩阵|困难|24%|997|
|[0912](https://leetcode-cn.com/problems/sort-an-array/)|按权重随机选择|中等|51%|16302|
|[0911](https://leetcode-cn.com/problems/online-election/)|盈利计划|中等|34%|1935|
|[0910](https://leetcode-cn.com/problems/smallest-range-ii/)|第 N 个神奇数字|中等|23%|5924|
|[0909](https://leetcode-cn.com/problems/snakes-and-ladders/)|石子游戏|中等|26%|2321|
|[0908](https://leetcode-cn.com/problems/smallest-range-i/)|链表的中间结点|简单|66%|10375|
|[0907](https://leetcode-cn.com/problems/sum-of-subarray-minimums/)|爱吃香蕉的珂珂|中等|25%|7933|
|[0906](https://leetcode-cn.com/problems/super-palindromes/)|模拟行走机器人|困难|18%|2942|
|[0905](https://leetcode-cn.com/problems/sort-array-by-parity/)|最长的斐波那契子序列的长度|简单|67%|35742|
|[0904](https://leetcode-cn.com/problems/fruit-into-baskets/)|叶子相似的树|中等|33%|4929|
|[0903](https://leetcode-cn.com/problems/valid-permutations-for-di-sequence/)|用 Rand7() 实现 Rand10()|困难|48%|1735|
|[0902](https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set/)|最低加油次数|困难|25%|3238|
|[0901](https://leetcode-cn.com/problems/online-stock-span/)|优势洗牌|中等|32%|4456|
|[0900](https://leetcode-cn.com/problems/rle-iterator/)|重新排序得到 2 的幂|中等|38%|2286|
|[0899](https://leetcode-cn.com/problems/orderly-queue/)|二进制间距|困难|47%|2445|
|[0898](https://leetcode-cn.com/problems/bitwise-ors-of-subarrays/)|转置矩阵|中等|24%|5784|
|[0897](https://leetcode-cn.com/problems/increasing-order-search-tree/)|回文素数|简单|62%|8090|
|[0896](https://leetcode-cn.com/problems/monotonic-array/)|具有所有最深结点的最小子树|简单|49%|18613|
|[0895](https://leetcode-cn.com/problems/maximum-frequency-stack/)|获取所有钥匙的最短路径|困难|42%|1909|
|[0894](https://leetcode-cn.com/problems/all-possible-full-binary-trees/)|黑名单中的随机数|中等|71%|3838|
|[0893](https://leetcode-cn.com/problems/groups-of-special-equivalent-strings/)|二叉树中所有距离为 K 的结点|简单|62%|4979|
|[0892](https://leetcode-cn.com/problems/surface-area-of-3d-shapes/)|和至少为 K 的最短子数组|简单|54%|6870|
|[0891](https://leetcode-cn.com/problems/sum-of-subsequence-widths/)|翻转矩阵后的得分|困难|23%|3071|
|[0890](https://leetcode-cn.com/problems/find-and-replace-pattern/)|柠檬水找零|中等|67%|3924|
|[0889](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/)|亲密字符串|中等|59%|3556|
|[0888](https://leetcode-cn.com/problems/fair-candy-swap/)|镜面反射|简单|49%|14111|
|[0887](https://leetcode-cn.com/problems/super-egg-drop/)|雇佣 K 名工人的最低成本|困难|20%|26786|
|[0886](https://leetcode-cn.com/problems/possible-bipartition/)|括号的分数|中等|33%|3452|
|[0885](https://leetcode-cn.com/problems/spiral-matrix-iii/)|考场就座|中等|60%|1792|
|[0884](https://leetcode-cn.com/problems/uncommon-words-from-two-sentences/)|相似度为 K 的字符串|简单|58%|10450|
|[0883](https://leetcode-cn.com/problems/projection-area-of-3d-shapes/)|车队|简单|62%|6389|
|[0882](https://leetcode-cn.com/problems/reachable-nodes-in-subdivided-graph/)|山脉数组的峰顶索引|困难|32%|762|
|[0881](https://leetcode-cn.com/problems/boats-to-save-people/)|喧闹和富有|中等|36%|10666|
|[0880](https://leetcode-cn.com/problems/decoded-string-at-index/)|矩形面积 II|中等|20%|7168|
|[0478](https://leetcode-cn.com/problems/generate-random-point-in-a-circle/)|Generate Random Point in a Circle|中等|37%|2808|
|[0497](https://leetcode-cn.com/problems/random-point-in-non-overlapping-rectangles/)|Random Point in Non-overlapping Rectangles|中等|34%|1413|
|[0519](https://leetcode-cn.com/problems/random-flip-matrix/)|Random Flip Matrix|中等|32%|1340|
|[0528](https://leetcode-cn.com/problems/random-pick-with-weight/)|Random Pick with Weight|中等|38%|2274|
|[0879](https://leetcode-cn.com/problems/profitable-schemes/)|到最近的人的最大距离|困难|27%|1744|
|[0878](https://leetcode-cn.com/problems/nth-magical-number/)|字母移位|困难|19%|5983|
|[0877](https://leetcode-cn.com/problems/stone-game/)|访问所有节点的最短路径|中等|66%|11528|
|[0876](https://leetcode-cn.com/problems/middle-of-the-linked-list/)|一手顺子|简单|62%|30818|
|[0875](https://leetcode-cn.com/problems/koko-eating-bananas/)|数组中的最长山脉|中等|39%|5888|
|[0874](https://leetcode-cn.com/problems/walking-robot-simulation/)|比较含退格的字符串|简单|30%|12835|
|[0873](https://leetcode-cn.com/problems/length-of-longest-fibonacci-subsequence/)|猜猜这个单词|中等|44%|6380|
|[0872](https://leetcode-cn.com/problems/leaf-similar-trees/)|将数组拆分成斐波那契序列|简单|60%|11303|
|[0470](https://leetcode-cn.com/problems/implement-rand10-using-rand7/)|Implement Rand10() Using Rand7()|中等|44%|6232|
|[0871](https://leetcode-cn.com/problems/minimum-number-of-refueling-stops/)|钥匙和房间|困难|28%|5849|
|[0870](https://leetcode-cn.com/problems/advantage-shuffle/)|矩阵中的幻方|中等|32%|8958|
|[0869](https://leetcode-cn.com/problems/reordered-power-of-2/)|相似字符串组|中等|46%|4112|
|[0868](https://leetcode-cn.com/problems/binary-gap/)|推多米诺|简单|58%|10593|
|[0867](https://leetcode-cn.com/problems/transpose-matrix/)|新21点|简单|65%|19636|
|[0866](https://leetcode-cn.com/problems/prime-palindrome/)|矩形重叠|中等|17%|11696|
|[0865](https://leetcode-cn.com/problems/smallest-subtree-with-all-the-deepest-nodes/)|扫地机器人|中等|46%|2750|
|[0864](https://leetcode-cn.com/problems/shortest-path-to-get-all-keys/)|图像重叠|困难|38%|1199|
|[0710](https://leetcode-cn.com/problems/random-pick-with-blacklist/)|Random Pick with Blacklist|困难|25%|2168|
|[0863](https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/)|树中距离之和|中等|43%|4853|
|[0862](https://leetcode-cn.com/problems/shortest-subarray-with-sum-at-least-k/)|字符串中的查找与替换|困难|11%|25939|
|[0861](https://leetcode-cn.com/problems/score-after-flipping-matrix/)|翻转图像|中等|71%|3982|
|[0860](https://leetcode-cn.com/problems/lemonade-change/)|设计循环队列|简单|52%|21119|
|[0859](https://leetcode-cn.com/problems/buddy-strings/)|设计循环双端队列|简单|27%|24750|
|[0858](https://leetcode-cn.com/problems/mirror-reflection/)|隐藏个人信息|中等|49%|1738|
|[0857](https://leetcode-cn.com/problems/minimum-cost-to-hire-k-workers/)|较大分组的位置|困难|37%|1336|
|[0856](https://leetcode-cn.com/problems/score-of-parentheses/)|连续整数求和|中等|55%|5249|
|[0855](https://leetcode-cn.com/problems/exam-room/)|独特字符串|中等|28%|2298|
|[0854](https://leetcode-cn.com/problems/k-similar-strings/)|最大人工岛|困难|25%|2927|
|[0853](https://leetcode-cn.com/problems/car-fleet/)|安排工作以达到最大收益|中等|33%|5373|
|[0852](https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/)|适龄的朋友|简单|68%|22960|
|[0851](https://leetcode-cn.com/problems/loud-and-rich/)|山羊拉丁文|中等|40%|2075|
|[0850](https://leetcode-cn.com/problems/rectangle-area-ii/)|循环有序列表的插入|困难|36%|1104|
|[0849](https://leetcode-cn.com/problems/maximize-distance-to-closest-person/)|Maximize Distance to Closest Person|简单|37%|14937|
|[0848](https://leetcode-cn.com/problems/shifting-letters/)|Shifting Letters|中等|36%|6809|
|[0847](https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/)|Shortest Path Visiting All Nodes|困难|46%|1590|
|[0846](https://leetcode-cn.com/problems/hand-of-straights/)|Hand of Straights|中等|43%|5401|
|[0845](https://leetcode-cn.com/problems/longest-mountain-in-array/)|Longest Mountain in Array|中等|32%|8322|
|[0844](https://leetcode-cn.com/problems/backspace-string-compare/)|Backspace String Compare|简单|48%|21008|
|[0843](https://leetcode-cn.com/problems/guess-the-word/)|带因子的二叉树|困难|30%|1527|
|[0842](https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/)|翻转卡片游戏|中等|32%|4487|
|[0841](https://leetcode-cn.com/problems/keys-and-rooms/)|字符的最短距离|中等|57%|9684|
|[0840](https://leetcode-cn.com/problems/magic-squares-in-grid/)|Magic Squares In Grid|简单|32%|9692|
|[0839](https://leetcode-cn.com/problems/similar-string-groups/)|单词的压缩编码|困难|27%|3276|
|[0838](https://leetcode-cn.com/problems/push-dominoes/)|设计链表|中等|39%|3595|
|[0837](https://leetcode-cn.com/problems/new-21-game/)|最常见的单词|中等|23%|4139|
|[0836](https://leetcode-cn.com/problems/rectangle-overlap/)|赛车|简单|43%|8804|
|[0489](https://leetcode-cn.com/problems/robot-room-cleaner/)|Robot Room Cleaner|困难|69%|208|
|[0835](https://leetcode-cn.com/problems/image-overlap/)|链表组件|中等|56%|1897|
|[0834](https://leetcode-cn.com/problems/sum-of-distances-in-tree/)|模糊坐标|困难|28%|1763|
|[0833](https://leetcode-cn.com/problems/find-and-replace-in-string/)|公交路线|中等|33%|3643|
|[0832](https://leetcode-cn.com/problems/flipping-an-image/)|二叉树剪枝|简单|73%|31429|
|[0622](https://leetcode-cn.com/problems/design-circular-queue/)|[Design Circular Queue](./src/0622-design-circular-queue/)|中等|38%|35702|
|[0641](https://leetcode-cn.com/problems/design-circular-deque/)|Design Circular Deque|中等|47%|4848|
|[0831](https://leetcode-cn.com/problems/masking-personal-information/)|最大平均值和的分组|中等|36%|7278|
|[0830](https://leetcode-cn.com/problems/positions-of-large-groups/)|最大三角形面积|简单|43%|13235|
|[0829](https://leetcode-cn.com/problems/consecutive-numbers-sum/)|子域名访问计数|困难|28%|8634|
|[0828](https://leetcode-cn.com/problems/unique-letter-string/)|黑板异或游戏|困难|34%|1934|
|[0827](https://leetcode-cn.com/problems/making-a-large-island/)|情感丰富的文字|困难|36%|1890|
|[0826](https://leetcode-cn.com/problems/most-profit-assigning-work/)|分汤|中等|35%|3880|
|[0825](https://leetcode-cn.com/problems/friends-of-appropriate-ages/)|保持城市天际线|中等|31%|4512|
|[0824](https://leetcode-cn.com/problems/goat-latin/)|写字符串需要的行数|简单|57%|9225|
|[0708](https://leetcode-cn.com/problems/insert-into-a-cyclic-sorted-list/)|Insert into a Cyclic Sorted List|中等|28%|479|
|[0823](https://leetcode-cn.com/problems/binary-trees-with-factors/)|数组的均值分割|中等|37%|1602|
|[0822](https://leetcode-cn.com/problems/card-flipping-game/)|唯一摩尔斯密码词|中等|45%|1406|
|[0821](https://leetcode-cn.com/problems/shortest-distance-to-a-character/)|打砖块|简单|65%|12147|
|[0820](https://leetcode-cn.com/problems/short-encoding-of-words/)|找到最终的安全状态|中等|38%|10825|
|[0707](https://leetcode-cn.com/problems/design-linked-list/)|Design Linked List|中等|22%|50353|
|[0819](https://leetcode-cn.com/problems/most-common-word/)|使序列递增的最小交换次数|简单|36%|12658|
|[0818](https://leetcode-cn.com/problems/race-car/)|相似 RGB 颜色|困难|29%|1780|
|[0817](https://leetcode-cn.com/problems/linked-list-components/)|设计哈希映射|中等|54%|5978|
|[0816](https://leetcode-cn.com/problems/ambiguous-coordinates/)|设计哈希集合|中等|45%|1383|
|[0815](https://leetcode-cn.com/problems/bus-routes/)|香槟塔|困难|33%|2970|
|[0814](https://leetcode-cn.com/problems/binary-tree-pruning/)|得分最高的最小轮调|中等|71%|5726|
|[0813](https://leetcode-cn.com/problems/largest-sum-of-averages/)|所有可能的路径|中等|44%|2722|
|[0812](https://leetcode-cn.com/problems/largest-triangle-area/)|旋转字符串|简单|57%|6576|
|[0811](https://leetcode-cn.com/problems/subdomain-visit-count/)|区间子数组个数|简单|64%|8976|
|[0810](https://leetcode-cn.com/problems/chalkboard-xor-game/)|有效的井字游戏|困难|50%|998|
|[0809](https://leetcode-cn.com/problems/expressive-words/)|阶乘函数后K个零|中等|36%|2020|
|[0808](https://leetcode-cn.com/problems/soup-servings/)|匹配子序列的单词数|中等|39%|2082|
|[0807](https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/)|自定义字符串排序|中等|78%|7316|
|[0806](https://leetcode-cn.com/problems/number-of-lines-to-write-string/)|多米诺和托米诺平铺|简单|62%|8349|
|[0805](https://leetcode-cn.com/problems/split-array-with-same-average/)|逃脱阻碍者|困难|22%|2452|
|[0804](https://leetcode-cn.com/problems/unique-morse-code-words/)|旋转数字|简单|73%|24213|
|[0803](https://leetcode-cn.com/problems/bricks-falling-when-hit/)|K 站中转内最便宜的航班|困难|21%|1400|
|[0802](https://leetcode-cn.com/problems/find-eventual-safe-states/)|第 K 个最小的素数分数|中等|42%|2666|
|[0801](https://leetcode-cn.com/problems/minimum-swaps-to-make-sequences-increasing/)|判断二分图|中等|35%|3435|
|[0800](https://leetcode-cn.com/problems/similar-rgb-color/)|字母大小写全排列|简单|61%|252|
|[0706](https://leetcode-cn.com/problems/design-hashmap/)|Design HashMap|简单|55%|10759|
|[0705](https://leetcode-cn.com/problems/design-hashset/)|Design HashSet|简单|54%|12983|
|[0799](https://leetcode-cn.com/problems/champagne-tower/)|二叉搜索树结点最小距离|中等|32%|2582|
|[0798](https://leetcode-cn.com/problems/smallest-rotation-with-highest-score/)|变为棋盘|困难|32%|818|
|[0797](https://leetcode-cn.com/problems/all-paths-from-source-to-target/)|森林中的兔子|中等|72%|2465|
|[0796](https://leetcode-cn.com/problems/rotate-string/)|到达终点|简单|48%|11942|
|[0795](https://leetcode-cn.com/problems/number-of-subarrays-with-bounded-maximum/)|第K个语法符号|中等|45%|3496|
|[0794](https://leetcode-cn.com/problems/valid-tic-tac-toe-state/)|水位上升的泳池中游泳|中等|28%|4399|
|[0793](https://leetcode-cn.com/problems/preimage-size-of-factorial-zeroes-function/)|在LR字符串中交换相邻字符|困难|30%|2556|
|[0792](https://leetcode-cn.com/problems/number-of-matching-subsequences/)|二分查找|中等|39%|4469|
|[0791](https://leetcode-cn.com/problems/custom-sort-string/)|拆分二叉搜索树|中等|62%|4078|
|[0790](https://leetcode-cn.com/problems/domino-and-tromino-tiling/)|全局倒置与局部倒置|中等|35%|2140|
|[0789](https://leetcode-cn.com/problems/escape-the-ghosts/)|数据流中的第K大元素|中等|57%|2704|
|[0788](https://leetcode-cn.com/problems/rotated-digits/)|最小化去加油站的最大距离|简单|57%|11511|
|[0787](https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/)|滑动谜题|中等|32%|6439|
|[0786](https://leetcode-cn.com/problems/k-th-smallest-prime-fraction/)|搜索长度未知的有序数组|困难|32%|1752|
|[0785](https://leetcode-cn.com/problems/is-graph-bipartite/)|基本计算器 III|中等|40%|5809|
|[0784](https://leetcode-cn.com/problems/letter-case-permutation/)|二叉搜索树中的插入操作|简单|60%|15525|
|[0783](https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/)|二叉搜索树中的搜索|简单|51%|11469|
|[0782](https://leetcode-cn.com/problems/transform-to-chessboard/)|宝石与石头|困难|38%|741|
|[0781](https://leetcode-cn.com/problems/rabbits-in-forest/)|基本计算器 IV|中等|51%|3717|
|[0780](https://leetcode-cn.com/problems/reaching-points/)|最多能完成排序的块|困难|26%|2571|
|[0779](https://leetcode-cn.com/problems/k-th-symbol-in-grammar/)|最多能完成排序的块 II|中等|40%|8940|
|[0778](https://leetcode-cn.com/problems/swim-in-rising-water/)|重构字符串|困难|44%|2327|
|[0777](https://leetcode-cn.com/problems/swap-adjacent-in-lr-string/)|托普利茨矩阵|中等|30%|4925|
|[0704](https://leetcode-cn.com/problems/binary-search/)|Binary Search|简单|50%|40919|
|[0776](https://leetcode-cn.com/problems/split-bst/)|N叉树的后序遍历|中等|48%|330|
|[0775](https://leetcode-cn.com/problems/global-and-local-inversions/)|N叉树的前序遍历|中等|39%|3161|
|[0703](https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/)|Kth Largest Element in a Stream|简单|40%|22675|
|[0774](https://leetcode-cn.com/problems/minimize-max-distance-to-gas-station/)|N叉树的最大深度|困难|54%|206|
|[0773](https://leetcode-cn.com/problems/sliding-puzzle/)|四叉树交集|困难|54%|1683|
|[0702](https://leetcode-cn.com/problems/search-in-a-sorted-array-of-unknown-size/)|Search in a Sorted Array of Unknown Size|中等|67%|406|
|[0772](https://leetcode-cn.com/problems/basic-calculator-iii/)|建立四叉树|困难|27%|909|
|[0701](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)|Insert into a Binary Search Tree|中等|69%|7666|
|[0700](https://leetcode-cn.com/problems/search-in-a-binary-search-tree/)|Search in a Binary Search Tree|简单|70%|17542|
|[0771](https://leetcode-cn.com/problems/jewels-and-stones/)|将 N 叉树编码为二叉树|简单|80%|69075|
|[0770](https://leetcode-cn.com/problems/basic-calculator-iv/)|情侣牵手|困难|46%|329|
|[0769](https://leetcode-cn.com/problems/max-chunks-to-make-sorted/)|最大加号标志|中等|51%|3267|
|[0768](https://leetcode-cn.com/problems/max-chunks-to-make-sorted-ii/)|划分字母区间|困难|42%|2736|
|[0767](https://leetcode-cn.com/problems/reorganize-string/)|二进制表示中质数个计算置位|中等|36%|6449|
|[0766](https://leetcode-cn.com/problems/toeplitz-matrix/)|扁平化多级双向链表|简单|61%|12934|
|[0590](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)|N-ary Tree Postorder Traversal|简单|70%|15593|
|[0589](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)|N-ary Tree Preorder Traversal|简单|70%|18049|
|[0559](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/)|Maximum Depth of N-ary Tree|简单|66%|17924|
|[0558](https://leetcode-cn.com/problems/quad-tree-intersection/)|Quad Tree Intersection|简单|44%|2993|
|[0427](https://leetcode-cn.com/problems/construct-quad-tree/)|Construct Quad Tree|中等|56%|3011|
|[0431](https://leetcode-cn.com/problems/encode-n-ary-tree-to-binary-tree/)|Encode N-ary Tree to Binary Tree|困难|63%|84|
|[0765](https://leetcode-cn.com/problems/couples-holding-hands/)|序列化和反序列化 N 叉树|困难|53%|5224|
|[0764](https://leetcode-cn.com/problems/largest-plus-sign/)|N叉树的层序遍历|中等|47%|2205|
|[0763](https://leetcode-cn.com/problems/partition-labels/)|特殊的二进制序列|中等|67%|6999|
|[0762](https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/)|找出变位映射|简单|64%|9090|
|[0430](https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/)|Flatten a Multilevel Doubly Linked List|中等|42%|9205|
|[0428](https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree/)|Serialize and Deserialize N-ary Tree|困难|67%|164|
|[0429](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)|N-ary Tree Level Order Traversal|简单|63%|17333|
|[0761](https://leetcode-cn.com/problems/special-binary-string/)|员工空闲时间|困难|51%|702|
|[0760](https://leetcode-cn.com/problems/find-anagram-mappings/)|字符串中的加粗单词|简单|78%|690|
|[0759](https://leetcode-cn.com/problems/employee-free-time/)| 设置交集大小至少为2|困难|49%|171|
|[0758](https://leetcode-cn.com/problems/bold-words-in-string/)|将二叉搜索树转化为排序的双向链表|简单|39%|671|
|[0757](https://leetcode-cn.com/problems/set-intersection-size-at-least-two/)|金字塔转换矩阵|困难|32%|1364|
|[0426](https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/)|Convert Binary Search Tree to Sorted Doubly Linked List|中等|55%|970|
|[0756](https://leetcode-cn.com/problems/pyramid-transition-matrix/)|倒水|中等|53%|1842|
|[0755](https://leetcode-cn.com/problems/pour-water/)|到达终点数字|中等|45%|361|
|[0754](https://leetcode-cn.com/problems/reach-a-number/)|破解保险箱|简单|38%|8337|
|[0753](https://leetcode-cn.com/problems/cracking-the-safe/)|打开转盘锁|困难|53%|751|
|[0752](https://leetcode-cn.com/problems/open-the-lock/)|IP 到 CIDR|中等|47%|11727|
|[0751](https://leetcode-cn.com/problems/ip-to-cidr/)|角矩形的数量|简单|57%|441|
|[0750](https://leetcode-cn.com/problems/number-of-corner-rectangles/)|隔离病毒|中等|70%|350|
|[0749](https://leetcode-cn.com/problems/contain-virus/)|最短完整词|困难|50%|501|
|[0748](https://leetcode-cn.com/problems/shortest-completing-word/)|至少是其他数字两倍的最大数|简单|55%|5469|
|[0747](https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/)|使用最小花费爬楼梯|简单|38%|41328|
|[0746](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)|[前缀和后缀搜索](./src/0746-min-cost-climbing-stairs/)|简单|45%|36749|
|[0745](https://leetcode-cn.com/problems/prefix-and-suffix-search/)|寻找比目标字母大的最小字母|困难|28%|1832|
|[0744](https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/)|网络延迟时间|简单|42%|19013|
|[0743](https://leetcode-cn.com/problems/network-delay-time/)|二叉树最近的叶节点|中等|44%|8102|
|[0742](https://leetcode-cn.com/problems/closest-leaf-in-a-binary-tree/)|转换成小写字母|中等|41%|329|
|[0709](https://leetcode-cn.com/problems/to-lower-case/)|To Lower Case|简单|74%|40044|
|[0741](https://leetcode-cn.com/problems/cherry-pickup/)|摘樱桃|困难|28%|2275|
|[0740](https://leetcode-cn.com/problems/delete-and-earn/)|删除与获得点数|中等|49%|4439|
|[0739](https://leetcode-cn.com/problems/daily-temperatures/)|[每日温度](./src/0739-daily-temperatures/)|中等|55%|31576|
|[0738](https://leetcode-cn.com/problems/monotone-increasing-digits/)|单调递增的数字|中等|40%|5221|
|[0737](https://leetcode-cn.com/problems/sentence-similarity-ii/)|句子相似性 II|中等|34%|564|
|[0736](https://leetcode-cn.com/problems/parse-lisp-expression/)|Lisp 语法解析|困难|39%|576|
|[0735](https://leetcode-cn.com/problems/asteroid-collision/)|行星碰撞|中等|35%|8309|
|[0734](https://leetcode-cn.com/problems/sentence-similarity/)|句子相似性|简单|49%|414|
|[0733](https://leetcode-cn.com/problems/flood-fill/)|图像渲染|简单|51%|11953|
|[0732](https://leetcode-cn.com/problems/my-calendar-iii/)|我的日程安排表 III|困难|51%|1105|
|[0731](https://leetcode-cn.com/problems/my-calendar-ii/)|我的日程安排表 II|中等|42%|1827|
|[0730](https://leetcode-cn.com/problems/count-different-palindromic-subsequences/)|统计不同回文子字符串|困难|40%|1106|
|[0729](https://leetcode-cn.com/problems/my-calendar-i/)|我的日程安排表 I|中等|43%|3431|
|[0728](https://leetcode-cn.com/problems/self-dividing-numbers/)|自除数|简单|70%|19200|
|[0727](https://leetcode-cn.com/problems/minimum-window-subsequence/)|最小窗口子序列|困难|35%|495|
|[0726](https://leetcode-cn.com/problems/number-of-atoms/)|原子的数量|困难|41%|2428|
|[0725](https://leetcode-cn.com/problems/split-linked-list-in-parts/)|分隔链表|中等|50%|7033|
|[0724](https://leetcode-cn.com/problems/find-pivot-index/)|寻找数组的中心索引|简单|35%|54179|
|[0723](https://leetcode-cn.com/problems/candy-crush/)|粉碎糖果|中等|57%|250|
|[0722](https://leetcode-cn.com/problems/remove-comments/)|删除注释|中等|26%|3441|
|[0721](https://leetcode-cn.com/problems/accounts-merge/)|账户合并|中等|30%|4238|
|[0720](https://leetcode-cn.com/problems/longest-word-in-dictionary/)|词典中最长的单词|简单|43%|9478|
|[0719](https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance/)|找出第 k 小的距离对|困难|28%|5395|
|[0718](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/)|最长重复子数组|中等|47%|12583|
|[0717](https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/)|1比特与2比特字符|简单|44%|22310|
|[0716](https://leetcode-cn.com/problems/max-stack/)|最大栈|简单|37%|667|
|[0715](https://leetcode-cn.com/problems/range-module/)|Range 模块|困难|30%|1181|
|[0714](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)|买卖股票的最佳时机含手续费|中等|58%|10363|
|[0713](https://leetcode-cn.com/problems/subarray-product-less-than-k/)|乘积小于K的子数组|中等|32%|8139|
|[0712](https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/)|两个字符串的最小ASCII删除和|中等|59%|4194|
|[0711](https://leetcode-cn.com/problems/number-of-distinct-islands-ii/)|不同岛屿的数量 II|困难|43%|187|
|[0699](https://leetcode-cn.com/problems/falling-squares/)|掉落的方块|困难|40%|931|
|[0698](https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/)|划分为k个相等的子集|中等|38%|9507|
|[0697](https://leetcode-cn.com/problems/degree-of-an-array/)|数组的度|简单|50%|13514|
|[0696](https://leetcode-cn.com/problems/count-binary-substrings/)|计数二进制子串|简单|47%|10890|
|[0695](https://leetcode-cn.com/problems/max-area-of-island/)|岛屿的最大面积|中等|57%|20240|
|[0694](https://leetcode-cn.com/problems/number-of-distinct-islands/)|不同岛屿的数量|中等|49%|551|
|[0693](https://leetcode-cn.com/problems/binary-number-with-alternating-bits/)|交替位二进制数|简单|60%|14161|
|[0692](https://leetcode-cn.com/problems/top-k-frequent-words/)|前K个高频单词|中等|44%|9668|
|[0691](https://leetcode-cn.com/problems/stickers-to-spell-word/)|贴纸拼词|困难|47%|1085|
|[0690](https://leetcode-cn.com/problems/employee-importance/)|员工的重要性|简单|54%|11372|
|[0689](https://leetcode-cn.com/problems/maximum-sum-of-3-non-overlapping-subarrays/)|三个无重叠子数组的最大和|困难|38%|1362|
|[0688](https://leetcode-cn.com/problems/knight-probability-in-chessboard/)|“马”在棋盘上的概率|中等|44%|3381|
|[0687](https://leetcode-cn.com/problems/longest-univalue-path/)|最长同值路径|简单|37%|18949|
|[0686](https://leetcode-cn.com/problems/repeated-string-match/)|重复叠加字符串匹配|简单|32%|15847|
|[0685](https://leetcode-cn.com/problems/redundant-connection-ii/)|冗余连接 II|困难|31%|3156|
|[0684](https://leetcode-cn.com/problems/redundant-connection/)|冗余连接|中等|53%|6956|
|[0683](https://leetcode-cn.com/problems/k-empty-slots/)|K 个空花盆|困难|43%|453|
|[0682](https://leetcode-cn.com/problems/baseball-game/)|棒球比赛|简单|64%|18781|
|[0681](https://leetcode-cn.com/problems/next-closest-time/)|最近时刻|中等|50%|350|
|[0680](https://leetcode-cn.com/problems/valid-palindrome-ii/)|验证回文字符串 Ⅱ|简单|33%|24426|
|[0679](https://leetcode-cn.com/problems/24-game/)|24 点游戏|困难|40%|4584|
|[0678](https://leetcode-cn.com/problems/valid-parenthesis-string/)|有效的括号字符串|中等|27%|6816|
|[0677](https://leetcode-cn.com/problems/map-sum-pairs/)|键值映射|中等|59%|5555|
|[0676](https://leetcode-cn.com/problems/implement-magic-dictionary/)|实现一个魔法字典|中等|53%|1976|
|[0675](https://leetcode-cn.com/problems/cut-off-trees-for-golf-event/)|为高尔夫比赛砍树|困难|29%|1324|
|[0674](https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/)|最长连续递增序列|简单|43%|28814|
|[0673](https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/)|最长递增子序列的个数|中等|31%|10357|
|[0672](https://leetcode-cn.com/problems/bulb-switcher-ii/)|灯泡开关 Ⅱ|中等|50%|1617|
|[0671](https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/)|二叉树中第二小的节点|简单|44%|11948|
|[0670](https://leetcode-cn.com/problems/maximum-swap/)|最大交换|中等|37%|6789|
|[0669](https://leetcode-cn.com/problems/trim-a-binary-search-tree/)|修剪二叉搜索树|简单|63%|10081|
|[0668](https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table/)|乘法表中第k小的数|困难|42%|3116|
|[0667](https://leetcode-cn.com/problems/beautiful-arrangement-ii/)|优美的排列 II|中等|56%|2927|
|[0666](https://leetcode-cn.com/problems/path-sum-iv/)|路径和 IV|中等|60%|319|
|[0665](https://leetcode-cn.com/problems/non-decreasing-array/)|非递减数列|简单|20%|45556|
|[0664](https://leetcode-cn.com/problems/strange-printer/)|奇怪的打印机|困难|38%|1855|
|[0663](https://leetcode-cn.com/problems/equal-tree-partition/)|均匀树划分|中等|36%|477|
|[0662](https://leetcode-cn.com/problems/maximum-width-of-binary-tree/)|二叉树最大宽度|中等|33%|8900|
|[0661](https://leetcode-cn.com/problems/image-smoother/)|图片平滑器|简单|50%|8821|
|[0660](https://leetcode-cn.com/problems/remove-9/)|移除 9|困难|45%|133|
|[0659](https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/)|分割数组为连续子序列|中等|40%|3305|
|[0658](https://leetcode-cn.com/problems/find-k-closest-elements/)|找到 K 个最接近的元素|中等|41%|9864|
|[0657](https://leetcode-cn.com/problems/robot-return-to-origin/)|机器人能否返回原点|简单|72%|30819|
|[0656](https://leetcode-cn.com/problems/coin-path/)|金币路径|困难|20%|507|
|[0655](https://leetcode-cn.com/problems/print-binary-tree/)|输出二叉树|中等|51%|3200|
|[0654](https://leetcode-cn.com/problems/maximum-binary-tree/)|最大二叉树|中等|77%|9560|
|[0653](https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/)|两数之和 IV - 输入 BST|简单|52%|14713|
|[0652](https://leetcode-cn.com/problems/find-duplicate-subtrees/)|寻找重复的子树|中等|50%|5823|
|[0651](https://leetcode-cn.com/problems/4-keys-keyboard/)|4键键盘|中等|55%|677|
|[0650](https://leetcode-cn.com/problems/2-keys-keyboard/)|只有两个键的键盘|中等|45%|9413|
|[0649](https://leetcode-cn.com/problems/dota2-senate/)|Dota2 参议院|中等|35%|4270|
|[0648](https://leetcode-cn.com/problems/replace-words/)|单词替换|中等|52%|4861|
|[0647](https://leetcode-cn.com/problems/palindromic-substrings/)|回文子串|中等|58%|17976|
|[0646](https://leetcode-cn.com/problems/maximum-length-of-pair-chain/)|最长数对链|中等|51%|6331|
|[0645](https://leetcode-cn.com/problems/set-mismatch/)|错误的集合|简单|38%|17220|
|[0644](https://leetcode-cn.com/problems/maximum-average-subarray-ii/)|最大平均子段和 II|困难|36%|241|
|[0643](https://leetcode-cn.com/problems/maximum-average-subarray-i/)|子数组最大平均数 I|简单|36%|21040|
|[0642](https://leetcode-cn.com/problems/design-search-autocomplete-system/)|设计搜索自动补全系统|困难|43%|327|
|[0640](https://leetcode-cn.com/problems/solve-the-equation/)|求解方程|中等|39%|2612|
|[0639](https://leetcode-cn.com/problems/decode-ways-ii/)|解码方法 2|困难|25%|2828|
|[0638](https://leetcode-cn.com/problems/shopping-offers/)|大礼包|中等|55%|4140|
|[0637](https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/)|二叉树的层平均值|简单|61%|15380|
|[0636](https://leetcode-cn.com/problems/exclusive-time-of-functions/)|函数的独占时间|中等|47%|2572|
|[0635](https://leetcode-cn.com/problems/design-log-storage-system/)|设计日志存储系统|中等|53%|294|
|[0634](https://leetcode-cn.com/problems/find-the-derangement-of-an-array/)|寻找数组的错位排列|中等|42%|145|
|[0633](https://leetcode-cn.com/problems/sum-of-square-numbers/)|平方数之和|简单|31%|29845|
|[0632](https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/)|最小区间|困难|32%|4311|
|[0631](https://leetcode-cn.com/problems/design-excel-sum-formula/)|设计 Excel 求和公式|困难|21%|343|
|[0630](https://leetcode-cn.com/problems/course-schedule-iii/)|课程表 III|困难|28%|2921|
|[0629](https://leetcode-cn.com/problems/k-inverse-pairs-array/)|K个逆序对数组|困难|35%|1700|
|[0628](https://leetcode-cn.com/problems/maximum-product-of-three-numbers/)|三个数的最大乘积|简单|47%|22335|
|[0627](https://leetcode-cn.com/problems/swap-salary/)|[交换工资](./src/0627-swap-salary/)|简单|72%|31097|
|[0626](https://leetcode-cn.com/problems/exchange-seats/)|[换座位](./src/0626-exchange-seats/)|中等|64%|17715|
|[0625](https://leetcode-cn.com/problems/minimum-factorization/)|最小因式分解|中等|30%|858|
|[0624](https://leetcode-cn.com/problems/maximum-distance-in-arrays/)|数组列表中的最大距离|简单|39%|681|
|[0623](https://leetcode-cn.com/problems/add-one-row-to-tree/)|在二叉树中增加一行|中等|48%|4000|
|[0621](https://leetcode-cn.com/problems/task-scheduler/)|任务调度器|中等|45%|17326|
|[0620](https://leetcode-cn.com/problems/not-boring-movies/)|[有趣的电影](./src/0620-not-boring-movies/)|简单|72%|36036|
|[0619](https://leetcode-cn.com/problems/biggest-single-number/)|只出现一次的最大数字|简单|44%|2110|
|[0618](https://leetcode-cn.com/problems/students-report-by-geography/)|学生地理信息报告|困难|44%|448|
|[0617](https://leetcode-cn.com/problems/merge-two-binary-trees/)|合并二叉树|简单|73%|31700|
|[0616](https://leetcode-cn.com/problems/add-bold-tag-in-string/)|给字符串添加加粗标签|中等|45%|949|
|[0615](https://leetcode-cn.com/problems/average-salary-departments-vs-company/)|平均工资：部门与公司比较|困难|39%|822|
|[0614](https://leetcode-cn.com/problems/second-degree-follower/)|二级关注者|中等|25%|1888|
|[0613](https://leetcode-cn.com/problems/shortest-distance-in-a-line/)|直线上的最近距离|简单|79%|1324|
|[0612](https://leetcode-cn.com/problems/shortest-distance-in-a-plane/)|平面上的最近距离|中等|61%|724|
|[0611](https://leetcode-cn.com/problems/valid-triangle-number/)|有效三角形的个数|中等|46%|6391|
|[0610](https://leetcode-cn.com/problems/triangle-judgement/)|判断三角形|简单|61%|1314|
|[0609](https://leetcode-cn.com/problems/find-duplicate-file-in-system/)|在系统中查找重复文件|中等|54%|1837|
|[0608](https://leetcode-cn.com/problems/tree-node/)|树节点|中等|63%|798|
|[0607](https://leetcode-cn.com/problems/sales-person/)|销售员|简单|63%|1401|
|[0606](https://leetcode-cn.com/problems/construct-string-from-binary-tree/)|根据二叉树创建字符串|简单|51%|11985|
|[0605](https://leetcode-cn.com/problems/can-place-flowers/)|种花问题|简单|30%|37119|
|[0604](https://leetcode-cn.com/problems/design-compressed-string-iterator/)|迭代压缩字符串|简单|30%|494|
|[0603](https://leetcode-cn.com/problems/consecutive-available-seats/)|连续空余座位|简单|64%|1343|
|[0602](https://leetcode-cn.com/problems/friend-requests-ii-who-has-the-most-friends/)|好友申请 II ：谁有最多的好友|中等|53%|899|
|[0601](https://leetcode-cn.com/problems/human-traffic-of-stadium/)|[体育馆的人流量](./src/0601-human-traffic-of-stadium/)|困难|43%|17262|
|[0600](https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/)|[不含连续1的非负整数](./src/0600-non-negative-integers-without-consecutive-ones/)|困难|26%|3107|
|[0599](https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/)|两个列表的最小索引总和|简单|47%|14876|
|[0598](https://leetcode-cn.com/problems/range-addition-ii/)|范围求和 II|简单|49%|7803|
|[0597](https://leetcode-cn.com/problems/friend-requests-i-overall-acceptance-rate/)|好友申请 I ：总体通过率|简单|41%|1677|
|[0596](https://leetcode-cn.com/problems/classes-more-than-5-students/)|[超过5名学生的课](./src/0596-classes-more-than-5-students/)|简单|38%|53317|
|[0595](https://leetcode-cn.com/problems/big-countries/)|[大的国家](./src/0595-big-countries/)|简单|75%|42065|
|[0594](https://leetcode-cn.com/problems/longest-harmonious-subsequence/)|最长和谐子序列|简单|43%|12416|
|[0593](https://leetcode-cn.com/problems/valid-square/)|有效的正方形|中等|40%|4715|
|[0592](https://leetcode-cn.com/problems/fraction-addition-and-subtraction/)|分数加减运算|中等|47%|1658|
|[0591](https://leetcode-cn.com/problems/tag-validator/)|标签验证器|困难|28%|809|
|[0588](https://leetcode-cn.com/problems/design-in-memory-file-system/)|设计内存文件系统|困难|35%|513|
|[0587](https://leetcode-cn.com/problems/erect-the-fence/)|安装栅栏|困难|28%|1045|
|[0586](https://leetcode-cn.com/problems/customer-placing-the-largest-number-of-orders/)|订单最多的客户|简单|71%|1546|
|[0585](https://leetcode-cn.com/problems/investments-in-2016/)|2016年的投资|中等|50%|999|
|[0584](https://leetcode-cn.com/problems/find-customer-referee/)|寻找用户推荐人|简单|73%|1510|
|[0583](https://leetcode-cn.com/problems/delete-operation-for-two-strings/)|两个字符串的删除操作|中等|46%|5918|
|[0582](https://leetcode-cn.com/problems/kill-process/)|杀死进程|中等|42%|403|
|[0581](https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/)|最短无序连续子数组|简单|33%|32830|
|[0580](https://leetcode-cn.com/problems/count-student-number-in-departments/)|统计各专业学生人数|中等|45%|1486|
|[0579](https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee/)|查询员工的累计薪水|困难|36%|1204|
|[0578](https://leetcode-cn.com/problems/get-highest-answer-rate-question/)|查询回答率最高的问题|中等|41%|1329|
|[0577](https://leetcode-cn.com/problems/employee-bonus/)|员工奖金|简单|68%|1678|
|[0576](https://leetcode-cn.com/problems/out-of-boundary-paths/)|出界的路径数|中等|34%|4411|
|[0575](https://leetcode-cn.com/problems/distribute-candies/)|分糖果|简单|63%|17006|
|[0574](https://leetcode-cn.com/problems/winning-candidate/)|当选者|中等|42%|1754|
|[0573](https://leetcode-cn.com/problems/squirrel-simulation/)|松鼠模拟|中等|56%|141|
|[0572](https://leetcode-cn.com/problems/subtree-of-another-tree/)|另一个树的子树|简单|41%|21606|
|[0571](https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers/)|给定数字的频率查询中位数|困难|45%|566|
|[0570](https://leetcode-cn.com/problems/managers-with-at-least-5-direct-reports/)|至少有5名直接下属的经理|中等|63%|1345|
|[0569](https://leetcode-cn.com/problems/median-employee-salary/)|员工薪水中位数|困难|41%|962|
|[0568](https://leetcode-cn.com/problems/maximum-vacation-days/)|最大休假天数|困难|42%|247|
|[0567](https://leetcode-cn.com/problems/permutation-in-string/)|字符串的排列|中等|32%|30163|
|[0566](https://leetcode-cn.com/problems/reshape-the-matrix/)|重塑矩阵|简单|62%|15815|
|[0565](https://leetcode-cn.com/problems/array-nesting/)|数组嵌套|中等|51%|4600|
|[0564](https://leetcode-cn.com/problems/find-the-closest-palindrome/)|寻找最近的回文数|困难|12%|7701|
|[0563](https://leetcode-cn.com/problems/binary-tree-tilt/)|二叉树的坡度|简单|50%|11431|
|[0562](https://leetcode-cn.com/problems/longest-line-of-consecutive-one-in-matrix/)|矩阵中最长的连续1线段|中等|35%|1626|
|[0561](https://leetcode-cn.com/problems/array-partition-i/)|数组拆分 I|简单|67%|32127|
|[0560](https://leetcode-cn.com/problems/subarray-sum-equals-k/)|和为K的子数组|中等|41%|22179|
|[0557](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)|反转字符串中的单词 III|简单|68%|46450|
|[0556](https://leetcode-cn.com/problems/next-greater-element-iii/)|下一个更大元素 III|中等|27%|6085|
|[0555](https://leetcode-cn.com/problems/split-concatenated-strings/)|分割连接字符串|中等|40%|335|
|[0554](https://leetcode-cn.com/problems/brick-wall/)|砖墙|中等|34%|13089|
|[0553](https://leetcode-cn.com/problems/optimal-division/)|最优除法|中等|52%|2463|
|[0552](https://leetcode-cn.com/problems/student-attendance-record-ii/)|学生出勤记录 II|困难|36%|2760|
|[0551](https://leetcode-cn.com/problems/student-attendance-record-i/)|学生出勤记录 I|简单|49%|16407|
|[0549](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence-ii/)|二叉树中最长的连续序列|中等|48%|311|
|[0548](https://leetcode-cn.com/problems/split-array-with-equal-sum/)|将数组分割成和相等的子数组|中等|28%|794|
|[0547](https://leetcode-cn.com/problems/friend-circles/)|朋友圈|中等|52%|28686|
|[0546](https://leetcode-cn.com/problems/remove-boxes/)|移除盒子|困难|49%|2148|
|[0545](https://leetcode-cn.com/problems/boundary-of-binary-tree/)|二叉树的边界|中等|30%|502|
|[0544](https://leetcode-cn.com/problems/output-contest-matches/)|输出比赛匹配对|中等|68%|307|
|[0543](https://leetcode-cn.com/problems/diameter-of-binary-tree/)|二叉树的直径|简单|46%|25595|
|[0542](https://leetcode-cn.com/problems/01-matrix/)|01 矩阵|中等|35%|17052|
|[0541](https://leetcode-cn.com/problems/reverse-string-ii/)|反转字符串 II|简单|49%|15925|
|[0540](https://leetcode-cn.com/problems/single-element-in-a-sorted-array/)|有序数组中的单一元素|中等|61%|7529|
|[0539](https://leetcode-cn.com/problems/minimum-time-difference/)|最小时间差|中等|50%|4670|
|[0538](https://leetcode-cn.com/problems/convert-bst-to-greater-tree/)|把二叉搜索树转换为累加树|简单|57%|17682|
|[0537](https://leetcode-cn.com/problems/complex-number-multiplication/)|复数乘法|中等|66%|4684|
|[0536](https://leetcode-cn.com/problems/construct-binary-tree-from-string/)|从字符串生成二叉树|中等|46%|553|
|[0535](https://leetcode-cn.com/problems/encode-and-decode-tinyurl/)|TinyURL 的加密与解密|中等|80%|6078|
|[0533](https://leetcode-cn.com/problems/lonely-pixel-ii/)|孤独像素 II|中等|40%|317|
|[0532](https://leetcode-cn.com/problems/k-diff-pairs-in-an-array/)|数组中的K-diff数对|简单|32%|21914|
|[0531](https://leetcode-cn.com/problems/lonely-pixel-i/)|孤独像素 I|中等|65%|388|
|[0530](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)|二叉搜索树的最小绝对差|简单|54%|12032|
|[0529](https://leetcode-cn.com/problems/minesweeper/)|扫雷游戏|中等|57%|3993|
|[0527](https://leetcode-cn.com/problems/word-abbreviation/)|单词缩写|困难|45%|746|
|[0526](https://leetcode-cn.com/problems/beautiful-arrangement/)|优美的排列|中等|55%|4523|
|[0525](https://leetcode-cn.com/problems/contiguous-array/)|连续数组|中等|38%|5327|
|[0524](https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/)|通过删除字母匹配到字典里最长单词|中等|44%|9358|
|[0523](https://leetcode-cn.com/problems/continuous-subarray-sum/)|连续的子数组和|中等|22%|18038|
|[0522](https://leetcode-cn.com/problems/longest-uncommon-subsequence-ii/)|最长特殊序列 II|中等|30%|3481|
|[0521](https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/)|最长特殊序列 Ⅰ|简单|64%|10351|
|[0520](https://leetcode-cn.com/problems/detect-capital/)|检测大写字母|简单|53%|21658|
|[0518](https://leetcode-cn.com/problems/coin-change-2/)|零钱兑换 II|中等|45%|9813|
|[0517](https://leetcode-cn.com/problems/super-washing-machines/)|超级洗衣机|困难|36%|1947|
|[0516](https://leetcode-cn.com/problems/longest-palindromic-subsequence/)|最长回文子序列|中等|46%|17846|
|[0515](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)|在每个树行中找最大值|中等|56%|7650|
|[0514](https://leetcode-cn.com/problems/freedom-trail/)|自由之路|困难|36%|2121|
|[0513](https://leetcode-cn.com/problems/find-bottom-left-tree-value/)|找树左下角的值|中等|66%|9371|
|[1059](https://leetcode-cn.com/problems/all-paths-from-source-lead-to-destination/)|有序数组中的缺失元素|中等|25%|1052|
|[0510](https://leetcode-cn.com/problems/inorder-successor-in-bst-ii/)|Inorder Successor in BST II|中等|55%|320|
|[0508](https://leetcode-cn.com/problems/most-frequent-subtree-sum/)|出现次数最多的子树元素和|中等|59%|4030|
|[0507](https://leetcode-cn.com/problems/perfect-number/)|完美数|简单|36%|22488|
|[0506](https://leetcode-cn.com/problems/relative-ranks/)|相对名次|简单|51%|10632|
|[0505](https://leetcode-cn.com/problems/the-maze-ii/)|迷宫 II|中等|40%|755|
|[0504](https://leetcode-cn.com/problems/base-7/)|七进制数|简单|46%|13734|
|[0503](https://leetcode-cn.com/problems/next-greater-element-ii/)|[下一个更大元素 II](./src/0503-next-greater-element-ii/)|中等|49%|10080|
|[0502](https://leetcode-cn.com/problems/ipo/)|IPO|困难|35%|2853|
|[0501](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)|二叉搜索树中的众数|简单|42%|13360|
|[0500](https://leetcode-cn.com/problems/keyboard-row/)|键盘行|简单|67%|16193|
|[0499](https://leetcode-cn.com/problems/the-maze-iii/)|迷宫 III|困难|39%|258|
|[0498](https://leetcode-cn.com/problems/diagonal-traverse/)|对角线遍历|中等|38%|19176|
|[0496](https://leetcode-cn.com/problems/next-greater-element-i/)|下一个更大元素 I|简单|62%|23935|
|[0495](https://leetcode-cn.com/problems/teemo-attacking/)|提莫攻击|中等|51%|9367|
|[0494](https://leetcode-cn.com/problems/target-sum/)|目标和|中等|41%|25540|
|[0493](https://leetcode-cn.com/problems/reverse-pairs/)|翻转对|困难|19%|8461|
|[0492](https://leetcode-cn.com/problems/construct-the-rectangle/)|构造矩形|简单|48%|10911|
|[0491](https://leetcode-cn.com/problems/increasing-subsequences/)|递增子序列|中等|43%|5647|
|[0490](https://leetcode-cn.com/problems/the-maze/)|迷宫|中等|42%|1033|
|[0488](https://leetcode-cn.com/problems/zuma-game/)|祖玛游戏|困难|39%|1550|
|[0487](https://leetcode-cn.com/problems/max-consecutive-ones-ii/)|最大连续1的个数 II|中等|51%|599|
|[0486](https://leetcode-cn.com/problems/predict-the-winner/)|预测赢家|中等|48%|7096|
|[0485](https://leetcode-cn.com/problems/max-consecutive-ones/)|最大连续1的个数|简单|54%|37394|
|[0484](https://leetcode-cn.com/problems/find-permutation/)|寻找排列|中等|58%|207|
|[0483](https://leetcode-cn.com/problems/smallest-good-base/)|最小好进制|困难|34%|1574|
|[0482](https://leetcode-cn.com/problems/license-key-formatting/)|密钥格式化|简单|37%|11598|
|[0481](https://leetcode-cn.com/problems/magical-string/)|神奇字符串|中等|47%|2236|
|[0480](https://leetcode-cn.com/problems/sliding-window-median/)|滑动窗口中位数|困难|33%|5115|
|[0479](https://leetcode-cn.com/problems/largest-palindrome-product/)|最大回文数乘积|困难|29%|3626|
|[0477](https://leetcode-cn.com/problems/total-hamming-distance/)|[汉明距离总和](./src/0477-total-hamming-distance/)|中等|45%|6359|
|[0476](https://leetcode-cn.com/problems/number-complement/)|数字的补数|简单|67%|20181|
|[0475](https://leetcode-cn.com/problems/heaters/)|供暖器|简单|28%|16774|
|[0474](https://leetcode-cn.com/problems/ones-and-zeroes/)|一和零|中等|46%|7900|
|[0473](https://leetcode-cn.com/problems/matchsticks-to-square/)|火柴拼正方形|中等|35%|8131|
|[0472](https://leetcode-cn.com/problems/concatenated-words/)|连接词|困难|39%|2446|
|[0471](https://leetcode-cn.com/problems/encode-string-with-shortest-length/)|编码最短长度的字符串|困难|38%|285|
|[0469](https://leetcode-cn.com/problems/convex-polygon/)|凸多边形|中等|38%|380|
|[0468](https://leetcode-cn.com/problems/validate-ip-address/)|验证IP地址|中等|20%|11256|
|[0467](https://leetcode-cn.com/problems/unique-substrings-in-wraparound-string/)|环绕字符串中唯一的子字符串|中等|36%|3325|
|[0466](https://leetcode-cn.com/problems/count-the-repetitions/)|统计重复个数|困难|27%|2028|
|[0465](https://leetcode-cn.com/problems/optimal-account-balancing/)|最优账单平衡|困难|46%|176|
|[0464](https://leetcode-cn.com/problems/can-i-win/)|我能赢吗|中等|32%|5290|
|[0463](https://leetcode-cn.com/problems/island-perimeter/)|岛屿的周长|简单|63%|14883|
|[0462](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/)|最少移动次数使数组元素相等 II|中等|54%|4952|
|[0461](https://leetcode-cn.com/problems/hamming-distance/)|汉明距离|简单|73%|42372|
|[0460](https://leetcode-cn.com/problems/lfu-cache/)|LFU缓存|困难|33%|4708|
|[0459](https://leetcode-cn.com/problems/repeated-substring-pattern/)|重复的子字符串|简单|43%|22633|
|[0458](https://leetcode-cn.com/problems/poor-pigs/)|可怜的小猪|困难|52%|4492|
|[0457](https://leetcode-cn.com/problems/circular-array-loop/)|环形数组循环|中等|30%|4377|
|[0456](https://leetcode-cn.com/problems/132-pattern/)|132模式|中等|24%|16499|
|[0455](https://leetcode-cn.com/problems/assign-cookies/)|分发饼干|简单|52%|33422|
|[0454](https://leetcode-cn.com/problems/4sum-ii/)|四数相加 II|中等|53%|20039|
|[0453](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/)|最小移动次数使数组元素相等|简单|52%|11443|
|[0452](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/)|用最少数量的箭引爆气球|中等|47%|11130|
|[0451](https://leetcode-cn.com/problems/sort-characters-by-frequency/)|根据字符出现频率排序|中等|60%|14818|
|[0450](https://leetcode-cn.com/problems/delete-node-in-a-bst/)|删除二叉搜索树中的节点|中等|36%|15193|
|[0449](https://leetcode-cn.com/problems/serialize-and-deserialize-bst/)|序列化和反序列化二叉搜索树|中等|49%|4289|
|[0448](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/)|[找到所有数组中消失的数字](./src/0448-find-all-numbers-disappeared-in-an-array/)|简单|53%|31836|
|[0447](https://leetcode-cn.com/problems/number-of-boomerangs/)|回旋镖的数量|简单|54%|12264|
|[0446](https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence/)|等差数列划分 II - 子序列|困难|29%|2412|
|[0445](https://leetcode-cn.com/problems/add-two-numbers-ii/)|两数相加 II|中等|51%|16015|
|[0444](https://leetcode-cn.com/problems/sequence-reconstruction/)|序列重建|中等|22%|664|
|[0443](https://leetcode-cn.com/problems/string-compression/)|压缩字符串|简单|37%|20986|
|[0442](https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/)|数组中重复的数据|中等|62%|13361|
|[0441](https://leetcode-cn.com/problems/arranging-coins/)|排列硬币|简单|39%|29991|
|[0440](https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/)|字典序的第K小数字|困难|29%|4451|
|[0439](https://leetcode-cn.com/problems/ternary-expression-parser/)|三元表达式解析器|中等|53%|330|
|[0438](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)|找到字符串中所有字母异位词|中等|38%|29673|
|[0437](https://leetcode-cn.com/problems/path-sum-iii/)|路径总和 III|简单|52%|25150|
|[0436](https://leetcode-cn.com/problems/find-right-interval/)|寻找右区间|中等|41%|3390|
|[0435](https://leetcode-cn.com/problems/non-overlapping-intervals/)|无重叠区间|中等|43%|11828|
|[0434](https://leetcode-cn.com/problems/number-of-segments-in-a-string/)|字符串中的单词数|简单|32%|27499|
|[0433](https://leetcode-cn.com/problems/minimum-genetic-mutation/)|最小基因变化|中等|44%|2176|
|[0432](https://leetcode-cn.com/problems/all-oone-data-structure/)|全 O(1) 的数据结构|困难|35%|4686|
|[0425](https://leetcode-cn.com/problems/word-squares/)|单词方块|困难|48%|203|
|[0424](https://leetcode-cn.com/problems/longest-repeating-character-replacement/)|替换后的最长重复字符|中等|43%|5715|
|[0423](https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/)|从英文中重建数字|中等|50%|3443|
|[0422](https://leetcode-cn.com/problems/valid-word-square/)|有效的单词方块|简单|39%|656|
|[0421](https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/)|[数组中两个数的最大异或值](./src/0421-maximum-xor-of-two-numbers-in-an-array/)|中等|57%|3455|
|[0420](https://leetcode-cn.com/problems/strong-password-checker/)|强密码检验器|困难|12%|4342|
|[0419](https://leetcode-cn.com/problems/battleships-in-a-board/)|甲板上的战舰|中等|68%|2738|
|[0418](https://leetcode-cn.com/problems/sentence-screen-fitting/)|屏幕可显示句子的数量|中等|28%|731|
|[0417](https://leetcode-cn.com/problems/pacific-atlantic-water-flow/)|太平洋大西洋水流问题|中等|37%|6462|
|[0416](https://leetcode-cn.com/problems/partition-equal-subset-sum/)|分割等和子集|中等|43%|25897|
|[0415](https://leetcode-cn.com/problems/add-strings/)|字符串相加|简单|47%|33055|
|[0414](https://leetcode-cn.com/problems/third-maximum-number/)|第三大的数|简单|33%|40321|
|[0413](https://leetcode-cn.com/problems/arithmetic-slices/)|等差数列划分|中等|57%|11705|
|[0412](https://leetcode-cn.com/problems/fizz-buzz/)|Fizz Buzz|简单|61%|34921|
|[0411](https://leetcode-cn.com/problems/minimum-unique-word-abbreviation/)|最短特异单词缩写|困难|52%|188|
|[0410](https://leetcode-cn.com/problems/split-array-largest-sum/)|分割数组的最大值|困难|38%|7517|
|[0409](https://leetcode-cn.com/problems/longest-palindrome/)|最长回文串|简单|50%|21544|
|[0408](https://leetcode-cn.com/problems/valid-word-abbreviation/)|有效单词缩写|简单|31%|1098|
|[0407](https://leetcode-cn.com/problems/trapping-rain-water-ii/)|接雨水 II|困难|32%|3593|
|[0406](https://leetcode-cn.com/problems/queue-reconstruction-by-height/)|根据身高重建队列|中等|62%|13521|
|[0405](https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/)|数字转换为十六进制数|简单|47%|12306|
|[0404](https://leetcode-cn.com/problems/sum-of-left-leaves/)|左叶子之和|简单|52%|22285|
|[0403](https://leetcode-cn.com/problems/frog-jump/)|青蛙过河|困难|29%|7587|
|[0402](https://leetcode-cn.com/problems/remove-k-digits/)|移掉K位数字|中等|26%|24946|
|[0401](https://leetcode-cn.com/problems/binary-watch/)|二进制手表|简单|49%|14569|
|[0400](https://leetcode-cn.com/problems/nth-digit/)|第N个数字|中等|33%|16156|
|[0399](https://leetcode-cn.com/problems/evaluate-division/)|除法求值|中等|51%|4677|
|[0398](https://leetcode-cn.com/problems/random-pick-index/)|随机数索引|中等|56%|3930|
|[0397](https://leetcode-cn.com/problems/integer-replacement/)|整数替换|中等|32%|10418|
|[0396](https://leetcode-cn.com/problems/rotate-function/)|旋转函数|中等|37%|4177|
|[0395](https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/)|至少有K个重复字符的最长子串|中等|40%|11174|
|[0394](https://leetcode-cn.com/problems/decode-string/)|字符串解码|中等|46%|20669|
|[0393](https://leetcode-cn.com/problems/utf-8-validation/)|UTF-8 编码验证|中等|36%|7006|
|[0392](https://leetcode-cn.com/problems/is-subsequence/)|判断子序列|简单|47%|29435|
|[0391](https://leetcode-cn.com/problems/perfect-rectangle/)|完美矩形|困难|23%|2698|
|[0390](https://leetcode-cn.com/problems/elimination-game/)|消除游戏|中等|40%|3701|
|[0389](https://leetcode-cn.com/problems/find-the-difference/)|找不同|简单|59%|26605|
|[0388](https://leetcode-cn.com/problems/longest-absolute-file-path/)|文件的最长绝对路径|中等|43%|1801|
|[0387](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)|字符串中的第一个唯一字符|简单|41%|112964|
|[0386](https://leetcode-cn.com/problems/lexicographical-numbers/)|字典序排数|中等|64%|5559|
|[0385](https://leetcode-cn.com/problems/mini-parser/)|迷你语法分析器|中等|39%|2753|
|[0384](https://leetcode-cn.com/problems/shuffle-an-array/)|打乱数组|中等|47%|23559|
|[0383](https://leetcode-cn.com/problems/ransom-note/)|赎金信|简单|50%|23140|
|[0382](https://leetcode-cn.com/problems/linked-list-random-node/)|链表随机节点|中等|54%|4183|
|[0381](https://leetcode-cn.com/problems/insert-delete-getrandom-o1-duplicates-allowed/)|O(1) 时间插入、删除和获取随机元素 - 允许重复|困难|36%|3394|
|[0380](https://leetcode-cn.com/problems/insert-delete-getrandom-o1/)|常数时间插入、删除和获取随机元素|中等|46%|12678|
|[0379](https://leetcode-cn.com/problems/design-phone-directory/)|电话目录管理系统|中等|62%|329|
|[0378](https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/)|有序矩阵中第K小的元素|中等|56%|16811|
|[0377](https://leetcode-cn.com/problems/combination-sum-iv/)|组合总和 Ⅳ|中等|39%|12038|
|[0376](https://leetcode-cn.com/problems/wiggle-subsequence/)|摆动序列|中等|39%|17901|
|[0375](https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/)|猜数字大小 II|中等|35%|7601|
|[0374](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)|[猜数字大小](./src/0374-guess-number-higher-or-lower/)|简单|41%|29370|
|[0373](https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/)|查找和最小的K对数字|中等|37%|8344|
|[0372](https://leetcode-cn.com/problems/super-pow/)|超级次方|中等|36%|5520|
|[0371](https://leetcode-cn.com/problems/sum-of-two-integers/)|两整数之和|简单|52%|32388|
|[0370](https://leetcode-cn.com/problems/range-addition/)|区间加法|中等|61%|381|
|[0369](https://leetcode-cn.com/problems/plus-one-linked-list/)|给单链表加一|中等|55%|661|
|[0368](https://leetcode-cn.com/problems/largest-divisible-subset/)|最大整除子集|中等|35%|8270|
|[0367](https://leetcode-cn.com/problems/valid-perfect-square/)|有效的完全平方数|简单|42%|38986|
|[0366](https://leetcode-cn.com/problems/find-leaves-of-binary-tree/)|寻找完全二叉树的叶子节点|中等|69%|399|
|[0365](https://leetcode-cn.com/problems/water-and-jug-problem/)|水壶问题|中等|28%|9961|
|[0364](https://leetcode-cn.com/problems/nested-list-weight-sum-ii/)|加权嵌套序列和 II|中等|71%|301|
|[0363](https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/)|矩形区域不超过 K 的最大数值和|困难|31%|3484|
|[0362](https://leetcode-cn.com/problems/design-hit-counter/)|敲击计数器|中等|70%|262|
|[0361](https://leetcode-cn.com/problems/bomb-enemy/)|轰炸敌人|中等|47%|861|
|[0360](https://leetcode-cn.com/problems/sort-transformed-array/)|有序转化数组|中等|51%|344|
|[0359](https://leetcode-cn.com/problems/logger-rate-limiter/)|日志速率限制器|简单|67%|604|
|[0358](https://leetcode-cn.com/problems/rearrange-string-k-distance-apart/)|K 距离间隔重排字符串|困难|31%|420|
|[0357](https://leetcode-cn.com/problems/count-numbers-with-unique-digits/)|计算各个位数不同的数字个数|中等|47%|8995|
|[0356](https://leetcode-cn.com/problems/line-reflection/)|直线镜像|中等|27%|542|
|[0355](https://leetcode-cn.com/problems/design-twitter/)|设计推特|中等|35%|4253|
|[0354](https://leetcode-cn.com/problems/russian-doll-envelopes/)|俄罗斯套娃信封问题|困难|32%|13738|
|[0353](https://leetcode-cn.com/problems/design-snake-game/)|贪吃蛇|中等|35%|297|
|[0352](https://leetcode-cn.com/problems/data-stream-as-disjoint-intervals/)|将数据流变为多个不相交区间|困难|49%|1693|
|[0351](https://leetcode-cn.com/problems/android-unlock-patterns/)|安卓系统手势解锁|中等|57%|433|
|[0350](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)|两个数组的交集 II|简单|44%|115822|
|[0349](https://leetcode-cn.com/problems/intersection-of-two-arrays/)|两个数组的交集|简单|66%|55770|
|[0348](https://leetcode-cn.com/problems/design-tic-tac-toe/)|判定井字棋胜负|中等|55%|487|
|[0347](https://leetcode-cn.com/problems/top-k-frequent-elements/)|前 K 个高频元素|中等|58%|42479|
|[0346](https://leetcode-cn.com/problems/moving-average-from-data-stream/)|数据流中的移动平均值|简单|70%|1571|
|[0345](https://leetcode-cn.com/problems/reverse-vowels-of-a-string/)|反转字符串中的元音字母|简单|48%|35228|
|[0344](https://leetcode-cn.com/problems/reverse-string/)|反转字符串|简单|67%|123509|
|[0343](https://leetcode-cn.com/problems/integer-break/)|整数拆分|中等|52%|22018|
|[0342](https://leetcode-cn.com/problems/power-of-four/)|4的幂|简单|46%|29955|
|[0341](https://leetcode-cn.com/problems/flatten-nested-list-iterator/)|扁平化嵌套列表迭代器|中等|60%|6169|
|[0340](https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters/)|至多包含 K 个不同字符的最长子串|困难|43%|1069|
|[0339](https://leetcode-cn.com/problems/nested-list-weight-sum/)|嵌套列表权重和|简单|75%|570|
|[0338](https://leetcode-cn.com/problems/counting-bits/)|比特位计数|中等|73%|24728|
|[0337](https://leetcode-cn.com/problems/house-robber-iii/)|打家劫舍 III|中等|54%|15982|
|[0336](https://leetcode-cn.com/problems/palindrome-pairs/)|回文对|困难|30%|5412|
|[0335](https://leetcode-cn.com/problems/self-crossing/)|路径交叉|困难|29%|2031|
|[0334](https://leetcode-cn.com/problems/increasing-triplet-subsequence/)|递增的三元子序列|中等|35%|28396|
|[0333](https://leetcode-cn.com/problems/largest-bst-subtree/)|最大 BST 子树|中等|41%|1327|
|[0332](https://leetcode-cn.com/problems/reconstruct-itinerary/)|重新安排行程|中等|32%|6929|
|[0331](https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/)|验证二叉树的前序序列化|中等|41%|6273|
|[0330](https://leetcode-cn.com/problems/patching-array/)|按要求补齐数组|困难|40%|2841|
|[0329](https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/)|矩阵中的最长递增路径|困难|39%|14438|
|[0328](https://leetcode-cn.com/problems/odd-even-linked-list/)|奇偶链表|中等|59%|31241|
|[0327](https://leetcode-cn.com/problems/count-of-range-sum/)|区间和的个数|困难|31%|4118|
|[0326](https://leetcode-cn.com/problems/power-of-three/)|3的幂|简单|45%|61373|
|[0325](https://leetcode-cn.com/problems/maximum-size-subarray-sum-equals-k/)|和等于 k 的最长子数组长度|中等|41%|1437|
|[0324](https://leetcode-cn.com/problems/wiggle-sort-ii/)|摆动排序 II|中等|34%|15995|
|[0323](https://leetcode-cn.com/problems/number-of-connected-components-in-an-undirected-graph/)|无向图中连通分量的数目|中等|51%|1124|
|[0322](https://leetcode-cn.com/problems/coin-change/)|零钱兑换|中等|36%|74553|
|[0321](https://leetcode-cn.com/problems/create-maximum-number/)|拼接最大数|困难|25%|4507|
|[0320](https://leetcode-cn.com/problems/generalized-abbreviation/)|列举单词的全部缩写|中等|60%|432|
|[0319](https://leetcode-cn.com/problems/bulb-switcher/)|灯泡开关|中等|43%|12597|
|[0318](https://leetcode-cn.com/problems/maximum-product-of-word-lengths/)|最大单词长度乘积|中等|58%|5665|
|[0317](https://leetcode-cn.com/problems/shortest-distance-from-all-buildings/)|离建筑物最近的距离|困难|35%|492|
|[0316](https://leetcode-cn.com/problems/remove-duplicate-letters/)|去除重复字母|困难|35%|7980|
|[0315](https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/)|计算右侧小于当前元素的个数|困难|37%|15463|
|[0314](https://leetcode-cn.com/problems/binary-tree-vertical-order-traversal/)|二叉树的垂直遍历|中等|54%|465|
|[0313](https://leetcode-cn.com/problems/super-ugly-number/)|超级丑数|中等|58%|5751|
|[0312](https://leetcode-cn.com/problems/burst-balloons/)|戳气球|困难|56%|8186|
|[0311](https://leetcode-cn.com/problems/sparse-matrix-multiplication/)|稀疏矩阵的乘法|中等|72%|363|
|[0310](https://leetcode-cn.com/problems/minimum-height-trees/)|最小高度树|中等|33%|8242|
|[0309](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)|最佳买卖股票时机含冷冻期|中等|49%|17612|
|[0308](https://leetcode-cn.com/problems/range-sum-query-2d-mutable/)|二维区域和检索 - 可变|困难|56%|285|
|[0307](https://leetcode-cn.com/problems/range-sum-query-mutable/)|区域和检索 - 数组可修改|中等|51%|9083|
|[0306](https://leetcode-cn.com/problems/additive-number/)|累加数|中等|31%|8176|
|[0305](https://leetcode-cn.com/problems/number-of-islands-ii/)|岛屿数量 II|困难|28%|873|
|[0304](https://leetcode-cn.com/problems/range-sum-query-2d-immutable/)|二维区域和检索 - 矩阵不可变|中等|41%|8411|
|[0303](https://leetcode-cn.com/problems/range-sum-query-immutable/)|[区域和检索 - 数组不可变](./src/0303-range-sum-query-immutable/)|简单|57%|35840|
|[0302](https://leetcode-cn.com/problems/smallest-rectangle-enclosing-black-pixels/)|包含全部黑色像素的最小矩形|困难|65%|282|
|[0301](https://leetcode-cn.com/problems/remove-invalid-parentheses/)|删除无效的括号|困难|43%|7389|
|[0300](https://leetcode-cn.com/problems/longest-increasing-subsequence/)|最长上升子序列|中等|43%|77281|
|[0299](https://leetcode-cn.com/problems/bulls-and-cows/)|猜数字游戏|简单|46%|9858|
|[0298](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence/)|二叉树最长连续序列|中等|50%|722|
|[0297](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)|二叉树的序列化与反序列化|困难|42%|21882|
|[0296](https://leetcode-cn.com/problems/best-meeting-point/)|最佳的碰头地点|困难|58%|276|
|[0295](https://leetcode-cn.com/problems/find-median-from-data-stream/)|数据流的中位数|困难|39%|16902|
|[0294](https://leetcode-cn.com/problems/flip-game-ii/)|翻转游戏 II|中等|48%|668|
|[0293](https://leetcode-cn.com/problems/flip-game/)|翻转游戏|简单|69%|966|
|[0292](https://leetcode-cn.com/problems/nim-game/)|[Nim 游戏](./src/0292-nim-game/)|简单|69%|41513|
|[0291](https://leetcode-cn.com/problems/word-pattern-ii/)|单词规律 II|困难|52%|263|
|[0290](https://leetcode-cn.com/problems/word-pattern/)|单词规律|简单|40%|30968|
|[0289](https://leetcode-cn.com/problems/game-of-life/)|生命游戏|中等|66%|10481|
|[0288](https://leetcode-cn.com/problems/unique-word-abbreviation/)|单词的唯一缩写|中等|28%|915|
|[0287](https://leetcode-cn.com/problems/find-the-duplicate-number/)|寻找重复数|中等|61%|38860|
|[0286](https://leetcode-cn.com/problems/walls-and-gates/)|墙与门|中等|43%|2270|
|[0285](https://leetcode-cn.com/problems/inorder-successor-in-bst/)|二叉搜索树中的顺序后继|中等|61%|1107|
|[0284](https://leetcode-cn.com/problems/peeking-iterator/)|顶端迭代器|中等|70%|2646|
|[0283](https://leetcode-cn.com/problems/move-zeroes/)|移动零|简单|57%|135850|
|[0282](https://leetcode-cn.com/problems/expression-add-operators/)|给表达式添加运算符|困难|29%|3621|
|[0281](https://leetcode-cn.com/problems/zigzag-iterator/)|锯齿迭代器|中等|71%|288|
|[0280](https://leetcode-cn.com/problems/wiggle-sort/)|摆动排序|中等|68%|806|
|[0279](https://leetcode-cn.com/problems/perfect-squares/)|完全平方数|中等|51%|46336|
|[0278](https://leetcode-cn.com/problems/first-bad-version/)|[第一个错误的版本](./src/0278-first-bad-version/)|简单|34%|78012|
|[0277](https://leetcode-cn.com/problems/find-the-celebrity/)|搜寻名人|中等|52%|1044|
|[0276](https://leetcode-cn.com/problems/paint-fence/)|栅栏涂色|简单|41%|1647|
|[0275](https://leetcode-cn.com/problems/h-index-ii/)|H指数 II|中等|38%|10252|
|[0274](https://leetcode-cn.com/problems/h-index/)|H指数|中等|36%|14200|
|[0273](https://leetcode-cn.com/problems/integer-to-english-words/)|整数转换英文表示|困难|24%|8311|
|[0272](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii/)|最接近的二叉搜索树值 II|困难|60%|478|
|[0271](https://leetcode-cn.com/problems/encode-and-decode-strings/)|字符串的编码与解码|中等|46%|421|
|[0270](https://leetcode-cn.com/problems/closest-binary-search-tree-value/)|最接近的二叉搜索树值|简单|49%|1631|
|[0269](https://leetcode-cn.com/problems/alien-dictionary/)|火星词典|困难|39%|1176|
|[0268](https://leetcode-cn.com/problems/missing-number/)|缺失数字|简单|52%|67881|
|[0267](https://leetcode-cn.com/problems/palindrome-permutation-ii/)|回文排列 II|中等|37%|637|
|[0266](https://leetcode-cn.com/problems/palindrome-permutation/)|回文排列|简单|61%|1056|
|[0265](https://leetcode-cn.com/problems/paint-house-ii/)|粉刷房子 II|困难|52%|898|
|[0264](https://leetcode-cn.com/problems/ugly-number-ii/)|丑数 II|中等|47%|23123|
|[0263](https://leetcode-cn.com/problems/ugly-number/)|丑数|简单|47%|36706|
|[0262](https://leetcode-cn.com/problems/trips-and-users/)|[行程和用户](./src/0262-trips-and-users/)|困难|43%|18299|
|[0261](https://leetcode-cn.com/problems/graph-valid-tree/)|以图判树|中等|41%|1418|
|[0260](https://leetcode-cn.com/problems/single-number-iii/)|只出现一次的数字 III|中等|67%|14851|
|[0259](https://leetcode-cn.com/problems/3sum-smaller/)|较小的三数之和|中等|53%|1044|
|[0258](https://leetcode-cn.com/problems/add-digits/)|各位相加|简单|64%|35923|
|[0257](https://leetcode-cn.com/problems/binary-tree-paths/)|二叉树的所有路径|简单|60%|28574|
|[0256](https://leetcode-cn.com/problems/paint-house/)|粉刷房子|简单|55%|1758|
|[0255](https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree/)|验证前序遍历序列二叉搜索树|中等|42%|1093|
|[0254](https://leetcode-cn.com/problems/factor-combinations/)|因子的组合|中等|51%|661|
|[0253](https://leetcode-cn.com/problems/meeting-rooms-ii/)|会议室 II|中等|41%|3000|
|[0252](https://leetcode-cn.com/problems/meeting-rooms/)|会议室|简单|51%|1621|
|[0251](https://leetcode-cn.com/problems/flatten-2d-vector/)|展开二维向量|中等|48%|706|
|[0250](https://leetcode-cn.com/problems/count-univalue-subtrees/)|统计同值子树|中等|59%|542|
|[0249](https://leetcode-cn.com/problems/group-shifted-strings/)|移位字符串分组|中等|56%|815|
|[0248](https://leetcode-cn.com/problems/strobogrammatic-number-iii/)|中心对称数 III|困难|42%|497|
|[0247](https://leetcode-cn.com/problems/strobogrammatic-number-ii/)|中心对称数 II|中等|45%|976|
|[0246](https://leetcode-cn.com/problems/strobogrammatic-number/)|中心对称数|简单|43%|1300|
|[0245](https://leetcode-cn.com/problems/shortest-word-distance-iii/)|最短单词距离 III|中等|62%|678|
|[0244](https://leetcode-cn.com/problems/shortest-word-distance-ii/)|最短单词距离 II|中等|54%|863|
|[0243](https://leetcode-cn.com/problems/shortest-word-distance/)|最短单词距离|简单|65%|1473|
|[0242](https://leetcode-cn.com/problems/valid-anagram/)|有效的字母异位词|简单|56%|93765|
|[0241](https://leetcode-cn.com/problems/different-ways-to-add-parentheses/)|为运算表达式设计优先级|中等|69%|6356|
|[0240](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)|搜索二维矩阵 II|中等|37%|77519|
|[0239](https://leetcode-cn.com/problems/sliding-window-maximum/)|滑动窗口最大值|困难|42%|42183|
|[0238](https://leetcode-cn.com/problems/product-of-array-except-self/)|除自身以外数组的乘积|中等|63%|30817|
|[0237](https://leetcode-cn.com/problems/delete-node-in-a-linked-list/)|删除链表中的节点|简单|78%|74547|
|[0236](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)|二叉树的最近公共祖先|中等|58%|51724|
|[0235](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)|二叉搜索树的最近公共祖先|简单|60%|45077|
|[0234](https://leetcode-cn.com/problems/palindrome-linked-list/)|回文链表|简单|38%|113386|
|[0233](https://leetcode-cn.com/problems/number-of-digit-one/)|数字 1 的个数|困难|30%|11128|
|[0232](https://leetcode-cn.com/problems/implement-queue-using-stacks/)|用栈实现队列|简单|61%|36411|
|[0231](https://leetcode-cn.com/problems/power-of-two/)|2的幂|简单|46%|72708|
|[0230](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)|二叉搜索树中第K小的元素|中等|66%|37079|
|[0229](https://leetcode-cn.com/problems/majority-element-ii/)|求众数 II|中等|41%|18086|
|[0228](https://leetcode-cn.com/problems/summary-ranges/)|汇总区间|中等|48%|9461|
|[0227](https://leetcode-cn.com/problems/basic-calculator-ii/)|基本计算器 II|中等|33%|20558|
|[0226](https://leetcode-cn.com/problems/invert-binary-tree/)|翻转二叉树|简单|72%|49391|
|[0225](https://leetcode-cn.com/problems/implement-stack-using-queues/)|用队列实现栈|简单|60%|32691|
|[0224](https://leetcode-cn.com/problems/basic-calculator/)|基本计算器|困难|35%|17806|
|[0223](https://leetcode-cn.com/problems/rectangle-area/)|矩形面积|中等|41%|10188|
|[0222](https://leetcode-cn.com/problems/count-complete-tree-nodes/)|完全二叉树的节点个数|中等|61%|14817|
|[0221](https://leetcode-cn.com/problems/maximal-square/)|最大正方形|中等|38%|35386|
|[0220](https://leetcode-cn.com/problems/contains-duplicate-iii/)|存在重复元素 III|中等|24%|33563|
|[0219](https://leetcode-cn.com/problems/contains-duplicate-ii/)|存在重复元素 II|简单|36%|66446|
|[0218](https://leetcode-cn.com/problems/the-skyline-problem/)|天际线问题|困难|39%|7460|
|[0217](https://leetcode-cn.com/problems/contains-duplicate/)|存在重复元素|简单|50%|170242|
|[0216](https://leetcode-cn.com/problems/combination-sum-iii/)|组合总和 III|中等|68%|13607|
|[0215](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)|数组中的第K个最大元素|中等|58%|91199|
|[0214](https://leetcode-cn.com/problems/shortest-palindrome/)|最短回文串|困难|30%|12754|
|[0213](https://leetcode-cn.com/problems/house-robber-ii/)|打家劫舍 II|中等|34%|37224|
|[0212](https://leetcode-cn.com/problems/word-search-ii/)|单词搜索 II|困难|36%|14755|
|[0211](https://leetcode-cn.com/problems/add-and-search-word-data-structure-design/)|添加与搜索单词 - 数据结构设计|中等|39%|11902|
|[0210](https://leetcode-cn.com/problems/course-schedule-ii/)|课程表 II|中等|47%|14999|
|[0209](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)|长度最小的子数组|中等|40%|45660|
|[0208](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)|实现 Trie (前缀树)|中等|62%|24171|
|[0207](https://leetcode-cn.com/problems/course-schedule/)|课程表|中等|48%|28156|
|[0206](https://leetcode-cn.com/problems/reverse-linked-list/)|反转链表|简单|65%|172180|
|[0205](https://leetcode-cn.com/problems/isomorphic-strings/)|同构字符串|简单|46%|42089|
|[0204](https://leetcode-cn.com/problems/count-primes/)|计数质数|简单|30%|100776|
|[0203](https://leetcode-cn.com/problems/remove-linked-list-elements/)|移除链表元素|简单|42%|101837|
|[0202](https://leetcode-cn.com/problems/happy-number/)|快乐数|简单|55%|58025|
|[0201](https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/)|数字范围按位与|中等|42%|12421|
|[0200](https://leetcode-cn.com/problems/number-of-islands/)|岛屿数量|中等|45%|74924|
|[0199](https://leetcode-cn.com/problems/binary-tree-right-side-view/)|二叉树的右视图|中等|61%|18174|
|[0198](https://leetcode-cn.com/problems/house-robber/)|[打家劫舍](./src/0198-house-robber/)|简单|41%|123501|
|[0197](https://leetcode-cn.com/problems/rising-temperature/)|[上升的温度](./src/0197-rising-temperature/)|简单|47%|44869|
|[0196](https://leetcode-cn.com/problems/delete-duplicate-emails/)|[删除重复的电子邮箱](./src/0196-delete-duplicate-emails/)|简单|57%|40472|
|[0195](https://leetcode-cn.com/problems/tenth-line/)|[第十行](./src/0195-tenth-line/)|简单|42%|24963|
|[0194](https://leetcode-cn.com/problems/transpose-file/)|[转置文件](./src/0194-transpose-file/)|中等|32%|9062|
|[0193](https://leetcode-cn.com/problems/valid-phone-numbers/)|[有效电话号码](./src/0193-valid-phone-numbers/)|简单|26%|24960|
|[0192](https://leetcode-cn.com/problems/word-frequency/)|[统计词频](./src/0192-word-frequency/)|中等|32%|18407|
|[0191](https://leetcode-cn.com/problems/number-of-1-bits/)|位1的个数|简单|60%|56099|
|[0190](https://leetcode-cn.com/problems/reverse-bits/)|颠倒二进制位|简单|49%|44238|
|[0189](https://leetcode-cn.com/problems/rotate-array/)|旋转数组|简单|39%|189992|
|[0188](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)|买卖股票的最佳时机 IV|困难|27%|24975|
|[0187](https://leetcode-cn.com/problems/repeated-dna-sequences/)|重复的DNA序列|中等|42%|17145|
|[0186](https://leetcode-cn.com/problems/reverse-words-in-a-string-ii/)|翻转字符串里的单词 II|中等|69%|926|
|[0185](https://leetcode-cn.com/problems/department-top-three-salaries/)|[部门工资前三高的所有员工](./src/0185-department-top-three-salaries/)|困难|39%|36325|
|[0184](https://leetcode-cn.com/problems/department-highest-salary/)|[部门工资最高的员工](./src/0184-department-highest-salary/)|中等|41%|51492|
|[0183](https://leetcode-cn.com/problems/customers-who-never-order/)|[从不订购的客户](./src/0183-customers-who-never-order/)|简单|63%|55845|
|[0182](https://leetcode-cn.com/problems/duplicate-emails/)|[查找重复的电子邮箱](./src/0182-duplicate-emails/)|简单|76%|61306|
|[0181](https://leetcode-cn.com/problems/employees-earning-more-than-their-managers/)|[超过经理收入的员工](./src/0181-employees-earning-more-than-their-managers/)|简单|66%|57028|
|[0180](https://leetcode-cn.com/problems/consecutive-numbers/)|[连续出现的数字](./src/0180-consecutive-numbers/)|中等|47%|39035|
|[0179](https://leetcode-cn.com/problems/largest-number/)|最大数|中等|34%|40554|
|[0178](https://leetcode-cn.com/problems/rank-scores/)|[分数排名](./src/0178-rank-scores/)|中等|58%|43139|
|[0177](https://leetcode-cn.com/problems/nth-highest-salary/)|[第N高的薪水](./src/0177-nth-highest-salary/)|中等|42%|59034|
|[0176](https://leetcode-cn.com/problems/second-highest-salary/)|[第二高的薪水](./src/0176-second-highest-salary/)|简单|33%|169350|
|[0175](https://leetcode-cn.com/problems/combine-two-tables/)|[组合两个表](./src/0175-combine-two-tables/)|简单|70%|108138|
|[0174](https://leetcode-cn.com/problems/dungeon-game/)|地下城游戏|困难|38%|12394|
|[0173](https://leetcode-cn.com/problems/binary-search-tree-iterator/)|二叉搜索树迭代器|中等|68%|12621|
|[0172](https://leetcode-cn.com/problems/factorial-trailing-zeroes/)|阶乘后的零|简单|39%|51940|
|[0171](https://leetcode-cn.com/problems/excel-sheet-column-number/)|Excel表列序号|简单|65%|35690|
|[0170](https://leetcode-cn.com/problems/two-sum-iii-data-structure-design/)|两数之和 III - 数据结构设计|简单|39%|1686|
|[0169](https://leetcode-cn.com/problems/majority-element/)|求众数|简单|60%|132744|
|[0168](https://leetcode-cn.com/problems/excel-sheet-column-title/)|Excel表列名称|简单|35%|40302|
|[0167](https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/)|两数之和 II - 输入有序数组|简单|50%|92368|
|[0166](https://leetcode-cn.com/problems/fraction-to-recurring-decimal/)|分数到小数|中等|25%|22363|
|[0165](https://leetcode-cn.com/problems/compare-version-numbers/)|比较版本号|中等|35%|21059|
|[0164](https://leetcode-cn.com/problems/maximum-gap/)|最大间距|困难|52%|12718|
|[0163](https://leetcode-cn.com/problems/missing-ranges/)|缺失的区间|中等|23%|3753|
|[0162](https://leetcode-cn.com/problems/find-peak-element/)|寻找峰值|中等|42%|46881|
|[0161](https://leetcode-cn.com/problems/one-edit-distance/)|相隔为 1 的编辑距离|中等|32%|2037|
|[0160](https://leetcode-cn.com/problems/intersection-of-two-linked-lists/)|相交链表|简单|49%|104454|
|[0159](https://leetcode-cn.com/problems/longest-substring-with-at-most-two-distinct-characters/)|至多包含两个不同字符的最长子串|困难|53%|1347|
|[0158](https://leetcode-cn.com/problems/read-n-characters-given-read4-ii-call-multiple-times/)|用 Read4 读取 N 个字符 II|困难|51%|450|
|[0157](https://leetcode-cn.com/problems/read-n-characters-given-read4/)|用 Read4 读取 N 个字符|简单|51%|929|
|[0156](https://leetcode-cn.com/problems/binary-tree-upside-down/)|上下翻转二叉树|中等|73%|711|
|[0155](https://leetcode-cn.com/problems/min-stack/)|最小栈|简单|50%|101573|
|[0154](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)|寻找旋转排序数组中的最小值 II|困难|45%|18611|
|[0153](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)|寻找旋转排序数组中的最小值|中等|49%|40790|
|[0152](https://leetcode-cn.com/problems/maximum-product-subarray/)|乘积最大子序列|中等|35%|65151|
|[0151](https://leetcode-cn.com/problems/reverse-words-in-a-string/)|翻转字符串里的单词|中等|32%|64939|
|[0150](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)|逆波兰表达式求值|中等|47%|37907|
|[0149](https://leetcode-cn.com/problems/max-points-on-a-line/)|直线上最多的点数|困难|18%|29321|
|[0148](https://leetcode-cn.com/problems/sort-list/)|排序链表|中等|62%|44672|
|[0147](https://leetcode-cn.com/problems/insertion-sort-list/)|对链表进行插入排序|中等|60%|21795|
|[0146](https://leetcode-cn.com/problems/lru-cache/)|[LRU缓存机制](./src/0146-lru-cache/)|中等|44%|48961|
|[0145](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)|二叉树的后序遍历|困难|68%|51732|
|[0144](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)|二叉树的前序遍历|中等|62%|74958|
|[0143](https://leetcode-cn.com/problems/reorder-list/)|重排链表|中等|52%|19962|
|[0142](https://leetcode-cn.com/problems/linked-list-cycle-ii/)|环形链表 II|中等|45%|75767|
|[0141](https://leetcode-cn.com/problems/linked-list-cycle/)|环形链表|简单|43%|173522|
|[0140](https://leetcode-cn.com/problems/word-break-ii/)|单词拆分 II|困难|36%|19600|
|[0139](https://leetcode-cn.com/problems/word-break/)|单词拆分|中等|42%|52628|
|[0138](https://leetcode-cn.com/problems/copy-list-with-random-pointer/)|复制带随机指针的链表|中等|38%|34127|
|[0137](https://leetcode-cn.com/problems/single-number-ii/)|只出现一次的数字 II|中等|64%|22739|
|[0136](https://leetcode-cn.com/problems/single-number/)|只出现一次的数字|简单|63%|178192|
|[0135](https://leetcode-cn.com/problems/candy/)|分发糖果|困难|40%|23011|
|[0134](https://leetcode-cn.com/problems/gas-station/)|加油站|中等|49%|28989|
|[0133](https://leetcode-cn.com/problems/clone-graph/)|克隆图|中等|43%|17350|
|[0132](https://leetcode-cn.com/problems/palindrome-partitioning-ii/)|分割回文串 II|困难|40%|11851|
|[0131](https://leetcode-cn.com/problems/palindrome-partitioning/)|分割回文串|中等|64%|23435|
|[0130](https://leetcode-cn.com/problems/surrounded-regions/)|被围绕的区域|中等|38%|32583|
|[0129](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/)|求根到叶子节点数字之和|中等|59%|19568|
|[0128](https://leetcode-cn.com/problems/longest-consecutive-sequence/)|最长连续序列|困难|46%|37474|
|[0127](https://leetcode-cn.com/problems/word-ladder/)|单词接龙|中等|37%|34114|
|[0126](https://leetcode-cn.com/problems/word-ladder-ii/)|单词接龙 II|困难|28%|12132|
|[0125](https://leetcode-cn.com/problems/valid-palindrome/)|验证回文串|简单|41%|142759|
|[0124](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/)|二叉树中的最大路径和|困难|37%|43264|
|[0123](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|买卖股票的最佳时机 III|困难|39%|35929|
|[0122](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|买卖股票的最佳时机 II|简单|55%|157286|
|[0121](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|[买卖股票的最佳时机](./src/0121-best-time-to-buy-and-sell-stock/)|简单|51%|175141|
|[0120](https://leetcode-cn.com/problems/triangle/)|三角形最小路径和|中等|62%|42852|
|[0119](https://leetcode-cn.com/problems/pascals-triangle-ii/)|杨辉三角 II|简单|58%|46712|
|[0118](https://leetcode-cn.com/problems/pascals-triangle/)|杨辉三角|简单|64%|66086|
|[0117](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)|填充每个节点的下一个右侧节点指针 II|中等|41%|22337|
|[0116](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/)|填充每个节点的下一个右侧节点指针|中等|48%|31710|
|[0115](https://leetcode-cn.com/problems/distinct-subsequences/)|不同的子序列|困难|44%|12799|
|[0114](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)|二叉树展开为链表|中等|65%|25051|
|[0113](https://leetcode-cn.com/problems/path-sum-ii/)|路径总和 II|中等|57%|33007|
|[0112](https://leetcode-cn.com/problems/path-sum/)|路径总和|简单|48%|72755|
|[0111](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)|二叉树的最小深度|简单|40%|86702|
|[0110](https://leetcode-cn.com/problems/balanced-binary-tree/)|平衡二叉树|简单|49%|69364|
|[0109](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/)|有序链表转换二叉搜索树|中等|68%|17715|
|[0108](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)|[将有序数组转换为二叉搜索树](./src/0108-convert-sorted-array-to-binary-search-tree/)|简单|67%|51638|
|[0107](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)|[二叉树的层次遍历 II](./src/0107-binary-tree-level-order-traversal-ii/)|简单|63%|45879|
|[0106](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)|从中序与后序遍历序列构造二叉树|中等|64%|24605|
|[0105](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)|从前序与中序遍历序列构造二叉树|中等|62%|45295|
|[0104](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)|[二叉树的最大深度](./src/0104-maximum-depth-of-binary-tree/)|简单|71%|123766|
|[0103](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)|二叉树的锯齿形层次遍历|中等|52%|40774|
|[0102](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)|二叉树的层次遍历|中等|59%|88864|
|[0101](https://leetcode-cn.com/problems/symmetric-tree/)|[对称二叉树](./src/0101-symmetric-tree/)|简单|48%|132878|
|[0100](https://leetcode-cn.com/problems/same-tree/)|[相同的树](./src/0100-same-tree/)|简单|55%|78165|
|[0099](https://leetcode-cn.com/problems/recover-binary-search-tree/)|恢复二叉搜索树|困难|53%|12704|
|[0098](https://leetcode-cn.com/problems/validate-binary-search-tree/)|验证二叉搜索树|中等|27%|166887|
|[0097](https://leetcode-cn.com/problems/interleaving-string/)|交错字符串|困难|36%|17510|
|[0096](https://leetcode-cn.com/problems/unique-binary-search-trees/)|不同的二叉搜索树|中等|63%|30919|
|[0095](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/)|不同的二叉搜索树 II|中等|59%|21571|
|[0094](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)|二叉树的中序遍历|中等|68%|102880|
|[0093](https://leetcode-cn.com/problems/restore-ip-addresses/)|复原IP地址|中等|45%|37466|
|[0092](https://leetcode-cn.com/problems/reverse-linked-list-ii/)|反转链表 II|中等|47%|49254|
|[0091](https://leetcode-cn.com/problems/decode-ways/)|解码方法|中等|21%|90858|
|[0090](https://leetcode-cn.com/problems/subsets-ii/)|子集 II|中等|57%|26400|
|[0089](https://leetcode-cn.com/problems/gray-code/)|格雷编码|中等|66%|19829|
|[0088](https://leetcode-cn.com/problems/merge-sorted-array/)|[合并两个有序数组](./src/0088-merge-sorted-array/)|简单|45%|173960|
|[0087](https://leetcode-cn.com/problems/scramble-string/)|扰乱字符串|困难|43%|10591|
|[0086](https://leetcode-cn.com/problems/partition-list/)|分隔链表|中等|53%|32142|
|[0085](https://leetcode-cn.com/problems/maximal-rectangle/)|最大矩形|困难|42%|26200|
|[0084](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)|柱状图中最大的矩形|困难|37%|45275|
|[0083](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)|[删除排序链表中的重复元素](./src/0083-remove-duplicates-from-sorted-list/)|简单|48%|108365|
|[0082](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)|删除排序链表中的重复元素 II|中等|44%|49770|
|[0081](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)|搜索旋转排序数组 II|中等|34%|35019|
|[0080](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/)|删除排序数组中的重复项 II|中等|52%|43606|
|[0079](https://leetcode-cn.com/problems/word-search/)|单词搜索|中等|39%|60384|
|[0078](https://leetcode-cn.com/problems/subsets/)|子集|中等|75%|56266|
|[0077](https://leetcode-cn.com/problems/combinations/)|组合|中等|70%|32415|
|[0076](https://leetcode-cn.com/problems/minimum-window-substring/)|最小覆盖子串|困难|34%|41661|
|[0075](https://leetcode-cn.com/problems/sort-colors/)|颜色分类|中等|53%|73351|
|[0074](https://leetcode-cn.com/problems/search-a-2d-matrix/)|搜索二维矩阵|中等|36%|54989|
|[0073](https://leetcode-cn.com/problems/set-matrix-zeroes/)|矩阵置零|中等|54%|38739|
|[0072](https://leetcode-cn.com/problems/edit-distance/)|编辑距离|困难|54%|32341|
|[0071](https://leetcode-cn.com/problems/simplify-path/)|简化路径|中等|37%|45511|
|[0070](https://leetcode-cn.com/problems/climbing-stairs/)|[爬楼梯](./src/0070-climbing-stairs/)|简单|47%|200748|
|[0069](https://leetcode-cn.com/problems/sqrtx/)|[x 的平方根](./src/0069-sqrtx/)|简单|37%|181219|
|[0068](https://leetcode-cn.com/problems/text-justification/)|文本左右对齐|困难|41%|9807|
|[0067](https://leetcode-cn.com/problems/add-binary/)|[二进制求和](./src/0067-add-binary/)|简单|51%|88707|
|[0066](https://leetcode-cn.com/problems/plus-one/)|[加一](./src/0066-plus-one/)|简单|41%|211764|
|[0065](https://leetcode-cn.com/problems/valid-number/)|有效数字|困难|16%|37551|
|[0064](https://leetcode-cn.com/problems/minimum-path-sum/)|最小路径和|中等|62%|60838|
|[0063](https://leetcode-cn.com/problems/unique-paths-ii/)|不同路径 II|中等|31%|83585|
|[0062](https://leetcode-cn.com/problems/unique-paths/)|不同路径|中等|56%|83119|
|[0061](https://leetcode-cn.com/problems/rotate-list/)|旋转链表|中等|39%|76684|
|[0060](https://leetcode-cn.com/problems/permutation-sequence/)|第k个排列|中等|47%|33808|
|[0059](https://leetcode-cn.com/problems/spiral-matrix-ii/)|螺旋矩阵 II|中等|75%|24214|
|[0058](https://leetcode-cn.com/problems/length-of-last-word/)|[最后一个单词的长度](./src/0058-length-of-last-word/)|简单|31%|155169|
|[0057](https://leetcode-cn.com/problems/insert-interval/)|[插入区间](./src/0057-insert-interval/)|困难|35%|28827|
|[0056](https://leetcode-cn.com/problems/merge-intervals/)|合并区间|中等|38%|88198|
|[0055](https://leetcode-cn.com/problems/jump-game/)|跳跃游戏|中等|36%|106201|
|[0054](https://leetcode-cn.com/problems/spiral-matrix/)|螺旋矩阵|中等|37%|75373|
|[0053](https://leetcode-cn.com/problems/maximum-subarray/)|[最大子序和](./src/0053-maximum-subarray/)|简单|47%|235859|
|[0052](https://leetcode-cn.com/problems/n-queens-ii/)|N皇后 II|困难|75%|14204|
|[0051](https://leetcode-cn.com/problems/n-queens/)|N皇后|困难|66%|25480|
|[0050](https://leetcode-cn.com/problems/powx-n/)|Pow(x, n)|中等|33%|107892|
|[0049](https://leetcode-cn.com/problems/group-anagrams/)|字母异位词分组|中等|59%|57332|
|[0048](https://leetcode-cn.com/problems/rotate-image/)|旋转图像|中等|65%|61912|
|[0047](https://leetcode-cn.com/problems/permutations-ii/)|全排列 II|中等|55%|47819|
|[0046](https://leetcode-cn.com/problems/permutations/)|全排列|中等|72%|74570|
|[0045](https://leetcode-cn.com/problems/jump-game-ii/)|跳跃游戏 II|困难|31%|63760|
|[0044](https://leetcode-cn.com/problems/wildcard-matching/)|通配符匹配|困难|25%|49818|
|[0043](https://leetcode-cn.com/problems/multiply-strings/)|[字符串相乘](./src/0043-multiply-strings/)|中等|40%|76515|
|[0042](https://leetcode-cn.com/problems/trapping-rain-water/)|[接雨水](./src/0042-trapping-rain-water/)|困难|46%|72366|
|[0041](https://leetcode-cn.com/problems/first-missing-positive/)|[缺失的第一个正数](./src/0041-first-missing-positive/)|困难|36%|70482|
|[0040](https://leetcode-cn.com/problems/combination-sum-ii/)|[组合总和 II](./src/0040-combination-sum-ii/)|中等|57%|46041|
|[0039](https://leetcode-cn.com/problems/combination-sum/)|[组合总和](./src/0039-combination-sum/)|中等|67%|62484|
|[0038](https://leetcode-cn.com/problems/count-and-say/)|[报数](./src/0038-count-and-say/)|简单|53%|103011|
|[0037](https://leetcode-cn.com/problems/sudoku-solver/)|[解数独](./src/0037-sudoku-solver/)|困难|57%|21752|
|[0036](https://leetcode-cn.com/problems/valid-sudoku/)|[有效的数独](./src/0036-valid-sudoku/)|中等|56%|72728|
|[0035](https://leetcode-cn.com/problems/search-insert-position/)|[搜索插入位置](./src/0035-search-insert-position/)|简单|44%|189694|
|[0034](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)|[在排序数组中查找元素的第一个和最后一个位置](./src/0034-find-first-and-last-position-of-element-in-sorted-array/)|中等|37%|107138|
|[0033](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)|[搜索旋转排序数组](./src/0033-search-in-rotated-sorted-array/)|中等|36%|139308|
|[0032](https://leetcode-cn.com/problems/longest-valid-parentheses/)|[最长有效括号](./src/0032-longest-valid-parentheses/)|困难|28%|89405|
|[0031](https://leetcode-cn.com/problems/next-permutation/)|[下一个排列](./src/0031-next-permutation/)|中等|32%|88330|
|[0030](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)|[串联所有单词的子串](./src/0030-substring-with-concatenation-of-all-words/)|困难|27%|56355|
|[0029](https://leetcode-cn.com/problems/divide-two-integers/)|[两数相除](./src/0029-divide-two-integers/)|中等|18%|129286|
|[0028](https://leetcode-cn.com/problems/implement-strstr/)|[实现 strStr()](./src/0028-implement-strstr/)|简单|39%|241384|
|[0027](https://leetcode-cn.com/problems/remove-element/)|[移除元素](./src/0027-remove-element/)|简单|56%|174879|
|[0026](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)|[删除排序数组中的重复项](./src/0026-remove-duplicates-from-sorted-array/)|简单|46%|397537|
|[0025](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)|[K 个一组翻转链表](./src/0025-reverse-nodes-in-k-group/)|困难|54%|45918|
|[0024](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)|[两两交换链表中的节点](./src/0024-swap-nodes-in-pairs/)|中等|62%|76577|
|[0023](https://leetcode-cn.com/problems/merge-k-sorted-lists/)|[合并K个排序链表](./src/0023-merge-k-sorted-lists/)|困难|47%|108344|
|[0022](https://leetcode-cn.com/problems/generate-parentheses/)|[括号生成](./src/0022-generate-parentheses/)|中等|72%|68544|
|[0021](https://leetcode-cn.com/problems/merge-two-sorted-lists/)|[合并两个有序链表](./src/0021-merge-two-sorted-lists/)|简单|57%|224853|
|[0020](https://leetcode-cn.com/problems/valid-parentheses/)|[有效的括号](./src/0020-valid-parentheses/)|简单|39%|362488|
|[0019](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)|[删除链表的倒数第N个节点](./src/0019-remove-nth-node-from-end-of-list/)|中等|36%|238487|
|[0018](https://leetcode-cn.com/problems/4sum/)|[四数之和](./src/0018-4sum/)|中等|36%|103495|
|[0017](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)|[电话号码的字母组合](./src/0017-letter-combinations-of-a-phone-number/)|中等|51%|103810|
|[0016](https://leetcode-cn.com/problems/3sum-closest/)|[最接近的三数之和](./src/0016-3sum-closest/)|中等|41%|121852|
|[0015](https://leetcode-cn.com/problems/3sum/)|[三数之和](./src/0015-3sum/)|中等|24%|456081|
|[0014](https://leetcode-cn.com/problems/longest-common-prefix/)|[最长公共前缀](./src/0014-longest-common-prefix/)|简单|35%|397977|
|[0013](https://leetcode-cn.com/problems/roman-to-integer/)|[罗马数字转整数](./src/0013-roman-to-integer/)|简单|59%|184783|
|[0012](https://leetcode-cn.com/problems/integer-to-roman/)|[整数转罗马数字](./src/0012-integer-to-roman/)|中等|61%|76769|
|[0011](https://leetcode-cn.com/problems/container-with-most-water/)|[盛最多水的容器](./src/0011-container-with-most-water/)|中等|58%|160857|
|[0010](https://leetcode-cn.com/problems/regular-expression-matching/)|[正则表达式匹配](./src/0010-regular-expression-matching/)|困难|25%|143137|
|[0009](https://leetcode-cn.com/problems/palindrome-number/)|[回文数](./src/0009-palindrome-number/)|简单|56%|331863|
|[0008](https://leetcode-cn.com/problems/string-to-integer-atoi/)|[字符串转换整数 (atoi)](./src/0008-string-to-integer-atoi/)|中等|18%|436877|
|[0007](https://leetcode-cn.com/problems/reverse-integer/)|[整数反转](./src/0007-reverse-integer/)|简单|33%|631460|
|[0006](https://leetcode-cn.com/problems/zigzag-conversion/)|[Z 字形变换](./src/0006-zigzag-conversion/)|中等|45%|149453|
|[0005](https://leetcode-cn.com/problems/longest-palindromic-substring/)|[最长回文子串](./src/0005-longest-palindromic-substring/)|中等|27%|474289|
|[0004](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)|[寻找两个有序数组的中位数](./src/0004-median-of-two-sorted-arrays/)|困难|36%|293466|
|[0003](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)|[无重复字符的最长子串](./src/0003-longest-substring-without-repeating-characters/)|中等|31%|790214|
|[0002](https://leetcode-cn.com/problems/add-two-numbers/)|[两数相加](./src/0002-add-two-numbers/)|中等|35%|674291|
|[0001](https://leetcode-cn.com/problems/two-sum/)|[两数之和](./src/0001-two-sum/)|简单|46%|1363241|

---

#### Similar Problem

|#|标题|
|:-:|:-|
|[0140](https://leetcode-cn.com/problems/word-break-ii/)|单词拆分 II|
|[0139](https://leetcode-cn.com/problems/word-break/)|单词拆分|
||||
||||
|[0096](https://leetcode-cn.com/problems/unique-binary-search-trees/)|不同的二叉搜索树|
|[0095](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/)|不同的二叉搜索树 II|
||||
||||
|[1174](https://leetcode-cn.com/problems/immediate-food-delivery-ii/)|销售分析III|
|[1173](https://leetcode-cn.com/problems/immediate-food-delivery-i/)|销售分析 II|
|[1172](https://leetcode-cn.com/problems/dinner-plate-stacks/)|销售分析 I |
|[1155](https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/)|产品销售分析 III|
|[1154](https://leetcode-cn.com/problems/day-of-the-year/)|产品销售分析 II|
|[1153](https://leetcode-cn.com/problems/string-transforms-into-another-string/)|产品销售分析 I|
||||
||||
|[1155](https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/)|产品销售分析 III|
|[1154](https://leetcode-cn.com/problems/day-of-the-year/)|产品销售分析 II|
|[1153](https://leetcode-cn.com/problems/string-transforms-into-another-string/)|产品销售分析 I|
||||
||||
|[1000](https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/)|删列造序 III|
|[0992](https://leetcode-cn.com/problems/subarrays-with-k-different-integers/)|删列造序 II|
|[0981](https://leetcode-cn.com/problems/time-based-key-value-store/)|[删列造序](./src/0981-time-based-key-value-store/)|
||||
||||
|[0533](https://leetcode-cn.com/problems/lonely-pixel-ii/)|孤独像素 II|
|[0531](https://leetcode-cn.com/problems/lonely-pixel-i/)|孤独像素 I|
||||
||||
|[0407](https://leetcode-cn.com/problems/trapping-rain-water-ii/)|接雨水 II|
|[0042](https://leetcode-cn.com/problems/trapping-rain-water/)|[接雨水](./src/0042-trapping-rain-water/)|
||||
||||
|[0248](https://leetcode-cn.com/problems/strobogrammatic-number-iii/)|中心对称数 III|
|[0247](https://leetcode-cn.com/problems/strobogrammatic-number-ii/)|中心对称数 II|
|[0246](https://leetcode-cn.com/problems/strobogrammatic-number/)|中心对称数|
||||
||||
|[0055](https://leetcode-cn.com/problems/jump-game/)|跳跃游戏|
|[0045](https://leetcode-cn.com/problems/jump-game-ii/)|跳跃游戏 II|
||||
||||
|[1130](https://leetcode-cn.com/problems/minimum-cost-tree-from-leaf-values/)|最后一块石头的重量 II|
|[1127](https://leetcode-cn.com/problems/user-purchase-platform/)|最后一块石头的重量|
||||
||||
|[0737](https://leetcode-cn.com/problems/sentence-similarity-ii/)|句子相似性 II|
|[0734](https://leetcode-cn.com/problems/sentence-similarity/)|句子相似性|
||||
||||
|[0253](https://leetcode-cn.com/problems/meeting-rooms-ii/)|会议室 II|
|[0252](https://leetcode-cn.com/problems/meeting-rooms/)|会议室|
||||
||||
|[0260](https://leetcode-cn.com/problems/single-number-iii/)|只出现一次的数字 III|
|[0137](https://leetcode-cn.com/problems/single-number-ii/)|只出现一次的数字 II|
|[0136](https://leetcode-cn.com/problems/single-number/)|只出现一次的数字|
||||
||||
|[0119](https://leetcode-cn.com/problems/pascals-triangle-ii/)|杨辉三角 II|
|[0118](https://leetcode-cn.com/problems/pascals-triangle/)|杨辉三角|
||||
||||
|[1040](https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/)|最大二叉树 II|
|[0654](https://leetcode-cn.com/problems/maximum-binary-tree/)|最大二叉树|
||||
||||
|[0518](https://leetcode-cn.com/problems/coin-change-2/)|零钱兑换 II|
|[0322](https://leetcode-cn.com/problems/coin-change/)|零钱兑换|
||||
||||
|[0220](https://leetcode-cn.com/problems/contains-duplicate-iii/)|存在重复元素 III|
|[0219](https://leetcode-cn.com/problems/contains-duplicate-ii/)|存在重复元素 II|
|[0217](https://leetcode-cn.com/problems/contains-duplicate/)|存在重复元素|
||||
||||
|[0142](https://leetcode-cn.com/problems/linked-list-cycle-ii/)|环形链表 II|
|[0141](https://leetcode-cn.com/problems/linked-list-cycle/)|环形链表|
||||
||||
|[0117](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)|填充每个节点的下一个右侧节点指针 II|
|[0116](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/)|填充每个节点的下一个右侧节点指针|
||||
||||
|[0080](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/)|删除排序数组中的重复项 II|
|[0026](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)|[删除排序数组中的重复项](./src/0026-remove-duplicates-from-sorted-array/)|
||||
||||
|[0711](https://leetcode-cn.com/problems/number-of-distinct-islands-ii/)|不同岛屿的数量 II|
|[0694](https://leetcode-cn.com/problems/number-of-distinct-islands/)|不同岛屿的数量|
||||
||||
|[0132](https://leetcode-cn.com/problems/palindrome-partitioning-ii/)|分割回文串 II|
|[0131](https://leetcode-cn.com/problems/palindrome-partitioning/)|分割回文串|
||||
||||
|[1022](https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/)|[不同路径 III](./src/1022-sum-of-root-to-leaf-binary-numbers/)|
|[0063](https://leetcode-cn.com/problems/unique-paths-ii/)|不同路径 II|
|[0062](https://leetcode-cn.com/problems/unique-paths/)|不同路径|
||||
||||
|[0509](https://leetcode-cn.com/problems/fibonacci-number/)|二叉搜索树中的中序后继 II|
||||
||||
|[0557](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)|反转字符串中的单词 III|
|[0541](https://leetcode-cn.com/problems/reverse-string-ii/)|反转字符串 II|
|[0345](https://leetcode-cn.com/problems/reverse-vowels-of-a-string/)|反转字符串中的元音字母|
|[0344](https://leetcode-cn.com/problems/reverse-string/)|反转字符串|
||||
||||
|[0294](https://leetcode-cn.com/problems/flip-game-ii/)|翻转游戏 II|
|[0293](https://leetcode-cn.com/problems/flip-game/)|翻转游戏|
||||
||||
|[0154](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)|寻找旋转排序数组中的最小值 II|
|[0153](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)|寻找旋转排序数组中的最小值|
||||
||||
|[0952](https://leetcode-cn.com/problems/largest-component-size-by-common-factor/)|单词子集|
|[0698](https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/)|划分为k个相等的子集|
|[0416](https://leetcode-cn.com/problems/partition-equal-subset-sum/)|分割等和子集|
|[0368](https://leetcode-cn.com/problems/largest-divisible-subset/)|最大整除子集|
|[0090](https://leetcode-cn.com/problems/subsets-ii/)|子集 II|
|[0078](https://leetcode-cn.com/problems/subsets/)|子集|
||||
||||
|[0800](https://leetcode-cn.com/problems/similar-rgb-color/)|字母大小写全排列|
|[0047](https://leetcode-cn.com/problems/permutations-ii/)|全排列 II|
|[0046](https://leetcode-cn.com/problems/permutations/)|全排列|
||||
||||
|[0714](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)|买卖股票的最佳时机含手续费|
|[0188](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)|买卖股票的最佳时机 IV|
|[0123](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|买卖股票的最佳时机 III|
|[0122](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|买卖股票的最佳时机 II|
|[0121](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|[买卖股票的最佳时机](./src/0121-best-time-to-buy-and-sell-stock/)|
||||
||||
|[0083](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)|[删除排序链表中的重复元素](./src/0083-remove-duplicates-from-sorted-list/)|
|[0082](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)|删除排序链表中的重复元素 II|
||||
||||
|[1113](https://leetcode-cn.com/problems/reported-posts/)|移动石子直到连续 II|
|[1103](https://leetcode-cn.com/problems/distribute-candies-to-people/)|移动石子直到连续|
||||
||||
|[0946](https://leetcode-cn.com/problems/validate-stack-sequences/)|最小差值 II|
|[0944](https://leetcode-cn.com/problems/delete-columns-to-make-sorted/)|最小差值 I|
||||
||||
|[0598](https://leetcode-cn.com/problems/range-addition-ii/)|范围求和 II|
||||
||||
|[0556](https://leetcode-cn.com/problems/next-greater-element-iii/)|下一个更大元素 III|
|[0503](https://leetcode-cn.com/problems/next-greater-element-ii/)|[下一个更大元素 II](./src/0503-next-greater-element-ii/)|
|[0496](https://leetcode-cn.com/problems/next-greater-element-i/)|下一个更大元素 I|
||||
||||
|[0350](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)|两个数组的交集 II|
|[0349](https://leetcode-cn.com/problems/intersection-of-two-arrays/)|两个数组的交集|
||||
||||
|[0245](https://leetcode-cn.com/problems/shortest-word-distance-iii/)|最短单词距离 III|
|[0244](https://leetcode-cn.com/problems/shortest-word-distance-ii/)|最短单词距离 II|
|[0243](https://leetcode-cn.com/problems/shortest-word-distance/)|最短单词距离|
||||
||||
|[0522](https://leetcode-cn.com/problems/longest-uncommon-subsequence-ii/)|最长特殊序列 II|
|[0521](https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/)|最长特殊序列 Ⅰ|
||||
||||
|[0375](https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/)|猜数字大小 II|
|[0374](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)|[猜数字大小](./src/0374-guess-number-higher-or-lower/)|
||||
||||
|[0324](https://leetcode-cn.com/problems/wiggle-sort-ii/)|摆动排序 II|
|[0280](https://leetcode-cn.com/problems/wiggle-sort/)|摆动排序|
||||
||||
|[0229](https://leetcode-cn.com/problems/majority-element-ii/)|求众数 II|
|[0169](https://leetcode-cn.com/problems/majority-element/)|求众数|
||||
||||
|[1193](https://leetcode-cn.com/problems/monthly-transactions-i/)|游戏玩法分析 V|
|[1179](https://leetcode-cn.com/problems/reformat-department-table/)|游戏玩法分析 I|
|[1182](https://leetcode-cn.com/problems/shortest-distance-to-target-color/)|游戏玩法分析 IV|
|[1181](https://leetcode-cn.com/problems/before-and-after-puzzle/)|游戏玩法分析 III|
|[1180](https://leetcode-cn.com/problems/count-substrings-with-only-one-distinct-letter/)|游戏玩法分析 II|
||||
||||
|[0977](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)|不同的子序列 II|
|[0115](https://leetcode-cn.com/problems/distinct-subsequences/)|不同的子序列|
||||
||||
|[0644](https://leetcode-cn.com/problems/maximum-average-subarray-ii/)|最大平均子段和 II|
||||
||||
|[0267](https://leetcode-cn.com/problems/palindrome-permutation-ii/)|回文排列 II|
|[0266](https://leetcode-cn.com/problems/palindrome-permutation/)|回文排列|
||||
||||
|[0212](https://leetcode-cn.com/problems/word-search-ii/)|单词搜索 II|
|[0079](https://leetcode-cn.com/problems/word-search/)|单词搜索|
||||
||||
|[0337](https://leetcode-cn.com/problems/house-robber-iii/)|打家劫舍 III|
|[0213](https://leetcode-cn.com/problems/house-robber-ii/)|打家劫舍 II|
|[0198](https://leetcode-cn.com/problems/house-robber/)|[打家劫舍](./src/0198-house-robber/)|
||||
||||
|[0880](https://leetcode-cn.com/problems/decoded-string-at-index/)|矩形面积 II|
|[0223](https://leetcode-cn.com/problems/rectangle-area/)|矩形面积|
||||
||||
|[0667](https://leetcode-cn.com/problems/beautiful-arrangement-ii/)|优美的排列 II|
|[0526](https://leetcode-cn.com/problems/beautiful-arrangement/)|优美的排列|
||||
||||
|[0454](https://leetcode-cn.com/problems/4sum-ii/)|四数相加 II|
||||
||||
|[0445](https://leetcode-cn.com/problems/add-two-numbers-ii/)|两数相加 II|
|[0002](https://leetcode-cn.com/problems/add-two-numbers/)|[两数相加](./src/0002-add-two-numbers/)|
||||
||||
|[0305](https://leetcode-cn.com/problems/number-of-islands-ii/)|岛屿数量 II|
|[0200](https://leetcode-cn.com/problems/number-of-islands/)|岛屿数量|
||||
||||
|[0785](https://leetcode-cn.com/problems/is-graph-bipartite/)|基本计算器 III|
|[0781](https://leetcode-cn.com/problems/rabbits-in-forest/)|基本计算器 IV|
|[0227](https://leetcode-cn.com/problems/basic-calculator-ii/)|基本计算器 II|
|[0224](https://leetcode-cn.com/problems/basic-calculator/)|基本计算器|
||||
||||
|[1003](https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions/)|最小面积矩形 II|
|[0976](https://leetcode-cn.com/problems/largest-perimeter-triangle/)|[最小面积矩形](./src/0976-largest-perimeter-triangle/)|
||||
||||
|[0552](https://leetcode-cn.com/problems/student-attendance-record-ii/)|学生出勤记录 II|
|[0551](https://leetcode-cn.com/problems/student-attendance-record-i/)|学生出勤记录 I|
||||
||||
|[0462](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/)|最少移动次数使数组元素相等 II|
||||
||||
|[0272](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii/)|最接近的二叉搜索树值 II|
|[0270](https://leetcode-cn.com/problems/closest-binary-search-tree-value/)|最接近的二叉搜索树值|
||||
||||
|[0921](https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/)|螺旋矩阵 III|
|[0059](https://leetcode-cn.com/problems/spiral-matrix-ii/)|螺旋矩阵 II|
|[0054](https://leetcode-cn.com/problems/spiral-matrix/)|螺旋矩阵|
||||
||||
|[1077](https://leetcode-cn.com/problems/project-employees-iii/)|易混淆数 II|
|[1069](https://leetcode-cn.com/problems/product-sales-analysis-ii/)|易混淆数|
||||
||||
|[1052](https://leetcode-cn.com/problems/grumpy-bookstore-owner/)|校园自行车分配|
|[1067](https://leetcode-cn.com/problems/digit-count-in-range/)|校园自行车分配 II|
||||
||||
|[0685](https://leetcode-cn.com/problems/redundant-connection-ii/)|冗余连接 II|
|[0684](https://leetcode-cn.com/problems/redundant-connection/)|冗余连接|
||||
||||
|[0364](https://leetcode-cn.com/problems/nested-list-weight-sum-ii/)|加权嵌套序列和 II|
||||
||||
|[0186](https://leetcode-cn.com/problems/reverse-words-in-a-string-ii/)|翻转字符串里的单词 II|
|[0151](https://leetcode-cn.com/problems/reverse-words-in-a-string/)|翻转字符串里的单词|
||||
||||
|[1188](https://leetcode-cn.com/problems/design-bounded-blocking-queue/)|花括号展开 II|
||||
||||
|[0780](https://leetcode-cn.com/problems/reaching-points/)|最多能完成排序的块|
|[0779](https://leetcode-cn.com/problems/k-th-symbol-in-grammar/)|最多能完成排序的块 II|
||||
||||
|[1106](https://leetcode-cn.com/problems/parsing-a-boolean-expression/)|逃离大迷宫|
|[0505](https://leetcode-cn.com/problems/the-maze-ii/)|迷宫 II|
|[0499](https://leetcode-cn.com/problems/the-maze-iii/)|迷宫 III|
|[0490](https://leetcode-cn.com/problems/the-maze/)|迷宫|
||||
||||
|[0275](https://leetcode-cn.com/problems/h-index-ii/)|H指数 II|
|[0274](https://leetcode-cn.com/problems/h-index/)|H指数|
||||
||||
|[0377](https://leetcode-cn.com/problems/combination-sum-iv/)|组合总和 Ⅳ|
|[0216](https://leetcode-cn.com/problems/combination-sum-iii/)|组合总和 III|
|[0040](https://leetcode-cn.com/problems/combination-sum-ii/)|[组合总和 II](./src/0040-combination-sum-ii/)|
|[0039](https://leetcode-cn.com/problems/combination-sum/)|[组合总和](./src/0039-combination-sum/)|
||||
||||
|[1195](https://leetcode-cn.com/problems/fizz-buzz-multithreaded/)|分糖果 II|
|[0575](https://leetcode-cn.com/problems/distribute-candies/)|分糖果|
||||
||||
|[0958](https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/)|按奇偶排序数组 II|
|[0941](https://leetcode-cn.com/problems/valid-mountain-array/)|按奇偶排序数组|
||||
||||
|[1046](https://leetcode-cn.com/problems/last-stone-weight/)|最大连续1的个数 III|
|[0487](https://leetcode-cn.com/problems/max-consecutive-ones-ii/)|最大连续1的个数 II|
|[0485](https://leetcode-cn.com/problems/max-consecutive-ones/)|最大连续1的个数|
||||
||||
|[0291](https://leetcode-cn.com/problems/word-pattern-ii/)|单词规律 II|
|[0290](https://leetcode-cn.com/problems/word-pattern/)|单词规律|
||||
||||
|[0265](https://leetcode-cn.com/problems/paint-house-ii/)|粉刷房子 II|
|[0256](https://leetcode-cn.com/problems/paint-house/)|粉刷房子|
||||
||||
|[0630](https://leetcode-cn.com/problems/course-schedule-iii/)|课程表 III|
|[0210](https://leetcode-cn.com/problems/course-schedule-ii/)|课程表 II|
|[0207](https://leetcode-cn.com/problems/course-schedule/)|课程表|
||||
||||
|[0964](https://leetcode-cn.com/problems/least-operators-to-express-number/)|尽量减少恶意软件的传播 II|
|[0960](https://leetcode-cn.com/problems/delete-columns-to-make-sorted-iii/)|尽量减少恶意软件的传播|
||||
||||
|[0313](https://leetcode-cn.com/problems/super-ugly-number/)|超级丑数|
|[0264](https://leetcode-cn.com/problems/ugly-number-ii/)|丑数 II|
|[0263](https://leetcode-cn.com/problems/ugly-number/)|丑数|
||||
||||
|[0081](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)|搜索旋转排序数组 II|
|[0033](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)|[搜索旋转排序数组](./src/0033-search-in-rotated-sorted-array/)|
||||
||||
|[0240](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)|搜索二维矩阵 II|
|[0074](https://leetcode-cn.com/problems/search-a-2d-matrix/)|搜索二维矩阵|
||||
||||
|[0127](https://leetcode-cn.com/problems/word-ladder/)|单词接龙|
|[0126](https://leetcode-cn.com/problems/word-ladder-ii/)|单词接龙 II|
||||
||||
|[0107](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)|[二叉树的层次遍历 II](./src/0107-binary-tree-level-order-traversal-ii/)|
|[0102](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)|二叉树的层次遍历|
||||
||||
|[0206](https://leetcode-cn.com/problems/reverse-linked-list/)|反转链表|
|[0092](https://leetcode-cn.com/problems/reverse-linked-list-ii/)|反转链表 II|
||||
||||
|[0052](https://leetcode-cn.com/problems/n-queens-ii/)|N皇后 II|
|[0051](https://leetcode-cn.com/problems/n-queens/)|N皇后|
||||
||||
|[0732](https://leetcode-cn.com/problems/my-calendar-iii/)|我的日程安排表 III|
|[0731](https://leetcode-cn.com/problems/my-calendar-ii/)|我的日程安排表 II|
|[0729](https://leetcode-cn.com/problems/my-calendar-i/)|我的日程安排表 I|
||||
||||
|[0437](https://leetcode-cn.com/problems/path-sum-iii/)|路径总和 III|
|[0113](https://leetcode-cn.com/problems/path-sum-ii/)|路径总和 II|
|[0112](https://leetcode-cn.com/problems/path-sum/)|路径总和|
||||
||||

---

[⬆️Top](#lccli)
