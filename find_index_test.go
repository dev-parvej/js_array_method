package js_array_method

import "testing"

func TestFindIndexWithCallback(t *testing.T) {
	users := []User{{Id: 10, Name: "John"}, {Id: 11, Name: "Doe"}, {Id: 12, Name: "Sabrina"}}

	index := FindIndex(users, func(user User, _ int) bool {
		return user.Name == "Doe"
	})

	if index == -1 || index == 2 {
		t.Errorf("Expected 1 got %d", index)
	}

	index = FindIndex(users, func(user User, _ int) bool {
		return user.Name == "creator"
	})

	if index != -1 {
		t.Errorf("Expected -1 got %d", index)
	}
}

func TestFindIndexWithoutCallback(t *testing.T) {
	users := []string{"User1", "User2", "User3", "User4"}

	index := FindIndex(users, "User1")

	if index != 0 {
		t.Errorf("Expected 0 got %d", index)
	}

	index = FindIndex(users, "User5")

	if index != -1 {
		t.Errorf("Expected -1 got %d", index)
	}
}
