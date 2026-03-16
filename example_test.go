package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func Example() {
	vr := All(
		Field("Name", User.GetName, And(
			StringNotEmpty(),
			RegexpMatch(`^[a-z]+$`),
		)),
		Field("Age", User.GetAge, Range(18, 130)),
		Field("Roles", User.GetRoles, And(
			SliceNotEmpty[[]string](),
			SliceEachValue[[]string](In("admin", "member")),
		)),
	)
	fmt.Println(vr)
	for _, u := range []User{
		{Name: "alice", Age: 42, Roles: []string{"admin"}},
		{Name: "Bob", Age: 15, Roles: []string{"admin", "guest"}},
	} {
		fmt.Println("User:", u.Name)
		fmt.Println(vr.Validate(u))
	}
	// Output:
	// All(
	// 	Field("Name", And(
	// 		StringNotEmpty,
	// 		RegexpMatch("^[a-z]+$"),
	// 	)),
	// 	Field("Age", Range(18, 130)),
	// 	Field("Roles", And(
	// 		SliceNotEmpty,
	// 		SliceEachValue(In([]string{"admin", "member"})),
	// 	)),
	// )
	// User: alice
	// <nil>
	// User: Bob
	// path field "Name": "Bob" does not match regexp "^[a-z]+$"
	// path field "Age": 15 is not in the range [18, 130]
	// path field "Roles": path index 1: "guest" is not in []string{"admin", "member"}
}

type User struct {
	Name  string
	Age   int
	Roles []string
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) GetRoles() []string {
	return u.Roles
}
