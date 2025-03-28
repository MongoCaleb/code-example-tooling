package main

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"pull-audit-data/aggregations"
	"pull-audit-data/utils"
)

// PerformAggregation executes several different aggregation operations for every collection in the DB, and logs the output to console.
func PerformAggregation(db *mongo.Database, ctx context.Context) {
	// The aggregations in this project use one of these data structures. Uncomment the corresponding data structure,
	// or make duplicates with appropriate names as needed
	simpleMap := make(map[string]int)
	//codeLengthMap := make(map[string]types.CodeLengthStats)
	//nestedOneLevelMap := make(map[string]map[string]int)
	//nestedTwoLevelMap := make(map[string]map[string]map[string]int)

	// If you just need to get data for a single collection, perform the aggregation using the collection name
	//simpleMap = aggregations.GetLanguageCounts(db, "pymongo", simpleMap, ctx)

	// If you need to get data across all the collections in the `code-examples` database, iterate through the collections
	emptyFilter := bson.D{}
	collectionNames, err := db.ListCollectionNames(ctx, emptyFilter)
	if err != nil {
		panic(err)
	}

	for _, collectionName := range collectionNames {
		//simpleMap = aggregations.GetCategoryCounts(db, collectionName, simpleMap, ctx)
		//simpleMap = aggregations.GetLanguageCounts(db, collectionName, simpleMap, ctx)
		//codeLengthMap = aggregations.GetCodeLengths(db, collectionName, codeLengthMap, ctx)
		//nestedOneLevelMap = aggregations.GetCategoryLanguageCounts(db, collectionName, nestedOneLevelMap, ctx)
		//nestedOneLevelMap = aggregations.GetProductCategoryCounts(db, collectionName, nestedOneLevelMap, ctx)
		//nestedTwoLevelMap = aggregations.GetSubProductCategoryCounts(db, collectionName, nestedTwoLevelMap, ctx)
		simpleMap = aggregations.GetOneLineUsageExampleCounts(db, collectionName, simpleMap, ctx)
		//nestedOneLevelMap = aggregations.GetProductLanguageCounts(db, collectionName, nestedOneLevelMap, ctx)
		//nestedTwoLevelMap = aggregations.GetSubProductLanguageCounts(db, collectionName, nestedTwoLevelMap, ctx)
		//simpleMap = aggregations.GetCollectionCount(db, collectionName, simpleMap, ctx)
		//simpleMap = aggregations.GetSpecificCategoryByProduct(db, collectionName, types.UsageExample, simpleMap, ctx)
		//langCount := aggregations.GetSpecificLanguageCount(db, collectionName, "go", ctx)
	}

	simpleTableLabel := "Category"
	simpleTableColumnNames := []interface{}{"Category", "Count"}
	simpleTableColumnWidths := []int{30, 15}
	utils.PrintSimpleCountDataToConsole(simpleMap, simpleTableLabel, simpleTableColumnNames, simpleTableColumnWidths)

	//nestedOneLevelTableLabel := "Product Language"
	//nestedOneLevelTableColumnNames := []interface{}{"Language", "Count"}
	//nestedOneLevelTableColumnWidths := []int{20, 15}
	//utils.PrintNestedOneLevelCountDataToConsole(nestedOneLevelMap, nestedOneLevelTableLabel, nestedOneLevelTableColumnNames, nestedOneLevelTableColumnWidths)
	//
	//nestedTwoLevelTableLabel := "Sub-Product Language"
	//nestedTwoLevelTableColumnNames := []interface{}{"Language", "Count"}
	//nestedTwoLevelTableColumnWidths := []int{20, 15}
	//utils.PrintNestedTwoLevelCountDataToConsole(nestedTwoLevelMap, nestedTwoLevelTableLabel, nestedTwoLevelTableColumnNames, nestedTwoLevelTableColumnWidths)

	// The length count map is a very specific fixed data structure, so this function has hard-coded title and column names/widths
	//utils.PrintCodeLengthMapToConsole(codeLengthMap)
}
