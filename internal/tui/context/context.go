package context

import (
    "context"
)

type Context struct {
    context.Context
}

func NewContext() *Context {
    return New(context.Background())
}

func New(parent context.Context) *Context {
    return &Context{Context: parent}
}
