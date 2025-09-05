// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package chat

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/components/model"

	kitlog "github.com/fsyyft-go/kit/log"

	appconf "github.com/fsyyft-ai/eino-wizard/internal/pkg/conf"
)

func CreateOllamaChatModel(ctx context.Context, logger kitlog.Logger, cfg *appconf.Config) model.ToolCallingChatModel {
	l := logger.WithField("chat", "ollama")
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: cfg.Ai.Ollama.BaseUrl, // Ollama 服务地址
		Model:   cfg.Ai.Ollama.Model,   // 模型名称
	})
	if err != nil {
		l.Fatalf("create ollama chat model failed: %v", err)
	}
	return chatModel
}
