package js_array_method

import "testing"

func TestSomeWithoutCallback(t *testing.T) {
	resultFalse := Some([]string{"hey", "hi"}, "hello")
	if resultFalse == true {
		t.Errorf("Expected %s received %s", "false", "true")
	}

	resultTrue := Some([]string{"hi", "hello", "hiya"}, "hi")

	if resultTrue == false {
		t.Errorf("Expected %s received %s", "true", "false")
	}
}

func TestSomeWithCallback(t *testing.T) {
	users := []User{{Id: 12, Name: "Go"}, {Id: 14, Name: "Lang"}}

	resultFalse := Some(users, func(user User, _ int) bool {
		return user.Id == 15
	})

	if resultFalse == true {
		t.Errorf("Expected %s received %s", "false", "true")
	}

	resultTrue := Some(users, func(user User, _ int) bool {
		return user.Name == "Go"
	})

	if resultTrue == false {
		t.Errorf("Expected %s received %s", "true", "false")
	}
}
