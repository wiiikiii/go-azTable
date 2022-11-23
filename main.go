package main

import (
	"fmt"
	"os"

	connectAzStorage "go-table/pkg/connect"
	helper "go-table/pkg/helper"
	manipulateAzTable "go-table/pkg/manipulate"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type ExportStruct struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

func main() {

	var args = os.Args[1:]
	function := args[0]
	partitionKey := args[1]
	rowKey := args[2]
	tableName := args[3]

	valid := true
	for _, k := range args {
		if !helper.ValidateParams(k) {
			valid = false
			break
		}
	}

	if valid {

		switch {

		case function == "get":
			var client *aztables.Client
			connectAzStorage.ConnectStorageAccount(tableName)
			client, err := connectAzStorage.ConnectStorageAccount(tableName)
			if err != nil {
				panic(err)
			}

			res, err := manipulateAzTable.GetTableData(client, partitionKey, rowKey, tableName)
			if err != nil {
				panic(err)
			}
			fmt.Println(res)

		case function == "update":

			propertyName := args[4]
			propertyValue := args[5]

			if helper.ValidateParams(propertyName) && helper.ValidateParams(propertyValue) {
				var client *aztables.Client
				connectAzStorage.ConnectStorageAccount(tableName)
				client, err := connectAzStorage.ConnectStorageAccount(tableName)
				if err != nil {
					panic(err)
				}

				res, err := manipulateAzTable.UpdateTableProperties(client, partitionKey, rowKey, tableName, propertyName, propertyValue)
				if err != nil {
					panic(err)
				}
				fmt.Println(res)
			}

		case function == "delete":

			propertyName := args[4]

			if helper.ValidateParams(propertyName) {
				var client *aztables.Client
				connectAzStorage.ConnectStorageAccount(tableName)
				client, err := connectAzStorage.ConnectStorageAccount(tableName)
				if err != nil {
					panic(err)
				}

				manipulateAzTable.DeleteTableProperties(client, partitionKey, rowKey, tableName, propertyName)
				if err != nil {
					panic(err)
				}
			}

		case function == "single":

			propertyName := args[4]

			if helper.ValidateParams(propertyName) {
				var client *aztables.Client
				connectAzStorage.ConnectStorageAccount(tableName)
				client, err := connectAzStorage.ConnectStorageAccount(tableName)
				if err != nil {
					panic(err)
				}
				res, err := manipulateAzTable.GetSingleTableValue(client, partitionKey, rowKey, tableName, propertyName)
				if err != nil {
					panic(err)
				}
				fmt.Println(res)
			}
		default:
			fmt.Printf("Unknown Parameter %q", function)
		}

	}

}
