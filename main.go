package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	m "go-table/pkg/manipulate"
)

type ExportStruct struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

var envs = []string{"TABLES_STORAGE_ACCOUNT_NAME", "TABLES_PRIMARY_STORAGE_ACCOUNT_KEY", "TABLE_NAME"}

func main() {

	var functions = []string{"server", "get", "update", "delete", "single"}

	var args = os.Args[1:]
	function := args[0]

	t := m.Table{
		Function:  function,
		Functions: functions,
	}

	if m.Contains(functions, function) {

		if function == "server" {

			listenAddr := ":8080"
			if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
				listenAddr = ":" + val
			}
			http.HandleFunc("/api/", t.GetHandler)
			log.Printf("Server started.\n")
			log.Printf("About to listen on Port%s.\nGo to https://127.0.0.1%s/", listenAddr, listenAddr)
			log.Fatal(http.ListenAndServe(listenAddr, nil))

		} else {

			valid := true

			for _, k := range args {
				if !t.ValidateParams(k) {
					valid = false
					break
				}
			}

			t.PartitionKey = args[1]
			t.RowKey = args[2]

			if valid {

				switch {

				case function == "get":

					if len(args) == 4 {
						var err error
						t.Client, err = t.Connect()
						if err != nil {
							panic(err)
						}

						res, err := t.Get()
						if err != nil {
							panic(err)
						}
						fmt.Println(res)

					} else {
						fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey and tablename")
						break
					}

				case function == "update":

					if len(args) == 6 {

						t.PropertyName = args[4]
						t.PropertyValue = args[5]

						if t.ValidateParams(t.PropertyName) && t.ValidateParams(t.PropertyValue) {
							var err error
							t.Client, err = t.Connect()
							if err != nil {
								panic(err)
							}

							res, err := t.Update()
							if err != nil {
								panic(err)
							}
							fmt.Println(res)
						}

					} else {
						fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey, tablename, propertyName and propertyValue")
						break
					}

				case function == "delete":

					if len(args) == 5 {

						t.PropertyName = args[4]

						if t.ValidateParams(t.PropertyName) {
							var err error
							t.Client, err = t.Connect()
							if err != nil {
								panic(err)
							}

							t.Delete()
							if err != nil {
								panic(err)
							}
							return
						}

					} else {
						fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey, tablename and propertyName")
						break

					}

				case function == "single":

					if len(args) == 5 {

						t.PropertyName = args[4]

						if t.ValidateParams(t.PropertyName) {
							var err error
							t.Client, err = t.Connect()
							if err != nil {
								panic(err)
							}
							res, err := t.GetSingle()
							if err != nil {
								panic(err)
							}
							fmt.Println(res)
						}

					} else {
						fmt.Printf("Parameters missing, you have to provide: partitionKey, rowKey, tablename and propertyName")
						break
					}

				default:
					fmt.Printf("Unknown Parameter %q", t.Function)
				}
			}
		}
	} else {
		fmt.Printf("%v is not a supported function, choose as first parameter from:\n %q", t.Function, t.Functions)
		return
	}

}
