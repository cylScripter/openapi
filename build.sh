#!/bin/bash

# 检查是否提供了版本参数
if [ -z "$1" ]; then
  echo "请提供服务名，例如: ./build.sh system"
  exit 1
fi



# 检查是否提供了版本参数
if [ -z "$2" ]; then
  echo "请提供版本号，例如: ./build.sh system  v1.0.1 "
  exit 1
fi

kitex -module github.com/cylScripter/openapi -template-dir /Users/cyl/project/kitex/client/standard2 -gen-path ./ -I ~/project/idl $1.thrift

git add .
git commit -m "feat: $1 $2"

git tag $2
git push origin $2
