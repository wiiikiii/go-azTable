package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

var storageaccountkey string = "OVCT1/fNFiMJSaFX38SEJT/kevTkyWLph5gZHRMre9ilguo52h3WlRrwpkc+SF/eC7sNPuZ0MJxu+AStayyzuQ=="

func main() {

	accountname := "axiansvdconfig"
	serviceURL := accountname + ".table.core.windows.net"

    cred, err := aztables.NewSharedKeyCredential(serviceURL, storageaccountkey)
    if err != nil {
        panic(err)
    }
    client, err := aztables.NewServiceClientWithSharedKey(serviceURL, cred, nil)
    if err != nil {
        panic(err)
    }

	filter := "PartitionKey eq 'AVD' or RowKey eq 'Input'"
    options := &aztables.ListEntitiesOptions{
        Filter: &filter,
        Select: to.Ptr("RowKey,Value,Product,Available"),
        Top: to.Ptr(int32(15)),
    }

    pager := client.NewListEntitiesPager(options)
    pageCount := 0
    for pager.More() {
        response, err := pager.NextPage(context.TODO())
        if err != nil {
            panic(err)
        }
        fmt.Printf("There are %d entities in page #%d\n", len(response.Entities), pageCount)
        pageCount += 1

        for _, entity := range response.Entities {
            var myEntity aztables.EDMEntity
            err = json.Unmarshal(entity, &myEntity)
            if err != nil {
                panic(err)
            }

            fmt.Printf("Received: %v, %v, %v, %v\n", myEntity.RowKey, myEntity.Properties["Value"], myEntity.Properties["Product"], myEntity.Properties["Available"])
        }
    }
}