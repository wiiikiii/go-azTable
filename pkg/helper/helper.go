package helperFuncs

import "encoding/json"

func RemoveKeysFromJSONObject(input *map[string]json.RawMessage, keys []string) {
	for _, key := range keys {
		delete(*input, key)
	}
}

func MapFuncToJsonObjectArray(fn func(input *map[string]json.RawMessage) error, jsonArray *[]map[string]json.RawMessage) error {
	for _, jsonArrayItem := range *jsonArray {
		err := fn(&jsonArrayItem)
		if err != nil {
			return err
		}
	}
	return nil
}

func MapFuncToJsonObjectArrayBytes(fn func(input *map[string]json.RawMessage) error, jsonArray *[]byte) error {
	var output []map[string]json.RawMessage
	if err := json.Unmarshal(*jsonArray, &output); err != nil {
		return err
	}
	err := MapFuncToJsonObjectArray(fn, &output)
	if err != nil {
		return err
	}
	// https://stackoverflow.com/a/24229303/4562156
	// The value passed to json.Marshal must be a pointer for json.RawMessage to work properly.
	outputBytes, err := json.Marshal(&output)
	if err != nil {
		return err
	}
	*jsonArray = outputBytes
	return nil
}

func RemoveKeysFromJSONObjectBytes(input *[]byte, keys []string) error {
	var output map[string]json.RawMessage
	if err := json.Unmarshal(*input, &output); err != nil {
		return err
	}
	err := RemoveKeysFromJSONObject(&output, keys)
	if err != nil {
		return err
	}
	outputBytes, err := json.Marshal(&output)
	if err != nil {
		return err
	}
	*input = outputBytes
	return nil
}