package helper

import "reflect"

func ValidateParams(value string) bool {
	res := reflect.ValueOf(value).IsValid()
	return res
}
