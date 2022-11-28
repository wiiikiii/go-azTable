package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getEnable := getCmd.Bool("true", true, "true")
	getRowKey := getCmd.String("rowKey", "", "rowKey")
	getPartitionKey := getCmd.String("partitionKey", "", "partitionKey")
	getStage := getCmd.String("stage", "", "stage")

	singleCmd := flag.NewFlagSet("single", flag.ExitOnError)
	singleEnable := singleCmd.Bool("true", true, "true")
	singleRowKey := singleCmd.String("rowKey", "", "rowKey")
	singlePartitionKey := singleCmd.String("partitionKey", "", "partitionKey")
	singlePropertyName := singleCmd.String("propertyName", "", "propertyName")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateEnable := updateCmd.Bool("true", true, "true")
	updateRowKey := updateCmd.String("rowKey", "", "rowKey")
	updatePartitionKey := updateCmd.String("partitionKey", "", "partitionKey")
	updatePropertyName := updateCmd.String("propertyName", "", "propertyName")
	updatePropertyValue := updateCmd.String("propertyValue", "", "propertyValue")

	switch os.Args[1] {

	case "get":
		getCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'get'")
		fmt.Println("  enable:", *getEnable)
		fmt.Println("  rowKey:", *getRowKey)
		fmt.Println("  partitionKey:", *getPartitionKey)
		fmt.Println("  stage:", *getStage)

	case "single":
		singleCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'single'")
		fmt.Println("  enable:", *singleEnable)
		fmt.Println("  rowKey:", *singleRowKey)
		fmt.Println("  partitionKey:", *singlePartitionKey)
		fmt.Println("  propertyName:", *singlePropertyName)
	case "update":
		updateCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'update'")
		fmt.Println("  enable:", *updateEnable)
		fmt.Println("  rowKey:", *updateRowKey)
		fmt.Println("  partitionKey:", *updatePartitionKey)
		fmt.Println("  propertyName:", *updatePropertyName)
		fmt.Println("  propertyValue:", *updatePropertyValue)
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
