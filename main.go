package main

import (
	manipulateTableData "go-table/pkg"
	"os"
)

func main(){
	partitionKey := os.Args[1]
	rowKey := os.Args[2]
	tableName := os.Args[3]

	manipulateTableData.GetTableData(partitionKey, rowKey, tableName)
}
