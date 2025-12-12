#!/bin/bash

# Zabbix 6.0/7.0 多版本支持测试脚本

set -e

echo "=== Zabbix 多版本支持测试 ==="
echo "开始时间: $(date)"
echo

# 检查 Go 环境
if ! command -v go &> /dev/null; then
    echo "错误: Go 编译器未找到，请先安装 Go"
    exit 1
fi

echo "Go 版本: $(go version)"
echo

# 运行单元测试
echo "=== 运行单元测试 ==="
echo "测试版本管理器和特性检测..."
go test -v -run TestVersionManager
go test -v -run TestFeatureDetection
go test -v -run TestHeaderConversion
go test -v -run TestItemValidation
go test -v -run TestHostValidation
go test -v -run TestBrowserItemValidation
go test -v -run TestVersionCompatibility
go test -v -run TestAdapterInterface
go test -v -run TestFeatureConstants

echo
echo "=== 运行集成测试 ==="
echo "测试多版本功能集成..."
go test -v -run TestMultiVersionIntegration
go test -v -run TestItemIntegration
go test -v -run TestHostIntegration
go test -v -run TestBrowserItemIntegration
go test -v -run TestMFAIntegration
go test -v -run TestProxyGroupIntegration
go test -v -run TestHistoryPushIntegration
go test -v -run TestSupportedFeaturesIntegration
go test -v -run TestErrorHandling

echo
echo "=== 运行性能测试 ==="
echo "测试 Header 转换性能..."
go test -v -bench=BenchmarkHeaderConversion -benchmem

echo
echo "测试版本检测性能..."
go test -v -bench=BenchmarkVersionDetection -benchmem

echo
echo "=== 运行所有测试 ==="
echo "运行完整的测试套件..."
go test -v -coverprofile=coverage.out

echo
echo "=== 生成覆盖率报告 ==="
if command -v go &> /dev/null; then
    go tool cover -html=coverage.out -o coverage.html
    echo "覆盖率报告已生成: coverage.html"
    echo "总体覆盖率: $(go tool cover -func=coverage.out | grep total | awk '{print $3}')"
fi

echo
echo "=== 代码检查 ==="
if command -v gofmt &> /dev/null; then
    echo "检查代码格式..."
    unformatted=$(gofmt -l .)
    if [ -n "$unformatted" ]; then
        echo "以下文件需要格式化:"
        echo "$unformatted"
        echo "运行: gofmt -w ."
    else
        echo "代码格式检查通过"
    fi
fi

if command -v go vet &> /dev/null; then
    echo "运行 go vet..."
    go vet ./...
fi

echo
echo "=== 编译检查 ==="
echo "检查代码编译..."
go build -v .

echo
echo "=== 测试完成 ==="
echo "结束时间: $(date)"
echo
echo "测试总结:"
echo "✓ 单元测试完成"
echo "✓ 集成测试完成"
echo "✓ 性能测试完成"
echo "✓ 代码检查完成"
echo "✓ 编译检查完成"
echo
echo "多版本支持实现完成！"
echo
echo "主要功能:"
echo "- 自动版本检测"
echo "- 适配器模式实现"
echo "- Zabbix 6.0/7.0 兼容"
echo "- 新特性支持 (MFA, Proxy Group, Browser Item, History Push)"
echo "- 完整的测试覆盖"
echo
echo "使用方法请参考:"
echo "- MULTIVERSION_README.md"
echo "- example_multiversion.go"