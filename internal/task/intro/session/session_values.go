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

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"

	appmodel "github.com/fsyyft-ai/eino-wizard/pkg/eino/adk/common/model"
	appprints "github.com/fsyyft-ai/eino-wizard/pkg/eino/adk/common/prints"
)

type (
	ToolAInput struct {
		Name string `json:"input" jsonschema:"description=用户姓名"`
	}
	ToolBInput struct {
		Age int `json:"input" jsonschema:"description=用户年龄"`
	}
	ToolCInput struct {
		Sex string `json:"input" jsonschema:"description=用户性别"`
	}
)

func toolAFn(ctx context.Context, in *ToolAInput) (string, error) {
	adk.AddSessionValue(ctx, "user-name", in.Name)
	return in.Name, nil
}

func toolBFn(ctx context.Context, in *ToolBInput) (string, error) {
	adk.AddSessionValue(ctx, "user-age", in.Age)
	userName, _ := adk.GetSessionValue(ctx, "user-name")
	return fmt.Sprintf("用户名: %v, 年龄: %v", userName, in.Age), nil
}

func toolCFn(ctx context.Context, in *ToolCInput) (string, error) {
	adk.AddSessionValue(ctx, "user-sex", in.Sex)
	userName, _ := adk.GetSessionValue(ctx, "user-name")
	age, _ := adk.GetSessionValue(ctx, "user-age")
	return fmt.Sprintf("用户名: %v, 年龄: %v，性别：%v", userName, age, in.Sex), nil
}

func main() {
	ctx := context.Background()

	toolA, err := utils.InferTool("tool_a", "设置用户名", toolAFn)
	if err != nil {
		log.Fatalf("InferTool 失败, err: %v", err)
	}

	toolB, err := utils.InferTool("tool_b", "设置用户年龄", toolBFn)
	if err != nil {
		log.Fatalf("InferTool 失败, err: %v", err)
	}

	toolC, err := utils.InferTool("tool_c", "设置用户性别", toolCFn)
	if err != nil {
		log.Fatalf("InferTool 失败, err: %v", err)
	}

	a, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "ChatModelAgent",
		Description: "一个聊天模型代理",
		Instruction: "你是一个聊天模型代理，请按顺序调用 toolA toolB toolC；然后给出一个第一次见面时的交谈语。根据同的年龄、性别，说出合适的信息，让对方感受到亲切。",
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []tool.BaseTool{
					toolA,
					toolB,
					toolC,
				},
			},
		},
		Model: appmodel.NewChatModel(ctx),
	})
	if err != nil {
		log.Fatalf("NewChatModelAgent 创建失败, err: %v", err)
	}

	r := adk.NewRunner(ctx, adk.RunnerConfig{
		Agent: a,
	})

	iter := r.Query(ctx, "我叫老鬼，我 18 岁，我是老男人")
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}

		appprints.Event(event)
	}
}
