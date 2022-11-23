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
	if err != nil{
		panic(err)
	}
	fmt.Fprint(w, message)
}
