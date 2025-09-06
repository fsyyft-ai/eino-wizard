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
	"log"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type BookSearchInput struct {
	Genre     string `json:"genre" jsonschema:"description=偏好的图书类型,enum=fiction,enum=sci-fi,enum=mystery,enum=biography,enum=business"`
	MaxPages  int    `json:"max_pages" jsonschema:"description=最大页数（0 表示不限制）"`
	MinRating int    `json:"min_rating" jsonschema:"description=最低用户评分（0-5）"`
}

type BookSearchOutput struct {
	Books []string
}

func NewBookRecommender() tool.InvokableTool {
	bookSearchTool, err := utils.InferTool("search_book", "根据用户偏好搜索图书",
		func(ctx context.Context, input *BookSearchInput) (output *BookSearchOutput, err error) {
			// 搜索代码
			// ...
			return &BookSearchOutput{Books: []string{"为美好的世界献上祝福！", "AI 开发指南"}}, nil
		},
	)
	if err != nil {
		log.Fatalf("创建搜索图书工具失败: %v", err)
	}
	return bookSearchTool
}
