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
// {{Id: 10, Name: "Go"}}

```

### Map

```
js.Map(users, func(user User, index int) Foo {
    return Foo{
        Bar: fmt.Sprintf("%d %s", user.Id, user.Name),
    }
})
// {{Bar: "10 Go"}, {Bar: "11 Lang"}}

```

### Reduce 
```
numbers := []int{10, 11, 12, 13}
js.Reduce(numbers, func(s int, n int, i int) int {
    return s + n
}, 0)

// 46
```

### Every

```
users := []User{{Id: 12, Name: "Go"}, {Id: 14, Name: "Go"}}
resultTrue := Every(users, func(user User, _ int) bool {
    return user.Name == "Go"
})
// true

Every([]string{"hey", "hi"}, "hi")
// false
```

### Foreach
```
var slices = []string{"Go", "Typescript", "Nodejs"}

Foreach(slices, func(ln string, index int) {
    // ln contains the item of the slice
    // index contains item index
})
```

### Includes

```
languages := []User{{Id: 12, Name: "Go"}, {Id: 14, Name: "Lang"}, {Id: 15, Name: "Typescript"}}

Includes(languages, func(ln User, _ int) bool { return ln.Name == "Go" })
//true
Includes(languages, func(ln User, _ int) bool { return ln.Name == "MySql" })
//false

languages := []string{"go", "PHP", "MySql", "TypeScript"}
Includes(languages, "TypeScript").
//true

```
### Includes

```
users := []User{{Id: 10, Name: "John"}, {Id: 11, Name: "Doe"}, {Id: 12, Name: "Sabrina"}}
Reverse(users)
// {{Id: 12, Name: "Sabrina"}, {Id: 11, Name: "Doe"}, {Id: 10, Name: "John"}}
```