package manipulateAzTable

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// TODO: implement mapping file, for json pass

func (t Table) GetHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t.PartitionKey = r.URL.Query().Get("PartitionKey")
		t.RowKey = r.URL.Query().Get("RowKey")

		message, err := t.Get()
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, message)
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func (t Table) GetSingleHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t.PartitionKey = r.URL.Query().Get("PartitionKey")
		t.RowKey = r.URL.Query().Get("RowKey")
		t.PropertyName = r.URL.Query().Get("PropertyName")

		message, err := t.GetSingle()
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, message)
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func (t Table) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		r.ParseForm()

		data := struct {
			Method      string
			URL         *url.URL
			Submissions url.Values
		}{
			r.Method,
			r.URL,
			r.Form,
		}

		s := data.Submissions
		t.PartitionKey = fmt.Sprintf(string(s.Get("PartitionKey")))
		t.RowKey = fmt.Sprintf(string(s.Get("RowKey")))
		t.PropertyName = fmt.Sprintf(string(s.Get("PropertyName")))
		t.PropertyValue = fmt.Sprintf(string(s.Get("PropertyValue")))

		if t.ValidateParams(t.PropertyName) && t.ValidateParams(t.PropertyValue) {
			message, err := t.Update()
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			fmt.Fprint(w, message)

		} else {
			http.Error(w, "not enough parameters.", http.StatusBadRequest)
		}

	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func (t Table) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {

		r.ParseForm()

		data := struct {
			Method      string
			URL         *url.URL
			Submissions url.Values
		}{
			r.Method,
			r.URL,
			r.Form,
		}

		s := data.Submissions
		t.PropertyName = fmt.Sprintln(string(s.Get("PropertyName")))

		if reflect.ValueOf(t.PropertyName).IsValid() {
			fmt.Print(w, "Here a delete property function should be implemented")
		} else {
			http.Error(w, "not enough parameters.", http.StatusBadRequest)
		}

	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}
