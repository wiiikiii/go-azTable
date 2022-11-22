package main

import (
	"os"

	connectAzStorage "go-table/pkg/connect"
	manipulateAzTable "go-table/pkg/manipulate"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type ExportStruct struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

func main() {

	var args = os.Args[1:]

	partitionKey := args[0]
	rowKey := args[1]
	tableName := args[2]
	propertyName := args[3]
	// propertyValue := args[4]
	var client *aztables.Client

	connectAzStorage.ConnectStorageAccount(tableName)

	client, err := connectAzStorage.ConnectStorageAccount(tableName)
	if err != nil {
		panic(err)
	}

	// var res string
	// res, err = manipulateAzTable.GetTableData(client, partitionKey, rowKey, tableName)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)

	//manipulateAzTable.GetSingleTableValue(client, partitionKey, rowKey, tableName, valueToQuery)

	// r, err  := manipulateAzTable.UpdateTableProperties(client, partitionKey, rowKey, tableName, propertyName, propertyValue)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(r)

	manipulateAzTable.DeleteTableProperties(client, partitionKey, rowKey, tableName, propertyName)
	if err != nil {
		panic(err)
	}
}