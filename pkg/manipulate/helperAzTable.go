package manipulateAzTable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
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

func (t Table) ParseJson(map[string]interface{}) ([]string, error) {

	var param string
	var export []string
	switch t.StageParamFile = t.Stage; t.StageParamFile {

	case "0_AVD-Landingzone":

		param = "./0_Landingzone.parameters.json"

	case "1_AVD-Structure":

		param = "./1_Structure.parameters.json"

	case "2_AVD-Network":

		param = "./2_Network.parameters.json"

	case "3_AVD-Infrastructure":

		param = "./3_Infrastructure_AVD.parameters.json"

	case "6_AVD-Sessionhosts":

		param = "./5_Sessionhosts.parameters.json"

	default:

		fmt.Printf("Error: couldn`t update or create value")
	}

	if len(param) > 0 {
		content, err := ioutil.ReadFile(param)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
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
