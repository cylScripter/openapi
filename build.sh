#!/bin/bash

# 检查是否提供了版本参数
if [ -z "$1" ]; then
  echo "请提供服务名，例如: ./build.sh system"
  exit 1
fi



## 随机生成版本号 版本号格式为 v25-01-01-000000
function generate_random_version() {
  echo "v$(date +%y%m%d%H%M%S)"
   return
}


version=$(generate_random_version)

echo "生成版本号: $version"

kitex -module github.com/cylScripter/openapi -template-dir /Users/cyl/project/kitex/client/standard2 -gen-path ./ -I ~/project/idl $1.thrift

git add .
git commit -m "feat: $1 $version"

git tag $version
git push origin $version
