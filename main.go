package main

import (
	"flag"
	"fmt"
	"os"

	m "go-table/pkg/manipulate"
)

type ExportStruct struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

var err error

func main() {

	env := (m.ReturnEnv([]string{
		"TABLES_STORAGE_ACCOUNT_NAME",
		"TABLES_PRIMARY_STORAGE_ACCOUNT_KEY",
		"TABLE_NAME"}))

	var functions = []string{"get", "update", "delete", "single"}
	var args = os.Args[1:]

	var f_key = flag.String("f", "", "Missing function name.")
	var p_key = flag.String("pk", "", "Missing Partition key.")
	var r_key = flag.String("rk", "", "Missing row key.")

	flag.Parse()

	t := m.Table{
		Function:    *f_key,
		Functions:   functions,
		AccountName: env["TABLES_STORAGE_ACCOUNT_NAME"],
		AccountKey:  env["TABLES_PRIMARY_STORAGE_ACCOUNT_KEY"],
		TableName:   env["TABLE_NAME"],
	}

	if !m.Contains(functions, *f_key) {
		fmt.Printf("Unknown Parameter %q", t.Function)
		return
	}

	t.Client, err = t.Connect()
	if err != nil {
		fmt.Printf("Got an error while connection table.")
		return
	}

	for _, k := range args {
		if !t.ValidateParams(k) {
			fmt.Printf("Parameter raised an error while parsing.")
			return
		}
	}

	t.PartitionKey = *p_key
	t.RowKey = *r_key

	switch *f_key {
	case "get":
		if len(args) < 3 {
			fmt.Printf("Parameters missing, you have to provide: partitionKey and rowKey")
			return
		}

		res, err := t.Get()
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

	case "update":
		if len(args) < 5 {
			fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey, propertyName and propertyValue")
			break
		}

		t.PropertyName = args[3]
		t.PropertyValue = args[4]

		if t.ValidateParams(t.PropertyName) && t.ValidateParams(t.PropertyValue) {

			res, err := t.Update()
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
		}

	case "delete":
		if len(args) != 3 {
			fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey and propertyName")
			break
		}

		t.PropertyName = args[3]
		if t.ValidateParams(t.PropertyName) {
			t.Delete()
			if err != nil {
				panic(err)
			}
			return
		}

	case "single":
		if len(args) != 4 {
			fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey, tablename and propertyName")
			break
		}

		t.PropertyName = args[3]

		if t.ValidateParams(t.PropertyName) {
			res, err := t.GetSingle()
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
		}

	default:
		fmt.Printf("Unknown Parameter: %q", t.Function)
		return
	}
}
