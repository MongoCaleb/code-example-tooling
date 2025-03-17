package snooty

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
	"snooty-api-parser/add-code-examples"
	"snooty-api-parser/types"
	"strings"
	"time"
)

func MakeCodeNodeFromSnootyAST(snootyNode types.ASTNode, llm *ollama.LLM, ctx context.Context, isDriverProject bool) types.CodeNode {
	whiteSpaceTrimmedCode := strings.TrimSpace(snootyNode.Value)
	hashString := MakeSha256HashForCode(whiteSpaceTrimmedCode)
	language := add_code_examples.GetNormalizedLanguageFromASTNode(snootyNode)
	fileExtension := add_code_examples.GetFileExtensionFromASTNode(snootyNode)
	category, llmCategorized := add_code_examples.GetCategory(whiteSpaceTrimmedCode, language, llm, ctx, isDriverProject)
	return types.CodeNode{
		Code:           whiteSpaceTrimmedCode,
		Language:       language,
		FileExtension:  fileExtension,
		Category:       category,
		SHA256Hash:     hashString,
		LLMCategorized: llmCategorized,
		DateAdded:      time.Now(),
	}
}
