<div align="center">

![leetcli](https://raw.githubusercontent.com/evercyan/cantor/master/resource/69/69f055fa7ccfe73114bf6608a2789d8f.png)

[![travis-ci](https://www.travis-ci.org/evercyan/leetcli.svg?branch=master)](https://www.travis-ci.org/evercyan/leetcli)
[![codecov](https://codecov.io/gh/evercyan/leetcli/branch/master/graph/badge.svg?token=RbJTUtAlvl)](https://codecov.io/gh/evercyan/leetcli)
[![goreportcard](https://goreportcard.com/badge/github.com/evercyan/leetcli)](https://goreportcard.com/report/github.com/evercyan/leetcli)

leetcode 刷题工具

</div>

---

## What's this?

刷题过程, 一般是酱紫的

- 1. 访问 leetcode, 选一道题准备攻克它
- 2. 本地建一个答题的目录
- 3. 进到答题目录里, 新建答题文件和测试文件
- 4. 切到浏览器, 复制答题对应语言的模板, 粘贴到答题文件中
- 5. 一毛一样, 复制相关测试用例, 粘贴到测试文件里
- 6. 开始攻克难关, 本地跑下单测, 然后提交到 leetcode

从 2-5 这几步的准备工作冗余繁琐, 每一道题都需要重复往之, 十分郁闷

他来了他来了, 他带着利器赶来了, 这个利器将缩短答题前置的流程, 可快速实现上述 2-5 步

- 1. 访问 leetcode, 选一道题准备攻克它
- 2. 复制题目标识, `question 标识` 直接生成答题目录, 答题文件, 测试文件
- 3. ......

---

## How to install?

- 有 `golang` 相关经验的, 可直接 `go get -v github.com/evercyan/leetcli`
- 亦可进入 https://github.com/evercyan/leetcli/releases/tag/v0.0.6 页面, 选择对应的平台软件包进行下载

---

## How to use?

终端 `leetcli --help` 或可 `leetcli` 进入交互模式

![leetcli-help](https://raw.githubusercontent.com/evercyan/cantor/master/resource/66/662dfac5a7e65a5d0ce48a5574ca71f6.png)

- `config` 答题配置
- `record` 生成答题记录文件 Record.md
- `question` 生成答题文件 [eg: question two-sum], two-sum 为题目链接中的标识
- `list` 问题列表 [eg: list two-sum; 最多显示 20 条]

### 配置

- 答题目录配置

初次使用需要配置答题目录, 通过 `config path xxx` 设置, 方便存放答题文件, 最好是个 git 仓库

![leetcli-help-config](https://raw.githubusercontent.com/evercyan/cantor/master/resource/aa/aafaa8f1330bb715116939be9e8ff834.png)

- 默认编程语言配置

一般题型是支持多种编程语言的, 如果你不想在生成答题文件时再选择一遍编辑语言, 可以使用 `config lang xxx` 设置默认编程语言

---

## How to test?

目前只支持生成 `golang` 和 `javascript` 的测试文件, 其余待支持

在生成答题文件时, 会在题目目录生成 `solution.xxx` 和 `solution_test.xxx`

你需要做的是, 是在 `solution.go` 中完成解题, 在 `solution_test.xxx` 补全测试用例

### golang

`go test -v ./src/1000023-the-masseuse-lcci`

```
=== RUN   TestSolution
=== RUN   TestSolution/Test
=== RUN   TestSolution/Test#01
=== RUN   TestSolution/Test#02
=== RUN   TestSolution/Test#03
=== RUN   TestSolution/Test#04
   PASS: TestSolution (0.00s)
        PASS: TestSolution/Test (0.00s)
        PASS: TestSolution/Test#01 (0.00s)
        PASS: TestSolution/Test#02 (0.00s)
        PASS: TestSolution/Test#03 (0.00s)
        PASS: TestSolution/Test#04 (0.00s)
PASS
ok  	github.com/evercyan/leetcli/src/1000023-the-masseuse-lcci	0.012s
```

### javascript

`javascript` 需要安装 `jest`

```sh
# 安装 jest
npm install -g jest

# 进入项目目录
cd xxx/leetcli

# 初始化
npm init

# 测试
jest

# 示例
 PASS  src/0142-linked-list-cycle-ii/solution.test.js
  ✓ detectCycle (2 ms)

Test Suites: 1 passed, 1 total
Tests:       1 passed, 1 total
Snapshots:   0 total
Time:        1.172 s
Ran all test suites.
```