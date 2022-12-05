package main

import (
	"fmt"
	m "go-table/pkg"
)

func main() {
	m.Table.ParseCli(m.Table{})

	ct := m.T{
		Status:    200,
		ErrorText: "Das ist ein Test",
		Meta: struct {
			Key        string `json:"key"`
			Name       string `json:"name"`
			LastUpdate string `json:"lastUpdate"`
		}{},
		Configurations: nil,
	}

	fmt.Println(ct)
}
