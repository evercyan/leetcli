#!/bin/sh

# 1, 读取文件内容
# 2, 将空格替换成换行
# 3, 过滤换行
# 4, 排序
# 5, 去重统计数字
# 6, 排序
# 7, 按格式输出

cat words.txt | tr " " "\n" | tr -s '\n' | sort | uniq -c | sort -rg | awk '{print $2" "$1}'
