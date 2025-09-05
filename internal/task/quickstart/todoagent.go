// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package quickstart

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"

	kitlog "github.com/fsyyft-go/kit/log"

	appconf "github.com/fsyyft-ai/eino-wizard/internal/pkg/conf"
	apptodoagent "github.com/fsyyft-ai/eino-wizard/internal/task/quickstart/todoagent"
	appbailian "github.com/fsyyft-ai/eino-wizard/pkg/ai/bailian"
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
	todoTools := h.agent.ToolInfos(ctx)
	toolsNodeConfig := &compose.ToolsNodeConfig{
		Tools: h.agent.BaseTools(ctx),
	}
	chatModelConfig := &openai.ChatModelConfig{
		BaseURL: appbailian.OpenAIURLBailian,
		Model:   appbailian.OpenAIModelBailianQwenPlusLatest,
		APIKey:  h.cfg.Ai.Openai.ApiKey,
	}

	chatModel, err := openai.NewChatModel(context.Background(), chatModelConfig)
	if nil != err {
		return err
	}

	err = chatModel.BindTools(todoTools)
	if nil != err {
		return err
	}

	todoToolsNode, err := compose.NewToolNode(ctx, toolsNodeConfig)
	if nil != err {
		return err
	}

	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))

	agent, err := chain.Compile(ctx)
	if nil != err {
		return err
	}

	taskMessages := []string{
		"添加一个学习 Eino 的 TODO，同时搜索一下 fsyyft-ai/eino-wizard 的仓库地址",
		"获取当前所有的 TODO 列表",
		"搜索一下 fsyyft-ai/eino-wizard 的仓库地址",
	}
	for _, task := range taskMessages {
		if err := h.invokeMessage(ctx, agent, task); nil != err {
			h.logger.Error(ctx, "invoke message failed", "task", task, "error", err)
			return err
		}
	}

	return nil
}

func (h *todoAgent) invokeMessage(ctx context.Context, agent compose.Runnable[[]*schema.Message, []*schema.Message], in string) error {
	msg := &schema.Message{
		Role:    schema.User,
		Content: in,
	}
	return h.invoke(ctx, agent, msg)
}

func (h *todoAgent) invoke(ctx context.Context, agent compose.Runnable[[]*schema.Message, []*schema.Message], in *schema.Message) error {
	resp, err := agent.Invoke(ctx, []*schema.Message{in})
	if err != nil {
		return err
	}

	for _, msg := range resp {
		fmt.Println(msg.Content)
	}

	return nil
}
