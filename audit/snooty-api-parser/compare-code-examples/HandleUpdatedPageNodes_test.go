package compare_code_examples

import (
	"snooty-api-parser/compare-code-examples/data"
	"snooty-api-parser/snooty"
	"snooty-api-parser/types"
	"strings"
	"testing"
	"time"
)

func TestHandleUpdatedPageNodesCorrectlyUpdatesValues(t *testing.T) {
	codeNode, astNode := data.GetUpdatedNodes()
	sha256HashCodeNodeLookupMap := make(map[string]types.CodeNode)
	whitespaceTrimmedString := strings.TrimSpace(astNode.Value)
	incomingSha26Hash := snooty.MakeSha256HashForCode(whitespaceTrimmedString)
	sha256HashCodeNodeLookupMap[incomingSha26Hash] = codeNode
	updatedCodeNodes := HandleUpdatedPageNodes([]types.ASTNode{astNode}, sha256HashCodeNodeLookupMap)
	updatedCodeNode := updatedCodeNodes[0]
	if updatedCodeNode.SHA256Hash != incomingSha26Hash {
		t.Errorf("FAILED: got %s on the code node hash, want %s", updatedCodeNode.SHA256Hash, incomingSha26Hash)
	}
	if updatedCodeNode.Code != whitespaceTrimmedString {
		t.Errorf("FAILED: got %s in the updated Code text, want %s", updatedCodeNode.Code, whitespaceTrimmedString)
	}
	tolerance := 2 * time.Second // Define tolerance of 2 seconds
	if !IsApproximatelyNow(updatedCodeNode.DateUpdated, tolerance) {
		t.Errorf("FAILED: updated node time is not approximately now")
	}
}
