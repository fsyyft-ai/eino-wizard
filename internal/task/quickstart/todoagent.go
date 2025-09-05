// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package quickstart

import (
	"context"
	"fmt"

	kitlog "github.com/fsyyft-go/kit/log"

	appconf "github.com/fsyyft-ai/eino-wizard/internal/conf"
	apptodoagent "github.com/fsyyft-ai/eino-wizard/internal/task/quickstart/todoagent"
)

type (
	// TodoAgent 定义了 TodoAgent 任务的接口。
	TodoAgent interface {
		// Run 执行 TodoAgent 任务。
		Run(ctx context.Context) error
	}

	// todoAgent 实现了 TodoAgent 接口。
	todoAgent struct {
		// logger 用于记录任务执行过程中的日志信息。
		logger kitlog.Logger
		// cfg 存储应用配置信息。
		cfg *appconf.Config

		agent apptodoagent.TodoAgent
	}
)

// NewTodoAgent 创建一个新的 TodoAgent 实例。
//
// 参数:
//   - logger: 用于记录日志的 logger 实例。
//   - cfg: 应用配置信息。
//
// 返回值:
//   - TodoAgent: 一个新的 TodoAgent 实例。
//   - func()：清理函数，用于在初始化失败时进行资源释放。
//   - error: 创建实例过程中可能发生的错误。
func NewTodoAgent(logger kitlog.Logger, cfg *appconf.Config, agent apptodoagent.TodoAgent) (TodoAgent, func(), error) {
	return &todoAgent{logger: logger, cfg: cfg, agent: agent}, func() {}, nil
}

// Run 执行 TodoAgent 任务。
//
// 参数:
//   - ctx: 上下文。
//
// 返回值:
//   - error: 执行过程中可能发生的错误。
func (h *todoAgent) Run(ctx context.Context) error {
	todoTools := h.agent.Tools()
	fmt.Println(len(todoTools))
	return nil
}
