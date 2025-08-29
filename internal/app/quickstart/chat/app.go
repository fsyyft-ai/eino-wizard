// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package chat

import (
	"context"
	"fmt"
)

func Run() {
	ctx := context.Background()
	fmt.Println("===create messages===")
	messages := createMessagesFromTemplate()
	fmt.Println("===create llm===")
	cm := createOllamaChatModel(ctx)
	fmt.Println("===llm stream generate===")
	streamResult := stream(ctx, cm, messages)
	reportStream(streamResult)
	fmt.Println()
	fmt.Println("===llm stream end===")
	fmt.Println("===llm generate===")
	result := generate(ctx, cm, messages)
	fmt.Println(result)
}
