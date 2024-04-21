package main

import (
	"fmt"
	"strings"

	"github.com/shortlink-org/shortlink/pkg/pattern/specification"
)

type User struct {
	ID   string
	Name string
	Age  int
}

type SpecificationError struct {
	Reason string
}

func (e *SpecificationError) Error() string {
	return fmt.Sprintf("Specification failed: %s", e.Reason)
}

type IsAdult struct{}

func (a *IsAdult) IsSatisfiedBy(user *User) error {
	if user.Age >= 18 { //nolint:revive,mnd // This is an example
		return nil
	}

	return &SpecificationError{Reason: fmt.Sprintf("User %s is not an adult. Age must be 18 or over", user.Name)}
}

type NameStartsWithA struct{}

func (n *NameStartsWithA) IsSatisfiedBy(user *User) error {
	if strings.HasPrefix(user.Name, "A") {
		return nil
	}

	return &SpecificationError{Reason: fmt.Sprintf("User %s's name does not start with 'A'", user.Name)}
}

func main() {
	// Create a list of users
	users := []*User{
		{ID: "1", Name: "Alice", Age: 25},   //nolint:revive,mnd // This is an example
		{ID: "2", Name: "Bob", Age: 30},     //nolint:revive,mnd // This is an example
		{ID: "3", Name: "Charlie", Age: 35}, //nolint:revive,mnd // This is an example
	}

	// Create a specification that checks if the user is an adult
	isAdult := &IsAdult{}
	nameStartsWithA := &NameStartsWithA{}
	spec := specification.NewAndSpecification[User](isAdult, nameStartsWithA)

	// Check if each user satisfies the specification
	for _, user := range users {
		err := spec.IsSatisfiedBy(user)
		if err != nil {
			fmt.Printf("User %s does not satisfy the specification: %v\n", user.Name, err)
		} else {
			fmt.Printf("User %s satisfies the specification\n", user.Name)
		}
	}

	// Filter the list of users using the specification
	filteredUsers, _ := specification.Filter(users, spec) //nolint:errcheck // This is an example
	fmt.Printf("Filtered users: %v\n", filteredUsers)
}
