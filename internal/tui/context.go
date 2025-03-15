package tui

import "github.com/jackitaliano/wayfinder/internal/tui/buffer"

type Context struct {
    Mode Mode
    PendingChord string
    LastKey string
}

func NewContext(activeBuffer *buffer.Buffer) *Context {
    return &Context{
        NORMAL,
        "",
        "",
    }
}
