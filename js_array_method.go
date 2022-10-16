package js_array_method

import (
	"fmt"
	"reflect"
)

func Map[T any, U any](input []T, callback func(element T, index int) U) []U {
	var output []U

	for index, value := range input {
		result := callback(value, index)

		output = append(output, result)
	}

	return output
}

// conditions can be callback or a value
// callback has two parameter one is element and another is index
// callback must return true or false
func Filter[T any](input []T, conditions interface{}) []T {
	var output []T
	conditionType := getType(conditions)

	for index, value := range input {
		if conditionType == "func" && callCallback(conditions, value, index) {
			output = append(output, value)
			continue
		} else if conditionType != "func" && compareValue(conditions, value) {
			output = append(output, value)
			continue
		}
	}

	return output
}

func Reduce[T any, U any](input []T, callback func(pre U, current T, index int) U, initialValue U) U {
	output := initialValue

	for index, value := range input {
		output = callback(output, value, index)
	}

	return output
}

// conditions can be callback or a value
// callback has two parameter one is element and another is index
// callback must return true or false
func Find[T any](input []T, conditions interface{}) T {
	var item T
	conditionType := getType(conditions)
	for index, value := range input {
		if conditionType == "func" && callCallback(conditions, value, index) {
			item = value
			break
		} else if conditionType != "func" && compareValue(conditions, value) {
			item = value
			break
		}
	}

	return item
}

func Foreach[T any](input []T, callback func(element T, index int)) {
	for index, value := range input {
		callback(value, index)
	}
}

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

// conditions can be callback or a value
// callback has two parameter one is element and another is index
// callback must return true or false
func Every[T any](input []T, conditions interface{}) any {
	output := calculateEveryElement(input, conditions)
	return len(output) == len(input)
}

// conditions can be callback or a value
// callback has two parameter one is element and another is index
// callback must return true or false
func Some[T any](input []T, conditions interface{}) any {
	output := calculateEveryElement(input, conditions)
	return len(output) > 0
}

// conditions can be callback or a value
// callback has two parameter one is element and another is index
// callback must return true or false
func Includes[T any](input []T, conditions interface{}) bool {
	result := Find(input, conditions)

	if getType(result) == "struct" && reflect.ValueOf(result).IsZero() {
		return false
	} else if getType(result) != "struct" && fmt.Sprintf("%v", result) == "" {
		return false
	}
	return true
}

func Reverse[T any](input []T) []T {
	length := len(input) - 1
	var output []T
	for i := length; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}
