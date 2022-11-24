package main

import (
	"fmt"
	"log"
	"net/http"
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

var functions = []string{"server", "get", "update", "delete", "single"}

func main() {

	var args = os.Args[1:]
	function := args[0]

	if helper.Contains(functions, function) {

		if function == "server" {
			listenAddr := ":8080"
			if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
				listenAddr = ":" + val
			}
			http.HandleFunc("/api/HttpExample", helper.HelloHandler)
			log.Printf("Server started.\n")
			log.Printf("About to listen on Port%s.\nGo to https://127.0.0.1%s/", listenAddr, listenAddr)
			log.Fatal(http.ListenAndServe(listenAddr, nil))

		} else {

			valid := true

			for _, k := range args {
				if !helper.ValidateParams(k) {
					valid = false
					break
				}
			}

			partitionKey := args[1]
			rowKey := args[2]
			tableName := args[3]

			if valid {

				switch {

				case function == "get":

					if len(args) == 4 {
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

					} else {
						fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey and tablename")
						break
					}

				case function == "update":

					if len(args) == 6 {

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

					} else {
						fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey, tablename, propertyName and propertyValue")
						break
					}

				case function == "delete":

					if len(args) == 5 {

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
							return
						}

					} else {
						fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey, tablename and propertyName")
						break

					}

				case function == "single":

					if len(args) == 5 {

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

					} else {
						fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey, tablename and propertyName")
						break
					}

				default:
					fmt.Printf("Unknown Parameter %q", function)
				}
			}
		}
	} else {
		fmt.Printf("%v is not a supported function, choose as first parameter from:\n %q", function, functions)
		return
	}

}
