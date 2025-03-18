package add_code_examples

import "snooty-api-parser/types"

func IsNewAppliedUsageExample(node types.CodeNode) bool {
	codeExampleCharacterCount := len([]rune(node.Code))
	if node.Category == UsageExample && codeExampleCharacterCount > 300 {
		return true
	}
	return false
}
