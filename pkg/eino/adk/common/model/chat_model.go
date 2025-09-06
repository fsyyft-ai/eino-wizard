/*
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"context"
	"os"
	"strings"
	"sync"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	arkModel "github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

type ChatModel struct {
	modelType     string
	arkAPIKey     string
	arkModel      string
	arkBaseURL    string
	openaiAPIKey  string
	openaiModel   string
	openaiBaseURL string
	openaiByAzure string
	once          sync.Once
	inner         model.ToolCallingChatModel
}

func NewChatModel(ctx context.Context, options ...ChatModelOption) model.ToolCallingChatModel {
	cm := &ChatModel{
		modelType:     os.Getenv("MODEL_TYPE"),
		arkAPIKey:     os.Getenv("ARK_API_KEY"),
		arkModel:      os.Getenv("ARK_MODEL"),
		arkBaseURL:    os.Getenv("ARK_BASE_URL"),
		openaiAPIKey:  os.Getenv("OPENAI_API_KEY"),
		openaiModel:   os.Getenv("OPENAI_MODEL"),
		openaiBaseURL: os.Getenv("OPENAI_BASE_URL"),
		openaiByAzure: os.Getenv("OPENAI_BY_AZURE"),
	}

	for _, opt := range options {
		opt(cm)
	}

	cm.init(ctx)

	return cm.inner
}

// Functional options (instead of methods) for configuring ChatModel before init.
type ChatModelOption func(*ChatModel)

func WithModelType(val string) ChatModelOption    { return func(c *ChatModel) { c.modelType = val } }
func WithArkAPIKey(val string) ChatModelOption    { return func(c *ChatModel) { c.arkAPIKey = val } }
func WithArkModel(val string) ChatModelOption     { return func(c *ChatModel) { c.arkModel = val } }
func WithArkBaseURL(val string) ChatModelOption   { return func(c *ChatModel) { c.arkBaseURL = val } }
func WithOpenAIAPIKey(val string) ChatModelOption { return func(c *ChatModel) { c.openaiAPIKey = val } }
func WithOpenAIModel(val string) ChatModelOption  { return func(c *ChatModel) { c.openaiModel = val } }
func WithOpenAIBaseURL(val string) ChatModelOption {
	return func(c *ChatModel) { c.openaiBaseURL = val }
}
func WithOpenAIByAzure(val string) ChatModelOption {
	return func(c *ChatModel) { c.openaiByAzure = val }
}

func (c *ChatModel) init(ctx context.Context) {
	c.once.Do(func() {
		modelType := strings.ToLower(c.modelType)
		if modelType == "ark" {
			cm, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
				APIKey:  c.arkAPIKey,
				Model:   c.arkModel,
				BaseURL: c.arkBaseURL,
				Thinking: &arkModel.Thinking{
					Type: arkModel.ThinkingTypeDisabled,
				},
			})
			if err == nil {
				c.inner = cm
			}
		} else {
			cm, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
				APIKey:  c.openaiAPIKey,
				Model:   c.openaiModel,
				BaseURL: c.openaiBaseURL,
				ByAzure: func() bool {
					return c.openaiByAzure == "true"
				}(),
			})
			if err == nil {
				c.inner = cm
			}
		}
	})
}
