package tui

import (
	"fmt"
	"io"
)

type Line struct {
    Fg string
    Bg string
    Gutter string
    Content string
    Len int
}

type StatusLine struct {
    Fg string
    Bg string
    Mode Mode
    Row int
    Col int
    Content string
    LastInput byte
    LastInputKey string
    LastInputMap string
}

func (c Line) Draw(io io.Writer) {
    fmt.Fprint(io, c.Fg, c.Bg, c.Content)
}
