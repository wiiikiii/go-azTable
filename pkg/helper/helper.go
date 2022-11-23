package helper

import (
	"fmt"
	"net/http"
	"reflect"

	connectAzStorage "go-table/pkg/connect"
	manipulateAzTable "go-table/pkg/manipulate"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func ValidateParams(value string) bool {
	res := reflect.ValueOf(value).IsValid()
	return res
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	tableName := "CorpTest"
	partitionKey := "AVD"
	rowKey := "DEPLOYMENT-IN"
	propertyName := ""
	propertyValue := ""
	
	var err error
	var client *aztables.Client
	client, err = connectAzStorage.ConnectStorageAccount(tableName)
	if err != nil {
		panic(err)
	}

	var message string
	name := r.URL.Query().Get("function")
	if name == "get" {
		message, err = manipulateAzTable.GetTableData(client, partitionKey, rowKey, tableName)
	}
	if name == "update" {
		message, err = manipulateAzTable.UpdateTableProperties(client, partitionKey, rowKey, tableName, propertyName, propertyValue)
	}
	if name == "single" {
		message, err = manipulateAzTable.GetSingleTableValue(client, partitionKey, rowKey, tableName, propertyName)
	}
	if name == "delete" {
		message, err = manipulateAzTable.DeleteTableProperties(client, partitionKey, rowKey, tableName, propertyName)
	}

	if err != nil{
		panic(err)
	}
	fmt.Fprint(w, message)
}
