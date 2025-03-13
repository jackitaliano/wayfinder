package tui

import "github.com/jackitaliano/wayfinder/internal/tui/buffer"

type Context struct {
    Mode Mode
    ActiveBuffer *buffer.Buffer
    PendingChord string
    LastKey string
}

func NewContext(activeBuffer *buffer.Buffer) *Context {
    return &Context{
        NORMAL,
        activeBuffer,
        "",
        "",
    }
}
