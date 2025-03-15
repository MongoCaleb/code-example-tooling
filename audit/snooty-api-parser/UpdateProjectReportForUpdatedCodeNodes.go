package main

import "snooty-api-parser/types"

func IncrementProjectCountsForExistingPage(incomingCodeNodeCount int, incomingLiteralIncludeNodeCount int, incomingIoCodeBlockNodeCount int, existingPage types.DocsPage, report types.ProjectReport) types.ProjectReport {
	report.Counter.IncomingCodeNodesCount += incomingCodeNodeCount
	report.Counter.IncomingLiteralIncludeCount += incomingLiteralIncludeNodeCount
	report.Counter.IncomingIoCodeBlockCount += incomingIoCodeBlockNodeCount
	report.Counter.ExistingCodeNodesCount += existingPage.CodeNodesTotal
	report.Counter.ExistingLiteralIncludeCount += existingPage.LiteralIncludesTotal
	report.Counter.ExistingIoCodeBlockCount += existingPage.IoCodeBlocksTotal
	return report
}
