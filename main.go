package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	m "go-table/pkg/manipulate"
)

type ExportStruct struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

func main() {

	var functions = []string{"server", "get", "update", "delete", "single"}

	env := (m.ReturnEnv([]string{
		"TABLES_STORAGE_ACCOUNT_NAME",
		"TABLES_PRIMARY_STORAGE_ACCOUNT_KEY",
		"TABLE_NAME"}))

	t := m.Table{
		Functions:   functions,
		AccountName: env["TABLES_STORAGE_ACCOUNT_NAME"],
		AccountKey:  env["TABLES_PRIMARY_STORAGE_ACCOUNT_KEY"],
		TableName:   env["TABLE_NAME"],
	}

	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getRowKey := getCmd.String("rowKey", "", "rowKey")
	getPartitionKey := getCmd.String("partitionKey", "", "partitionKey")
	getStage := getCmd.String("stage", "", "stage")

	singleCmd := flag.NewFlagSet("single", flag.ExitOnError)
	singleRowKey := singleCmd.String("rowKey", "", "rowKey")
	singlePartitionKey := singleCmd.String("partitionKey", "", "partitionKey")
	singlePropertyName := singleCmd.String("propertyName", "", "propertyName")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateRowKey := updateCmd.String("rowKey", "", "rowKey")
	updatePartitionKey := updateCmd.String("partitionKey", "", "partitionKey")
	updatePropertyName := updateCmd.String("propertyName", "", "propertyName")
	updatePropertyValue := updateCmd.String("propertyValue", "", "propertyValue")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteRowKey := deleteCmd.String("rowKey", "", "rowKey")
	deletePartitionKey := deleteCmd.String("partitionKey", "", "partitionKey")
	deletePropertyName := deleteCmd.String("propertyName", "", "propertyName")

	switch os.Args[1] {

	case "server":

		serverCmd.Parse(os.Args[2:])
		t.Function = "server"
		t.Client, _ = t.Connect()

		listenAddr := ":8080"
		if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
			listenAddr = ":" + val
		}
		http.HandleFunc("/api/table/get", t.GetHandler)
		http.HandleFunc("/api/table/getsingle", t.GetSingleHandler)
		http.HandleFunc("/api/table/update", t.UpdateHandler)
		http.HandleFunc("/api/table/delete", t.DeleteHandler)
		log.Printf("Server started.\n")
		log.Printf("About to listen on Port%s.\nGo to https://127.0.0.1%s/", listenAddr, listenAddr)
		log.Fatal(http.ListenAndServe(listenAddr, nil))

	case "get":

		var err error
		getCmd.Parse(os.Args[2:])

		t.Function = "get"
		t.RowKey = *getRowKey
		t.PartitionKey = *getPartitionKey
		t.Stage = *getStage

		t.Client, _ = t.Connect()

		res, err := t.Get()
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

	case "single":

		var err error
		t.Function = "single"
		t.RowKey = *singleRowKey
		t.PartitionKey = *singlePartitionKey
		t.PropertyName = *singlePropertyName

		t.Client, _ = t.Connect()

		res, err := t.GetSingle()
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

	case "update":

		var err error
		t.Function = "single"
		t.RowKey = *updateRowKey
		t.PartitionKey = *updatePartitionKey
		t.PropertyName = *updatePropertyName
		t.PropertyValue = *updatePropertyValue

		t.Client, _ = t.Connect()

		res, err := t.Update()
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

	case "delete":

		t.Function = "single"
		t.RowKey = *deleteRowKey
		t.PartitionKey = *deletePartitionKey
		t.PropertyName = *deletePropertyName

		t.Client, _ = t.Connect()

	default:
		fmt.Println("expected 'get','single' or 'update' subcommands")
		os.Exit(1)
	}
}
