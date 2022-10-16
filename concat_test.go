package js_array_method

import (
	"testing"
)

func TestConcat(t *testing.T) {
	users1 := []User{{Id: 10, Name: "John"}, {Id: 11, Name: "Doe"}, {Id: 12, Name: "Sabrina"}}
	users2 := []User{{Id: 10, Name: "John"}, {Id: 11, Name: "Doe"}, {Id: 12, Name: "Sabrina"}}

	users := Concat(users1, users2)

	if len(users) != 6 {
		t.Errorf("Expected 6 got %d", len(users))
	}
}
