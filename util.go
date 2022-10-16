package js_array_method

import (
	"fmt"
	"reflect"
)

func callCallback[T any](conditions interface{}, value T, index int) bool {
	res := reflect.ValueOf(conditions).Call([]reflect.Value{reflect.ValueOf(value), reflect.ValueOf(index)})

	return res[0].Bool()
}

func compareValue[T any](conditions interface{}, value T) bool {
	res := reflect.ValueOf(conditions).String()
	return res == fmt.Sprintf("%v", value)
}

func getType(conditions interface{}) string {
	return reflect.TypeOf(conditions).Kind().String()
}

func calculateEveryElement[T any](input []T, conditions interface{}) []bool {
	output := []bool{}

	conditionType := getType(conditions)
	for index, value := range input {
		result := false
		if conditionType == "func" {
			result = callCallback(conditions, value, index)
		} else {
			result = compareValue(conditions, value)
		}
		if result {
			output = append(output, true)
		}
	}

	return output
}
