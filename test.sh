#!/bin/bash

# eg:
# 执行全部测试用例
# sh test.sh
# 执行单个测试用例(目录名中包含 0001 且数量为 1 时才会执行)
# sh test.sh 0001

set -e

tag="leetcli"
if [ "$1" != "" ];
then
    tag=$1
fi

cfile="coverage.txt"
hfile="coverage.html"
echo "" > $cfile
echo "" > $hfile

if [ "${tag}" != "leetcli" -a $(go list ./src/... | grep $tag | wc -l) == 1 ];
then
    # 如果符合条件的只有一个测试用例, 测试之, 生成覆盖率可视 html
    f=$(go list ./src/... | grep $tag)
    go test -v -coverprofile=$cfile -covermode=atomic $f
    go tool cover -html=$cfile -o $hfile
else
    # 执行全部测试用例
    for d in $(go list ./src/... | grep -v vendor); do
        go test -coverprofile=$hfile -covermode=atomic $d
        if [ -f $hfile ];
        then
            cat $hfile >> $cfile
        fi
    done
fi

