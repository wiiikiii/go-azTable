package manipulateAzTable

import (
	"fmt"
	"os"
	"reflect"
)

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

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (t Table) ValidateParams(string) bool {
	res := reflect.ValueOf(t.Function).IsValid()
	return res
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
