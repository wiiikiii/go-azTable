package manipulateAzTable

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func GetTableData(client *aztables.Client, partitionKey string, rowKey string, tableName string) (string, error) {

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
				}

				err = ioutil.WriteFile("data.json", jsonStr, 0644)
				if err != nil {
					log.Fatal(err)
				}

				export = fmt.Sprintln(string(jsonStr))
			}
		}
	}
	return export, nil
}

func GetSingleTableValue(client *aztables.Client, partitionKey string, rowKey string, tableName string, tableProperty string) (string, error) {

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
					if k == tableProperty {

						r := make(map[string]string)
						r[k] = v.(string)

						jsonStr, err := json.Marshal(r)
						if err != nil {
							fmt.Printf("Error: %s", err.Error())
						}
						export = fmt.Sprintln(string(jsonStr))
					}
				}
			}
		}
	}
	return export, nil
}

func UpdateTableProperties(client *aztables.Client, partitionKey string, rowKey string, tableName string, propertyName string, propertyValue string) (string, error) {

	myAddEntity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: partitionKey,
			RowKey:       rowKey,
		},
		Properties: map[string]interface{}{
			propertyName: propertyValue,
		},
	}

	upsertEntityOptions := aztables.UpsertEntityOptions{
		UpdateMode: "merge",
	}

	marshalled, err := json.Marshal(myAddEntity)
	if err != nil {
		return "", errors.New("couldnt convert to json")
	}

	_, err = client.UpsertEntity(context.TODO(), marshalled, &upsertEntityOptions)
	if err != nil {
		return "", errors.New("couldnt update or create value")
	}

	var export string

	r := make(map[string]string)
	r[propertyName] = propertyValue

	jsonStr, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	export = fmt.Sprintln(string(jsonStr))
	return export, nil
}

func DeleteTableProperties(client *aztables.Client, partitionKey string, rowKey string, tableName string, propertyName string) (string, error) {

	updateEntityOptions := aztables.UpdateEntityOptions{
		UpdateMode: "replace",
	}

	var res string
	res, _ = GetTableData(client, partitionKey, rowKey, tableName)

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(res), &jsonMap)

	delete(jsonMap, propertyName)

	marshalled, err := json.Marshal(res)
	if err != nil {
		return "", errors.New("couldnt convert to json")
	}

	_, err = client.UpdateEntity(context.TODO(), marshalled, &updateEntityOptions)
	if err != nil {
		return "", errors.New("couldnt update or create value")
	}

	var export string

	r := make(map[string]string)
	r[propertyName] = "success"

	jsonStr, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	export = fmt.Sprintln(string(jsonStr))
	return export, nil
}
