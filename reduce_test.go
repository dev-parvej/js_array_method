package js_array_method

import "testing"

func TestReduce(t *testing.T) {
	numbers := []int{10, 11, 12, 13}

	sum := Reduce(numbers, func(s int, n int, i int) int {
		return s + n
	}, 0)

	if sum != 46 {
		t.Errorf("Expected %d received %d", 46, sum)
	}
}
