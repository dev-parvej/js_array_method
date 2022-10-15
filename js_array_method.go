package js_array_method

func Map[T any, U any](input []T, callback func(element T, index int) U) []U {
	var output []U

	for index, value := range input {
		result := callback(value, index)

		output = append(output, result)
	}

	return output
}

func Filter[T any](input []T, callback func(element T, index int) bool) []T {
	var output []T

	for index, value := range input {
		if callback(value, index) {
			output = append(output, value)
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

func Find[T any](input []T, callback func(element T, index int) bool) T {
	var item T
	for index, value := range input {
		if callback(value, index) {
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
