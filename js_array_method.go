package js_array_method

import (
	"errors"
)

func Foreach[T any](input []T, callback func(element T, index int)) {
	for index, value := range input {
		callback(value, index)
	}
}

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
func Find[T any](input []T, conditions interface{}) (T, error) {
	var item T
	index := FindIndex(input, conditions)

	if index < 0 {
		return item, errors.New("nothing found")
	}

	item = input[index]

	return item, nil
}

// conditions can be callback or a value
// callback has two parameter one is element and another is index
// callback must return true or false
func FindIndex[T any](input []T, conditions interface{}) int {
	var itemIndex int = -1
	conditionType := getType(conditions)
	for index, value := range input {
		if conditionType == "func" && callCallback(conditions, value, index) {
			itemIndex = index
			break
		} else if conditionType != "func" && compareValue(conditions, value) {
			itemIndex = index
			break
		}
	}

	return itemIndex
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
	result := FindIndex(input, conditions)
	return result >= 0
}

func Reverse[T any](input []T) []T {
	length := len(input) - 1
	var output []T
	for i := length; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}
