package manipulateAzTable

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type Table struct {
	Client        *aztables.Client
	Function      string
	Functions     []string
	AccountName   string
	AccountKey    string
	TableName     string
	PropertyName  string
	PropertyValue string
	PartitionKey  string
	RowKey        string
	Stage         string
	JSonString    string
}

var dateTime = time.Now().Format("02-01-2006 15:04:05")

func (t Table) Get() ([]byte, error) {

	filter := fmt.Sprintf("PartitionKey eq '%s'", t.PartitionKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(500)),
	}

	pager := t.Client.NewListEntitiesPager(options)
	pageCount := 0

	var export string

	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		pageCount += 1

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}

			if myEntity.RowKey == t.RowKey {

				jsonStr, err := json.Marshal(myEntity.Properties)
				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				}

				export = fmt.Sprintln(string(jsonStr))

			}
		}
	}
	return export, nil
}

func (t Table) GetSingle() ([]byte, error) {

	filter := fmt.Sprintf("PartitionKey eq '%s'", t.PartitionKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(500)),
	}

	pager := t.Client.NewListEntitiesPager(options)
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

			if myEntity.RowKey == t.RowKey {

				for k, v := range myEntity.Properties {
					if k == t.PropertyName {

						r := make(map[string]string)
						r[k] = v.(string)

						jsonStr, err := json.Marshal(r[k])
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

func (t Table) GetStage() (string, error) {

	filter := fmt.Sprintf("PartitionKey eq '%s'", t.PartitionKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(500)),
	}

	pager := t.Client.NewListEntitiesPager(options)
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

			if myEntity.RowKey == t.RowKey {

				r := make(map[string]string)

				xi, err := t.ParseJson()
				if err != nil {
					panic(err)
				}

				for _, i := range xi {

					for k, v := range myEntity.Properties {

						if i == k {

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
	}
	return export, nil
}

func (t Table) Update() (string, error) {

	myAddEntity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: t.PartitionKey,
			RowKey:       t.RowKey,
		},
		Properties: map[string]interface{}{
			t.PropertyName: t.PropertyValue,
		},
	}

	upsertEntityOptions := aztables.UpsertEntityOptions{
		UpdateMode: "merge",
	}

	marshalled, err := json.Marshal(myAddEntity)
	if err != nil {
		return "", errors.New("couldn`t convert to json")
	}

	_, err = t.Client.UpsertEntity(context.TODO(), marshalled, &upsertEntityOptions)
	if err != nil {
		return "", errors.New("couldn`t update or create value")
	}

	var export string

	r := make(map[string]string)
	r[t.PropertyName] = t.PropertyValue

	jsonStr, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	export = fmt.Sprintln(string(jsonStr))
	return export, nil
}

func (t Table) UpdateJSON() (string, error) {

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(t.JSonString), &jsonMap)

	if err != nil {
		panic(err)
	}

	var export string
	r := make(map[string]string)

	for key, value := range jsonMap {

		myAddEntity := aztables.EDMEntity{
			Entity: aztables.Entity{
				PartitionKey: t.PartitionKey,
				RowKey:       t.RowKey,
			},
			Properties: map[string]interface{}{
				key: value,
			},
		}

		upsertEntityOptions := aztables.UpsertEntityOptions{
			UpdateMode: "merge",
		}

		marshalled, err := json.Marshal(myAddEntity)
		if err != nil {
			return "", errors.New("couldn`t convert to json")
		}

		_, err = t.Client.UpsertEntity(context.TODO(), marshalled, &upsertEntityOptions)
		if err != nil {
			return "", errors.New("couldn`t update or create value")
		}

		valStr := fmt.Sprint(value)

		r[key] = string(valStr)
	}

	jsonStr, err := json.Marshal(r)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	export = fmt.Sprintln(string(jsonStr))

	return export, nil
}

func (t Table) Delete() (string, error) {

	updateEntityOptions := aztables.UpdateEntityOptions{
		UpdateMode: "replace",
	}

	var res string
	res, _ = t.Get()

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(res), &jsonMap)

	delete(jsonMap, t.PropertyName)

	marshalled, err := json.Marshal(res)
	if err != nil {
		return "", errors.New("couldn`t convert to json")
	}

	_, err = t.Client.UpdateEntity(context.TODO(), marshalled, &updateEntityOptions)
	if err != nil {
		return "", errors.New("couldn`t update or create value")
	}

	var export string

	r := make(map[string]string)
	r[t.PropertyName] = "success"

	jsonStr, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	export = fmt.Sprintln(string(jsonStr))
	return export, nil
}

func (t Table) GetConfig() ([]byte, error) {
	t.RowKey = "SELECT-VALUES"
	var jsonStruct JsonStruct

	filter := fmt.Sprintf("PartitionKey eq '%s'", t.PartitionKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(500)),
	}

	pager := t.Client.NewListEntitiesPager(options)
	pageCount := 0

	var err error

	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		pageCount += 1

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}

			if myEntity.RowKey == t.RowKey {

				jsonStruct.Status = "200"
				jsonStruct.ErrorText = ""
				jsonStruct.Meta.LastUpdate = dateTime
				jsonStruct.Configurations.LuUpdate = dateTime
				jsonStruct.Configurations.LuProcessed = dateTime
				jsonStruct.Configurations.Fields.Tier.SelectValues = []string{fmt.Sprint(myEntity.Properties["tierSelectValues"])}
				jsonStruct.Configurations.Fields.Location.SelectValues = []string{fmt.Sprint(myEntity.Properties["locationSelectValues"])}
				jsonStruct.Configurations.Fields.Usercount.SelectValues = []string{fmt.Sprint(myEntity.Properties["userCountSelectValues"])}
				jsonStruct.Configurations.Fields.Maintenance.SelectValues = []string{fmt.Sprint(myEntity.Properties["maintenanceSelectValues"])}
				jsonStruct.Configurations.Fields.Environment.SelectValues = []string{fmt.Sprint(myEntity.Properties["environmentSelectValues"])}
				jsonStruct.Configurations.Fields.Backup.SelectValues = []string{fmt.Sprint(myEntity.Properties["backupSelectValues"])}
				jsonStruct.Configurations.Fields.Recovery.SelectValues = []string{fmt.Sprint(myEntity.Properties["recoverySelectValues"])}
				jsonStruct.Configurations.Fields.Applications.SelectValues = []string{fmt.Sprint(myEntity.Properties["applicationsSelectValues"])}
			}
		}
	}

	t.RowKey = "CURRENT-CONFIG"

	pager = t.Client.NewListEntitiesPager(options)
	pageCount = 0

	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		pageCount += 1

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}

			if myEntity.RowKey == t.RowKey {

				jsonStruct.Meta.Key = "6aad6be5-e01f-4e1b-a8af-887eb7851957"
				jsonStruct.Meta.Name = "e332cdc2-15db-4280-807d-ed119fedd95f"
				jsonStruct.Configurations.Name = "HP-POOL-01"
				jsonStruct.Configurations.Fields.Tier.Values = fmt.Sprint(myEntity.Properties["performanceTier"])
				jsonStruct.Configurations.Fields.Location.Values = fmt.Sprint(myEntity.Properties["location"])
				jsonStruct.Configurations.Fields.Usercount.Values = fmt.Sprint(myEntity.Properties["userCount"])
				jsonStruct.Configurations.Fields.Maintenance.Values = fmt.Sprint(myEntity.Properties["maintenance"])
				jsonStruct.Configurations.Fields.Environment.Values = fmt.Sprint(myEntity.Properties["environment"])
				jsonStruct.Configurations.Fields.Backup.Values = fmt.Sprint(myEntity.Properties["backup"])
				jsonStruct.Configurations.Fields.Recovery.Values = fmt.Sprint(myEntity.Properties["recovery"])
				jsonStruct.Configurations.Fields.Hostpools.Values = fmt.Sprint(myEntity.Properties["hostPools"])
				jsonStruct.Configurations.Fields.Applications.Values = []string{fmt.Sprint(myEntity.Properties["applications"])}

				i, err := json.Marshal(jsonStruct)
				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				}

				return i, nil
			}
		}
	}
	return nil, err
}
