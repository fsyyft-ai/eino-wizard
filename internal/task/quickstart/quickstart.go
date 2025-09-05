// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package quickstart

import (
	"github.com/google/wire"

	apptodoagent "github.com/fsyyft-ai/eino-wizard/internal/task/quickstart/todoagent"
)

var (
	// ProviderSet 是 wire 的依赖注入提供者集合。
	// 包含了创建任务实例所需的所有依赖。
	ProviderSet = wire.NewSet(
		NewChat,
		NewTodoAgent,
		apptodoagent.NewTodoAgent,
	)
)
