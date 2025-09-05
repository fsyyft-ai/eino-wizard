// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package todoagent

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"

	kitlog "github.com/fsyyft-go/kit/log"

	appconf "github.com/fsyyft-ai/eino-wizard/internal/conf"
)

type (
	TodoAgent interface {
		BaseTools() []tool.BaseTool
		ToolInfos(ctx context.Context) []*schema.ToolInfo
	}

	todoAgent struct {
		// logger 用于记录任务执行过程中的日志信息。
		logger kitlog.Logger
		// cfg 存储应用配置信息。
		cfg *appconf.Config

		tools     []tool.BaseTool
		toolInfos []*schema.ToolInfo
	}

	TodoAddParams struct {
		Content  string `json:"content"`
		StartAt  *int64 `json:"started_at,omitempty"` // 开始时间
		Deadline *int64 `json:"deadline,omitempty"`
	}
	TodoUpdateParams struct {
		ID        string  `json:"id" jsonschema:"description=id of the todo"`
		Content   *string `json:"content,omitempty" jsonschema:"description=content of the todo"`
		StartedAt *int64  `json:"started_at,omitempty" jsonschema:"description=start time in unix timestamp"`
		Deadline  *int64  `json:"deadline,omitempty" jsonschema:"description=deadline of the todo in unix timestamp"`
		Done      *bool   `json:"done,omitempty" jsonschema:"description=done status"`
	}
)

func NewTodoAgent(logger kitlog.Logger, cfg *appconf.Config) (TodoAgent, func(), error) {
	return &todoAgent{
		logger: logger,
		cfg:    cfg,
	}, func() {}, nil
}

func (a *todoAgent) ToolInfos(ctx context.Context) []*schema.ToolInfo {
	if nil == a.toolInfos {
		bts := a.BaseTools()
		toolInfos := make([]*schema.ToolInfo, 0, len(bts))
		for _, tool := range bts {
			info, err := tool.Info(ctx)
			if err != nil {
				a.logger.Fatal(err)
			}
			toolInfos = append(toolInfos, info)
		}
		a.toolInfos = toolInfos
	}
	return a.toolInfos
}

func (a *todoAgent) BaseTools() []tool.BaseTool {
	if nil == a.tools {
		a.tools = []tool.BaseTool{
			a.getAddTodoTool(),
			a.getUpdateTodoTool(),
		}
	}
	return a.tools
}

// -----------------------------------------------------------------------------
// 方式一：使用 NewTool 构建。
// -----------------------------------------------------------------------------
func (a *todoAgent) getAddTodoTool() tool.InvokableTool {
	// 工具信息
	info := &schema.ToolInfo{
		Name: "add_todo",
		Desc: "Add a todo item",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"content": {
				Desc:     "The content of the todo item",
				Type:     schema.String,
				Required: true,
			},
			"started_at": {
				Desc: "The started time of the todo item, in unix timestamp",
				Type: schema.Integer,
			},
			"deadline": {
				Desc: "The deadline of the todo item, in unix timestamp",
				Type: schema.Integer,
			},
		}),
	}

	// i utils.InvokeFunc[*TodoAddParams, string]
	return utils.NewTool(info, a.addTodoFunc)
}

func (a *todoAgent) addTodoFunc(_ context.Context, params *TodoAddParams) (string, error) {
	a.logger.Infof("invoke tool add_todo: %+v", params)
	return `{"msg": "add todo success"}`, nil
}

// -----------------------------------------------------------------------------
// 方式二：使用 InferTool 构建。
// -----------------------------------------------------------------------------

func (a *todoAgent) getUpdateTodoTool() tool.InvokableTool {
	updateTool, err := utils.InferTool("update_todo", "Update a todo item, eg: content,deadline...", a.UpdateTodoFunc)
	if err != nil {
		a.logger.Errorf("InferTool failed, err=%v", err)
		return nil
	}
	a.logger.Info("update_todo tool inferred successfully")
	return updateTool
}
func (a *todoAgent) UpdateTodoFunc(_ context.Context, params *TodoUpdateParams) (string, error) {
	a.logger.Infof("invoke tool update_todo: %+v", params)
	// Mock 处理逻辑。
	return `{"msg": "update todo success"}`, nil
}
