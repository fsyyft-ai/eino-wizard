// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package chat

import (
	"testing"
)

func TestCreateMessagesFromTemplate(t *testing.T) {
	t.Run("TestCreateMessagesFromTemplate", func(t *testing.T) {
		// 创建一个测试用例。
		testCases := []struct {
			name     string
			template string
			input    string
			expected []string
		}{
			{
				name: "Basic",
			},
		}

		for _, tc := range testCases {

			messages := createMessagesFromTemplate()
			t.Logf("Testing %s：%v", tc.name, messages)
		}
	})
}
