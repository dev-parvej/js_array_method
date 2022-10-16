# go-js-array-methods

This package is to implement javacript array methods which is missing in golang

## Installation

```
go get github.com/dev-parvej/js_array_method
```

## Usage

```
import js "github.com/dev-parvej/js_array_method"

type User struct {
	Id   int
	Name string
}

type Foo struct {
	Bar string
}

users := []User{{Id: 10, Name: "Go"}, {Id: 11, Name: "Lang"}}

```

### Filter

```
filteredUsers := js.Filter(users, func(user User, index int) bool {
    return user.Id == 10
})
// filteredUsers
{{Id: 10, Name: "Go"}}

```

### Map

```
mapedUsers := js.Map(users, func(user User, index int) Foo {
    return Foo{
        Bar: fmt.Sprintf("%d %s", user.Id, user.Name),
    }
})
// mapedUsers
{{Bar: "10 Go"}, {Bar: "11 Lang"}}

```

### Reduce 
```
numbers := []int{10, 11, 12, 13}
sum := js.Reduce(numbers, func(s int, n int, i int) int {
    return s + n
}, 0)

// sum
46
```

### Every

```
users := []User{{Id: 12, Name: "Go"}, {Id: 14, Name: "Go"}}
resultTrue := Every(users, func(user User, _ int) bool {
    return user.Name == "Go"
})
// resultTrue
true

resultFalse := Every([]string{"hey", "hi"}, "hi")
// resultFalse
false
```

### Foreach
```
var slices = []string{"Go", "Typescript", "Nodejs"}

Foreach(slices, func(ln string, index int) {
    // ln contains the item of the slice
    // index contains item index
})
```

