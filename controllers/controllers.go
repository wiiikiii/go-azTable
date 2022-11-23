package controllers

// import (
// 	"encoding/json"
// 	"net/http"

// 	manipulateAzTable "go-table/pkg/manipulate"
// )

// func GetData(w http.ResponseWriter, r *http.Request){

// 	b, err := manipulateAzTable.GetTableData(client, partitionKey, rowKey, tableName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	res, _ := json.Marshal(b)
// 	w.Header().Set("Content-Type","application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
