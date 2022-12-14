package manipulateAzTable

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var DateTime = time.Now().Format("2006-01-02 15:04:05")

func (t Table) ParseCli() {

	env := (ReturnEnv([]string{
		"TABLES_STORAGE_ACCOUNT_NAME",
		"TABLES_PRIMARY_STORAGE_ACCOUNT_KEY",
		"TABLE_NAME"}))

	t = Table{
		Functions:   []string{"server", "get", "update", "delete", "single", "json", "getstage"},
		AccountName: env["TABLES_STORAGE_ACCOUNT_NAME"],
		AccountKey:  env["TABLES_PRIMARY_STORAGE_ACCOUNT_KEY"],
		TableName:   env["TABLE_NAME"],
	}

	// server
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)

	// get
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getRowKey := getCmd.String("rowKey", "", "rowKey")
	getPartitionKey := getCmd.String("partitionKey", "", "partitionKey")

	//getstage
	getStageCmd := flag.NewFlagSet("getstage", flag.ExitOnError)
	getStageRowKey := getStageCmd.String("rowKey", "", "rowKey")
	getStagePartitionKey := getStageCmd.String("partitionKey", "", "partitionKey")
	getStageStage := getStageCmd.String("stage", "", "stage")

	// single
	singleCmd := flag.NewFlagSet("single", flag.ExitOnError)
	singleRowKey := singleCmd.String("rowKey", "", "rowKey")
	singlePartitionKey := singleCmd.String("partitionKey", "", "partitionKey")
	singlePropertyName := singleCmd.String("propertyName", "", "propertyName")

	// update
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateRowKey := updateCmd.String("rowKey", "", "rowKey")
	updatePartitionKey := updateCmd.String("partitionKey", "", "partitionKey")
	updatePropertyName := updateCmd.String("propertyName", "", "propertyName")
	updatePropertyValue := updateCmd.String("propertyValue", "", "propertyValue")

	// delete
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteRowKey := deleteCmd.String("rowKey", "", "rowKey")
	deletePartitionKey := deleteCmd.String("partitionKey", "", "partitionKey")
	deletePropertyName := deleteCmd.String("propertyName", "", "propertyName")

	//get select values
	configCmd := flag.NewFlagSet("getconfig", flag.ExitOnError)

	switch os.Args[1] {

	case "server":

		serverCmd.Parse(os.Args[2:])
		t.Function = "server"
		t.Client, _ = t.Connect()

		listenAddr := ":8080"
		if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
			listenAddr = ":" + val
		}
		http.HandleFunc("/api/v1/table/get", t.MakeHttpHandler(t.GetHttpHandler))
		http.HandleFunc("/api/v1/table/getsingle", t.MakeHttpHandler(t.GetSingleHttpHandler))
		http.HandleFunc("/api/v1/table/update", t.MakeHttpHandler(t.UpdateHttpHandler))
		http.HandleFunc("/api/v1/table/delete", t.MakeHttpHandler(t.DeleteHttpHandler))
		http.HandleFunc("/api/v1/table/config", t.MakeHttpHandler(t.GetConfigHttpHandler))
		http.HandleFunc("/api/v1/table/updateconfig", t.MakeHttpHandler(t.UpdateConfigHttpHandler))
		log.Printf("Server started.\n")
		log.Printf("About to listen on Port%s.\nGo to https://127.0.0.1%s/", listenAddr, listenAddr)
		log.Fatal(http.ListenAndServe(listenAddr, nil))

	case "get":

		getCmd.Parse(os.Args[2:])
		t.Function = "get"
		t.RowKey = fmt.Sprintf(*getRowKey)
		t.PartitionKey = fmt.Sprintf(*getPartitionKey)
		t.Client, _ = t.Connect()

		//fmt.Printf("%+v\n", t) <-- Print all Values from Struct

		res, err := t.Get()
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

	case "getstage":

		getStageCmd.Parse(os.Args[2:])
		t.Function = "getstage"
		t.RowKey = fmt.Sprintf(*getStageRowKey)
		t.PartitionKey = fmt.Sprintf(*getStagePartitionKey)
		t.Stage = fmt.Sprintf(*getStageStage)
		t.Client, _ = t.Connect()

		res, err := t.GetStage()
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

	case "single":

		singleCmd.Parse(os.Args[2:])
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

		updateCmd.Parse(os.Args[2:])
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

		deleteCmd.Parse(os.Args[2:])
		t.Function = "single"
		t.RowKey = *deleteRowKey
		t.PartitionKey = *deletePartitionKey
		t.PropertyName = *deletePropertyName
		t.Client, _ = t.Connect()

	case "getconfig":
		configCmd.Parse(os.Args[2:])
		t.Function = "select"
		t.PartitionKey = "AVD"
		t.Client, _ = t.Connect()

		current, err := t.GetConfig()
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}

		fmt.Printf("%s", current)

	default:
		fmt.Println("expected 'server', 'get', 'single', 'getstage', 'json' or 'update' subcommands")
		os.Exit(1)
	}
}
