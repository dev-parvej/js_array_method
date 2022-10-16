package js_array_method

import "testing"

func TestReverse(t *testing.T) {
	users := []User{{Id: 10, Name: "John"}, {Id: 11, Name: "Doe"}, {Id: 12, Name: "Sabrina"}}

	reversedUsers := Reverse(users)
	if reversedUsers[0].Id != 12 {
		t.Errorf("Expected %d received %d", 12, reversedUsers[0].Id)
	}

	if reversedUsers[1].Id != 11 {
		t.Errorf("Expected %d received %d", 11, reversedUsers[0].Id)
	}
}
