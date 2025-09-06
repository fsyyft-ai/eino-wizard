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

package subagents

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"

	appmodel "github.com/fsyyft-ai/eino-wizard/pkg/eino/adk/common/model"
)

type (
	GetWeatherInput struct {
		City string `json:"city"`
	}
)

func NewWeatherAgent(ctx context.Context) adk.Agent {
	weatherTool, err := utils.InferTool(
		"get_weather",
		"获取指定城市的当前天气。",
		func(ctx context.Context, input *GetWeatherInput) (string, error) {
			return fmt.Sprintf(`当前 %s 的气温是 25°C`, input.City), nil
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	a, err := adk.NewChatModelAgent(context.Background(), &adk.ChatModelAgentConfig{
		Name:        "WeatherAgent",
		Description: "该 Agent 可获取指定城市的当前天气。",
		Instruction: `你的唯一目的，是使用 'get_weather' 工具获取指定城市的当前天气。
			调用工具后，直接将结果告知用户。`,
		Model: appmodel.NewChatModel(ctx),
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []tool.BaseTool{weatherTool},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func NewChatAgent(ctx context.Context) adk.Agent {
	a, err := adk.NewChatModelAgent(context.Background(), &adk.ChatModelAgentConfig{
		Name:        "ChatAgent",
		Description: "用于处理一般会话聊天的通用 Agent。",
		Instruction: `你是一个友好的对话助手。
			你的职责是处理一般性的闲聊，并回答与特定工具任务无关的问题。`,
		Model: appmodel.NewChatModel(ctx),
	})
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func NewRouterAgent(ctx context.Context) adk.Agent {
	a, err := adk.NewChatModelAgent(context.Background(), &adk.ChatModelAgentConfig{
		Name:        "RouterAgent",
		Description: "一个将任务转交给其他专家 Agent 的手动路由器。",
		Instruction: `你是一个智能任务路由器。
			你的责任是分析用户请求，并将其委派给最合适的专家 Agent。
			如果没有任何 Agent 可以处理该任务，就直接告知用户无法处理。`,
		Model: appmodel.NewChatModel(ctx),
	})
	if err != nil {
		log.Fatal(err)
	}
	return a
}
