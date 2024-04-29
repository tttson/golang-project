package graph

import (
	"context"

	"github.com/tttson/golang-project/server/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}

// In the generated schema.resolvers.go file, the todoResolver type
// is automatically defined by gqlgen based on the
// resolver interface you've defined in this resolver.go file.
// The todoResolver type is a struct that implements
// the resolver interface methods for the Todo type.

// When you define a resolver method in your resolver interface for a specific field,
// such as User for the Todo type, gqlgen automatically generates the corresponding
// resolver implementation in the schema.resolvers.go file.
// The resolver interface should include methods for resolving the fields of the Todo type,
// including the User field if needed.
type TodoResolver interface {
	User(ctx context.Context, obj *model.Todo) (*model.User, error)
}

// User is a method of the Resolver struct that resolves the User field of the Todo type.
// It receives the context.Context and a pointer to the Todo object as arguments.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	// Implement logic to resolve the User field of the Todo type
	// Here you can fetch the user data based on the UserID field of the Todo
	// For example:
	// user, err := fetchUserByID(obj.UserID)
	// if err != nil {
	//     return nil, err
	// }
	// return user, nil
	return &model.User{ID: obj.UserID, Name: "placeholder name" + obj.UserID}, nil
}
