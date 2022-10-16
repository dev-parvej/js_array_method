package js_array_method

import (
	"testing"
)

func TestFilterWithStructSlice(t *testing.T) {
	users := []User{{Id: 10, Name: "Go"}, {Id: 11, Name: "Lang"}}

	mapedUsers := Filter(users, func(user User, index int) bool {
		return user.Id == 10
	})
	expected := 10

	if mapedUsers[0].Id != 10 {
		t.Errorf("Expected %d received %d", expected, mapedUsers[0].Id)
	}
}

func TestFilterWithStringSlice(t *testing.T) {
	var slices = []string{"Go", "Typescript", "Nodejs"}

	ln := Filter(slices, "Go")

	expected := "Go"

	if ln[0] != expected {
		t.Errorf("Expected %s received %s", expected, ln[0])
	}
}
