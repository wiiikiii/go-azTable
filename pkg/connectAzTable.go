package manipulateAzTable

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

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
