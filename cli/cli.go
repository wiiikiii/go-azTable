package cli

import (
	"flag"
	"fmt"
	"os"
)

type ArgStruct struct {
	Name          string
	Enable        bool
	RowKey        string
	PartitionKey  string
	Stage         string
	PropertyName  string
	PropertyValue string
}

func (r ArgStruct) ParseInputArgs() *ArgStruct{

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

		g := ArgStruct{
			Name:         "get",
			Enable:       *getEnable,
			RowKey:       *getRowKey,
			PartitionKey: *getPartitionKey,
			Stage:        *getStage,
		}

		return &g

	case "single":

		g := ArgStruct{
			Name:         "single",
			Enable:       *singleEnable,
			RowKey:       *singleRowKey,
			PartitionKey: *singlePartitionKey,
			PropertyName: *singlePropertyName,
		}

		return &g

	case "update":

		g := ArgStruct{
			Name:          "single",
			Enable:        *updateEnable,
			RowKey:        *updateRowKey,
			PartitionKey:  *updatePartitionKey,
			PropertyName:  *updatePropertyName,
			PropertyValue: *updatePropertyValue,
		}

		return &g

	default:
		fmt.Println("expected 'get','single' or 'update' subcommands")
		os.Exit(1)
		return nil
	}
}
