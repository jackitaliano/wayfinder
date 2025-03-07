package tui

import (
	"fmt"
	"io"
)

type Line struct {
    Fg string
    Bg string
    Content string
}

func (c Line) Draw(io io.Writer) {
    fmt.Fprint(io, c.Fg, c.Bg, c.Content)
}
