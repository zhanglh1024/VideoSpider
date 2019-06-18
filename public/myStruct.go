package public

import "context"

type MyContext struct {
	context.Context
	UserList    []string
	wokeList    []string
}