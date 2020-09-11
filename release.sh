#!/bin/bash

set -e

version=$1
if [ "${version}" == "" ];
then
    echo "sh release.sh v0.0.1"
    exit
fi

rm -rf ./dist
git tag -a $version -m $version
git push origin master
git push origin $version
goreleaser
