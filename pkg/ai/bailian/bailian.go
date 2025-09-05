// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package bailian

// 百炼控制台：https://bailian.console.aliyun.com
// 所有模型列表与价格：https://help.aliyun.com/zh/model-studio/models

const (
	OpenAIURLBailian = "https://dashscope.aliyuncs.com/compatible-mode/v1"

	OpenAIModelBailianQwenTurboLatest = "qwen-turbo-latest" // 文本生成。
	OpenAIModelBailianQwenPlusLatest  = "qwen-plus-latest"  // 文本生成。
	OpenAIModelBailianQwenMaxLatest   = "qwen-max-latest"   // 文本生成。

	OpenAIModelBailianTextEmbeddingV4       = "text-embedding-v4"       // 文本向量。
	OpenAIModelBailianMultimodalEmbeddingV1 = "multimodal-embedding-v1" // 多模态向量。

	OpenAIModelBailianWan22T2iFlash = "wan2.2-t2i-flash " // 文生图。
	OpenAIModelBailianWan22T2iPlus  = "wan2.2-t2i-plus"   // 文生图。

	OpenAIModelBailianQwen3CoderPlus = "qwen3-coder-plus" // 代码模型。
)
