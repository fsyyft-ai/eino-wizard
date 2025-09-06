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
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"

	appsubagents "github.com/fsyyft-ai/eino-wizard/internal/task/intro/chatmodel/subagents"
	appprints "github.com/fsyyft-ai/eino-wizard/pkg/eino/adk/common/prints"
)

func main() {
	checkPointID := "1"
	checkPointStore := newInMemoryStore()

	ctx := context.Background()
	a := appsubagents.NewBookRecommendAgent()
	runnerConfig := adk.RunnerConfig{
		EnableStreaming: true, // 你可以在这里关闭流式输出。
		Agent:           a,
		CheckPointStore: checkPointStore,
	}

	runner := adk.NewRunner(ctx, runnerConfig)
	// 设置了状态持久化 (Checkpoint)，Runner 捕获到这个带有 Interrupted Action 的 Event 时，会立即终止当前的执行流程。
	iter := runner.Query(ctx, "给我推荐一本书", adk.WithCheckPointID(checkPointID))
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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n请输入内容: ")
	scanner.Scan()
	fmt.Println()
	nInput := scanner.Text()

	toolOptions := []tool.Option{
		appsubagents.WithNewInput(nInput),
	}

	// 运行中断，调用 Runner 的 Resume 接口传入中断时的 CheckPointID 可以恢复运行。
	iter, err := runner.Resume(ctx, checkPointID, adk.WithToolOptions(toolOptions))
	if err != nil {
		log.Fatal(err)
	}
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

func newInMemoryStore() compose.CheckPointStore {
	return &inMemoryStore{
		mem: map[string][]byte{},
	}
}

type inMemoryStore struct {
	mem map[string][]byte
}

func (i *inMemoryStore) Set(ctx context.Context, key string, value []byte) error {
	i.mem[key] = value
	return nil
}

func (i *inMemoryStore) Get(ctx context.Context, key string) ([]byte, bool, error) {
	v, ok := i.mem[key]
	return v, ok, nil
}
