#!/bin/bash

# 检查protoc是否安装
if ! command -v protoc &> /dev/null; then
    echo "错误: protoc 未安装，请先安装 Protocol Buffers 编译器"
    exit 1
fi

# 检查是否安装了Go的protoc插件
if ! protoc --go_out=. --go_opt=paths=source_relative --version &> /dev/null; then
    echo "错误: 未安装Go的protoc插件，请先安装: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
    exit 1
fi

# 查找所有.proto文件并编译
find . -name "*.proto" | while read -r proto_file; do
    echo "正在编译: $proto_file"
    protoc --go_out=. --go_opt=paths=source_relative "$proto_file"
    
    if [ $? -eq 0 ]; then
        echo "成功编译: $proto_file"
    else
        echo "错误: 编译失败: $proto_file"
    fi
done

echo "所有.proto文件处理完成"
