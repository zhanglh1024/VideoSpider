package public

import "context"

type MyContext struct {
	context.Context
	UserList    []string
	wokeList    []string
}

type MyActor struct {
	Actor  string`json:"actor"`
}

type MyTitle struct {
	Title string`json:"title"`
}