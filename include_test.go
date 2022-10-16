package js_array_method

import (
	"testing"
)

func TestIncludeWithoutCallback(t *testing.T) {

	languages := []string{"go", "PHP", "MySql", "TypeScript"}

	if !Includes(languages, "TypeScript") {
		t.Errorf("Expected %s received %s", "true", "false")
	}

	if Includes(languages, "Hack") {
		t.Errorf("Expected %s received %s", "false", "true")
	}
}

func TestIncludeWithCallback(t *testing.T) {

	languages := []User{{Id: 12, Name: "Go"}, {Id: 14, Name: "Lang"}, {Id: 15, Name: "Typescript"}}

	if !Includes(languages, func(ln User, _ int) bool { return ln.Name == "Go" }) {
		t.Errorf("Expected %s received %s", "true", "false")
	}

	if Includes(languages, func(ln User, _ int) bool { return ln.Name == "MySql" }) {
		t.Errorf("Expected %s received %s", "false", "true")
	}
}
