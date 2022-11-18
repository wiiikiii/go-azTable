package manipulateTableData

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func ConnectStorageAccount(tableName string) (cl *aztables.Client, err error) {

	accountName, ok := os.LookupEnv("TABLES_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic(" TABLES_STORAGE_ACCOUNT_NAME could not be found")
	}
	accountKey, ok := os.LookupEnv("TABLES_PRIMARY_STORAGE_ACCOUNT_KEY")
	if !ok {
		panic(" TABLES_PRIMARY_STORAGE_ACCOUNT_KEY could not be found")
	}

	serviceURL := fmt.Sprintf("https://%s.table.core.windows.net/%s", accountName, tableName)

	cred, err := aztables.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		panic(err)
	}
	client, err := aztables.NewClientWithSharedKey(serviceURL, cred, nil)
	if err != nil {
		panic(err)
	}

	return client, err
}

func GetTableData(client *aztables.Client, partitionKey string, rowKey string, tableName string) string {

	type ListEntitiesOptions struct {
		Filter           *string
		Select           *string
		Top              int32
		NextPartitionKey *string
		NextRowKey       *string
	}

	filter := fmt.Sprintf("PartitionKey eq '%s' or RowKey eq '%s'", partitionKey, rowKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(15)),
	}

	pager := client.NewListEntitiesPager(options)
	pageCount := 0

	var Export string

	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}
		pageCount += 1

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				panic(err)
			}

			jsonStr, err := json.Marshal(myEntity.Properties)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			} else {
				fmt.Println(string(jsonStr))
			}

			err = ioutil.WriteFile("data.json", jsonStr, 0644)
			if err != nil {
				log.Fatal(err)
			}

			Export := fmt.Sprintln(string(jsonStr))
			return Export
		}
		return Export
	}
	return Export
}

func WriteTableData(partitionKey string, rowKey string, tableName string) {
}

func UpdateTableData(partitionKey string, rowKey string, tableName string) {
}

func UpdateTableProperties(client *aztables.Client, partitionKey string, rowKey string, tableName string) {
}

func DeleteTableProperties(partitionKey string, rowKey string, tableName string) {
}

func WriteTableProperties(client *aztables.Client, partitionKey string, rowKey string, tableName string) {

	type InventoryEntity struct {
		aztables.Entity
		Price       float32
		Inventory   int32
		ProductName string
		OnSale      bool
	}

	//TODO: Check access policy, Storage Blob Data Contributor role needed
	_, err := client.CreateTable(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	myEntity := InventoryEntity{
		Entity: aztables.Entity{
			PartitionKey: partitionKey,
			RowKey:       rowKey,
		},
		Price:       3.99,
		Inventory:   20,
		ProductName: "Markers",
		OnSale:      false,
	}

	marshalled, err := json.Marshal(myEntity)
	if err != nil {
		panic(err)
	}

	_, err = client.AddEntity(context.TODO(), marshalled, nil) // TODO: Check access policy, need Storage Table Data Contributor role
	if err != nil {
		panic(err)
	}

}

func GetSingleTableValue(client *aztables.Client, partitionKey string, rowKey string, tableName string, valueToQuery string) string {

	filter := fmt.Sprintf("PartitionKey eq '%s' or RowKey eq '%s'", partitionKey, rowKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(15)),
	}

	pager := client.NewListEntitiesPager(options)
	pageCount := 0

	var export string

	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			fmt.Printf(err.Error())
			panic(err)
		}
		pageCount += 1

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				fmt.Printf(err.Error())
				panic(err)
			}

			jsonStr, err := json.Marshal(myEntity.Properties[valueToQuery])
						if err != nil {
				fmt.Printf("Error: %s", err.Error())
			} else {
				fmt.Println(string(jsonStr))
			}

			err = ioutil.WriteFile("data.json", jsonStr, 0644)
			if err != nil {
				log.Fatal(err)
			}

			Export := fmt.Sprintln(string(jsonStr))
			return Export
		}
	}
	return export
}
