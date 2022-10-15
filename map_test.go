package js_array_method

import (
	"fmt"
	"testing"
)

type User struct {
	Id   int
	Name string
}

type Foo struct {
	Bar string
}

func TestMapWithStructSlice(t *testing.T) {
	users := []User{{Id: 10, Name: "Go"}, {Id: 11, Name: "Lang"}}

	mapedUsers := Map(users, func(user User, index int) Foo {
		return Foo{
			Bar: fmt.Sprintf("%d %s", user.Id, user.Name),
		}
	})

	expected := "10 Go"
	if mapedUsers[0].Bar != expected {
		t.Errorf("Expected %s received %s", expected, mapedUsers[0].Bar)
	}
	expected = "11 Lang"
	if mapedUsers[1].Bar != expected {
		t.Errorf("Expected %s received %s", expected, mapedUsers[1].Bar)
	}
}

func TestMapWithStringSlice(t *testing.T) {
	var slices = []string{"Go", "Typescript", "Nodejs"}

	mappedSlice := Map(slices, func(ln string, index int) string {
		return fmt.Sprintf("%d/%s", index, ln)
	})

	expected := "0/Go"

	if mappedSlice[0] != expected {
		t.Errorf("Expected %s received %s", expected, mappedSlice[0])
	}

	expected = "1/Typescript"

	if mappedSlice[1] != expected {
		t.Errorf("Expected %s received %s", expected, mappedSlice[1])
	}
}
