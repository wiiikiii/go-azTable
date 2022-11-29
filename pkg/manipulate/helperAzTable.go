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

func JsonToMap(data map[string]interface{}) map[string][]string {
	// final output
	out := make(map[string][]string)

	// check all keys in data
	for key, value := range data {
		// check if key not exist in out variable, add it
		if _, ok := out[key]; !ok {
			out[key] = []string{}
		}

		if valueA, ok := value.(map[string]interface{}); ok { // if value is map
			out[key] = append(out[key], "")
			for keyB, valueB := range JsonToMap(valueA) {
				if _, ok := out[keyB]; !ok {
					out[keyB] = []string{}
				}
				out[keyB] = append(out[keyB], valueB...)
			}
		} else if valueA, ok := value.([]interface{}); ok { // if value is array
			for _, valueB := range valueA {
				if valueC, ok := valueB.(map[string]interface{}); ok {
					for keyD, valueD := range JsonToMap(valueC) {
						if _, ok := out[keyD]; !ok {
							out[keyD] = []string{}
						}
						out[keyD] = append(out[keyD], valueD...)
					}
				} else {
					out[key] = append(out[key], fmt.Sprintf("%v", valueB))
				}
			}
		} else { // if string and numbers and other ...
			out[key] = append(out[key], fmt.Sprintf("%v", value))
		}
	}
	return out
}


// func main() {

// 	content, err := ioutil.ReadFile("./test.json")
// 	if err != nil {
// 		log.Fatal("Error when opening file: ", err)
// 	}

// 	var JSON map[string]interface{}
// 	json.Unmarshal([]byte(content), &JSON)

// 	neededOutput := jsonToMap(JSON)

// 	for key := range neededOutput {
// 		if strings.HasPrefix(key, "PARAM"){
// 			fmt.Println(key)
// 		}
		
// 	}
// }