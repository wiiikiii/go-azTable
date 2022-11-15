package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type InventoryEntity struct {
	aztables.Entity
	Price       float32
	Inventory   int32
	ProductName string
	OnSale      bool
}

type PurchasedEntity struct {
	aztables.Entity
	Price float32
	ProductName string
	OnSale bool
}

func getClient() *aztables.Client {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT environment variable not found")
	}

	tableName, ok := os.LookupEnv("AZURE_TABLE_NAME")
	if !ok {
		panic("AZURE_TABLE_NAME environment variable not found")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	serviceURL := fmt.Sprintf("https://%s.table.core.windows.net/%s", accountName, tableName)
	client, err := aztables.NewClient(serviceURL, cred, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func listEntities(client *aztables.Client) {
	listPager := client.GetEntity("AVD", "Input" )
	pageCount := 0
	for listPager.More() {
		response, err := listPager.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}
		fmt.Printf("There are %d entities in page #%d\n", len(response.Entities), pageCount)
		pageCount += 1
	}
}

func main() {

	fmt.Println("Authenticating...")
	client := getClient()

	fmt.Println("Calculating all entities in the table...")
	listEntities(client)

}