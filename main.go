package main

import (
	"fmt"
	manipulateTableData "go-table/pkg"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type ExportStruct struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

func main(){
	partitionKey := os.Args[1]
	rowKey := os.Args[2]
	tableName := os.Args[3]
	var client *aztables.Client

	//valueToQuery := "FRONT_ApplicationsToInstall"

	client, err := manipulateTableData.ConnectStorageAccount(tableName)
	if(err != nil){
		panic(err)
	}

	export := manipulateTableData.GetTableData(client, partitionKey, rowKey, tableName)
	//manipulateTableData.GetSingleTableValue(client, partitionKey, rowKey, tableName, valueToQuery)

	fmt.Println(export)
}
