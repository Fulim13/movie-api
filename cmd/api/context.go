package main

import (
	"context"
	"net/http"

	"github.com/Fulim13/movie-api/internal/data"
)

// Every http.Request that our application processes
// has a context.Context embedded in it,
// which we can use to store key-value pairs
// containing arbitrary data during the lifetime of the request.
// In this case we want to store a User struct
// containing the current user’s information.

// Any values stored in the request context
// have the type any.
// This means that after retrieving a value from the request context
// you need to assert it back to its original type before using it.

// It’s good practice to use your own custom type
// for the request context keys.
// This helps prevent naming collisions
// between your code and any third-party packages
// which are also using the request context to store information.

type contextKey string

const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(context.Background(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetuser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
