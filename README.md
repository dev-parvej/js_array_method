# go-js-array-methods

This package is to implement javacript array methods which is missing in golang

## Installation

```
go get github.com/dev-parvej/go-js-array-methods

```

## Usage

```
import js "github.com/dev-parvej/go-js-array-methods"

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

```

### Map

```
mapedUsers := js.Map(users, func(user User, index int) Foo {
    return Foo{
        Bar: fmt.Sprintf("%d %s", user.Id, user.Name),
    }
})

```

### Reduce 
```
numbers := []int{10, 11, 12, 13}
sum := js.Reduce(numbers, func(s int, n int, i int) int {
    return s + n
}, 0)

```

