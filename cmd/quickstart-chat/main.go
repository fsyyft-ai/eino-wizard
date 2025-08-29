// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package main

import (
	appchat "github.com/fsyyft-ai/eino-wizard/internal/app/quickstart/chat"
)

func main() {
	// 应用程序入口。
	// 测试过在某些情况下，使用 wire 生成代码时，会报错，可能是因为这时 main 包的原因，所以这里只包含入口。
	appchat.Run()
}
