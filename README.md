<div align="center">

![leetcli](https://raw.githubusercontent.com/evercyan/cantor/master/resource/69/69f055fa7ccfe73114bf6608a2789d8f.png)

[![travis-ci](https://www.travis-ci.org/evercyan/leetcli.svg?branch=master)](https://www.travis-ci.org/evercyan/leetcli)
[![codecov](https://codecov.io/gh/evercyan/leetcli/branch/master/graph/badge.svg?token=RbJTUtAlvl)](https://codecov.io/gh/evercyan/leetcli)
[![goreportcard](https://goreportcard.com/badge/github.com/evercyan/leetcli)](https://goreportcard.com/report/github.com/evercyan/leetcli)

</div>

---

## What's this?

这是一个有点小帅的 `leetcode` 刷题工具

普通的刷题过程, 一般是酱紫的:

- 1. 打开 leetcode 网站, 精挑细选一道难题准备攻克它
- 2. 本地找个地方建一个答题目录
- 3. 进到答题目录里, 新建答题文件和测试文件
- 4. 切到浏览器, 复制题止页面右侧对应语言的模板代码, 粘贴到答题文件中
- 5. 再切到浏览器, 复制题目左侧题目描述中测试用例, 粘贴到测试文件里
- 6. 开始攻克难关, 本地跑单测, 再提交到 leetcode, 祈祷一遍过

解题仍需你的努力, 但中间几步的冗余繁琐的准备工作, 此可取而代之

- 1. 选题...
- 2. 复制题目标识, 打开工具, `question 问题标识` 直接生成答题目录, 答题文件, 测试文件
- 3. 解题...

---

## How to install?

- `go get -v github.com/evercyan/leetcli`
- 或者进到 https://github.com/evercyan/leetcli/releases/tag/v0.0.5 页面, 下载相应平台软件包

---

## How to use?

如果是通过下载软件包方式安装, 建议可放至全局目录, 方便随时随地启动

- 终端 `leetcli --help` 查看帮助手册, 可单步式使用
- 终端 `leetcli` 进入交互模式, 持续刷题少不了

![leetcli-help](https://raw.githubusercontent.com/evercyan/cantor/master/resource/fb/fb9d9228546d63375b4522cbd55806ea.png)

- 问题标识

如 `https://leetcode-cn.com/problems/two-sum/` 这道 leetcode 届的 `hello world` 题

其问题标识为 `two-sum`, 生成答题文件可直接 `question two-sum`

- 答题目录配置

初次使用需要配置答题目录, 通过 `config path xxx` 设置, 方便存放答题文件

![leetcli-help-config](https://raw.githubusercontent.com/evercyan/cantor/master/resource/aa/aafaa8f1330bb715116939be9e8ff834.png)

- 默认编程语言配置

如果你不想在生成某道答题文件时再选择一遍编辑语言, 可以使用 `config lang xxx` 设置默认编程语言

---

## How to test?

目前只支持生成 `golang` 和 `javascript` 的答题测试文件

执行 `question two-sum` 时, 会在题目目录生成 `solution.xxx` 和 `solution_test.xxx`

你需要做的是, 是在 `solution.xxx` 中完成解题, 在 `solution_test.xxx` 补全相关测试用例

### golang

`go test -v ./src/1000023-the-masseuse-lcci`

```sh
# 单测输出示例
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

执行 `javascript` 测试文件需要本地安装 `npm` 包 `jest`

```sh
# 安装 jest
npm install -g jest

# 初始化
npm init

# 启动单测
jest
```

```
# 单测输出示例
 PASS  src/0142-linked-list-cycle-ii/solution.test.js
  ✓ detectCycle (2 ms)

Test Suites: 1 passed, 1 total
Tests:       1 passed, 1 total
Snapshots:   0 total
Time:        1.172 s
Ran all test suites.
```
