package server

import "context"

type Options struct {
	Folder    string
	Recursive bool
	Port      int
	Host      string
	Ctx       context.Context
}
