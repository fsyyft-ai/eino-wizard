// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package chat

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"

	kitlog "github.com/fsyyft-go/kit/log"

	appconf "github.com/fsyyft-ai/eino-wizard/internal/pkg/conf"
	appbailian "github.com/fsyyft-ai/eino-wizard/pkg/ai/bailian"
)

func CreateOpenAIChatModel(ctx context.Context, logger kitlog.Logger, cfg *appconf.Config) model.ToolCallingChatModel {
	l := logger.WithField("chat", "openai")
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: appbailian.OpenAIURLBailian,
		Model:   appbailian.OpenAIModelBailianQwenPlusLatest,
		APIKey:  cfg.Ai.Openai.ApiKey,
	})
	if err != nil {
		l.Fatalf("create openai chat model failed, err=%v", err)
	}
	return chatModel
}
