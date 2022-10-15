package js_array_method

import "testing"

func TestForeachWithStructSlice(t *testing.T) {
	users := []User{{Id: 10, Name: "Go"}, {Id: 11, Name: "Lang"}}

	Foreach(users, func(user User, index int) {
		if index == 0 && user.Name != "Go" {
			t.Errorf("Expected %s received %s", "Go", user.Name)
		}
	})

}

func TestForeachWithStringSlice(t *testing.T) {
	var slices = []string{"Go", "Typescript", "Nodejs"}

	Foreach(slices, func(ln string, index int) {
		if index == 0 && ln != "Go" {
			t.Errorf("Expected %s received %s", "Go", ln)
		}
	})
}
