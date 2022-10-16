package js_array_method

import "testing"

func TestFindWithStructSlice(t *testing.T) {
	users := []User{{Id: 10, Name: "Go"}, {Id: 11, Name: "Lang"}}

	user, err := Find(users, func(user User, index int) bool {
		return user.Id == 10
	})

	expected := "Go"

	if err != nil {
		t.Errorf("Expected %s received %s", expected, user.Name)
	}
}

func TestFindWithStringSlice(t *testing.T) {
	var slices = []string{"Go", "Typescript", "Nodejs"}

	ln, _ := Find(slices, "Go")

	expected := "Go"

	if ln != expected {
		t.Errorf("Expected %s received %s", expected, ln)
	}
}
