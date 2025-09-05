// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package task

import (
	"context"

	kitlog "github.com/fsyyft-go/kit/log"

	appconf "github.com/fsyyft-ai/eino-wizard/internal/conf"
	appquickstart "github.com/fsyyft-ai/eino-wizard/internal/task/quickstart"
)

type (
	// QuickStart 定义了 QuickStart 任务的接口。
	QuickStart interface {
		// Run 执行 QuickStart 任务。
		Run(ctx context.Context) error
	}

	// quickStart 实现了 QuickStart 接口。
	quickStart struct {
		// logger 用于记录任务执行过程中的日志信息。
		logger kitlog.Logger
		// cfg 存储应用配置信息。
		cfg *appconf.Config

		chat appquickstart.Chat
	}
)

// NewQuickStart 创建一个新的 QuickStart 实例。
//
// 参数:
//   - logger: 用于记录日志的 logger 实例。
//   - cfg: 应用配置信息。
//
// 返回值:
//   - QuickStart: 一个新的 QuickStart 实例。
//   - error: 创建实例过程中可能发生的错误。
func NewQuickStart(logger kitlog.Logger, cfg *appconf.Config, chat appquickstart.Chat) (QuickStart, error) {
	return &quickStart{logger: logger, cfg: cfg, chat: chat}, nil
}

// Run 执行 QuickStart 任务。
//
// 参数:
//   - ctx: 上下文。
//
// 返回值:
//   - error: 执行过程中可能发生的错误。
func (h *quickStart) Run(ctx context.Context) error {
	err := h.chat.Run(ctx)
	return err
}
