// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package task

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/model"

	kitlog "github.com/fsyyft-go/kit/log"

	appconf "github.com/fsyyft-ai/eino-wizard/internal/conf"
	appchat "github.com/fsyyft-ai/eino-wizard/internal/task/quickstart/chat"
)

type (
	// QuickStartChat 定义了 QuickStartChat 任务的接口。
	QuickStartChat interface {
		// Run 执行 QuickStartChat 任务。
		Run(ctx context.Context) error
	}

	// quickStartChat 实现了 QuickStartChat 接口。
	quickStartChat struct {
		// logger 用于记录任务执行过程中的日志信息。
		logger kitlog.Logger
		// cfg 存储应用配置信息。
		cfg *appconf.Config
	}
)

// NewQuickStartChat 创建一个新的 QuickStartChat 实例。
//
// 参数:
//   - logger: 用于记录日志的 logger 实例。
//   - cfg: 应用配置信息。
//
// 返回值:
//   - QuickStartChat: 一个新的 QuickStartChat 实例。
//   - error: 创建实例过程中可能发生的错误。
func NewQuickStartChat(logger kitlog.Logger, cfg *appconf.Config) (QuickStartChat, error) {
	return &quickStartChat{logger: logger, cfg: cfg}, nil
}

// Run 执行 QuickStartChat 任务。
//
// 参数:
//   - ctx: 上下文。
//
// 返回值:
//   - error: 执行过程中可能发生的错误。
func (h *quickStartChat) Run(ctx context.Context) error {
	fmt.Println("===create messages===")
	messages := appchat.CreateMessagesFromTemplate()
	fmt.Println("===create llm===")
	var cm model.ToolCallingChatModel
	if h.cfg.Ai.LocalTest {
		cm = appchat.CreateOllamaChatModel(ctx, h.logger, h.cfg)
	} else {
		cm = appchat.CreateOpenAIChatModel(ctx, h.logger, h.cfg)
	}
	fmt.Println("===llm stream generate===")
	streamResult := appchat.Stream(ctx, cm, messages)
	appchat.ReportStream(streamResult)
	fmt.Println()
	fmt.Println("===llm stream end===")
	fmt.Println("===llm generate===")
	result := appchat.Generate(ctx, cm, messages)
	fmt.Println(result)

	return ctx.Err()
}
