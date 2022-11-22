package manipulateAzTable

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func GetTableData(client *aztables.Client, partitionKey string, rowKey string, tableName string) *string {

	filter := fmt.Sprintf("PartitionKey eq '%s'", partitionKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(500)),
	}

	pager := client.NewListEntitiesPager(options)
	pageCount := 0

	var export string

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

			if myEntity.RowKey == rowKey {

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
				export = fmt.Sprintln(string(jsonStr))
			}
		}
	}
	return &export
}

func GetSingleTableValue(client *aztables.Client, partitionKey string, rowKey string, tableName string, valueToQuery string) *string {

	type ExportStruct struct {
		Name  string `json:"Key"`
		Value string `json:"Value"`
	}

	filter := fmt.Sprintf("PartitionKey eq '%s'", partitionKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(500)),
	}

	pager := client.NewListEntitiesPager(options)
	pageCount := 0

	var export string

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

			if myEntity.RowKey == rowKey {

				for k, v := range myEntity.Properties {
					if k == valueToQuery {

						jsonStr, err := json.Marshal(ExportStruct{k, v.(string)})

						if err != nil {
							fmt.Printf("Error: %s", err.Error())
						} else {
							fmt.Println(string(jsonStr))
						}

						err = ioutil.WriteFile("data.json", jsonStr, 0644)
						if err != nil {
							log.Fatal(err)
						}
						export = fmt.Sprintln(string(jsonStr))
					}
				}
			}
		}
	}
	return &export
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



