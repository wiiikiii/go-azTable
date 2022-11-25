package manipulateAzTable

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

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
}

func (t Table) Connect() (cl *aztables.Client, err error) {

	serviceURL := fmt.Sprintf("https://%s.table.core.windows.net/%s", t.AccountName, t.TableName)

	cred, err := aztables.NewSharedKeyCredential(t.AccountName, t.AccountKey)
	if err != nil {
		panic(err)
	}
	client, err := aztables.NewClientWithSharedKey(serviceURL, cred, nil)
	if err != nil {
		panic(err)
	}

	return client, err
}

func (t Table) Get() (string, error) {

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

func (t Table) GetSingle() (string, error) {

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

func (t Table) ValidateParams(string) bool {
	res := reflect.ValueOf(t.Function).IsValid()
	return res
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (t Table) ReturnEnv(s []string) map[string]string {

	m := make(map[string]string)

	for _, v := range s {

		k, ok := os.LookupEnv(v)
		if !ok {
			fmt.Println("Not found:", v)
			m[v] = "false"
		} else {
			m[v] = k
		}
	}
	return m
}

func (t Table) GetHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	var message string
	name := r.URL.Query().Get("function")
	if name == "get" {
		message, err = t.Get()
	}
	if name == "update" {
		message, err = t.Update()
	}
	if name == "single" {
		message, err = t.GetSingle()
	}
	if name == "delete" {
		message, err = t.Delete()
	}

	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, message)
}

func ReturnEnv(s []string) map[string]string {

	m := make(map[string]string)

	for _, v := range s {

		k, ok := os.LookupEnv(v)
		if !ok {
			fmt.Println("Not found:", v)
			m[v] = "false"
		} else {
			m[v] = k
		}
	}
	return m
}

func ValidateParams(value string) bool {
	res := reflect.ValueOf(value).IsValid()
	return res
}
