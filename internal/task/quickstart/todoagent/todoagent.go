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
		}
	}
	return a.tools
}

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
