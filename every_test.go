package js_array_method

import (
	"testing"
)

func TestEveryWithSlice(t *testing.T) {
	resultFalse := Every([]string{"hey", "hi"}, "hi")
	if resultFalse == true {
		t.Errorf("Expected %s received %s", "false", "true")
	}

	resultTrue := Every([]string{"hi", "hi", "hi"}, "hi")

	if resultTrue == false {
		t.Errorf("Expected %s received %s", "true", "false")
	}
}

func TestEveryWithStruct(t *testing.T) {
	users := []User{{Id: 12, Name: "Go"}, {Id: 14, Name: "Go"}}

	resultFalse := Every(users, func(user User, _ int) bool {
		return user.Id == 12
	})

	if resultFalse == true {
		t.Errorf("Expected %s received %s", "false", "true")
	}

	resultTrue := Every(users, func(user User, _ int) bool {
		return user.Name == "Go"
	})

	if resultTrue == false {
		t.Errorf("Expected %s received %s", "true", "false")
	}
}
