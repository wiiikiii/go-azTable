package manipulateAzTable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
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

func (t Table) JsonToMap(data map[string]interface{}) map[string][]string {

	out := make(map[string][]string)

	for key, value := range data {

		if _, ok := out[key]; !ok {
			out[key] = []string{}
		}

		if valueA, ok := value.(map[string]interface{}); ok {
			out[key] = append(out[key], "")
			for keyB, valueB := range t.JsonToMap(valueA) {
				if _, ok := out[keyB]; !ok {
					out[keyB] = []string{}
				}
				out[keyB] = append(out[keyB], valueB...)
			}
		} else if valueA, ok := value.([]interface{}); ok {
			for _, valueB := range valueA {
				if valueC, ok := valueB.(map[string]interface{}); ok {
					for keyD, valueD := range t.JsonToMap(valueC) {
						if _, ok := out[keyD]; !ok {
							out[keyD] = []string{}
						}
						out[keyD] = append(out[keyD], valueD...)
					}
				} else {
					out[key] = append(out[key], fmt.Sprintf("%v", valueB))
				}
			}
		} else {
			out[key] = append(out[key], fmt.Sprintf("%v", value))
		}
	}
	return out
}

func (t Table) FindFile(ext string) string {
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var file string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				file = f.Name()
			}
		}
		return err
	})
	return file
}

func (t Table) ParseJson() ([]string, error) {

	var export []string
	var param string

	if len(t.PartitionKey) != 0 {

		param = fmt.Sprintf("%v-%v.parameters.json", t.PartitionKey, t.Stage)
	}

	s := t.FindFile(param)

	if len(s) > 0 {
		content, err := ioutil.ReadFile(s)
		if err != nil {
			log.Fatal("error when opening file: ", err)
		}

		var JSON map[string]interface{}
		json.Unmarshal([]byte(content), &JSON)

		output := t.JsonToMap(JSON)
		for key := range output {
			if strings.HasPrefix(key, "PARAM") {
				export = append(export, key)
			}
		}
	}
	return export, nil
}

func (t Table) GetValueFromMap(v string) (string, error) {
	c, err := ioutil.ReadFile("map.json")
	if err != nil {
		log.Fatal("error when opening file: ", err)
	}

	var J, P map[string]interface{}

	json.Unmarshal([]byte(c), &J)
	json.Unmarshal([]byte(v), &P)

	for key := range J {
		for k, v := range P {
			if k == key {
				return key, nil
			}
			return fmt.Sprint(v), nil
		}
	}
	log.Fatal("error, value not found: ", err)
	return "", err
}
