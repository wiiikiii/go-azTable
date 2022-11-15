package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"golang.org/x/exp/slices"
)

var tableName string = "CorpA"

landingZoneStages := []string{
	"structure",
	"connectivity_er",
	"connectivity_ipsec",
	"identity",
	"managementmonitoring",
	"azuread",
	"variables"
}

var workPlaceStages := []string{
	"landingzone",
	"structure",
	"network",
	"infrastructure_avd",
	"sessionhosts"
}

func main() {
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
	fmt.Println(client)

    filter := "PartitionKey eq 'AVD' or RowKey eq 'DeploymentOut'"
    options := &aztables.ListEntitiesOptions{
        Filter: &filter,
        Select: to.Ptr("RowKey,ADDCGuid,ADDSName,AVDGroupID,ApplicationsToInstall"),
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

            fmt.Printf("Received: %v, %v, %v, %v\n", myEntity.RowKey, myEntity.Properties["ADDCGuid"], myEntity.Properties["ADDSName"], myEntity.Properties["AVDGroupID"], myEntity.Properties["ApplicationsToInstall"])
        }
    }
}