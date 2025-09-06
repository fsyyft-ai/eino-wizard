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
	"log"

	"github.com/cloudwego/eino/adk"

	appsubagents "github.com/fsyyft-ai/eino-wizard/internal/task/intro/transfer/subagents"
	appprints "github.com/fsyyft-ai/eino-wizard/pkg/eino/adk/common/prints"
)

func main() {
	ctx := context.Background()

	weatherAgent := appsubagents.NewWeatherAgent(ctx)
	chatAgent := appsubagents.NewChatAgent(ctx)
	routerAgent := appsubagents.NewRouterAgent(ctx)

	a, err := adk.SetSubAgents(ctx, routerAgent, []adk.Agent{chatAgent, weatherAgent})
	if err != nil {
		log.Fatal(err)
	}

	runner := adk.NewRunner(ctx, adk.RunnerConfig{
		EnableStreaming: true,
		Agent:           a,
	})

	println("\n\n>>>>>>>>> 查询天气 <<<<<<<<<")
	iter := runner.Query(ctx, "杭州的天气如何？")
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			log.Fatal(event.Err)
		}

		appprints.Event(event)
	}

	println("\n\n>>>>>>>>> 路由失败 <<<<<<<<<")
	iter = runner.Query(ctx, "帮我订一张明天从纽约飞往伦敦的机票。")
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			log.Fatal(event.Err)
		}
		appprints.Event(event)
	}
}
